package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello, what is your name ?")
	// var name = ""
	// fmt.Scan(&name) //yeh sirf ek he word read karega
	// fmt.Println("hello", name)
	fmt.Println("hello, what is your name ?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("hello", name)
}

// buffer is temporary memory in storage
