package global

import (
	"os"

	"github.com/rivo/tview"

  "github.com/otaviocosta2110/goEx/pkg/capturekeys"
)

var App *tview.Application
var Dir string
var Flex *tview.Flex
var TextV *tview.Box

func Init() {
	Dir = os.Getenv("HOME")

	if len(os.Args) > 1 {
		Dir = os.Args[1]
	}

	App = tview.NewApplication()

  CaptureKeys()

  Flex = tview.NewFlex().SetDirection(tview.FlexRow)
  TextV = tview.NewTextView().SetBorder(true).SetTitle(Dir)

  UpdateAndDisplayTable()

	App.SetRoot(Flex, true)
	if err := App.Run(); err != nil {
		panic(err)
	}
}
