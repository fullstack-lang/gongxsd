package models

import "log"

type ComplexType struct {
	Name string

	// analysis
	IsInlined        bool // it has been defined by the enclosing element
	EnclosingElement *Element

	ElementWithAnnotation
	ElementWithNameAttribute
	Sequence        *Sequence         `xml:"sequence"`
	All             *All              `xml:"all"`
	Choice          *Choice           `xml:"choice"`
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
			switch st.Restriction.Base {

			case "xs:string":
				fields += "\n\n\t// generated from element " + elem.NameXSD + " with type simple type " + st.NameXSD +
					"\n\t" + elem.GoIdentifier + " string " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
			case "xs:integer":
				fields += "\n\n\t// generated from element " + elem.NameXSD + " with type simple type " + st.NameXSD +
					"\n\t" + elem.GoIdentifier + " int " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
			}
		}
	}

	return
}
