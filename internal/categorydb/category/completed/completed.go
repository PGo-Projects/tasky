package completed

import (
	"github.com/PGo-Projects/tasky/internal/categorydb"
	"github.com/PGo-Projects/tasky/internal/categorydb/category"
	"github.com/PGo-Projects/tasky/internal/taskdb"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
)

func init() {
	category.RegisterCategories([]category.Category{
		&completed{},
	})
}

type completed struct {
}

func Get() *completed {
	return &completed{}
}

func (c *completed) Name() string {
	return "completed"
}

func (c *completed) HasPrevious() bool {
	return false
}

func (c *completed) HasNext() bool {
	return false
}

func (c *completed) Previous() category.Category {
	return nil
}

func (c *completed) Next() category.Category {
	return nil
}

func (c *completed) Valid(t task.Task) bool {
	return true
}

func Mark(t *task.Task) error {
	completed := Get()
	completedFirst, err := categorydb.GetFirst(t.Username, completed.Name())
	if err != nil {
		f := *t
		f.Category = t.OldCategory

		t.OldCategory = t.Category
		t.Category = completed.Name()
		t.Predecessor = -1
		t.Successor = -1
		return taskdb.UpdateTask(&f, t)
	} else {
		f := *t
		f.Category = t.OldCategory

		t.Predecessor = -1
		t.Successor = completedFirst.Index
		t.Category = completed.Name()
		err := taskdb.UpdateTask(&f, t)
		if err != nil {
			return err
		}

		f = *completedFirst
		completedFirst.Predecessor = t.Index
		return taskdb.ReplaceTask(&f, completedFirst)
	}
}

func Unmark(t *task.Task) error {
	completed := Get()
	oldT, err := taskdb.GetTaskWithCategory(completed.Name(), t.Username, t.Index)
	if err != nil {
		return err
	}
	return taskdb.UpdateTask(oldT, t)
}
