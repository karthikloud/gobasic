package main

import "fmt"

func main() {

	day := "monday"

	switch day {
	case "monday":
		fmt.Println("today is monday")
	case "tuesday":
		fmt.Println("today is tuesday")
	case "wednesday":
		fmt.Println("today is wednesday")
	case "thursday":
		fmt.Println("today is thursday")
	case "friday":
		fmt.Println("today is friday")
	case "saturday":
		fmt.Println("it is weekend")
	default:
		fmt.Println("invalid data")
	}

	switch {

	case day == "monday":
		fmt.Println("week days")

	}
}
