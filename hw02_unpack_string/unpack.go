package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrInvalidEscaping = errors.New("only digits or slashes can be escaped")

func Unpack(str string) (string, error) {
	var result strings.Builder
	var prevSymbol rune
	var slashed bool
	prevSymbolIsDigit := true

	for _, currentSymbol := range str {
		if unicode.IsDigit(currentSymbol) && !slashed {
			if prevSymbolIsDigit {
				return "", ErrInvalidString
			}

			repeatCount, err := strconv.Atoi(string(currentSymbol))
			if err != nil {
				return "", err
			}

			result.WriteString(strings.Repeat(string(prevSymbol), repeatCount))
			prevSymbolIsDigit = true
			continue
		}

		if slashed && !unicode.IsDigit(currentSymbol) && currentSymbol != '\\' {
			return "", ErrInvalidEscaping
		}

		slashed = currentSymbol == '\\' && !slashed

		if slashed {
			continue
		}

		if !prevSymbolIsDigit {
			result.WriteRune(prevSymbol)
		}

		prevSymbol = currentSymbol
		prevSymbolIsDigit = false
	}

	if !prevSymbolIsDigit {
		result.WriteRune(prevSymbol)
	}

	return result.String(), nil
}
