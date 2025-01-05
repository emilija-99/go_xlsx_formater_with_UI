package main

import (
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func file() {
	fmt.Println("Init: import_export_file")

	// file_path := ""

	f, err := excelize.OpenFile("")
	if err != nil {
		fmt.Println(err)
		return
	}

	getUploadedFile(f)

}

func getUploadedFile(f *excelize.File) {

}

func findFile(fileName string, f *excelize.File) error {
	if fileName == "" {
		return errors.New("File name is empty")
	}

	return nil
}

func saveFile(fileName string, f *excelize.File) {

}

func openFile(fileName string, f *excelize.File) {
	if fileName == "" {
		return
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
