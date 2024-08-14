package models

type Extension struct {
	Name string
	ModelGroup

	Base string `xml:"base,attr"`

	Attributes []*Attribute `xml:"attribute"`
}
