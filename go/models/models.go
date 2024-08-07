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

type Schema struct {
	Name         string
	Elements     []*Element     `xml:"element"`
	SimpleTypes  []*SimpleType  `xml:"simpleType"`
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type Element struct {
	Name string
	ElementWithNameAttribute
	ElementWithTypeAttribute
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
}

type SimpleType struct {
	Name string
	ElementWithNameAttribute
	Restriction *Restriction `xml:"restriction"`
}

type Restriction struct {
	Name         string
	Base         string         `xml:"base,attr"`
	Enumerations []*Enumeration `xml:"enumeration"`
	MinInclusive *MinInclusive  `xml:"minInclusive"`
	MaxInclusive *MaxInclusive  `xml:"maxInclusive"`
	Pattern      *Pattern       `xml:"pattern"`
}

type ComplexType struct {
	Name string
	ElementWithNameAttribute
	Sequence *Sequence `xml:"sequence"`
}

type Sequence struct {
	Name     string
	Elements []*Element `xml:"element"`
}
