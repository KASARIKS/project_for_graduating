package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func Task(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, errors.New("wrong request format").Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		taskPost(w, r)
	}
}

func taskPost(w http.ResponseWriter, r *http.Request) {
	task, err := getTaskFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send id into repsonse in json
	if err := json.NewEncoder(w).Encode(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTaskFromRequest(r *http.Request) (*dbtask.DbTask, error) {
	var task dbtask.DbTask

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return nil, err
	}

	if task.Title == "" {
		return nil, errors.New("empty title")
	}

	return filterTaskDate(&task)
}

func filterTaskDate(task *dbtask.DbTask) (*dbtask.DbTask, error) {
	// Validate task date
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(nextdate.DateFormat)
	}
	t, err := time.Parse(nextdate.DateFormat, task.Date)
	if err != nil {
		return nil, errors.New("incorrect date")
	}
	if task.Repeat != "" {
		next, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return nil, err
		}

		if now.After(t) {
			if len(task.Repeat) == 0 {
				task.Date = now.Format(nextdate.DateFormat)
			} else {
				task.Date = next
			}
		}
	}

	return task, nil
}
