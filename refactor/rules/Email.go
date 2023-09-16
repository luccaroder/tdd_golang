package rules

import "net/mail"

type Email struct {
	Value string
}

func (p *Email) IsValid() *RuleErrors {
	errs := make([]error, 0)

	if p.Value == "" {
		errs = append(errs, ErrEmailRequired)
	}

	if !valid(p.Value) {
		errs = append(errs, ErrEmailInvalid)
	}
	if len(errs) > 0 {
		return &RuleErrors{Errors: errs}
	}

	return nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
