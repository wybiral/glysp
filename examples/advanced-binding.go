package main

import (
	"fmt"
	"github.com/wybiral/glysp"
)

type Matrix struct {
	m, n int
	v    []float64
}

func NewMatrix(args lisp.List) lisp.E {
	m := int(args[0].(lisp.Int))
	n := int(args[1].(lisp.Int))
	return &Matrix{m, n, make([]float64, m*n)}
}

func (x *Matrix) Eval(s *lisp.Scope) lisp.E {
	return x
}

func (x *Matrix) String() string {
	s := ""
	for i := 0; i < x.m; i++ {
		for j := 0; j < x.n; j++ {
			s += fmt.Sprintf("\t%v", x.v[i*x.n+j])
		}
		s += "\n"
	}
	return s
}

func main() {
	runtime := lisp.NewRuntime()
	runtime.InstallFunc("matrix", NewMatrix)
	runtime.Eval("(set x (matrix 3 3)) (print x)")
}
