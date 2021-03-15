package database

import (
	"context"
	"fmt"

	"github.com/Hofsiedge/Organizer/src/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetTask(dbPool *pgxpool.Pool, taskId int) (models.Task, error) {
	rows, err := dbPool.Query(
		context.Background(),
		"SELECT * from tasktracker.task WHERE id = $1;",
		taskId,
	)
	if err != nil {
		return models.Task{}, fmt.Errorf("Query failed in GetCharacter: %v\n", err)
	}

	var task models.Task
	if err := pgxscan.ScanOne(&task, rows); err != nil {
		return models.Task{}, fmt.Errorf("Row scan failed in GetCharacter: %v\n", err)
	}

	return task, nil
}

func CreateTask(dbPool *pgxpool.Pool, name string, description string, ownerId int64) (int64, error) {
	row := dbPool.QueryRow(
		context.Background(),
		"SELECT * from tasktracker.create_task($1, $2, $3);",
		name,
		description,
		ownerId,
	)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Query failed in CreateTask: %v\n", err)
	}
	return id, nil
}
