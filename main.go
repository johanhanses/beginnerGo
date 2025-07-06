package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 25}

	fmt.Printf("This is our person %+v\n", person)

	employee := struct {
		Name string
		id   int
	}{
		Name: "Jane",
		id:   33,
	}

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Mark",
		Address: Address{
			Street: "123 Main street",
			City:   "Anytown",
		},
	}

	fmt.Printf("%+v\n", contact)

	fmt.Printf("employee %+v\n", employee)
}
