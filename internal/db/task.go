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
	id, err = res.LastInsertId()

	return id, err
}

func GetTasks(quantity int) ([]dbtask.DbTask, error) {
	query := `SELECT * FROM scheduler LIMIT :quantity;`
	rows, err := db.Query(query, sql.Named("quantity", quantity))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dbtask.DbTask
	for rows.Next() {
		task := dbtask.DbTask{}
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
