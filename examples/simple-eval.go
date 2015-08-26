package main

import (
	"github.com/wybiral/glysp"
)

func main() {
	runtime := lisp.NewRuntime()
	runtime.Eval("(print 'Hello world!')")
}
