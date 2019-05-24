package utils

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"

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
		log.Fatalf("RSA Error: %s", err)
		os.Exit(1)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("RSA Error: %s", err)
		os.Exit(1)
	}
	return signKey
}

// GetRSAPublicKey returns rsa verify key
func GetRSAPublicKey() *rsa.PublicKey {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("RSA Error: %s", err)
		os.Exit(1)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("RSA Error: %s", err)
		os.Exit(1)
	}
	return verifyKey
}
