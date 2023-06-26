package crypta

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func loadPrivateKeyFromDir(path string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("coult not get private key: %w", err)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}
	return privateKey, nil
}

func loadPublicKeyFromDir(path string) (*rsa.PublicKey, error) {
	publicKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("coult not get public key: %w", err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}
	switch publicKey := publicKey.(type) {
	case *rsa.PublicKey:
		return publicKey, nil
	default:
		return nil, fmt.Errorf("error casting publicKey to *rsa.PublicKey")
	}
}
