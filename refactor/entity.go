package refactor

import "tdd/refactor/rules"

type RegisterUserFlow struct {
	Rules map[rules.RulesType]Rules
}

type RulesInput struct {
	Type  rules.RulesType
	Value string
}
