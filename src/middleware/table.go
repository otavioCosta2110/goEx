package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

var table *tview.Table

func GetFilesTable(files []os.FileInfo) *tview.Table {
	table := tview.NewTable()
	table.SetSelectable(true, false)

	fileInfos := GetFilesStruct(files)

	for i, file := range fileInfos {
		table.SetCell(i, 0,
			tview.NewTableCell(file.Name).
				SetTextColor(GetColor(file)).
				SetAlign(tview.AlignLeft).
				SetSelectable(true))
	}

	return table
}

func UpdateTable(dir string) *tview.Table {
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return nil
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	return GetFilesTable(files)
}

func UpdateAndDisplayTable(dir string, app *tview.Application) {
	table = UpdateTable(dir)
	if table == nil {
		return
	}

	table.SetSelectedFunc(func(row, column int) {
		d, err := os.Open(dir)

		if err != nil {
			fmt.Println("Error opening directory:", err)
			return
		}

		defer d.Close()

		files, err := d.Readdir(-1)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		fileInfos := GetFilesStruct(files)

		if row >= 0 && row < len(fileInfos) {
			selectedFile := fileInfos[row]
			if selectedFile.IsDir {
				dir = filepath.Join(dir, selectedFile.Name)
				UpdateAndDisplayTable(dir, app)
			}
		}
	})

	app.SetRoot(table, true).SetFocus(table)
}
