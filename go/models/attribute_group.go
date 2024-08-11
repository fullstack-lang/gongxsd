package models

type AttributeGroup struct {
	Name string
	ElementWithNameAttribute
	ElementWithAnnotation

	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`

	Ref string `xml:"ref,attr"`

	Attributes []*Attribute `xml:"attribute"`
}

func (ag *AttributeGroup) generateAttributes(agMap map[string]*AttributeGroup, stMap map[string]*SimpleType, fields *string) {

	for _, referencedAg := range ag.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			generateAttributes(namedAg.Attributes, stMap, fields)
			namedAg.generateAttributes(agMap, stMap, fields)
		}
	}

}
