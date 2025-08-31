package taskhandlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
)

func taskPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var dbTask dbtask.DbTask

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

	err = db.UpdateTask(&dbTask)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	byteResp, err := json.Marshal("{}")
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteResp)
}
