package taskhandlers

import (
	"encoding/json"
	"net/http"

	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
)

func writeErrorInJson(w http.ResponseWriter, err error) {
	byteErr, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(byteErr)
}

func getJsonFromTasks(dbTasks []dbtask.Task) ([]byte, error) {
	var tasksMap map[string][]map[string]string = map[string][]map[string]string{}

	for _, v := range dbTasks {
		tasksMap["tasks"] = append(tasksMap["tasks"], map[string]string{
			"id":      v.Id,
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

	return byteTasks, err
}
