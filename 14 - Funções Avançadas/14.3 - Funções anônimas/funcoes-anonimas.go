package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf(texto)
	}("Olá mundo")

	fmt.Println(retorno)
}
