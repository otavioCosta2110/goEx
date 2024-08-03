package action

import (
	"github.com/gdamore/tcell/v2"
	"github.com/otavioCosta2110/goex/pkg/files"
	"github.com/otavioCosta2110/goex/pkg/global"
	"github.com/otavioCosta2110/goex/pkg/table"
	"github.com/rivo/tview"
)

func Delete(lastKeyPtr *rune){
  lastKey := *lastKeyPtr
  if lastKey == 'd' {
    files.DeleteFile(global.Dir, table.GetSelectedFile())
    table.UpdateAndDisplayTable()
    *lastKeyPtr = 0
  } else {
    *lastKeyPtr = 'd'
  }
}

func Create(){
  input := tview.NewInputField().
    SetLabel("File Name: ").
    SetFieldWidth(30)

  input.SetDoneFunc(func(key tcell.Key){
    if key == tcell.KeyEnter {
      files.CreateFile(global.Dir, input.GetText())
      table.UpdateAndDisplayTable()
    }
  })

  global.App.SetRoot(input, false)
}
