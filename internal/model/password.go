package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type password []byte

type key []byte

func (p password) decrypt(k key) (string, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(p) < nonceSize {
		return "", err
	}

	nonce, ciphertext := p[:nonceSize], p[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func encrypt(password string, k key) (password, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(password), nil), nil
}

func keyFromString(s string) key {
	k := []byte(s)
	if len(k) >= 32 {
		return k[:32]
	}

	i := 0
	for len(k) < 32 {
		k = append(k, k[i])
		i++
		if i > len(k) {
			i = 0
		}
	}

	return k
}
