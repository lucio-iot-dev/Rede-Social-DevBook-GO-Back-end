package main

import (
	"api/src/config"
	"api/src/router"
	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
)
/*----------------------------ATENÃO----------------------------------------------------------------*/

// ESTA FUNÇÃO SÓ SERVIU PARA GERAR A CHAVE SECRET PARA COLOCAR NO ARQUIVO .ENV PARA UTILIZARMOS NA APLICAÇÃO. ELA FOI COMENDADA PORQUE ELA NÃO PRECISA FICAR INICIALIZANDO TODA VEZ QUE RODAR A APLICAÇÃO.PARA NÃO FICAR GERANDO CHAVES E NÃO SENDO UTILIZADAS.
// DEIXEI COMENTADA PARA VER COMO A CHAVE SECRET FOI GERADA PARA FUTURAS CONSULTAS!

/*======================================================================================================*/

// func init() {
// 	chave := make([]byte, 64)
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

/*======================================================================================================*/


func main() {
	config.Carregar()
	r := router.Gerar()

  fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	
}