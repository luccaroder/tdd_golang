package refactor

import "tdd/refactor/rules"

type RegisterUserFlow struct {
	Rules map[rules.RulesType]Rules
}

func Validate(r []Rules) *rules.RuleErrors {
	ruleErr := &rules.RuleErrors{}
	for _, v := range r {
		if err := v.IsValid(); err != nil {
			ruleErr.Errors = append(ruleErr.Errors, err.Errors...)
		}
	}
	if len(ruleErr.Errors) > 0 {
		return ruleErr
	}
	return nil
}
