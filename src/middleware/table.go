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

func UpdateTable(dirPtr *string) *tview.Table {
  dir := *dirPtr
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

func UpdateAndDisplayTable(dirPtr *string, app *tview.Application) string {
  dir := *dirPtr
	table = UpdateTable(dirPtr)
	if table == nil {
		return ""
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
        dir, err := filepath.Abs(dir)
        if err != nil {
          panic(err)
        }
        *dirPtr = dir
				UpdateAndDisplayTable(dirPtr, app)
			}
		}
	})

	app.SetRoot(table, true).SetFocus(table)
  return dir
}

func GetSelectedFile() string {
  row, _ := table.GetSelection()
  return table.GetCell(row, 0).Text
}
