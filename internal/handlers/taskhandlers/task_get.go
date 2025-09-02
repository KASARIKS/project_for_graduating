package taskhandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kasariks/project_for_graduating/internal/db"
)

func taskGet(w http.ResponseWriter, r *http.Request) {
	id, err := getIdentifier(r)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	dbTask, err := db.GetTaskById(id)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	_, err = strconv.Atoi(dbTask.Id)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if err = json.NewEncoder(w).Encode(dbTask); err != nil {
		writeErrorInJson(w, err)
		return
	}
}
