// Teste de unidade
package enderecos_test

import (
	. "introducao-testes/enderecos" //O ponto . faz o Go entender esse pacote com um padrão, não precisando ser referenciado Ex: x.y("Exemplo") -> y("Exemplo")
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) { //Test"NomeDaFunção"(t *testing.T)
	t.Parallel()

	cenarioDeTestes := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada Qualquer", "Estrada"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOLSAS", "Avenida"},
		{"", "Tipo Inválido"},
		//{"Rodovia dos Imigrantes", "Rua"}, Forçando o erro
	}

	for _, cenario := range cenarioDeTestes {
		tipoDeEndereco := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEndereco != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEndereco, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou!!")
	}
}
