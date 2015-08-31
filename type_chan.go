package lisp

// Lisp Chan type
type Chan chan T

func (x Chan) Iter() Chan { return x }

func (x Chan) Apply(s *Scope, args List) T {
	if len(args) == 0 {
		return <-x
	}
	for _, v := range args {
		x <- v.(Evalable).Eval(s)
	}
	return nil
}
