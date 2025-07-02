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

const (
	clientesTotales = 100
	cajas           = 3
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	// Crear una cola por caja
	colas := make([]chan Cliente, cajas) // slice de canales para las cajas
	for i := 0; i < cajas; i++ {
		colas[i] = make(chan Cliente, clientesTotales) // cada canal puede contener hasta clientesTotales clientes
	}

	// Lanzar cajeros
	for i := 0; i < cajas; i++ {
		wg.Add(1)
		go atenderCaja(i+1, colas[i]) // i+1 para que las cajas arranquen desde 1
	}

	// Enviar clientes por round-robin
	for i := 0; i < clientesTotales; i++ {
		cliente := Cliente{ID: i + 1}
		cajaAsignada := i % cajas                                                    // round-robin si hay 3 cajas, el cliente 1 va a la caja 1, el 2 a la caja 2, el 3 a la caja 3, el 4 a la caja 1, etc.
		colas[cajaAsignada] <- cliente                                               // enviar el cliente a la caja asignada
		fmt.Printf("Cliente %d asignado a la caja %d\n", cliente.ID, cajaAsignada+1) // +1 para que las cajas arranquen desde 1
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
