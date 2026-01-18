package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

type Display struct {
	CursorXCoordinate int
	screen            tcell.Screen
}

func NewDisplay() Display {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	s.SetStyle(defStyle)

	return Display{
		screen: s,
	}
}
func (d *Display) Init() {
	if err := d.screen.Init(); err != nil {
		log.Fatal(err)
	}
}
func (d *Display) Close() {
	d.screen.Fini()
}

func (d *Display) PrintStack(calculator Calculator, numberOfLevels int) {
	for i := range numberOfLevels {
		y := numberOfLevels - i
		stackLevel := i + 1
		d.ClearLine(y)
		d.screen.PutStr(0, y, fmt.Sprintf("%d:", stackLevel))
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

func (d *Display) TypeCharacterOnScreen(character string) {
	d.screen.PutStr(d.CursorXCoordinate, 11, character)
	d.CursorXCoordinate++
	d.screen.Show()
}

func (d *Display) Backspace() {
	if d.CursorXCoordinate < 1 {
		d.screen.Show()
		return
	}
	d.screen.PutStr(d.CursorXCoordinate-1, 11, " ")
	d.CursorXCoordinate--
	d.screen.Show()
}
