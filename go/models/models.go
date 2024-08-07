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
	Source string `xml:"source,attr"`
	Lang   string `xml:"lang,attr"`
}

type Schema struct {
	Name string
	ElementWithAnnotation
	Elements     []*Element     `xml:"element"`
	SimpleTypes  []*SimpleType  `xml:"simpleType"`
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type Element struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	ElementWithTypeAttribute
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
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
}

type ComplexType struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	Sequence *Sequence `xml:"sequence"`
}

type Sequence struct {
	Name string
	ElementWithAnnotation
	Elements []*Element `xml:"element"`
}
