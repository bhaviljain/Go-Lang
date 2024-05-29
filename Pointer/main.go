package main

import "fmt"

func modifyvalu(num *int) {
	*num = *num + 20
}

func main() {
	// var num int
	// num = 2
	// var ptr *int
	// ptr = &num

	//short mein likha
	num := 2
	ptr := &num

	fmt.Println(num)
	fmt.Println(*ptr)

	value := 10
	modifyvalu(&value)
	fmt.Println((value))
	//modifies the value thru refrenece
}

//defualt value of prt is nil
