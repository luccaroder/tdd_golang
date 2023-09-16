package refactor

import (
	"tdd/refactor/rules"
)

type Rules interface {
	IsValid() *rules.RuleErrors
}
