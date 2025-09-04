package taskhandlers

import (
	"errors"
	"net/http"
)

func Task(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		taskPost(w, r)
	case http.MethodGet:
		taskGet(w, r)
	case http.MethodPut:
		taskPut(w, r)
	case http.MethodDelete:
		taskDelete(w, r)
	default:
		writeErrorInJson(w, errors.New("wrong method"), http.StatusBadRequest)
	}
}
