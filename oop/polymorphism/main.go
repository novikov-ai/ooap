package main

import "fmt"

// Animal
// Common interface for all structures
type Animal interface {
	Eat()
}

// Dog
// The Dog structure which implemented Eat()
type Dog struct {
	name string
}

func (d Dog) Eat() {
	fmt.Println("Dog", d.name, "is eating")
}

// Cat
// The Cat structure which implemented Eat()
type Cat struct {
	name string
}

func (c Cat) Eat() {
	fmt.Println("Cat", c.name, "is eating")
}

func main() {
	// create an Animal interface
	var animal Animal

	// initialize a dog struct
	dog := Dog{name: "Rex"}
	// initialize a cat struct
	cat := Cat{name: "Bars"}

	// assign var dog with animal interface
	animal = dog
	animal.Eat()

	// assign var cat with animal interface
	animal = cat
	animal.Eat()

	// got polymorphic behaviour
}
