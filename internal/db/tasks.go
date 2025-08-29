package db

import (
	"database/sql"

	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
)

func GetTasks(quantity int) ([]dbtask.DbTask, error) {
	query := `SELECT * FROM scheduler ORDER BY date LIMIT :quantity;`
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

func GetTasksByWord(quantity int, word string) ([]dbtask.DbTask, error) {
	word = "%" + word + "%"
	rows, err := db.Query(`SELECT * FROM scheduler WHERE title LIKE :word OR comment LIKE :word ORDER BY date LIMIT :quantity;`,
		sql.Named("word", word),
		sql.Named("quantity", quantity))
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

func GetTasksByDate(quantity int, date string) ([]dbtask.DbTask, error) {
	query := `SELECT * FROM scheduler WHERE date = :date LIMIT :quantity;`
	rows, err := db.Query(query,
		sql.Named("date", date),
		sql.Named("quantity", quantity))

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
