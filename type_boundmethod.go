package lisp

import (
	"fmt"
)

type BoundMethod struct {
	method   T
	instance T
}

func (x *BoundMethod) String() string {
	return fmt.Sprintf("%v", x.method)
}

func (x *BoundMethod) Apply(s *Scope, args List) T {
	selfargs := append(List{x.instance}, args...)
	return Apply(s, x.method, selfargs)
}
