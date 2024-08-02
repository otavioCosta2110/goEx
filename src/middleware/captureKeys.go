package middleware

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var lastKey rune

func CaptureKeys() {
	App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		focused := App.GetFocus()

		if _, ok := focused.(*tview.InputField); ok {
			return event
		}

		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				Stop(App)
			case 'd':
				Delete(App, &lastKey)
			case 'a':
				Create(App)
			default:
				lastKey = 0
			}
		}
		return event
	})
}

func Stop(App *tview.Application) {
	App.Stop()
}
