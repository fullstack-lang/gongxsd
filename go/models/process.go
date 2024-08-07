package models

func prefix(s string) string {
	return s + "_E"
}

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

		if x, ok := map_EmbeddedComplexStruct[complexType]; ok {
			complexType.Name = prefix(x.Name)
		}

		if complexType.Sequence != nil {
			complexType.Sequence.Name = prefix(complexType.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[SimpleType](stage) {
		x.Name = x.NameXSD

		if x.Restriction != nil {
			x.Restriction.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Schema](stage) {
		x.Name = "Schema"
	}
	for r := range *GetGongstructInstancesSet[Restriction](stage) {
		for _, e := range r.Enumerations {
			e.Name = prefix(r.Name)
		}
		if r.MinInclusive != nil {
			r.MinInclusive.Name = prefix(r.Name)
		}
		if r.MaxInclusive != nil {
			r.MaxInclusive.Name = prefix(r.Name)
		}
	}

}
