package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine { /** Função que configura as rotas para testes **/
	rotas := gin.Default() /** Cria uma nova instância do roteador Gin com o middleware padrão (logger e recovery) **/
	return rotas           /** Retorna o roteador configurado **/
}

func TestFalhador(t *testing.T) { /** Função de teste que falha propositalmente **/
	t.Fatalf("Teste falhou de propósito. Não se preocupe") /** Força uma falha no teste e exibe a mensagem especificada **/
}
