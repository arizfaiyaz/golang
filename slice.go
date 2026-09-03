package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit: ", s , s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: s", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("slice", l)

	l = s[:5]
	fmt.Println("slice", l)

	l = s[2:]
	fmt.Println("slice", l)

	t := []string{"g", "h", "i"} 
	fmt.Println("declare slice:- ", t)

	t2 := []string{"j", "k", "l"}
	if slices.Equal(t, t2){
		fmt.Println("t1 == t2")
	} else {
		fmt.Println("t1 != t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j:= range innerlen {
			twoD[i][j] = j + 1
		}
	}
	fmt.Println("twoD", twoD)
}