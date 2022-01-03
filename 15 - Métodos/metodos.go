package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Vyctor", 27}
	usuario2 := usuario{"Davi", 30}

	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorDeIdade()
	usuario2.salvar()
	usuario2.maiorDeIdade()
}
