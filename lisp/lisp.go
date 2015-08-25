package lisp

import (
	"fmt"
	"os"
)

// All objects in the language are required to have a Eval method.
type E interface {
	Eval(*Scope) E
}

func FatalError(x string) {
	fmt.Println(x)
	os.Exit(1)
}

func Repr(x E) E {
	switch v := x.(type) {
	case interface {
		Repr() E
	}:
		return v.Repr()
	default:
		return String(fmt.Sprintf("%v", x))
	}
}
