package authenticate

import (
	"fmt"
	"os"
	"time"

	"github.com/Jardielson-s/api-task/configs"
	"github.com/golang-jwt/jwt/v5"
)

type TokenInfo struct {
	ID          int
	Username    string
	Email       string
	Permissions []interface{}
	Roles       []interface{}
}

func CreateToken(input TokenInfo) (string, error) {
	configs.Envs()

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":          input.ID,
			"username":    input.Username,
			"email":       input.Email,
			"permissions": input.Permissions,
			"roles":       input.Roles,
			"exp":         time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	configs.Envs()

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
