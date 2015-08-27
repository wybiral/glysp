package lisp

import (
	"strings"
)

type DottedSymbol []string

func (x DottedSymbol) String() string {
	return strings.Join(x, ".")
}

func (x DottedSymbol) Apply(s *Scope, args List) E {
	last := len(x) - 1
	obj := Symbol(x[0]).Eval(s)
	for _, v := range x[1:last] {
		obj = obj.(GetAttrable).GetAttr(Symbol(v))
	}
	obj.(SetAttrable).SetAttr(Symbol(x[last]), args[0].Eval(s))
	return nil
}

func (x DottedSymbol) Eval(s *Scope) E {
	out := Symbol(x[0]).Eval(s)
	obj := out.(GetAttrable)
	for _, v := range x[1:] {
		obj := out.(GetAttrable)
		out = obj.GetAttr(Symbol(v))
	}
	switch v := out.(type) {
	case *Method:
		return &BoundMethod{v, obj.(*Object)}
	}
	return out
}
