package lisp

import (
	"fmt"
)

type List []T

func (x List) String() string {
	out := "["
	for _, v := range x {
		out += fmt.Sprintf("%v, ", Repr(v))
	}
	if len(x) > 0 {
		out = out[:len(out)-2]
	}
	out += "]"
	return out
}

func (x List) Iter() Chan {
	ch := make(Chan)
	go func() {
		for _, v := range x {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

/* TODO: Check for Applyable report errors */
func (x List) Eval(s *Scope) T {
	fn := Eval(s, x[0])
	return Apply(s, fn, x[1:])
}

func (x List) Apply(s *Scope, args List) T {
	index := Eval(s, args[0])
	if len(args) == 1 {
		return x[index.(Int)]
	} else if len(args) == 2 {
		x[index.(Int)] = Eval(s, args[1])
	}
	return nil
}
