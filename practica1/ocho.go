/*8. Realizar un programa que lea el punto cardinal (como caracter o
string) del cual viene el viento (‘N’, ‘S’, ‘E’, ‘O’) y envíe a la salida
estándar hacia cuál se dirigiría.
Sub-objetivo: Uso de case con la opción por default. E/S
caracteres o strings.
a. ¿Cómo se escribe el default en el case de otros lenguajes?
*/

package main

import "fmt"

func main() {
	var c string
	fmt.Println("Ingrese el punto cardinal")
	fmt.Scan(&c)
	switch c {
	case "N":
		fmt.Println("Se dirige hacia el norte")
	case "S":
		fmt.Println("Se dirige hacia el sur")
	case "O":
		fmt.Println("Se dirige hacia el oeste")
	case "E":
		fmt.Println("Se dirige hacia el este")
	default:
		fmt.Println("No se ingreso un punto cardinal")
	}
}
