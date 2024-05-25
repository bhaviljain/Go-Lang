package main

import "fmt"

func main() {
	numbers := []int{12, 3, 4}
	numbers = append(numbers, 1, 2, 3, 5, 6, 6)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Printf("%T\n", numbers)

	num := make([]int, 3, 5)
	num = append(num, 1)
	num = append(num, 3)
	num = append(num, 5)
	num = append(num, 1)
	num = append(num, 9)
	num = append(num, 13)
	num = append(num, 88)
	num = append(num, 88)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	num1 := make([]int, 0)
	fmt.Println(num1)

}

//slice  to init by make or like this
// number:=[]int{}
//you cant init like array
