package util

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"sipamit-be/internal/pkg/log"
	"time"
	"unsafe"
)

func RandomString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	const str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var letterIdxBits int64 = 6
	var letterIdxMask int64 = 1<<letterIdxBits - 1
	letterIdxMax := 63 / letterIdxBits

	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(str) {
			b[i] = str[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// CheckPassword is function to compare hashed and plain text
func CheckPassword(hashed, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	if err != nil {
		return false
	}
	return true
}

// CryptPassword is a function to encrypt plain text to bcrypt
func CryptPassword(text string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	return string(hashed)
}

func TimeToMilis(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
