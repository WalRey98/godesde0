package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "CDMX"
	paises["Argentina"] = "Buenos Aires"

	/*fmt.Println(paises)
	fmt.Println(paises["Argentina"])*/

	campeonato := map[string]int{
		"Barcelona":       39,
		"Real Madrid":     36,
		"Colo Colo":       35,
		"Atletico Madrid": 32,
		"Valencia":        31,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}
	*/

	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Colo Colo"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
