// package main

// import "fmt"

// /*Here in this function we are accepting A POINTER TO A STRING as a parameter
// Whenever you see a * infront of a type, it means you are accepting a
// pointer to whatever value is stored at that location.*/
// func updateName(p *string) {
// 	// Here we are dereferencing the pointer with an asterisk
// 	//Basically we are getting the value of whatever the pointer is pointing at.
// 	*p = "Lele"
// }

// func main() {

// 	name := "Niloy"

// 	m_pointer := &name
// 	// fmt.Println("Memory Address is ", m_pointer)

// 	// fmt.Println("Value is ", *m_pointer)
// 	fmt.Println(name)

// 	updateName(m_pointer)

// 	fmt.Println(name)

// }

// // REVIEW IF STILL CONFUSED WITH POINTERS
// // https://www.youtube.com/watch?v=4B2rwYvuiBo&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=14
