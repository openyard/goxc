package goxc_test

import (
	"testing"

	"github.com/openyard/goxc"
)

func TestStringIsBaseType(t *testing.T) {
	if !goxc.IsBaseType("string") {
		t.Errorf("expected 'string' is base type")
	}
}

func TestIDIsBaseType(t *testing.T) {
	if !goxc.IsBaseType("ID") {
		t.Errorf("expected 'ID' is base type")
	}
}

func TestBooleanIsBaseType(t *testing.T) {
	if !goxc.IsBaseType("boolean") {
		t.Errorf("expected 'boolean' is base type")
	}
}

func TestWithPrefix(t *testing.T) {
	if goxc.IsBaseType("foo:Bar") {
		t.Errorf("expected 'foo:Bar' isn't base type")
	}
}
