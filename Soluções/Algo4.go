package main

import "fmt"

// Struct Musica com campos título, artista e duração
type Musica struct {
	titulo  string
	artista string
	duracao int
}

// Struct Playlist com campos nome e músicas
type Playlist struct {
	nome    string
	musicas []Musica
}

// Função para imprimir informações da playlist
func imprimirPlaylist(p Playlist) {
	fmt.Printf("Playlist: %s\n", p.nome)

	tempoTotal := 0
	for _, musica := range p.musicas {
		fmt.Printf("Música: %s\n", musica.titulo)
		fmt.Printf("Artista: %s\n", musica.artista)
		fmt.Printf("Duração: %d minutos\n", musica.duracao)

		tempoTotal += musica.duracao
	}

	fmt.Printf("Tempo total da playlist: %d minutos\n", tempoTotal)
}

func main() {
	// Exemplo de uso
	musicas := []Musica{
		{titulo: "Música 1", artista: "Artista 1", duracao: 4},
		{titulo: "Música 2", artista: "Artista 2", duracao: 3},
		{titulo: "Música 3", artista: "Artista 3", duracao: 5},
	}

	playlist := Playlist{
		nome:    "Minha Playlist",
		musicas: musicas,
	}

	imprimirPlaylist(playlist)
}
