package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()                                /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/
	r.GET("/alunos", controllers.ExibeTodosAlunos)    /** Define a rota "/alunos" que usa o método GET para exibir todos os alunos **/
	r.GET("/:nome", controllers.Saudacao)             /** Define a rota dinâmica "/:nome" que usa o método GET para exibir uma saudação personalizada **/
	r.POST("/alunos", controllers.CriaNovoAluno)      /** Define a rota "/alunos" para criar um novo aluno usando o método POST **/
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID) /** Define a rota "/alunos/:id" que usa o método GET para buscar um aluno pelo ID, associando-a à função BuscaAlunoPorID no pacote controllers **/
	r.Run()                                           /** Inicia o servidor web na porta padrão (8080) **/

}
