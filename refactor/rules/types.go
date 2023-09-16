package rules

type RulesType int

func (RulesType) Values() []string {
	return []string{"_", "phoneValidate", "emailValidate"}
}

func (a RulesType) String() string {
	return a.Values()[a]
}

const (
	_ RulesType = iota
	PhoneValidate
	EmailValidate
)
