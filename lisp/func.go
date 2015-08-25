package lisp

// Lisp Func type
type Func func(List) E

func (x Func) Eval(s *Scope) E { return x }
func (x Func) Apply(s *Scope, args List) E {
	temp := make([]E, len(args))
	for i, v := range args {
		temp[i] = v.Eval(s)
	}
	return x(temp)
}
