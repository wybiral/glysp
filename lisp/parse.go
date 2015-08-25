package lisp

import (
	"strconv"
	"strings"
)

type parserStruct struct {
	code   string
	tokens []interface{}
	buffer string
	index  int
}

func typify(token string) E {
	x, err := strconv.Atoi(token)
	if err == nil {
		return Int(x)
	}
	if strings.Contains(token, ".") {
		return DottedSymbol(strings.Split(token, "."))
	}
	return Symbol(token)
}

func (p *parserStruct) push(t interface{}) {
	p.tokens = append(p.tokens, t)
}

func (p *parserStruct) flush() {
	if len(p.buffer) > 0 {
		p.push(typify(p.buffer))
		p.buffer = ""
	}
}

func (p *parserStruct) parseString(quote string) {
	p.flush()
	p.index += 1
	for p.index < len(p.code) {
		c := string(p.code[p.index])
		if c == quote {
			break
		} else if c == "\\" {
			p.index += 1
			c = string(p.code[p.index])
			switch c {
			case "n":
				c = "\n"
			case "r":
				c = "\r"
			case "a":
				c = "\a"
			case "t":
				c = "\t"
			}
			p.buffer += c
		} else {
			p.buffer += c
		}
		p.index += 1
	}
	p.push(String(p.buffer))
	p.buffer = ""
}

func parseTokens(code string) []interface{} {
	p := new(parserStruct)
	p.code = code
	p.tokens = make([]interface{}, 0)
	p.buffer = ""
	p.index = 0
	for p.index < len(p.code) {
		c := string(p.code[p.index])
		if strings.Contains("()[]{}", c) {
			p.flush()
			p.push(c)
		} else if c == "\"" {
			p.parseString("\"")
		} else if c == "'" {
			p.parseString("'")
		} else if c == "`" {
			p.parseString("`")
		} else if c == ":" {
			p.flush()
			p.push(c)
		} else if strings.Contains(" \n\t,", c) {
			p.flush()
		} else {
			p.buffer += c
		}
		p.index++
	}
	p.flush()
	return p.tokens
}

func buildRecurse(tokens []interface{}, index int) (E, int) {
	out := make(List, 0)
	mapping := false
	for index < len(tokens) {
		token := tokens[index]
		if token == "(" {
			x, next := buildRecurse(tokens, index+1)
			index = next
			if mapping {
				out[len(out)-1] = List{Symbol("list"), out[len(out)-1], x}
				mapping = false
			} else {
				out = append(out, x)
			}
		} else if token == ")" {
			break
		} else if token == "[" {
			x, next := buildRecurse(tokens, index+1)
			x = append(List{Symbol("list")}, x.(List)...)
			index = next
			if mapping {
				out[len(out)-1] = List{Symbol("list"), out[len(out)-1], x}
				mapping = false
			} else {
				out = append(out, x)
			}
		} else if token == "]" {
			break
		} else if token == "{" {
			x, next := buildRecurse(tokens, index+1)
			x = append(List{Symbol("hash")}, x.(List)...)
			index = next
			if mapping {
				out[len(out)-1] = List{Symbol("list"), out[len(out)-1], x}
				mapping = false
			} else {
				out = append(out, x)
			}
		} else if token == "}" {
			break
		} else if token == ":" {
			mapping = true
		} else {
			if mapping {
				out[len(out)-1] = List{Symbol("list"), out[len(out)-1], token.(E)}
				mapping = false
			} else {
				out = append(out, token.(E))
			}
		}
		index++
	}
	return out, index
}

func buildCode(tokens []interface{}) E {
	x, _ := buildRecurse(tokens, 0)
	return x
}

func Parse(code string) List {
	return buildCode(parseTokens(code)).(List)
}
