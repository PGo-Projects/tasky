package category

import (
	"time"

	"github.com/PGo-Projects/tasky/internal/taskdb/task"
)

func init() {
	RegisterCategories([]Category{
		&incomplete{},
		&today{},
		&upcoming{},
		&longTerm{},
		&completed{},
	})
}

func GetChronological(category string) Category {
	switch category {
	case "incomplete":
		return &incomplete{}
	case "today":
		return &today{}
	case "upcoming":
		return &upcoming{}
	case "longTerm":
		return &longTerm{}
	case "completed":
		return &completed{}
	default:
		return nil
	}
}

type incomplete struct {
}

func (i *incomplete) Name() string {
	return "incomplete"
}

func (i *incomplete) HasPrevious() bool {
	return false
}

func (i *incomplete) HasNext() bool {
	return true
}

func (i *incomplete) Previous() Category {
	return nil
}

func (i *incomplete) Next() Category {
	return &today{}
}

func (i *incomplete) Valid(t task.Task) bool {
	bound := time.Now()
	boundTask := &task.Task{
		Date: bound.Format("2006-01-02"),
		Time: bound.Format("15:04"),
	}

	return t.DateTime().IsBefore(boundTask.DateTime())
}

type today struct {
}

func (t *today) Name() string {
	return "today"
}

func (t *today) HasPrevious() bool {
	return true
}

func (t *today) HasNext() bool {
	return true
}

func (t *today) Previous() Category {
	return &incomplete{}
}

func (t *today) Next() Category {
	return &upcoming{}
}

func (to *today) Valid(t task.Task) bool {
	now := time.Now()
	nowTask := &task.Task{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04"),
	}

	boundTask := &task.Task{
		Date: now.Format("2006-01-02"),
		Time: "23:59",
	}

	return t.DateTime().IsAfterOrEqual(nowTask.DateTime()) &&
		t.DateTime().IsBeforeOrEqual(boundTask.DateTime())
}

type upcoming struct {
}

func (u *upcoming) Name() string {
	return "upcoming"
}

func (u *upcoming) HasPrevious() bool {
	return true
}

func (u *upcoming) HasNext() bool {
	return true
}

func (u *upcoming) Previous() Category {
	return &today{}
}

func (u *upcoming) Next() Category {
	return &longTerm{}
}

func (u *upcoming) Valid(t task.Task) bool {
	now := time.Now()
	nowTask := &task.Task{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04"),
	}

	bound := now.AddDate(0, 0, 7)
	boundTask := &task.Task{
		Date: bound.Format("2006-01-02"),
		Time: "23:59",
	}

	return t.DateTime().IsAfterOrEqual(nowTask.DateTime()) &&
		t.DateTime().IsBeforeOrEqual(boundTask.DateTime())
}

type longTerm struct {
}

func (lt *longTerm) Name() string {
	return "longTerm"
}

func (lt *longTerm) HasPrevious() bool {
	return true
}

func (lt *longTerm) HasNext() bool {
	return false
}

func (lt *longTerm) Previous() Category {
	return &upcoming{}
}

func (lt *longTerm) Next() Category {
	return nil
}

func (lt *longTerm) Valid(t task.Task) bool {
	bound := time.Now().AddDate(0, 0, 7)
	boundTask := &task.Task{
		Date: bound.Format("2006-01-02"),
		Time: "23:59",
	}

	return t.DateTime().IsAfter(boundTask.DateTime())

}

type completed struct {
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

func (c *completed) Previous() Category {
	return nil
}

func (c *completed) Next() Category {
	return nil
}

func (c *completed) Valid(t task.Task) bool {
	return true
}
