package casadevinhos

type clienteAPI struct {
	ID   string `json:"id"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
}

type item struct {
	Produto   string `json:"prduto"`
	Variedade string `json:"variedade"`
	Pais      string `json:"pais"`
	Categoria string `json:"categoria"`
	Safra     string `json:"safra"`
	Preco     string `json:"preco"`
}

type compra struct {
	Itens      []item  `json:"itens"`
	Codigo     string  `json:"codigo"`
	Data       string  `json:"data"`
	Cliente    string  `json:"cliente"`
	ValorTotal float64 `json:"valorTotal"`
}
