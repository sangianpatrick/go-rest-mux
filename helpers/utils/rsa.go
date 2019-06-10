package utils

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/sangianpatrick/go-rest-mux/helpers/logger"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "keys/id_rsa"
	publicKeyPath  = "keys/id_rsa.pub"
)

// GetRSAPrivateKey returns rsa sign key
func GetRSAPrivateKey() *rsa.PrivateKey {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		logger.Fatal("GetRSAPrivateKey", err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Fatal("GetRSAPrivateKey", err)
	}
	logger.Info("GetRSAPrivateKey", "Successfuly loaded")
	return signKey
}

// GetRSAPublicKey returns rsa verify key
func GetRSAPublicKey() *rsa.PublicKey {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		logger.Fatal("GetRSAPublicKey", err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Fatal("GetRSAPublicKey", err)
	}
	return verifyKey
}
