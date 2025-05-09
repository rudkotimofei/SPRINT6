package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(text string) bool {
	for _, char := range text {
		if char == ' ' {
			continue
		}
		if char != '.' && char != '-' {
			return false
		}
	}
	return true
}

func Detection(text string) (string, error) {
	if IsMorse(text) == true {
		return morse.ToText(text), nil
	} else {
		return morse.ToMorse(text), nil
	}
}
