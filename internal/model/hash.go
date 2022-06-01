package model

import (
	"crypto/sha512"
)

type hash []byte

func hashString(secret string) (hash, error) {
	sha512 := sha512.New()
	_, err := sha512.Write([]byte(secret))
	if err != nil {
		return nil, err
	}

	return sha512.Sum(nil), nil
}
