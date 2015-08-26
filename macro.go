package lisp

// Lisp Macro type
type Macro func(*Scope, List) E

func (x Macro) Eval(s *Scope) E { return x }
func (x Macro) Apply(s *Scope, args List) E {
	return x(s, args)
}
