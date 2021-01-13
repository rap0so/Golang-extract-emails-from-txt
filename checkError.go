package main

func checkError(eventError error) {
	if eventError != nil {
		panic(eventError)
	}
}
