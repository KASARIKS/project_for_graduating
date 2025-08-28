package handlers

import (
	"errors"
	"net/http"
)

func AddNewTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, errors.New("wrong request format").Error(), http.StatusBadRequest)
	}

}
