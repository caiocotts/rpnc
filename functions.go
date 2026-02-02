package main

import (
	"errors"
	"math"
	"strconv"
)

const ( // error messages
	TooFewArguments = "too few arguments"
	StackEmpty      = "stack is empty"
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
	_, err := stack.Pop()
	if err != nil && err.Error() == StackEmpty {
		return errors.New("drop error: " + TooFewArguments)
	} else if err != nil {
		return errors.New("drop error: " + err.Error())
	}
	return nil
}

func Dup(stack *Stack[string]) error {
	number, err := stack.Pop()
	if err != nil && err.Error() == StackEmpty {
		return errors.New("dup error: " + TooFewArguments)
	} else if err != nil {
		return errors.New("dup error: " + err.Error())
	}
	stack.Push(number)
	stack.Push(number)
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
	stackElements, err := pullFromStack(stack, numOfLevelsToRoll)
	if err != nil {
		return errors.New("roll error: " + err.Error())
	}
	bottomMostElement := stackElements[0]

	for i := 1; i < len(stackElements); i++ {
		stack.Push(stackElements[i])
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
