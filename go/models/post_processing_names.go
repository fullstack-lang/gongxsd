package models

func prefix(s string) string {
	return s + "_Inlined"
}

func PostProcessingNames(stage *StageStruct) {

	// map of embedded complex struct within elements
	map_EmbeddedComplexType := make(map[*ComplexType]*Element)
	map_EmbeddedGroup := make(map[*Group]*Element)
	setOfGoIdentifiers := make(map[string]any)

	for x := range *GetGongstructInstancesSet[ComplexType](stage) {

		if x.NameXSD == "" {
			continue
		}
		x.Name = x.NameXSD

		computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if _x, ok := map_EmbeddedComplexType[x]; ok {
			x.Name = prefix(_x.Name)
		}

		x.ModelGroup.nameRecursively(x.Name)

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
		for _, ag := range x.AttributeGroups {
			ag.Name = prefix(x.Name)
		}
		if x.SimpleContent != nil {
			x.SimpleContent.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Group](stage) {

		if x.Ref != "" {
			continue
		}
		x.Name = x.NameXSD

		computeGoIdentifier("Group_"+x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if _x, ok := map_EmbeddedGroup[x]; ok {
			x.Name = prefix(_x.Name)
		}

		x.ModelGroup.nameRecursively(x.Name)

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
		if x.WhiteSpace != nil {
			x.WhiteSpace.Name = prefix(x.Name)
		}
		if x.MinLength != nil {
			x.MinLength.Name = prefix(x.Name)
		}
		if x.MaxLength != nil {
			x.MaxLength.Name = prefix(x.Name)
		}
		if x.Length != nil {
			x.Length.Name = prefix(x.Name)
		}
		if x.TotalDigit != nil {
			x.TotalDigit.Name = prefix(x.Name)
		}

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	//
	// Restrictions
	//
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
	for x := range *GetGongstructInstancesSet[WhiteSpace](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[MinLength](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[MaxLength](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Length](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[TotalDigit](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	for x := range *GetGongstructInstancesSet[Annotation](stage) {
		for _, d := range x.Documentations {
			d.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[Attribute](stage) {
		x.Name = x.NameXSD
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for x := range *GetGongstructInstancesSet[AttributeGroup](stage) {
		x.Name = x.NameXSD

		computeGoIdentifier("AttributeGroup_"+x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	for x := range *GetGongstructInstancesSet[SimpleContent](stage) {

		if x.Extension != nil {
			x.Extension.Name = prefix(x.Name)
		}
		if x.Restriction != nil {
			x.Restriction.Name = prefix(x.Name)
		}
	}

	for x := range *GetGongstructInstancesSet[Element](stage) {
		x.Name = x.NameXSD
		x.GoIdentifier = xsdNameToGoIdentifier(x.Name)

		if x.ComplexType != nil {
			map_EmbeddedComplexType[x.ComplexType] = x

			if x.ComplexType.Name == "" {
				x.ComplexType.Name = "A_" + x.Name
			}
			computeGoIdentifier(x.ComplexType.Name, &x.ComplexType.WithGoIdentifier, setOfGoIdentifiers)

			setOfGoIdentifiers := make(map[string]any)
			computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)
		}
		for _, group := range x.Groups {
			map_EmbeddedGroup[group] = x
			computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)
		}

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

}
