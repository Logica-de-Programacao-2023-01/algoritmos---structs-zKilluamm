package main

import (
	"fmt"
	"math"
)

// Struct Filme com campos título, diretor, ano e avaliações
type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

// Função para adicionar uma avaliação ao Filme
func adicionarAvaliacao(filme *Filme, avaliacao int) {
	filme.avaliacoes = append(filme.avaliacoes, avaliacao)
}

// Função para remover uma avaliação do Filme
func removerAvaliacao(filme *Filme, indice int) {
	if indice >= 0 && indice < len(filme.avaliacoes) {
		filme.avaliacoes = append(filme.avaliacoes[:indice], filme.avaliacoes[indice+1:]...)
	}
}

// Função para calcular a média das avaliações do Filme
func calcularMediaAvaliacoes(filme Filme) float64 {
	total := 0
	for _, avaliacao := range filme.avaliacoes {
		total += avaliacao
	}
	media := float64(total) / float64(len(filme.avaliacoes))
	return math.Round(media*10) / 10 // Arredonda para uma casa decimal
}

// Função para imprimir as informações do Filme e sua média de avaliações
func imprimirFilme(filme Filme) {
	fmt.Printf("Título: %s\n", filme.titulo)
	fmt.Printf("Diretor: %s\n", filme.diretor)
	fmt.Printf("Ano: %d\n", filme.ano)
	fmt.Printf("Média das avaliações: %.1f\n", calcularMediaAvaliacoes(filme))
}

func main() {
	// Exemplo de uso
	filme := Filme{
		titulo:     "Inception",
		diretor:    "Christopher Nolan",
		ano:        2010,
		avaliacoes: []int{9, 8, 10, 7, 9},
	}

	imprimirFilme(filme)

	adicionarAvaliacao(&filme, 8)
	fmt.Println("Após adicionar uma avaliação:")
	imprimirFilme(filme)

	removerAvaliacao(&filme, 2)
	fmt.Println("Após remover a terceira avaliação:")
	imprimirFilme(filme)
}
