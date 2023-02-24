package helpers

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret = os.Getenv("JWT_KEYS")

type claims struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}

func NewToken(userId string, role string) *claims {
	times := time.Now().Add(time.Minute * 10).Unix()
	return &claims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(times, 0)),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func CheckToken(token string) (*claims, error) {
	log.Println(token)
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims := tokens.Claims.(*claims)
	return claims, nil
}
