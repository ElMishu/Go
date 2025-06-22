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
	const canttareas int = 5                       // quiero que se generen aleatoriamente X tareas
	var tareas []Tarea                             // creo slice en donde voy a guardar las tareas a generar
	os.WriteFile("prioridad0.txt", []byte{}, 0644) // crea el archivo prioridad0.txt si no existe y lo vacía
	os.WriteFile("prioridad1.txt", []byte{}, 0644) // crea el archivo prioridad1.txt si no existe y lo vacía

	for i := 0; i < canttareas; i++ {
		valor := rand.Intn(500) + 1
		prio := rand.Intn(4)
		tareas = append(tareas, Tarea{prio, valor}) // agrega la tarea al slice de tareas
	}
	// muestra las tareas generadas
	fmt.Printf("Se generaron %d tareas aleatorias:\n", len(tareas))
	for _, t := range tareas {
		fmt.Printf("Tarea Prioridad: %d, Valor: %d\n", t.Prioridad, t.Valor)
	}

	tareaCh := make(chan Tarea) // crea el canal donde se van a ir mandando las tareas

	// Iniciar a los 4 workers
	for i := 1; i <= 4; i++ {
		go worker(i, tareaCh, &wg) // i= numero de identificador del worker, le mando el canal de tareas y el &wg para luego avisar que finalizo
	}

	// Scheduler:
	//  envía tareas en orden de prioridad
	for prioridad := 0; prioridad <= 3; prioridad++ {
		for _, tarea := range tareas {
			if tarea.Prioridad == prioridad {
				wg.Add(1)
				tareaCh <- tarea
			}
		}
		wg.Wait() // Espera a que terminen TODAS las tareas de esta prioridad
	}
	close(tareaCh)
	fmt.Println("Todas las tareas fueron procesadas.")
	mostrarContenidoArchivo("prioridad0.txt")
	mostrarContenidoArchivo("prioridad1.txt")
}

func worker(id int, tareas <-chan Tarea, wg *sync.WaitGroup) {
	for tarea := range tareas { // recibe tareas del canal
		fmt.Printf("Worker %d procesando tarea = Prioridad: %d, Valor: %d\n", id, tarea.Prioridad, tarea.Valor)
		switch tarea.Prioridad {

		/* prioridad 0= Sumar los dígitos del número y almacenar el resultado en un archivo llamado "prioridad0.txt" en formato
		de par ordenado (0, resultado)
		*/
		case 0:
			resultado := sumarDigitos(tarea.Valor)
			mutexArchivo0.Lock()                                                                  // bloquea el acceso a el archivo para evitar conflictos de escritura
			file, err := os.OpenFile("prioridad0.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // abre el archivo en modo append, crea si no existe y con permisos de escritura
			if err == nil {                                                                       // si no hay error al abrir el archivo
				fmt.Fprintf(file, "(0, %d)\n", resultado) // escribe el resultado en el archivo
				file.Close()
			} else {
				fmt.Println("Error escribiendo prioridad0.txt:", err) // si hay error, muestra un mensaje de error
			}
			mutexArchivo0.Unlock() // desbloquea el acceso a el archivo

		case 1: // prioridad 1= Invertir los dígitos del número y almacenar el resultado en un archivo llamado "prioridad1.txt" en formato de par ordenado (1, resultado)
			resultado := invertirDigitos(tarea.Valor)
			mutexArchivo1.Lock()                                                                  // bloquea el acceso a el archivo para evitar conflictos de escritura
			file, err := os.OpenFile("prioridad1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // abre el archivo en modo append, crea si no existe y con permisos de escritura
			if err == nil {                                                                       // si no hay error al abrir el archivo
				fmt.Fprintf(file, "(1, %d)\n", resultado) // escribe el resultado en el archivo
				file.Close()
			} else {
				fmt.Println("Error escribiendo prioridad1.txt:", err) // si hay error , muestra un mensaje de error
			}
			mutexArchivo1.Unlock() // desbloquea el acceso a el archivo

		case 2: // prioridad 2= Multiplicar el número por una constante X (10) y mostrar el resultado en la consola
			resultado := tarea.Valor * 10                                                // m
			fmt.Printf("Prioridad 2= %d multiplicado es = %d\n", tarea.Valor, resultado) // muestra el resultado

		case 3: // prioridad 3= Acumular el valor de la tarea en una variable compartida y mostrar el acumulado en la consola
			mutexAcumulado.Lock()                                                   // bloquea el acceso a la variable acumulado para evitar conflictos de escritura
			acumulado += tarea.Valor                                                // acumula el valor de la tarea
			fmt.Printf("Prioridad 3 acumulado (+%d): %d\n", tarea.Valor, acumulado) // muestra el acumulado en la consola
			mutexAcumulado.Unlock()                                                 // desbloquea acceso a la variable acumulado
		}
		time.Sleep(200 * time.Millisecond) // simulacion de trabajo
		wg.Done()
	}
}

func sumarDigitos(n int) int { //usada en datos de prioridad 0
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(n int) int { //usada en datos de prioridad 1
	invertido := 0
	for n > 0 {
		invertido = invertido*10 + n%10
		n /= 10
	}
	return invertido
}

func mostrarContenidoArchivo(nombre string) {
	fmt.Printf("\nContenido de %s:\n", nombre)
	data, _ := os.ReadFile(nombre)
	if len(data) == 0 {
		fmt.Println("Archivo vacío")
		return
	}
	fmt.Println(string(data))
}
