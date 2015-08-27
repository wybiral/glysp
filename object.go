package lisp

type Object struct {
	class *Class
	attrs map[Symbol]E
}

func (x *Object) String() string {
	fn := x.GetAttr(Symbol("__str__"))
	if fn != nil {
		c := fn.(*Method).closure
		str := c.Apply(c.Local, List{x})
		return string(str.(String))
	}
	return ""
}

func binaryMethod(x *Object, y E, key string) E {
	fn := x.GetAttr(Symbol(key))
	if fn != nil {
		c := fn.(*Method).closure
		return c.Apply(c.Local, List{x, y})
	}
	return nil
}

func (x *Object) Add(y E) E { return binaryMethod(x, y, "__add__") }
func (x *Object) Sub(y E) E { return binaryMethod(x, y, "__sub__") }
func (x *Object) Mul(y E) E { return binaryMethod(x, y, "__mul__") }
func (x *Object) Div(y E) E { return binaryMethod(x, y, "__div__") }
func (x *Object) Pow(y E) E { return binaryMethod(x, y, "__pow__") }

func (x *Object) Eval(s *Scope) E { return x }

func (x *Object) Apply(s *Scope, args List) E {
	fn := x.GetAttr(Symbol("__call__"))
	if fn != nil {
		c := fn.(*Method).closure
		return c.Apply(c.Local, append(List{x}, args...))
	}
	return nil
}

func (x *Object) GetAttr(key Symbol) E {
	cls := x.class
	out := x.attrs[key]
	for out == nil {
		out = cls.attrs[key]
		cls = cls.parent
	}
	return out
}

func (x *Object) SetAttr(key Symbol, val E) {
	x.attrs[key] = val
}
