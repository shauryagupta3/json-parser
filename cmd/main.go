package main

import (
	"log/slog"

	"github.com/shauryagupta3/json-parser/internal/parser"
)

func main() {
	logger := slog.Default()

	textFromFile, err := parser.LoadFile("input.txt")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	checkValidJson := parser.ValidateJson(textFromFile)
	if !checkValidJson {
		logger.Error("Input file is not valid")
		return
	}

	logger.Info("Input file is valid")

	json, err := parser.ParseJson(textFromFile)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	for key, value := range json {
		logger.Info("Logging key-value", slog.String("key", key), slog.Any("value", value))
	}
	return
}
