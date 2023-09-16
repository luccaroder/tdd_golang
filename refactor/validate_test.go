package refactor

import (
	"testing"

	"tdd/refactor/rules"

	"gotest.tools/v3/assert"
)

func TestValidate(t *testing.T) {
	scenarios := []struct {
		Name    string
		input   []Rules
		wantErr *rules.RuleErrors
	}{
		{
			Name:    "should return error when phone number and email are empties",
			input:   []Rules{&rules.Phone{Number: ""}, &rules.Email{Value: ""}},
			wantErr: &rules.RuleErrors{Errors: []error{rules.ErrPhoneRequired, rules.ErrEmailRequired, rules.ErrEmailInvalid}},
		},
	}
	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			err := Validate(s.input)
			assert.Equal(t, err.Error(), s.wantErr.Error())
		})
	}
}
