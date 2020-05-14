package casadevinhos

import (
	"github.com/bradfitz/slice"
	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

var vinhosCompradosCliente = make(map[int][]item)

// RetornaVinhoMaisCompradoPeloCliente retorna o vinho mais comprado por um cliente determinado por parâmetro de CPF.
func RetornaVinhoMaisCompradoPeloCliente(cpf string) *models.VinhoRecomendado {
	processaDados()
	// Ordena descentente os itens comprados pelo Cliente recebido no CPF para indicar o vinho mais comprado
	Dados := vinhosCompradosCliente[converteCPF(cpf)]
	slice.Sort(Dados[:], func(i, j int) bool {
		return Dados[i].Quantidade > Dados[j].Quantidade
	})
	// Popula o retorno do vinho a ser indicado
	recomendaVinho := models.VinhoRecomendado{Produto: Dados[0].Produto,
		Variedade: Dados[0].Variedade,
		Pais:      Dados[0].Pais,
		Categoria: Dados[0].Categoria,
		Safra:     Dados[0].Safra,
		Preco:     Dados[0].Preco}

	return &recomendaVinho
}
