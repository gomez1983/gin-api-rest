package main

import (
	"github.com/gomez1983/api-go-gin/database"
	"github.com/gomez1983/api-go-gin/models"
	"github.com/gomez1983/api-go-gin/routes"
)

/**
O GIN é um framework web escrito em Go que facilita o desenvolvimento de APIs e aplicações web.
Ele é utilizado neste projeto porque é leve, rápido e possui uma sintaxe simples, o que permite criar rotas, middlewares e lidar com requisições HTTP de maneira eficiente.
O GIN também inclui funcionalidades como recuperação de erros, logging e serialização de JSON, tornando o desenvolvimento mais ágil e organizado.
**/

func main() { /** Função principal que inicia a aplicação **/
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{ /** Inicializa um slice de alunos com dados fictícios **/
		{Nome: "André Gomez", CPF: "00000000000", RG: "11111111"}, /** Adiciona o aluno André Gomez com CPF e RG **/
		{Nome: "Carol", CPF: "11122233314", RG: "212352168"},      /** Adiciona a aluna Carol com CPF e RG **/
	}

	routes.HandleRequests() /** Chama a função que configura as rotas e inicia o servidor **/
}
