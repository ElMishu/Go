package main

import (
	"errors"
	"fmt"
)

type Elemento struct {
	numero   int
	cantidad int
}

type OptimumSlice struct {
	Info []Elemento
}

func New(slice []int) OptimumSlice {
	if len(slice) == 0 {
		return OptimumSlice{} // Retorna un OptimumSlice vacía si el slice de entrada está vacío
	}

	resultado := OptimumSlice{} // Inicializa un OptimumSlice vacío
	nroAct := slice[0]          // Toma el primer número como el número actual
	cont := 1                   // Inicializa el contador en 1

	for i := 1; i < len(slice); i++ { // Recorre el slice desde el segundo elemento
		if slice[i] == nroAct { // Si el número actual es igual al número anterior
			cont++ // Incrementa el contador
		} else {
			// Agrega el número actual con su cantidad
			resultado.Info = append(resultado.Info, Elemento{numero: nroAct, cantidad: cont})
			// Actualiza con el nuevo número
			nroAct = slice[i]
			cont = 1
		}
	}

	// Agrega el último grupo que quedó sin guardar
	resultado.Info = append(resultado.Info, Elemento{numero: nroAct, cantidad: cont})

	return resultado // Retorna el OptimumSlice con los elementos agrupados
}

// inciso ii
func IsEmpty(o OptimumSlice) bool {
	return len(o.Info) == 0 // Verifica si el OptimumSlice está vacío devolviendo true si no tiene elementos
}

// inciso iii
func Len(o OptimumSlice) int {
	var elementos int = 0              // Inicializa un contador de elementos a 0
	for i := 0; i < len(o.Info); i++ { // Recorre cada elemento en Info
		elementos += o.Info[i].cantidad // Suma la cantidad de cada elemento al contador
	}
	return elementos // Retorna el total de elementos contando las cantidades
}

// inciso iv
func FrontElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("el OptimumSlice está vacío") // Retorna un error si el OptimumSlice está vacío
	}
	return o.Info[0].numero, nil // retorna el número del primer elemento del OptimumSlice
}

// inciso v
func LastElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("el OptimumSlice está vacío") // Retorna un error si el OptimumSlice está vacío
	}

	return o.Info[len(o.Info)-1].numero, nil // Retorna el número del último elemento del OptimumSlice
}

// inciso vi
func Insert(os *OptimumSlice, elem int, pos int) (bool, error) {
	opt := (*os)
	if pos < 0 || pos > Len(opt) {
		return false, errors.New("No se pudo insertar") // Retorna un error si la posición es inválida
	}

	if pos == Len(opt) { // Si la posición es igual al largo del OptimumSlice, se agrega al final directamente
		if len(opt.Info) > 0 && opt.Info[len(opt.Info)-1].numero == elem {
			opt.Info[len(opt.Info)-1].cantidad++
		} else {
			opt.Info = append(opt.Info, Elemento{elem, 1})
		}
		*os = opt
		return true, nil
	}

	var ocurrencias int = 0 // Inicializa un contador de ocurrencias a 0
	var valorActual int     // Inicializa una variable para almacenar el valor actual del elemento
	for i := 0; i < len(opt.Info); i++ {
		valorActual = opt.Info[i].numero // Obtiene el número del elemento actual
		for j := 0; j < opt.Info[i].cantidad; j++ {
			if pos == ocurrencias { // Si la posición actual coincide con la posición donde se quiere insertar
				if valorActual == elem { // Si el elemento ya existe en la posición indicada
					opt.Info[i].cantidad++ // Incrementa la cantidad del elemento existente
				} else if (i-1) >= 0 && opt.Info[i-1].numero == elem { // Si el elemento anterior es igual al elemento a insertar
					opt.Info[i-1].cantidad++ // Incrementa la cantidad del elemento anterior
				} else if (i+1) < len(opt.Info) && opt.Info[i+1].numero == elem { // Si el elemento siguiente es igual al elemento a insertar
					opt.Info[i+1].cantidad++ // Incrementa la cantidad del elemento siguiente
				} else {
					*os = insert(opt, elem, pos) // Llama a la función insert para insertar el elemento en la posición indicada y lo devuelve
				}
				return true, nil // Retorna 0 si se pudo insertar el valor en la posición indicada
			}
			ocurrencias++
		}
	}
	return false, errors.New("No se pudo insertar") // Retorna false si no se pudo insertar el elem en la posición indicada
}

func insert(o OptimumSlice, v int, p int) OptimumSlice {
	var s OptimumSlice
	ocurrencias := 0

	for i := 0; i < len(o.Info); i++ {
		posInicio := ocurrencias           // Posición de inicio del bloque actual
		posFin := ocurrencias + o.Info[i].cantidad - 1    // Posición de fin del bloque actual

		if p >= posInicio && p <= posFin {
			// Dividir el bloque actual en dos partes si hace falta
			before := p - posInicio
			after := o.Info[i].cantidad - before

			if before > 0 {      // Si hay elementos antes de la posición p
				s.Info = append(s.Info, Elemento{o.Info[i].numero, before})      
			}
			s.Info = append(s.Info, Elemento{v, 1})
			if after > 0 {    // Si hay elementos después de la posición p
				s.Info = append(s.Info, Elemento{o.Info[i].numero, after})   // Agrega el resto del bloque actual
			}
			// Agregar el resto del slice
			s.Info = append(s.Info, o.Info[i+1:]...)   //
			return s
		}

		s.Info = append(s.Info, o.Info[i])
		ocurrencias += o.Info[i].cantidad
	}

	// Si llega acá es porque se debe insertar al final (caso p == Len)
	s.Info = append(s.Info, Elemento{v, 1})
	return s
}

// inciso vii
func SliceArray(o OptimumSlice) []int {
	var slice []int            // Inicializa un slice vacío para almacenar los números
	for _, r := range o.Info { // Recorre cada elemento en Info
		for i := 0; i < r.cantidad; i++ { // Por cada cantidad del elemento, agrega el número al slice
			slice = append(slice, r.numero) // Agrega el número al slice
		}
	}
	return slice // retorna el slice con todos los numeros de los elementos
}

func main() {
	slice := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}
	opt := New(slice)                            // variable opt con slice cargado a testear
	fmt.Println("OptimumSlice:", opt)            //imprimo para verificar que deberia devolver (1 [2], 2 [3], 3 [2], 4 [1])
	fmt.Println("Testeo IsEmpty:", IsEmpty(opt)) // imprime falso ya que no esta vacio
	fmt.Println("Testeo Length:", Len(opt))      // imprime el ancho, 8
	front, _ := FrontElement(opt)
	fmt.Println("Testeo Front Element:", front) //imprime el primer elemento, 1
	last, _ := LastElement(opt)
	fmt.Println("Testeo Last Element:", last) //imprime el ultimo elemento, 4
	inserted, err := Insert(&opt, 9, 6)
	if err != nil {
		fmt.Println(err) // si no se puede insertar imprime el error
	} else {
		fmt.Println("Insertado:", inserted) // imprime "insertado" si el elemento se inserta correctamente
	}

	inserted2, err2 := Insert(&opt, 3, 32) // prueba de insertado falso
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Insertado:", inserted2)
	}
	fmt.Println("OptimumSlice:", opt)
	fmt.Println("Slice Array:", SliceArray(opt)) // imprime el resultado final
}
