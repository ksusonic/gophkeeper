package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

const keySize = 2048

func main() {
	err := os.Mkdir("cert", 0744)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}
	_ = os.Chdir("cert")

	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = os.WriteFile("private.pem", privateKeyPEM, 0600)
	if err != nil {
		panic(err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	err = os.WriteFile("public.pem", publicKeyPEM, 0644)
	if err != nil {
		panic(err)
	}
	currentPath, _ := os.Getwd()
	fmt.Printf("Your public and private key are generated to %s\n", currentPath)
}
