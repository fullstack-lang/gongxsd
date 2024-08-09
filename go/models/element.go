package models

type Element struct {
	Name string

	// analysis

	// HasNameConflict
	//
	// In XML Schema Definition (XSD), it is technically possible to have a
	// named complex type and an element with the same name or elements with
	// the same (see musicxml.xsd for instance),
	HasNameConflict bool
	GoIdentifier    string

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
