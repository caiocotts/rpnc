package main

import (
	"errors"
	"math"
	"strconv"
)

const ( // error messages
	TooFewArguments = "too few arguments"
)

func Add(stack *Stack[string]) error {
	numbers, err := pullFromStackAsNumbers(stack, 2)
	if err != nil {
		return errors.New("+ error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] + numbers[1]))
	return nil
}

func Subtract(stack *Stack[string]) error {
	numbers, err := pullFromStackAsNumbers(stack, 2)
	if err != nil {
		return errors.New("- error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] - numbers[1]))
	return nil
}

func Multiply(stack *Stack[string]) error {
	numbers, err := pullFromStackAsNumbers(stack, 2)
	if err != nil {
		return errors.New("* error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] * numbers[1]))
	return nil
}

func Divide(stack *Stack[string]) error {
	numbers, err := pullFromStackAsNumbers(stack, 2)
	if err != nil {
		return errors.New("/ error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] / numbers[1]))
	return nil
}

func Drop(stack *Stack[string]) error {
	if stack.Empty() {
		return errors.New("drop error: " + TooFewArguments)
	}
	_, err := stack.Pop()
	if err != nil {
		return errors.New("drop error: " + err.Error())
	}
	return nil
}

func Dup(stack *Stack[string]) error {
	if stack.Empty() {
		return errors.New("dup error: " + TooFewArguments)
	}
	stack.Push(stack.Peek())
	return nil
}

func Clear(stack *Stack[string]) error {
	stack.Clear()
	return nil
}

func Swap(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("swap error: " + err.Error())
	}
	stack.Push(numbers[1])
	stack.Push(numbers[0])
	return nil
}

func Roll(stack *Stack[string]) error {
	n, err := pullFromStackAsNumbers(stack, 1)
	if err != nil {
		return errors.New("roll error: " + err.Error())
	}
	numOfLevelsToRoll := int(math.Round(n[0]))
	if numOfLevelsToRoll <= 0 {
		return nil
	}
	elements, err := pullFromStack(stack, numOfLevelsToRoll)
	if err != nil {
		return errors.New("roll error: " + err.Error())
	}
	bottomMostElement := elements[0]
	for i := 1; i < len(elements); i++ {
		stack.Push(elements[i])
	}
	stack.Push(bottomMostElement)
	return nil
}

func Rot(stack *Stack[string]) error {
	if stack.Size() < 3 {
		return errors.New("rot error: " + TooFewArguments)
	}
	elements, err := pullFromStack(stack, 3)
	if err != nil {
		return errors.New("rot error: " + err.Error())
	}
	bottomMostElement := elements[0]
	for i := 1; i < 3; i++ {
		stack.Push(elements[i])
	}
	stack.Push(bottomMostElement)
	return nil
}

func floatToString(number float64) string {
	return strconv.FormatFloat(number, 'g', -1, 64)
}

func pullFromStackAsNumbers(stack *Stack[string], numberOfElements int) ([]float64, error) {
	elements, err := pullFromStack(stack, numberOfElements)
	if err != nil {
		return nil, err
	}
	numbers := make([]float64, numberOfElements)
	for i, element := range elements {
		numbers[i], err = strconv.ParseFloat(element, 64)
		if err != nil {
			return nil, err
		}
	}
	return numbers, nil
}

func pullFromStack(stack *Stack[string], numberOfElements int) ([]string, error) {
	if stack.Size() < numberOfElements {
		return nil, errors.New(TooFewArguments)
	}
	elements := stack.ToSlice()[stack.Size()-numberOfElements:]
	for range numberOfElements {
		_, err := stack.Pop()
		if err != nil {
			return nil, err
		}
	}
	return elements, nil
}
