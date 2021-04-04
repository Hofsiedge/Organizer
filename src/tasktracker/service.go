package tasktracker

type IService interface {
	CreateTask(string, string, int64) (int64, error)
	DeleteTask(int64) error
	ChangeTaskStatus(int64, TaskStatus) error
	GetTask(int64) (*Task, error)
}

type Middleware func(IService) IService

type Repository interface {
	CreateTask(name string, description string, ownerId int64) (taskId int64, err error)
	GetTask(taskId int64) (task *Task, err error)
	DeleteTask(taskId int64) error
	ChangeTaskStatus(taskId int64, status TaskStatus) error
}

type Service struct {
	Repository Repository
}

func (src Service) CreateTask(name, description string, ownerId int64) (id int64, err error) {
	id, err = src.Repository.CreateTask(name, description, ownerId)
	return
}

func (src Service) GetTask(id int64) (task *Task, err error) {
	task, err = src.Repository.GetTask(id)
	return
}

func (src Service) DeleteTask(id int64) error {
	err := src.Repository.DeleteTask(id)
	return err
}

func (src Service) ChangeTaskStatus(id int64, status TaskStatus) error {
	err := src.Repository.ChangeTaskStatus(id, status)
	return err
}
