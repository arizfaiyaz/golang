package main

import "fmt"

func sum(a int, b int) int {

	// c = a + b and then return c
	c := a + b
	return c
}

func plus(a,b,c int) int {
	return a + b + c
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)
	res2 := plus(1, 2, 3)
	fmt.Println(res2)
}