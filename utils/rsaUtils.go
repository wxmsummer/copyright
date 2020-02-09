package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

var (
	Base64Encoding = base64.StdEncoding
)

func GenerateKeys() (string, string, error) {
	private, err := rsa.GenerateKey(rand.Reader, keyBit)
	if err != nil {
		return "", "", err
	}
	err = private.Validate()
	if err != nil {
		return "", "", err
	}
	privateBs := x509.MarshalPKCS1PrivateKey(private)
	privateKey := Base64Encoding.EncodeToString(privateBs)
	publicBs := x509.MarshalPKCS1PublicKey(&private.PublicKey)
	publicKey := Base64Encoding.EncodeToString(publicBs)
	return privateKey, publicKey, nil
}

func DecodePrivateKey(key string) (*rsa.PrivateKey, error) {
	bs, err := Base64Encoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(bs)
}

func DecodePublicKey(key string) (*rsa.PublicKey, error) {
	bs, err := Base64Encoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(bs)
}
