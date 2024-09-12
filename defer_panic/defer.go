package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")

	/*el defer es una instruccion que nos da go que nos permite configurar
	cual va a ser la ultima instruccion en ejecutarse cuando salga de la funcion*/

}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
