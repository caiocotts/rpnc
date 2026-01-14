package main

import (
	"strconv"
)

var functions = map[string]func(stack *Stack[string]) error{
	"+":     Add,
	"-":     Subtract,
	"*":     Multiply,
	"/":     Divide,
	"drop":  Drop,
	"dup":   Dup,
	"clear": Clear,
	"swap":  Swap,
}

type Calculator struct {
	Stack Stack[string]
}

func NewCalculator() Calculator {
	return Calculator{
		Stack: Stack[string]{},
	}
}
func (c *Calculator) Enter(value string) error {
	if value == "" {
		value = "dup"
	}
	if f, ok := functions[value]; ok {
		err := f(&c.Stack)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	c.Stack.Push(value)
	return nil
}
