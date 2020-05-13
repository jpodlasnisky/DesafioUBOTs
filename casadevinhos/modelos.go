package casadevinhos

type clienteAPI struct {
	ID   string `json:"id"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
}

type item struct {
	Produto    string  `json:"produto"`
	Variedade  string  `json:"variedade"`
	Pais       string  `json:"pais"`
	Categoria  string  `json:"categoria"`
	Safra      int64   `json:"safra"`
	Preco      float64 `json:"preco"`
	Quantidade int
}

type compra struct {
	Itens      []item  `json:"itens"`
	Codigo     string  `json:"codigo"`
	Data       string  `json:"data"`
	Cliente    string  `json:"cliente"`
	ValorTotal float64 `json:"valorTotal"`
}
