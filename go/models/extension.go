package models

type Extension struct {
	Name string
	ModelGroup

	Attributes []*Attribute `xml:"attribute"`
}
