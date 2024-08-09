package models

import "log"

type ComplexType struct {
	Name string

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

	stMap := make(map[string]*SimpleType)

	for st := range *GetGongstructInstancesSet[SimpleType](stage) {
		stMap[st.Name] = st
	}

	if ct.Sequence == nil {
		return
	}

	seq := ct.Sequence

	for _, elem := range seq.Elements {
		switch elem.Type {
		case "xs:string":
			fields += "\n\n\t// generated from element " + elem.NameXSD + " with type xs:string" +
				"\n\t" + elem.GoIdentifier + " string " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		case "xs:integer":
			fields += "\n\n\t// generated from element " + elem.NameXSD + " with type xs:integer" +
				"\n\t" + elem.GoIdentifier + " int " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		}

		// element type is a simple type
		if st, ok := stMap[elem.Type]; ok {
			if st.Restriction == nil {
				log.Fatalln("simple type without restriction", st.NameXSD)
			}

			goType := generateGoType(st.Restriction.Base, stMap)

			fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" with type simple type " + st.NameXSD +
				"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		}
	}

	return
}

func generateGoType(base string, stMap map[string]*SimpleType) (goType string) {

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
			return generateGoType(st.Restriction.Base, stMap)
		} else {
			log.Fatalln("unknown restriction base", base)
		}
	}
	return
}
