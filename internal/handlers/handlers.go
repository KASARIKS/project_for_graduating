package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	//"api/nextdate?now=20240126&date=20240126&repeat=y"
	now, err := time.Parse(nextdate.DateFormat, r.FormValue("now"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	finalDate, err := nextdate.NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	io.WriteString(w, finalDate)
}
