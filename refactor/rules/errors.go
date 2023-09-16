package rules

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrPhoneRequired = errors.New("The phone is required")
	ErrEmailRequired = errors.New("The email is required")
	ErrEmailInvalid  = errors.New("The email is invalid")
)

type RuleErrors struct {
	Errors []error
}

func (e *RuleErrors) Error() string {
	if e == nil || len(e.Errors) == 0 {
		return ""
	}
	var errors []string
	for _, e := range e.Errors {
		errors = append(errors, e.Error())
	}

	return strings.Join(errors, ", ")
}
