package goxc_test

import (
	"testing"

	"github.com/openyard/goxc"
)

func TestNamespace(t *testing.T) {
	namespaces := map[string]string{
		"foo": "foo",
	}

	if "foo " != goxc.Namespace("foo.Bar", namespaces) {
		t.Errorf("expected 'foo' got %s", goxc.Namespace("foo.Bar", namespaces) )
	}
	if "" != goxc.Namespace("FooBar", namespaces) {
		t.Errorf("expected '' got %s", goxc.Namespace("FooBar", namespaces) )
	}
}
