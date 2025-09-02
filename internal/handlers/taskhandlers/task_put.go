package taskhandlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func taskPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var dbTask dbtask.DbTask

	// Get task from request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	err = json.Unmarshal(body, &dbTask)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	// Validate task
	_, err = strconv.Atoi(dbTask.Id)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if len(dbTask.Title) == 0 {
		writeErrorInJson(w, errors.New("empty title"))
		return
	}

	_, err = time.Parse(nextdate.DateFormat, dbTask.Date)
	if err != nil {
		writeErrorInJson(w, errors.New("incorrect date"))
		return
	}

	now := time.Now()
	_, err = nextdate.NextDate(now, dbTask.Date, dbTask.Repeat)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	// Update task
	err = db.UpdateTask(&dbTask)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		writeErrorInJson(w, err)
		return
	}
}
