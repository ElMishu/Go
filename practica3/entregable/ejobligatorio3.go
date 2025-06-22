package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Tarea struct {
	Prioridad int
	Valor     int
}

var (
	wg             sync.WaitGroup
	mutexArchivo0  sync.Mutex
	mutexArchivo1  sync.Mutex
	mutexAcumulado sync.Mutex
	acumulado      int
)

func main() {
	const canttareas int = 20 // quiero que se generen aleatoriamente X tareas
	var tareas []Tarea        // creo slice en donde voy a guardar las tareas a generar
	for i := 0; i < canttareas; i++ {
		valor := rand.Intn(500) + 1                 // genera numeros aleatorios entre 1 y 500
		prio := rand.Intn(4)                        // genera prioridades entre 0 y 3
		tareas = append(tareas, Tarea{prio, valor}) // agrega la tarea al slice de tareas
	}

	for _, t := range tareas { // muestra las tareas generadas
		fmt.Printf("Tarea generada -> Prioridad: %d, Valor: %d\n", t.Prioridad, t.Valor)
	}

	tareaCh := make(chan Tarea) // crea el canal donde se van a ir mandando las tareas

	// Inicializo wg con la cantidad total de tareas
	wg.Add(len(tareas))

	// Iniciar a los 4 workers
	for i := 1; i <= 4; i++ {
		go worker(i, tareaCh, &wg) // i= numero de identificador del worker, le mando el canal de tareas y el &wg para luego avisar que finalizo
	}

	// Scheduler: envía tareas en orden de prioridad
	for prioridad := 0; prioridad <= 3; prioridad++ { // itera prioridades de 0 a 3
		for _, tarea := range tareas { // recorre el slice de tareas
			if tarea.Prioridad == prioridad { // si la prioridad de la tarea es igual a la prioridad que estoy iterando
				tareaCh <- tarea // envía la tarea al canal para que un worker la procese
			}
		}
	}
	close(tareaCh) // cierra el canal de tareas para que los workers sepan que no hay más tareas

	wg.Wait()
	fmt.Println("Todas las tareas fueron procesadas.")
}

func worker(id int, tareas <-chan Tarea, wg *sync.WaitGroup) { // worker
	for tarea := range tareas { // recibe tareas del canal
		switch tarea.Prioridad { // chequea la prioridad de la tarea

		/* prioridad 0= Sumar los dígitos del número y almacenar el resultado en un archivo llamado "prioridad0.txt" en formato
		de par ordenado (0, resultado)
		*/

		case 0:
			resultado := sumarDigitos(tarea.Valor)                                                // suma los dígitos del número
			mutexArchivo0.Lock()                                                                  // bloquea el acceso a el archivo para evitar conflictos de escritura
			file, err := os.OpenFile("prioridad0.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // abre el archivo en modo append, crea si no existe y con permisos de escritura
			if err == nil {                                                                       // si no hay error al abrir el archivo
				fmt.Fprintf(file, "(0, %d)\n", resultado) // escribe el resultado en el archivo
				file.Close()                              // cierra el archivo
			} else {
				fmt.Println("Error escribiendo prioridad0.txt:", err) // si hay error al abrir el archivo, muestra un mensaje de error
			}
			mutexArchivo0.Unlock() // desbloquea el acceso a el archivo

		case 1: // prioridad 1= Invertir los dígitos del número y almacenar el resultado en un archivo llamado "prioridad1.txt" en formato de par ordenado (1, resultado)
			resultado := invertirDigitos(tarea.Valor)                                             // invierte los dígitos del numero
			mutexArchivo1.Lock()                                                                  // bloquea el acceso a el archivo para evitar conflictos de escritura
			file, err := os.OpenFile("prioridad1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // abre el archivo en modo append, crea si no existe y con permisos de escritura
			if err == nil {                                                                       // si no hay error al abrir el archivo
				fmt.Fprintf(file, "(1, %d)\n", resultado) // escribe el resultado en el archivo
				file.Close()                              // cierra el archivo
			} else {
				fmt.Println("Error escribiendo prioridad1.txt:", err) // si hay error al abrir el archivo, muestra un mensaje de error
			}
			mutexArchivo1.Unlock() // desbloquea el acceso a el archivo

		case 2: // prioridad 2= Multiplicar el número por 10 y mostrar el resultado en la consola
			resultado := tarea.Valor * 10                                       // multiplica el número por 10
			fmt.Printf("Prioridad 2 -> %d * 10 = %d\n", tarea.Valor, resultado) // muestra el resultado en la consola

		case 3: // prioridad 3= Acumular el valor de la tarea en una variable compartida y mostrar el acumulado en la consola
			mutexAcumulado.Lock()                                   // bloquea el acceso a la variable acumulado para evitar conflictos de escritura
			acumulado += tarea.Valor                                // acumula el valor de la tarea
			fmt.Printf("Prioridad 3 -> acumulado: %d\n", acumulado) // muestra el acumulado en la consola
			mutexAcumulado.Unlock()                                 // desbloquea acceso a la variable acumulado
		}
		time.Sleep(200 * time.Millisecond) // simulacion de trabajo
		wg.Done()
	}
}

func sumarDigitos(n int) int { //utilizada en datos de prioridad 0
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(n int) int { //utilizada en datos de prioridad 1
	invertido := 0
	for n > 0 {
		invertido = invertido*10 + n%10
		n /= 10
	}
	return invertido
}
