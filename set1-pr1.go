package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func hextoBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}
