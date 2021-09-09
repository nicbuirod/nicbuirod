package main

import (
	"fmt"

	Perro "gitlab.com/aprendiendo-go/ejercicios12/perro"
)

type canino struct {
	Nombre string
	Edad   int
}

func main() {
	c1 := canino{
		Nombre: "Sonic",
		Edad:   Perro.Edadperro(2),
	}
	fmt.Println(c1)
}
