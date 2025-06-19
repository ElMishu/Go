package main

import "fmt"

type Contact struct{
	nombre string
	apellido string
	correo string
	telefono string
}

type Agenda struct {
	Contact map [string] Contact
}

func (a*Agenda) AgregarContacto(c Contacto){
	
}

func main(){
	miAgenda=
}