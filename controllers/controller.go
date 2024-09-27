package controllers

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	c.JSON(200, gin.H{ /** Responde com um status HTTP 200 e um JSON contendo um mapa de chave-valor **/
		"id":   "1",           /** Chave "id" com valor "1" **/
		"nome": "André Gomez", /** Chave "nome" com valor "André Gomez" **/
	})
}

func Saudacao(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	nome := c.Params.ByName("nome") /** Obtém o valor do parâmetro "nome" da rota **/
	c.JSON(200, gin.H{              /** Responde com um status HTTP 200 e um JSON contendo uma saudação **/
		"API diz:": "E aí " + nome + ", tudo beleza?", /** Monta a mensagem personalizada usando o valor do nome **/
	})
}
