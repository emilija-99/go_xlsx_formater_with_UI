package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	_btnUpload = widget.NewButton("Upload File", func() {
		fmt.Println("Upload File!")
		showFilePicker(_window)

	})

	_lblOr := widget.NewLabelWithStyle("OR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	_inputFileName := widget.NewEntry()
	_inputFileName.SetPlaceHolder("Enter File Name")
	_btnSubmitFile := widget.NewButton("Submit File", func() {
		fmt.Println("Submit File")
		// importExportFile.findFile(_inputFileName.Text)
	})

	_container := container.NewVBox(_btnUpload, _selectedFile, _lblOr, _inputFileName, _btnSubmitFile)
	_window.SetContent(_container)
	_window.Resize(fyne.NewSize(500, 300))
	_window.ShowAndRun()
}
func showFilePicker(_window fyne.Window) {
	dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if f == nil {
			return
		}

		saveFile := f.URI().Path()
		fmt.Println("File URI: ", saveFile)

		_fileURI = f.URI()

		fmt.Println("File URI: ", _fileURI)

		_selectedFile.SetText(saveFile)
		checkFile()

	}, _window)
}
func checkFile() {
	fmt.Println("Init: import_export_file")

	if file != nil {
		return
	}

	fmt.Println("File: ", _fileURI)
	file, err := excelize.OpenFile(_fileURI.Path())

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Rows: ", file)

	sheets := file.GetSheetList()

	rows, err := file.GetRows(sheets[0])

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// file_path := ""
	fmt.Println("Sheet Length: ", sheets[0], rows[0], len(rows[0]))

	fmt.Println("Rows LEN: ", len(rows))
	if len(rows[0]) == 1 {
		fmt.Println("Format this file.")
		formatFile(file, rows)
	} else {
		fmt.Println("File is already formatted.")
	}
	// f, err := excelize.OpenFile("")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// getUploadedFile(f)

}

func formatFile(f *excelize.File, rows [][]string) {
	f.NewSheet("Formatted")
	// rows_len := len(rows)
	for rowIndex, row := range rows {
		rowData := strings.Split(row[0], ",")
		for colIndex, cellValue := range rowData {
			cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			f.SetCellValue("Formatted", cellName, cellValue)
		}
	}

	f.Save()
}

func saveFile(fileName string, f *excelize.File) {

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
