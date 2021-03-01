package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Hmac256(password string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
