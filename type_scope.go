package lisp

type Scope struct {
	parent *Scope
	scope  map[Symbol]E
}

func NewScope(parent *Scope) *Scope {
	return &Scope{parent, make(map[Symbol]E)}
}

func (s *Scope) Get(key Symbol) E {
	var out E
	for s != nil {
		out = s.scope[key]
		if out != nil {
			goto found
		}
		s = s.parent
	}
found:
	return out
}

func (s *Scope) Set(key Symbol, value E) {
	s.scope[key] = value
}
