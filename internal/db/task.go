package db

import (
	"database/sql"

	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
)

func AddTask(task *dbtask.DbTask) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat);`
	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	// if err != nil {
	// 	return id, err
	// }

	id, err = res.LastInsertId()

	return id, err
}
