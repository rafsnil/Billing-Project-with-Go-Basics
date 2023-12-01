// package main

// import "fmt"

// /* Here ven though we are passing "name" into the function, as GO is a
// PASS BY VALUE language, it actually creates a copy of the "name"
// variable, and it is this copied variable that gets updated in
// the fucntion below. */

// // This happens for all data types except maps slices
// // e.g. int, float, bool, string, arrays, structs
// // LETS CALL THEM NON_POINTER VALUES

// // func updateName(x string) {
// // 	x = "ChodnaMagi"
// // }

// func updateName(x string) string {
// 	x = "ChodnaMagi"
// 	//To actually change the value of name, we would return the desired value
// 	return "ChodnaMagi"
// }

// func updateMenu(menu map[string]float64) {
// 	menu["coffee"] = 5.99
// }

// func main() {

// 	// name := "Mahadi"

// 	// name = updateName(name)

// 	// fmt.Println(name)

// 	// For types like slices, maps, functions.
// 	// LETS CALL THEM POINTER WRAPPER VALUES
// 	menu := map[string]float64{
// 		"ice-cream": 4.99,
// 		"wedges":    3.55,
// 	}

// 	updateMenu(menu)

// 	fmt.Println(menu)

// }
