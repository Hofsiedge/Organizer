package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Hofsiedge/Organizer/src/tasktracker"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type JSONResponse struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(*tasktracker.ApplicationError); ok {
		w.WriteHeader(err.StatusCode())
		return json.NewEncoder(w).Encode(JSONResponse{
			Data:  nil,
			Error: err,
		})
	}
	return json.NewEncoder(w).Encode(JSONResponse{
		Data:  response,
		Error: nil,
	})
}

/* 						CreateTask							*/
type createTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     int64  `json:"owner_id"`
}

type createTaskResponse struct {
	Id  int64 `json:"id"`
	Err error `json:"err,omitempty"`
}

func MakeCreateTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTaskRequest)
		id, err := svc.CreateTask(ctx, req.Name, req.Description, req.OwnerId)
		if err != nil {
			return nil, err
		}
		return createTaskResponse{id, nil}, nil
	}
}

func DecodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, tasktracker.ErrBadRequest
	}
	return request, nil
}

/* 							GetTask							*/
type getTaskRequest struct {
	Id int64 `json:"id"`
}

type getTaskResponse struct {
	Task *tasktracker.Task `json:"task"`
}

func MakeGetTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTaskRequest)
		task, err := svc.GetTask(ctx, req.Id)
		if err != nil {
			if appErr, ok := err.(*tasktracker.ApplicationError); ok {
				return appErr, nil
			}
			return nil, err
		}
		return getTaskResponse{task}, nil
	}
}

func DecodeGetTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	strId, found := vars["id"]
	if !found || strId == "" {
		return nil, fmt.Errorf("invalid id parameter")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return nil, err
	}
	return getTaskRequest{id}, nil
}

/* 							DeleteTask						*/
type deleteTaskRequest struct {
	Id int64 `json:"id"`
}

func MakeDeleteTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteTaskRequest)
		err := svc.DeleteTask(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return struct{}{}, nil
	}
}

func DecodeDeleteTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, tasktracker.ErrBadRequest
	}
	return request, nil
}

/* 							UpdateTask						*/
type updateTaskRequest struct {
	Id          int64
	Name        *string                 `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Status      *tasktracker.TaskStatus `json:"status,omitempty"`
}

func DecodeUpdateTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	strId, found := vars["id"]
	if !found || strId == "" {
		return nil, tasktracker.ErrBadRequest
	}
	var id int64
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return nil, tasktracker.ErrBadRequest
	}

	var request updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, tasktracker.ErrBadRequest
	}
	request.Id = id
	if request.Name == nil && request.Description == nil && request.Status == nil {
		return nil, &tasktracker.ApplicationError{
			Message: "Missing parameters: at least one of 'name', 'description', 'status' must be specified",
			Status:  400,
		}
	}
	return request, nil
}

func MakeUpdateTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateTaskRequest)
		// TODO: obtain userId from JWT
		err := svc.UpdateTask(ctx, req.Id, req.Name, req.Description, req.Status, 1)
		if err != nil {
			return nil, err
		}
		return struct{}{}, nil
	}
}
