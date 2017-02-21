package goxc_test

import (
	"encoding/xml"
	"testing"

	"go/goxc"
)

func TestEnumerationWithBaseType(t *testing.T) {
	var s goxc.SimpleType
	withTempDir(func(packageName string) {
		xml.Unmarshal(simpleTypeWithEnumeration(), &s)
		s.PackageName = packageName
		s.Generate("foo", Namespaces())
	}, "foo")
}

func TestEnumerationWithOwnType(t *testing.T) {
	var s goxc.SimpleType
	withTempDir(func(packageName string) {
		xml.Unmarshal(simpleTypeWithRestrictionBase(), &s)
		s.PackageName = packageName
		s.Generate("foo", Namespaces())
	}, "foo")
}


func TestEnumerationWithBaseTypeAndParent(t *testing.T) {
	var s goxc.SimpleType
	withTempDir(func(packageName string) {
		xml.Unmarshal(simpleTypeWithEnumeration(), &s)
		s.PackageName = packageName
		s.Parent = "P"
		s.Generate("foo", Namespaces())
	}, "foo")
}

func TestEnumerationWithOwnTypeAndParent(t *testing.T) {
	var s goxc.SimpleType
	withTempDir(func(packageName string) {
		xml.Unmarshal(simpleTypeWithRestrictionBase(), &s)
		s.PackageName = packageName
		s.Parent = "P"
		s.Generate("foo", Namespaces())
	}, "foo")
}

func simpleTypeWithEnumeration() []byte {
	return []byte(
		`<simpleType name="E">
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
	`)
}

func simpleTypeWithRestrictionBase() []byte {
	return []byte(
		`<simpleType name="E">
			<annotation>
				<documentation xml:lang="de">Dokumentation</documentation>
			</annotation>
			<restriction base="foo:C">
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
	`)
}