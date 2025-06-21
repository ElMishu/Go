package main

import "fmt"

type Fecha struct {
	dia  int
	mes  int
	anio int
}

type Ingresante struct {
	apellido string
	nombre   string
	ciudad   string
	fecha    Fecha
	titulo   bool
	carrera  int
}

func (i Ingresante) String() string {
	return fmt.Sprintf("Apellido: %s, Nombre: %s, Ciudad: %s, Fecha: %02d/%02d/%04d, Titulo: %t, Carrera: %d",
		i.apellido, i.nombre, i.ciudad, i.fecha.dia, i.fecha.mes, i.fecha.anio, i.titulo, i.carrera)
}

func main() {
	f := Fecha{21, 5, 2003}
	i := Ingresante{
		apellido: "Tranquillini",
		nombre:   "Nahuel",
		ciudad:   "La Plata",
		fecha:    f,
		titulo:   true,
		carrera:  42, // por ejemplo
	}
	fmt.Println(i)
}
