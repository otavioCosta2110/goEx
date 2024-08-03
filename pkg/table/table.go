package table

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/otavioCosta2110/goex/pkg/color"
	filepkg "github.com/otavioCosta2110/goex/pkg/files"
	"github.com/otavioCosta2110/goex/pkg/global"
	"github.com/rivo/tview"
)

var table *tview.Table

func GetFilesTable(files []os.FileInfo) *tview.Table {
	table := tview.NewTable()
	table.SetSelectable(true, false)

	fileInfos := filepkg.GetFilesStruct(files)

	for i, file := range fileInfos {
		table.SetCell(i, 0,
			tview.NewTableCell(file.Name).
				SetTextColor(color.GetColor(file)).
				SetAlign(tview.AlignLeft).
				SetSelectable(true))
	}

	return table
}

func UpdateTable() *tview.Table {
	d, err := os.Open(global.Dir)
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
		d, err := os.Open(global.Dir)
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

		fileInfos := filepkg.GetFilesStruct(files)

		if row >= 0 && row < len(fileInfos) {
			selectedFile := fileInfos[row]
			if selectedFile.IsDir {
				newDir := filepath.Join(global.Dir, selectedFile.Name)
				newDir, err := filepath.Abs(newDir)
				if err != nil {
					fmt.Println("Error getting absolute path:", err)
					return
				}
				global.Dir = newDir

				global.Flex.Clear()
				UpdateAndDisplayTable()
			}
		}
	})

  global.Dir, _ = filepath.Abs(global.Dir)

  global.Flex.Clear()
  global.TextV.SetTitle(global.Dir)

  global.Flex.AddItem(table, 0, 20, true)
  global.Flex.AddItem(global.TextV, 0, 1, true)

	global.App.SetRoot(global.Flex, true).SetFocus(table)
}

func GetSelectedFile() string {
  row, _ := table.GetSelection()
  return table.GetCell(row, 0).Text
}
