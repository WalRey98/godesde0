package main

import (
	"fmt"

	"github.com/walRey98/godesde0/gorutines"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/

	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("Esto no es Windows")
		} else {
			fmt.Println("Esto es Windows")
		}
		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	}*/

	/*numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)
	*/

	/*teclado.IngresoNumeros()*/

	//fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()

	//funciones.LlamarClousure()

	//funciones.Exponencia(2)

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	/*Pedro := new(modelos.Hombre)  // e "github.com/walRey98/godesde0/ejer_interfaces" "github.com/walRey98/godesde0/modelos"

	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//defer_panic.EjemploPanic()

	canal1 := make(chan bool)
	go gorutines.MiNombreLentooo("Walter Reyes", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")

}
