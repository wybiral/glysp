package lisp

type Symbol string

func (x Symbol) Apply(s *Scope, args List) T {
	s.scope[x] = Eval(s, args[0])
	return nil
}

func (x Symbol) Eval(s *Scope) T {
	var out T
	for s != nil {
		out = s.scope[x]
		if out != nil {
			goto found
		}
		s = s.parent
	}
	FatalError("Symbol not found: " + string(x))
found:
	return out
}
