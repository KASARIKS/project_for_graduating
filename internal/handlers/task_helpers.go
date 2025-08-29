package handlers

import (
	"encoding/json"
	"net/http"
)

func writeErrorInJson(w http.ResponseWriter, err error) {
	byteErr, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(byteErr)
}
