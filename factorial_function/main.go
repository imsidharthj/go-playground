package main

import "fmt"

func main() {
	// reverse a name
	// 1. input name
	var name string
	fmt.Scanln(&name)
	// 2. reverse given input name
	reversedName := reverseName(name)
	// 3. print reversed name
	fmt.Println(reversedName)
}

func reverseName(name string) string {
	return name
}
