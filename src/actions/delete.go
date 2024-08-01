package actions

import (
	"otaviocosta2110/goEx/src/middleware"

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
