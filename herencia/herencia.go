package main

import "fmt"

type Curso struct {
	habilidad   []string
	nombre, url string
}

func main() {
	// curso1 := Curso{
	// 	nombre:    "Go",
	// 	url:       "GoCurso",
	// 	habilidad: []string{"backend", "2"},
	// }
	// fmt.Println(curso1)
	// curso1.Inscribirse("Mateo")

	carrera1 := new(Carrera)

	carrera1.nombreCarrera = "Administracion de empresas y sistemas"
	carrera1.duracion = 5
	carrera1.nombre = "DB"
	carrera1.url = "BasesDeDatos"
	carrera1.habilidad = []string{"Programacion", "Bases de datos"}

	carrera1.Inscribirse("Mateo Cava")

	fmt.Println(carrera1)
}

func (c Curso) Inscribirse(nombre string) {
	fmt.Printf("La persona %s se ha inscrito al curso %s \n", nombre, c.nombre)
}

type Carrera struct {
	nombreCarrera string
	duracion      int
	Curso
}
