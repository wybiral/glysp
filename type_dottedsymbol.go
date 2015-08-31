package lisp

import (
	"strings"
)

type DottedSymbol []string

func (x DottedSymbol) String() string {
	return strings.Join(x, ".")
}

func (x DottedSymbol) Apply(s *Scope, args List) T {
	last := len(x) - 1
	obj := Eval(s, Symbol(x[0]))
	for _, v := range x[1:last] {
		obj = obj.(GetAttrable).GetAttr(Symbol(v))
	}
	obj.(SetAttrable).SetAttr(Symbol(x[last]), Eval(s, args[0]))
	return nil
}

func (x DottedSymbol) Eval(s *Scope) T {
	out := Symbol(x[0]).Eval(s)
	obj := out.(GetAttrable)
	for _, v := range x[1:] {
		obj := out.(GetAttrable)
		out = obj.GetAttr(Symbol(v))
	}
	switch v := out.(type) {
	case *Method:
		return &BoundMethod{v, obj}
	}
	return out
}
