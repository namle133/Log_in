package service

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound           = errNotFound{}
	ErrUnknown            = errUnknown{}
	ErrNameIsRequired     = errNameIsRequired{}
	ErrRecordNotFound     = errRecordNotFound{}
	ErrPasswordIsRequired = errPasswordIsRequired{}
	ErrEmailIsRequired    = errEmailIsRequired{}
	ErrUserIsExist        = errUserIsExist{}
	ErrTokenIsInvalid     = errTokenIsInvalid{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "username is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errPasswordIsRequired struct{}

func (errPasswordIsRequired) Error() string {
	return "password is required"
}

func (errPasswordIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errEmailIsRequired struct{}

func (errEmailIsRequired) Error() string {
	return "email is required"
}

func (errEmailIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errUserIsExist struct{}

func (errUserIsExist) Error() string {
	return "user is exist"
}

func (errUserIsExist) StatusCode() int {
	return http.StatusBadRequest
}

type errTokenIsInvalid struct{}

func (errTokenIsInvalid) Error() string {
	return "token is invalid"
}

func (errTokenIsInvalid) StatusCode() int {
	return http.StatusBadRequest
}
