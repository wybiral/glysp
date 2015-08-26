package lisp

import (
	"fmt"
)

type List []E

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

func (x List) Eval(s *Scope) E {
	return x[0].Eval(s).(interface {
		Apply(*Scope, List) E
	}).Apply(s, x[1:])
}

func (x List) Apply(s *Scope, args List) E {
	if len(args) == 1 {
		return x[args[0].Eval(s).(Int)]
	} else if len(args) == 2 {
		x[args[0].Eval(s).(Int)] = args[1].Eval(s)
	}
	return nil
}
