package main

import "fmt"

func checkError(msg string, eventError error) {
	if eventError != nil {
		fmt.Print(msg + "\n")
		panic(eventError)
	}
}
