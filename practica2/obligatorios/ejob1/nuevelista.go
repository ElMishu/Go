package ejob1 //ejercico nueve
import "fmt"

type Nodo struct {
	data      Estudiante
	siguiente *Nodo
}

type List struct {
	head *Nodo
}

// New crea y devuelve una lista vacía.
func New() List {
	return List{head: nil}
}

// isEmpty verifica si la lista está vacía.
// Devuelve true si la lista está vacía, false en caso contrario.
func IsEmpty(l *List) bool {
	return l.head == nil
}

// agregar agrega un nuevo nodo al final de la lista.
func (l *List) Agregar(valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(l) {
		l.head = nuevoNodo
	} else {
		actual := l.head
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
	}
}

func FrontElement(self List) Estudiante {
	if self.head == nil {
		fmt.Println("lista vacía") // error si la lista no tiene nodos
	}
	return self.head.data
}

func Len(self List) int {
	if self.head == nil {
		return 0
	} else {
		cant := 0
		actual := self.head
		for actual != nil {
			cant++
			actual = actual.siguiente
		}
		return cant
	}
}

func ToString(lista List) string {
	if lista.head == nil {
		return "La lista está vacía"
	}
	actual := lista.head
	var resultado string
	for actual != nil {
		resultado += fmt.Sprintf("%v ", actual.data)
		actual = actual.siguiente
	}
	return resultado

}

func PushFront(self *List, valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(self) {
		self.head = nuevoNodo
	} else {
		nuevoNodo.siguiente = self.head
		self.head = nuevoNodo
	}
}

func PushBack(self *List, valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(self) {
		self.head = nuevoNodo
	} else {
		actual := self.head
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
	}
}

func (self List) Iterate() {
	if IsEmpty(&self) {
		fmt.Println("La lista está vacía")
		return
	}
	actual := self.head
	for actual != nil {
		fmt.Println(actual.data)
		actual = actual.siguiente
	}
}
