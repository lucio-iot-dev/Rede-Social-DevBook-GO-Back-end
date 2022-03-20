package autenticacao

import (
	// "go/token"
	"api/src/config"
	"errors"
	"strconv"

	// "api/src/respostas"
	"fmt"
	// "go/token"
	"net/http"
	"strings"
	"time"

	// "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
   permissoes := jwt.MapClaims{}
	 permissoes["authorized"] = true
	 permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //expirar em 6horas por exemplo
	 permissoes["usuarioId"] = usuarioID
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	 return token.SignedString([]byte(config.SecretKey)) //secret
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		  return erro
	}
  if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	 return errors.New("Token inválido")
}

// ExtrairusuarioID retorna o usuarioId que está salvo no token
func ExtrairusuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		  return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f",permissoes["usuarioId"]), 10, 64)// O primeiro parametro  que o ParseUint espera é uma string, mas permissoes[usuarioId] é uma interface. Então tem que ser convertido em uma string para depois ser convertido em um unit.Mas existe ainda uma particularidade na documentação do json web token. Na função CriarToken quando passamos o parametro usuarioID ele vai por padrão ficar salvo como Float la dentro. Então teremos que converter o Float para String para passar essa string para o ParseUint e ai o PàrseUint irá converter essa string para Uint de fato.
		// Sprintf retorna uma string . "%.0f" porque so quero uma casa decimal

		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	  token := r.Header.Get("Authorization")

		//Bearer 123
    // Função para verificar se existe duas strings no token ex: Bearer lldfihgkjxchsjkdçf...
		if len(strings.Split(token, " ")) == 2 {
        return strings.Split(token, " ")[1] 
		}
		return ""
}
  
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil 
}