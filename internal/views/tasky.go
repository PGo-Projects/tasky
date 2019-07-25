package views

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/taskdb"
	"github.com/PGo-Projects/tasky/internal/taskdb/task"
	"github.com/PGo-Projects/webauth"
	response "github.com/PGo-Projects/webresponse"
	"github.com/go-chi/chi"
	"github.com/lpar/gzipped"
	"github.com/spf13/viper"
)

type taskOp int

const (
	successMessage = "We have processed the task successfully!"

	insert taskOp = iota
	update
	delete
)

func RegisterTaskyEndpoints(mux *chi.Mux) {
	mux.Get("/today", taskyHandler)
	mux.Get("/upcoming", taskyHandler)
	mux.Get("/long_term", taskyHandler)
	mux.Get("/incomplete", taskyHandler)
	mux.Get("/completed", taskyHandler)
	mux.Get("/thought_cloud", taskyHandler)

	mux.Post("/insert_task", insertTaskHandler)
	mux.Post("/update_task", updateTaskHandler)
	mux.Post("/delete_task", deleteTaskHandler)
}

func taskyHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/tasky.html"
	webAssetsDirectory := http.Dir(viper.GetString(config.WebAssetsPathKey))
	gzipped.FileServer(webAssetsDirectory).ServeHTTP(w, r)
}

func insertTaskHandler(w http.ResponseWriter, r *http.Request) {
	t, err := handleGenericTask(w, r, insert)
	if err == response.ErrForbidden {
		return
	}

	var responseJSON []byte
	if err != nil {
		responseJSON = response.Error(err)
	} else {
		responseJSON = response.General(map[string]string{
			"status":     successMessage,
			"statusType": response.StatusSuccess,
			"index":      strconv.FormatInt(t.Index, 10),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	_, err := handleGenericTask(w, r, update)
	if err == response.ErrForbidden {
		return
	}

	var responseJSON []byte
	if err != nil {
		responseJSON = response.Error(err)
	} else {
		responseJSON = response.Status(successMessage, response.StatusSuccess)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	_, err := handleGenericTask(w, r, delete)
	if err == response.ErrForbidden {
		return
	}

	var responseJSON []byte
	if err != nil {
		responseJSON = response.Error(err)
	} else {
		responseJSON = response.Status(successMessage, response.StatusSuccess)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func handleGenericTask(w http.ResponseWriter, r *http.Request, op taskOp) (*task.Task, error) {
	username, isLoggedIn := webauth.IsLoggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, response.ErrInternalServer
	}

	t, err := task.DecodeFromRequest(r)
	if err != nil {
		output.Errorln(err)
		return nil, response.ErrInternalServer
	} else {
		t.Username = username
		filter := &task.Task{
			Index:    t.Index,
			Username: t.Username,
		}
		if err := executeTaskOp(op, filter, t); err != nil {
			output.Errorln(err)
			return nil, response.ErrBadRequest
		} else {
			return t, response.Successful
		}
	}
}

func executeTaskOp(op taskOp, filter *task.Task, t *task.Task) error {
	switch op {
	case insert:
		return taskdb.InsertTask(t)
	case update:
		current, err := taskdb.GetTask(filter.Username, filter.Index)
		if err != nil {
			return err
		}
		t.Predecessor = current.Predecessor
		t.Successor = current.Successor
		return taskdb.ReplaceTask(filter, t)
	case delete:
		return taskdb.DeleteTask(t)
	default:
		return errors.New("Not a valid option!")
	}
}
