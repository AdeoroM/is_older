package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastName string
	age      int
}

func isAdult(person Person) bool {
	return person.age >= 18
}
func main() {
	people := []Person{
		Person{
			name:     "Andres",
			lastName: "De Oro",
			age:      20,
		},
		Person{
			name:     "Juliana",
			lastName: "Chois",
			age:      17,
		},
		Person{
			name:     "Daniel",
			lastName: "Muriel",
			age:      24,
		},
	}
	for i := 0; i < len(people); i++ {
		if isAdult(people[i]) {
			fmt.Printf("%v %v - es mayor de edad\n", people[i].name, people[i].lastName)
		} else {
			fmt.Printf("%v %v - es menor de edad\n", people[i].name, people[i].lastName)
		}
	}
}
