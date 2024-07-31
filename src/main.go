package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)


func main() {
  app := tview.NewApplication()

  dir := "."

  d, err := os.Open(dir)

  if err != nil{
    fmt.Println("Error: ", err)
    return
  }

  defer d.Close()

  files, err := d.Readdir(-1)

  table := tview.NewTable()
  table.SetSelectable(true, false)


  for i, file := range files {
    table.SetCell(i, 0,
    tview.NewTableCell(file.Name()).
    SetTextColor(getColor(file)).
    SetAlign(tview.AlignLeft).
    SetSelectable(true))
  }

  if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}


}
func getColor(file os.FileInfo) tcell.Color {
  if file.IsDir() {
    return tcell.ColorBlue
  }
  return tcell.ColorGreen
}
