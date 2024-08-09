package models

type Composer struct {
	Sequence *Sequence `xml:"sequence"`
	All      *All      `xml:"all"`
	Choice   *Choice   `xml:"choice"`
}

type Sequence struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
	Composer
}

type All struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
	Composer
}

type Choice struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Groups    []*Group   `xml:"group"`
	Composer
}
