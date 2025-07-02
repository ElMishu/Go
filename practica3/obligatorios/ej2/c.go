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

var wg sync.WaitGroup

func main() {
	cajas := 3
	clientesTotales := 6
	start := time.Now()

	colas := make([]chan Cliente, cajas) // slice de canales para las cajas
	for i := 0; i < cajas; i++ {
		colas[i] = make(chan Cliente, clientesTotales) // cada canal puede contener hasta clientesTotales clientes
	}

	for i := 0; i < cajas; i++ {
		wg.Add(1)                     // agregar un contador de goroutines para cada caja
		go atenderCaja(i+1, colas[i]) // i+1 para que las cajas arranquen desde 1 y la correspondiente caja
	}

	for i := 0; i < clientesTotales; i++ {
		cliente := Cliente{ID: i + 1}
		mejor := 0
		for j := 1; j < cajas; j++ {
			if len(colas[j]) < len(colas[mejor]) { // comparar la longitud de las colas
				// si la cola j es más corta que la mejor encontrada hasta ahora, actualizar mejor
				mejor = j
			}
		}

		colas[mejor] <- cliente // enviar el cliente a la caja con la cola más corta
		fmt.Printf("Cliente %d asignado a la caja %d\n", cliente.ID, mejor+1)
	}

	// Cerrar las colas
	for i := 0; i < cajas; i++ {
		close(colas[i])
	}

	wg.Wait()
	fmt.Println("Todos los cajeros terminaron de atender.")
	fmt.Println("Todos los clientes fueron atendidos.")
	fmt.Println("Tiempo total:", time.Since(start))
}

func atenderCaja(id int, cola chan Cliente) {
	defer wg.Done()
	for cliente := range cola {
		tiempoAtencion := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Cajero %d atendiendo al cliente %d durante %v\n", id, cliente.ID, tiempoAtencion)
		time.Sleep(tiempoAtencion)
	}
}
