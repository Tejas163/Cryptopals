package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func hextoBase64(s string) (string, error)  {
	val, err := hex.DecodeString(s)
	if err !=nil {
		return "",err
	}
	log.Printf("%s",val)
	return base64.StdEncoding.EncodeToString(val), nil
	}