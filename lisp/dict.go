package lisp

import (
	"fmt"
)

type Dict map[E]E

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

func (x Dict) Eval(s *Scope) E { return x }

func (x Dict) Apply(s *Scope, args List) E {
	if len(args) == 1 {
		return x[args[0].Eval(s)]
	} else if len(args) == 2 {
		x[args[0].Eval(s)] = args[1].Eval(s)
	}
	return nil
}
