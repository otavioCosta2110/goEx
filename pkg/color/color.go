package color

import "github.com/gdamore/tcell/v2"

func GetColor(file FileInfo) tcell.Color {
	if file.IsDir {
		return tcell.ColorBlue
	}
	return tcell.ColorGreen
}
