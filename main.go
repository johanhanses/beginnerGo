package main

import (
	"fmt"
)

func main() {
	// A comment
	var name string = "Johan"
	fmt.Printf("This is my name %s\n", name)

	age := 41
	fmt.Printf("This is my age %d\n", age)

	var city string
	city = "Falun"
	fmt.Printf("This is my city %s\n", city)

	var country, continent string = "Sweden", "Europe"
	fmt.Printf("This is my country %s and this is my continent %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("isEmployed: %t this is my salary: %d and this is my position %s\n", isEmployed, salary, position)
}
