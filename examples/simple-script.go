package main

import (
	"github.com/wybiral/glysp"
)

func main() {
	// Create new Runtime object
	runtime := lisp.NewRuntime()

	// Execute script
	runtime.EvalScript("simple-script.lisp")
}
