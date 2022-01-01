package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Vyctor", "Guimaraes", 27, 178}

	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}

	fmt.Println(e1.nome, e1.sobrenome)
}
