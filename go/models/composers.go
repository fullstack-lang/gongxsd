package models

type Composer struct {
	Sequence *Sequence `xml:"sequence"`
	All      *All      `xml:"all"`
	Choice   *Choice   `xml:"choice"`
}

func (composer *Composer) getElements(map_Name_Elems map[string]*Element) (elems []*Element) {

	if composer.Sequence != nil {
		for _, e := range composer.Sequence.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, composer.Sequence.Composer.getElements(map_Name_Elems)...)
	}
	if composer.Choice != nil {
		for _, e := range composer.Choice.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, composer.Choice.Composer.getElements(map_Name_Elems)...)
	}
	if composer.All != nil {
		for _, e := range composer.All.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, composer.All.Composer.getElements(map_Name_Elems)...)
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
