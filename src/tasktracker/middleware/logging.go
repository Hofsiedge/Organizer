package middleware

import (
	"context"
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

func (mw LoggingMiddleware) CreateTask(ctx context.Context, name, description string, ownerId int64) (taskId int64, err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called method CreateTask", begin,
			map[string]interface{}{
				"name":        name,
				"description": description,
				"ownerId":     ownerId,
			}, taskId, err)
	}(time.Now())

	taskId, err = mw.Next.CreateTask(ctx, name, description, ownerId)
	return taskId, err
}

func (mw LoggingMiddleware) GetTask(ctx context.Context, id int64) (task *tasktracker.Task, err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called GetTask", begin, map[string]interface{}{
			"id": id,
		}, task, err)
	}(time.Now())

	task, err = mw.Next.GetTask(ctx, id)
	return task, err
}

func (mw LoggingMiddleware) GetTasks(ctx context.Context, userId int64) (tasks []*tasktracker.Task, err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called GetTasks", begin, nil, tasks, err)
	}(time.Now())

	tasks, err = mw.Next.GetTasks(ctx, userId)
	return tasks, err
}

func (mw LoggingMiddleware) UpdateTask(ctx context.Context, taskId int64, name, description *string, status *tasktracker.TaskStatus, userId int64) (err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called UpdateTask", begin, map[string]interface{}{
			"taskId":      taskId,
			"name":        name,
			"description": description,
			"status":      status,
			"userId":      userId,
		}, nil, err)
	}(time.Now())

	err = mw.Next.UpdateTask(ctx, taskId, name, description, status, userId)
	return err
}

func (mw LoggingMiddleware) DeleteTask(ctx context.Context, id int64) (err error) {
	defer func(begin time.Time) {
		err = mw.ProcessMethodOutput("called DeleteTask", begin, map[string]interface{}{
			"id": id,
		}, nil, err)
	}(time.Now())

	err = mw.Next.DeleteTask(ctx, id)
	return
}
