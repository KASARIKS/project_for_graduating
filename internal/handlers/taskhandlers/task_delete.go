package taskhandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/kasariks/project_for_graduating/internal/db"
)

func taskDelete(w http.ResponseWriter, r *http.Request) {
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

	if err := db.DeleteTask(intId); err != nil {
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
