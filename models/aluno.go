package models

type Aluno struct { /** Define a estrutura "Aluno" com três campos **/
	Nome string `json:"nome"` /** Campo "Nome" do tipo string, que será convertido para "nome" no JSON **/
	CPF  string `json:"cpf"`  /** Campo "CPF" do tipo string, que será convertido para "cpf" no JSON **/
	RG   string `json:"rg"`   /** Campo "RG" do tipo string, que será convertido para "rg" no JSON **/
}

var Alunos []Aluno /** Declara uma variável global "Alunos" como um slice do tipo "Aluno" **/

/**
Um *slice* é uma estrutura de dados em Go que funciona como uma lista dinâmica.
Ele pode aumentar ou diminuir de tamanho conforme necessário, ao contrário de um array, que tem tamanho fixo.
**/
