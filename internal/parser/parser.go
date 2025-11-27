package parser

import (
	"errors"
	"os"
	"strings"
)

func LoadFile(pathToFile string) (string, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ValidateJson(text string) bool {
	if text == "" {
		return false
	}

	if text[0] == '{' && text[len(text)-1] == '}' {
		return true
	}

	return false
}

func ParseJson(text string) (map[string]interface{}, error) {
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "{")
	text = strings.TrimSuffix(text, "}")

	if text == "" {
		return make(map[string]interface{}), nil
	}

	parts := strings.SplitN(text, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid json input")
	}

	mapData := make(map[string]interface{})
	mapData[clean(parts[0])] = clean(parts[1])

	return mapData, nil
}

func clean(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"`)
	return s
}
