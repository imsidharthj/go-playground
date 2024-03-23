package main

import "fmt"

func main() {
	for {
		var choice int
		var km float32
		var meter float32
		fmt.Println("km to meter press 1")
		fmt.Println("mete1r to km press 2")
		fmt.Printf("enter value = ")
		fmt.Scanln(&choice)

		if choice == 1 {

			fmt.Printf("enter km = ")
			fmt.Scanln(&km)
			meter = km * 1000
			fmt.Println("total value=", meter)

		} else if choice == 2 {

			fmt.Printf("enter meter = ")
			fmt.Scanln(&meter)
			km = meter / 1000
			fmt.Println("total value=", km)
		}

	}
}
