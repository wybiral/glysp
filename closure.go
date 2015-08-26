package lisp

type Closure struct {
	Local *Scope
	Args  List
	Body  List
}

func (x *Closure) Eval(s *Scope) E { return x }

func (x *Closure) Apply(s *Scope, args List) E {
	var out E
	local := &Scope{x.Local, make(map[Symbol]E)}
	for i, v := range x.Args {
		local.scope[v.(Symbol)] = args[i].Eval(s)
	}
	for _, v := range x.Body {
		out = v.Eval(local)
	}
	return out
}
