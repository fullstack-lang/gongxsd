package models

import (
	"log"
	"strings"
)

type ComplexType struct {
	Name string

	WithGoIdentifier

	// analysis
	IsInlined        bool // it has been defined by the enclosing element
	EnclosingElement *Element

	ElementWithAnnotation
	ElementWithNameAttribute
	Composer

	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}

func (ct *ComplexType) Fields(stage *StageStruct) (fields string) {
	setOfGoIdentifiers := make(map[string]any)

	stMap := make(map[string]*SimpleType)
	for st := range *GetGongstructInstancesSet[SimpleType](stage) {
		stMap[st.Name] = st
	}
	ctMap := make(map[string]*ComplexType)
	for st := range *GetGongstructInstancesSet[ComplexType](stage) {
		ctMap[st.Name] = st
	}
	agMap := make(map[string]*AttributeGroup)
	for ag := range *GetGongstructInstancesSet[AttributeGroup](stage) {
		agMap[ag.Name] = ag
	}

	generateAttributes(ct.Attributes, stMap, setOfGoIdentifiers, &fields)
	for _, referencedAg := range ct.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			generateAttributes(namedAg.Attributes, stMap, setOfGoIdentifiers, &fields)
			namedAg.generateAttributes(agMap, stMap, setOfGoIdentifiers, &fields)
		}
	}

	map_Name_Elems := make(map[string]*Element)
	elems := ct.Composer.getElements(map_Name_Elems)

	for _, elem := range elems {

		computeGoIdentifier(elem.GoIdentifier, &elem.WithGoIdentifier, setOfGoIdentifiers)

		goType := generateGoTypeFromSimpleType(elem.Type, stMap)
		if goType != "" {
			fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + elem.Type +
				"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		} else {
			if ct, ok := ctMap[elem.Type]; ok {
				fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + ct.Name +
					"\n\t" + elem.GoIdentifier + " []*" + ct.GoIdentifier + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
			}
		}
	}

	// parse all fields and post fix if necessary to avoid name conflicts

	return
}

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
			}

			prefix = "xml:"
			if strings.HasPrefix(attr.Ref, prefix) {
				goType = "string"
				nameDec := strings.TrimPrefix(attr.Ref, prefix)
				name = xsdNameToGoIdentifier(nameDec)
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

		*fields += "\n\n\t// generated from attribute \"" + attr.NameXSD + "\" of type " + attr.Type +
			"\n\t" + attr.GoIdentifier + " " + goType + " " + "`" + `xml:"` + attr.NameXSD + `,attr"` + "`"
	}
}

func generateGoTypeFromSimpleType(base string, stMap map[string]*SimpleType) (goType string) {

	switch base {
	// String types
	case "xs:string", "xs:normalizedString", "xs:token", "xs:language", "xs:Name", "xs:NCName",
		"xs:NMTOKEN", "xs:NMTOKENS", "xs:ID", "xs:IDREF", "xs:IDREFS", "xs:ENTITY", "xs:ENTITIES",
		"xs:QName", "xs:anyURI", "xs:NOTATION":
		goType = "string"

	// Numeric types
	case "xs:integer", "xs:nonPositiveInteger", "xs:negativeInteger", "xs:long",
		"xs:int", "xs:short", "xs:byte", "xs:nonNegativeInteger", "xs:unsignedLong", "xs:unsignedInt",
		"xs:unsignedShort", "xs:unsignedByte", "xs:positiveInteger":
		goType = "int"

		// Numeric types
	case "xs:decimal", "xs:double", "xs:float":
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
