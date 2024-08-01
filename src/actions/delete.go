package actions

import (
	"otaviocosta2110/goEx/src/middleware"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Delete(app *tview.Application, dirPtr *string, lastKeyPtr *rune){
  if lastKey == 'd' {
    dir := *dirPtr
    middleware.DeleteFile(dir, middleware.GetSelectedFile())
    middleware.UpdateAndDisplayTable(dirPtr, app)
    *lastKeyPtr = 0
  } else {
    *lastKeyPtr = 'd'
  }
}

func Create(app *tview.Application, dirPtr *string){
  dir := *dirPtr

  bunda := tview.NewInputField().
    SetLabel("File Name: ").
    SetFieldWidth(30)

  bunda.SetDoneFunc(func(key tcell.Key){
    if key == tcell.KeyEnter {
      middleware.CreateFile(dir, bunda.GetText())
      middleware.UpdateAndDisplayTable(dirPtr, app)
    }
  })

  app.SetRoot(bunda, true)
}
