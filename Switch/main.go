package main

import "fmt"

func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("tuesday")

	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("unknown day")
	}
	month := "sept"

	switch month {
	case "january", "feb", "march":
		fmt.Println("winter")
	case "April", "may", "june":
		fmt.Println("summer")
	case "july", "august", "sept":
		fmt.Println("rainy")
	}
	temp := 30
	switch {
	case temp < 0:
		fmt.Println("verycold")
	case temp > 0 && temp <= 10:
		fmt.Println("Cold")
	case temp >= 10 && temp < 15:
		fmt.Println("Winds")
	case temp > 15 && temp < 25:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
