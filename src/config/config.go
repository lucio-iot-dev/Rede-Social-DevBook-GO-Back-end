package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = ""

	// Porta onde a API vai estar rodando
	Porta = 0
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
   var erro error

	 if erro = godotenv.Load(); erro != nil {
		 log.Fatal(erro)
	 }

	 Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))  //  strconv.Atoi ---> converte de string para numero inteiro ( em .env 5000 é uma string)
	 if erro != nil {
		   Porta = 9000
	 }

	 //CRIAR STRING DE CONEXÃO COM O BANCO
	 StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	       os.Getenv("DB_USUARIO"),
	       os.Getenv("DB_SENHA"),
	       os.Getenv("DB_NOME"),
	)


}