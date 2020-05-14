package casadevinhos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

// Para uso na busca dos dados
var myClient = &http.Client{Timeout: 10 * time.Second}

// DadosClientes lista contendo dados processados
var DadosClientes []*models.Cliente

var clienteMaiorCompra clienteAPI

var valorMaiorCompra float64 = 0.

// processaDados retorna dados para uso das funções
func processaDados() {
	if DadosClientes == nil {

		// Busca dados das APIs
		listaDeClientes := make([]clienteAPI, 0)

		r, err := myClient.Get("http://www.mocky.io/v2/598b16291100004705515ec5")
		if err != nil {
			print(err)
		}
		defer r.Body.Close()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		json.Unmarshal(bodyBytes, &listaDeClientes)

		listaDeCompras := make([]compra, 0)

		r, err = myClient.Get("http://www.mocky.io/v2/598b16861100004905515ec7")
		if err != nil {
			print(err)
		}
		defer r.Body.Close()
		bodyBytes, err = ioutil.ReadAll(r.Body)
		json.Unmarshal(bodyBytes, &listaDeCompras)
		// Fim busca dados das APIs

		// Percorre todos Clientes e para cada cliente, percorre todas as compras
		for _, cliente := range listaDeClientes {
			// Criação de um novoCliente para poder adicionar na lista de dados processados
			var novoCliente models.Cliente
			novoCliente.NomeCliente = cliente.Nome
			novoCliente.Cpf = cliente.CPF
			var quantasVezesComprou int64
			var valorTotalDeCompraHistoricoCliente float64 = 0
			var maiorCompraUnicaDoCliente float64

			// map de clientes x vinhos comprados
			vinhosCompradosCliente[converteCPF(novoCliente.Cpf)] = make([]item, 0)

			for _, compra := range listaDeCompras {

				//verificar se a compra é do cliente e armazena a maior compra unica do cliente (histórico)
				if converteCPF(compra.Cliente) == converteCPF(novoCliente.Cpf) {
					valorTotalDeCompraHistoricoCliente += compra.ValorTotal
					quantasVezesComprou++
					if compra.ValorTotal > maiorCompraUnicaDoCliente {
						maiorCompraUnicaDoCliente = compra.ValorTotal
					}
				}

				// verifica se a compra é do cliente, se é de 2016, e verifica a maior compra
				if converteCPF(compra.Cliente) == converteCPF(novoCliente.Cpf) {
					if converteData(compra.Data) == 2016 {
						if compra.ValorTotal > valorMaiorCompra {
							valorMaiorCompra = compra.ValorTotal
							clienteMaiorCompra = cliente
						}
					}
				}

				// verifica se a compra é do cliente e cria um array de vinhos comprados pelo cliente, incrementando um contador de quantidade de vinhos identicos comprados.
				if converteCPF(compra.Cliente) == converteCPF(novoCliente.Cpf) {
					// percorre itens da compra
					for _, item := range compra.Itens {
						var jaComprouEsseVinho = false
						// percorre todos os itens do cliente para comparar com o item da compra
						for indice, itemCliente := range vinhosCompradosCliente[converteCPF(compra.Cliente)] {
							// Se o cliente já comprou esse vinho, incrementa a quantidade
							if ComparaItens(item, itemCliente) {
								vinhosCompradosCliente[converteCPF(compra.Cliente)][indice].Quantidade++
								jaComprouEsseVinho = true
							}
						}
						// Se o cliente ainda não comprou esse vinho antes, adiciona o item na lista de vinhos comprados pelo cliente
						if !jaComprouEsseVinho {
							vinhosCompradosCliente[converteCPF(compra.Cliente)] = append(vinhosCompradosCliente[converteCPF(compra.Cliente)], item)
						}
					}
				}

			}
			// Adiciona os dados no novo cliente
			novoCliente.MaiorCompraUnicaDoCliente = maiorCompraUnicaDoCliente
			novoCliente.QuantidadeDeComprasDoCliente = quantasVezesComprou
			novoCliente.TotalEmComprasDoCliente = valorTotalDeCompraHistoricoCliente
			// Armazena os dados tratados em uma variável usada pelas outras funções
			DadosClientes = append(DadosClientes, &novoCliente)
		}

	}

}

// Fim da organização dos dados
