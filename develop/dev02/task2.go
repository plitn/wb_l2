package dev02

import (
	"strconv"
	"strings"
	"unicode"
)

func Run(str string) string {
	var sBuilder strings.Builder
	r := strings.NewReader(str)
	lastRune, _, _ := r.ReadRune()
	if unicode.IsDigit(lastRune) || str == "" {
		return ""
	}
	for {
		currChar, _, readErr := r.ReadRune()
		if readErr != nil {
			sBuilder.WriteRune(lastRune)
			break
		}
		var quantity int
		if unicode.IsDigit(currChar) {
			quantity, _ = strconv.Atoi(string(currChar))
		}
		if quantity != 0 {
			sBuilder.WriteString(strings.Repeat(string(lastRune), quantity))
		} else {
			sBuilder.WriteRune(lastRune)
		}

		if currChar == '\\' || quantity != 0 {
			lastRune, _, readErr = r.ReadRune()
			if readErr != nil {
				break
			}
		} else {
			lastRune = currChar
		}
	}
	return sBuilder.String()
}
