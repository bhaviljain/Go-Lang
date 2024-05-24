package main

import "fmt"

func simple() {
	fmt.Println("simple")
}

func main() {
	fmt.Println("Function")
	simple()
	result := add(2, 3)
	fmt.Println(result)
	res := mul(2.3, 3.2)
	fmt.Printf("%.1f\n", res)

}

//data type of input    //data type of output
func add(a, b int) (result int) { // curly braces ko same line mein karo , next line mein error dega
	result = a + b
	return
	// return  a + b (aisa bhi likh sakte hain)

}

func mul(c, d float64) (res float64) {
	res = c * d
	return
}
