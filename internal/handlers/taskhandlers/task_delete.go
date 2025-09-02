package taskhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/kasariks/project_for_graduating/internal/db"
)

func taskDelete(w http.ResponseWriter, r *http.Request) {
	id, err := getIdentifier(r)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeErrorInJson(w, err)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		writeErrorInJson(w, err)
		return
	}
}
