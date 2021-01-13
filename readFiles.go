package main

import (
	"io/ioutil"
	"os"
)

func readFiles(folder string) []os.FileInfo {
	inputFiles, err := ioutil.ReadDir(folder)
	checkError("A pasta 'input' não foi criada corretamente :(", err)

	return inputFiles
}
