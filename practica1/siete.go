/*
7. Las temperaturas de los pacientes de un hospital se dividen en 3
grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja
(menor de 36). Se deben leer 10 temperaturas de pacientes e
informar el porcentaje de pacientes de cada grupo. Luego se
debe imprimir el promedio entero entre la temperatura máxima y
la temperatura mínima.
a. ¿Se puede utilizar el case para tipos reales en otros
lenguajes?
b. ¿ Cómo se realizan las conversiones entre reales (punto
flotante) y enteros en otros lenguajes ?
Sub-objetivo: El tipado fuerte, usar casting. Operaciones y
E/S con float. Casting en otros lenguajes.
*/
package main

import "fmt"

func sumargrupo(temperatura float32, normales, bajas, altas *int) {
	switch {
	case temperatura > 37.5:
		*altas++
	case temperatura >= 36 && temperatura <= 37.5:
		*normales++
	case temperatura < 36:
		*bajas++

	}
}

func calcularmax(temperatura float32, mayor *float32) {
	if temperatura > *mayor {
		*mayor = temperatura
	}
}

func calcularmin(temperatura float32, minimo *float32) {
	if temperatura < *minimo {
		*minimo = temperatura
	}
}

func leertemperaturas(mayor, menor *float32) {
	var temperatura float32
	var normales, bajas, altas int
	for i := 0; i < 10; i++ {
		fmt.Scan(&temperatura)
		sumargrupo(temperatura, &normales, &bajas, &altas)
		calcularmax(temperatura, mayor)
		calcularmin(temperatura, menor)
	}

	fmt.Println("Porcentaje de pacientes con temperatura baja:", float32(bajas)*100/10, "%")
	fmt.Println("Porcentaje de pacientes con temperatura normal:", float32(normales)*100/10, "%")
	fmt.Println("Porcentaje de pacientes con temperatura alta:", float32(altas)*100/10, "%")
}

func main() {
	var mayor float32 = -9999
	var menor float32 = 9999
	leertemperaturas(&mayor, &menor)
	promedio := int((mayor + menor) / 2)
	fmt.Printf("Promedio entero entre máx y mín: %d\n", promedio)

}
