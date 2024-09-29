package main

import (
	"github.com/gomez1983/api-go-gin/database"
	"github.com/gomez1983/api-go-gin/routes"
)

/**
O GIN é um framework web escrito em Go que facilita o desenvolvimento de APIs e aplicações web.
Ele é utilizado neste projeto porque é leve, rápido e possui uma sintaxe simples, o que permite criar rotas, middlewares e lidar com requisições HTTP de maneira eficiente.
O GIN também inclui funcionalidades como recuperação de erros, logging e serialização de JSON, tornando o desenvolvimento mais ágil e organizado.
**/

func main() { /** Função principal que inicia a aplicação **/
	database.ConectaComBancoDeDados() /** Conecta com o banco de dados utilizando a função definida no pacote database **/
	routes.HandleRequests()           /** Configura as rotas e inicia o servidor web **/
}
