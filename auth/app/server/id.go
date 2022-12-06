package server

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func CreateClientID() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}

func CreateCode() (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, 30)
	for i := range buf {
		len := int64(len(letters))
		num, err := rand.Int(rand.Reader, big.NewInt(len))
		if err != nil {
			return "", err
		}
		buf[i] = letters[int(num.Int64())]
	}
	return string(buf), nil
}

func CreateCodeChallenge() (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	buf := make([]rune, 45)
	for i := range buf {
		len := int64(len(letters))
		num, err := rand.Int(rand.Reader, big.NewInt(len))
		if err != nil {
			return "", err
		}
		buf[i] = letters[int(num.Int64())]
	}
	return string(buf), nil
}
