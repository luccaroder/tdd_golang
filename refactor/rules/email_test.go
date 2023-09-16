package rules

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestEmail(t *testing.T) {
	scenarios := []struct {
		Name    string
		input   Email
		wantErr *RuleErrors
	}{
		{
			Name:    "should return error when email is empty",
			input:   Email{Value: ""},
			wantErr: &RuleErrors{Errors: []error{ErrEmailRequired, ErrEmailInvalid}},
		},
		{
			Name:    "should return error when email is invalid",
			input:   Email{Value: "test-test"},
			wantErr: &RuleErrors{Errors: []error{ErrEmailInvalid}},
		},
		{
			Name:    "should return empty error when has an valid email",
			input:   Email{Value: "test-test@test.com"},
			wantErr: &RuleErrors{Errors: []error{}},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			err := s.input.IsValid()
			assert.Equal(t, err.Error(), s.wantErr.Error())
		})
	}
}
