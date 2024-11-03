package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	s, err := GetScanner(&ScannerTmpl{
		Package:   "scanner",
		OptString: true,
		OptInt:    true,
		OptInt64:  true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	// os.WriteFile("result.txt", s, os.ModeExclusive|os.ModePerm)

	if string(s) != `package scanner

func (oS *OptString) Scan(src any) error {
	if src == nil {
		return nil
	}
	oS.Set = true
	return convertAssignRows(&oS.Value, src)
}


func (oInt *OptInt) Scan(src any) error {
	if src == nil {
		return nil
	}
	oInt.Set = true	
	return convertAssignRows(&oInt.Value, src)
}


func (oInt *OptInt64) Scan(src any) error {
	if src == nil {
		return nil
	}
	oInt.Set = true	
	return convertAssignRows(&oInt.Value, src)
}

` {
		t.Error("wrong")
	}
}
