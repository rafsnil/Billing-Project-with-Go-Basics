package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input_name, err := r.ReadString('\n')
	return strings.TrimSpace(input_name), err
}

func createBill() bill {

	scanner := bufio.NewReader(os.Stdin)
	// fmt.Print("Create A New Bill Name: ")
	// name, _ := scanner.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create A New Bill Name: ", scanner)

	b := newBill(name)

	fmt.Println("Created The Bill - ", name)

	return b
}

func promtOptions(b bill) {
	scanner := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip): ", scanner)
	// fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", scanner)
		price, _ := getInput("Item Price: ", scanner)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be in number...")
			promtOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Items added: ", name, p)
		promtOptions(b)

	case "s":
		b.saveBill()
	case "t":

		tip, _ := getInput("Tip: ", scanner)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be in number...")
			promtOptions(b)
		}

		b.updateTip(t)

		fmt.Println("Tip added! BDT", t)
		promtOptions(b)

	default:
		fmt.Println("Incorrect Input...")
		promtOptions(b)
	}
}

func main() {
	mybill := createBill()
	promtOptions(mybill)
}
