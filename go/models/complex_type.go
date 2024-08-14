package models

type ComplexType struct {
	Name string

	WithGoIdentifier

	// analysis
	IsAnonymous bool // it has been defined inline by the enclosing element
	DerivedFrom *Element

	Annotated
	ElementWithNameAttribute

	ModelGroup

	Extension *Extension `xml:"extension"`

	// xs:simpleContent element is used exclusively within an xs:complexType element
	// to define complex types that contain simple content along with
	// attributes or restrictions.
	SimpleContent *SimpleContent `xml:"simpleContent"`

	ComplexContent *ComplexContent `xml:"complexContent"`

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

	if ct.SimpleContent != nil {
		if ct.SimpleContent.Extension != nil {
			generateAttributes(ct.SimpleContent.Extension.Attributes, stMap, setOfGoIdentifiers, &fields)
		}
	}

	map_Name_Elems := make(map[string]*Element)

	ct.ModelGroup.generateElements(map_Name_Elems, stMap, ctMap, groupMap, setOfGoIdentifiers, &fields)

	return
}
