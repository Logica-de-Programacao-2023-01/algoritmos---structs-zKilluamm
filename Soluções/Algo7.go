package main

import "fmt"

// Struct Animal com campos nome, espécie, idade e som
type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

// Função para modificar o som que o animal faz
func modificarSom(a *Animal, novoSom string) {
	a.som = novoSom
}

// Função para imprimir as informações do animal e o som que ele faz
func imprimirAnimal(a Animal) {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Espécie: %s\n", a.especie)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Som: %s\n", a.som)
}

func main() {
	// Exemplo de uso
	animal := Animal{
		nome:    "Simba",
		especie: "Leão",
		idade:   3,
		som:     "Rugido",
	}

	imprimirAnimal(animal)

	modificarSom(&animal, "Miado")
	fmt.Println("Após modificar o som:")
	imprimirAnimal(animal)
}
