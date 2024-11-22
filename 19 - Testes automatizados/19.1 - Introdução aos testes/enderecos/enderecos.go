package enderecos

import "strings"

// TipoDeEndereço Verifica se um endereço tem um valor válido e o retorna
func TipoDeEndereco(enderecos string) string {
	tiposValidos := []string{"rua", "estrada", "avenida", "rodovia"}

	enderecoComLetraMinuscula := strings.ToLower(enderecos)
	primeiraPalavra := strings.Split(enderecoComLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}
