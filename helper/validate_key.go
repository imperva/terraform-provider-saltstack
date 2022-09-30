package helper

import (
	"errors"
	"strings"

	"crypto/x509"
	"encoding/pem"
)

func ValidateRsaPrivateKey(privateKeyPem string) error {
	block, _ := pem.Decode([]byte(privateKeyPem))

	if block == nil || !strings.Contains(block.Type, "PRIVATE KEY") {
		return errors.New("failed to parse PEM block containing the private key")
	}

	_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	return nil
}

func ValidateRsaPublicKey(publicKeyPem string) error {
	block, _ := pem.Decode([]byte(publicKeyPem))

	if block == nil || !strings.Contains(block.Type, "PUBLIC KEY") {
		return errors.New("failed to parse PEM block containing the public key")
	}

	_, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	return nil
}
