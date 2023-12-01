package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

// score is accessed from main.go
func showScore() {
	fmt.Printf("You scored %0.2f points", score)
}
