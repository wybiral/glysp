package lisp

type Symbol string

func (x Symbol) Apply(s *Scope, args List) E {
	s.scope[x] = args[0].Eval(s)
	return nil
}

func (x Symbol) Eval(s *Scope) E {
	var out E
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
