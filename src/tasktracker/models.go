package tasktracker

import "github.com/jackc/pgtype"

// Task model
type Task struct {
	Id          int64       `json:"id"`
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
	Status      TaskStatus  `json:"status"`
	OwnerId     int64       `db:"owner_id" json:"owner_id"`
}

type TaskStatus string

const (
	Created   TaskStatus = "created"
	InProcess TaskStatus = "in process"
	Failed    TaskStatus = "failed"
	Done      TaskStatus = "done"
	Planned   TaskStatus = "planned"
	Archived  TaskStatus = "archived"
)
