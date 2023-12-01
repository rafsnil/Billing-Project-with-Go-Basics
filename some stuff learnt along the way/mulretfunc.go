// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // paramter      return type(multiple in this case)
// func getInitails(name string) (string, string) {
// 	s := strings.ToUpper(name)
// 	words := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range words {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

// func main() {
// 	fn1, sn1 := getInitails("Mahadi Chodna")
// 	fmt.Println(fn1, sn1)

// 	fn2, sn2 := getInitails("Farhan Magi")
// 	fmt.Println(fn2, sn2)

// 	fn3, sn3 := getInitails("Niloy")
// 	fmt.Println(fn3, sn3)

// }
