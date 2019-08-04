package taskdb

import (
	"context"
	"errors"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/database"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
	"github.com/PGo-Projects/tasky/internal/userdb"
	"github.com/PGo-Projects/tasky/internal/userdb/user"
	"github.com/PGo-Projects/webauth"
	"github.com/spf13/viper"
)

var (
	dbClient = database.MustMongoClient(context.TODO(), "mongodb://localhost:27017")
)

func GetTask(username string, index int64) (*task.Task, error) {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	filter := task.Task{
		Username: username,
		Index:    index,
	}
	var t task.Task
	err := tasks.FindOne(context.TODO(), filter).Decode(&t)
	return &t, err
}

func ReplaceTask(filter *task.Task, t *task.Task) error {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	result := tasks.FindOneAndReplace(context.TODO(), filter, t)
	return result.Err()
}

func InsertTask(t *task.Task) error {
	u, err := userdb.GetUser(t.Username)
	if err != nil {
		return err
	}

	t.Index = u.PopNextUnusedIndex()
	if !attemptToInsert(t) {
		return errors.New("The task is malformed.")
	}

	filter := &user.User{
		Username: t.Username,
	}
	if err := userdb.ReplaceUser(filter, u); err != nil {
		return err
	}

	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")
	_, err = tasks.InsertOne(context.TODO(), t)
	return err
}

func DeleteTask(t *task.Task) error {
	// The index guard task should not be deletable
	if t.Index == -1 {
		return errors.New("Task cannot be found.")
	}

	// Restore missing values from the database
	// (in particular we need predecessor and successor)
	t, err := GetTask(t.Username, t.Index)
	if err != nil {
		return err
	}

	u, err := userdb.GetUser(t.Username)
	if err != nil {
		return err
	}

	u.InsertUnusedIndex(t.Index)
	filter := &user.User{
		Username: t.Username,
	}
	if err := userdb.ReplaceUser(filter, u); err != nil {
		return err
	}

	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")
	result, err := tasks.DeleteOne(context.TODO(), t)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("Task could not be found and so was not deleted.")
	}

	predecessor, err := GetTask(t.Username, t.Predecessor)
	if err != nil {
		return err
	}
	successor, err := GetTask(t.Username, t.Successor)
	if err != nil {
		return err
	}
	return linkTasks(predecessor, successor)
}

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

	var tasks_found []task.Task
	for cursor.Next(context.TODO()) {
		var t task.Task
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}

		// If not the index guard, append it
		if t.Index != -1 {
			tasks_found = append(tasks_found, t)
		}
	}

	return tasks_found, nil
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

	var ordered_tasks []task.Task
	ordered_tasks = append(ordered_tasks, firstTask)

	nextTask := tasks[indexMap[firstTask.Successor]]
	for nextTask.Successor != -1 {
		ordered_tasks = append(ordered_tasks, nextTask)
		nextTask = tasks[indexMap[nextTask.Successor]]
	}
	return ordered_tasks, nil
}

func attemptToInsert(t *task.Task) bool {
	predecessor, err := GetTask(t.Username, t.Predecessor)
	if err != nil {
		return false
	}
	successor, err := GetTask(t.Username, t.Successor)
	if err != nil {
		return false
	}
	return linkTasks(predecessor, t) == nil && linkTasks(t, successor) == nil
}

func linkTasks(p *task.Task, s *task.Task) error {
	pf := *p
	sf := *s

	if p.Index == s.Index {
		// p and s references the same task, so just update the predecessor and
		// successor for one of them and update the database
		if p.SetPredecessor(p) != nil || p.SetSuccessor(p) != nil {
			return errors.New("Linking failed!")
		}
		return ReplaceTask(&pf, p)
	} else {
		if p.SetSuccessor(s) != nil || s.SetPredecessor(p) != nil {
			return errors.New("Linked failed!")
		}
		if err := ReplaceTask(&pf, p); err != nil {
			return err
		}
		return ReplaceTask(&sf, s)
	}
}

func AddIndexGuardHook(w http.ResponseWriter, r *http.Request, c webauth.Credentials) error {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	_, err := tasks.InsertOne(context.TODO(), &task.Task{
		Username:    c.Username,
		Index:       -1,
		Predecessor: -1,
		Successor:   -1,
	})
	return err
}
