package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Name        string `bson:"name,omitempty" json:"name"`
	Date        string `bson:"date,omitempty" json:"date"`
	Time        string `bson:"time,omitempty" json:"time"`
	Description string `bson:"description,omitempty" json:"description"`
	Category    string `bson:"category,omitempty" json:"category"`

	Username    string `bson:"username" json:"-"`
	Index       int64  `bson:"index,omitempty" json:"index"`
	Predecessor int64  `bson:"predecessor,omitempty" json:"-"`
	Successor   int64  `bson:"successor,omitempty" json:"-"`
}

func (t *Task) DateTime() *dateTime {
	dt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%s:00Z", t.Date, t.Time))
	return &dateTime{
		value: dt,
		err:   err,
	}
}

func (t *Task) HasPredecessor() bool {
	return t.Predecessor != -1
}

func (t *Task) HasSuccessor() bool {
	return t.Successor != -1
}

func (t *Task) SetPredecessor(p *Task) error {
	predecessorDateTime := p.DateTime()
	// Index of -1 means a guard index is involved which is a special case
	if t.Index == -1 || p.Index == -1 ||
		predecessorDateTime.IsBeforeOrEqual(t.DateTime()) {
		t.Predecessor = p.Index
		return nil
	}
	return errors.New("Not a valid predecessor!")
}

func (t *Task) SetSuccessor(s *Task) error {
	successorDateTime := s.DateTime()
	// Index of -1 means a guard index is involved which is a special case
	if t.Index == -1 || s.Index == -1 || successorDateTime.IsAfterOrEqual(t.DateTime()) {
		t.Successor = s.Index
		return nil
	}
	return errors.New("Not a valid successor!")
}

type dateTime struct {
	value time.Time
	err   error
}

func (dt *dateTime) IsAfterOrEqual(bound *dateTime) bool {
	return dt.err == nil && bound.err == nil &&
		(dt.value.After(bound.value) || dt.value.Equal(bound.value))
}

func (dt *dateTime) IsBeforeOrEqual(bound *dateTime) bool {
	return dt.err == nil && bound.err == nil &&
		(dt.value.Before(bound.value) || dt.value.Equal(bound.value))
}

func DecodeFromRequest(r *http.Request) (*Task, error) {
	var task Task

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
