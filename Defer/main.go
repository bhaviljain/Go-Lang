package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("Hello2")
	res := add(2, 3)
	fmt.Println(res)

}

//allows a function to postpone the execution of a statement until the surrounding function has completed
//do defer hoga toh LIFO mein jayega
