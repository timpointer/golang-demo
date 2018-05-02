package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

const (
	SESSION_LENGTH = 168
)

func generateRandomSession() (string, error) {
	b, err := generateRandomBytes()
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func generateRandomBytes() ([]byte, error) {
	r := make([]byte, SESSION_LENGTH)
	if _, err := rand.Read(r); err != nil {
		return nil, err
	}
	return r, nil
}
