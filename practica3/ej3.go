package main

import "fmt"

func main() {
	done := make(chan bool) // Canal para sincronización
	fmt.Println("Inicia Goroutine del main")
	go hello(done)
	<-done // Espera a que la Goroutine hello termine
	fmt.Println("Termina Goroutine del main")

}

func hello(done chan bool) {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
	done <- true
}

/*
a) ¿Cuántas veces se imprime Hello world?
Se imprime 0 veces porque la Goroutine del main termina antes de que la Goroutine hello tenga tiempo de ejecutarse completamente.

b) ¿Cuántas Goroutines tiene el programa?
El programa tiene 2 Goroutines: una del main y otra de hello.

c) ¿Cómo cambiaría el programa (con la misma cantidad de Goroutines) para que imprima 3 veces Hello world?
i) Hágalo usando time.Sleep
func main() {
	fmt.Println("Inicia Goroutine del main")
	go hello()
	fmt.Println("Termina Goroutine del main")
	time.Sleep(1 * time.Second) // Espera 1 segundo para que la Goroutine hello tenga tiempo de ejecutarse
}
ii) Hágalo usando Channel Synchronization
func main() {
	done := make(chan bool) // Canal para sincronización
	fmt.Println("Inicia Goroutine del main")
	go hello(done)
	<-done  // Espera a que la Goroutine hello termine
	fmt.Println("Termina Goroutine del main")

}

func hello(done chan bool) {
	defer done <- true // Envía true al canal cuando la Goroutine hello termine
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
}


*/
