package lisp

// Lisp Chan type
type Chan chan E

func (x Chan) Eval(s *Scope) E { return x }

func (x Chan) Iter() Chan { return x }

func (x Chan) Apply(s *Scope, args List) E {
	if len(args) == 0 {
		return <-x
	}
	for _, v := range args {
		x <- v.Eval(s)
	}
	return nil
}
