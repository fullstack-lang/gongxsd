package models

func Process(stage *StageStruct) {

	// map of embedded complex struct within elements
	map_EmbeddedComplexStruct := make(map[*ComplexType]*Element)

	for x := range *GetGongstructInstancesSet[Element](stage) {
		x.Name = x.NameXSD

		if x.ComplexType != nil {
			map_EmbeddedComplexStruct[x.ComplexType] = x
		}
	}
	for complexType := range *GetGongstructInstancesSet[ComplexType](stage) {
		complexType.Name = complexType.NameXSD

		if element, ok := map_EmbeddedComplexStruct[complexType]; ok {
			complexType.Name = "within " + element.Name
		}

		if complexType.Sequence != nil {
			complexType.Sequence.Name = "within " + complexType.Name
		}
	}
	for x := range *GetGongstructInstancesSet[SimpleType](stage) {
		x.Name = x.NameXSD

		if x.Restriction != nil {
			x.Restriction.Name = "within " + x.Name
		}
	}
	for x := range *GetGongstructInstancesSet[Schema](stage) {
		x.Name = "Schema"
	}
	for x := range *GetGongstructInstancesSet[Restriction](stage) {
		for _, e := range x.Enumerations {
			e.Name = "within " + x.Name
		}
	}

}
