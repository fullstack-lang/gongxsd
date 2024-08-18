package models

import (
	"log"
	"strings"
)

func generateAttributes(
	attrs []*Attribute,
	stMap map[string]*SimpleType,
	setOfGoIdentifiers map[string]any,
	fields *string) {
	for _, attr := range attrs {
		goType := generateGoTypeFromSimpleType(attr.Type, stMap)

		var name string

		if goType == "" {

			prefix := "xlink:"
			if strings.HasPrefix(attr.Ref, prefix) {
				goType = "string"
				nameDec := strings.TrimPrefix(attr.Ref, prefix)
				name = xsdNameToGoIdentifier(nameDec)

				// overide the name of the attr
				attr.NameXSD = "http://www.w3.org/1999/xlink " + nameDec
			}

			prefix = "xml:"
			if strings.HasPrefix(attr.Ref, prefix) {
				goType = "string"
				nameDec := strings.TrimPrefix(attr.Ref, prefix)
				name = xsdNameToGoIdentifier(nameDec)

				// overide the name of the attr
				attr.NameXSD = "http://www.w3.org/XML/1998/namespace " + nameDec
			}
		} else {
			name = xsdNameToGoIdentifier(attr.Name)
		}

		if goType == "" {
			log.Fatalln("No resolution of attribute type", attr.Type)
		}

		switch name {
		case "Name":
			name = "NameXSD"
		}

		computeGoIdentifier(name, &attr.WithGoIdentifier, setOfGoIdentifiers)

		*fields += "\n\n\t// generated from attribute \"" + attr.NameXSD +
			"\n\t" + attr.GoIdentifier + " " + goType + " " + "`" + `xml:"` + attr.NameXSD + `,attr,omitempty"` + "`"
	}
}

func generateGoTypeFromSimpleType(base string, stMap map[string]*SimpleType) (goType string) {

	switch base {
	// String types
	case "xs:string", "xs:normalizedString", "xs:token", "xs:language", "xs:Name", "xs:NCName",
		"xs:NMTOKEN", "xs:NMTOKENS", "xs:ID", "xs:IDREF", "xs:IDREFS", "xs:ENTITY", "xs:ENTITIES",
		"xs:QName", "xs:anyURI", "xs:NOTATION", "xs:decimal":
		goType = "string"

	// Numeric types
	case "xs:integer", "xs:nonPositiveInteger", "xs:negativeInteger", "xs:long",
		"xs:int", "xs:short", "xs:byte", "xs:nonNegativeInteger", "xs:unsignedLong", "xs:unsignedInt",
		"xs:unsignedShort", "xs:unsignedByte", "xs:positiveInteger":
		goType = "int"

		// Numeric types
	case "xs:double", "xs:float":
		goType = "float64"

	// Boolean type
	case "xs:boolean":
		goType = "bool"

	}

	// a base can refer a simple type
	if goType == "" {

		if st, ok := stMap[base]; ok {
			if st.Restriction != nil {
				return generateGoTypeFromSimpleType(st.Restriction.Base, stMap)
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
