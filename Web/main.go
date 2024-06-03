package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer res.Body.Close() //ek bar data agaya fir close kardo
	fmt.Println(res)

	//read

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(data))
	//data ko string mein convert karne ke liye
}
