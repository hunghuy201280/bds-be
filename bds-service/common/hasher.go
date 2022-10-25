package common

import (
	"crypto/md5"
	"golang.org/x/crypto/bcrypt"
	"hash"
)

type Hasher interface {
	HashPassword(data string) (string, error)
	ValidatePassword(hashedPassword string, rawPassword string) bool
}

type passwordHash struct {
	hasher hash.Hash
}

func NewMd5Hash() *passwordHash {
	hasher := md5.New()

	return &passwordHash{
		hasher,
	}
}

func (h *passwordHash) HashPassword(data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (h *passwordHash) ValidatePassword(hashedPassword string, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err != nil
}
