package main

import "fmt"

func main() {
	// matriz := []int{1, 2}

	// if matriz == nil {
	// 	fmt.Println("Esta vacio")
	// } else {
	// 	fmt.Println(len(matriz))
	// }

	arreglo := [3]int{1, 2, 3}
	fmt.Println(arreglo[2])
	slice := arreglo[0:1]
	fmt.Println(slice)
}
