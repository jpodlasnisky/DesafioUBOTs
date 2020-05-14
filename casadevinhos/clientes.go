package casadevinhos

import (
	"errors"

	"github.com/bradfitz/slice"
	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

// RetornaMaiorCompraUnica Busca o cliente dos dados processados com a maior compra única e retorna
func RetornaMaiorCompraUnica() (*models.Cliente, error) {
	processaDados()
	Dados := DadosClientes
	for _, busca := range Dados {
		if converteCPF(busca.Cpf) == converteCPF(clienteMaiorCompra.CPF) {
			return busca, nil
		}
	}
	return nil, errors.New("Erro ao processar as compras")
}

// variavel utilizada para cache dos dados
var clientesOrdenadosMaiorValorTotalEmCompras []*models.Cliente = nil

// OrdenaClientesMaiorValorTotalEmCompras Ordena lista de clientes por maior valor total em compras
func OrdenaClientesMaiorValorTotalEmCompras() []*models.Cliente {
	if clientesOrdenadosMaiorValorTotalEmCompras == nil {
		processaDados()
		clientesOrdenadosMaiorValorTotalEmCompras = DadosClientes
		slice.Sort(clientesOrdenadosMaiorValorTotalEmCompras[:], func(i, j int) bool {
			return clientesOrdenadosMaiorValorTotalEmCompras[i].TotalEmComprasDoCliente > clientesOrdenadosMaiorValorTotalEmCompras[j].TotalEmComprasDoCliente
		})
	}

	return clientesOrdenadosMaiorValorTotalEmCompras
}

// RetornaClientesMaisFieis Ordena lista de clientes por numero total de compras realizadas com base em um parâmetro inteiro que define o tamanho do retorno
func RetornaClientesMaisFieis(tamanhoLista int64) []*models.Cliente {
	processaDados()
	Dados := DadosClientes
	slice.Sort(Dados[:], func(i, j int) bool {
		return Dados[i].QuantidadeDeComprasDoCliente > Dados[j].QuantidadeDeComprasDoCliente
	})
	return Dados[:tamanhoLista]
}
