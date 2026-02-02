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
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("+ error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] + numbers[1]))
	return nil
}

func Subtract(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("- error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] - numbers[1]))
	return nil
}

func Multiply(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("* error: " + err.Error())
	}
	stack.Push(floatToString(numbers[0] * numbers[1]))
	return nil
}

func Divide(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
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
	*stack = Stack[string]{}
	return nil
}

func Swap(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("swap error: " + err.Error())
	}

	stack.Push(floatToString(numbers[1]))
	stack.Push(floatToString(numbers[0]))
	return nil
}

func Roll(stack *Stack[string]) error {
	n, err := pullFromStack(stack, 1)
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
		stack.Push(floatToString(stackElements[i]))
	}
	stack.Push(floatToString(bottomMostElement))
	return nil
}

func floatToString(number float64) string {
	return strconv.FormatFloat(number, 'g', -1, 64)
}

func pullFromStack(stack *Stack[string], numberOfValues int) ([]float64, error) {
	if stack.Size() < numberOfValues {
		return nil, errors.New(TooFewArguments)
	}
	numbers := make([]float64, numberOfValues)
	for i := numberOfValues - 1; i >= 0; i-- {
		value, err := stack.Pop()
		if err != nil {
			return nil, err
		}
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}
