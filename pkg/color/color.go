package color

import (
	"github.com/gdamore/tcell/v2"
	"github.com/otavioCosta2110/goex/pkg/files"
)

func GetColor(file files.FileInfo) tcell.Color {
	if file.IsDir {
		return tcell.ColorBlue
	}
	return tcell.ColorGreen
}
