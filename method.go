package lisp

type Method struct {
	name    String
	closure *Closure
}

func (x *Method) String() string { return string(x.name) }

func (x *Method) Eval(s *Scope) E { return x }
