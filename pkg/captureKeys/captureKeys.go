package capturekeys

import (
	"github.com/gdamore/tcell/v2"
	"github.com/otavioCosta2110/goex/pkg/action"
	"github.com/otavioCosta2110/goex/pkg/global"
	"github.com/rivo/tview"
)

var lastKey rune

func CaptureKeys() {
	global.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		focused := global.App.GetFocus()

		if _, ok := focused.(*tview.InputField); ok {
			return event
		}

		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				Stop()
			case 'd':
				action.Delete(&lastKey)
			case 'a':
				action.Create()
			default:
				lastKey = 0
			}
		}
		return event
	})
}

func Stop() {
	global.App.Stop()
}
