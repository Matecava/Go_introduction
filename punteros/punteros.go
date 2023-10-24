package main

import "fmt"

func main() {
	// a := 200
	// b := &a
	// *b++
	// fmt.Println(a, b)
	// fmt.Println(fn())
	v := 1
	increment(&v)
	fmt.Println(v)
	increment(&v)
	fmt.Println(v)
}

// func fn() *int {
// 	v := 1
// 	return &v
// }

func increment(p *int) int {
	*p++
	return *p
}
