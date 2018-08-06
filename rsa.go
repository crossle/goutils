package goutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func RsaDecrypt(pinToken, sessionId, secret string) (string, error) {
	token, err := base64.StdEncoding.DecodeString(pinToken)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode([]byte(secret))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	keyBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, token, []byte(sessionId))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(keyBytes), nil
}
