package main

import (
	"fmt"
	// "strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	input  string
	tokens []string
	pos    int
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:  input,
		tokens: make([]string, 0),
		pos:    0,
	}
	lexer.tokenize()
	return lexer
}

func (l *Lexer) tokenize() {
	tokens := strings.FieldsFunc(l.input, func(c rune) bool {
		return unicode.IsSpace(c) || c == '(' || c == ')'
	})

	l.tokens = append(l.tokens, tokens...)
}

func (l *Lexer) getNextToken() string {
	if l.pos < len(l.tokens) {
		token := l.tokens[l.pos]
		l.pos++
		return token
	}
	return ""
}

func (l *Lexer) peekNextToken() string {
	if l.pos < len(l.tokens) {
		return l.tokens[l.pos]
	}
	return ""
}

type LispToGoTranspiler struct {
	lexer *Lexer
}

func NewLispToGoTranspiler(input string) *LispToGoTranspiler {
	return &LispToGoTranspiler{
		lexer: NewLexer(input),
	}
}

func (t *LispToGoTranspiler) transpile() string {
	return t.transpileS()
}

func (t *LispToGoTranspiler) transpileS() string {
	token := t.lexer.getNextToken() // Consume "("
    fmt.Println(">>" + token)
	if token != "(" {
		panic("Expected '(' at the beginning")
	}

	token = t.lexer.getNextToken() // Get the next token

	switch token {
	case "eq":
		return t.transpileEq()
	default:
		panic("Unsupported operation: " + token)
	}
}

func (t *LispToGoTranspiler) transpileEq() string {
	left := t.transpileExpr()
	operator := t.lexer.getNextToken() // Consume operator
	right := t.transpileExpr()
	t.lexer.getNextToken() // Consume ")"

	return fmt.Sprintf("%s %s %s", left, operator, right)
}

func (t *LispToGoTranspiler) transpileExpr() string {
	token := t.lexer.peekNextToken()

	if token == "(" {
		return t.transpileS()
	}

	// Atom
	t.lexer.getNextToken() // Consume atom
	return token
}

func main() {
	input := "(eq a (+ 23 23 23))"
	transpiler := NewLispToGoTranspiler(input)
	goCode := transpiler.transpile()

	fmt.Println(goCode)
}

