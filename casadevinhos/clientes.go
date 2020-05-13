package casadevinhos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/bradfitz/slice"
	"std/github.com/jpodlasnisky/DesafioUBOTs/models"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// DadosClientes lista contendo dados processados
var DadosClientes []*models.Cliente

var valorMaiorCompra float64 = 0.

var clienteMaiorCompra string

var vinhosCompradosCliente map[int][]item

// processaDados Calculo o total em compras realizado por cada cliente
func processaDados() {

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
	// retorna dados para uso das funções
	for _, cliente := range listaDeClientes {
		var novoCliente models.Cliente
		novoCliente.NomeCliente = cliente.Nome
		novoCliente.Cpf = cliente.CPF
		var quantasVezesComprou int64
		var valorTotalDeCompraHistoricoCliente float64 = 0
		var maiorCompraUnicaDoCliente float64

		vinhosCompradosCliente[converteCPF(novoCliente.Cpf)] = make([]item, 0)
		//a := make([]int, 5)

		for _, compra := range listaDeCompras {
			//verificar se a compra é do cliente e soma o valor total de compra

			if converteCPF(compra.Cliente) == converteCPF(novoCliente.Cpf) {
				valorTotalDeCompraHistoricoCliente += compra.ValorTotal
				quantasVezesComprou++
				if compra.ValorTotal > maiorCompraUnicaDoCliente {
					maiorCompraUnicaDoCliente = compra.ValorTotal
				}
			}

			if converteData(compra.Data) == 2016 {
				if compra.ValorTotal > valorMaiorCompra {
					valorMaiorCompra = compra.ValorTotal
					clienteMaiorCompra = compra.Cliente
				}
			}

			// percorre itens da compra
			for _, item := range compra.Itens {
				var jaComprouEsseVinho = false
				// percorre todos os itens do cliente para comparar com o item da compra
				for _, itemCliente := range vinhosCompradosCliente[converteCPF(compra.Cliente)] {
					// Se o cliente já comprou esse vinho, incrementa a quantidade
					if ComparaItens(item, itemCliente) {
						itemCliente.Quantidade++
						jaComprouEsseVinho = true
					}
				}
				// Se o cliente ainda não comprou esse vinho antes, adiciona o item na lista de vinhos comprados pelo cliente
				if !jaComprouEsseVinho {
					vinhosCompradosCliente[converteCPF(compra.Cliente)] = append(vinhosCompradosCliente[converteCPF(compra.Cliente)], item)
				}
			}

		}
		novoCliente.MaiorCompraUnicaDoCliente = maiorCompraUnicaDoCliente
		novoCliente.QuantidadeDeComprasDoCliente = quantasVezesComprou
		novoCliente.TotalEmComprasDoCliente = valorTotalDeCompraHistoricoCliente
		DadosClientes = append(DadosClientes, &novoCliente)

		// fmt.Println("cliente:", novoCliente.CpfClientes, "valor total:", valorTotalDeCompraHistoricoCliente)

		// fmt.Println("o cliente", novoCliente.CpfClientes, "comprou", quantasVezesComprou, "vezes")

		// fmt.Println("o cliente", clienteMaiorCompra, "foi o que fez a maior compra única em 2016, gastando:", valorMaiorCompra, "Reais")
	}

}

// Fim da organização dos dados

// RetornaDados Retorna os dados para API
func RetornaDados() []*models.Cliente {
	processaDados()
	return DadosClientes
}

func converteData(data string) int {
	layout := "02-01-2006"
	t, err := time.Parse(layout, data)
	if err != nil {
		fmt.Println(err)
	}
	return t.Year()
}
func converteCPF(cpf string) int {

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(cpf, "")
	resultado, err := strconv.Atoi(processedString)
	if err != nil {
		fmt.Println(err)
	}

	return resultado

}

// RetornaMaiorCompraUnica ordena o array por compra unica do cliente
func RetornaMaiorCompraUnica() *models.Cliente {
	processaDados()
	Dados := DadosClientes
	slice.Sort(Dados[:], func(i, j int) bool {
		return Dados[i].MaiorCompraUnicaDoCliente > Dados[j].MaiorCompraUnicaDoCliente
	})
	return Dados[0]
}

// OrdenaClientesMaiorValorTotalEmCompras Ordena lista de clientes por maior valor total em compras
func OrdenaClientesMaiorValorTotalEmCompras() []*models.Cliente {
	processaDados()
	Dados := DadosClientes
	slice.Sort(Dados[:], func(i, j int) bool {
		return Dados[i].TotalEmComprasDoCliente > Dados[j].TotalEmComprasDoCliente
	})
	return Dados
}

// RetornaClientesMaisFieis Ordena lista de clientes por numero total de compras realizadas
func RetornaClientesMaisFieis() []*models.Cliente {
	processaDados()
	Dados := DadosClientes
	slice.Sort(Dados[:], func(i, j int) bool {
		return Dados[i].QuantidadeDeComprasDoCliente > Dados[j].QuantidadeDeComprasDoCliente
	})
	return Dados

	// fmt.Println("cliente:", novoCliente.CpfClientes, "valor total:", valorTotalDeCompraHistoricoCliente)

	// fmt.Println("o cliente", novoCliente.CpfClientes, "comprou", quantasVezesComprou, "vezes")

	// fmt.Println("o cliente", clienteMaiorCompra, "foi o que fez a maior compra única em 2016, gastando:", valorMaiorCompra, "Reais")
}

// ComparaItens compara item a item para contagem de vinhos unicos por cliente
func ComparaItens(itemA, itemB item) bool {
	if itemA.Produto == itemB.Produto &&
		itemA.Categoria == itemB.Categoria &&
		itemA.Variedade == itemB.Variedade &&
		itemA.Safra == itemB.Safra &&
		itemA.Pais == itemB.Pais {
		return true
	}
	return false
}

// RetornaVinhoMaisCompradoPeloCliente ordena o array por compra unica do cliente
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
