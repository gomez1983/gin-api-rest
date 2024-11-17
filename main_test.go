package main

import (
	"bytes"
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

func TestListaTodosOsAlunosHandler(t *testing.T) { /** Função de teste que verifica o retorno da listagem de todos os alunos **/
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

func TestBuscaAlunoPorIDHandler(t *testing.T) { /** Função que testa o controlador de busca de aluno por ID **/
	database.ConectaComBancoDeDados()                                                   /** Estabelece a conexão com o banco de dados para o teste **/
	CriaAlunoMock()                                                                     /** Cria um aluno mock para realizar o teste **/
	defer DeletaAlunoMock()                                                             /** Garante que o aluno mock será deletado após o teste **/
	r := SetupDasRotasDeTeste()                                                         /** Configura as rotas para testes utilizando o Gin **/
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)                                   /** Define a rota GET para "/alunos/:id", que chama o controlador BuscaAlunoPorID **/
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)                                        /** Monta o caminho da requisição utilizando o ID do aluno mock **/
	req, _ := http.NewRequest("GET", pathDeBusca, nil)                                  /** Cria uma requisição HTTP GET para buscar o aluno mock pelo ID **/
	resposta := httptest.NewRecorder()                                                  /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                                                          /** Envia a requisição e obtém a resposta simulada **/
	var alunoMock models.Aluno                                                          /** Declara uma variável para armazenar os dados retornados pela API **/
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)                                   /** Converte o JSON da resposta para a estrutura Aluno **/
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais") /** Verifica se o nome retornado corresponde ao mock **/
	assert.Equal(t, "12345678901", alunoMock.CPF)                                       /** Verifica se o CPF retornado corresponde ao mock **/
	assert.Equal(t, "123456789", alunoMock.RG)                                          /** Verifica se o RG retornado corresponde ao mock **/
}

func TestDeletaAlunoHandler(t *testing.T) { /** Função que testa a funcionalidade de deletar aluno pelo ID **/
	database.ConectaComBancoDeDados()                     /** Conecta ao banco de dados para realizar operações durante o teste **/
	CriaAlunoMock()                                       /** Cria um aluno mock para ser deletado no teste **/
	r := SetupDasRotasDeTeste()                           /** Configura o roteador para os testes **/
	r.DELETE("/alunos/:id", controllers.DeletaAluno)      /** Define a rota DELETE para "/alunos/:id", chamando o controlador DeletaAluno **/
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)          /** Constrói dinamicamente o caminho da requisição usando o ID do aluno mock **/
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil) /** Cria a requisição HTTP DELETE para a rota "/alunos/:id" **/
	resposta := httptest.NewRecorder()                    /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                            /** Executa a requisição no roteador e obtém a resposta simulada **/
	assert.Equal(t, http.StatusOK, resposta.Code)         /** Verifica se o código de status retornado é 200 OK **/
}

func TestEditaUmlunoHandler(t *testing.T) { /** Função de teste que verifica a edição de um aluno **/
	database.ConectaComBancoDeDados() /** Conecta ao banco de dados para realizar operações durante o teste **/
	CriaAlunoMock()                   /** Cria um aluno mock para ser editado no teste **/
	defer DeletaAlunoMock()           /** Garante que o aluno mock será deletado após o teste, mantendo o banco de dados limpo **/

	r := SetupDasRotasDeTeste()                    /** Configura as rotas para testes utilizando o Gin **/
	r.PATCH("/alunos/:id", controllers.EditaAluno) /** Define a rota PATCH para "/alunos/:id", que chama o controlador EditaAluno **/

	aluno := models.Aluno{Nome: "Nome do Aluno Teste",
		CPF: "47123456789",
		RG:  "123456700"} /** Cria uma instância da struct Aluno com os novos dados que serão atualizados **/

	valorJson, _ := json.Marshal(aluno)             /** Converte o objeto aluno para JSON para enviar na requisição PATCH **/
	pathParaEditar := "/alunos/" + strconv.Itoa(ID) /** Constrói o caminho da requisição PATCH usando o ID do aluno mock criado **/

	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson)) /** Cria uma nova requisição HTTP PATCH para a rota "/alunos/:id" com o corpo JSON dos dados atualizados **/
	resposta := httptest.NewRecorder()                                             /** Cria um gravador de resposta para simular o comportamento do servidor durante o teste **/

	r.ServeHTTP(resposta, req) /** Envia a requisição PATCH ao roteador configurado e captura a resposta simulada **/

	var alunoMockAtualizado models.Aluno                        /** Declara uma variável para armazenar os dados do aluno atualizado retornados pela resposta **/
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado) /** Converte o corpo da resposta JSON para a struct Aluno, preenchendo alunoMockAtualizado **/

	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)          /** Verifica se o CPF do aluno atualizado corresponde ao esperado ("47123456789") **/
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)             /** Verifica se o RG do aluno atualizado corresponde ao esperado ("123456700") **/
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome) /** Verifica se o Nome do aluno atualizado corresponde ao esperado ("Nome do Aluno Teste") **/

	fmt.Println(alunoMockAtualizado.CPF) /** Imprime o CPF do aluno atualizado no console para fins de depuração **/
}
