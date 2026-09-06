package main

import "fmt"

type person struct {
	name string
	age  int
}
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	//You can name the fields when initializing a struct.
	
    fmt.Println(person{name: "Alice", age: 30})
	//Omitted fields will be zero-valued.
	
    fmt.Println(person{name: "Fred"})
	//An & prefix yields a pointer to the struct.
	
    fmt.Println(&person{name: "Ann", age: 40})
	//It's idiomatic to encapsulate new struct creation in constructor functions
	
    fmt.Println(newPerson("Jon"))

    s := person{name: "sean", age: 50}
    fmt.Println(s.name)
    // fmt.Println(s.age)

    sp := &s
    fmt.Println(sp.age)
    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name string
        isGood bool
    }{
        "trex",
        true,
    }
    fmt.Println(dog)
}
// struct are mutable
// access struct fiels with dot notation
//  ommited fields will be zero-valued
// its idiomatc to encapsulate struct creation in constructor functions
// 