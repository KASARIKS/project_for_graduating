package taskhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func TaskDone(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErrorInJson(w, errors.New("no identifier"))
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	task, err := db.GetTaskById(intId)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if len(task.Repeat) == 0 {
		if err := db.DeleteTask(intId); err != nil {
			writeErrorInJson(w, err)
			return
		}
	} else {
		now := time.Now()

		next, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeErrorInJson(w, err)
			return
		}

		task.Date = next

		if err := db.UpdateTask(task); err != nil {
			writeErrorInJson(w, err)
			return
		}
	}

	if err = json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		writeErrorInJson(w, err)
		return
	}
}
