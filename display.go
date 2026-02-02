package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

const ( //special runes
	backspace = '\b'
	reset     = 0
)

const (
	inputLinePosition = 11
)

var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

type Display struct {
	screen tcell.Screen
	output chan rune
}

func NewDisplay() Display {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	s.SetStyle(defStyle)

	return Display{
		screen: s,
		output: make(chan rune),
	}
}
func (d *Display) Init() {
	if err := d.screen.Init(); err != nil {
		log.Fatal(err)
	}
	d.initInput()
}

func (d *Display) initInput() {
	go func() {
		cursorVisible := false
		cursorPosition := 0
		for {
			select {
			case <-time.After(500 * time.Millisecond):
			case output := <-d.output:
				cursorVisible = true
				switch output {
				case reset:
					cursorPosition = 0
				case backspace:
					if cursorPosition >= 1 {
						d.screen.PutStr(cursorPosition-1, inputLinePosition, " ")
						cursorPosition--
					}
				default:
					d.screen.PutStr(cursorPosition, inputLinePosition, string(output))
					cursorPosition++
				}
			}
			if cursorVisible {
				d.screen.ShowCursor(cursorPosition, inputLinePosition)
			} else {
				d.screen.HideCursor()
			}
			cursorVisible = !cursorVisible
			d.screen.Show()
		}
	}()
}

func (d *Display) Close() {
	d.screen.Fini()
}

func (d *Display) PrintStack(calculator Calculator, numberOfLevels int, refresh bool) {
	for i := range numberOfLevels {
		y := numberOfLevels - i
		stackLevel := i + 1
		if refresh {
			d.ClearLine(y)
			d.screen.PutStr(0, y, fmt.Sprintf("%d:", stackLevel))
		} else {
			width, _ := d.screen.Size()
			d.ClearRangeInLine(y, 5, width)
		}
		val, err := calculator.Stack.Pop()
		if err == nil {
			d.screen.PutStr(5, y, val)
		}
	}
	d.screen.Show()
}

func (d *Display) PrintMessage(message string) {
	d.ClearLine(0)
	d.screen.PutStr(0, 0, message)
	d.screen.Show()
}

func (d *Display) ClearRangeInLine(y, xStart, xEnd int) {
	for x := xStart; x <= xEnd; x++ {
		d.screen.PutStr(x, y, " ")
	}
	d.screen.Show()
}

func (d *Display) ClearLine(y int) {
	width, _ := d.screen.Size()
	for x := range width {
		d.screen.PutStr(x, y, " ")
	}
	d.screen.Show()
}

func (d *Display) PollEvent() tcell.Event {
	return d.screen.PollEvent()
}

func (d *Display) InputCharacter(character rune) {
	d.output <- character
}

func (d *Display) Backspace() {
	d.output <- backspace
}

func (d *Display) ClearInput() {
	d.ClearLine(inputLinePosition)
	d.output <- reset
}
