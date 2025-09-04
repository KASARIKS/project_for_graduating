package taskhandlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func taskPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var dbTask dbtask.Task

	// Get task from request
	if err := json.NewDecoder(r.Body).Decode(&dbTask); err != nil {
		writeErrorInJson(w, err, http.StatusInternalServerError)
		return
	}

	// Validate task
	if err := validateTask(dbTask); err != nil {
		writeErrorInJson(w, err, http.StatusBadRequest)
		return
	}

	// NOT IN "validateTask" CAUSE BY MISTERY REASON RUIN TEST IF NOT HERE
	now := time.Now()
	_, err := nextdate.NextDate(now, dbTask.Date, dbTask.Repeat)
	if err != nil {
		writeErrorInJson(w, err, http.StatusBadRequest)
		return
	}

	// Update task
	if err := db.UpdateTask(&dbTask); err != nil {
		writeErrorInJson(w, err, http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		writeErrorInJson(w, err, http.StatusInternalServerError)
		return
	}
}
