package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple. bananan. pineabpple"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	str := "one two three four one"
	count := strings.Count(str, "one")
	fmt.Println(count)

	str = "    xjnsj whbdb wjnwj                   " //start aur end mein se he space niklega
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "bhavil"
	str2 := "jain"
	res := strings.Join([]string{str1, "Mukesh", str2}, ",")
	fmt.Println(res)
}
