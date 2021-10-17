package utils

import (
	"crypto/rand"
	"encoding/base32"
)


const (
	// LowerLetters is the list of lowercase letters.
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// UpperLetters is the list of uppercase letters.
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits is the list of permitted digits.
	Digits = "0123456789"

	// Symbols is the list of symbols.
	Symbols = "~!@_+:"
)

func Generate(length int, symbool bool, upper bool, digits bool) string {

	var base = ""+ LowerLetters

	if symbool == true {
		base = base + Symbols
	}
	if upper == true {
		base = base + UpperLetters
	}
	if digits == true {
		base = base + Digits
	}

	return base

	//return getToken(50)

}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}