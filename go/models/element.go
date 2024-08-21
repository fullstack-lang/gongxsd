package models

import "encoding/xml"

var Order int
var Depth int

type Element struct {
	Name string

	ParticleAbstract

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

	// OuterParticle is the particle that reference the element
	// this is used in the factoring process of non normalized xsd
	OuterParticle     Particle `xml:"-"`
	IsDuplicatedInXSD bool
}

func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	e.Depth = Depth
	Order = Order + 1

	type Alias Element
	aux := (*Alias)(e)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
