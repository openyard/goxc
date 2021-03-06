package goxc_test

import (
	"testing"

	"github.com/openyard/goxc"
)

func TestAppend(t *testing.T) {
	schemaImports := make(map[string]string, 0)
	schemaImports["foo"] = "bar"
	schemaImports["bar"] = "foo"
	imports := make([]*goxc.Import, 0)
	imports = goxc.Append(imports, "foo.Bar", schemaImports)
	imports = goxc.Append(imports, "bar.Bar", schemaImports)
	if len(imports) == 0 {
		t.Error("expected imports > 0")
	}
	sample := imports[0]
	if sample.Alias != "foo" {
		t.Errorf("expected 'foo' got %s", sample.Alias)
	}
	if sample.Package != "bar" {
		t.Errorf("expected 'bar' got %s", sample.Package)
	}
	sample = imports[1]
	if sample.Alias != "bar" {
		t.Errorf("expected 'bar' got %s", sample.Alias)
	}
	if sample.Package != "foo" {
		t.Errorf("expected 'foo' got %s", sample.Package)
	}
}

func TestAppendTwice(t *testing.T) {
	schemaImports := make(map[string]string, 0)
	schemaImports["foo"] = "bar"
	imports := make([]*goxc.Import, 0)
	imports = goxc.Append(imports, "foo.Bar", schemaImports)
	imports = goxc.Append(imports, "foo.Bar", schemaImports)

	if len(imports) == 2 {
		t.Fail()
	}

}
