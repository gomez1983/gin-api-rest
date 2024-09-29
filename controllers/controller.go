package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/database"
	"github.com/gomez1983/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) { /** Define a função para exibir todos os alunos **/
	var alunos []models.Aluno /** Declara uma variável "alunos" como um slice de modelos do tipo Aluno **/
	database.DB.Find(&alunos) /** Busca todos os registros da tabela correspondente a "Aluno" no banco de dados e os armazena em "alunos" **/
	c.JSON(200, alunos)       /** Retorna os alunos no formato JSON com status HTTP 200 (OK) **/
}

func Saudacao(c *gin.Context) { /** Define a função que recebe o contexto da requisição **/
	nome := c.Params.ByName("nome") /** Obtém o valor do parâmetro "nome" da rota **/
	c.JSON(200, gin.H{              /** Responde com um status HTTP 200 e um JSON contendo uma saudação **/
		"API diz:": "E aí " + nome + ", tudo beleza?", /** Monta a mensagem personalizada usando o valor do nome **/
	})
}
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno                           // Declara uma variável do tipo Aluno para armazenar os dados do novo aluno
	if err := c.ShouldBindJSON(&aluno); err != nil { // Tenta vincular os dados JSON da requisição à estrutura Aluno
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Retorna um erro 400 se a vinculação falhar
		return
	}
	database.DB.Create(&aluno)   // Cria um novo registro de aluno no banco de dados
	c.JSON(http.StatusOK, aluno) // Retorna os dados do aluno recém-criado
}

func BuscaAlunoPorID(c *gin.Context) { /** Define a função para buscar um aluno pelo ID **/
	var aluno models.Aluno        /** Declara uma variável "aluno" do tipo Aluno **/
	id := c.Params.ByName("id")   /** Obtém o parâmetro "id" da URL **/
	database.DB.First(&aluno, id) /** Busca o primeiro registro de Aluno que corresponde ao ID no banco de dados **/

	if aluno.ID == 0 { /** Verifica se o ID do aluno é zero, indicando que nenhum aluno foi encontrado **/
		c.JSON(http.StatusNotFound, gin.H{ /** Retorna um JSON com status 404 (Não encontrado) **/
			"Not found": "Aluno não encontrado"})
		return /** Interrompe a execução da função se o aluno não for encontrado **/
	}

	c.JSON(http.StatusOK, aluno) /** Retorna o aluno encontrado no formato JSON com status HTTP 200 (OK) **/
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
