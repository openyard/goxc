package goxc_test

import (
	"testing"

	"github.com/openyard/goxc"
)

func TestXMLDSig(t *testing.T) {
	targetNamespace := "http://www.w3.org/2000/09/xmldsig#"
	result := goxc.PackageName(targetNamespace)
	if "xmldsig" != result {
		t.Errorf("expected 'xmldsig' got '%s'", result)
	}
}

func TestH004(t *testing.T) {
	targetNamespace := "urn:org:ebics:H004"
	result := goxc.PackageName(targetNamespace)
	if "h004" != result {
		t.Errorf("expected 'h004' got '%s'", result)
	}
}
