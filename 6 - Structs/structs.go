package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario

	u.nome = "Vyctor"
	u.idade = 27
	fmt.Println(u)

	enderecoUsuario := endereco{
		"Rua dos Bobos", 0,
	}

	usuario2 := usuario{"Vyctor", 27, enderecoUsuario}

	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Vyctor"}

	fmt.Println(usuario3)
}
