package postgres

import (
	"context"
	"fmt"
	"github.com/Hofsiedge/Organizer/src/tasktracker"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type TaskRepository struct {
	Db *pgxpool.Pool
}

func (r TaskRepository) CreateTask(name string, description string, ownerId int64) (taskId int64, err error) {
	row := r.Db.QueryRow(
		context.Background(),
		"SELECT * from tasktracker.create_task($1, $2, $3);",
		name,
		description,
		ownerId,
	)
	// TODO: can an ApplicationError occur?
	if err := row.Scan(&taskId); err != nil {
		return 0, fmt.Errorf("Query failed in DBCreateTask: %v\n", err)
	}
	return taskId, nil
}

func (r TaskRepository) GetTask(taskId int64) (task *tasktracker.Task, err error) {
	rows, err := r.Db.Query(
		context.Background(),
		"SELECT * from tasktracker.task WHERE id = $1;",
		taskId,
	)
	if err != nil {
		return nil, fmt.Errorf("Query failed in GetCharacter: %v\n", err)
	}

	task = new(tasktracker.Task)
	if err := pgxscan.ScanOne(task, rows); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, tasktracker.ErrTaskNotFound
		}
		// return nil, fmt.Errorf("Row scan failed in GetCharacter: %v\n", err)
		// TODO: check if WithStack is required
		// TODO: check if stack is printed in logs
		return nil, errors.WithStack(err)
	}
	return task, nil
}

func (r TaskRepository) DeleteTask(taskId int64) error {
	row := r.Db.QueryRow(
		context.Background(),
		"SELECT * from tasktracker.delete_task($1);",
		taskId,
	)
	if err := row.Scan(); err != nil {
		// TODO: check what happens if task not found
		return fmt.Errorf("Query failed in DBDeleteTask: %v\n", err)
	}
	return nil
}

func (r TaskRepository) ChangeTaskStatus(taskId int64, status tasktracker.TaskStatus) error {
	panic("implement me")
}
