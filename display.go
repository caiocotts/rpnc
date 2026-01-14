package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

var defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

func CloseDisplay(screen tcell.Screen) {
	screen.Fini()
}

func InitDisplay() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Init(); err != nil {
		log.Fatal(err)
	}
	s.SetStyle(defStyle)
	return s
}

func PrintStack(calculator Calculator, screen tcell.Screen, numberOfLevels int) {
	for i := 0; i < numberOfLevels; i++ {
		y := numberOfLevels - i
		stackLevel := i + 1
		ClearLine(screen, y)
		screen.PutStr(0, y, fmt.Sprintf("%d:", stackLevel))
		val, err := calculator.Stack.Pop()
		if err == nil {
			screen.PutStr(5, y, val)
		}
	}
	screen.Show()
}

func PrintMessage(screen tcell.Screen, message string) {
	screen.PutStr(0, 0, message)
	screen.Show()
}

func ClearLine(screen tcell.Screen, y int) {
	width, _ := screen.Size()
	for i := 0; i < width; i++ {
		screen.PutStr(i, y, " ")
	}
	screen.Show()
}
