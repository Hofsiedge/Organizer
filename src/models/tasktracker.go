package models

import "github.com/jackc/pgtype"

// Task model
type Task struct {
	Id          int64
	Name        pgtype.Text
	Description pgtype.Text
	Status      TaskStatus
	OwnerId     int64 `db:"owner_id"`
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
