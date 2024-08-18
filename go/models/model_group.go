package models

import (
	"cmp"
	"log"
	"slices"
)

// ModelGroup is a construct that allows for
// the specification of complex structures within an XML document.
// Model groups define how elements within an XML schema are composed and ordered.
// There are three main types of model groups in XSD
type ModelGroup struct {
	OuterElementName string

	Sequences []*Sequence `xml:"sequence"`
	Alls      []*All      `xml:"all"`
	Choices   []*Choice   `xml:"choice"`
	Groups    []*Group    `xml:"group"`

	Elements []*Element `xml:"element"`
}

func (modelGroup *ModelGroup) getElements(groupMap map[string]*Group, map_Name_Elems map[string]*Element) (elems []*Element) {

	for _, gRef := range modelGroup.Groups {

		if gRef.Ref != "" {
			// log.Println("Processing group", gRef.Name, gRef.Ref)
			if namedGroup, ok := groupMap[gRef.Ref]; ok {
				elems = append(elems, namedGroup.getElements(
					groupMap, map_Name_Elems)...)
			}
		} else {
			for _, e := range gRef.Elements {
				elems = append(elems, gRef.ModelGroup.getElements(groupMap, map_Name_Elems)...)
				if _, ok := map_Name_Elems[e.Name]; ok {
					continue
				}
				map_Name_Elems[e.Name] = e
				elems = append(elems, e)
			}
		}
	}
	for _, s := range modelGroup.Sequences {
		elems = append(elems, s.getElements(groupMap, map_Name_Elems)...)
		for _, e := range s.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
	}
	for _, c := range modelGroup.Choices {
		elems = append(elems, c.getElements(groupMap, map_Name_Elems)...)
		for _, e := range c.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
	}
	for _, a := range modelGroup.Alls {
		elems = append(elems, a.ModelGroup.getElements(groupMap, map_Name_Elems)...)
		for _, e := range a.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			elems = append(elems, e)
		}
	}

	// reoder elements according to their rank
	slices.SortFunc(elems, func(a, b *Element) int {
		return cmp.Compare(a.Order, b.Order)
	})

	return
}

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

type Choice struct {
	Name string
	Annotated
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`

	ModelGroup
}

func (composer *ModelGroup) generateElements(
	map_Name_Elems map[string]*Element,
	stMap map[string]*SimpleType,
	ctMap map[string]*ComplexType,
	groupMap map[string]*Group,
	setOfGoIdentifiers map[string]any,
	fields *string,
) {
	elems := composer.getElements(groupMap, map_Name_Elems)

	for _, elem := range elems {

		computeGoIdentifier(elem.GoIdentifier, &elem.WithGoIdentifier, setOfGoIdentifiers)

		// an element can be of 3 types:
		// 1. a simple type
		// 2. a named complex type
		// 3. an anonmous complex type
		goType := generateGoTypeFromSimpleType(elem.Type, stMap)
		if goType != "" {
			// 1. a simple type
			*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + elem.Type +
				"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
		} else {
			if elem.Type != "" {
				if ct, ok := ctMap[elem.Type]; ok {
					*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + ct.Name +
						"\n\t" + elem.GoIdentifier + " []*" + ct.GoIdentifier + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
				} else {
					log.Println("element", elem.NameXSD, "unkown type", elem.Type)
				}
			} else {
				if elem.ComplexType == nil {
					log.Println("element", elem.NameXSD, "should have an anonymous complex type", elem.Type)
				} else {
					ct := elem.ComplexType
					*fields += "\n\n\t// generated from anonymous type within outer element \"" + elem.NameXSD + "\" of type " + ct.Name +
						"\n\t" + elem.GoIdentifier + " []*" + ct.GoIdentifier + " " + "`" + `xml:"` + elem.NameXSD + `"` + "`"
				}
			}

		}
	}
}
