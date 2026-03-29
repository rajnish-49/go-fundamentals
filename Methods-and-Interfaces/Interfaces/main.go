package main

import "fmt"

// An interface is just a set of method signatures.
// anything that has Speak() method is an Animal
type Animal interface {
	Speak()
}

// Dog struct
type Dog struct{ name string }

// Cat struct
type Cat struct{ name string }

// Dog satisfies Animal — has Speak()
func (d Dog) Speak() { fmt.Println("woof") }

// Cat satisfies Animal — has Speak()
func (c Cat) Speak() { fmt.Println("meow") }

// one function works for ANY Animal
// doesn't care if it's Dog, Cat, or any future animal
func makeSpeak(a Animal) {
	a.Speak()
}

func main() {
	d := Dog{"rex"}
	c := Cat{"tom"}

	fmt.Println("Dog:")
	makeSpeak(d)

	fmt.Println("Cat:")
	makeSpeak(c)

	// interface variable — can hold any Animal
	var a Animal

	a = Dog{"buddy"}
	fmt.Println("\nInterface variable holding Dog:")
	a.Speak()

	a = Cat{"whiskers"}
	fmt.Println("Interface variable holding Cat:")
	a.Speak()
}