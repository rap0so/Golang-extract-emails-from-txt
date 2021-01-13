package main

import (
	"io/ioutil"
	"os"
)

func readFiles(folder string) []os.FileInfo {
	inputFiles, err := ioutil.ReadDir(folder)
	checkError(err)

	return inputFiles
}
