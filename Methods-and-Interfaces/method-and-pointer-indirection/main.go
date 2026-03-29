package main

import "fmt"

type Person struct {
    age int
}

// regular function — strict
func birthdayFunc(p *Person) {
    p.age += 1
}

// method — flexible
func (p *Person) birthdayMethod() {
    p.age += 1
}

func main() {
    john := Person{age: 25}
    // birthdayFunc(john)     // COMPILE ERROR — must pass pointer
    birthdayFunc(&john)       // OK — manually pass pointer
    fmt.Println(john.age)     // 26

    ron := Person{age: 30}
    ron.birthdayMethod()      // OK — Go adds & automatically
    ptr := &ron
    ptr.birthdayMethod()      // OK — already a pointer
    fmt.Println(ron.age)      // 32
}