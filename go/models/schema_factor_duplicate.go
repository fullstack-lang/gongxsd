package models

import (
	"log"
	"slices"
	"sort"

	"golang.org/x/exp/maps"
)

// FactorDuplicates performs a simplification of the xsd schema.
//
// For instance, in REQIF schema, the same "ALTERNATIVE-ID" element and its
// inner anonymous complex type are repeated 24 times.
// in order to avoid duplicated go code, one need to factor the xsd before
//
// This function is currently suited for REQIF schema and could be generalized later
func (schema *Schema) FactorDuplicates() {

	map_NameXSD_Type_Element := make(map[string]map[string][]*Element)

	schema.extractMapOfElementsWithinChoice(map_NameXSD_Type_Element)
	schema.analyse(map_NameXSD_Type_Element)
	log.Println()

	// perform factoring of elements within the choice particle
	schema.factorChoicesFromMap(map_NameXSD_Type_Element)

	log.Println("")

	map_NameXSD_Type_Element = make(map[string]map[string][]*Element)

	schema.extractMapOfElementsWithinChoice(map_NameXSD_Type_Element)
	schema.analyse(map_NameXSD_Type_Element)

	log.Println("")
}

func (schema *Schema) extractMapOfElementsWithinChoice(map_Elements map[string]map[string][]*Element) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					for _, choice := range _complexType.Choices {
						for _, _e := range choice.Elements {
							_e.OuterParticle = choice
							if _, ok := map_Elements[_e.NameXSD]; !ok {
								map_Elements[_e.NameXSD] = make(map[string][]*Element)
							}

							if _, ok := map_Elements[_e.NameXSD][_e.Type]; ok {

								// check that _element is not in the slice already
								if !slices.Contains(map_Elements[_e.NameXSD][_e.Type], _e) {
									map_Elements[_e.NameXSD][_e.Type] = append(
										map_Elements[_e.NameXSD][_e.Type], _e)
								}
							} else {
								map_Elements[_e.NameXSD][_e.Type] = []*Element{_e}
							}
						}
					}
				}
			}
		}
	}
}

func (*Schema) analyse(map_NameXSD_Type_Element map[string]map[string][]*Element) {
	elementsNames := maps.Keys(map_NameXSD_Type_Element)
	sort.Strings(elementsNames)

	for _, elementName := range elementsNames {

		for elementType := range map_NameXSD_Type_Element[elementName] {
			log.Println(elementName, elementType, len(map_NameXSD_Type_Element[elementName][elementType]))
		}
	}
}

func (*Schema) factorChoicesFromMap(map_NameXSD_Type_Element map[string]map[string][]*Element) {
	for elementNameXSD := range map_NameXSD_Type_Element {
		for elementType := range map_NameXSD_Type_Element[elementNameXSD] {
			if sliceElements, ok := map_NameXSD_Type_Element[elementNameXSD][elementType]; ok {
				firstOfTypeElement := sliceElements[0]

				// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
				for _, elementToRemoveFromChoice := range sliceElements[1:] {
					elementToRemoveFromChoice.IsDuplicatedInXSD = true
					if choice, ok := elementToRemoveFromChoice.OuterParticle.(*Choice); ok {
						var newElementsWithinChoice []*Element
						for _, _elem := range choice.Elements {
							if _elem == elementToRemoveFromChoice {
								newElementsWithinChoice = append(newElementsWithinChoice, firstOfTypeElement)
							} else {
								newElementsWithinChoice = append(newElementsWithinChoice, _elem)
							}
						}
						choice.Elements = newElementsWithinChoice
					}
				}
			}
		}
	}
}
