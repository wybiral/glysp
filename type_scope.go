package lisp

type Scope struct {
	parent *Scope
	scope  map[Symbol]T
}

func NewScope(parent *Scope) *Scope {
	return &Scope{parent, make(map[Symbol]T)}
}

func (s *Scope) Get(key Symbol) T {
	var out T
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

func (s *Scope) Set(key Symbol, value T) {
	s.scope[key] = value
}
