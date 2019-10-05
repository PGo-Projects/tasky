package views

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/tasky/internal/categorydb"
	"github.com/PGo-Projects/tasky/internal/categorydb/category/completed"
	"github.com/PGo-Projects/tasky/internal/taskdb"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
	"github.com/PGo-Projects/webauth"
	response "github.com/PGo-Projects/webresponse"
	"github.com/go-chi/chi"
)

type taskOp int

const (
	successCategoryMessage = "We have processed the category successfully!"
	successTaskMessage     = "We have processed the task successfully!"

	insert taskOp = iota
	update
	delete
)

func RegisterTaskyEndpoints(mux *chi.Mux) {
	mux.Get("/get_category/{category}", getCategoryHandler)
	mux.Post("/update_category/{category}", updateCategoryHandler)
	mux.Post("/is_completed/{mark}", isCompletedHandler)

	mux.Post("/insert_task", insertTaskHandler)
	mux.Post("/update_task", updateTaskHandler)
	mux.Post("/delete_task", deleteTaskHandler)
}

func getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	username, isLoggedIn := webauth.IsLoggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, webauth.LoginRoute, http.StatusFound)
		return
	}

	var responseJSON []byte
	taskCategory := chi.URLParam(r, "category")
	tasks, err := categorydb.GetOrderedCategory(username, taskCategory)
	if err != nil {
		responseJSON = response.Error(response.ErrInternalServer)
	} else {
		jsonTasks, err := json.Marshal(tasks)
		if err != nil {
			responseJSON = response.Error(response.ErrInternalServer)
		} else {
			responseJSON, err = json.Marshal(map[string]string{
				"status":     successCategoryMessage,
				"statusType": response.StatusSuccess,
				"tasks":      string(jsonTasks),
			})
			if err != nil {
				responseJSON = response.Error(response.ErrInternalServer)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	username, isLoggedIn := webauth.IsLoggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, webauth.LoginRoute, http.StatusFound)
		return
	}

	var responseJSON []byte
	taskCategory := chi.URLParam(r, "category")
	err := categorydb.Update(username, taskCategory)
	if err != nil {
		responseJSON = response.Error(response.ErrInternalServer)
	} else {
		responseJSON = response.Success(successCategoryMessage)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func isCompletedHandler(w http.ResponseWriter, r *http.Request) {
	username, isLoggedIn := webauth.IsLoggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, webauth.LoginRoute, http.StatusFound)
		return
	}

	var responseJSON []byte
	t, err := task.DecodeFromRequest(r)
	if err != nil {
		output.Errorln(err)
		responseJSON = response.Error(response.ErrInternalServer)
	} else {
		t.Username = username

		var err error
		isCompleted := chi.URLParam(r, "mark")
		if isCompleted == "true" {
			err = completed.Mark(t)
		} else if isCompleted == "false" {
			err = completed.Unmark(t)
		} else {
			err = errors.New("Not a valid option")
		}
		if err != nil {
			output.Errorln(err)
			responseJSON = response.Error(response.ErrInternalServer)
		} else {
			responseJSON = response.Success(successTaskMessage)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func insertTaskHandler(w http.ResponseWriter, r *http.Request) {
	handleGenericTask(w, r, insert)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	handleGenericTask(w, r, update)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	handleGenericTask(w, r, delete)
}

func handleGenericTask(w http.ResponseWriter, r *http.Request, op taskOp) {
	username, isLoggedIn := webauth.IsLoggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, webauth.LoginRoute, http.StatusFound)
		return
	}

	var responseJSON []byte
	t, err := task.DecodeFromRequest(r)
	if err != nil {
		output.Errorln(err)
		responseJSON = response.Error(response.ErrInternalServer)
	} else {
		t.Username = username
		filter := &task.Task{
			Index:    t.Index,
			Username: t.Username,
			Category: t.OldCategory,
		}
		data, err := executeTaskOp(op, filter, t)
		if err != nil {
			output.Errorln(err)
			responseJSON = response.Error(response.ErrBadRequest)
		} else {
			data["status"] = successTaskMessage
			data["statusType"] = response.StatusSuccess
			responseJSON = response.General(data)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func executeTaskOp(op taskOp, filter *task.Task, t *task.Task) (map[string]string, error) {
	data := make(map[string]string)
	switch op {
	case insert:
		err := taskdb.InsertTask(t)
		if err != nil {
			return data, err
		}
		data["index"] = strconv.FormatInt(t.Index, 10)
		return data, nil
	case update:
		return data, taskdb.UpdateTask(filter, t)
	case delete:
		return data, taskdb.DeleteTask(t)
	default:
		return data, errors.New("Not a valid option!")
	}
}
