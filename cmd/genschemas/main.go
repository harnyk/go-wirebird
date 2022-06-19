package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"io/ioutil"
	"reflect"

	"github.com/danielgtaylor/huma/schema"
	"github.com/harnyk/go-wirebird/internal/models"
)

func main() {
	fileNamePtr := flag.String("o", "", "File name to write to")
	flag.Parse()
	if fileNamePtr == nil || *fileNamePtr == "" {
		panic("No file name specified")
	}
	fileName := *fileNamePtr

	t := reflect.TypeOf(models.LoggerEvent{})

	sch, err := schema.Generate(t)
	if err != nil {
		panic(err)
	}

	schJson, err := json.MarshalIndent(sch, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the schema to the file
	err = ioutil.WriteFile(fileName, schJson, fs.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
