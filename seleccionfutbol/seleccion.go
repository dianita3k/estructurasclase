package seleccionfutbol

import "fmt"

type Seleccion struct {
	Nombre string
	Edad   int
}

func (s Seleccion) Viajar() {
	fmt.Println("La Seleccion de Fútbol Viaja")
}
