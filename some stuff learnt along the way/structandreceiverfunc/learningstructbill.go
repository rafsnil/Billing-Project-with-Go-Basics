// /*
// Sure, here is a summary of the transcript into points:

// - When calling a function with arguments, a copy of the arguments is created.
// - For receiver functions, a copy of the receiver object is also created.
// - When updating a value inside a function, passing a pointer to the value allows the original value to be updated.
// - Passing pointers to functions is efficient as it avoids creating unnecessary copies of large objects.
// - Receiver functions typically use either pointers or no pointers.
// */

// package main

// import "fmt"

// type bill struct {
// 	name  string
// 	items map[string]float64
// 	tip   float64
// }

// // Make new bills
// func newBill(name string) bill {
// 	b := bill{
// 		name:  name,
// 		items: map[string]float64{},
// 		tip:   0,
// 	}
// 	return b
// }

// // Receiver Functions
// // Format the bill
// // This binds the function to bill object, which enables
// // b.format(~)
// func (b *bill) format() string {
// 	fs := "Bill Breakdown: \n"
// 	var total float64 = 0

// 	// List items and their prices
// 	for i, v := range b.items {
// 		fs += fmt.Sprintf("%-20v ...BDT %v \n", i+": ", v)
// 		total += v
// 	}

// 	//tip
// 	fs += fmt.Sprintf("%-20v ...BDT %v\n", "Tip:", b.tip)

// 	//total
// 	fs += fmt.Sprintf("%-20v ...BDT %v", "Total:", total+b.tip)

// 	return fs
// }

// // Update Tip
// func (b *bill) updateTip(tip float64) {
// 	b.tip = tip
// }

// // Add Item to the list
// func (b *bill) addItem(name string, price float64) {
// 	b.items[name] = price
// }
