package main

import (
	"fmt"
)

func main() {
	age := 30

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("its the weekend")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("This is i", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("This is the counter", counter)
		counter++
	}

	iterations := 0
	for {
		if iterations > 3 {
			break
		}
		iterations++
	}

	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50}
	numbers[0] = 60 //reassign the first position

	fmt.Printf("this is our array %v\n", numbers)
	fmt.Printf("this is our array %v\n", numbers[2])
	fmt.Printf("this is our array %v\n", len(numbers))

	fmt.Println("this is the last value", numbers[len(numbers)-1])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("this is our matrix %v\n", matrix)

	allNumbers := numbers[:]
	firstThree := numbers[0:3

	allNumbers

}

func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
