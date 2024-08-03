package main

import (
	capturekeys "github.com/otavioCosta2110/goex/pkg/captureKeys"
	"github.com/otavioCosta2110/goex/pkg/global"
	"github.com/otavioCosta2110/goex/pkg/table"

	"os"

	"github.com/rivo/tview"
)

func main() {
  Init()
}

func Init() {
	global.Dir = os.Getenv("HOME")

	if len(os.Args) > 1 {
		global.Dir = os.Args[1]
	}

	global.App = tview.NewApplication()

  global.Flex = tview.NewFlex().SetDirection(tview.FlexRow)
  global.TextV = tview.NewTextView().SetBorder(true).SetTitle(global.Dir)

  capturekeys.CaptureKeys()
  table.UpdateAndDisplayTable()

	global.App.SetRoot(global.Flex, true)
	if err := global.App.Run(); err != nil {
		panic(err)
	}
}
