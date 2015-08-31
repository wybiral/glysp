package lisp

// Lisp Int type
type Int int64

func (x Int) Int() Int {
	return x
}

func (x Int) Float() Float {
	return Float(x)
}

func (x Int) String() String {
	return String(x)
}

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

func (x Int) Apply(s *Scope, args List) T {
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

func (x Int) Add(y T) T { return x + y.(Int) }
func (x Int) Sub(y T) T { return x - y.(Int) }
func (x Int) Mul(y T) T { return x * y.(Int) }
func (x Int) Div(y T) T { return x / y.(Int) }

func (x Int) Lt(y T) T { return Bool(x < y.(Int)) }
func (x Int) Eq(y T) T { return Bool(x == y.(Int)) }
func (x Int) Gt(y T) T { return Bool(x > y.(Int)) }
