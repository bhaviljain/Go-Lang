package main

import (
	"fmt"
	"strconv"
)

func main() {
	// num := 42
	// // num = num + 10 //can be added
	// fmt.Println(num)
	// fmt.Printf("Type %T", num)

	// num = 123
	// str := strconv.Itoa(num)
	// fmt.Println("str is", str)
	// fmt.Printf("type of %T\n", str)

	// number := "1234"
	// nums, _ := strconv.Atoi(number)
	// nums = nums + 100
	// fmt.Println("str is", nums)
	// fmt.Printf("type of %T\n", nums)

	number1 := "-123"
	nums1, _ := strconv.ParseInt(number1, 10, 64)
	fmt.Println("str is", nums1)
	fmt.Printf("type of %T\n", nums1)
}
