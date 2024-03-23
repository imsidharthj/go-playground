package main

import "fmt"

func main() {
	for {
		var addition float32
		var subtraction float32
		var division float32
		var multiplication float32
		var choice int
		var a float32
		var b float32
		fmt.Println("for addition press 1")
		fmt.Println("for subtraction press 2")
		fmt.Println("for division press 3")
		fmt.Println("for multiplication press 4")
		fmt.Println("press 0 to cancel")
		fmt.Printf("select choice: ")
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		} else if choice < 0 || choice > 4 {
			fmt.Println("invalid input")
			fmt.Println()
			continue
		}

		fmt.Printf("enter a = ")
		fmt.Scanln(&a)
		fmt.Printf("enter b = ")
		fmt.Scanln(&b)

		if choice == 1 {
			addition = a + b
			fmt.Println("total addition value = ", addition)
		} else if choice == 2 {
			subtraction = a - b
			fmt.Println("total subtraction value = ", subtraction)
		} else if choice == 3 {
			division = a / b
			fmt.Println("total division value = ", division)
		} else if choice == 4 {
			multiplication = a * b
			fmt.Println("total multiplication value = ", multiplication)
		}
		fmt.Println()
	}
}
