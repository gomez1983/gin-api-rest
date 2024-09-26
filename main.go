package main

import "github.com/gin-gonic/gin"

func main() { /** Função principal que inicia o servidor **/
	r := gin.Default() /** Cria uma nova instância do roteador Gin com middleware padrão (logger e recovery) **/
	r.Run(":5000")     /** Inicia o servidor na porta 5000 **/
}
