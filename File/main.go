package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("done")
	// }
	// defer file.Close()
	// fmt.Println("file done")
	// con := "hello bhavil"
	// _, errs := io.WriteString(file, con+"\n")
	// if errs != nil {
	// 	fmt.Println("errrs")
	// 	return
	// }
	// fmt.Println("file done")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("erros")
	// 	return
	// }
	// defer file.Close()

	// buffer := make([]byte, 1024)
	// for {
	// 	n, errr := file.Read(buffer)
	// 	if errr == io.EOF {
	// 		break
	// 	}

	// 	if errr != nil {
	// 		fmt.Println("err")
	// 		return
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }
	con, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("erros")
		return
	}
	fmt.Println(string(con))
}
