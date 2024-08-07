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

type Element struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	ElementWithTypeAttribute

	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
	Default   string `xml:"default,attr"`
	Fixed     string `xml:"fixed,attr"`
	Nillable  string `xml:"nillable,attr"`
	Ref       string `xml:"ref,attr"`
	Abstract  string `xml:"abstract,attr"`
	Form      string `xml:"form,attr"`
	Block     string `xml:"block,attr"`
	Final     string `xml:"final,attr"`

	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
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

type ComplexType struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	Sequence        *Sequence         `xml:"sequence"`
	All             *All              `xml:"all"`
	Choice          *Choice           `xml:"choice"`
	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
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

type AttributeGroup struct {
	Name string
	ElementWithNameAttribute
	ElementWithAnnotation

	AttributeGroup *AttributeGroup `xml:"attributeGroup"`

	Ref string `xml:"ref,attr"`
}

type Sequence struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
}

type All struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
}

type Choice struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
}
