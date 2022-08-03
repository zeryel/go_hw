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
	var lastRune rune = -1

	runeParts := []rune(encodedString)

	for i := 0; i < len(encodedString); i++ {
		switch {
		case unicode.IsDigit(runeParts[i]):
			if lastRune == -1 {
				return "", ErrInvalidString
			}

			multiplier, err := strconv.Atoi(string(runeParts[i]))
			if err != nil {
				return "", ErrInvalidString
			}

			builder.WriteString(strings.Repeat(string(lastRune), multiplier))
			lastRune = -1
		case unicode.IsLetter(runeParts[i]):
			if unicode.IsLetter(lastRune) {
				builder.WriteString(string(lastRune))
			}

			lastRune = runeParts[i]
		default:
			return "", ErrInvalidString
		}
	}

	if unicode.IsLetter(lastRune) {
		builder.WriteString(string(lastRune))
	}

	return builder.String(), nil
}
