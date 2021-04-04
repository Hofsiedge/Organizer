package middleware

import (
	"github.com/Hofsiedge/Organizer/src/tasktracker"
	"go.uber.org/zap"
	"time"
)

type LoggingMiddleware struct {
	Logger *zap.Logger
	Next   tasktracker.IService
}

// TODO: check if it makes sense to make lambda a pointer
func (mw LoggingMiddleware) HandleError(err error, defaultMessage string) (message string, appErr error, loggingFunc func(string, ...zap.Field)) {
	if err == nil {
		return defaultMessage, nil, mw.Logger.Debug
	}
	if applicationError, ok := err.(*tasktracker.ApplicationError); ok {
		appErr = applicationError
		loggingFunc = mw.Logger.Info
	} else {
		appErr = tasktracker.ErrInternalServerError
		loggingFunc = mw.Logger.Error
	}
	return err.Error(), appErr, loggingFunc // TODO: check if stack trace is included on unexpected errors
}

func (mw LoggingMiddleware) ProcessMethodOutput(defaultMessage string, startTime time.Time, input map[string]interface{}, output interface{}, serviceError error) (err error) {
	var (
		message string
		logger  func(string, ...zap.Field)
	)
	message, err, logger = mw.HandleError(serviceError, defaultMessage)
	logger(message,
		zap.Any("input", input),
		zap.Any("output", output),
		zap.Duration("took", time.Since(startTime)),
	)
	return
}

func (mw LoggingMiddleware) ChangeTaskStatus(id int64, status tasktracker.TaskStatus) error {
	panic("implement me")
}

func (mw LoggingMiddleware) CreateTask(name, description string, ownerId int64) (taskId int64, err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called method CreateTask", begin,
			map[string]interface{}{
				"name":        name,
				"description": description,
				"ownerId":     ownerId,
			}, taskId, err)
	}(time.Now())

	taskId, err = mw.Next.CreateTask(name, description, ownerId)
	return taskId, err
}

func (mw LoggingMiddleware) GetTask(id int64) (task *tasktracker.Task, err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called GetTask", begin, map[string]interface{}{"id": id}, task, err)
	}(time.Now())

	task, err = mw.Next.GetTask(id)
	return task, err
}

// TODO
func (mw LoggingMiddleware) DeleteTask(id int64) (err error) {
	defer func(begin time.Time) {
		panic("not implemented")
	}(time.Now())

	err = mw.Next.DeleteTask(id)
	return
}
