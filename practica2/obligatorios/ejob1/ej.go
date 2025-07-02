package ejob1

import "fmt"

func informarbariloche(lista List) {
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		if actual.data.ciudad == "Bariloche" {
			fmt.Println(actual.data.nombre, "", actual.data.apellido)
		}
	}
}

func calcularanio(lista List) int {
	contador := make(map[int]int)
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		contador[actual.data.fecha.anio]++
	}
	max := 0
	anio := 0
	for k, v := range contador {
		if v > max {
			max = v
			anio = k
		}
	}
	return anio
}

func calcularcarrera(lista List) string {
	contador := make(map[int]int)
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		contador[actual.data.codigo]++
	}
	max := 0
	codigo := 0
	for k, v := range contador {
		if v > max {
			max = v
			codigo = k
		}
	}
	switch codigo {
	case 1:
		return "APU"
	case 2:
		return "LI"
	case 3:
		return "LS"
	default:
		return "No se encontró la carrera"
	}
}

func EliminarSinTitulo(lista *List) {
	// Eliminar nodos desde el principio mientras el título sea falso
	for lista.head != nil && !lista.head.data.titulo {
		lista.head = lista.head.siguiente
	}

	// Si la lista quedó vacía, no hay nada más que hacer
	if lista.head == nil {
		return
	}

	// Recorrer la lista a partir del primer nodo con título
	actual := lista.head
	for actual != nil && actual.siguiente != nil {
		if !actual.siguiente.data.titulo {
			// Saltar el nodo sin título
			actual.siguiente = actual.siguiente.siguiente
		} else {
			actual = actual.siguiente
		}
	}
}

func main() {
	lista := New()
	// ya cargada
	informarbariloche(lista)
	fmt.Println("El año que mas ingresantes nacieron fue", calcularanio(lista))
	fmt.Println("La carrera con mas ingresantes inscriptos fue", calcularcarrera(lista))
	lista.Iterate()
	EliminarSinTitulo(&lista)
	fmt.Println("Lista nueva solo con los ingresantes que tienen titulo")
	fmt.Println(ToString(lista))
}
