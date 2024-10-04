package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct { /** Define a estrutura "Aluno" com quatro campos **/
	gorm.Model        /** `gorm.Model` adiciona automaticamente os campos ID, CreatedAt, UpdatedAt e DeletedAt aos modelos. **/
	Nome       string `json:"nome" validate:"nonzero"`                /** Campo "Nome" do tipo string, que será convertido para "nome" no JSON e não pode ser vazio **/
	CPF        string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"` /** Campo "CPF" do tipo string, que será convertido para "cpf" no JSON e deve ter exatamente 9 caracteres **/
	RG         string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`   /** Campo "RG" do tipo string, que será convertido para "rg" no JSON e deve ter exatamente 11 caracteres **/
}

func ValidaDadosDeAluno(aluno *Aluno) error { /** Função que valida os dados de um aluno usando um validador **/
	if err := validator.Validate(aluno); err != nil { /** Verifica se há erros de validação nos dados do aluno **/
		return err /** Retorna o erro, se houver **/
	}
	return nil /** Retorna nil se não houver erros, indicando que os dados são válidos **/
}
