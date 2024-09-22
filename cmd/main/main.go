package main

import (
	"log/slog"
	"os"

	"github.com/wildwind123/extogen/internal/builder"
)

func main() {
	logger := slog.Default()

	sT := builder.ScannerTmpl{
		Package:   "res",
		OptString: true,
		OptInt:    true,
		OptInt64:  true,
	}

	scannerBytes, err := builder.GetScanner(sT)
	if err != nil {
		logger.Error("cant GetScanner", slog.Any("err", err))
		return
	}
	err = os.WriteFile("res/scanner.go", scannerBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}

	convertBytes, err := builder.GetConvert(sT)
	if err != nil {
		logger.Error("cant GetConvert", slog.Any("err", err))
		return
	}
	err = os.WriteFile("res/convert.go", convertBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}
}
