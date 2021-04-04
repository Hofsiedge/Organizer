package tasktracker

type ApplicationError struct {
	Message string `json:"message"`
	Status  int    `json:"status_code"`
}

// TODO: is this still needed?
// implementation of go-kit's http.StatusCoder interface
func (e *ApplicationError) StatusCode() int {
	return e.Status
}

// implementation of error interface
func (e *ApplicationError) Error() string {
	return e.Message
}

var ErrTaskNotFound = &ApplicationError{
	"Task not found",
	404,
}

var ErrInternalServerError = &ApplicationError{
	"Internal server error",
	500,
}
