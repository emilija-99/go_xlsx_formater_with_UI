package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

var _btnUpload *widget.Button

var _selectedFile *widget.Label
var _fileURI fyne.URI
var file *excelize.File

func main() {
	runUI()
}

func runUI() {

	app := app.NewWithID("com.example.import_export_file")
	_window := app.NewWindow("Format Excel File")

	_selectedFile = widget.NewLabel("No file selected")

	_btnUpload = widget.NewButton("Upload file from you ", func() {
		// fmt.Println("Upload File!")
		showFilePicker(_window)

	})

	_lblOr := widget.NewLabelWithStyle("OR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	_inputFileName := widget.NewEntry()
	_inputFileName.SetPlaceHolder("Enter File Name")
	_btnSubmitFile := widget.NewButton("Submit File", func() {
		// fmt.Println("Submit File")
		findFile(_window, _inputFileName.Text)
	})
	_toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			dialog.ShowInformation("Toolbar Help", "Please make sure you file is in ../Downlaod dir", _window)
			// fmt.Println("Display help")
			// TODO: please make sure you add a help section for toolbar
		}),
	)

	_content := container.NewBorder(_toolbar, nil, nil, nil, nil)
	_container := container.NewVBox(_btnUpload, _selectedFile, _lblOr, _inputFileName, _content, _btnSubmitFile)
	_window.SetContent(_container)
	_window.Resize(fyne.NewSize(500, 300))
	_window.ShowAndRun()
}
func showFilePicker(_window fyne.Window) {
	dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			// fmt.Println("Error: ", err)
			return
		}

		if f == nil {
			return
		}

		saveFile := f.URI().Path()
		// fmt.Println("File URI: ", saveFile)

		_fileURI = f.URI()

		// fmt.Println("File URI: ", _fileURI)

		_selectedFile.SetText(saveFile)
		checkFile(_window, "")

	}, _window)
}
func checkFile(_window fyne.Window, file_path string) {
	fmt.Println("Init: import_export_file", file_path)

	if file != nil {
		return
	}

	if file_path != "" {
		fmt.Println("File Path: ", file_path)
		var err error
		file, err = excelize.OpenFile(file_path)
		if err != nil {
			fmt.Println("Error: ", err)
			// dialog.ShowError(err, _window)
			return
		}
	} else {
		fmt.Println("File: ", _fileURI)
		var err error
		file, err = excelize.OpenFile(_fileURI.Path())
		if err != nil {
			fmt.Println("Error: ", err)
			dialog.ShowError(err, _window)
			return
		}
	}

	fmt.Println("Rows: ", file)

	sheets := file.GetSheetList()

	rows, err := file.GetRows(sheets[0])

	if err != nil {
		fmt.Println("Error: ", err)
		dialog.ShowError(err, _window)
		return
	}
	// file_path := ""
	fmt.Println("Sheet Length: ", sheets[0], rows[0], len(rows[0]))

	fmt.Println("Rows LEN: ", len(rows))
	if len(rows[0]) == 1 {
		fmt.Println("Format this file.")
		formatFile(file, rows, _window)
	} else {
		fmt.Println("File is already formatted.")
	}
}

func formatFile(f *excelize.File, rows [][]string, _window fyne.Window) {
	f.NewSheet("Formatted")
	// rows_len := len(rows)
	for rowIndex, row := range rows {
		rowData := strings.Split(row[0], ",")
		for colIndex, cellValue := range rowData {
			cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			f.SetCellValue("Formatted", cellName, cellValue)
		}
	}

	if err := f.Save(); err != nil {
		fmt.Println("Error: ", err)
		dialog.ShowError(err, _window)
		return
	} else {
		dialog.ShowInformation("Success", "File has been formatted successfully.", _window)
	}
}

func findFile(_window fyne.Window, name string) {
	DownloadDirNames := []string{"Downloads", "download", "Download", "downloads"}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	var _err error
	for _, dirName := range DownloadDirNames {
		filePath := filepath.Join(homeDir, dirName, name)
		fmt.Println("File Path: ", filePath)

		for _, ddn := range DownloadDirNames {
			var dir = filepath.Join(homeDir, ddn)

			if _, err := os.Stat(dir); os.IsNotExist(err) {
				fmt.Println(dir, "does not exist")
				dialog.ShowError(err, _window)
				break

			} else {
				fmt.Println("The provided directory named", dir, "exists")
				// downloadDir = dir
				_fileURI = storage.NewFileURI(dir + "/" + name)

				if _, err := os.Stat(_fileURI.Path()); os.IsNotExist(err) {
					fmt.Println("File does not exist")
					_err = errors.New("File does not exist")
				} else {
					checkFile(_window, _fileURI.Path())
					break
				}

				break
			}

		}

	}

	// fmt.Println("Error: ", _err)
	if _err != nil {
		dialog.ShowError(_err, _window)
	}
}
