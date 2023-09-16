package rules

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPhone(t *testing.T) {
	scenarios := []struct {
		Name    string
		input   Phone
		wantErr *RuleErrors
	}{
		{
			Name:    "should return error when phone number is empty",
			input:   Phone{Number: ""},
			wantErr: &RuleErrors{Errors: []error{ErrPhoneRequired}},
		},
		{
			Name:    "should return empty error when has an valid phone number",
			input:   Phone{Number: "1888888"},
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
