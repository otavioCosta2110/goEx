package main

import (
	"fmt"
	"os"

	table "otaviocosta2110/goEx/src/middleware"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var dir string
var lastKey rune

func main() {
	dir = "."

	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	app = tview.NewApplication()

	table.UpdateAndDisplayTable(dir, app)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				Stop()
			}
			if event.Rune() == 'd' {
				if lastKey == 'd' {
					fmt.Println("Delete")
					lastKey = 0
				} else {
					lastKey = 'd'
				}
			} else {
				lastKey = 0
			}
		}
		return event
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func Stop() {
	app.Stop()
}

