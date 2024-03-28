package main

import "fmt"

func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func main() {
	a := fact(7)
	fmt.Println(a)
}
