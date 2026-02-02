package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

const numberOfLevelsToDisplay = 10

func main() {
	calc := NewCalculator()

	display := NewDisplay()
	display.Init()

	inputBuffer := Stack[string]{}

	display.PrintMessage("====(rpnc)====")
	display.PrintStack(calc, numberOfLevelsToDisplay, true)

	for {
		e := display.PollEvent()
		switch ev := e.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				typeKeyIntoInputField(ev.Rune(), display, &inputBuffer)
			case tcell.KeyEnter:
				enterValueIntoCalculator(display, &calc, &inputBuffer)
			case tcell.KeyBackspace:
				display.Backspace()
				inputBuffer.Pop()
			case tcell.KeyCtrlC:
				closeApplication(display, calc)
				return
			case tcell.KeyCtrlD:
				dropAValueFromTheStack(display, &calc)
			}
		}
	}
}

func dropAValueFromTheStack(d Display, calc *Calculator) {
	err := calc.Enter("drop")
	if err != nil {
		d.PrintMessage(err.Error())
		return
	}
	d.PrintMessage("---> drop")
	d.PrintStack(*calc, numberOfLevelsToDisplay, false)
}

func enterValueIntoCalculator(d Display, calc *Calculator, inputBuffer *Stack[string]) {
	value := strings.Join(inputBuffer.ToSlice(), "")
	d.ClearLine(0)
	err := calc.Enter(value)
	if err != nil {
		d.PrintMessage(err.Error())
	} else {
		if value == "" {
			value = "dup"
		}
		d.PrintMessage("---> " + value)
	}
	inputBuffer.Clear()
	d.ClearInput()
	d.PrintStack(*calc, numberOfLevelsToDisplay, false)
}

func typeKeyIntoInputField(c rune, d Display, inputBuffer *Stack[string]) {
	d.InputCharacter(c)
	inputBuffer.Push(string(c))
}

func closeApplication(d Display, calc Calculator) {
	d.Close()
	val, _ := calc.Stack.Pop()
	if val != "" {
		fmt.Println(val)
	}
}
