package uwu

import (
	"bytes"
	"unicode"
)

func UwUify(original string) string {
	var buff bytes.Buffer
	strLen := len(original)
	currentIndex := 0
	for currentIndex < strLen {
		char := rune(original[currentIndex])
		isUpper := unicode.IsUpper(char)
		loweredChar := unicode.ToLower(char)
		switch loweredChar {
		case 'r', 'l':
			if isUpper {
				buff.WriteRune('W')
			} else {
				buff.WriteRune('w')
			}
		case 't':
			var nextChar rune
			if currentIndex+1 == strLen {
				buff.WriteRune(char)
				break
			}
			currentIndex++
			nextChar = rune(original[currentIndex])
			if nextChar == 'h' {
				if isUpper {
					buff.WriteRune('D')
				} else {
					buff.WriteRune('d')
				}
			} else {
				buff.WriteRune(char)
				currentIndex--
			}
		case 'a', 'e', 'i', 'o', 'u':
			var nextChar rune
			if currentIndex+1 == strLen {
				buff.WriteRune(char)
				break
			}
			currentIndex++
			nextChar = rune(original[currentIndex])
			if nextChar == 't' {
				buff.WriteString(string([]rune{char, 'w', nextChar}))
			} else {
				buff.WriteRune(char)
				currentIndex--
			}
		case '!':
			buff.WriteString(" UwU")
		default:
			buff.WriteRune(char)
		}
		currentIndex++
	}

	return buff.String()
}
