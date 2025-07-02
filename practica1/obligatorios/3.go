package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
)

// Separa los signos adelante o atrás de una palabra
func limpiarPalabra(palabra string) (letras string, signos string, atras bool) {
	runes := []rune(palabra)

	if len(runes) == 0 {
		return "", "", true
	}

	if !unicode.IsLetter(runes[len(runes)-1]) {
		i := len(runes) - 1
		for i >= 0 && !unicode.IsLetter(runes[i]) {
			i--
		}
		return string(runes[:i+1]), string(runes[i+1:]), true
	}

	if !unicode.IsLetter(runes[0]) {
		i := 0
		for i < len(runes) && !unicode.IsLetter(runes[i]) {
			i++
		}
		return string(runes[i:]), string(runes[:i]), false
	}

	return palabra, "", true
}

// Invierte las mayúsculas y minúsculas
func modificarMayusculasYMinusculas(palabra string) string {
	aux := []rune(palabra)
	for i := 0; i < len(aux); i++ {
		if unicode.IsLower(aux[i]) {
			aux[i] = unicode.ToUpper(aux[i])
		} else if unicode.IsUpper(aux[i]) {
			aux[i] = unicode.ToLower(aux[i])
		}
	}
	return string(aux)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Tenes que pasar una palabra como argumento.")
		fmt.Println("Ejemplo: go run programa.go pequeño")
		return
	}

	word := os.Args[1]
	fmt.Println("Palabra a buscar/modificar:", word)

	fmt.Println("Escribí una frase:")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')

	partes := strings.Fields(frase)

	for i, palabra := range partes {
		palabraLimpia, signos, estaAlFinal := limpiarPalabra(palabra)

		// Paso ambas palabras a minúsculas y comparo si la palabra buscada está dentro de la otra
		if strings.Contains(strings.ToLower(palabraLimpia), strings.ToLower(word)) {
			palabraLimpia = modificarMayusculasYMinusculas(palabraLimpia)

			if estaAlFinal {
				partes[i] = palabraLimpia + signos
			} else {
				partes[i] = signos + palabraLimpia
			}
		}
	}

	nueva := strings.Join(partes, " ")
	fmt.Println("Frase modificada:")
	fmt.Println(nueva)
}
