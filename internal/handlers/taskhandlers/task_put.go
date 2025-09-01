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

	var taskMap map[string]string

	err = json.Unmarshal(body, &taskMap)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	// Validate task
	id, err := strconv.Atoi(taskMap["id"])
	if err != nil {
		writeErrorInJson(w, err)
		return
	}
	dbTask.Id = id
	dbTask.Date = taskMap["date"]
	dbTask.Title = taskMap["title"]
	dbTask.Comment = taskMap["comment"]
	dbTask.Repeat = taskMap["repeat"]

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

	// if now.After(t) {
	// 	if len(dbTask.Repeat) == 0 || dbTask.Repeat == "d 1" {
	// 		dbTask.Date = now.Format(nextdate.DateFormat)
	// 	} else {
	// 		next, err := nextdate.NextDate(now, dbTask.Date, dbTask.Repeat)
	// 		if err != nil {
	// 			writeErrorInJson(w, err)
	// 			return
	// 		}

	// 		dbTask.Date = next
	// 	}
	// }

	// Update task
	err = db.UpdateTask(&dbTask)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	byteResp, err := json.Marshal(map[string]string{})
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteResp)
}
