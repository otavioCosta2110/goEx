package actions

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var lastKey rune
func CaptureKeys(app *tview.Application, dirPtr *string) {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				Stop(app)
			}
			if event.Rune() == 'd' {
        Delete(app, dirPtr, &lastKey)

			} else {
				lastKey = 0
			}
		}
		return event
	})
}

func Stop(app *tview.Application) {
	app.Stop()
}
