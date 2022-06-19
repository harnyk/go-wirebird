package main

import (
	"flag"

	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	fileNamePtr := flag.String("o", "", "File name to write to")
	flag.Parse()
	if fileNamePtr == nil || *fileNamePtr == "" {
		panic("No file name specified")
	}
	fileName := *fileNamePtr

	converter := typescriptify.New().WithInterface(true).WithBackupDir("")

	converter.Add(models.LoggerEvent{})
	err := converter.ConvertToFile(fileName)
	if err != nil {
		panic(err.Error())
	}
}
