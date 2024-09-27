package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()                             /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/
	r.GET("/alunos", controllers.ExibeTodosAlunos) /** Define a rota "/alunos" que usa o método GET para exibir todos os alunos **/
	r.GET("/:nome", controllers.Saudacao)          /** Define a rota dinâmica "/:nome" que usa o método GET para exibir uma saudação personalizada **/
	r.Run(":5000")                                 /** Inicia o servidor na porta 5000 **/
}
