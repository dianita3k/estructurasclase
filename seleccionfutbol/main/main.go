package main

import (
	"estructurasclase/seleccionfutbol"
	"fmt"
)

func main() {

	entrenador := new(seleccionfutbol.Entrenador)
	entrenador.Nombre = "Alex"
	entrenador.Edad = 25

	entrenador.Federacion = "FPF"
	fmt.Println(entrenador)
	entrenador.DirigirPartido

	jugador1 := new(seleccionfutbol.Futbolista)
	jugador1.Nombre = "Roel"
	jugador1.Edad = 24
	jugador1.Dorsal = 10
	fmt.Println(jugador1)
	jugador1.JugarPartido

}
