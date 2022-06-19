package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	b := strings.Builder{}

	length := len(str)
	isSlash := false

	for k, s := range str {
		if string(s) == `\` {
			isSlash = true
			continue
		}
		if k == 0 && unicode.IsDigit(s) {
			return "", ErrInvalidString
		}
		if k == length-1 && !unicode.IsDigit(s) {
			b.WriteRune(s)
			continue
		}
		if !unicode.IsDigit(s) {
			if unicode.IsDigit(rune(str[k+1])) {
				continue
			}
			if isSlash && string(s) == "n" {
				b.WriteString("\n")
				isSlash = false
				continue
			}
			b.WriteRune(s)
			continue
		}
		if unicode.IsDigit(rune(str[k-1])) {
			return "", ErrInvalidString
		}
		n, _ := strconv.Atoi(string(s))
		if n == 0 {
			isSlash = false
			continue
		}
		if string(rune(str[k-1])) == "n" && isSlash {
			isSlash = false
			b.WriteString(strings.Repeat("\n", n))
			continue
		}
		b.WriteString(strings.Repeat(string(rune(str[k-1])), n))
	}
	return b.String(), nil
}
