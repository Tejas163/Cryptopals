package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"unicode/utf8"
)

//HextoBase64 converts a string of hexadecimal value to base64
func HextoBase64(hst string) (string, error) {
	v, err := hex.DecodeString(hst)
	if err != nil {
		return "", err
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}

//xor to XOR both bytes
func xor(a, b []byte) []byte {
	if len(a) > len(b) {
		a = a[:len(b)]
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

//buildCorpus build a  corpus of english letters
func buildCorpus(text string) map[rune]float64 {
	t := make(map[rune]float64)
	for _, char := range text {
		t[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range t {
		t[char] = t[char] / float64(total)
	}
	return t
}

//scoreEnglish-to frequency scoring of english alphabets in the corpus
func scoreEnglish(text string, t map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += t[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}

//SingleXOR single byte XOR key finding
func SingleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, t := range in {
		res[i] = t ^ key
	}
	return res
}

//findsingleXOR function to find the key
func findsingleXOR(in []byte, t map[rune]float64) (res []byte, key byte, score float64) {
	for k := 0; k < 256; k++ {
		out := SingleXOR(in, byte(k))
		s := scoreEnglish(string(out), t)
		if s > score {
			res = out
			score = s
			key = byte(k)
		}
	}
	return
}
