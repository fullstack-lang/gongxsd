package models

type Choice struct {
	Name string
	Annotated
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`

	ModelGroup
}
