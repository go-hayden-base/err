package err

import (
	"errors"
	"strconv"
)

// Err is interface to extend error interface by adding method Code()
type Err interface {
	error
	Code() int
}

// ErrBase is struct that implement Err interface
type ErrBase struct {
	CodeNumber  int
	ErrorObject error
}

func (s *ErrBase) Error() string {
	return "Error: { Code:" + strconv.Itoa(s.CodeNumber) + "  Message: \"" + s.ErrorObject.Error() + "\" }"
}

// Code return error code
func (s *ErrBase) Code() int {
	return s.CodeNumber
}

// NewErr reutrn a new Err implementation
func NewErr(code int, err error) Err {
	return &ErrBase{CodeNumber: code, ErrorObject: err}
}

// NewErrMessage reutrn a new Err implementation
func NewErrMessage(code int, message string) Err {
	return &ErrBase{CodeNumber: code, ErrorObject: errors.New(message)}
}
