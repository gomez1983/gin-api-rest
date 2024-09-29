package database

import (
	"log"

	"github.com/gomez1983/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB /** Declara uma variável global para armazenar a conexão com o banco de dados **/
	err error    /** Declara uma variável global para armazenar possíveis erros **/
)

func ConectaComBancoDeDados() { /** Função que conecta com o banco de dados **/
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable" /** String de conexão com o PostgreSQL **/
	DB, err = gorm.Open(postgres.Open(stringDeConexao))                                               /** Abre a conexão com o banco usando GORM **/
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados") /** Exibe um erro e encerra o programa se a conexão falhar **/
	}
	DB.AutoMigrate(&models.Aluno{}) /** Cria ou atualiza a tabela Alunos no banco de dados com base na struct Aluno **/
}
