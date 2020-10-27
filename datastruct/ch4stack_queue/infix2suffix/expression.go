// Author: xufei
// Date: 2019-07-29

package infix2suffix

import (
	"fmt"
)

func New(exps []string) *Expression {
	return &Expression{
		exps: exps,
		ans:  make([]string, 0),
		link: InitStack(),
	}
}

func (e *Expression) Infix2Suffix() ([]string, error) {
	for _, value := range e.exps {
		switch value {
		case "+":
			if err := e.add(value); err != nil {
				return nil, err
			}
		case "-":
			if err := e.sub(value); err != nil {
				return nil, err
			}
		case "*":
			if err := e.mul(value); err != nil {
				return nil, err
			}
		case "/":
			if err := e.div(value); err != nil {
				return nil, err
			}
		case "(":
			if !e.link.Push("(") {
				return nil, fmt.Errorf("expression error")
			}
		case ")":
			for {
				var val string
				if !e.link.Pop(&val) {
					return nil, fmt.Errorf("expression error")
				}
				if val == "(" {
					break
				}
				e.ans = append(e.ans, val)

				if e.link.Count() == 0 {
					break
				}
			}
		default:
			e.ans = append(e.ans, value)
		}
	}

	for {
		if e.link.Count() == 0 {
			break
		}

		var val string
		if !e.link.Pop(&val) {
			return nil, fmt.Errorf("expression error")
		}
		e.ans = append(e.ans, val)
	}

	return e.ans, nil
}

func (e *Expression) add(symbol string) error {
	for {
		if e.link.Count() == 0 || e.link.Top() == "(" {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		var val string
		if !e.link.Pop(&val) {
			return fmt.Errorf("expression error")
		}
		e.ans = append(e.ans, val)
	}
	return nil
}

func (e *Expression) sub(symbol string) error {
	for {
		if e.link.Count() == 0 || e.link.Top() == "(" {
			if !e.link.Push("-") {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Top() == "-" || e.link.Top() == "*" || e.link.Top() == "/" {
			var val string
			if !e.link.Pop(&val) {
				return fmt.Errorf("expression error")
			}
			e.ans = append(e.ans, val)
		} else {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Count() == 0 {
			break
		}
	}
	return nil
}

func (e *Expression) mul(symbol string) error {
	for {
		if e.link.Count() == 0 || e.link.Top() == "(" {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Top() == "*" || e.link.Top() == "/" {
			var val string
			if !e.link.Pop(&val) {
				return fmt.Errorf("expression error")
			}
			e.ans = append(e.ans, val)
		} else {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Count() == 0 {
			break
		}
	}
	return nil
}

func (e *Expression) div(symbol string) error {
	for {
		if e.link.Count() == 0 || e.link.Top() == "(" {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Top() == "/" {
			var val string
			if !e.link.Pop(&val) {
				return fmt.Errorf("expression error")
			}
			e.ans = append(e.ans, val)
		} else {
			if !e.link.Push(symbol) {
				return fmt.Errorf("expression error")
			}
			break
		}

		if e.link.Count() == 0 {
			break
		}
	}
	return nil
}
