package rules

type Email struct {
	Value string
}

func (p *Email) IsValid() []error {
	errs := make([]error, 0)

	if p.Value == "" {
		errs = append(errs, ErrEmailRequired)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
