package main

import (
	"github.com/wybiral/glysp"
)

func main() {
	// Create new Runtime object
	runtime := lisp.NewRuntime()

	// Execute script
	runtime.EvalScript("simple-script2.lisp")

	// Grab object with name "main"
	main := runtime.Global.Get("main")

	// Create argument list
	args := lisp.List{lisp.String("world!!!")}

	// Apply "main" object to args list
	lisp.Apply(runtime.Global, main, args)
}
