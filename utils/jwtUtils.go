package utils

import (
	"crypto/rsa"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id    string `json:"username"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func LoadPrivateKey(path string) error {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return err
	}
	return nil
}

func LoadPublicKey(path string) error {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return err
	}
	return nil
}

func VerifyJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
