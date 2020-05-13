package main

import (
	"fmt"

	"std/github.com/jpodlasnisky/DesafioUBOTs/casadevinhos"
)

func main() {
	casadevinhos.RetornaDados()
	fmt.Println("------------Dados retornados----------------")
	fmt.Println("------------Imprime ordanado----------------")
	casadevinhos.OrdenaMaiorCompraUnica()

}
