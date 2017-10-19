package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func hextoBase64(s string) (string, error) {
	val, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	log.Printf("%s", val)
	return base64.StdEncoding.EncodeToString(val), nil
}

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor mismatch length")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
