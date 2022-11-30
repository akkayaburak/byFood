package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net/mail"
	"unicode"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func IsValidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func IsValidUsername(s string) bool {
	var hasMinLen = false
	if len(s) >= 7 {
		hasMinLen = true
	}
	return hasMinLen
}

func IsValidMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
