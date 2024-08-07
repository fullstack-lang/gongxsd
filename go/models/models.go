package models

type Schema struct {
	Name         string
	Elements     []*Element     `xml:"element"`
	SimpleTypes  []*SimpleType  `xml:"simpleType"`
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type Element struct {
	Name        string
	NameXSD     string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
}

type SimpleType struct {
	Name        string
	NameXSD     string       `xml:"name,attr"`
	Restriction *Restriction `xml:"restriction"`
}

type Restriction struct {
	Name         string
	Base         string         `xml:"base,attr"`
	Enumerations []*Enumeration `xml:"enumeration"`
	MinInclusive *MinInclusive  `xml:"minInclusive"`
	MaxInclusive *MaxInclusive  `xml:"maxInclusive"`
}

type Enumeration struct {
	Name  string
	Value string `xml:"value,attr"`
}

type MinInclusive struct {
	Name  string
	Value string `xml:"value,attr"`
}

type MaxInclusive struct {
	Name  string
	Value string `xml:"value,attr"`
}

type ComplexType struct {
	Name     string
	NameXSD  string    `xml:"name,attr"`
	Sequence *Sequence `xml:"sequence"`
}

type Sequence struct {
	Name     string
	Elements []*Element `xml:"element"`
}
