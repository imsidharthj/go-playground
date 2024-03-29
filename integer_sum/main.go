package main

import "fmt"

func add(a, b float32) (float32, float32) {
	c := a + b
	d := a - b
	return c, d
}

func main() {
	c, d := add(4, 3)
	fmt.Println("addresult:", c)
	fmt.Println("addresult:", d)
}
