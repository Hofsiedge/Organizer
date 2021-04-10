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

func (r TaskRepository) CreateTask(ctx context.Context, name string, description string, ownerId int64) (taskId int64, err error) {
	row := r.Db.QueryRow(
		ctx,
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

func (r TaskRepository) GetTask(ctx context.Context, taskId int64) (*tasktracker.Task, error) {
	rows, err := r.Db.Query(
		ctx,
		"SELECT * FROM tasktracker.task WHERE id = $1;",
		taskId,
	)
	if err != nil {
		return nil, fmt.Errorf("Query failed in GetCharacter: %v\n", err)
	}

	task := new(tasktracker.Task)
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

func (r TaskRepository) GetTasks(ctx context.Context, userId int64) ([]*tasktracker.Task, error) {
	rows, err := r.Db.Query(
		ctx,
		"SELECT * FROM tasktracker.get_tasks($1);",
		userId,
	)
	if err != nil {
		return nil, fmt.Errorf("query failed in GetTasks: %v", err)
	}
	var tasks []*tasktracker.Task
	if err := pgxscan.ScanAll(&tasks, rows); err != nil {
		return nil, fmt.Errorf("query failed in GetTasks: %v", err)
	}
	return tasks, nil
}

func (r TaskRepository) UpdateTask(ctx context.Context, taskId int64, name, description *string, status *tasktracker.TaskStatus, userId int64) error {
	if _, err := r.Db.Query(
		ctx,
		"CALL tasktracker.update_task($1, $2, $3, $4, $5);",
		taskId, name, description, status, userId,
	); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r TaskRepository) DeleteTask(ctx context.Context, taskId int64) error {
	if _, err := r.Db.Query(
		ctx,
		"CALL tasktracker.delete_task($1);",
		taskId,
	); err != nil {
		// TODO: handle 'task not found'
		return errors.WithStack(err)
	}
	return nil
}
