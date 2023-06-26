package crypta

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type Service struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewService(keysDir string) (*Service, error) {
	publicKey, err := loadPublicKeyFromDir(keysDir + "public.pem")
	if err != nil {
		return nil, fmt.Errorf("error loading public key: %w", err)
	}
	privateKey, err := loadPrivateKeyFromDir(keysDir + "private.pem")
	if err != nil {
		return nil, fmt.Errorf("error loading private key: %w", err)
	}

	return &Service{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

func (s *Service) Encrypt(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, s.publicKey, data)
}

func (s *Service) Decrypt(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, s.privateKey, data)
}
