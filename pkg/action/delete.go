package action

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Delete(app *tview.Application, lastKeyPtr *rune){
  lastKey := *lastKeyPtr
  if lastKey == 'd' {
    DeleteFile(Dir, GetSelectedFile())
    UpdateAndDisplayTable()
    *lastKeyPtr = 0
  } else {
    *lastKeyPtr = 'd'
  }
}

func Create(app *tview.Application){
  input := tview.NewInputField().
    SetLabel("File Name: ").
    SetFieldWidth(30)

  input.SetDoneFunc(func(key tcell.Key){
    if key == tcell.KeyEnter {
      CreateFile(Dir, input.GetText())
      UpdateAndDisplayTable()
    }
  })

  app.SetRoot(input, false)
}
