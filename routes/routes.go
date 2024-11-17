package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()              /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/
	r.Static("/assets", "./assets") // Serve arquivos estáticos localizados no diretório "./assets" quando acessados pela URL "/assets"

	r.LoadHTMLGlob("templates/*") /** Carrega todos os arquivos HTML da pasta "templates" para serem usados em renderizações de página **/

	/** Rotas definidas para a aplicação **/
	r.GET("/alunos", controllers.ExibeTodosAlunos) /**
	  - Rota: "/alunos"
	  - Método: GET
	  - Ação: Exibe todos os alunos cadastrados
	  - Controlador: ExibeTodosAlunos
	  **/

	r.GET("/:nome", controllers.Saudacao) /**
	  - Rota: "/:nome" (dinâmica)
	  - Método: GET
	  - Ação: Exibe uma saudação personalizada com base no nome fornecido na URL
	  - Controlador: Saudacao
	  **/

	r.POST("/alunos", controllers.CriaNovoAluno) /**
	  - Rota: "/alunos"
	  - Método: POST
	  - Ação: Cria um novo registro de aluno no banco de dados
	  - Controlador: CriaNovoAluno
	  **/

	r.GET("/alunos/:id", controllers.BuscaAlunoPorID) /**
	  - Rota: "/alunos/:id" (dinâmica)
	  - Método: GET
	  - Ação: Busca um aluno pelo ID fornecido na URL
	  - Controlador: BuscaAlunoPorID
	  **/

	r.DELETE("/alunos/:id", controllers.DeletaAluno) /**
	  - Rota: "/alunos/:id" (dinâmica)
	  - Método: DELETE
	  - Ação: Exclui o registro de um aluno com base no ID fornecido
	  - Controlador: DeletaAluno
	  **/

	r.PATCH("/alunos/:id", controllers.EditaAluno) /**
	  - Rota: "/alunos/:id" (dinâmica)
	  - Método: PATCH
	  - Ação: Atualiza parcialmente os dados de um aluno com base no ID fornecido
	  - Controlador: EditaAluno
	  **/

	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF) /**
	  - Rota: "/alunos/cpf/:cpf" (dinâmica)
	  - Método: GET
	  - Ação: Busca um aluno pelo CPF fornecido na URL
	  - Controlador: BuscaAlunoPorCPF
	  **/

	r.GET("/index", controllers.ExibePaginaIndex) /**
	  - Rota: "/index"
	  - Método: GET
	  - Ação: Renderiza uma página inicial utilizando um arquivo HTML carregado em "templates"
	  - Controlador: ExibePaginaIndex
	  **/

	r.Run() /**
	  - Inicia o servidor web usando a porta padrão (8080) ou uma porta especificada pela variável de ambiente `PORT`
	  - Exemplo: Pode ser acessado localmente em http://localhost:8080
	  **/
}
