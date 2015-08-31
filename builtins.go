package lisp

import (
	"fmt"
	"reflect"
)

func MacroClassMethod(s *Scope, args List) T {
	name := args[0].(Symbol)
	closure := MacroFn(s, args[1:])
	s.scope[name] = &Method{String(name), closure.(*Closure)}
	return nil
}

func MacroClass(s *Scope, args List) T {
	s_def := NewScope(s)
	s_def.scope[Symbol("method")] = Macro(MacroClassMethod)
	s_cls := NewScope(s_def)
	name := args[0].(Symbol)
	class := &Class{nil, String(name), s_cls.scope}
	s.scope[name] = class
	for _, v := range args[1:] {
		Eval(s_cls, v)
	}
	return class
}

func MacroDo(s *Scope, body List) T {
	var out T
	for _, v := range body {
		out = Eval(s, v)
	}
	return out
}

func MacroFor(s *Scope, args List) T {
	var out T
	pair := args[0].(List)
	name := pair[0].(Symbol)
	iter := Eval(s, pair[1])
	body := args[1:]
	for x := range iter.(interface {
		Iter() Chan
	}).Iter() {
		s.scope[name] = x
		for _, expr := range body {
			out = Eval(s, expr)
		}
	}
	return out
}

func MacroGo(s *Scope, args List) T {
	go func() {
		MacroDo(s, args)
	}()
	return nil
}

func MacroIf(s *Scope, args List) T {
	if bool(Eval(s, args[0]).(Bool)) {
		return Eval(s, args[1])
	}
	return Eval(s, args[2])
}

func MacroSet(s *Scope, args List) T {
	obj := args[0].(interface {
		Apply(*Scope, List) T
	})
	obj.Apply(s, args[1:])
	return nil
}

func MacroEval(s *Scope, args List) T {
	var out T
	for _, v := range args {
		out = Eval(s, Eval(s, v))
	}
	return out
}

func MacroFn(s *Scope, args List) T {
	return &Closure{s, args[0].(List), args[1:]}
}

func MacroFunc(s *Scope, args List) T {
	x := &Closure{s, args[1].(List), args[2:]}
	s.scope[args[0].(Symbol)] = x
	return x
}

func FuncAdd(args ...T) T {
	out := args[0]
	for _, v := range args[1:] {
		out = out.(interface {
			Add(T) T
		}).Add(v)
	}
	return out
}

func FuncSub(args ...T) T {
	return args[0].(interface {
		Sub(T) T
	}).Sub(args[1])
}

func FuncMul(args ...T) T {
	return args[0].(interface {
		Mul(T) T
	}).Mul(args[1])
}

func FuncDiv(args ...T) T {
	return args[0].(interface {
		Div(T) T
	}).Div(args[1])
}

func FuncLt(args ...T) T {
	return args[0].(interface {
		Lt(T) T
	}).Lt(args[1])
}

func FuncEq(args ...T) T {
	return args[0].(interface {
		Eq(T) T
	}).Eq(args[1])
}

func FuncGt(args ...T) T {
	return args[0].(interface {
		Gt(T) T
	}).Gt(args[1])
}

func FuncList(args ...T) T {
	return List(args)
}

func FuncDict(args ...T) T {
	out := make(Dict)
	for _, v := range args {
		pair := v.(List)
		out[pair[0]] = pair[1]
	}
	return out
}

func FuncPrint(args ...T) T {
	line := ""
	for _, v := range args {
		line += fmt.Sprintf("%v", v)
		line += " "
	}
	if len(args) > 0 {
		line = line[:len(line)-1]
	}
	fmt.Println(line)
	return nil
}

func FuncStr(args ...T) T {
	return String(fmt.Sprintf("%v", args[0]))
}

func FuncRepr(args ...T) T {
	return Repr(args[0])
}

func FuncType(args ...T) T {
	var out String
	switch v := args[0].(type) {
	case *Object:
		out = v.class.name
		break
	case *Closure:
		out = String("Closure")
		break
	case List:
		out = String("List")
		break
	case Dict:
		out = String("Dict")
		break
	default:
		out = String(reflect.TypeOf(args[0]).Name())
	}
	return out
}

func FuncChan(args ...T) T {
	return make(Chan)
}

func FuncClose(args ...T) T {
	close(args[0].(Chan))
	return nil
}

func MacroSetAttr(s *Scope, args List) T {
	obj := Eval(s, args[0]).(SetAttrable)
	key := args[1].(Symbol)
	val := Eval(s, args[2])
	obj.SetAttr(key, val)
	return val
}

func MacroGetAttr(s *Scope, args List) T {
	obj := Eval(s, args[0]).(GetAttrable)
	key := args[1].(Symbol)
	return obj.GetAttr(key)
}

func MacroQuote(s *Scope, args List) T {
	return args[0]
}

func MacroParse(s *Scope, args List) T {
	arg := Eval(s, args[0]).(String)
	out := Parse(string(arg))
	return Eval(s, out[0])
}

func InstallBuiltins(s *Scope) {

	s.Set(Symbol("nil"), nil)
	s.Set(Symbol("true"), Bool(true))
	s.Set(Symbol("false"), Bool(false))

	s.Set(Symbol("parse"), Macro(MacroParse))
	s.Set(Symbol("quote"), Macro(MacroQuote))
	s.Set(Symbol("eval"), Macro(MacroEval))

	s.Set(Symbol("class"), Macro(MacroClass))
	s.Set(Symbol("setattr"), Macro(MacroSetAttr))
	s.Set(Symbol("getattr"), Macro(MacroGetAttr))

	s.Set(Symbol("set"), Macro(MacroSet))

	s.Set(Symbol("do"), Macro(MacroDo))
	s.Set(Symbol("if"), Macro(MacroIf))
	s.Set(Symbol("for"), Macro(MacroFor))

	s.Set(Symbol("go"), Macro(MacroGo))
	s.Set(Symbol("chan"), Func(FuncChan))
	s.Set(Symbol("close"), Func(FuncClose))

	s.Set(Symbol("fn"), Macro(MacroFn))
	s.Set(Symbol("func"), Macro(MacroFunc))

	s.Set(Symbol("str"), Func(FuncStr))
	s.Set(Symbol("print"), Func(FuncPrint))

	s.Set(Symbol("+"), Func(FuncAdd))
	s.Set(Symbol("-"), Func(FuncSub))
	s.Set(Symbol("*"), Func(FuncMul))
	s.Set(Symbol("/"), Func(FuncDiv))
	s.Set(Symbol("<"), Func(FuncLt))
	s.Set(Symbol("="), Func(FuncEq))
	s.Set(Symbol(">"), Func(FuncGt))

	s.Set(Symbol("list"), Func(FuncList))
	s.Set(Symbol("dict"), Func(FuncDict))

	s.Set(Symbol("type"), Func(FuncType))
}
