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
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[ComplexType](stage) {
		x.Name = x.NameXSD

		if _x, ok := map_EmbeddedComplexStruct[x]; ok {
			x.Name = prefix(_x.Name)
		}

		if x.Sequence != nil {
			x.Sequence.Name = prefix(x.Name)
		}
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[SimpleType](stage) {
		x.Name = x.NameXSD

		if x.Restriction != nil {
			x.Restriction.Name = prefix(x.Name)
		}
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Schema](stage) {
		x.Name = "Schema"
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Restriction](stage) {
		for _, e := range x.Enumerations {
			e.Name = prefix(x.Name)
		}
		if x.MinInclusive != nil {
			x.MinInclusive.Name = prefix(x.Name)
		}
		if x.MaxInclusive != nil {
			x.MaxInclusive.Name = prefix(x.Name)
		}
		if x.Pattern != nil {
			x.Pattern.Name = prefix(x.Name)
		}
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	for x := range *GetGongstructInstancesSet[Enumeration](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Enumeration](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[MinInclusive](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[MaxInclusive](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Pattern](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Annotation](stage) {
		for _, d := range x.Documentations {
			d.Name = prefix(x.Name)
		}
	}
}
