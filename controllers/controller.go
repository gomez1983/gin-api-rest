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
		"API diz": "E aí " + nome + ", tudo beleza?", /** Monta a mensagem personalizada usando o valor do nome **/
	})
}

func CriaNovoAluno(c *gin.Context) { /** Função para criar um novo aluno no banco de dados **/
	var aluno models.Aluno /** Declara uma variável "aluno" do tipo Aluno **/

	if err := c.ShouldBindJSON(&aluno); err != nil { /** Tenta vincular o corpo da solicitação JSON aos campos do aluno **/
		c.JSON(http.StatusBadRequest, gin.H{ /** Se houver erro de vinculação, retorna um status 400 (Bad Request) com a mensagem de erro **/
			"erro": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil { /** Valida os dados do aluno utilizando a função ValidaDadosDeAluno **/
		c.JSON(http.StatusBadRequest, gin.H{ /** Se houver erro de validação, retorna um status 400 (Bad Request) com a mensagem de erro **/
			"erro": err.Error()})
		return
	}

	database.DB.Create(&aluno)   /** Insere o novo aluno no banco de dados **/
	c.JSON(http.StatusOK, aluno) /** Retorna uma resposta de sucesso com o aluno recém-criado **/
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

func DeletaAluno(c *gin.Context) { /** Define a função para deletar um aluno **/
	var aluno models.Aluno                                             /** Declara uma variável "aluno" do tipo Aluno **/
	id := c.Params.ByName("id")                                        /** Obtém o parâmetro "id" da URL **/
	database.DB.Delete(&aluno, id)                                     /** Deleta o registro do aluno no banco de dados que corresponde ao ID **/
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"}) /** Retorna um JSON com status 200 (OK) confirmando a exclusão do aluno **/
}

func EditaAluno(c *gin.Context) { /** Define a função para editar parcialmente um aluno **/
	var aluno models.Aluno        /** Declara uma variável "aluno" do tipo Aluno **/
	id := c.Params.ByName("id")   /** Obtém o parâmetro "id" da URL **/
	database.DB.First(&aluno, id) /** Busca o registro do aluno no banco de dados com base no ID **/

	if err := c.ShouldBindJSON(&aluno); err != nil { /** Tenta vincular o corpo da solicitação JSON aos campos do aluno **/
		c.JSON(http.StatusBadRequest, gin.H{ /** Se ocorrer um erro de vinculação, retorna um status 400 (Bad Request) com a mensagem de erro **/
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosDeAluno(&aluno); err != nil { /** Valida os dados do aluno usando a função ValidaDadosDeAluno **/
		c.JSON(http.StatusBadRequest, gin.H{ /** Se houver erro de validação, retorna um status 400 (Bad Request) com a mensagem de erro **/
			"erro": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno) /** Atualiza as colunas do aluno no banco de dados com os novos valores **/
	c.JSON(http.StatusOK, aluno)                   /** Retorna um JSON com status 200 (OK) e os dados atualizados do aluno **/
}

func BuscaAlunoPorCPF(c *gin.Context) { /** Define a função que recebe um Context do Gin para tratar a requisição **/
	var aluno models.Aluno                                   /** Declara uma variável do tipo Aluno para armazenar o resultado da busca **/
	cpf := c.Param("cpf")                                    /** Extrai o valor do parâmetro 'cpf' da URL **/
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno) /** Realiza uma consulta no banco de dados para encontrar o aluno com o CPF fornecido **/

	if aluno.ID == 0 { /** Verifica se o aluno foi encontrado, se não foi, ID será 0 **/
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"}) /** Retorna um erro 404 indicando que o aluno não foi encontrado **/
		return
	}

	c.JSON(http.StatusOK, aluno) /** Retorna o aluno encontrado com o status 200 (OK) **/
}

func ExibePaginaIndex(c *gin.Context) {
	// Declara a variável 'alunos' que será preenchida com os dados do banco
	var alunos []models.Aluno

	// Realiza a consulta no banco de dados para obter todos os alunos
	// A função 'Find' recupera todos os registros da tabela de alunos
	// e os armazena na variável 'alunos'
	database.DB.Find(&alunos)

	// Renderiza a página HTML com o código de status 200 (OK) e passa os dados
	// para o template 'index.html'. Os dados são passados em um mapa gin.H
	c.HTML(
		http.StatusOK, // Código de status HTTP para a resposta (200 OK)
		"index.html",  // Nome do arquivo HTML que será renderizado
		gin.H{ // Mapa de dados a ser passado para o template
			"alunos": alunos, // A variável 'alunos' contém a lista de alunos
		},
	)
}

func RotaNaoEncontrada(c *gin.Context) { // Função que lida com requisições para rotas não encontradas (erro 404)
	c.HTML(http.StatusNotFound, "404.html", nil) // Retorna uma resposta HTML com o código de status 404 e renderiza o template "404.html"
}
