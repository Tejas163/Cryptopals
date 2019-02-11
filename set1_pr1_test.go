package cryptopals

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProblem1(t *testing.T) {
	res, err := HextoBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}
	if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error("wrong string", res)
	}
}

func TestProblem2(t *testing.T) {
	res := xor(decodeHex(t, "1c0111001f010100061a024b53535009181c"), decodeHex(t, "686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, decodeHex(t, "746865206b696420646f6e277420706c6179")) {
		t.Errorf("wrong result: %x", res)
	}

}

func decodeHex(t *testing.T, s string) []byte {
	t.Helper()
	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("failed to decode hex", s)
	}
	return v
}

func corpusfromFile(name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintln("failed to read the corpus file", err))
	}
	return buildCorpus(string(text))
}

var corpus = corpusfromFile("E:/cryptopals/testdata/alice_in_wonderland.txt")

func TestProblem3(t *testing.T) {
	res, _, _ := findsingleXOR(decodeHex(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), corpus)
	t.Logf("%s", res)
}
