package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Println(myurl)
	fmt.Printf("type of %T\n", myurl)

	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(parsedURL)
	fmt.Printf("type of %T\n", parsedURL)
	fmt.Println("scheme", parsedURL.Scheme)
	fmt.Println("host", parsedURL.Host)
	fmt.Println("path", parsedURL.Path)
	fmt.Println("qw", parsedURL.RawQuery)

	parsedURL.Path = "/gandu"
	parsedURL.Host = "bhavil.com"
	parsedURL.Scheme = "xnxx"
	newUrl := parsedURL.String()
	fmt.Println("newUrl:", newUrl)
}
