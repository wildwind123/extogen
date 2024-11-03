package builder

import (
	"bufio"
	"bytes"
	"html/template"

	"github.com/go-faster/errors"
)

type Builder struct {
}

type ScannerTmpl struct {
	Package    string
	OptString  bool
	OptInt     bool
	OptInt64   bool
	OptFloat64 bool
}

func GetScanner(sT *ScannerTmpl) ([]byte, error) {
	var b bytes.Buffer
	res := bufio.NewWriter(&b)

	tmpl := template.Must(template.New("scanner").Parse(scannerTemplate))

	err := tmpl.Execute(res, sT)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	err = res.Flush()
	if err != nil {
		return nil, errors.Wrap(err, " cant res.Flush")
	}
	return b.Bytes(), nil
}

func GetConvert(sT *ScannerTmpl) ([]byte, error) {
	var b bytes.Buffer
	res := bufio.NewWriter(&b)

	tmpl := template.Must(template.New("scanner").Parse(convertTemplate))

	err := tmpl.Execute(res, sT)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	err = res.Flush()
	if err != nil {
		return nil, errors.Wrap(err, " cant res.Flush")
	}
	return b.Bytes(), nil
}
