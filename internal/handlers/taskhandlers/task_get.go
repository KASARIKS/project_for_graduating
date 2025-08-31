package taskhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/kasariks/project_for_graduating/internal/db"
)

func taskGet(w http.ResponseWriter, r *http.Request) {
	getId := r.URL.Query().Get("id")
	if len(getId) == 0 {
		writeErrorInJson(w, errors.New("identifier not specified"))
		return
	}

	id, err := strconv.Atoi(getId)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	dbTask, err := db.GetTaskById(id)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	var mapTask map[string]string = map[string]string{
		"id":      strconv.Itoa(dbTask.Id),
		"date":    dbTask.Date,
		"title":   dbTask.Title,
		"comment": dbTask.Comment,
		"repeat":  dbTask.Repeat,
	}

	byteTask, err := json.Marshal(mapTask)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteTask)
}
