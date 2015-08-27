package lisp

import (
	"fmt"
	"reflect"
)

func MacroClassMethod(s *Scope, args List) E {
	name := args[0].(Symbol)
	closure := MacroFn(s, args[1:])
	s.scope[name] = &Method{String(name), closure.(*Closure)}
	return nil
}

func MacroClass(s *Scope, args List) E {
	s_def := NewScope(s)
	s_def.scope[Symbol("method")] = Macro(MacroClassMethod)
	s_cls := NewScope(s_def)
	name := args[0].(Symbol)
	class := &Class{nil, String(name), s_cls.scope}
	s.scope[name] = class
	for _, v := range args[1:] {
		v.Eval(s_cls)
	}
	return class
}

func MacroDo(s *Scope, body List) E {
	var out E
	for _, v := range body {
		out = v.Eval(s)
	}
	return out
}

func MacroFor(s *Scope, args List) E {
	var out E
	pair := args[0].(List)
	name := pair[0].(Symbol)
	iter := pair[1].Eval(s)
	body := args[1:]
	for x := range iter.(interface {
		Iter() Chan
	}).Iter() {
		s.scope[name] = x
		for _, expr := range body {
			out = expr.Eval(s)
		}
	}
	return out
}

func MacroWhile(s *Scope, args List) E {
	var out E
	cond := args[0]
	body := args[1:]
	for cond.Eval(s).(Bool) {
		for _, expr := range body {
			out = expr.Eval(s)
		}
	}
	return out
}

func MacroGo(s *Scope, args List) E {
	go func() {
		MacroDo(s, args)
	}()
	return nil
}

func MacroIf(s *Scope, args List) E {
	if bool(args[0].Eval(s).(Bool)) {
		return args[1].Eval(s)
	}
	return args[2].Eval(s)
}

func MacroSet(s *Scope, args List) E {
	obj := args[0].(interface {
		Apply(*Scope, List) E
	})
	obj.Apply(s, args[1:])
	return nil
}

func MacroEval(s *Scope, args List) E {
	var out E
	for _, v := range args {
		out = v.Eval(s).Eval(s)
	}
	return out
}

func MacroFn(s *Scope, args List) E {
	return &Closure{s, args[0].(List), args[1:]}
}

func MacroFunc(s *Scope, args List) E {
	x := &Closure{s, args[1].(List), args[2:]}
	s.scope[args[0].(Symbol)] = x
	return x
}

func FuncAdd(args List) E {
	out := args[0]
	for _, v := range args[1:] {
		out = out.(interface {
			Add(E) E
		}).Add(v)
	}
	return out
}

func FuncSub(args List) E {
	return args[0].(interface {
		Sub(E) E
	}).Sub(args[1])
}

func FuncMul(args List) E {
	return args[0].(interface {
		Mul(E) E
	}).Mul(args[1])
}

func FuncDiv(args List) E {
	return args[0].(interface {
		Div(E) E
	}).Div(args[1])
}

func FuncLt(args List) E {
	return args[0].(interface {
		Lt(E) E
	}).Lt(args[1])
}

func FuncEq(args List) E {
	return args[0].(interface {
		Eq(E) E
	}).Eq(args[1])
}

func FuncGt(args List) E {
	return args[0].(interface {
		Gt(E) E
	}).Gt(args[1])
}

func FuncList(args List) E {
	return args
}

func FuncDict(args List) E {
	out := make(Dict)
	for _, v := range args {
		pair := v.(List)
		out[pair[0]] = pair[1]
	}
	return out
}

func FuncPrint(args List) E {
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

func FuncStr(args List) E {
	return String(fmt.Sprintf("%v", args[0]))
}

func FuncRepr(args List) E {
	return Repr(args[0])
}

func FuncType(args List) E {
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

func FuncChan(args List) E {
	return make(Chan)
}

func FuncClose(args List) E {
	close(args[0].(Chan))
	return nil
}

func MacroSetAttr(s *Scope, args List) E {
	obj := args[0].Eval(s).(SetAttrable)
	key := args[1].(Symbol)
	val := args[2].Eval(s)
	obj.SetAttr(key, val)
	return val
}

func MacroGetAttr(s *Scope, args List) E {
	obj := args[0].Eval(s).(GetAttrable)
	key := args[1].(Symbol)
	return obj.GetAttr(key)
}

func MacroQuote(s *Scope, args List) E {
	return args[0]
}

func MacroParse(s *Scope, args List) E {
	arg := args[0].Eval(s).(String)
	out := Parse(string(arg))
	return out[0].Eval(s)
}

func InstallBuiltins(s *Scope) {

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
	s.Set(Symbol("while"), Macro(MacroWhile))

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
