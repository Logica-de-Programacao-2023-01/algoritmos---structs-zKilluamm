package main

import "fmt"

// Struct Triângulo com campos base e altura
type Triangulo struct {
	base   float64
	altura float64
}

// Função para calcular a área do triângulo
func calcularAreaTriangulo(t Triangulo) float64 {
	return (t.base * t.altura) / 2
}

func main() {
	// Exemplo de uso
	triangulo := Triangulo{
		base:   4.5,
		altura: 3.2,
	}

	area := calcularAreaTriangulo(triangulo)
	fmt.Printf("Área do triângulo: %.2f\n", area)
}
