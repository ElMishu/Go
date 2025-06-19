/*5) Realice un programa que tenga 2 productores y 2 consumidores. Cada
productor y consumidor debe ser una Goroutine. Cada productor
producirá 3 números y cada consumidor consumirá 3 números. Los
productores deben esperar un tiempo aleatorio entre 0 y 1 segundo
para producir un número aleatorio entre 0 y 100. Los consumidores
deben consumirlos inmediatamente e imprimirlos por pantalla indicando
cual es el consumidor que lo consumió.
Objetivo: WaitGroup*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 6) // Canal con capacidad para 6 elementos (3 de cada productor)
	var wg sync.WaitGroup   // WaitGroup para esperar a que terminen todas las goroutines
	wg.Add(4)               // 2 productores + 2 consumidores
	go productor(1, ch, &wg)
	go productor(2, ch, &wg)
	go consumidor(1, ch, &wg)
	go consumidor(2, ch, &wg)
	wg.Wait()
	fmt.Println("Todos los productores y consumidores han terminado.")
}

func productor(id int, ch chan<- int, wg *sync.WaitGroup) {
	// Cada vez que una goroutine termina, llama a wg.Done(), que resta 1 al contador.
	// El wg.Done, cuando se ejecuta, le resta uno al contador de wg de goroutines activas
	defer wg.Done() // el defer pospone la ejecución de la función wg.Done hasta que la función que lo contiene termina.
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // duerme entre 0 y 1 segundo
		num := rand.Intn(101)                                         // Genera un número entre 0 y 100
		fmt.Printf("Productor %d produjo: %d\n", id, num)
		ch <- num //Manda el número al canal (ch <- num).
	}
}

func consumidor(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		num := <-ch // recibe un número si hay en el canal, sino se queda esperando hasta que haya un número
		fmt.Printf("Consumidor %d consumió: %d\n", id, num)
	}
}
