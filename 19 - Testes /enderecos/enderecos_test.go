package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnte(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{
			enderecoInserido: "Rua ABC 123",
			retornoEsperado:  "Rua",
		},
		{
			enderecoInserido: "Avenida Paulista",
			retornoEsperado:  "Avenida",
		},
		{
			enderecoInserido: "Rodovia dos imigrantes Paulista",
			retornoEsperado:  "Rodovia",
		},
		{
			enderecoInserido: "Praça das Rosas",
			retornoEsperado:  "Tipo inválido",
		},
		{
			enderecoInserido: "Estrada 171",
			retornoEsperado:  "Estrada",
		},
		{
			enderecoInserido: "RUA DOS BOBOS",
			retornoEsperado:  "Rua",
		},
		{
			enderecoInserido: "AVENIDA REBOLSAS",
			retornoEsperado:  "Avenida",
		},
		{
			enderecoInserido: "",
			retornoEsperado:  "Tipo inválido",
		},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			if cenario.retornoEsperado != tipoDeEnderecoRecebido {
				t.Errorf("O tipo recebido é diferente do esperado! Esperado: %s, recebido: %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
			}
		}
	}
}
