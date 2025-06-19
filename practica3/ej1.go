package main

import "fmt"

type Map[K comparable, V any] map[K]V

func main() {
	m1 := Map[string, int]{"nahuel": 21, "lucero": 22, "juan": 23} // Clave nombre valor edad
	fmt.Println(m1)                                                // imprimo el mapa

	for nombre, edad := range m1 { // otra forma de recorrer un mapa
		fmt.Printf("%s: %d años\n", nombre, edad) // imprimo
	}

	m2 := Map[string, []string]{ // Clave MATERIA, valor ALUMNO
		"FOD":    {"Nahuel", "Lucía", "Pedro"},
		"Goland": {"Marta", "Juan"},
		"DBD":    {"Carlos", "Lucía"},
	}

	fmt.Println("Materias con más de 2 alumnos:")
	for materia, alumnos := range m2 {
		if len(alumnos) > 2 {
			fmt.Printf("→ %s (%d alumnos)\n", materia, len(alumnos))
		}
	}

}
