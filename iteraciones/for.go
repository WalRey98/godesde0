package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i++ {
		if i == 3 { // si llega hasta 3 breack osea para
			break // si quiero que se salte un numero se coloca 'continue'
		}
		fmt.Println(i)
	}
}
