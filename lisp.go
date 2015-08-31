package lisp

import (
	"fmt"
	"io/ioutil"
	"os"
)

/* Just a shorthand for empty interface */
type T interface{}

type Evalable interface {
	Eval(*Scope) T
}

type Applyable interface {
	Apply(*Scope, List) T
}

type GetAttrable interface {
	GetAttr(Symbol) T
}

type SetAttrable interface {
	SetAttr(Symbol, T)
}

type Intable interface {
	Int() Int
}

type Floatable interface {
	Float() Float
}

type Stringable interface {
	String() String
}

type Iterable interface {
	Iter() Chan
}

/*
TODO: Actual error handling and reporting :)
*/
func FatalError(x string) {
	fmt.Println(x)
	os.Exit(1)
}

func Eval(s *Scope, x T) T {
	switch v := x.(type) {
	case Evalable:
		return v.Eval(s)
	}
	return x
}

func Apply(s *Scope, x T, args List) T {
	switch v := x.(type) {
	case Applyable:
		return v.Apply(s, args)
	}
	panic(fmt.Sprintf("No apply method for %s", Repr(x)))
}

/*
Return representation as String
*/
func Repr(x T) T {
	switch v := x.(type) {
	case interface {
		Repr() T
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
func (x *Runtime) Eval(code string) T {
	body := Parse(code)
	body = append(List{Symbol("do")}, body...)
	return Eval(x.Global, body)
}

/*
Evaluate a script file as code in runtime.
*/
func (x *Runtime) EvalScript(filename string) T {
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error loading script")
	}
	return x.Eval(string(code))
}

func (x *Runtime) InstallFunc(symbol string, value T) {
	var fn Func
	switch v := value.(type) {
	case func(...T) T:
		fn = Func(v)
	case func(T) T:
		fn = Func(func(a ...T) T { return v(a[0]) })
	case func(T, T) T:
		fn = Func(func(a ...T) T { return v(a[0], a[1]) })
	case func(T, T, T) T:
		fn = Func(func(a ...T) T { return v(a[0], a[1], a[2]) })
	case func(T, T, T, T) T:
		fn = Func(func(a ...T) T { return v(a[0], a[1], a[2], a[3]) })
	default:
		panic("Unable to install: " + symbol)
	}
	x.Global.Set(Symbol(symbol), fn)
}
