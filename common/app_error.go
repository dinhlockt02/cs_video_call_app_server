package common

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Entity string

var (
	UserEntity    Entity = "User"
	MeetingEntity Entity = "Meeting"
	CallEntity    Entity = "Call"
	GroupEntity   Entity = "Group"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

type ValidationError []error

func (err ValidationError) Error() string {
	if len(err) == 0 {
		return ""
	}
	rs := "["
	for i := range err {
		rs += err[i].Error() + ", "
	}
	rs += "]"
	return rs
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (err *AppError) RootError() error {
	if e, ok := err.RootErr.(*AppError); ok {
		return e.RootErr
	}
	return err.RootErr
}

func (err *AppError) Error() string {
	return err.RootError().Error()
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal server error", err.Error(), "ErrInternal")
}

func ErrEntityNotFound(entity Entity, err error) *AppError {
	return NewFullErrorResponse(http.StatusNotFound, err, fmt.Sprintf("%v not found", entity), err.Error(), "ErrNotfound")
}

func ErrForbidden(err error) *AppError {
	return NewFullErrorResponse(http.StatusForbidden, err, "Forbidden Request", err.Error(), "ErrForbidden")
}

func ErrUnauthorized(err error) *AppError {
	return NewFullErrorResponse(http.StatusUnauthorized, err, "Unauthorized", err.Error(), "ErrUnauthorized")
}

var ErrInvalidObjectId = errors.New("invalid object id")
