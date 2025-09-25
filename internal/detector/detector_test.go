package detector

import (
	"testing"
)

func TestScannerTmpl(t *testing.T) {
	tm, err := ScannerTmpl("fixture/oas_schemas_gen.go")
	if err != nil {
		t.Error(err)
		return
	}
	if tm.Package != "ogenapi" {
		t.Error("wrong package name")
	}
	if tm.OptString != true {
		t.Error("OptString should be true")
	}
	if tm.OptInt != true {
		t.Error("OptInt should be true")
	}
	if tm.OptInt64 {
		t.Error("OptInt64 should be false")
	}
	if !tm.OptFloat64 {
		t.Error("OptFloat64 should be true")
	}
	if !tm.OptNilInt {
		t.Error("OptNilInt should be true")
	}
}
