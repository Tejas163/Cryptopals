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

func EnglishScore(b []byte) float64{
    expected freqs :=map[rune] float64 {
        'a': 0.08167,
        'b': 0.01492,
        'c': 0.02782,
        'd': 0.04253,
        'e': 0.12702,
        'f': 0.02228,
        'g': 0.02015,
        'h': 0.06094,
        'i': 0.06966,
        'j': 0.00153,
        'k': 0.00772,
        'l': 0.04025,
        'm': 0.02406,
        'n': 0.06749,
        'o': 0.07507,
        'p': 0.01929,
        'q': 0.00095,
        'r': 0.05987,
        's': 0.06327,
        't': 0.09056,
        'u': 0.02758,
        'v': 0.00978,
        'w': 0.02360,
        'x': 0.00150,
        'y': 0.01974,
        'z': 0.00074,
    }

    score := float64(0)
    freqs := RuneFrequencies(b)

    for ch := range freqs {
        if expectedFreq, ok := expectedFreqs[ch]; ok {
            score += math.Pow(freqs[ch]-expectedFreq, 2.0)
        } else {
            score += freqs[ch]
        }

    }

    return score
}

func RuneFrequencies(b []byte) map[rune]float64 {

    counts := make(map[rune]int)
    for i := 0; i < len(b); i++ {
        counts[rune(b[i])]++
    }

    freqs := make(map[rune]float64)
    for ch := range counts {
        freqs[ch] = float64(counts[ch]) / float64(len(b))
    }

    return freqs
}

func singleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}
	return res
}

func FindSingleCharForXor(b []byte) byte {
	bestScore := float64(-1)
	var bestChar byte

	for ch := 0; ch <= 255; ch++ {
		Xored, _ := SingleXor(byte(ch), b)
		currentScore := EnglishScore(Xored)
		if currentScore < bestScore || bestScore < 0 {
			bestScore = currentScore
			bestChar = byte(ch)
		}
	}

	return byte(bestChar)
}