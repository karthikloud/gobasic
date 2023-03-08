package main

import "fmt"

func main() {
	fmt.Println("hi hello")

	//value

	fmt.Println("123+43:", 123+43)
	fmt.Println("17.2/2.3:", 17.2/2.3)

	//  variables
	var name string = "Alice"
	var age int = 25
	var isMarried bool = false
	var height float64 = 5.5

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Marital status:", isMarried)
	fmt.Println("Height:", height)

	//contant
	const m string = "karthi"
	fmt.Println(m)

	// for
	for a := 0; a < 10; a++ {
		fmt.Print(a)

		//if/else
		a := 10
		if a == 10 {
			fmt.Printf("a is equal to 10")
		} else {
			fmt.Printf("a is not equal to 10")
		}
		//arrays
		var num [5]int
		num[0] = 10
		num[1] = 20
		num[2] = 30
		num[3] = 40
		num[4] = 50

		fmt.Println("array of num:", num)
	}
	//switch
	var dayOfWeek int = 5

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("invalid day")
		//slice
		s := make([]string, 3)
		fmt.Println("emp:", s)
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])
		fmt.Println("length:", len(s))
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

	}
}
