package main

import "fmt"

func convert(from string, to string, value float32) float32 {
	return 2.5
}

func main() {
	for {
		var rupee float32
		var dollar float32
		var euro float32
		var japaneseyen float32
		var britishpound float32
		var choice int

		fmt.Println("to convert rupee in doller press = 1")
		fmt.Println("to convert dollar in rupee press = 2")
		fmt.Println("to contert rupee in euro press = 3")
		fmt.Println("to convert euro in rupee press = 4")
		fmt.Println("to convert rupee in japaneseyen press = 5")
		fmt.Println("to convert japaneseyen in rupee press = 6")
		fmt.Println("to convert rupee in britishpound press = 5")
		fmt.Println("to convert britishpound in rupee press = 6")
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		}

		val := convert("USD", "INR", 10)
		fmt.Println(val)

		switch choice {
		case 1:
			fmt.Printf("enter rupee = ")
			fmt.Scanln(&rupee)
			dollar = rupee / 83.13
			fmt.Println("total dollar value = ", dollar, "$")
		case 2:
			fmt.Printf("enter dollar = ")
			fmt.Scanln(&dollar)
			rupee = dollar * 83.13
			fmt.Println("total rupee value = ", rupee)
		case 3:
			fmt.Printf("enter rupee = ")
			fmt.Scanln(&rupee)
			euro = rupee / 90.68
			fmt.Println("total euro value = ", euro)
		case 4:
			fmt.Printf("enter euro = ")
			fmt.Scanln(&euro)
			rupee = euro * 90.68
			fmt.Println("total euro value = ", rupee)
		case 5:
			fmt.Printf("enter japaneseyen = ")
			fmt.Scanln(&japaneseyen)
			rupee = japaneseyen * 1.82
			fmt.Println("total rupee value = ", rupee)
		case 6:
			fmt.Printf("enter japaneseyen = ")
			fmt.Scanln(&japaneseyen)
			rupee = japaneseyen * 1.82
			fmt.Println("total rupee value = ", rupee)
		case 7:
			fmt.Printf("enter rupee = ")
			fmt.Scanln(&rupee)
			britishpound = rupee / 99.87
			fmt.Println("total rupee value = ", britishpound)
		case 8:
			fmt.Printf("enter britishpound = ")
			fmt.Scanln(&britishpound)
			rupee = britishpound * 99.87
			fmt.Println("total rupee value = ", rupee)
		}
		fmt.Println()
	}
}
