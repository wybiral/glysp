package lisp

// Lisp Func type
type Func func(...T) T

func (x Func) Apply(s *Scope, args List) T {
	temp := make([]T, len(args))
	for i, v := range args {
		temp[i] = Eval(s, v)
	}
	return x(temp...)
}
