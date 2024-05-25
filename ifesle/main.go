package main

import "fmt"

func main() {
	x := 10
	if x > 10 {
		fmt.Println("X is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5 but smaller than 10")
	} else {
		fmt.Println("x is smaller")
	}

	y := 6
	if y > 10 && y < 20 {
		fmt.Println("how are you")
	} else {
		fmt.Println("go away")
	}
}
