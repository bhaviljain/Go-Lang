package main

import (
	"fmt"
	"time"
)

// 02-01.2006 is date ko go launch hua tha aur time :15:04:05
func main() {
	current := time.Now()
	fmt.Println(current)
	// format := current.Format("2006-01-02, Monday, 15:04:05")
	// format := current.Format("2006-01-02, Monday, 3:04 PM") //pm HE LIKHA SIRF dono ke liye am and pm ke liye
	format := current.Format("2006-01-02, Monday, 15:04 PM") //pm HE LIKHA SIRF dono ke liye am and pm ke liye
	fmt.Println(format)

	layout := "02/01/2006"
	datestr := "25/11/2000"
	formatted_Time, _ := time.Parse(layout, datestr)
	fmt.Println(formatted_Time)
	newdate := current.Add(24 * time.Hour) //adding 24 hours to the current day
	fmt.Println(newdate)
	format_data := newdate.Format("02/01/2006, Monday")
	fmt.Println(format_data)
}
