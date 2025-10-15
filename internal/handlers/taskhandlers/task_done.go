package taskhandlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func TaskDone(w http.ResponseWriter, r *http.Request) {
	id, err := getIdentifier(r)
	if err != nil {
		writeErrorInJson(w, err, http.StatusBadRequest)
		return
	}

	if err := writeTask(id); err != nil {
		writeErrorInJson(w, err, http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		writeErrorInJson(w, err, http.StatusInternalServerError)
		return
	}
}

func writeTask(id int) error {
	task, err := db.GetTaskById(id)
	if err != nil {
		return err
	}

	return changeTaskDateOrDelete(task, id)
}

func changeTaskDateOrDelete(task *dbtask.Task, id int) error {
	if len(task.Repeat) == 0 {
		if err := db.DeleteTask(id); err != nil {
			return err
		}
	} else {
		now := time.Now()

		next, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		task.Date = next

		if err := db.UpdateTask(task); err != nil {
			return err
		}
	}

	return nil
}
