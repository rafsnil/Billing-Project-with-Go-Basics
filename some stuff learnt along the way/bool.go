// package main

// import "fmt"

// func main() {

// 	age := 25

// 	fmt.Println(age <= 50)
// 	fmt.Println(age >= 50)
// 	fmt.Println(age == 45)
// 	fmt.Println(age != 50)

// 	if age < 30 {
// 		fmt.Println("Age is less than 30")
// 	} else if age < 40 {
// 		fmt.Println("Age is less than 40")
// 	} else {
// 		fmt.Println("Age is not less than 45")
// 	}

// 	names := []string{"yoshi", "mario", "bowser", "peaches", "luigi"}

// 	// What does continue do in FOR LOOP in the code below?
// 	// It tells to break out of the current iteration but continue with the loop
// 	// Break totally gets you out of the loop
// 	for index, value := range names {
// 		if index == 1 {
// 			fmt.Println("Continuing at position", index)
// 			continue
// 		}

// 		if index > 2 {
// 			fmt.Println("Breaking at position", index)
// 			break
// 		}

// 		fmt.Printf("The value at position %d is %s\n", index, value)
// 	}

// }
