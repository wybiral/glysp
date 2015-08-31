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

func (x String) Repr() T {
	return String("\"" + x + "\"")
}

func (x String) Add(y T) T {
	return x + y.(String)
}

func (x String) Eq(y T) T {
	return Bool(x == y.(String))
}
