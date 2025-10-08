package utility

import "encoding/base64"

func DecodeB64Url(str string) ([]byte, error) {
	return base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(str)
}

func EncodeB64Url(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
