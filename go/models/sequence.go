package models

type Sequence struct {
	Name string
	Annotated
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`

	ModelGroup
}

type All struct {
	Name string
	Annotated
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`

	ModelGroup
}
