package lisp

type Class struct {
	parent *Class
	name   String
	attrs  map[Symbol]T
}

func (x *Class) Apply(s *Scope, args List) T {
	self := &Object{x, make(map[Symbol]T)}
	init := x.attrs[Symbol("__init__")]
	if init != nil {
		init.(*Method).closure.Apply(s, append(List{self}, args...))
	}
	return self
}

func (x *Class) GetAttr(key Symbol) T {
	out := x.attrs[key]
	for out == nil {
		out = x.attrs[key]
		x = x.parent
	}
	return out
}

func (x *Class) SetAttr(key Symbol, val T) {
	x.attrs[key] = val
}
