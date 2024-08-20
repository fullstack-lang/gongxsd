package models

func generateGoTypeFromBase(base string, stMap map[string]*SimpleType) (goType string) {

	switch base {
	// String types
	case
		"xs:dateTime",
		"xs:string", "xs:normalizedString", "xs:token", "xs:language", "xs:Name", "xs:NCName",
		"xs:NMTOKEN", "xs:NMTOKENS", "xs:ID", "xs:IDREF", "xs:IDREFS", "xs:ENTITY", "xs:ENTITIES",
		"xs:QName", "xs:anyURI", "xs:NOTATION", "xs:decimal",

		// ugly hack

		"xsd:dateTime",
		"xsd:string", "xsd:normalizedString", "xsd:token", "xsd:language", "xsd:Name", "xsd:NCName",
		"xsd:NMTOKEN", "xsd:NMTOKENS", "xsd:ID", "xsd:IDREF", "xsd:IDREFS", "xsd:ENTITY", "xsd:ENTITIES",
		"xsd:QName", "xsd:anyURI", "xsd:NOTATION", "xsd:decimal":
		goType = "string"

	// Numeric types
	case "xs:integer", "xs:nonPositiveInteger", "xs:negativeInteger", "xs:long",
		"xs:int", "xs:short", "xs:byte", "xs:nonNegativeInteger", "xs:unsignedLong", "xs:unsignedInt",
		"xs:unsignedShort", "xs:unsignedByte", "xs:positiveInteger",

		"xsd:integer", "xsd:nonPositiveInteger", "xsd:negativeInteger", "xsd:long",
		"xsd:int", "xsd:short", "xsd:byte", "xsd:nonNegativeInteger", "xsd:unsignedLong", "xsd:unsignedInt",
		"xsd:unsignedShort", "xsd:unsignedByte", "xsd:positiveInteger":
		goType = "int"

		// Numeric types
	case "xs:double", "xs:float",

		"xsd:double", "xsd:float":
		goType = "float64"

	// Boolean type
	case "xs:boolean",

		"xsd:boolean":
		goType = "bool"

	}

	// a base can refer a simple type
	if goType == "" {

		if st, ok := stMap[base]; ok {
			if st.Restriction != nil {
				return generateGoTypeFromBase(st.Restriction.Base, stMap)
			} else {
				//
				// the union of type can be a go type int, float64, string or boolean
				// on simplify to string
				return "string"
				// types := strings.Split(st.Union.MemberTypes, " ")
				// return generateGoTypeFromSimpleType(types[0], stMap)
			}
		}
		//
		// type="empty": This indicates that the element is expected to be empty
		// (i.e., it does not have any child elements or text content).
		if base == "empty" {
			return "string"
		}
	}
	return
}
