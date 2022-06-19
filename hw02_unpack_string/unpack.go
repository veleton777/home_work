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
		if k == length-1 && !unicode.IsDigit(s) {
			b.WriteRune(s)
			continue
		}
		if k == 0 {
			if unicode.IsDigit(s) {
				return "", ErrInvalidString
			}
		}
		if !unicode.IsDigit(s) {
			if k != length-1 && unicode.IsDigit(rune(str[k+1])) {
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
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return "", err
		}
		if n == 0 {
			isSlash = false
			continue
		}
		tmpStr := ""
		if string(rune(str[k-1])) == "n" && isSlash {
			tmpStr = "\n"
			isSlash = false
		} else {
			tmpStr = string(rune(str[k-1]))
		}
		b.WriteString(strings.Repeat(tmpStr, n))
	}
	return b.String(), nil
}
