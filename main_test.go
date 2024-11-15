package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
	"github.com/gomez1983/api-go-gin/database"
	"github.com/gomez1983/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine { /** Função que configura as rotas para testes **/
	gin.SetMode(gin.ReleaseMode) /** Define o modo de operação do Gin para "ReleaseMode", reduzindo logs durante os testes **/
	rotas := gin.Default()       /** Cria uma nova instância do roteador Gin com o middleware padrão (logger e recovery) **/
	return rotas                 /** Retorna o roteador configurado **/
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Nome do Aluno Teste",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	database.DB.Create(&aluno) // Cria o aluno no banco de dados
	ID = int(aluno.ID)         // Armazena o ID gerado para uso na exclusão do mock
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID) // Usa o ID armazenado para deletar o aluno mock
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) { /** Função de teste que verifica o status code da rota de saudação **/
	r := SetupDasRotasDeTeste()                                          /** Configura as rotas de teste utilizando o Gin **/
	r.GET("/:nome", controllers.Saudacao)                                /** Define a rota GET que recebe um parâmetro "nome" e chama a função Saudacao no controlador **/
	req, _ := http.NewRequest("GET", "/andre", nil)                      /** Cria uma nova requisição HTTP GET com o parâmetro "andre" na URL **/
	resposta := httptest.NewRecorder()                                   /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                                           /** Envia a requisição e obtém a resposta simulada **/
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") /** Verifica se o código de status é 200 OK e exibe a mensagem personalizada se falhar **/
	mockDaResposta := `{"API diz":"E aí andre, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)          /** Lê o corpo da resposta retornada e armazena na variável respostaBody **/
	assert.Equal(t, mockDaResposta, string(respostaBody)) /** Verifica se o corpo da resposta é igual ao esperado e lança erro se não for **/
	fmt.Println(string(respostaBody))                     /** Imprime o corpo da resposta recebida no console para fins de depuração **/
	fmt.Println(mockDaResposta)                           /** Imprime o corpo mock esperado para comparação visual no console **/
}

func TestListandoTodosOsAlunosHandler(t *testing.T) { /** Função de teste que verifica o retorno da listagem de todos os alunos **/
	database.ConectaComBancoDeDados() /** Conecta ao banco de dados para tornar os dados acessíveis durante o teste **/
	CriaAlunoMock()                   /** Cria um aluno de teste (mock) para garantir que há dados disponíveis na resposta **/
	defer DeletaAlunoMock()           /** Deleta o aluno de teste após o término do teste, garantindo a limpeza do banco de dados **/

	r := SetupDasRotasDeTeste()                    /** Configura as rotas de teste usando o Gin **/
	r.GET("/alunos", controllers.ExibeTodosAlunos) /** Define a rota GET para "/alunos", que chama o controlador ExibeTodosAlunos **/

	req, _ := http.NewRequest("GET", "/alunos", nil) /** Cria uma nova requisição HTTP GET para a rota "/alunos" sem corpo **/
	resposta := httptest.NewRecorder()               /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                       /** Envia a requisição e obtém a resposta simulada **/

	assert.Equal(t, http.StatusOK, resposta.Code) /** Verifica se o código de status é 200 OK, indicando sucesso na chamada da rota **/

	fmt.Println(resposta.Body) /** Imprime o corpo da resposta no console para depuração e verificação manual **/
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) { /** Função de teste para verificar a busca de um aluno pelo CPF **/
	database.ConectaComBancoDeDados()                                /** Conecta ao banco de dados para acesso aos dados durante o teste **/
	CriaAlunoMock()                                                  /** Cria um registro de aluno mock para simular dados reais **/
	defer DeletaAlunoMock()                                          /** Remove o aluno mock após o teste para limpar o banco de dados **/
	r := SetupDasRotasDeTeste()                                      /** Configura as rotas de teste utilizando o Gin **/
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)          /** Define a rota GET para "/alunos/cpf/:cpf", chamando o controlador BuscaAlunoPorCPF **/
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil) /** Cria uma nova requisição HTTP GET para a rota com CPF de teste **/
	resposta := httptest.NewRecorder()                               /** Cria um gravador de resposta para capturar a resposta do servidor **/
	r.ServeHTTP(resposta, req)                                       /** Envia a requisição e armazena a resposta simulada **/

	assert.Equal(t, http.StatusOK, resposta.Code) /** Verifica se o código de status é 200 OK, indicando sucesso na chamada da rota **/
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)

}
