package main

import (
	"./lisp"
	"flag"
	"io/ioutil"
)

func main() {

	flag.Parse()
	filename := flag.Args()[0]
	code, err := ioutil.ReadFile(filename)

	if err != nil {
		panic("Error loading script")
	}

	body := lisp.Parse(string(code))

	body = append(lisp.List{lisp.Symbol("do")}, body...)

	s := lisp.NewScope(nil)

	lisp.InstallBuiltins(s)
	_ = body.Eval(s)
}
