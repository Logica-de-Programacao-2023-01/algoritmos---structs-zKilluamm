package main

import "fmt"

// Struct Endereco com campos rua, número, cidade e estado
type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

// Struct Pessoa com campos nome, idade e endereço
type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

// Função para imprimir o endereço completo da pessoa
func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço de %s:\n", p.nome)
	fmt.Printf("Rua: %s, Número: %d\n", p.endereco.rua, p.endereco.numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", p.endereco.cidade, p.endereco.estado)
}

func main() {
	// Exemplo de uso
	endereco := Endereco{
		rua:    "Rua das Flores",
		numero: 123,
		cidade: "São Paulo",
		estado: "SP",
	}

	pessoa := Pessoa{
		nome:     "João",
		idade:    30,
		endereco: endereco,
	}

	imprimirEnderecoCompleto(pessoa)
}
