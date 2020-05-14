package casadevinhos

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

// converte a data para o formato do GO e retorna somente o ano
func converteData(data string) int {
	layout := "02-01-2006"
	t, err := time.Parse(layout, data)
	if err != nil {
		fmt.Println(err)
	}
	return t.Year()
}

// corrige o problema de CPF com 12 dígitos para permitir a comparação correta
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

// ComparaItens compara 2 itens para identificar vinhos únicos
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
