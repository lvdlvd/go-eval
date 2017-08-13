package eval

import (
	"errors"
	"fmt"
	"math"
)

var divByZero = errors.New("division by zero")

type Evaluator interface {
	Evaluate(params map[string]float64) (float64, error)
	String() string
}

type constval float64

func (n constval) Evaluate(params map[string]float64) (float64, error) { return float64(n), nil }

func (n constval) String() string { return fmt.Sprintf("=%f", float64(n)) }

type varval string

func (n varval) Evaluate(params map[string]float64) (float64, error) {
	return params[string(n)], nil
}

func (n varval) String() string { return fmt.Sprintf("$%s", string(n)) }

type unop struct{ a Evaluator }

type neg unop

func (n neg) Evaluate(params map[string]float64) (float64, error) {
	a, err := n.a.Evaluate(params)
	return -a, err
}

func (n neg) String() string { return fmt.Sprintf("- %v", n.a) }

type inv unop

func (n inv) Evaluate(params map[string]float64) (float64, error) {
	a, err := n.a.Evaluate(params)
	if a == 0 {
		return math.Inf(1), divByZero
	}
	return 1 / a, err
}

func (n inv) String() string { return fmt.Sprintf("(1/%v)", n.a) }

type binop struct{ a, b Evaluator }

type add binop

func (n add) Evaluate(params map[string]float64) (float64, error) {
	a, err := n.a.Evaluate(params)
	if err != nil {
		return 0, err
	}
	b, err := n.b.Evaluate(params)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

func (n add) String() string { return fmt.Sprintf("%v + %v", n.a, n.b) }

type mul binop

func (n mul) Evaluate(params map[string]float64) (float64, error) {
	a, err := n.a.Evaluate(params)
	if err != nil {
		return 0, err
	}
	b, err := n.b.Evaluate(params)
	if err != nil {
		return 0, err
	}
	return a * b, nil
}

func (n mul) String() string { return fmt.Sprintf("%v x %v", n.a, n.b) }
