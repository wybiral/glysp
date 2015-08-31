package lisp

type Method struct {
	name    String
	closure *Closure
}

func (x *Method) String() string { return string(x.name) }

func (x *Method) Apply(s *Scope, args List) T {
	return x.closure.Apply(s, args)
}
