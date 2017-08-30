package goxc_test

import (
	"go/goxc"
	"testing"
)

var namespaces = map[string]string{
	"xs": "http://www.w3.org/2001/XMLSchema",
	"foo": "http://foo",
	"bar": "http://bar",
}

func TestWithTargetPrefix(t *testing.T) {
	n := goxc.Replace("foo", "foo:Bar", namespaces)
	if "Bar" != n {
		t.Errorf("expected 'Bar' got '%s'", n)
	}
}

func TestWithOtherPrefix(t *testing.T) {
	n := goxc.Replace("foo", "bar:Bar", namespaces)
	if "bar.Bar" != n {
		t.Errorf("expected 'bar.Bar' got '%s'", n)
	}
}

func TestWithBaseTypePrefix(t *testing.T) {
	n := goxc.Replace("foo", "w3c.Bar", namespaces)
	if "w3c.Bar" != n {
		t.Errorf("expected 'bar.Bar' got '%s'", n)
	}
}

func TestWithWithoutPrefix(t *testing.T) {
	n := goxc.Replace("", "boolean", namespaces)
	if "w3c.Boolean" != n {
		t.Errorf("expected 'w3c.Boolean' got '%s'", n)
	}
	n = goxc.Replace("", "string", namespaces)
	if "w3c.String" != n {
		t.Errorf("expected 'w3c.String' got '%s'", n)
	}
}

func TestWithWithDifferentW3CPrefix(t *testing.T) {
	n := goxc.Replace("", "xs:boolean", namespaces)
	if "w3c.Boolean" != n {
		t.Errorf("expected 'w3c.Boolean' got '%s'", n)
	}
	n = goxc.Replace("", "xs:string", namespaces)
	if "w3c.String" != n {
		t.Errorf("expected 'w3c.String' got '%s'", n)
	}
}