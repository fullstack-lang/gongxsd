package models

type Composer struct {
	Sequences []*Sequence `xml:"sequence"`
	Alls      []*All      `xml:"all"`
	Choices   []*Choice   `xml:"choice"`
	Groups    []*Group    `xml:"group"`
}

func (composer *Composer) getElements(map_Name_Elems map[string]*Element) (elems []*Element) {

	for _, s := range composer.Sequences {
		for _, e := range s.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, s.getElements(map_Name_Elems)...)
	}
	for _, c := range composer.Choices {
		for _, e := range c.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, c.getElements(map_Name_Elems)...)
	}
	for _, a := range composer.Alls {
		for _, e := range a.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
		elems = append(elems, a.Composer.getElements(map_Name_Elems)...)
	}

	return
}

type Sequence struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Composer
}

type All struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Composer
}

type Choice struct {
	Name string
	ElementWithAnnotation
	MinOccurs string     `xml:"minOccurs,attr"`
	MaxOccurs string     `xml:"maxOccurs,attr"`
	Elements  []*Element `xml:"element"`
	Composer
}

func (composer *Composer) generateElements(
	map_Name_Elems map[string]*Element,
	stMap map[string]*SimpleType,
	ctMap map[string]*ComplexType,
	setOfGoIdentifiers map[string]any,
	fields *string,
) {
	elems := composer.getElements(map_Name_Elems)

	for _, elem := range elems {

		computeGoIdentifier(elem.GoIdentifier, &elem.WithGoIdentifier, setOfGoIdentifiers)

		goType := generateGoTypeFromSimpleType(elem.Type, stMap)
		if goType != "" {
			*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + elem.Type +
				"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		} else {
			if ct, ok := ctMap[elem.Type]; ok {
				*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + ct.Name +
					"\n\t" + elem.GoIdentifier + " []*" + ct.GoIdentifier + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
			}
		}
	}
}
