package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(input string) bool {
	allowedChars := ".-/ "
	for _, char := range input {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}

func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input")
	}

	if IsMorse(input) {
		return morse.DefaultConverter.ToText(input), nil
	}
	return morse.DefaultConverter.ToMorse(input), nil
}
