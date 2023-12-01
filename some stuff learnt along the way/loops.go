// package main

// import "fmt"

// func main() {

// 	// // WHILE LOOP
// 	// x := 0

// 	// //The code below is a while loop equivalent in Go
// 	// //It reads as: while x<5 run the shit inside the curly braces
// 	// for x < 5 {
// 	// 	fmt.Println("Value of x is:", x)
// 	// 	x++
// 	// }

// 	// // TRADITIONAL FOR LOOP
// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println("Value of i is:", i)
// 	// }

// 	// //Traversing through an array using loop
// 	// names := [4]string{"yoshi", "mario", "bowser", "peaches"}
// 	// for i := 0; i < len(names); i++ {
// 	// 	fmt.Printf("Name of Person %d is %s\n", (i + 1), names[i])
// 	// }

// 	// // Traversing through an array using "range"
// 	// names := [4]string{"yoshi", "mario", "bowser", "peaches"}
// 	// for index, value := range names {
// 	// 	fmt.Printf("The value at position %d is %s", index, value)
// 	// }

// 	// If you want to use only value use "_", same for name as well
// 	names := [4]string{"yoshi", "mario", "bowser", "peaches"}
// 	for _, value := range names {
// 		fmt.Printf("The value is %s\n", value)
// 		// If you change "value" here, it will not alter the original array
// 		// value = "new string" doesn't change any shit
// 	}

// }
