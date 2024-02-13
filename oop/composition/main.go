package main

import "fmt"

type Dog struct {
	Animal
}

type Animal struct {
	name   string
	weight int
}

func (a Animal) Eat() {
	fmt.Println(a.name, "is eating")
}

func main() {
	dog := Dog{}
	dog.name = "Rex"
	dog.weight = 200

	dog.Eat()
}
