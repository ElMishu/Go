package main

import "fmt"

type cliente struct {
	nombre   string
	apellido string
	dni      int
	impuesto string
	monto    float64
}

func impuestoMasPagado(conteo map[string]int) string {
	var maxClave string
	maxValor := 0

	for clave, valor := range conteo {
		if valor > maxValor {
			maxValor = valor
			maxClave = clave
		}
	}
	return maxClave
}

func atencionClientes(clientes []cliente) {
	conteo := make(map[string]int)
	conteo["A"]
	conteo["B"]
	conteo["C"]
	conteo["D"]
	var montoAcumulado float64
	i := 0

	for i < len(clientes) && montoAcumulado < 10000 {
		montoAcumulado += clientes[i].monto
		conteo[clientes[i].impuesto]++
		i++
	}
	if montoAcumulado >= 10000 {
		fmt.Println("La cantidad de clientes que quedaron sin antender es ", len(clientes)-i)
	}
	fmt.Println("El impuesto más pagado fue:", impuestoMasPagado(conteo))
	fmt.Printf("Monto recaudado: %.2f\n", montoAcumulado)

}

func main() {

	clientes := []cliente{{nombre: "Juan", apellido: "Perez", dni: 1234, impuesto: 1, monto: 5000.50},
		{nombre: "Ana", apellido: "Garcia", dni: 5678, impuesto: 2, monto: 3000.00},
		{nombre: "Luis", apellido: "Martinez", dni: 9101, impuesto: 1, monto: 7000.75}}
	atencionClientes(clientes)

}
