package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func Task(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		taskPost(w, r)
	}
}

func taskPost(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	byteId, err := json.Marshal(map[string]int64{"id": id})
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteId)
}

func getTaskFromRequest(r *http.Request) (*dbtask.DbTask, error) {
	var task dbtask.DbTask

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf.Bytes(), &task); err != nil {
		return nil, err
	}

	if len(task.Title) == 0 {
		return nil, errors.New("empty title")
	}

	return filterTaskDate(&task)
}

func filterTaskDate(task *dbtask.DbTask) (*dbtask.DbTask, error) {
	// Validate task date
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(nextdate.DateFormat)
	}
	t, err := time.Parse(nextdate.DateFormat, task.Date)
	if err != nil {
		return nil, errors.New("incorrect date")
	}

	// Date cannot be smaller than today's date
	if now.After(t) {
		task.Date = now.Format(nextdate.DateFormat)
	}

	if len(task.Repeat) != 0 {
		next, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return nil, err
		}

		if now.After(t) {
			if len(task.Repeat) == 0 || task.Repeat == "d 1" {
				task.Date = now.Format(nextdate.DateFormat)
			} else {
				task.Date = next
			}
		} else {
			task.Date = now.Format(nextdate.DateFormat)
		}
	}

	return task, nil
}
