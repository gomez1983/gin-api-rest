package main

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	c.JSON(200, gin.H{ /** Responde com um status HTTP 200 e um JSON contendo um mapa de chave-valor **/
		"id":   "1",           /** Chave "id" com valor "1" **/
		"nome": "André Gomez", /** Chave "nome" com valor "André Gomez" **/
	})
}

func main() { /** Função principal que inicia o servidor **/
	r := gin.Default() /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/

	r.GET("/alunos", ExibeTodosAlunos)
	r.Run(":5000") /** Inicia o servidor na porta 5000 **/
}
