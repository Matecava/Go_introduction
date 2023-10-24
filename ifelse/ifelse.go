package main

import "fmt"

func main() {

	var n int

	fmt.Print("Ingrese su numero: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(n, "Es neutro")
	} else if n%2 == 0 {
		fmt.Println(n, "Es par")
	} else {
		fmt.Println(n, "Es impar")
	}
	// var (
	// 	nombre string
	// 	edad   int
	// 	pi     float64
	// )

	// fmt.Print("Ingrese su Nombre: ")
	// fmt.Scanln(&nombre)

	// fmt.Print("Ingrese su Edad: ")
	// fmt.Scanln(&edad)

	// fmt.Print("Ingrese el Valor de Pi: ")
	// fmt.Scanln(&pi)

	// fmt.Printf("Nombre: %s Edad: %d \nValor de Pi: %f", nombre, edad, pi)
	// var (
	// 	nombre string
	// )
	// fmt.Print("mi nombre es: ")
	// fmt.Scanln(&nombre)

	// fmt.Printf("Nombre: %s ", nombre)

}
