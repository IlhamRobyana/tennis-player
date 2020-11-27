package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilhamrobyana/tennis-player/entity"
)

func GenerateToken(player entity.Player) (token string, e error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Hour * 24 * 7)
	claims["username"] = player.Username
	claims["id"] = player.ID
	claims["exp"] = exp.Unix()

	secret := os.Getenv("TOKEN_SECRET")
	token, e = t.SignedString([]byte(secret))
	return
}

func ParseToken(token string) (claims jwt.MapClaims, e error) {
	parsed, e := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("TOKEN_SECRET")
		return []byte(secret), nil
	})

	if e != nil {
		return
	}

	claims, _ = parsed.Claims.(jwt.MapClaims)
	return
}
