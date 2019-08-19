package categorydb

import (
	"context"
	"math"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/categorydb/category"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/database"
	"github.com/PGo-Projects/tasky/internal/taskdb"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
	"github.com/PGo-Projects/webauth"
	"github.com/spf13/viper"
)

var (
	dbClient = database.MustMongoClient(context.TODO(), "mongodb://localhost:27017")
)

func GetCategory(username, category string) ([]task.Task, error) {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	f := &task.Task{
		Username: username,
		Category: category,
	}
	// TODO: Ignore the index guard in the initial query
	cursor, err := tasks.Find(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	var tasksFound []task.Task
	for cursor.Next(context.TODO()) {
		var t task.Task
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}

		// If not the index guard, append it
		if t.Index != -1 {
			tasksFound = append(tasksFound, t)
		}
	}

	return tasksFound, nil
}

func GetOrderedCategory(username, category string) ([]task.Task, error) {
	tasks, err := GetCategory(username, category)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return tasks, nil
	}

	var firstTask task.Task
	indexMap := make(map[int64]int)
	for index, t := range tasks {
		indexMap[t.Index] = index
		if t.Predecessor == -1 {
			firstTask = t
		}
	}

	var orderedTasks []task.Task
	orderedTasks = append(orderedTasks, firstTask)

	nextTask := tasks[indexMap[firstTask.Successor]]
	for nextTask.Successor != -1 {
		orderedTasks = append(orderedTasks, nextTask)
		nextTask = tasks[indexMap[nextTask.Successor]]
	}

	if firstTask != nextTask {
		orderedTasks = append(orderedTasks, nextTask)
	}

	return orderedTasks, nil
}

func GetFirst(username, category string) (*task.Task, error) {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	filter := task.Task{
		Username:    username,
		Category:    category,
		Predecessor: -1,
	}

	var t task.Task
	err := tasks.FindOne(context.TODO(), filter).Decode(&t)
	return &t, err
}

func GetLast(username, category string) (*task.Task, error) {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	filter := task.Task{
		Username:  username,
		Category:  category,
		Successor: -1,
	}

	var t task.Task
	err := tasks.FindOne(context.TODO(), filter).Decode(&t)
	return &t, err
}

func Update(username, taskCategory string) error {
	if current := category.GetChronological(taskCategory); current != nil {
		if current.HasPrevious() {
			previous := current.Previous()
			update(username, taskCategory, previous, current)
		}
		if current.HasNext() {
			next := current.Next()
			update(username, taskCategory, current, next)
		}
	}
	return nil
}

func update(username, category string, previous, current category.Category) error {
	tasks, err := GetOrderedCategory(username, current.Name())
	if err != nil {
		return err
	}

	previousLast, err := GetLast(username, previous.Name())
	if err != nil {
		return err
	}

	// Update the successor of the last task in the database for given category
	if len(tasks) > 0 && previous.Valid(tasks[0]) {
		filter := &task.Task{
			Username: previousLast.Username,
			Index:    previousLast.Index,
		}
		previousLast.Successor = tasks[0].Index
		if err := taskdb.ReplaceTask(filter, previousLast); err != nil {
			return err
		}
	}

	for _, elem := range tasks {
		if previous.Valid(elem) {
			if err := taskdb.DeleteTask(&elem); err != nil {
				return err
			}

			elem.Predecessor = previousLast.Index
			elem.Successor = -1
			if err := taskdb.InsertTask(&elem); err != nil {
				return err
			}

			previousLast = &elem
		}
	}
	return nil
}

func AddIndexGuardHook(w http.ResponseWriter, r *http.Request, c webauth.Credentials) error {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	for _, taskCategory := range category.List {
		_, err := tasks.InsertOne(context.TODO(), &task.Task{
			Username:    c.Username,
			Category:    taskCategory.Name(),
			Index:       -1,
			Predecessor: math.MinInt64,
			Successor:   -1,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
