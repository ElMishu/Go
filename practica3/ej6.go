/*
 6. Realice un programa que utilice select para recibir valores desde tres

canales diferentes. Cada canal debe recibir una secuencia de números
enteros. El programa debe recibir un valor de cada canal y mostrar el
valor recibido. Debes usar select para determinar cuál canal tiene un
valor disponible y recibir ese valor. El programa debe continuar hasta
haber recibido todos los valores enviados a cada canal. Al final debe
mostrar el total de valores recibidos desde cada canal.
Objetivo: select
*/
package main

import (
	"fmt"
)

func enviarSecuencias(ch chan int, secuencia []int) {
	for _, v := range secuencia {
		ch <- v
	}
	//close(ch) // Cierra el canal al finalizar el envío de la secuencia
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	contador1 := 0
	contador2 := 0
	contador3 := 0

	go enviarSecuencias(ch1, []int{1, 2, 3, 4, 5})
	go enviarSecuencias(ch2, []int{6, 7, 8, 9, 10})
	go enviarSecuencias(ch3, []int{11, 12, 13, 14, 15})

	for contador1 < 5 || contador2 < 5 || contador3 < 5 {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Println("Enviado 1 a ch1:", v)
				contador1++
			} else {
				fmt.Println("Canal ch1 cerrado")
			}

		case v, ok := <-ch2:
			if ok {
				fmt.Println("Enviado 2 a ch2:", v)
				contador2++
			} else {
				fmt.Println("Canal ch2 cerrado")
			}
		case v, ok := <-ch3:
			if ok {
				fmt.Println("Enviado 3 a ch3:", v)
				contador3++
			} else {
				fmt.Println("Canal ch3 cerrado")
			}
		}
	}
	fmt.Printf("Total valores recibidos desde ch1: %d\n", contador1)
	fmt.Printf("Total valores recibidos desde ch2: %d\n", contador2)
	fmt.Printf("Total valores recibidos desde ch3: %d\n", contador3)
}
