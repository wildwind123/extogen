package detector

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/go-faster/errors"
	"github.com/wildwind123/extogen/internal/builder"
)

func ScannerTmpl(ogenSchemaPath string) (*builder.ScannerTmpl, error) {

	b, err := os.ReadFile(ogenSchemaPath)
	if err != nil {
		return nil, errors.Wrap(err, "cant os.ReadFile")
	}

	packageName, err := getPackageNameFromSource(string(b))
	if err != nil {
		return nil, errors.Wrap(err, "cant getPackageNameFromSource")
	}

	tm := builder.ScannerTmpl{
		Package:    packageName,
		OptString:  bytes.Contains(b, []byte(` OptString `)),
		OptInt:     bytes.Contains(b, []byte(` OptInt `)),
		OptInt64:   bytes.Contains(b, []byte(` OptInt64 `)),
		OptFloat64: bytes.Contains(b, []byte(` OptFloat64 `)),
		OptBool:    bytes.Contains(b, []byte(` OptBool `)),
		// nil
		OptNilString:  bytes.Contains(b, []byte(` OptNilString `)),
		OptNilInt:     bytes.Contains(b, []byte(` OptNilInt `)),
		OptNilInt64:   bytes.Contains(b, []byte(` OptNilInt64 `)),
		OptNilFloat64: bytes.Contains(b, []byte(` OptNilFloat64 `)),
		OptNilBool:    bytes.Contains(b, []byte(` OptNilBool `)),
	}

	return &tm, nil
}

func getPackageNameFromSource(src string) (string, error) {
	// Define a regex pattern to match the 'package' keyword followed by the package name
	re := regexp.MustCompile(`(?m)^\s*package\s+([a-zA-Z0-9_]+)\s*$`)

	// Find the match
	match := re.FindStringSubmatch(src)

	// Check if a match was found
	if len(match) < 2 {
		return "", fmt.Errorf("package declaration not found")
	}

	// Return the captured package name
	return match[1], nil
}
