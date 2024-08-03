package table

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

func UpdateTable() *tview.Table {
	d, err := os.Open(Dir)
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

func UpdateAndDisplayTable() {
	table = UpdateTable()
	if table == nil {
		return
	}

	table.SetSelectedFunc(func(row, column int) {
		d, err := os.Open(Dir)
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
				newDir := filepath.Join(Dir, selectedFile.Name)
				newDir, err := filepath.Abs(newDir)
				if err != nil {
					fmt.Println("Error getting absolute path:", err)
					return
				}
				Dir = newDir

				Flex.Clear()
				UpdateAndDisplayTable()
			}
		}
	})

  Dir, err := filepath.Abs(Dir)
  if err != nil {
    panic(err)
  }

  Flex.Clear()
  TextV.SetTitle(Dir)

  Flex.AddItem(table, 0, 20, true)
  Flex.AddItem(TextV, 0, 1, true)

	App.SetRoot(Flex, true).SetFocus(table)
}

func GetSelectedFile() string {
  row, _ := table.GetSelection()
  return table.GetCell(row, 0).Text
}
