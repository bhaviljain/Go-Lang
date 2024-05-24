package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("value cant be 0")
	}
	return a / b, nil
}

func main() {
	// fmt.Println("error handling")
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
