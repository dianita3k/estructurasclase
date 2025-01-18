package seleccionfutbol

import "fmt"

type Futbolista struct {
	Seleccion
	Dorsal int
}

func (f Futbolista) JugarPartido() {

	fmt.Printf("El futbolista %S Jueva un partido\n", f.Nombre)

}
