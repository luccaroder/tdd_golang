package refactor

import "tdd/refactor/rules"

func NewRegisterUserFlow() *RegisterUserFlow {
	rulesManager := map[rules.RulesType]Rules{
		rules.EmailValidate: &rules.Email{},
		rules.PhoneValidate: &rules.Phone{},
	}

	return &RegisterUserFlow{
		Rules: rulesManager,
	}
}
