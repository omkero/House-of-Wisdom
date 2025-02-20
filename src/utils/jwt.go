package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Json-Web-Token
type Jwt struct{}

func (j *Jwt) SignAuthToken(username string, user_id int, secreteKey []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"bis": user_id,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration time 1day
		"iat": time.Now().Unix(),                     // Issued at
	})

	tokenString, err := token.SignedString(secreteKey)
	fmt.Println(err)

	if err != nil {
		return ""
	}
	return tokenString
}

func (j *Jwt) VerifyToken(token_string string, secreteKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
		return secreteKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}

func (j *Jwt) DecodeJwtToken(ctx *gin.Context) (*jwt.Token, error) {

	keyAuth := ctx.GetHeader("Authorization")

	if keyAuth == "" {
		return nil, errors.New("no token")
	}

	headerToken := keyAuth[len("Bearer "):]

	if headerToken == "" {
		return nil, errors.New("no token")
	}

	isValidToken, err := j.VerifyToken(headerToken, []byte(os.Getenv("LOGIN_PRIVATE_KEY")))

	if err != nil {
		return nil, err
	}

	if !isValidToken.Valid {
		return nil, errors.New("token is not valid")
	}

	return isValidToken, nil
}
