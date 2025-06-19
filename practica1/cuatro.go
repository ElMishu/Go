/*Escriba un programa que imprima en la salida estándar la suma
de los primeros números positivos pares menores o iguales a
250. Cambiar el programa para que itere en el sentido contrario
pero obtener el mismo resultado. Cambiar el programa para que
en lugar de usar un literal como tope se use una constante. Si lo
desea, investigue la herramienta gofmt y pruebe sobre el código
escrito.
Sub-objetivo: Uso de E/S de valores numéricos en Go,
estructuras de control básicas, constantes y variables.*/

package main

import "fmt"

func forcomun() {
	i := 0
	suma := 0
	for i <= 250 {
		if i%2 == 0 {
			suma += i
		}
		i++
	}
	fmt.Println("la suma es ", suma)
}

func downto() {
	i := 250
	suma := 0
	for i >= 1 {
		if i%2 == 0 {
			suma += i
		}
		i--
	}
	fmt.Println("la suma es ", suma)
}

func main() {
	downto()
	forcomun()
}
