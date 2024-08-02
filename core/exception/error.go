package exception

import (
	"errors"
)

const (
	TypeInternal         = "ErrInternal"
	TypeValidation       = "ErrValidation"
	TypeNotFound         = "ErrNotFound"
	TypePermissionDenied = "ErrPermissionDenied"
	TypeTokenExpired     = "TokenExpired"
	TypeTokenInvalid     = "TokenInvalid"
)

type Err = map[string][]string

type Exception struct {
	Type    string `json:"-"`
	Message string `json:"-"`
	Cause   error  `json:"-"`
	Errors  Err    `json:"errors"`
}

func New(kind, message string, err error) *Exception {
	return &Exception{
		Type:    kind,
		Cause:   err,
		Message: message,
		Errors:  make(map[string][]string),
	}
}

func Validation() *Exception {
	return New(TypeValidation, "validation error", nil)
}

func (fail *Exception) HasError() bool {
	return len(fail.Errors) > 0
}

func (fail *Exception) AddError(key, msg string) *Exception {
	fail.Errors[key] = append(fail.Errors[key], msg)
	return fail
}

func Into(err error) *Exception {
	if err == nil {
		return nil
	}
	var fail *Exception
	ok := errors.As(err, &fail)
	if ok {
		return fail
	}
	return New(TypeInternal, err.Error(), err)
}

func (fail *Exception) Error() string {
	if fail.Cause == nil {
		return fail.Message
	}
	return fail.Cause.Error()
}
