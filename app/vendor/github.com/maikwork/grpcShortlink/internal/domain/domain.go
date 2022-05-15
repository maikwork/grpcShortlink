package domain

import (
	"errors"
	"log"
	"math"
	"regexp"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func checkLink(link string) bool {
	b, err := regexp.MatchString("/^[a-zA-Z0-9_]{3,10}$/", link)
	if err != nil {
		log.Printf("Could not perform")
	}
	return b
}

func Encode(number uint64) string {
	var encode strings.Builder
	encode.Grow(10)

	for ; number > 0; number /= length {
		encode.WriteByte(alphabet[(number % length)])
	}

	return encode.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("Invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
