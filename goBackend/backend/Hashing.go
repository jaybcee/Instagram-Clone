package main

import (
	"golang.org/x/crypto/bcrypt"
)

//Bhash hashed a string with bcrypt
func Bhash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckHash compares a password string and it's corresponding hash
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
