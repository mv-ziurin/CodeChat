package util

import (
	"strings"
	"errors"
)

var (
	parsingEmailError = errors.New("bad email")
)

func EmailParse(userEmail string) (string, string, error) {
	components := strings.Split(userEmail, "@")
	if len(components) != 2 {
		return "", "", parsingEmailError
	}

	return components[0], components[1], nil
}
