package utility

import (
	"crypto/hmac"
	"crypto/sha256"
)

func VerifyHS(headerPart, payloadPart, signPart string, secret []byte) bool {
	mssg := headerPart + "." + payloadPart
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(mssg))
	res := h.Sum(nil)

	signBytes, err := DecodeB64Url(signPart)
	if err != nil {
		return false
	}
	return hmac.Equal(res, signBytes)
}
