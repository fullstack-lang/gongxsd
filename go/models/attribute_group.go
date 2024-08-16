package models

type AttributeGroup struct {
	Name string
	ElementWithNameAttribute
	Annotated
	WithGoIdentifier

	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`

	Ref string `xml:"ref,attr"`

	Attributes []*Attribute `xml:"attribute"`
}

func (ag *AttributeGroup) generateAttributes(
	agMap map[string]*AttributeGroup,
	stMap map[string]*SimpleType,
	setOfGoIdentifiers map[string]any,
	fields *string) {

	for _, referencedAg := range ag.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			generateAttributes(namedAg.Attributes, stMap, setOfGoIdentifiers, fields)
			namedAg.generateAttributes(agMap, stMap, setOfGoIdentifiers, fields)
		}
	}

}
