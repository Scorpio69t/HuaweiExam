package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var plaintext, ciphertext string

	if scanner.Scan() {
		plaintext = scanner.Text()
	}

	if scanner.Scan() {
		ciphertext = scanner.Text()
	}

	encrypted := encrypt(plaintext)
	decrypted := decrypt(ciphertext)

	fmt.Println(encrypted)
	fmt.Println(decrypted)
}

func encrypt(s string) string {
	result := make([]rune, len(s))
	for i, char := range s {
		if unicode.IsLetter(char) {
			switch char {
			case 'Z':
				result[i] = 'a'
			case 'z':
				result[i] = 'A'
			default:
				result[i] = char + 1
				if unicode.IsUpper(char) {
					result[i] = unicode.ToLower(result[i])
				} else {
					result[i] = unicode.ToUpper(result[i])
				}
			}
		} else if unicode.IsDigit(char) {
			if char == '9' {
				result[i] = '0'
			} else {
				result[i] = char + 1
			}
		} else {
			result[i] = char
		}
	}
	return string(result)
}

func decrypt(s string) string {
	result := make([]rune, len(s))
	for i, char := range s {
		if unicode.IsLetter(char) {
			switch char {
			case 'A':
				result[i] = 'z'
			case 'a':
				result[i] = 'Z'
			default:
				result[i] = char - 1
				if unicode.IsUpper(char) {
					result[i] = unicode.ToLower(result[i])
				} else {
					result[i] = unicode.ToUpper(result[i])
				}
			}
		} else if unicode.IsDigit(char) {
			if char == '0' {
				result[i] = '9'
			} else {
				result[i] = char - 1
			}
		} else {
			result[i] = char
		}
	}
	return string(result)
}
