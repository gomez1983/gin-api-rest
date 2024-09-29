package models

import "gorm.io/gorm"

type Aluno struct { /** Define a estrutura "Aluno" com três campos **/
	gorm.Model        /** `gorm.Model` adiciona automaticamente os campos ID, CreatedAt, UpdatedAt e DeletedAt aos modelos. **/
	Nome       string `json:"nome"` /** Campo "Nome" do tipo string, que será convertido para "nome" no JSON **/
	CPF        string `json:"cpf"`  /** Campo "CPF" do tipo string, que será convertido para "cpf" no JSON **/
	RG         string `json:"rg"`   /** Campo "RG" do tipo string, que será convertido para "rg" no JSON **/
}
