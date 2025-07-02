package main

import (
	"fmt"
	"strings"
	"unicode"
)


func cambiar(orig, remp string) string {
	resultado := []rune(remp) // uso runas para descomponer letra por letra
	for i, palab := range orig {
		if i < len(resultado) { //recorre hasta la dim del resultado
			if unicode.IsUpper(palab) {
				// si  está en mayúscula, paso la nueva a mayúscula
				resultado[i] = unicode.ToUpper(resultado[i])
			} else {
				// si está en minúscula, la dejo en minúscula
				resultado[i] = unicode.ToLower(resultado[i])
			}
		}
	}
	return string(resultado)
}

func main() {
	frase := "los jueves   voy en martes miércoles"
	fmt.Println("Frase original:", frase)

	// separo las palabras ignorando los espacios
	palabras := strings.Fields(frase)

	for i, palabra := range palabras {
		// si encuentro jueves reemplazo por martes cumpliendo las minusculas y mayusculas
		if strings.EqualFold(palabra, "jueves") {
			palabras[i] = cambiar(palabra, "martes")
		// si encuentro miercoles reemplazo por automovil
		} else if strings.EqualFold(palabra, "miércoles") {
			palabras[i] = cambiar(palabra, "automóvil")
		}
	}

	// uno las palabras en una sola frase
	nuevaFrase := strings.Join(palabras, " ")
	fmt.Println("Frase modificada:", nuevaFrase)
}
