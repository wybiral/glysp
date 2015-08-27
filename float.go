package lisp

// Lisp Int type
type Float float64

func (x Float) Eval(s *Scope) E { return x }

func (x Float) Add(y E) E { return x + y.(Float) }
func (x Float) Sub(y E) E { return x - y.(Float) }
func (x Float) Mul(y E) E { return x * y.(Float) }

func (x Float) Div(y E) E {
	switch v := y.(type) {
	case Float:
		return x / v
	case Int:
		return x / Float(v)
	}
	return nil
}

func (x Float) Lt(y E) E { return Bool(x < y.(Float)) }
func (x Float) Eq(y E) E { return Bool(x == y.(Float)) }
func (x Float) Gt(y E) E { return Bool(x > y.(Float)) }
