package taskdb

import (
	"context"
	"errors"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/database"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
	"github.com/PGo-Projects/tasky/internal/userdb"
	"github.com/PGo-Projects/tasky/internal/userdb/user"
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

func GetTaskWithCategory(category, username string, index int64) (*task.Task, error) {
	dbName := viper.GetString(config.DBName)
	tasks := dbClient.Database(dbName).Collection("tasks")

	filter := task.Task{
		Username: username,
		Category: category,
		Index:    index,
	}
	var t task.Task
	err := tasks.FindOne(context.TODO(), filter).Decode(&t)
	return &t, err
}

func UpdateTask(f *task.Task, t *task.Task) error {
	f, err := GetTaskWithCategory(f.Category, f.Username, f.Index)
	if err != nil {
		return err
	}
	predecessor, err := GetTaskWithCategory(f.Category, t.Username, f.Predecessor)
	if err != nil {
		return err
	}
	successor, err := GetTaskWithCategory(f.Category, t.Username, f.Successor)
	if err != nil {
		return err
	}
	if err := linkTasks(predecessor, successor); err != nil {
		return err
	}

	newPredecessor, err := GetTaskWithCategory(t.Category, t.Username, t.Predecessor)
	if err != nil {
		return err
	}
	newSuccessor, err := GetTaskWithCategory(t.Category, t.Username, t.Successor)
	if err != nil {
		return err
	}
	if linkTasks(newPredecessor, t) != nil || linkTasks(t, newSuccessor) != nil {
		return errors.New("Linking failed")
	}

	return ReplaceTask(f, t)
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

	predecessor, err := GetTaskWithCategory(t.Category, t.Username, t.Predecessor)
	if err != nil {
		return err
	}
	successor, err := GetTaskWithCategory(t.Category, t.Username, t.Successor)
	if err != nil {
		return err
	}
	return linkTasks(predecessor, successor)
}

func attemptToInsert(t *task.Task) bool {
	predecessor, err := GetTaskWithCategory(t.Category, t.Username, t.Predecessor)
	if err != nil {
		return false
	}
	successor, err := GetTaskWithCategory(t.Category, t.Username, t.Successor)
	if err != nil {
		return false
	}
	return linkTasks(predecessor, t) == nil && linkTasks(t, successor) == nil
}

func linkTasks(p *task.Task, s *task.Task) error {
	pf := *p
	sf := *s

	if s.Index == -1 {
		// s refers to the index guard, so just update the predecessor and update
		// the database
		if p.SetSuccessor(s) != nil {
			return errors.New("Linking failed!")
		}
		return ReplaceTask(&pf, p)
	} else if p.Index == s.Index {
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
