/*
	crea una funcion publica que devuelva 2 valores

(int y string) y que reciba de parametro un valor 'string'
2- el parametro de tipo string recibido debera ser convertido a entero y si el entero es > a 100, el string retornado debe decir "Es mayor a 100"
de lo contrario devolver el mensaje "Es menor a 100"
5- El valor numerico entero retornado debe corresponder al string convertido
6- En main llamar a dicha afuncion asignandola a 2 variables y luego mostrar dichas variables por consola
7- grabar y ejecutar y subir a github
*/
package ejercicios

import (
	"strconv"
)

func ConvNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
