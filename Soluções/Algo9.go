package main

import (
	"fmt"
	"math"
)

// Struct Aluno com campos nome, idade e notas
type Aluno struct {
	nome  string
	idade int
	notas []float64
}

// Função para adicionar uma nota ao Aluno
func adicionarNota(aluno *Aluno, nota float64) {
	aluno.notas = append(aluno.notas, nota)
}

// Função para remover uma nota do Aluno
func removerNota(aluno *Aluno, indice int) {
	if indice >= 0 && indice < len(aluno.notas) {
		aluno.notas = append(aluno.notas[:indice], aluno.notas[indice+1:]...)
	}
}

// Função para calcular a média das notas do Aluno
func calcularMedia(aluno Aluno) float64 {
	total := 0.0
	for _, nota := range aluno.notas {
		total += nota
	}
	media := total / float64(len(aluno.notas))
	return math.Round(media*100) / 100 // Arredonda para duas casas decimais
}

// Função para imprimir o nome, idade e média do Aluno
func imprimirAluno(aluno Aluno) {
	fmt.Printf("Nome: %s\n", aluno.nome)
	fmt.Printf("Idade: %d\n", aluno.idade)
	fmt.Printf("Média das notas: %.2f\n", calcularMedia(aluno))
}

func main() {
	// Exemplo de uso
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{7.5, 8.0, 9.5},
	}

	imprimirAluno(aluno)

	adicionarNota(&aluno, 8.5)
	fmt.Println("Após adicionar uma nota:")
	imprimirAluno(aluno)

	removerNota(&aluno, 1)
	fmt.Println("Após remover a segunda nota:")
	imprimirAluno(aluno)
}
