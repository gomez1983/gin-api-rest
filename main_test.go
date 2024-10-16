package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gomez1983/api-go-gin/controllers"
)

func SetupDasRotasDeTeste() *gin.Engine { /** Função que configura as rotas para testes **/
	rotas := gin.Default() /** Cria uma nova instância do roteador Gin com o middleware padrão (logger e recovery) **/
	return rotas           /** Retorna o roteador configurado **/
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) { /** Função de teste que verifica o status code da rota de saudação **/
	r := SetupDasRotasDeTeste()                     /** Configura as rotas de teste utilizando o Gin **/
	r.GET("/:nome", controllers.Saudacao)           /** Define a rota GET que recebe um parâmetro "nome" e chama a função Saudacao no controlador **/
	req, _ := http.NewRequest("GET", "/andre", nil) /** Cria uma nova requisição HTTP GET com o parâmetro "andre" na URL **/
	resposta := httptest.NewRecorder()              /** Cria um gravador de resposta para simular o comportamento do servidor **/
	r.ServeHTTP(resposta, req)                      /** Envia a requisição e obtém a resposta simulada **/
	if resposta.Code != http.StatusOK {             /** Verifica se o código de status da resposta é diferente de 200 OK **/
		t.Fatalf("Status error: valor recebido foi %d e o esperado foi %d", resposta.Code, http.StatusOK) /** Falha no teste se o status code não for 200 OK, mostrando a diferença **/
	}
}
