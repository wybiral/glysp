package lisp

import "fmt"

// Lisp Int type
type Float float64

func ToFloat(x T) Float {
	switch v := x.(type) {
	case Floatable:
		return v.Float()
	}
	panic("Can't convert to float")
}

func (x Float) Int() Int {
	return Int(x)
}

func (x Float) Float() Float {
	return x
}

func (x Float) String() String {
	return String(fmt.Sprintf("%f", x))
}

func (x Float) Add(y T) T { return x + ToFloat(y) }
func (x Float) Sub(y T) T { return x - ToFloat(y) }
func (x Float) Mul(y T) T { return x * ToFloat(y) }

func (x Float) Div(y T) T {
	switch v := y.(type) {
	case Float:
		return x / v
	case Int:
		return x / Float(v)
	}
	return nil
}

func (x Float) Lt(y T) T { return Bool(x < y.(Float)) }
func (x Float) Eq(y T) T { return Bool(x == y.(Float)) }
func (x Float) Gt(y T) T { return Bool(x > y.(Float)) }
