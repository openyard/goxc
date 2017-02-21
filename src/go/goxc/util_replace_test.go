package goxc_test

import (
	"go/goxc"
	"testing"
)

func TestWithTargetPrefix(t *testing.T) {
	n := goxc.Replace("foo", "foo:Bar")
	if "Bar" != n {
		t.Errorf("expected 'Bar' got '%s'", n)
	}
}

func TestWithOtherPrefix(t *testing.T) {
	n := goxc.Replace("foo", "bar:Bar")
	if "bar.Bar" != n {
		t.Errorf("expected 'bar.Bar' got '%s'", n)
	}
}

func TestWithBaseTypePrefix(t *testing.T) {
	n := goxc.Replace("foo", "w3c.Bar")
	if "w3c.Bar" != n {
		t.Errorf("expected 'bar.Bar' got '%s'", n)
	}
}

func TestWithWithoutPrefix(t *testing.T) {
	n := goxc.Replace("", "boolean")
	if "w3c.Boolean" != n {
		t.Errorf("expected 'w3c.Boolean' got '%s'", n)
	}
	n = goxc.Replace("", "string")
	if "w3c.String" != n {
		t.Errorf("expected 'w3c.String' got '%s'", n)
	}
}