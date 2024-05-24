package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var number = [8]int{1, 2, 3, 4, 4, 2, 5, 6}
	fmt.Println(number)

	var names = [5]string{"bhavil", "jain"}
	fmt.Println(names)
	fmt.Println(len(names)) //length
	fmt.Println(names[0])   // parrticular index

	var hello = [4]string{} //default val for string: empty, array :0 , bool : false
	fmt.Println(hello)

	var n = [4]string{"bhavil", "jain"}
	fmt.Println(n)
	fmt.Printf("%q\n", n) //quoted ways mein ayega

}
