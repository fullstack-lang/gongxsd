package models

type ComplexType struct {
	Name string

	WithGoIdentifier

	// analysis
	IsInlined        bool // it has been defined by the enclosing element
	EnclosingElement *Element

	ElementWithAnnotation
	ElementWithNameAttribute
	Composer

	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}

func (ct *ComplexType) GetFields(stage *StageStruct) (fields string) {
	setOfGoIdentifiers := make(map[string]any)

	stMap := make(map[string]*SimpleType)
	for st := range *GetGongstructInstancesSet[SimpleType](stage) {
		stMap[st.Name] = st
	}
	ctMap := make(map[string]*ComplexType)
	for st := range *GetGongstructInstancesSet[ComplexType](stage) {
		ctMap[st.Name] = st
	}
	agMap := make(map[string]*AttributeGroup)
	for ag := range *GetGongstructInstancesSet[AttributeGroup](stage) {
		agMap[ag.Name] = ag
	}
	groupMap := make(map[string]*Group)
	for group := range *GetGongstructInstancesSet[Group](stage) {
		groupMap[group.Name] = group
	}

	generateAttributes(ct.Attributes, stMap, setOfGoIdentifiers, &fields)
	for _, referencedAg := range ct.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			generateAttributes(namedAg.Attributes, stMap, setOfGoIdentifiers, &fields)
			namedAg.generateAttributes(agMap, stMap, setOfGoIdentifiers, &fields)
		}
	}

	map_Name_Elems := make(map[string]*Element)

	ct.Composer.generateElements(map_Name_Elems, stMap, ctMap, groupMap, setOfGoIdentifiers, &fields)

	return
}
