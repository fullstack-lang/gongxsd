package models

type ElementWithNameAttribute struct {
	NameXSD string `xml:"name,attr"`
}

type ElementWithTypeAttribute struct {
	Type string `xml:"type,attr"`
}

type ElementWithValueAttribute struct {
	Value string `xml:"value,attr"`
}

type ElementWithAnnotation struct {
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	Name           string
	Documentations []*Documentation `xml:"documentation"`
}

type Documentation struct {
	Name   string
	Text   string `xml:",chardata"`
	Source string `xml:"source,attr"`
	Lang   string `xml:"lang,attr"`
}

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`
	ElementWithAnnotation
	Elements        []*Element        `xml:"element"`
	SimpleTypes     []*SimpleType     `xml:"simpleType"`
	ComplexTypes    []*ComplexType    `xml:"complexType"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	Groups          []*Group          `xml:"group"`
}

type Group struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	Ref string `xml:"ref,attr"`
}

type SimpleType struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	Restriction *Restriction `xml:"restriction"`
}

type Restriction struct {
	Name string
	ElementWithAnnotation
	Base         string         `xml:"base,attr"`
	Enumerations []*Enumeration `xml:"enumeration"`
	MinInclusive *MinInclusive  `xml:"minInclusive"`
	MaxInclusive *MaxInclusive  `xml:"maxInclusive"`
	Pattern      *Pattern       `xml:"pattern"`
	WhiteSpace   *WhiteSpace    `xml:"whiteSpace"`
	MinLength    *MinLength     `xml:"minLength"`
	MaxLength    *MaxLength     `xml:"maxLength"`
	Length       *Length        `xml:"length"`
	TotalDigit   *TotalDigit    `xml:"totalDigits"`
}

type Attribute struct {
	Name string
	ElementWithNameAttribute
	ElementWithTypeAttribute
	ElementWithAnnotation

	Default         string `xml:"default,attr"`
	Use             string `xml:"use,attr"`
	Form            string `xml:"form,attr"`
	Fixed           string `xml:"fixed,attr"`
	Ref             string `xml:"ref,attr"`
	TargetNamespace string `xml:"targetNamespace,attr"`
	SimpleType      string `xml:"simpleType,attr"`
	IDXSD           string `xml:"id,attr"`
}
