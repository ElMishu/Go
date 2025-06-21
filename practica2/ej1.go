package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const totalPacientes = 10
	var temperaturas [totalPacientes]float64
	var alta, normal, baja int
	var maxTemp, minTemp float64

	// Leer temperaturas desde la entrada estándar
	for i := 0; i < totalPacientes; i++ {
		var input string
		fmt.Fscan(os.Stdin, &input)
		temp, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error al leer la temperatura:", err)
			return
		}
		temperaturas[i] = temp
	}

	// Inicializar valores de maxTemp y minTemp
	maxTemp = temperaturas[0]
	minTemp = temperaturas[0]

	// Clasificar temperaturas y calcular máximos y mínimos
	for _, temp := range temperaturas {
		if temp > 37.5 {
			alta++
		} else if temp >= 36 && temp <= 37.5 {
			normal++
		} else {
			baja++
		}

		if temp > maxTemp {
			maxTemp = temp
		}
		if temp < minTemp {
			minTemp = temp
		}
	}

	// Calcular porcentajes
	porcentajeAlta := float64(alta) * 100 / totalPacientes
	porcentajeNormal := float64(normal) * 100 / totalPacientes
	porcentajeBaja := float64(baja) * 100 / totalPacientes

	// Calcular promedio entero entre temperatura máxima y mínima
	promedio := int((maxTemp + minTemp) / 2)

	// Imprimir resultados
	fmt.Printf("Porcentaje de pacientes con temperatura alta: %.2f%%\n", porcentajeAlta)
	fmt.Printf("Porcentaje de pacientes con temperatura normal: %.2f%%\n", porcentajeNormal)
	fmt.Printf("Porcentaje de pacientes con temperatura baja: %.2f%%\n", porcentajeBaja)
	fmt.Printf("Promedio entero entre la temperatura máxima y mínima: %d\n", promedio)
}
