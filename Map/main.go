package main

import (
	"fmt"
)

func main() {
	student := make(map[string]int)
	student["Bhavil"] = 100
	student["Roham"] = 80
	student["rahul"] = 120
	student["vimal"] = 100
	fmt.Println(student["Bhavil"])
	// delete(student, "Bhavil")
	fmt.Println(student["Bhavil"])

	grades, exists := student["Bhavil"]
	fmt.Println(grades, exists)

	grade, exist := student["rajesh"]
	fmt.Println("grade", grade, "does exists ?:", exist) //checks whether the key is avail in the map, if not returns false

	for name, grage := range student {
		fmt.Printf("name is %s and marks is %d\n", name, grage)
	}

}
