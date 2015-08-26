package lisp

type BoundMethod struct {
	method   *Method
	instance *Object
}

func (x *BoundMethod) Eval(s *Scope) E { return x }

func (x *BoundMethod) String() string { return string(x.method.name) }

func (x *BoundMethod) Apply(s *Scope, args List) E {
	selfargs := append(List{x.instance}, args...)
	return x.method.closure.Apply(s, selfargs)
}
