package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	// mateo := User{edad: 19 ,nombre: "Mateo", apellido: "Cava"}
	mateo := new(User)

	mateo.nombre = "Mateo"
	mateo.apellido = "Cava"
	mateo.edad = 19

	fmt.Println(mateo)

}
