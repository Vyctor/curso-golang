package enderecos

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua",
		"avenida",
		"estrada",
		"rodovia",
	}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavradoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavradoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavradoEndereco)
	}

	return "Tipo inválido"
}
