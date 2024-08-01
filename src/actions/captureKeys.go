package actions

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var lastKey rune

func CaptureKeys(app *tview.Application, dirPtr *string) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		focused := app.GetFocus()

		if _, ok := focused.(*tview.InputField); ok {
			return event
		}

		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				Stop(app)
			case 'd':
				Delete(app, dirPtr, &lastKey)
			case 'a':
				Create(app, dirPtr)
			default:
				lastKey = 0
			}
		}
		return event
	})
}

func Stop(app *tview.Application) {
	app.Stop()
}
