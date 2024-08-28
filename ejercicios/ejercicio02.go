/*
1- Crear un archivo GO en el paquete 'ejercicios'
llamado 'ejercicio02.go'
2-Crear una funcion publica que pida por teclado un numero valide si hay
error o no (si hay error que vuelva a pedirlo)
3- En base al numero recibido crear la tabla numerica del mismo de 1 a 10
y la muestre por pantalla
4- En main llamar a dicha funcion
5- Grabar, ejecutar y luego subir a github
*/

package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)

	}

	return texto
}
