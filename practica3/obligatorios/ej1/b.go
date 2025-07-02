/*b) Realice el programa utilizando más de una goroutine para dividir
el trabajo de comprobación de primos entre múltiples goroutines
en paralelo
i) Cada goroutine debe recibir un rango de números a
comprobar y devolver una lista de los números primos
encontrados
ii) El programa principal debe recibir el número N y dividir el
trabajo en goroutines, asignando a cada una un rango de
números a comprobar
iii) Una vez que todas las goroutines hayan finalizado, el
programa principal debe recopilar los resultados y mostrar
en pantalla la lista de números primos encontrados.

c) Realice la ejecución con N igual a 1.000, 100.000 y 1.000.000
tanto del programa a) como del b). Para cada caso calcule el
speed-up con la siguiente fórmula:
S(p) = T(1) / T(p)
donde, S(p) es el speed-up, T(1) es el tiempo que tarda la
ejecución con una única goroutine y T(p) es el tiempo de
ejecución con p goroutines.*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func esPrimo(n int) bool { // función que verifica si un número es primo
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ { // verifica si n es divisible por algún número entre 2 y la raíz cuadrada de n
		if n%i == 0 { // si n es divisible por i, no es primo
			return false
		}
	}
	return true // si no es divisible por ningún número, es primo
}

func buscarPrimos(desde, hasta int, resultados chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done() // función que busca números primos en un rango y envía los resultados a un canal
	var primos []int
	for i := desde; i <= hasta; i++ {
		if esPrimo(i) {
			primos = append(primos, i) // si el número es primo, lo agrega a la lista de primos
		}
	}
	resultados <- primos // envía la lista de primos encontrados al canal
}

func main() {
	start := time.Now() // inicia el cronómetro para medir el tiempo de ejecución
	var input string
	fmt.Print("Ingresá un número entero positivo: ")
	fmt.Scanln(&input)

	N, err := strconv.Atoi(input) // convierte lo que se leyo de string a int
	if err != nil || N < 1 {      // si hay un error en la conversión o el número es menor que 1, muestra un mensaje de error
		fmt.Println("Entrada inválida. Debe ser un entero positivo.")
		return
	}

	const goroutines = 4 // número de goroutines a utilizar
	tamBloque := N / goroutines
	var wg sync.WaitGroup
	resultados := make(chan []int, goroutines)

	for i := 0; i < goroutines; i++ { // divide el trabajo en goroutines
		desde := 2 + i*tamBloque       // el primer número primo es 2, así que empezamos desde ahí
		hasta := desde + tamBloque - 1 // calcula el rango de números a comprobar para cada goroutine
		if i == goroutines-1 {         // si la última goroutine llega hasta N
			hasta = N
		}
		wg.Add(1)                                      // incrementa el contador de goroutines
		go buscarPrimos(desde, hasta, resultados, &wg) // inicia la goroutine para buscar primos en el rango
	}

	wg.Wait()         // espera a que todas las goroutines terminen
	close(resultados) // cierra el canal de resultados

	var todos []int                 // slice para almacenar todos los números primos encontrados
	for lista := range resultados { // recibe los resultados del canal
		todos = append(todos, lista...) // agrega los números primos encontrados a la lista total
	}

	fmt.Printf("numeros primos menores o iguales a %d:\n", N) // muestra los números primos encontrados
	for _, p := range todos {
		fmt.Print(p, " ")
	}
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Println("Tiempo total:", elapsed)
}
