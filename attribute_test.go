package goxc_test

import (
	"encoding/xml"
	"testing"

	"github.com/openyard/goxc"
)

func TestAttributeOutput(t *testing.T) {
	var a goxc.Attribute
	withTempDir(func(packageName string) {
		xml.Unmarshal(attributeWithSimpleTypeWithEnumeration(), &a)
		a.PackageName = packageName
		a.Generate("foo", Namespaces())
	}, "foo")
}

func attributeWithSimpleTypeWithEnumeration() []byte {
	return []byte(
		`<attribute name="a" use="required">
			<annotation>
				<documentation xml:lang="de">Beschreibungstyp.</documentation>
			</annotation>
			<simpleType name="B">
				<annotation>
					<documentation xml:lang="de">Dokumentation</documentation>
				</annotation>
				<restriction base="token">
					<enumeration value="One">
						<annotation>
							<documentation xml:lang="de">Wert 1</documentation>
						</annotation>
					</enumeration>
					<enumeration value="Two">
						<annotation>
							<documentation xml:lang="de">Wert 2</documentation>
						</annotation>
					</enumeration>
					<enumeration value="Three">
						<annotation>
							<documentation xml:lang="de">Wert 3</documentation>
						</annotation>
					</enumeration>
				</restriction>
			</simpleType>
		</attribute>
		`)
}
