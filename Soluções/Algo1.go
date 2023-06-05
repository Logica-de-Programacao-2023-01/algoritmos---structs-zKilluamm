package main

import (
	"fmt"
	"math"
)

// Struct Circulo com campo raio
type Circulo struct {
	raio float64
}

// Função para calcular a área do círculo
func calcularAreaCirculo(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	// Exemplo de uso
	c := Circulo{raio: 5}
	area := calcularAreaCirculo(c)
	fmt.Printf("Área do círculo: %.2f\n", area)
}
