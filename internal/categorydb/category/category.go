package category

import (
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
)

var (
	List []Category
)

type Category interface {
	Name() string
	HasPrevious() bool
	HasNext() bool
	Previous() Category
	Next() Category
	Valid(t task.Task) bool
}

func RegisterCategories(categories []Category) {
	List = append(List, categories...)
}
