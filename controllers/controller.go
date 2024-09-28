package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	nome := c.Params.ByName("nome") /** Obtém o valor do parâmetro "nome" da rota **/
	c.JSON(200, gin.H{              /** Responde com um status HTTP 200 e um JSON contendo uma saudação **/
		"API diz:": "E aí " + nome + ", tudo beleza?", /** Monta a mensagem personalizada usando o valor do nome **/
	})
}
