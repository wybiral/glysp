package lisp

import (
	"fmt"
)

type Dict map[T]T

func (x Dict) Iter() Chan {
	ch := make(Chan)
	go func() {
		for k, v := range x {
			ch <- &List{k, v}
		}
		close(ch)
	}()
	return ch
}

func (x Dict) String() string {
	out := "{"
	for k, v := range x {
		out += fmt.Sprintf("%v: %v, ", Repr(k), Repr(v))
	}
	if len(x) > 0 {
		out = out[:len(out)-2]
	}
	out += "}"
	return out
}

func (x Dict) Apply(s *Scope, args List) T {
	key := Eval(s, args[0])
	if len(args) == 1 {
		return x[key]
	} else if len(args) == 2 {
		x[key] = Eval(s, args[1])
	}
	return nil
}
