package main

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Var string

func (v Var) String() string {
	return string(v)
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

type literal float64

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

type unary struct {
	op rune
	x  Expr
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

type abs struct {
	args []Expr
}

func (a abs) Eval(env Env) float64 {
	return math.Abs(a.args[0].Eval(env))
}

func (a abs) Check(vars map[Var]bool) error {
	if len(a.args) != 1 {
		return fmt.Errorf("call to abs has %d args, want %d", len(a.args), 1)
	}
	for _, arg := range a.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (a abs) String() string {
	return fmt.Sprintf("abs(%s)", a.args[0])
}

type binary struct {
	op   rune
	x, y Expr
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x, string(b.op), b.y)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+=*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

type call struct {
	fn   string
	args []Expr
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("pow(%s, %s)", c.args[0], c.args[1])
	case "sin":
		return fmt.Sprintf("sin(%s)", c.args[0])
	case "sqrt":
		return fmt.Sprintf("sqrt(%s)", c.args[0])
	}
	return ""
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

type Env map[Var]float64