package tasktracker

type ApplicationError struct {
	Message string `json:"message"`
	Status  int    `json:"status_code"`
}

// implementation of go-kit's http.StatusCoder interface
func (e *ApplicationError) StatusCode() int {
	return e.Status
}

// implementation of error interface
func (e *ApplicationError) Error() string {
	return e.Message
}

var (
	ErrBadRequest = &ApplicationError{
		Message: "Bad request",
		Status:  400,
	}

	ErrTaskNotFound = &ApplicationError{
		Message: "Task not found",
		Status:  404,
	}

	ErrInternalServerError = &ApplicationError{
		Message: "Internal server error",
		Status:  500,
	}
)
