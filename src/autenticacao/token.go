package autenticacao

import (
	// "go/token"
	"api/src/config"
	"time"

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