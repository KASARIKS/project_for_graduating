package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var dbTasks []dbtask.DbTask
	var err error

	// Choose search method
	searchParam := r.URL.Query().Get("search")

	if searchParam == "" {
		dbTasks, err = db.GetTasks(50)
		if err != nil {
			writeErrorInJson(w, err)
			return
		}
	} else if t, err := time.Parse("02.01.2006", searchParam); err == nil {
		dbTime := t.Format(nextdate.DateFormat)
		dbTasks, err = db.GetTasksByDate(50, dbTime)
		if err != nil {
			writeErrorInJson(w, err)
			return
		}
	} else {
		dbTasks, err = db.GetTasksByWord(50, searchParam)
		if err != nil {
			writeErrorInJson(w, err)
			return
		}
	}

	if len(dbTasks) == 0 {
		dbTasks = []dbtask.DbTask{}
	}

	var tasksMap map[string][]map[string]string = map[string][]map[string]string{}

	for _, v := range dbTasks {
		tasksMap["tasks"] = append(tasksMap["tasks"], map[string]string{
			"id":      strconv.Itoa(v.Id),
			"date":    v.Date,
			"title":   v.Title,
			"comment": v.Comment,
			"repeat":  v.Repeat,
		})
	}

	if tasksMap["tasks"] == nil {
		tasksMap["tasks"] = []map[string]string{}
	}

	byteTasks, err := json.Marshal(tasksMap)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteTasks)
}
