// package main

// import (
// 	"fmt"
// 	"math"
// )

// func sayGreeting(name string) {
// 	fmt.Printf("Hello Chodna %s\n", name)
// }

// func sayBye(name string) {
// 	fmt.Printf("Goodbye Chodna %s\n", name)
// }

// // Here the parameter of the functions are: a slice and a function
// // YES YOU CAN HAVE FUNCCTIONS AS PARAMETERS IN GO
// func cycleNames(names []string, f func(string)) {
// 	for _, value := range names {
// 		f(value)
// 	}
// }

// func circleArea(rad float64) float64 {
// 	return math.Pi * rad * rad
// }

// func main() {
// 	// sayGreeting("Mahadi")
// 	// sayBye("Mahadi")

// 	// names := []string{"yoshi", "mario", "bowser", "peaches"}
// 	// cycleNames(names, sayGreeting)
// 	// cycleNames(names, sayBye)

// 	area1 := circleArea(3.5)
// 	area2 := circleArea(4.5)
// 	fmt.Printf("%0.2f\n", area1)
// 	fmt.Printf("%0.2f", area2)
// }
