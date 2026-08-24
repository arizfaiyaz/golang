package main

import "fmt"

func main() {
	var a = "initialized"
	fmt.Println(a)

	var b,c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := 7  // this is the true way to declare a variable in go easier and simpler syntax
	fmt.Println(f)
}
