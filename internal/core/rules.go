package core

type Rule struct {
	ID      string
	Version string

	Source      []byte
	SymbolTable *SymbolTable

	Condition *Condition
	Runner    *Runner
}

func NewRule() *Rule {
	return &Rule{
		ID:          "",
		Version:     "",
		Source:      nil,
		SymbolTable: NewSymbolTable(),
		Condition:   NewCondition(),
		Runner:      NewRunner(),
	}
}
