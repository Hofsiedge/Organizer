package tasktracker

import "context"

type IService interface {
	CreateTask(ctx context.Context, name, description string, ownerId int64) (int64, error)
	GetTask(ctx context.Context, taskId int64) (*Task, error)
	// TODO: filtering
	GetTasks(ctx context.Context, userId int64) ([]*Task, error)
	UpdateTask(ctx context.Context, taskId int64, name, description *string, status *TaskStatus, userId int64) error
	DeleteTask(ctx context.Context, taskId int64) error
}

/*
Ways to implement filtering:
1. Use Filter objects that generate a part of the WHERE clause and pass all of them to the DB
	PL/pgSQL argument - WHERE clause
2. Use Filter objects on predefined criteria:
	* date
	* status
	* name
	* tag
	Explicitly map filters to distinct arguments of a PL/pgSQL function and construct the WHERE clause in PL/pgSQL
2 is probably the way to go, but I'll have to decide taking into consideration prepared statement behaviour and profiling.
*/

type Middleware func(IService) IService

type Repository interface {
	CreateTask(ctx context.Context, name string, description string, ownerId int64) (taskId int64, err error)
	GetTask(ctx context.Context, taskId int64) (task *Task, err error)
	GetTasks(ctx context.Context, userId int64) (tasks []*Task, err error)
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

func (src Service) GetTasks(ctx context.Context, userId int64) (tasks []*Task, err error) {
	tasks, err = src.Repository.GetTasks(context.Background(), userId)
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
