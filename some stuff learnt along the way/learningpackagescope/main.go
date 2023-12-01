package main

import "fmt"

// score is accessed in greetings.go as both this and that file are
// in the same package
var score float64 = 90.55

func main() {

	// points is accessed from greeting.go
	for _, v := range points {
		fmt.Println(v)
	}

	// Both are functions below are accessed from greetings.go
	sayHello("Mahadi Chodna")
	showScore()
}
