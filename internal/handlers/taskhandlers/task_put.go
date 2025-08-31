package taskhandlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
)

func taskPut(w http.ResponseWriter, r *http.Request) {
	var dbTask dbtask.DbTask

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
