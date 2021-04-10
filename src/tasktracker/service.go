package tasktracker

import "context"

type IService interface {
	CreateTask(ctx context.Context, name, description string, ownerId int64) (int64, error)
	DeleteTask(ctx context.Context, taskId int64) error
	UpdateTask(ctx context.Context, taskId int64, name, description *string, status *TaskStatus, userId int64) error
	GetTask(ctx context.Context, taskId int64) (*Task, error)
}

type Middleware func(IService) IService

type Repository interface {
	CreateTask(ctx context.Context, name string, description string, ownerId int64) (taskId int64, err error)
	GetTask(ctx context.Context, taskId int64) (task *Task, err error)
	DeleteTask(ctx context.Context, taskId int64) error
	UpdateTask(ctx context.Context, taskId int64, name, description *string, status *TaskStatus, userId int64) error
}

type Service struct {
	Repository Repository
}

func (src Service) CreateTask(ctx context.Context, name, description string, ownerId int64) (id int64, err error) {
	id, err = src.Repository.CreateTask(context.Background(), name, description, ownerId)
	return
}

func (src Service) GetTask(ctx context.Context, id int64) (task *Task, err error) {
	task, err = src.Repository.GetTask(context.Background(), id)
	return
}

func (src Service) DeleteTask(ctx context.Context, id int64) error {
	err := src.Repository.DeleteTask(context.Background(), id)
	return err
}

func (src Service) UpdateTask(ctx context.Context, id int64, name, description *string, status *TaskStatus, userId int64) error {
	/*
		Potential errors:
		- Not found
		- Incorrect status
		- Statuses are the same
		- [Invalid change]
		- Permission denial (the user must be the task owner)
	*/
	err := src.Repository.UpdateTask(context.Background(), id, name, description, status, userId)
	return err
}
