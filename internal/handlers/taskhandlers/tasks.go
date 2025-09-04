package taskhandlers

import (
	"net/http"
	"time"

	"github.com/kasariks/project_for_graduating/internal/db"
	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var dbTasks []dbtask.Task
	var err error

	searchParam := r.URL.Query().Get("search")

	dbTasks, err = getTasksFromDbBySearch(searchParam)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	byteTasks, err := getJsonFromTasks(dbTasks)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	w.Write(byteTasks)
}

func getTasksFromDbBySearch(searchParam string) ([]dbtask.Task, error) {
	var dbTasks []dbtask.Task
	var err error

	if searchParam == "" {
		dbTasks, err = db.GetTasks(50)
		if err != nil {
			return dbTasks, err
		}
	} else if t, err := time.Parse("02.01.2006", searchParam); err == nil {
		dbTime := t.Format(nextdate.DateFormat)
		dbTasks, err = db.GetTasksByDate(50, dbTime)
		if err != nil {
			return dbTasks, err
		}
	} else {
		dbTasks, err = db.GetTasksByWord(50, searchParam)
		if err != nil {
			return dbTasks, err
		}
	}

	if len(dbTasks) == 0 {
		dbTasks = []dbtask.Task{}
	}

	return dbTasks, err
}
