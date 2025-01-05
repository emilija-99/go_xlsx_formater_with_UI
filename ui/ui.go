package main

import (
	"fmt"

	"go_xlsx_formater.com/importExportFile"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func runUI() {

	app := app.New()
	_window := app.NewWindow("Format Excel File")

	_btnUpload := widget.NewButton("Upload File", func() {
		fmt.Println("Upload File!")
		importExportFile.GetUploadedFile()
	})

	_lblOr := widget.NewLabelWithStyle("OR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	_inputFileName := widget.NewEntry()
	_inputFileName.SetPlaceHolder("Enter File Name")
	_btnSubmitFile := widget.NewButton("Submit File", func() {
		fmt.Println("Submit File")
		importExportFile.findFile(_inputFileName.Text)
	})

	_container := container.NewVBox(_btnUpload, _lblOr, _inputFileName, _btnSubmitFile)
	_window.SetContent(_container)
	_window.Resize(fyne.NewSize(500, 300))
	_window.ShowAndRun()

}
