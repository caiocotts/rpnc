package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func main() {
	const numberOfLevelsToDisplay = 10

	calc := NewCalculator()

	display := NewDisplay()
	display.Init()

	inputBuffer := Stack[string]{}

	display.PrintStack(calc, numberOfLevelsToDisplay)

	for {
		e := display.PollEvent()
		switch ev := e.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlC:
				closeApplication(display, calc)
				return
			case tcell.KeyCtrlD:
				err := calc.Enter("drop")
				if err != nil {
					display.PrintMessage(err.Error())
					break
				}
				display.PrintMessage("---> drop")
				display.PrintStack(calc, numberOfLevelsToDisplay)
			case tcell.KeyRune:
				character := ev.Rune()
				display.InputCharacter(character)
				inputBuffer.Push(string(character))
			case tcell.KeyEnter:
				value := strings.Join(inputBuffer.ToSlice(), "")
				display.ClearLine(0)
				err := calc.Enter(value)
				if err != nil {
					display.PrintMessage(err.Error())
				} else {
					if value == "" {
						value = "dup"
					}
					display.PrintMessage("---> " + value)
				}
				inputBuffer = Stack[string]{}
				display.ClearInput()
				display.PrintStack(calc, numberOfLevelsToDisplay)
			case tcell.KeyBackspace:
				display.Backspace()
				inputBuffer.Pop()
			}
		}
	}
}

func closeApplication(display Display, calc Calculator) {
	display.Close()
	val, _ := calc.Stack.Pop()
	if val != "" {
		fmt.Println(val)
	}
}
