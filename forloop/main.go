package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	//infinite loop
	count := 0
	for {
		fmt.Println(count)
		count++
		if count == 1 {
			break
		}
	}
	number := []int{11, 23, 13, 94}
	for ok, oks := range number {
		fmt.Println(ok, oks) //range gives you index and value
	}
	data := "hello, bhai"

	for index, char := range data {
		fmt.Printf("index is %d, and value is %c\n", index, char) //range gives you index and value
	}

}

// go mein sirf ek he loop hain, for loop, (while aur do while loop nahi hota hai)
