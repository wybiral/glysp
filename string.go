package lisp

type String string

func (x String) Iter() Chan {
	ch := make(Chan)
	go func() {
		for _, c := range x {
			ch <- String(c)
		}
		close(ch)
	}()
	return ch
}

func (x String) Repr() E {
	return String("\"" + x + "\"")
}

func (x String) Eval(s *Scope) E {
	return x
}

func (x String) Add(y E) E {
	return x + y.(String)
}

func (x String) Eq(y E) E {
	return Bool(x == y.(String))
}
