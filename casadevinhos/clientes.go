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

// processaDados Calculo o total em compras realizado por cada cliente
func processaDados() {
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

	for _, cliente := range listaDeClientes {
		var novoCliente models.Cliente
		novoCliente.NomeCliente = cliente.Nome
		novoCliente.CpfClientes = cliente.CPF
		var quantasVezesComprou int64
		var valorTotalDeCompraHistoricoCliente float64 = 0
		var maiorCompraUnicaDoCliente float64
		for _, compra := range listaDeCompras {
			//verificar se a compra é do cliente e soma o valor total de compra

			if converteCPF(compra.Cliente) == converteCPF(novoCliente.CpfClientes) {
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

		}
		novoCliente.MaiorCompraUnicaDoCliente = maiorCompraUnicaDoCliente
		novoCliente.QuantidadeDeComprasDoCliente = quantasVezesComprou
		novoCliente.TotalEmComprasDoCliente = valorTotalDeCompraHistoricoCliente
		DadosClientes = append(DadosClientes, &novoCliente)

		// fmt.Println("cliente:", novoCliente.CpfClientes, "valor total:", valorTotalDeCompraHistoricoCliente)

		// fmt.Println("o cliente", novoCliente.CpfClientes, "comprou", quantasVezesComprou, "vezes")

		// fmt.Println("o cliente", clienteMaiorCompra, "foi o que fez a maior compra única em 2016, gastando:", valorMaiorCompra, "Reais")
	}
	fmt.Println(DadosClientes)

}

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
