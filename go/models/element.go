package models

type Element struct {
	Name string

	// analysis
	WithGoIdentifier

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
	Groups      []*Group     `xml:"group"`
}
