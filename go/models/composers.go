package models

type Composer struct {
	Sequence *Sequence `xml:"sequence"`
	All      *All      `xml:"all"`
	Choice   *Choice   `xml:"choice"`
}

func (composer *Composer) getElements() (elems []*Element) {

	if composer.Sequence != nil {
		for _, e := range composer.Sequence.Elements {
			elems = append(elems, e)
		}
		elems = append(elems, composer.Sequence.Composer.getElements()...)
	}
	if composer.Choice != nil {
		for _, e := range composer.Choice.Elements {
			elems = append(elems, e)
		}
		elems = append(elems, composer.Choice.Composer.getElements()...)
	}
	if composer.All != nil {
		for _, e := range composer.All.Elements {
			elems = append(elems, e)
		}
		elems = append(elems, composer.All.Composer.getElements()...)
	}

	return
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
