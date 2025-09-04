package taskhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func taskPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	task, err := getTaskFromRequest(r)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	id, err := db.AddTask(task)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	// Send id into repsonse in json
	if err = json.NewEncoder(w).Encode(map[string]int64{"id": id}); err != nil {
		writeErrorInJson(w, err)
		return
	}
}

func getTaskFromRequest(r *http.Request) (*dbtask.Task, error) {
	var task dbtask.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return nil, err
	}

	if len(task.Title) == 0 {
		return nil, errors.New("empty title")
	}

	return filterTaskDate(&task)
}

func filterTaskDate(task *dbtask.Task) (*dbtask.Task, error) {
	// Validate task date
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(nextdate.DateFormat)
	}

	t, err := time.Parse(nextdate.DateFormat, task.Date)
	if err != nil {
		return nil, errors.New("incorrect date")
	}

	if now.After(t) {
		if len(task.Repeat) == 0 || task.Repeat == "d 1" || task.Date == now.Format(nextdate.DateFormat) {
			task.Date = now.Format(nextdate.DateFormat)
		} else {
			next, err := nextdate.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return nil, err
			}

			task.Date = next
		}
	}

	return task, nil
}
