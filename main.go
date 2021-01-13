package main

import (
	"io/ioutil"
	"strings"
	"time"
)

const fromFile = "./txt/a.txt"

// const toFile = "result.csv"

func main() {
	file, err := ioutil.ReadFile(fromFile)
	checkError(err)

	emails := extractEmail(file)
	unduplicated := removeDuplicatesUnordered(emails)

	csv := []byte(strings.Join(unduplicated, ",\n"))

	a := time.Now().Format("20060102150405")
	// toFile := strconv.Itoa(hours) + strconv.Itoa(minutes) + strconv.Itoa(sec) + ".csv"
	// fmt.Print(hours)
	ioutil.WriteFile(a, csv, 0644)
}
