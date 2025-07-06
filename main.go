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

	fmt.Printf("contact %+v\n", contact)
	fmt.Printf("employee %+v\n", employee)

	fmt.Println("Name before:", person.Name)

	person.modifyPersonName("Johan")

	fmt.Println("Name after:", person.Name)

	x := 20
	ptr := &x

	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)

	*ptr = 30

	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)

}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope new name:", p.Name)
}
