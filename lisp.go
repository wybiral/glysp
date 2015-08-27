package lisp

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
E is the universal interface for lisp objects. All objects must support Eval.

TODO: E is ambiguous, should be Evalable
*/
type E interface {
	Eval(*Scope) E
}

type Evalable interface {
	Eval(*Scope) E
}

type Applyable interface {
	Apply(*Scope, List) E
}

type GetAttrable interface {
	GetAttr(Symbol) E
}

type SetAttrable interface {
	SetAttr(Symbol, E)
}

/*
TODO: Actual error handling and reporting :)
*/
func FatalError(x string) {
	fmt.Println(x)
	os.Exit(1)
}

/*
Return representation as String
*/
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

/*
Runtime objects manage execution of code and global scope.
*/
type Runtime struct {
	Global *Scope
}

/*
Instantiate a new Runtime and preload global with builtins.
*/
func NewRuntime() *Runtime {
	x := &Runtime{NewScope(nil)}
	InstallBuiltins(x.Global)
	return x
}

/*
Evaluate some string as code in runtime.
*/
func (x *Runtime) Eval(code string) E {
	body := Parse(code)
	body = append(List{Symbol("do")}, body...)
	return body.Eval(x.Global)
}

/*
Evaluate a script file as code in runtime.
*/
func (x *Runtime) EvalScript(filename string) E {
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error loading script")
	}
	return x.Eval(string(code))
}

func (x *Runtime) InstallFunc(symbol string, fn func(List) E) {
	x.Global.Set(Symbol(symbol), Func(fn))
}
