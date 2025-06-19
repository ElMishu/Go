/*6. Escriba un programa que lee desde la entrada estándar dos
enteros y retorne la división entre el mayor de ellos y el menor.
Realizar el mismo programa considerando que se leen dos
enteros sin signo. Luego modifique para que trabaje con reales
(punto flotante). Ver que sucede con las división por cero.
*/

package main

import "fmt"

func division(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		fmt.Println("No se puede dividir por 0")
		return 0
	} else {
		if num1 > num2 {
			return (num1 / num2)
		} else {
			return num2 / num1
		}
	}
}

func main() {
	var num1, num2 int
	fmt.Println("Ingrese dos numeros")
	fmt.Scan(&num1, &num2)
	fmt.Println("La divisioon entre el mayor de ellos y el menor es=", division(num1, num2))
}

/*

func divisionsinsigno(num1, num2 uint) uint {
	if num1 == 0 || num2 == 0 {
		fmt.Println("No se puede dividir por 0")
		return 0
	} else {
		if num1 > num2 {
			return (num1 / num2)
		} else {
			return num2 / num1
		}
	}
}

func main() {
	var num1, num2 uint
	fmt.Println("Ingrese dos numeros sin signo")
	fmt.Scan(&num1, &num2)
	fmt.Println("La divisioon entre el mayor de ellos y el menor es=", division(num1, num2))
}








func division(num1, num2 float32) float32 {
	if num1 == 0 || num2 == 0 {
		fmt.Println("No se puede dividir por 0")
		return 0
	} else {
		if num1 > num2 {
			return (num1 / num2)
		} else {
			return num2 / num1
		}
	}
}

func main() {
	var num1, num2 float32
	fmt.Println("Ingrese dos numeros reales")
	fmt.Scan(&num1, &num2)
	fmt.Println("La divisioon entre el mayor de ellos y el menor es=", division(num1, num2))
}
*/
