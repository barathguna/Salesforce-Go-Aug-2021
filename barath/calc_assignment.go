package main

import "fmt"

func main() {
	var mode int
	printMenu()
	fmt.Scanf("%d", &mode)
	for mode != 5 {
		var a, b int
		if mode >= 1 && mode <= 5 {
			fmt.Println("Enter Value 1:")
			fmt.Scanf("%d", &a)
			fmt.Println("Enter Value 2:")
			fmt.Scanf("%d", &b)
		}

		switch mode {
		case 1:

			fmt.Println("Add value for ", a, "and", b, "is", a+b)
		case 2:
			fmt.Println("Substract value for ", a, "and", b, "is", a-b)
		case 3:
			fmt.Println("Multiply value for ", a, "and", b, "is", a*b)
		case 4:
			if b != 0 {
				fmt.Println("Divide value for ", a, "and", b, "is", a/b)
			} else {
				fmt.Println("Divide value for ", a, "and", b, "is not possible")
			}
		default:
			fmt.Println("Invalid mode")
		}
		// repeat
		printMenu()
		fmt.Scanf("%d", &mode)
	}
	fmt.Println("Thank you !!!")

}

func printMenu() {
	fmt.Println("Calculator")
	fmt.Println("\t 1. Add")
	fmt.Println("\t 2. Substract")
	fmt.Println("\t 3. Multiply")
	fmt.Println("\t 4. Divide")
	fmt.Println("\t 5. Exit")
	fmt.Println("Select mode (1/2/3/4/5):")
}
