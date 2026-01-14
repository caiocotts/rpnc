package main

import (
	"errors"
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
	stack.Push(strconv.FormatFloat(numbers[1]+numbers[0], 'g', -1, 64))
	return nil
}

func Subtract(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("- error: " + err.Error())
	}
	stack.Push(strconv.FormatFloat(numbers[1]-numbers[0], 'g', -1, 64))
	return nil
}

func Multiply(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("* error: " + err.Error())
	}
	stack.Push(strconv.FormatFloat(numbers[1]*numbers[0], 'g', -1, 64))
	return nil
}

func Divide(stack *Stack[string]) error {
	numbers, err := pullFromStack(stack, 2)
	if err != nil {
		return errors.New("/ error: " + err.Error())
	}
	stack.Push(strconv.FormatFloat(numbers[1]/numbers[0], 'g', -1, 64))
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

	stack.Push(strconv.FormatFloat(numbers[0], 'g', -1, 64))
	stack.Push(strconv.FormatFloat(numbers[1], 'g', -1, 64))
	return nil
}

func pullFromStack(stack *Stack[string], numberOfValues int) ([]float64, error) {
	if stack.Size() < numberOfValues {
		return nil, errors.New(TooFewArguments)
	}
	numbers := make([]float64, 0, numberOfValues)
	for i := 0; i < numberOfValues; i++ {
		value, err := stack.Pop()
		if err != nil {
			return nil, err
		}
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}
