package lisp

// Lisp Macro type
type Macro func(*Scope, List) T

func (x Macro) Apply(s *Scope, args List) T {
	return x(s, args)
}
