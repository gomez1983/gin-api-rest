package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine { /** Função que configura as rotas para testes **/
	rotas := gin.Default() /** Cria uma nova instância do roteador Gin com o middleware padrão (logger e recovery) **/
	return rotas           /** Retorna o roteador configurado **/
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) { /** Função de teste que verifica o status code da rota de saudação **/
	r := SetupDasRotasDeTeste()                                          /** Configura as rotas de teste utilizando o Gin **/
	r.GET("/:nome", controllers.Saudacao)                                /** Define a rota GET que recebe um parâmetro "nome" e chama a função Saudacao no controlador **/
	req, _ := http.NewRequest("GET", "/andre", nil)                      /** Cria uma nova requisição HTTP GET com o parâmetro "andre" na URL **/
	resposta := httptest.NewRecorder()                                   /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                                           /** Envia a requisição e obtém a resposta simulada **/
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") /** Verifica se o código de status é 200 OK e exibe a mensagem personalizada se falhar **/
	mockDaResposta := `{"API diz":"E aí andre, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body) /** Lê o corpo da resposta retornada e armazena na variável respostaBody **/

	assert.Equal(t, mockDaResposta, string(respostaBody)) /** Verifica se o corpo da resposta é igual ao esperado e lança erro se não for **/

	fmt.Println(string(respostaBody)) /** Imprime o corpo da resposta recebida no console para fins de depuração **/
	fmt.Println(mockDaResposta)       /** Imprime o corpo mock esperado para comparação visual no console **/
}
