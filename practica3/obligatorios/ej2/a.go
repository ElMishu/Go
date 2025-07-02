/*2) Realice un programa que simule la atención de clientes en las cajas de
un supermercado. La atención de cada cliente por parte del cajero se
debe simular con un timer entre 0 y 1 segundo.
a) Realice el programa haciendo esperar a los clientes en un única
cola global y luego enviándolo

d) Imprima los tiempos de ejecución de cada uno de los programas
implementados en a), b) y c)
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	ID int
}

var colaClientes = make(chan Cliente, 100)
var wg sync.WaitGroup

func main() {
	cajas := 3
	clientesTotales := 100 // cantidad de clientes que se van a agregar a la cola

	start := time.Now() // inicia el cronómetro para medir el tiempo de ejecución

	// Simular la llegada de clientes
	for i := 1; i <= clientesTotales; i++ {
		cliente := Cliente{ID: i}
		colaClientes <- cliente
		fmt.Printf("Cliente %d en la cola\n", cliente.ID) // para chequear que lo agregue a la cola
	}

	// Cerrar la cola de clientes
	close(colaClientes)

	// Iniciar los cajeros
	for i := 1; i <= cajas; i++ {
		wg.Add(1)
		go atenderClientes(i)
	}

	wg.Wait() // espera a que todos los cajeros terminen
	fmt.Println("todos los cajeros terminaron de atender ")
	fmt.Println("todos los clientes fueron atendidos")
	elapsed := time.Since(start)
	fmt.Println("Tiempo total:", elapsed)
}

func atenderClientes(id int) {
	defer wg.Done()
	for cliente := range colaClientes {
		tiempoAtencion := time.Duration(rand.Intn(1000)) * time.Millisecond // simula el tiempo de atención entre 0 y 1 segundo
		fmt.Printf("cajero %d atendiendo al cliente %d durante %v\n", id, cliente.ID, tiempoAtencion)
		time.Sleep(tiempoAtencion)
	}
}
