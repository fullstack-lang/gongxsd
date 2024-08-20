package models

import (
	"encoding/xml"
)

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`

	Xmlns           Xmlns              `xml:"-"`
	importedModules map[string]*Schema `xml:"-"`

	Annotated
	Elements        []*Element        `xml:"element"`
	SimpleTypes     []*SimpleType     `xml:"simpleType"`
	ComplexTypes    []*ComplexType    `xml:"complexType"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	Groups          []*Group          `xml:"group"`
}

func (sch *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	sch.Xmlns = parseXmlns(start)

	type s Schema
	ss := (*s)(sch)
	return d.DecodeElement(ss, &start)
}
