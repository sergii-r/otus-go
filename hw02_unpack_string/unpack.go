package hw02_unpack_string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

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

			repeatCount, _ := strconv.Atoi(string(currentSymbol))
			result.WriteString(strings.Repeat(string(prevSymbol), repeatCount))
			prevSymbolIsDigit = true
			continue
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
