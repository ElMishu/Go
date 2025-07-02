package main
// VER STRINGS.FIELD 
//split
import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func invertir (palabra string) string{
	aux := [] rune (palabra) //runa para descomponer la palabra
	aux2 := []rune{}; // runa para guardar al reves 
	for j:=len(aux)-1;j>=0; j--{ 
		aux2= append(aux2, aux[j]) //invierto 
	}
	return string(aux2) //retorno runa aux2 convertido en un string
}


func main() {
	var cadena string
	fmt.Println("Leer una frase: ")
	// leo la frase completa desde teclado (hasta que apretás enter)
	reader := bufio.NewReader(os.Stdin)
	cadena, _ = reader.ReadString('\n') //subestima el error
	
//cadena:= "hola soy juana y ando en patineta"
	rcadenas := strings.Fields(cadena) // separo el string palabra porpalabra
	contador:=0 //para saber pos
	for i, palabra:=  range rcadenas {
		contador++; //
		if !(contador%2 == 0) { //me fijo si esta en una pos par o impar
			rcadenas[i] = invertir (palabra) //si esta en pos impar iinvierto todo
		}
		
	}
	// uno todas las palabras otra vez con espacios
	nueva:= strings.Join(rcadenas, " ")
	fmt.Println("la nueva cadena es" , nueva)
}

