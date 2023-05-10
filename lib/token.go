package lib

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type AuthorizationInfo struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	LastLogin string `json:"last_login"`
}

var secretKey = []byte("Terminal_123")

func ClaimToken(tokens []string) (response AuthorizationInfo, serr serror.SError) {

	if tokens == nil {
		return response, serror.New("Token tidak ditemukan")
	}

	tokenString, err := parseToken(tokens[0])
	if err != nil {
		return response, serror.NewFromError(err)
	}
	decode, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, serror.NewFromError(fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return response, serror.NewFromError(err)
	}

	resToken := decode.Claims.(jwt.MapClaims)
	response = AuthorizationInfo{
		Id:        fmt.Sprintf("%v", resToken["id"]),
		Username:  fmt.Sprintf("%v", resToken["username"]),
		Status:    fmt.Sprintf("%v", resToken["status"]),
		LastLogin: fmt.Sprintf("%v", resToken["last_login"]),
	}

	return response, nil
}

func parseToken(source string) (token string, err error) {

	separator := " "
	valueSection := 1
	expectedTokenLength := 2

	if source == "" {
		return token, errors.New("Token tidak ditemukan")
	}

	tokens := strings.Split(source, separator)
	if len(tokens) != expectedTokenLength {
		return token, errors.New("Token tidak valid")
	}

	token = tokens[valueSection]
	return token, nil
}

func GenerateJWT(id string, username string, status int, lastLogin string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["status"] = status
	claims["last_login"] = lastLogin

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
