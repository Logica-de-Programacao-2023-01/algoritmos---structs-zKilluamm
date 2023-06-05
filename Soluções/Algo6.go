package main

import (
	"fmt"
	"time"
)

// Struct Funcionario com campos nome, salario e idade
type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

// Função para aumentar o salário do funcionário em uma determinada porcentagem
func aumentarSalario(f *Funcionario, porcentagem float64) {
	f.salario += f.salario * (porcentagem / 100)
}

// Função para diminuir o salário do funcionário em uma determinada porcentagem
func diminuirSalario(f *Funcionario, porcentagem float64) {
	f.salario -= f.salario * (porcentagem / 100)
}

// Função para calcular o tempo de serviço do funcionário
func calcularTempoServico(f Funcionario) int {
	idadeInicioTrabalho := 18
	anosTrabalho := time.Now().Year() - (f.idade + idadeInicioTrabalho)
	return anosTrabalho
}

func main() {
	// Exemplo de uso
	funcionario := Funcionario{
		nome:    "João",
		salario: 5000,
		idade:   30,
	}

	fmt.Println("Salário inicial:", funcionario.salario)

	aumentarSalario(&funcionario, 10)
	fmt.Println("Salário após aumento:", funcionario.salario)

	diminuirSalario(&funcionario, 5)
	fmt.Println("Salário após diminuição:", funcionario.salario)

	tempoServico := calcularTempoServico(funcionario)
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
