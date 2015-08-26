package main

import (
	"github.com/wybiral/glysp"
)

func main() {
	runtime := lisp.NewRuntime()
	runtime.EvalScript("simple-script.lisp")
}
