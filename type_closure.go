package lisp

type Closure struct {
	Local *Scope
	Args  List
	Body  List
}

func (x *Closure) Apply(s *Scope, args List) T {
	var out T
	local := &Scope{x.Local, make(map[Symbol]T)}
	for i, v := range x.Args {
		local.scope[v.(Symbol)] = Eval(s, args[i])
	}
	for _, v := range x.Body {
		out = Eval(local, v)
	}
	return out
}
