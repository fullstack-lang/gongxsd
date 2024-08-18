package models

type Element struct {
	Name string

	Line   int `xml:"-"`
	Column int `xml:"-"`

	// analysis
	WithGoIdentifier

	Annotated
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

// func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	type Alias Element
// 	aux := (*Alias)(e)
// 	return d.DecodeElement(aux, &start)
// }
