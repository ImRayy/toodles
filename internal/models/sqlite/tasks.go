package sqlite

import (
	"database/sql"
	"log"
	"toodles/db"
	"toodles/internal/models"

	_ "modernc.org/sqlite"
)

func parseTasks(rows *sql.Rows) ([]models.Task, error) {
	tasks := []models.Task{}

	for rows.Next() {
		t := models.Task{}

		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Status, &t.Priority,
			&t.CreatedAt,
			&t.DoneAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}
	return tasks, nil
}

func AllTasks() ([]models.Task, error) {
	rows, err := db.Sqlite.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(rows)
	if err != nil {
		log.Fatal(err)
	}

	return tasks, nil
}

func CompletedTasks() ([]models.Task, error) {
	rows, err := db.Sqlite.Query("SELECT * FROM tasks WHERE status = 'done'")
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(rows)
	if err != nil {
		log.Fatal(err)
	}

	return tasks, nil
}

func PendingTasks() ([]models.Task, error) {
	rows, err := db.Sqlite.Query("SELECT * FROM tasks WHERE status = 'pending'")
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(rows)
	if err != nil {
		log.Fatal(err)
	}

	return tasks, nil
}
