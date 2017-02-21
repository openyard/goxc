package goxc_test

import (
	"encoding/xml"
	"testing"

	"go/goxc"
)

func TestComplexTypeWithSimpleContentContainingAttributeWithEnumeration(t *testing.T) {
	var c goxc.ComplexType
	withTempDir(func(packageName string) {
		xml.Unmarshal(example(), &c)
		c.PackageName = packageName
		c.Parent = "P"
		c.Name = "C"
		c.Generate("foo", Namespaces())
	}, "foo")
}

func example() []byte {
	return []byte(`
		<complexType>
			<simpleContent>
				<extension base="string">
					<attribute name="a" use="required">
						<annotation>
							<documentation xml:lang="de">Beschreibungstyp.</documentation>
						</annotation>
						<simpleType>
							<restriction base="token">
								<enumeration value="Purpose">
									<annotation>
										<documentation xml:lang="de">Verwendungszweck</documentation>
									</annotation>
								</enumeration>
								<enumeration value="Details">
									<annotation>
										<documentation xml:lang="de">Auftragsdetails</documentation>
									</annotation>
								</enumeration>
								<enumeration value="Comment">
									<annotation>
										<documentation xml:lang="de">Kommentar</documentation>
									</annotation>
								</enumeration>
							</restriction>
						</simpleType>
					</attribute>
				</extension>
			</simpleContent>
		</complexType>

	`)
}