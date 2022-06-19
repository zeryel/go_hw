package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(encodedString string) (string, error) {
	var builder strings.Builder
	var runeParts = []rune(encodedString)
	var lastRune rune = -1

	for i := 0; i < len(encodedString); i++ {
		if unicode.IsDigit(runeParts[i]) {
			if lastRune == -1 {
				return "", ErrInvalidString
			} else {
				multiplier, err := strconv.Atoi(string(runeParts[i]))
				if err != nil {
					return "", ErrInvalidString
				}

				builder.WriteString(strings.Repeat(string(lastRune), multiplier))
				lastRune = -1
			}
		} else if unicode.IsLetter(runeParts[i]) {
			if unicode.IsLetter(lastRune) {
				builder.WriteString(string(lastRune))
			}

			lastRune = runeParts[i]
		} else {
			return "", ErrInvalidString
		}
	}

	if unicode.IsLetter(lastRune) {
		builder.WriteString(string(lastRune))
	}

	return builder.String(), nil
}
