package db

import (
	"database/sql"

	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
)

func GetTasks(quantity int) ([]dbtask.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :quantity;`
	rows, err := db.Query(query, sql.Named("quantity", quantity))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dbtask.Task
	for rows.Next() {
		task := dbtask.Task{}
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTasksByWord(quantity int, word string) ([]dbtask.Task, error) {
	word = "%" + word + "%"
	rows, err := db.Query(`SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :word OR comment LIKE :word ORDER BY date LIMIT :quantity;`,
		sql.Named("word", word),
		sql.Named("quantity", quantity))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dbtask.Task
	for rows.Next() {
		task := dbtask.Task{}
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTasksByDate(quantity int, date string) ([]dbtask.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date LIMIT :quantity;`
	rows, err := db.Query(query,
		sql.Named("date", date),
		sql.Named("quantity", quantity))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []dbtask.Task
	for rows.Next() {
		task := dbtask.Task{}
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}
