package main

import "fmt"

func main() {
	/*
		copia el minimo de elementos de ambos arreglos
	*/
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice))

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
