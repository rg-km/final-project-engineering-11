package secure

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/config"
)

func Authentication(c *gin.Context) error {
	token, err := VerifyToken(c)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := TakeToken(c)
	if len(tokenString) == 0 {
		return nil, errors.New("Token not found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TakeToken(c *gin.Context) string {
	keys := c.GetHeader("Authorization")

	barearkeys := strings.Split(keys, " ")
	if len(barearkeys) > 1 {
		return barearkeys[1]
	} else {
		return ""
	}

}
