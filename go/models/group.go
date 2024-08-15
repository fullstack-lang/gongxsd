package models

type Group struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Ref string `xml:"ref,attr"`

	// analysis
	IsAnonymous  bool // it has been defined by the enclosing element
	OuterElement *Element
	WithGoIdentifier

	ModelGroup
}

func (group *Group) GetFields(stage *StageStruct) (fields string) {
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

	map_Name_Elems := make(map[string]*Element)

	group.ModelGroup.generateElements(map_Name_Elems, stMap, ctMap, groupMap, setOfGoIdentifiers, &fields)

	return
}
