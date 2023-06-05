package main

import "fmt"

// Struct Viagem com campos origem, destino, data e preço
type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

// Função para encontrar a viagem mais cara em um slice de Viagens
func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		// Retorna uma viagem vazia se não houver viagens no slice
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	// Exemplo de uso
	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: "2023-06-01", preco: 200.50},
		{origem: "Rio de Janeiro", destino: "Salvador", data: "2023-07-15", preco: 350.00},
		{origem: "Salvador", destino: "Fortaleza", data: "2023-08-10", preco: 400.75},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.origem)
	fmt.Println("Destino:", viagemMaisCara.destino)
	fmt.Println("Data:", viagemMaisCara.data)
	fmt.Println("Preço:", viagemMaisCara.preco)
}
