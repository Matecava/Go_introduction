package main

import "fmt"

type Persona struct {
	nombre           string
	edad             int
	numero_documento int
}

func (persona Persona) GetName() string {
	return persona.nombre
}

func (persona Persona) GetEdad() int {
	return persona.edad
}

func (persona Persona) GetNumDoc() int {
	return persona.numero_documento
}

func main() {
	persona1 := new(Persona)

	persona1.nombre = "Mateo"
	persona1.edad = 19
	persona1.numero_documento = 12345678

	persona1.Mostrar("Mateo", 19, 12345678)

	fmt.Println(persona1)
}
func (c Persona) Mostrar(nombre string, edad, numero_documento int) {
	fmt.Printf("Nombre: %s, edad:  %d, numero de documento: %d \n", nombre, edad, numero_documento)
}
