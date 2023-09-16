package rules

import "github.com/pkg/errors"

var (
	ErrPhoneRequired = errors.New("ErrPhoneRequired")
	ErrEmailRequired = errors.New("ErrEmailRequired")
)
