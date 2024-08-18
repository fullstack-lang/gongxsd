package models

import "encoding/xml"

var Order int

type Element struct {
	Name string

	// Order is the order at wich the element was unmarshalled in the xsd
	// It is important to preserve the order output that is defined in the xsd
	Order int `xml:"-"`

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

func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	Order += 1

	type Alias Element
	aux := (*Alias)(e)
	err := d.DecodeElement(aux, &start)

	return err
}
