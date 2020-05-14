# Desafio Técnico - UBOTs

Esta é a entrega do "Problema #1: Venda de Vinhos" da UBOTs. Realizado entre 11/05/2020 e 13/05/2020. 


## Motivações

Quando disponibilizado o desafio, foi liberada a escolha da linguagem de preferência, e determinado que seriam avaliados: 

 - Capacidade de resolução de problemas
 - Lógica de Programação
 - Design de Software
 - Clareza e simplicidade de Código
 - Testes
 - Performance

Desta forma procurei atender os requisitos, com base em algo atual e performático. Dispendi algumas horas na pesquisa de qual tecnologia melhor me atenderia na resolução do desafio. Acabei escolhendo trabalhar com Golang, que é uma linguagem que vem crescendo muito rapidamente e oferece condições de um desenvolvimento rápido, com resultados bastante performáticos. Entretanto, havia o porém de eu não conhecer a linguagem. Passei a estudar a linguagem para entender como utilizá-la e entregar o desafio com base nesta linguagem. 
Como havia a possibilidade de entregar uma API REST ou uma UI, foquei primeiramente na API REST. Buscando a melhor forma de documentar a API, encontrei o Swagger, que possui uma implementação para Golang, facilitando assim o trabalho. Mais uma vez, me vi na necessidade de estudar e entender o funcionamento de algo até então novo para mim. Mais algumas horas alocadas no entendimento do Swagger. 
Acredito que a capacidade de resolver problemas não está diretamente vinculada à quantidade de código escrito, e sim na criatividade e praticidade na busca por soluções. Assim, criei a documentação da API necessária para resolver os desafios e comecei a desenvolver as soluções. 
Todo o código foi documentado durante o desenvolvimento.
## O desafio em si

O desafio consiste em desenvolver uma solução que atenda 4 requisitos:

 1. Listar os clientes, ordenando pelo maior valor total em compras;
 2. Mostrar o cliente com maior compra única no ano de 2016;
 3. Listar os clientes mais fiéis;
 4. Recomendar um vinho para um cliente determinado à partir do histórico de compras.

Para tal desenvolvimento foram disponibilizados 2 datasets para validar as funções. 
Com a utilização do Swagger, foi criado todo o webserver e a validação da API que fora documentada conforme o arquivo `swagger.yaml`. 
Tendo o webserver criado e rodando, foram criadas as funções de processamento de dados, tratamento de dados e retornos aos Endpoints. 
## Deploy

Para poder realizar o deploy deste desafio é necessário atender alguns requisitos mínimos de ferramentas instaladas:

 - Golang
 - Git
 
Realize um clone do projeto, dentro do `$GOPATH` definido na sua instalação do Golang. Dê o comando `go get` no diretório do arquivo `main.go` *`$GOPATH\DesafioUBOTs\cmd\desafio-ubots-server>` *.
Após a realização da instalação das bibliotecas necessárias, ainda dentro do diretório, dê o comando `go build`. 
Agora já é possível executar o backend.


# Execução
Para utilizar o backend desenvolvido, basta executar o arquivo *desafio-ubots-server.exe* com o parâmetro *`--port=<porta>`*. 

    C:\Users\User\Desafios\DesafioUBOTs\desafio-ubots-server.exe --port=8080

Foi também disponibilizado o arquivo de exportação do Postman, para ser carregado e realizar os testes do backend.

# A API
Foram desenvolvidos quatro endpoints nesta API REST. 

 - /clientesGastoTotal
	 - Retorna um json com todos os clientes ordenados pelo gasto total no seu histórico de compras
 - /clienteMaiorCompraUnica
	 - Retorna um json contendo o cliente com a maior compra única realizada no ano de 2016
 - /clientesMaisFieis/{tamanhoLista}
	 - Retorna um json de tamanho definido pelo número inteiro {tamanhoLista} passado como parâmetro na URL, contendo os clientes que realizaram mais compras, ordenados de forma decrescente.  
 - /recomendacaoVinho/{cpfCliente}
	 - Retorna um json com o vinho indicado ao cliente, com base no histórico de compras de um cliente determinado pelo CPF passado como parâmetro na URL.

# Próximos passos

 - Ainda que tenha sido criado um cache primário no processamento de dados, ainda cabe realizar a implementação de novos caches que permitirão ao software uma maior performance no tratamento dos dados.
 - Criação de uma rotina de reprocessamento dos dados em intervalos pré estabelecidos, à fim de não ser necessário reiniciar o servidor toda vez que houver uma alteração nos Datasets.
 - Tratar todos os erros possíveis nas funções que o necessitarem. Por exemplo, informar CPF inválido, ou erros internos do servidor. 
 - Desenvolver uma UI que apresente os dados em formato Dashboard para consulta dos dados.
