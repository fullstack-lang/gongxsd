package models

import (
	"cmp"
	"log"
	"slices"
	"sort"
	"strconv"

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

	// Step 1. elements within choices
	schema.extractMapOfElementsWithinChoice(map_NameXSD_Type_Element)
	schema.analyseMapOfElements(map_NameXSD_Type_Element)
	schema.factorElementsWithinChoices(map_NameXSD_Type_Element)

	// Step 2 choices within complex types
	map_Choices := make(map[choiceId][]*Choice)
	schema.extractMapOfChoicesWithinComplexTypes(map_Choices)
	schema.analyseMapOfChoices(map_Choices)
	schema.factorChoicesWithinComplexType(map_Choices)

	log.Println("")
}

// if choice is with only those elements, one can factor them
type choiceId struct {
	minOccurs   int
	maxOccurs   int
	elementName string
	elementType string
}

// choices within complex type

func genChoiceId(choice *Choice) (res choiceId) {
	minOccurs, _ := strconv.Atoi(choice.MinOccurs)
	maxOccurs, _ := strconv.Atoi(choice.MaxOccurs)
	res = choiceId{
		minOccurs:   minOccurs,
		maxOccurs:   maxOccurs,
		elementName: choice.Elements[0].NameXSD,
		elementType: choice.Elements[0].Type,
	}
	return
}

func (schema *Schema) extractMapOfChoicesWithinComplexTypes(map_Choices map[choiceId][]*Choice) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					for _, choice := range _complexType.Choices {
						// get rid of choices that are not with one element only
						if len(choice.Alls) > 0 ||
							len(choice.Sequences) > 0 ||
							len(choice.Choices) > 0 ||
							len(choice.Elements) != 1 {
							continue
						}
						choiceId := genChoiceId(choice)
						if _, ok := map_Choices[choiceId]; ok {
							if !slices.Contains(map_Choices[choiceId], choice) {
								map_Choices[choiceId] = append(map_Choices[choiceId], choice)
							}
						} else {
							map_Choices[choiceId] = []*Choice{choice}
						}
					}
				}
			}
		}
	}
}

func (*Schema) analyseMapOfChoices(map_Choices map[choiceId][]*Choice) {
	choices := maps.Keys(map_Choices)
	slices.SortFunc(choices, func(a, b choiceId) int {
		return cmp.Compare(a.elementName, b.elementName)
	})

	for _, choiceId := range choices {
		log.Println("choice", choiceId.elementName, choiceId.elementType, len(map_Choices[choiceId]))
	}
}

func (*Schema) factorChoicesWithinComplexType(map_Choices map[choiceId][]*Choice) {
	for _, sliceElements := range map_Choices {

		firstOfTypeElement := sliceElements[0]

		// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
		for _, choiceToRemoveFromComplexType := range sliceElements[1:] {
			choiceToRemoveFromComplexType.IsDuplicatedInXSD = true
			if complexType, ok := choiceToRemoveFromComplexType.OuterParticle.(*Choice); ok {
				var newChoicesWithinElement []*Choice
				for _, _elem := range complexType.Choices {
					if _elem == choiceToRemoveFromComplexType {
						newChoicesWithinElement = append(newChoicesWithinElement, firstOfTypeElement)
					} else {
						newChoicesWithinElement = append(newChoicesWithinElement, _elem)
					}
				}
				complexType.Choices = newChoicesWithinElement
			}
		}
	}
}

// elements within choices

func (schema *Schema) extractMapOfElementsWithinChoice(map_Elements map[string]map[string][]*Element) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					for _, choice := range _complexType.Choices {
						choice.OuterParticle = complexType
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

func (*Schema) analyseMapOfElements(map_NameXSD_Type_Element map[string]map[string][]*Element) {
	elementsNames := maps.Keys(map_NameXSD_Type_Element)
	sort.Strings(elementsNames)

	for _, elementName := range elementsNames {

		for elementType := range map_NameXSD_Type_Element[elementName] {
			log.Println(elementName, elementType, len(map_NameXSD_Type_Element[elementName][elementType]))
		}
	}
}

func (*Schema) factorElementsWithinChoices(map_NameXSD_Type_Element map[string]map[string][]*Element) {
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
