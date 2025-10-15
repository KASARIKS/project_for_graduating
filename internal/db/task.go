package db

import (
	"database/sql"
	"fmt"

	dbtask "github.com/kasariks/project_for_graduating/internal/dbEntites/db_task"
)

func AddTask(task *dbtask.Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat);`
	res, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return id, err
	}

	id, err = res.LastInsertId()

	return id, err
}

func GetTaskById(id int) (*dbtask.Task, error) {
	var dbTask dbtask.Task
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id;`
	row := db.QueryRow(query, sql.Named("id", id))
	err := row.Scan(&dbTask.Id, &dbTask.Date, &dbTask.Title, &dbTask.Comment, &dbTask.Repeat)

	return &dbTask, err
}

func UpdateTask(dbtask *dbtask.Task) error {
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id;`
	res, err := db.Exec(query,
		sql.Named("date", dbtask.Date),
		sql.Named("title", dbtask.Title),
		sql.Named("comment", dbtask.Comment),
		sql.Named("repeat", dbtask.Repeat),
		sql.Named("id", dbtask.Id),
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}

	return nil
}

func DeleteTask(id int) error {
	query := `DELETE FROM scheduler WHERE id = :id;`
	_, err := db.Exec(query, sql.Named("id", id))

	return err
}
