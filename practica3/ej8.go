package main

import (
	"fmt"
	"sync"
)

var mu sync.RWMutex
var a Agenda

type Contact struct {
	nombre   string
	apellido string
	correo   string
	telefono string
}

type Agenda struct {
	Contact map[string]Contact
}

func (a *Agenda) AgregarContacto(c Contact) {
	mu.Lock()
	a.Contact[c.correo] = c
	mu.Unlock()
}

func (a *Agenda) EliminarContacto(c Contact) {
	mu.Lock()
	_, ok := a.Contact[c.correo]
	if !ok {
		fmt.Println(c.nombre, "no está en la agenda.")
	} else {
		delete(a.Contact, c.correo)
		fmt.Println("Contacto eliminado:", c.nombre)
	}
	mu.Unlock()
}

func (a *Agenda) BuscarContacto(correo string) {
	mu.RLock()
	valor, existe := a.Contact[correo]
	if existe {
		fmt.Println("Encontrado:", valor)
	} else {
		fmt.Println("No existe esa clave.")
	}
	mu.RUnlock()
}

func main() {
	a = Agenda{Contact: make(map[string]Contact)}

	juan := Contact{"Juan", "Pérez", "juan@gmail.com", "1234"}
	ana := Contact{"Ana", "López", "ana@gmail.com", "5678"}

	a.AgregarContacto(juan)
	a.AgregarContacto(ana)

	a.BuscarContacto("juan@gmail.com")
	a.EliminarContacto(juan)
	a.BuscarContacto("juan@gmail.com")
}
