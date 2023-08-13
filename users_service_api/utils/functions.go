package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type FunctionsInterface interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashed_password string, password string) error
}

type Functions struct {
}

func NewFunctions() *Functions {
	return &Functions{}
}

func (f Functions) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (f Functions) ComparePassword(hashed_password string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
