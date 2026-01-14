package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func main() {
	calc := NewCalculator()
	display := InitDisplay()
	cursorPosition := 0
	inputBuffer := Stack[string]{}
	const numberOfLevelsToDisplay = 10

	PrintStack(calc, display, numberOfLevelsToDisplay)

	for {
		e := display.PollEvent()
		switch ev := e.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlC:
				CloseDisplay(display)
				val, _ := calc.Stack.Pop()
				if val != "" {
					fmt.Println(val)
				}
				return
			case tcell.KeyCtrlD:
				err := calc.Enter("drop")
				if err != nil {
					PrintMessage(display, err.Error())
					break
				}
				PrintMessage(display, "---> drop")
				PrintStack(calc, display, numberOfLevelsToDisplay)
				break
			case tcell.KeyRune:
				character := string(ev.Rune())
				display.PutStr(cursorPosition, 11, character)
				inputBuffer.Push(character)
				cursorPosition++
				break
			case tcell.KeyEnter:
				value := strings.Join(inputBuffer.ToSlice(), "")
				ClearLine(display, 0)
				err := calc.Enter(value)
				if err != nil {
					PrintMessage(display, err.Error())
				} else {
					if value == "" {
						value = "dup"
					}
					PrintMessage(display, "---> "+value)
				}
				inputBuffer = Stack[string]{}
				cursorPosition = 0
				ClearLine(display, 11)
				PrintStack(calc, display, numberOfLevelsToDisplay)
				break
			case tcell.KeyBackspace:
				if cursorPosition < 1 {
					break
				}
				display.PutStr(cursorPosition-1, 11, " ")
				_, _ = inputBuffer.Pop()
				cursorPosition--
				break
			}
			break
		}
		display.Show()
	}
}
