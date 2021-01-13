package main

import (
	"io/ioutil"
	"strings"
)

const toFile = "result.csv"
const inputFolder = "./input/"

var result = []string{}

func main() {
	for _, fileInFolder := range readFiles(inputFolder) {
		fileName := inputFolder + fileInFolder.Name()
		file, err := ioutil.ReadFile(fileName)
		checkError(err)

		emails := extractEmail(file)
		unduplicated := removeDuplicatesUnordered(emails)

		result = append(result, unduplicated...)
	}

	csv := []byte(strings.Join(result, ",\n"))
	ioutil.WriteFile(toFile, csv, 0644)
}
