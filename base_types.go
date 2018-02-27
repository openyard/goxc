package goxc

var baseTypes []*BaseType = make([]*BaseType, 0)

func init() {
	baseTypes = append(baseTypes, &BaseType{Name: "String", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "NCName", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "NormalizedString", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Token", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "AnyURI", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "AnySimpleType", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Decimal", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Base64Binary", Type: "[]byte"})
	baseTypes = append(baseTypes, &BaseType{Name: "HexBinary", Type: "[]byte"})
	baseTypes = append(baseTypes, &BaseType{Name: "ID", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Date", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "DateTime", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Language", Type: "string"})
	baseTypes = append(baseTypes, &BaseType{Name: "Integer", Type: "int"})
	baseTypes = append(baseTypes, &BaseType{Name: "NonNegativeInteger", Type: "uint"})
	baseTypes = append(baseTypes, &BaseType{Name: "PositiveInteger", Type: "uint"})
	baseTypes = append(baseTypes, &BaseType{Name: "Boolean", Type: "bool"})
}
