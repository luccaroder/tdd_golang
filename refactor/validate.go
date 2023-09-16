package refactor

import "tdd/refactor/rules"

func Validate(r []RulesInput) *rules.RuleErrors {
	ruleErr := &rules.RuleErrors{}
	register := NewRegisterUserFlow()
	for _, rule := range r {
		rule := register.Rules[rule.Type]
		err := rule.IsValid()
		if err != nil {
			ruleErr.Errors = append(ruleErr.Errors, err.Errors...)
		}
	}

	if len(ruleErr.Errors) > 0 {
		return ruleErr
	}

	return nil
}
