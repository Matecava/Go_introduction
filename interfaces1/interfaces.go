package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
}

func (admin Admin) Permisos() int {
	return 5
}

func (admin Admin) Nombre() string {
	return admin.nombre
}

type Editor struct {
	nombre string
}

func (editor Editor) Permisos() int {
	return 3
}

func (editor Editor) Nombre() string {
	return editor.nombre
}

func auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de admin"
	} else if permisos == 3 {
		return user.Nombre() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	usuarios := []User{Admin{"Mateo"}, Editor{"Fulano"}}

	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
