package seleccionfutbol

import "fmt"

type Entrenador struct {
	Seleccion
	Federacion string
}

func (e Entrenador) DirigirPartido() {
	fmt.Printf("El entrenador %s Dirige Partido \n", e.Nombre)
}
