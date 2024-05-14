// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}
                    //    pacote de teste padrão
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // roda em paralelo com outros testes
	// enderecoParaTeste := "Rua ABC"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// }

		cenarioDeTeste:= []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA Dos BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}   

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}