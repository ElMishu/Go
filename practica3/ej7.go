/*7) Realice un programa que envíe datos a dos canales desde dos
goroutines y estos sean recibidos en el programa principal por un
select. Los datos se deben recibir en uno de los canales por el lapso de
5 segundos y por el otro en el lapso de 10 segundos */

package main

import (
	"fmt"
	"time"
)

func enviar(ch chan int, secuencia []int) {
	for _, v := range secuencia {
		ch <- v
		time.Sleep(2 * time.Second) // Simula un tiempo de envío de 1 segundo por valor
	}
	close(ch) // Cierra el canal al finalizar el envío de la secuencia
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go enviar(ch1, []int{1, 2, 3, 4, 5})
	go enviar(ch2, []int{6, 7, 8, 9, 10})

	timeout1 := time.After(5 * time.Second)
	timeout2 := time.After(10 * time.Second)

	for {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Println("Recibido de ch1:", v)
			} else {
				fmt.Println("Canal ch1 cerrado")
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Println("Recibido de ch2:", v)
			} else {
				fmt.Println("Canal ch2 cerrado")
				ch2 = nil
			}
		case <-timeout1:
			fmt.Println("Tiempo finalizado para el primer canal")
			ch1 = nil
		case <-timeout2:
			fmt.Println("Tiempo finalizado para el segundo canal")
			return
		}
	}
}
