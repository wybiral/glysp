package lisp

// Lisp Int type
type Int int64

func (x Int) Eval(s *Scope) E { return x }

func (x Int) Iter() Chan {
	ch := make(Chan)
	go func() {
		for i := Int(0); i < x; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func (x Int) Apply(s *Scope, args List) E {
	ch := make(Chan)
	if len(args) == 0 {
		ch = x.Iter()
	} else if len(args) == 1 {
		y := args[0].(Int)
		go func() {
			for i := Int(x); i < y; i++ {
				ch <- i
			}
			close(ch)
		}()
	} else if len(args) == 2 {
		y := args[0].(Int)
		z := args[1].(Int)
		go func() {
			for i := Int(x); i < y; i += z {
				ch <- i
			}
			close(ch)
		}()
	}
	return ch
}

func (x Int) Add(y E) E { return x + y.(Int) }
func (x Int) Sub(y E) E { return x - y.(Int) }
func (x Int) Mul(y E) E { return x * y.(Int) }
func (x Int) Div(y E) E { return x / y.(Int) }

func (x Int) Lt(y E) E { return Bool(x < y.(Int)) }
func (x Int) Eq(y E) E { return Bool(x == y.(Int)) }
func (x Int) Gt(y E) E { return Bool(x > y.(Int)) }
