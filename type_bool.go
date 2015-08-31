package lisp

// Lisp Bool type
type Bool bool

func (x Bool) Eval(s *Scope) E { return x }
