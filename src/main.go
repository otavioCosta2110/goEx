package main

import (
	"os"
	"path/filepath"

	"otaviocosta2110/goEx/src/actions"
	table "otaviocosta2110/goEx/src/middleware"

	"github.com/rivo/tview"
)

var app *tview.Application
var dir string

func main() {
	dir = "."

	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
  dir, err := filepath.Abs(dir)
  if err != nil {
    panic(err)
  }

	app = tview.NewApplication()

  dirPtr := &dir

  actions.CaptureKeys(app, dirPtr)

	table.UpdateAndDisplayTable(dirPtr, app)


	if err := app.Run(); err != nil {
		panic(err)
	}
}

