package scanner

import (
	"testing"
)

func TestOptString(t *testing.T) {
	var value OptString
	var v string = "test_value"

	err := value.Scan(v)
	if err != nil {
		t.Error("cant scan", err)
		return
	}
	if value.Set != true && value.Value == v {
		t.Error("wrong value")
		return
	}

	var value2 OptString
	err = value2.Scan(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if value2.Set != false && value2.Value != "" {
		t.Error("wrong result")
	}
}

func TestOptInt(t *testing.T) {
	var value OptInt
	var v string = "1"

	err := value.Scan(v)
	if err != nil {
		t.Error("cant scan", err)
		return
	}
	if value.Set != true && value.Value != 1 {
		t.Error("wrong value")
		return
	}

	var value2 OptInt
	err = value2.Scan(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if value2.Set != false && value2.Value != 0 {
		t.Error("wrong result")
	}
}

func TestOptInt64(t *testing.T) {
	var value OptInt64
	var v string = "1"

	err := value.Scan(v)
	if err != nil {
		t.Error("cant scan", err)
		return
	}
	if value.Set != true && value.Value != 1 {
		t.Error("wrong value")
		return
	}

	var value2 OptInt64
	err = value2.Scan(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if value2.Set != false && value2.Value != 0 {
		t.Error("wrong result")
	}
}
