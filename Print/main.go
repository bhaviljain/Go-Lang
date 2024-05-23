package main

import "fmt"

func main() {
	age := 22
	name := "Bhavil"
	height := 5.9
	fmt.Println(age, name, height)
	fmt.Printf("age is %d", age)
	fmt.Printf("height is %.1f\n", height)
	fmt.Printf("type of age is %T\n", age)
	fmt.Printf("type of name is %T\n", name)

	fmt.Printf("name %s, age %.1d\n", name, age) //print all together.
	//printf is format specifier , no space, no break line
}
