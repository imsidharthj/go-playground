package main

import "fmt"

func saygreeting(n string) {
	fmt.Printf("hii %v\n", n)
}

func saygreeting1(n string) {
	fmt.Printf("by %v\n", n)
}

func cyclenames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	cyclenames([]string{"a, b, c, d"}, saygreeting)
	cyclenames([]string{"a, b, c, d"}, saygreeting1)
}
