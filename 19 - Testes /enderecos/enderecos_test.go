package enderecos

import "testing"

func TestTipoDeEnte(t *testing.T) {
	enderecoParaTeste := "Rua Paulista 286"

	tipoDeEnderecoEsperado := "Rua"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Error("Esperado:", tipoDeEnderecoEsperado, "Recebido:", tipoDeEnderecoRecebido)
	}
}
