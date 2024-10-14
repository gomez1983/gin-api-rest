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

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
