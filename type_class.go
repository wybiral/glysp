package lisp

type Class struct {
	parent *Class
	name   String
	attrs  map[Symbol]E
}

func (x *Class) Eval(s *Scope) E { return x }

func (x *Class) Apply(s *Scope, args List) E {
	self := &Object{x, make(map[Symbol]E)}
	init := x.attrs[Symbol("__init__")]
	if init != nil {
		init.(*Method).closure.Apply(s, append(List{self}, args...))
	}
	return self
}
