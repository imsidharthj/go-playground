package main

import "fmt"

func main() {
	for {
		var celcius float32
		var feranite float32
		var choice int
		fmt.Println("from celcius to feranite press 1")
		fmt.Println("from feranite to celcius press 2")
		fmt.Printf("enter value = ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Printf("enter the celcius = ")
			fmt.Scanln(&celcius)
			feranite = (celcius * 9.0 / 5.0) + 32
			fmt.Println("total feranite value = ", feranite)
		} else if choice == 2 {
			fmt.Printf("enter the feranite = ")
			fmt.Scanln(&feranite)
			celcius = (feranite - 32) * 5.0 / 9.0
			fmt.Println(" total celcius value = ", celcius)
		}
		if celcius > 38 || feranite > 100 {
			fmt.Println()
			fmt.Println("you have fever . get docter")
			fmt.Println()
		}
	}

}
