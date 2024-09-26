package main

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "André Gomez",
	})
}
func main() { /** Função principal que inicia o servidor **/
	r := gin.Default() /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/

	r.GET("/alunos", ExibeTodosAlunos)
	r.Run(":5000") /** Inicia o servidor na porta 5000 **/
}
