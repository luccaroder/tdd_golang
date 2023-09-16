package rules

type Phone struct {
	Number string
}

func (p *Phone) IsValid() []error {
	errs := make([]error, 0)

	if p.Number == "" {
		errs = append(errs, ErrPhoneRequired)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
