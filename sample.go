package main

import "fmt"

func main() {

	//array
	array := []int{1, 23, 45, 56, 54, 56, 78, 89, 90}
	sum := 0
	var total int
	for k := range array {
		sum += k
		total = sum
		fmt.Println(total)

	}

	//map
	hm := map[string]int{"siva": 76, "sanjay": 45}
	hm["dinesh"] = 87
	fmt.Println(hm)
	fmt.Println(hm["siva"])
	delete(hm, "dinesh")
	fmt.Println(hm)
	v, ok := hm["vengat"]
	fmt.Println(v, ok)
	//pointer
	ptr := new(int)
	*ptr = 708
	fmt.Println(*ptr)

	//slice
	number := []int{1, 3, 5, 6, 7}
	fmt.Println("number")
	number = append(number, 8)
	println(number)
	sliceNumber := number[2:4]
	fmt.Println(sliceNumber)
	sliceNumber[0] = 28
	fmt.Println(sliceNumber)
	fmt.Println(number)

	//range
	ages := map[string]int{
		"siva": 25,
		"arun": 30,
		"mano": 35,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

}
