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

// Função para buscar playlists por título de música
func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	// Exemplo de uso
	playlist1 := Playlist{
		nome: "Playlist 1",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 4},
			{titulo: "Música 2", artista: "Artista 2", duracao: 3},
		},
	}

	playlist2 := Playlist{
		nome: "Playlist 2",
		musicas: []Musica{
			{titulo: "Música 3", artista: "Artista 3", duracao: 5},
			{titulo: "Música 4", artista: "Artista 4", duracao: 2},
		},
	}

	playlist3 := Playlist{
		nome: "Playlist 3",
		musicas: []Musica{
			{titulo: "Música 5", artista: "Artista 5", duracao: 3},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	tituloBuscado := "Música 3"
	resultado := buscarPlaylistsPorTitulo(tituloBuscado, playlists)

	fmt.Printf("Playlists com a música \"%s\":\n", tituloBuscado)
	for _, playlist := range resultado {
		fmt.Println(playlist.nome)
	}
}
