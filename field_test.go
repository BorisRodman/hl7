package hl7_test

import (
	"testing"

	"github.com/freemed/hl7"
)

func TestFieldParse(t *testing.T) {
	val := []rune("520 51st Street^^Denver^CO^80020^USA")
	seps := NewDelimeters()
	fld := &Field{Value: val}
	fld.parse(seps)
	if len(fld.Components) != 6 {
		t.Errorf("Expected 6 components got %d\n", len(fld.Components))
	}
}

func TestFieldSet(t *testing.T) {
	seps := hl7.NewDelimeters()
	fld := &hl7.Field{}
	loc := "ZZZ.1.10"
	l := hl7.NewLocation(loc)
	err := fld.Set(l, "TEST", seps)
	if err != nil {
		t.Error(err)
	}
	if len(fld.Components) != 11 {
		t.Fatalf("Expected 11 got %d\n", len(fld.Components))
	}
	if string(fld.Components[10].SubComponents[0].Value) != "TEST" {
		t.Errorf("Expected TEST got %s\n", string(fld.Components[10].SubComponents[0].Value))
	}
}
