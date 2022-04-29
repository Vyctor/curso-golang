package main

import "fmt"

func generica(interf interface{}) {
	fmt.Print(interf)
}

func main() {
	generica("Oi")
	generica(10)
	generica(true)
}
