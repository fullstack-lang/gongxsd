package models

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
	ctMap := make(map[string]*ComplexType)
	for st := range *GetGongstructInstancesSet[ComplexType](stage) {
		ctMap[st.Name] = st
	}

	if ct.Sequence == nil {
		return
	}

	elems := ct.Composer.getElements()
	// for _, e := range elems {
	// 	fields += "\n\t// elem " + e.GoIdentifier
	// }

	for _, elem := range elems {

		goType := generateGoTypeFromSimpleType(elem.Type, stMap)
		if goType != "" {
			fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + elem.Type +
				"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		} else {
			if ct, ok := ctMap[elem.Type]; ok {
				fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + ct.Name +
					"\n\t" + elem.GoIdentifier + " []*" + xsdNameToGoIdentifier(ct.Name) + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
			}
		}
	}

	return
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
			return generateGoTypeFromSimpleType(st.Restriction.Base, stMap)
		}
	}
	return
}
