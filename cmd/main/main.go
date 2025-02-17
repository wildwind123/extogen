package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/wildwind123/extogen/internal/builder"
	"github.com/wildwind123/extogen/internal/detector"
)

func main() {
	logger := slog.Default()

	sourceDir := flag.String("source_dir", "", "output of generated files")

	flag.Parse()

	if *sourceDir == "" {
		logger.Error("--source_dir  command required")
		return
	}

	sT, err := detector.ScannerTmpl(path.Join(*sourceDir, "oas_schemas_gen.go"))
	if err != nil {
		logger.Error("cant ScannerTmpl", slog.Any("err", err))
		return
	}

	scannerBytes, err := builder.GetScanner(sT)
	if err != nil {
		logger.Error("cant GetScanner", slog.Any("err", err))
		return
	}
	err = os.WriteFile(fmt.Sprintf("%s/gogen_scanner.go", *sourceDir), scannerBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}

	convertBytes, err := builder.GetConvert(sT)
	if err != nil {
		logger.Error("cant GetConvert", slog.Any("err", err))
		return
	}
	err = os.WriteFile(fmt.Sprintf("%s/gogen_convert.go", *sourceDir), convertBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}

	sqlNullBytes, err := builder.GetSqlNull(sT)
	if err != nil {
		logger.Error("cant GetSqlNull", slog.Any("err", err))
		return
	}
	err = os.WriteFile(fmt.Sprintf("%s/gogen_sqlnull.go", *sourceDir), sqlNullBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}

	pointerBytes, err := builder.GetPointer(sT)
	if err != nil {
		logger.Error("cant GetPointer", slog.Any("err", err))
		return
	}
	err = os.WriteFile(fmt.Sprintf("%s/gogen_pointer.go", *sourceDir), pointerBytes, os.ModeExclusive|os.ModePerm)
	if err != nil {
		logger.Error("cant write file", slog.Any("err", err))
		return
	}
}
