package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Vyctor",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Vyctor",
			"ultimo":   "Guimarães",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2["nome"])
}
