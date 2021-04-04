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
	Payload interface{} `json:"payload"`
	Error   error       `json:"error"`
}

func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(*tasktracker.ApplicationError); ok {
		w.WriteHeader(err.StatusCode())
		return json.NewEncoder(w).Encode(JSONResponse{
			Payload: nil,
			Error:   err,
		})
	}
	return json.NewEncoder(w).Encode(JSONResponse{
		Payload: response,
		Error:   nil,
	})
}

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
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createTaskRequest)
		id, err := svc.CreateTask(req.Name, req.Description, req.OwnerId)
		if err != nil {
			return nil, err
		}
		return createTaskResponse{id, nil}, nil
	}
}

func DecodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type getTaskRequest struct {
	Id int64 `json:"id"`
}

type getTaskResponse struct {
	Task *tasktracker.Task `json:"task"`
}

func MakeGetTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getTaskRequest)
		task, err := svc.GetTask(req.Id)
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

type deleteTaskRequest struct {
	Id int64 `json:"id"`
}

type deleteTaskResponse struct{}

func MakeDeleteTaskEndpoint(svc tasktracker.IService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteTaskRequest)
		err := svc.DeleteTask(req.Id)
		if err != nil {
			return nil, err
		}
		return deleteTaskResponse{}, nil
	}
}

func DecodeDeleteTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeChangeTaskStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("not implemented")
}
