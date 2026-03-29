package main 

// fmt says — "if your type has a method called String() that returns a string,
//  I will call it when printing instead of the default ugly output."

import (
	"fmt"
)

// Stringer interface lets you control how your type prints
// String() method is called when you print a value of this type

type Person struct {
	name string
	age  int
}

// It's just an interface — implicit like everything else
// You don't say "Person implements Stringer". You just define String() on Person and fmt automatically uses it.


func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	p := Person{name: "Alice", age: 30}
	fmt.Println(p)
	fmt.Println("john is 25")           // prints directly
    fmt.Sprintf("%v is %v", "john", 25) // returns a string, doesn't print
}