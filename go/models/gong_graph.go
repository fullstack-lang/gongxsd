// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *All:
		ok = stage.IsStagedAll(target)

	case *Annotation:
		ok = stage.IsStagedAnnotation(target)

	case *Attribute:
		ok = stage.IsStagedAttribute(target)

	case *AttributeGroup:
		ok = stage.IsStagedAttributeGroup(target)

	case *Choice:
		ok = stage.IsStagedChoice(target)

	case *ComplexContent:
		ok = stage.IsStagedComplexContent(target)

	case *ComplexType:
		ok = stage.IsStagedComplexType(target)

	case *Documentation:
		ok = stage.IsStagedDocumentation(target)

	case *Element:
		ok = stage.IsStagedElement(target)

	case *Enumeration:
		ok = stage.IsStagedEnumeration(target)

	case *Extension:
		ok = stage.IsStagedExtension(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Length:
		ok = stage.IsStagedLength(target)

	case *MaxInclusive:
		ok = stage.IsStagedMaxInclusive(target)

	case *MaxLength:
		ok = stage.IsStagedMaxLength(target)

	case *MinInclusive:
		ok = stage.IsStagedMinInclusive(target)

	case *MinLength:
		ok = stage.IsStagedMinLength(target)

	case *Pattern:
		ok = stage.IsStagedPattern(target)

	case *Restriction:
		ok = stage.IsStagedRestriction(target)

	case *Schema:
		ok = stage.IsStagedSchema(target)

	case *Sequence:
		ok = stage.IsStagedSequence(target)

	case *SimpleContent:
		ok = stage.IsStagedSimpleContent(target)

	case *SimpleType:
		ok = stage.IsStagedSimpleType(target)

	case *TotalDigit:
		ok = stage.IsStagedTotalDigit(target)

	case *Union:
		ok = stage.IsStagedUnion(target)

	case *WhiteSpace:
		ok = stage.IsStagedWhiteSpace(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAll(all *All) (ok bool) {

	_, ok = stage.Alls[all]

	return
}

func (stage *StageStruct) IsStagedAnnotation(annotation *Annotation) (ok bool) {

	_, ok = stage.Annotations[annotation]

	return
}

func (stage *StageStruct) IsStagedAttribute(attribute *Attribute) (ok bool) {

	_, ok = stage.Attributes[attribute]

	return
}

func (stage *StageStruct) IsStagedAttributeGroup(attributegroup *AttributeGroup) (ok bool) {

	_, ok = stage.AttributeGroups[attributegroup]

	return
}

func (stage *StageStruct) IsStagedChoice(choice *Choice) (ok bool) {

	_, ok = stage.Choices[choice]

	return
}

func (stage *StageStruct) IsStagedComplexContent(complexcontent *ComplexContent) (ok bool) {

	_, ok = stage.ComplexContents[complexcontent]

	return
}

func (stage *StageStruct) IsStagedComplexType(complextype *ComplexType) (ok bool) {

	_, ok = stage.ComplexTypes[complextype]

	return
}

func (stage *StageStruct) IsStagedDocumentation(documentation *Documentation) (ok bool) {

	_, ok = stage.Documentations[documentation]

	return
}

func (stage *StageStruct) IsStagedElement(element *Element) (ok bool) {

	_, ok = stage.Elements[element]

	return
}

func (stage *StageStruct) IsStagedEnumeration(enumeration *Enumeration) (ok bool) {

	_, ok = stage.Enumerations[enumeration]

	return
}

func (stage *StageStruct) IsStagedExtension(extension *Extension) (ok bool) {

	_, ok = stage.Extensions[extension]

	return
}

func (stage *StageStruct) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *StageStruct) IsStagedLength(length *Length) (ok bool) {

	_, ok = stage.Lengths[length]

	return
}

func (stage *StageStruct) IsStagedMaxInclusive(maxinclusive *MaxInclusive) (ok bool) {

	_, ok = stage.MaxInclusives[maxinclusive]

	return
}

func (stage *StageStruct) IsStagedMaxLength(maxlength *MaxLength) (ok bool) {

	_, ok = stage.MaxLengths[maxlength]

	return
}

func (stage *StageStruct) IsStagedMinInclusive(mininclusive *MinInclusive) (ok bool) {

	_, ok = stage.MinInclusives[mininclusive]

	return
}

func (stage *StageStruct) IsStagedMinLength(minlength *MinLength) (ok bool) {

	_, ok = stage.MinLengths[minlength]

	return
}

func (stage *StageStruct) IsStagedPattern(pattern *Pattern) (ok bool) {

	_, ok = stage.Patterns[pattern]

	return
}

func (stage *StageStruct) IsStagedRestriction(restriction *Restriction) (ok bool) {

	_, ok = stage.Restrictions[restriction]

	return
}

func (stage *StageStruct) IsStagedSchema(schema *Schema) (ok bool) {

	_, ok = stage.Schemas[schema]

	return
}

func (stage *StageStruct) IsStagedSequence(sequence *Sequence) (ok bool) {

	_, ok = stage.Sequences[sequence]

	return
}

func (stage *StageStruct) IsStagedSimpleContent(simplecontent *SimpleContent) (ok bool) {

	_, ok = stage.SimpleContents[simplecontent]

	return
}

func (stage *StageStruct) IsStagedSimpleType(simpletype *SimpleType) (ok bool) {

	_, ok = stage.SimpleTypes[simpletype]

	return
}

func (stage *StageStruct) IsStagedTotalDigit(totaldigit *TotalDigit) (ok bool) {

	_, ok = stage.TotalDigits[totaldigit]

	return
}

func (stage *StageStruct) IsStagedUnion(union *Union) (ok bool) {

	_, ok = stage.Unions[union]

	return
}

func (stage *StageStruct) IsStagedWhiteSpace(whitespace *WhiteSpace) (ok bool) {

	_, ok = stage.WhiteSpaces[whitespace]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *All:
		stage.StageBranchAll(target)

	case *Annotation:
		stage.StageBranchAnnotation(target)

	case *Attribute:
		stage.StageBranchAttribute(target)

	case *AttributeGroup:
		stage.StageBranchAttributeGroup(target)

	case *Choice:
		stage.StageBranchChoice(target)

	case *ComplexContent:
		stage.StageBranchComplexContent(target)

	case *ComplexType:
		stage.StageBranchComplexType(target)

	case *Documentation:
		stage.StageBranchDocumentation(target)

	case *Element:
		stage.StageBranchElement(target)

	case *Enumeration:
		stage.StageBranchEnumeration(target)

	case *Extension:
		stage.StageBranchExtension(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *Length:
		stage.StageBranchLength(target)

	case *MaxInclusive:
		stage.StageBranchMaxInclusive(target)

	case *MaxLength:
		stage.StageBranchMaxLength(target)

	case *MinInclusive:
		stage.StageBranchMinInclusive(target)

	case *MinLength:
		stage.StageBranchMinLength(target)

	case *Pattern:
		stage.StageBranchPattern(target)

	case *Restriction:
		stage.StageBranchRestriction(target)

	case *Schema:
		stage.StageBranchSchema(target)

	case *Sequence:
		stage.StageBranchSequence(target)

	case *SimpleContent:
		stage.StageBranchSimpleContent(target)

	case *SimpleType:
		stage.StageBranchSimpleType(target)

	case *TotalDigit:
		stage.StageBranchTotalDigit(target)

	case *Union:
		stage.StageBranchUnion(target)

	case *WhiteSpace:
		stage.StageBranchWhiteSpace(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAll(all *All) {

	// check if instance is already staged
	if IsStaged(stage, all) {
		return
	}

	all.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		StageBranch(stage, all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range all.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range all.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range all.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range all.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *StageStruct) StageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if IsStaged(stage, annotation) {
		return
	}

	annotation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		StageBranch(stage, _documentation)
	}

}

func (stage *StageStruct) StageBranchAttribute(attribute *Attribute) {

	// check if instance is already staged
	if IsStaged(stage, attribute) {
		return
	}

	attribute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		StageBranch(stage, attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAttributeGroup(attributegroup *AttributeGroup) {

	// check if instance is already staged
	if IsStaged(stage, attributegroup) {
		return
	}

	attributegroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		StageBranch(stage, attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		StageBranch(stage, _attribute)
	}

}

func (stage *StageStruct) StageBranchChoice(choice *Choice) {

	// check if instance is already staged
	if IsStaged(stage, choice) {
		return
	}

	choice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		StageBranch(stage, choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range choice.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range choice.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range choice.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range choice.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *StageStruct) StageBranchComplexContent(complexcontent *ComplexContent) {

	// check if instance is already staged
	if IsStaged(stage, complexcontent) {
		return
	}

	complexcontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchComplexType(complextype *ComplexType) {

	// check if instance is already staged
	if IsStaged(stage, complextype) {
		return
	}

	complextype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		StageBranch(stage, complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		StageBranch(stage, complextype.Annotation)
	}
	if complextype.Extension != nil {
		StageBranch(stage, complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		StageBranch(stage, complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		StageBranch(stage, complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range complextype.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range complextype.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range complextype.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range complextype.Elements {
		StageBranch(stage, _element)
	}
	for _, _attribute := range complextype.Attributes {
		StageBranch(stage, _attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}

}

func (stage *StageStruct) StageBranchDocumentation(documentation *Documentation) {

	// check if instance is already staged
	if IsStaged(stage, documentation) {
		return
	}

	documentation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchElement(element *Element) {

	// check if instance is already staged
	if IsStaged(stage, element) {
		return
	}

	element.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		StageBranch(stage, element.Annotation)
	}
	if element.SimpleType != nil {
		StageBranch(stage, element.SimpleType)
	}
	if element.ComplexType != nil {
		StageBranch(stage, element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		StageBranch(stage, _group)
	}

}

func (stage *StageStruct) StageBranchEnumeration(enumeration *Enumeration) {

	// check if instance is already staged
	if IsStaged(stage, enumeration) {
		return
	}

	enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		StageBranch(stage, enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchExtension(extension *Extension) {

	// check if instance is already staged
	if IsStaged(stage, extension) {
		return
	}

	extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range extension.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range extension.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range extension.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range extension.Elements {
		StageBranch(stage, _element)
	}
	for _, _attribute := range extension.Attributes {
		StageBranch(stage, _attribute)
	}

}

func (stage *StageStruct) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		StageBranch(stage, group.Annotation)
	}
	if group.OuterElement != nil {
		StageBranch(stage, group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range group.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range group.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range group.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range group.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *StageStruct) StageBranchLength(length *Length) {

	// check if instance is already staged
	if IsStaged(stage, length) {
		return
	}

	length.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		StageBranch(stage, length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMaxInclusive(maxinclusive *MaxInclusive) {

	// check if instance is already staged
	if IsStaged(stage, maxinclusive) {
		return
	}

	maxinclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		StageBranch(stage, maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMaxLength(maxlength *MaxLength) {

	// check if instance is already staged
	if IsStaged(stage, maxlength) {
		return
	}

	maxlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		StageBranch(stage, maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMinInclusive(mininclusive *MinInclusive) {

	// check if instance is already staged
	if IsStaged(stage, mininclusive) {
		return
	}

	mininclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		StageBranch(stage, mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMinLength(minlength *MinLength) {

	// check if instance is already staged
	if IsStaged(stage, minlength) {
		return
	}

	minlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		StageBranch(stage, minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPattern(pattern *Pattern) {

	// check if instance is already staged
	if IsStaged(stage, pattern) {
		return
	}

	pattern.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		StageBranch(stage, pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRestriction(restriction *Restriction) {

	// check if instance is already staged
	if IsStaged(stage, restriction) {
		return
	}

	restriction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		StageBranch(stage, restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		StageBranch(stage, restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		StageBranch(stage, restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		StageBranch(stage, restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		StageBranch(stage, restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		StageBranch(stage, restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		StageBranch(stage, restriction.MaxLength)
	}
	if restriction.Length != nil {
		StageBranch(stage, restriction.Length)
	}
	if restriction.TotalDigit != nil {
		StageBranch(stage, restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		StageBranch(stage, _enumeration)
	}

}

func (stage *StageStruct) StageBranchSchema(schema *Schema) {

	// check if instance is already staged
	if IsStaged(stage, schema) {
		return
	}

	schema.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		StageBranch(stage, schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		StageBranch(stage, _element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		StageBranch(stage, _simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		StageBranch(stage, _complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}
	for _, _group := range schema.Groups {
		StageBranch(stage, _group)
	}

}

func (stage *StageStruct) StageBranchSequence(sequence *Sequence) {

	// check if instance is already staged
	if IsStaged(stage, sequence) {
		return
	}

	sequence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		StageBranch(stage, sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range sequence.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range sequence.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range sequence.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range sequence.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *StageStruct) StageBranchSimpleContent(simplecontent *SimpleContent) {

	// check if instance is already staged
	if IsStaged(stage, simplecontent) {
		return
	}

	simplecontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		StageBranch(stage, simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		StageBranch(stage, simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSimpleType(simpletype *SimpleType) {

	// check if instance is already staged
	if IsStaged(stage, simpletype) {
		return
	}

	simpletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		StageBranch(stage, simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		StageBranch(stage, simpletype.Restriction)
	}
	if simpletype.Union != nil {
		StageBranch(stage, simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTotalDigit(totaldigit *TotalDigit) {

	// check if instance is already staged
	if IsStaged(stage, totaldigit) {
		return
	}

	totaldigit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		StageBranch(stage, totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUnion(union *Union) {

	// check if instance is already staged
	if IsStaged(stage, union) {
		return
	}

	union.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		StageBranch(stage, union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWhiteSpace(whitespace *WhiteSpace) {

	// check if instance is already staged
	if IsStaged(stage, whitespace) {
		return
	}

	whitespace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		StageBranch(stage, whitespace.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *All:
		toT := CopyBranchAll(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Annotation:
		toT := CopyBranchAnnotation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Attribute:
		toT := CopyBranchAttribute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AttributeGroup:
		toT := CopyBranchAttributeGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Choice:
		toT := CopyBranchChoice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexContent:
		toT := CopyBranchComplexContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexType:
		toT := CopyBranchComplexType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Documentation:
		toT := CopyBranchDocumentation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Element:
		toT := CopyBranchElement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Enumeration:
		toT := CopyBranchEnumeration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extension:
		toT := CopyBranchExtension(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Length:
		toT := CopyBranchLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxInclusive:
		toT := CopyBranchMaxInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxLength:
		toT := CopyBranchMaxLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinInclusive:
		toT := CopyBranchMinInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinLength:
		toT := CopyBranchMinLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pattern:
		toT := CopyBranchPattern(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Restriction:
		toT := CopyBranchRestriction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Schema:
		toT := CopyBranchSchema(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sequence:
		toT := CopyBranchSequence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleContent:
		toT := CopyBranchSimpleContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleType:
		toT := CopyBranchSimpleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TotalDigit:
		toT := CopyBranchTotalDigit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Union:
		toT := CopyBranchUnion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *WhiteSpace:
		toT := CopyBranchWhiteSpace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAll(mapOrigCopy map[any]any, allFrom *All) (allTo *All) {

	// allFrom has already been copied
	if _allTo, ok := mapOrigCopy[allFrom]; ok {
		allTo = _allTo.(*All)
		return
	}

	allTo = new(All)
	mapOrigCopy[allFrom] = allTo
	allFrom.CopyBasicFields(allTo)

	//insertion point for the staging of instances referenced by pointers
	if allFrom.Annotation != nil {
		allTo.Annotation = CopyBranchAnnotation(mapOrigCopy, allFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range allFrom.Sequences {
		allTo.Sequences = append(allTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range allFrom.Alls {
		allTo.Alls = append(allTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range allFrom.Choices {
		allTo.Choices = append(allTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range allFrom.Groups {
		allTo.Groups = append(allTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range allFrom.Elements {
		allTo.Elements = append(allTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchAnnotation(mapOrigCopy map[any]any, annotationFrom *Annotation) (annotationTo *Annotation) {

	// annotationFrom has already been copied
	if _annotationTo, ok := mapOrigCopy[annotationFrom]; ok {
		annotationTo = _annotationTo.(*Annotation)
		return
	}

	annotationTo = new(Annotation)
	mapOrigCopy[annotationFrom] = annotationTo
	annotationFrom.CopyBasicFields(annotationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotationFrom.Documentations {
		annotationTo.Documentations = append(annotationTo.Documentations, CopyBranchDocumentation(mapOrigCopy, _documentation))
	}

	return
}

func CopyBranchAttribute(mapOrigCopy map[any]any, attributeFrom *Attribute) (attributeTo *Attribute) {

	// attributeFrom has already been copied
	if _attributeTo, ok := mapOrigCopy[attributeFrom]; ok {
		attributeTo = _attributeTo.(*Attribute)
		return
	}

	attributeTo = new(Attribute)
	mapOrigCopy[attributeFrom] = attributeTo
	attributeFrom.CopyBasicFields(attributeTo)

	//insertion point for the staging of instances referenced by pointers
	if attributeFrom.Annotation != nil {
		attributeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, attributeFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAttributeGroup(mapOrigCopy map[any]any, attributegroupFrom *AttributeGroup) (attributegroupTo *AttributeGroup) {

	// attributegroupFrom has already been copied
	if _attributegroupTo, ok := mapOrigCopy[attributegroupFrom]; ok {
		attributegroupTo = _attributegroupTo.(*AttributeGroup)
		return
	}

	attributegroupTo = new(AttributeGroup)
	mapOrigCopy[attributegroupFrom] = attributegroupTo
	attributegroupFrom.CopyBasicFields(attributegroupTo)

	//insertion point for the staging of instances referenced by pointers
	if attributegroupFrom.Annotation != nil {
		attributegroupTo.Annotation = CopyBranchAnnotation(mapOrigCopy, attributegroupFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroupFrom.AttributeGroups {
		attributegroupTo.AttributeGroups = append(attributegroupTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _attribute := range attributegroupFrom.Attributes {
		attributegroupTo.Attributes = append(attributegroupTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}

	return
}

func CopyBranchChoice(mapOrigCopy map[any]any, choiceFrom *Choice) (choiceTo *Choice) {

	// choiceFrom has already been copied
	if _choiceTo, ok := mapOrigCopy[choiceFrom]; ok {
		choiceTo = _choiceTo.(*Choice)
		return
	}

	choiceTo = new(Choice)
	mapOrigCopy[choiceFrom] = choiceTo
	choiceFrom.CopyBasicFields(choiceTo)

	//insertion point for the staging of instances referenced by pointers
	if choiceFrom.Annotation != nil {
		choiceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, choiceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choiceFrom.Sequences {
		choiceTo.Sequences = append(choiceTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range choiceFrom.Alls {
		choiceTo.Alls = append(choiceTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range choiceFrom.Choices {
		choiceTo.Choices = append(choiceTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range choiceFrom.Groups {
		choiceTo.Groups = append(choiceTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range choiceFrom.Elements {
		choiceTo.Elements = append(choiceTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchComplexContent(mapOrigCopy map[any]any, complexcontentFrom *ComplexContent) (complexcontentTo *ComplexContent) {

	// complexcontentFrom has already been copied
	if _complexcontentTo, ok := mapOrigCopy[complexcontentFrom]; ok {
		complexcontentTo = _complexcontentTo.(*ComplexContent)
		return
	}

	complexcontentTo = new(ComplexContent)
	mapOrigCopy[complexcontentFrom] = complexcontentTo
	complexcontentFrom.CopyBasicFields(complexcontentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchComplexType(mapOrigCopy map[any]any, complextypeFrom *ComplexType) (complextypeTo *ComplexType) {

	// complextypeFrom has already been copied
	if _complextypeTo, ok := mapOrigCopy[complextypeFrom]; ok {
		complextypeTo = _complextypeTo.(*ComplexType)
		return
	}

	complextypeTo = new(ComplexType)
	mapOrigCopy[complextypeFrom] = complextypeTo
	complextypeFrom.CopyBasicFields(complextypeTo)

	//insertion point for the staging of instances referenced by pointers
	if complextypeFrom.OuterElement != nil {
		complextypeTo.OuterElement = CopyBranchElement(mapOrigCopy, complextypeFrom.OuterElement)
	}
	if complextypeFrom.Annotation != nil {
		complextypeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, complextypeFrom.Annotation)
	}
	if complextypeFrom.Extension != nil {
		complextypeTo.Extension = CopyBranchExtension(mapOrigCopy, complextypeFrom.Extension)
	}
	if complextypeFrom.SimpleContent != nil {
		complextypeTo.SimpleContent = CopyBranchSimpleContent(mapOrigCopy, complextypeFrom.SimpleContent)
	}
	if complextypeFrom.ComplexContent != nil {
		complextypeTo.ComplexContent = CopyBranchComplexContent(mapOrigCopy, complextypeFrom.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextypeFrom.Sequences {
		complextypeTo.Sequences = append(complextypeTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range complextypeFrom.Alls {
		complextypeTo.Alls = append(complextypeTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range complextypeFrom.Choices {
		complextypeTo.Choices = append(complextypeTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range complextypeFrom.Groups {
		complextypeTo.Groups = append(complextypeTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range complextypeFrom.Elements {
		complextypeTo.Elements = append(complextypeTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range complextypeFrom.Attributes {
		complextypeTo.Attributes = append(complextypeTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}
	for _, _attributegroup := range complextypeFrom.AttributeGroups {
		complextypeTo.AttributeGroups = append(complextypeTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}

	return
}

func CopyBranchDocumentation(mapOrigCopy map[any]any, documentationFrom *Documentation) (documentationTo *Documentation) {

	// documentationFrom has already been copied
	if _documentationTo, ok := mapOrigCopy[documentationFrom]; ok {
		documentationTo = _documentationTo.(*Documentation)
		return
	}

	documentationTo = new(Documentation)
	mapOrigCopy[documentationFrom] = documentationTo
	documentationFrom.CopyBasicFields(documentationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchElement(mapOrigCopy map[any]any, elementFrom *Element) (elementTo *Element) {

	// elementFrom has already been copied
	if _elementTo, ok := mapOrigCopy[elementFrom]; ok {
		elementTo = _elementTo.(*Element)
		return
	}

	elementTo = new(Element)
	mapOrigCopy[elementFrom] = elementTo
	elementFrom.CopyBasicFields(elementTo)

	//insertion point for the staging of instances referenced by pointers
	if elementFrom.Annotation != nil {
		elementTo.Annotation = CopyBranchAnnotation(mapOrigCopy, elementFrom.Annotation)
	}
	if elementFrom.SimpleType != nil {
		elementTo.SimpleType = CopyBranchSimpleType(mapOrigCopy, elementFrom.SimpleType)
	}
	if elementFrom.ComplexType != nil {
		elementTo.ComplexType = CopyBranchComplexType(mapOrigCopy, elementFrom.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range elementFrom.Groups {
		elementTo.Groups = append(elementTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func CopyBranchEnumeration(mapOrigCopy map[any]any, enumerationFrom *Enumeration) (enumerationTo *Enumeration) {

	// enumerationFrom has already been copied
	if _enumerationTo, ok := mapOrigCopy[enumerationFrom]; ok {
		enumerationTo = _enumerationTo.(*Enumeration)
		return
	}

	enumerationTo = new(Enumeration)
	mapOrigCopy[enumerationFrom] = enumerationTo
	enumerationFrom.CopyBasicFields(enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if enumerationFrom.Annotation != nil {
		enumerationTo.Annotation = CopyBranchAnnotation(mapOrigCopy, enumerationFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExtension(mapOrigCopy map[any]any, extensionFrom *Extension) (extensionTo *Extension) {

	// extensionFrom has already been copied
	if _extensionTo, ok := mapOrigCopy[extensionFrom]; ok {
		extensionTo = _extensionTo.(*Extension)
		return
	}

	extensionTo = new(Extension)
	mapOrigCopy[extensionFrom] = extensionTo
	extensionFrom.CopyBasicFields(extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extensionFrom.Sequences {
		extensionTo.Sequences = append(extensionTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range extensionFrom.Alls {
		extensionTo.Alls = append(extensionTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range extensionFrom.Choices {
		extensionTo.Choices = append(extensionTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range extensionFrom.Groups {
		extensionTo.Groups = append(extensionTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range extensionFrom.Elements {
		extensionTo.Elements = append(extensionTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range extensionFrom.Attributes {
		extensionTo.Attributes = append(extensionTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}

	return
}

func CopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.CopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers
	if groupFrom.Annotation != nil {
		groupTo.Annotation = CopyBranchAnnotation(mapOrigCopy, groupFrom.Annotation)
	}
	if groupFrom.OuterElement != nil {
		groupTo.OuterElement = CopyBranchElement(mapOrigCopy, groupFrom.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range groupFrom.Sequences {
		groupTo.Sequences = append(groupTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range groupFrom.Alls {
		groupTo.Alls = append(groupTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range groupFrom.Choices {
		groupTo.Choices = append(groupTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range groupFrom.Groups {
		groupTo.Groups = append(groupTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range groupFrom.Elements {
		groupTo.Elements = append(groupTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchLength(mapOrigCopy map[any]any, lengthFrom *Length) (lengthTo *Length) {

	// lengthFrom has already been copied
	if _lengthTo, ok := mapOrigCopy[lengthFrom]; ok {
		lengthTo = _lengthTo.(*Length)
		return
	}

	lengthTo = new(Length)
	mapOrigCopy[lengthFrom] = lengthTo
	lengthFrom.CopyBasicFields(lengthTo)

	//insertion point for the staging of instances referenced by pointers
	if lengthFrom.Annotation != nil {
		lengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, lengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMaxInclusive(mapOrigCopy map[any]any, maxinclusiveFrom *MaxInclusive) (maxinclusiveTo *MaxInclusive) {

	// maxinclusiveFrom has already been copied
	if _maxinclusiveTo, ok := mapOrigCopy[maxinclusiveFrom]; ok {
		maxinclusiveTo = _maxinclusiveTo.(*MaxInclusive)
		return
	}

	maxinclusiveTo = new(MaxInclusive)
	mapOrigCopy[maxinclusiveFrom] = maxinclusiveTo
	maxinclusiveFrom.CopyBasicFields(maxinclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusiveFrom.Annotation != nil {
		maxinclusiveTo.Annotation = CopyBranchAnnotation(mapOrigCopy, maxinclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMaxLength(mapOrigCopy map[any]any, maxlengthFrom *MaxLength) (maxlengthTo *MaxLength) {

	// maxlengthFrom has already been copied
	if _maxlengthTo, ok := mapOrigCopy[maxlengthFrom]; ok {
		maxlengthTo = _maxlengthTo.(*MaxLength)
		return
	}

	maxlengthTo = new(MaxLength)
	mapOrigCopy[maxlengthFrom] = maxlengthTo
	maxlengthFrom.CopyBasicFields(maxlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if maxlengthFrom.Annotation != nil {
		maxlengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, maxlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMinInclusive(mapOrigCopy map[any]any, mininclusiveFrom *MinInclusive) (mininclusiveTo *MinInclusive) {

	// mininclusiveFrom has already been copied
	if _mininclusiveTo, ok := mapOrigCopy[mininclusiveFrom]; ok {
		mininclusiveTo = _mininclusiveTo.(*MinInclusive)
		return
	}

	mininclusiveTo = new(MinInclusive)
	mapOrigCopy[mininclusiveFrom] = mininclusiveTo
	mininclusiveFrom.CopyBasicFields(mininclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if mininclusiveFrom.Annotation != nil {
		mininclusiveTo.Annotation = CopyBranchAnnotation(mapOrigCopy, mininclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMinLength(mapOrigCopy map[any]any, minlengthFrom *MinLength) (minlengthTo *MinLength) {

	// minlengthFrom has already been copied
	if _minlengthTo, ok := mapOrigCopy[minlengthFrom]; ok {
		minlengthTo = _minlengthTo.(*MinLength)
		return
	}

	minlengthTo = new(MinLength)
	mapOrigCopy[minlengthFrom] = minlengthTo
	minlengthFrom.CopyBasicFields(minlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if minlengthFrom.Annotation != nil {
		minlengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, minlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPattern(mapOrigCopy map[any]any, patternFrom *Pattern) (patternTo *Pattern) {

	// patternFrom has already been copied
	if _patternTo, ok := mapOrigCopy[patternFrom]; ok {
		patternTo = _patternTo.(*Pattern)
		return
	}

	patternTo = new(Pattern)
	mapOrigCopy[patternFrom] = patternTo
	patternFrom.CopyBasicFields(patternTo)

	//insertion point for the staging of instances referenced by pointers
	if patternFrom.Annotation != nil {
		patternTo.Annotation = CopyBranchAnnotation(mapOrigCopy, patternFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRestriction(mapOrigCopy map[any]any, restrictionFrom *Restriction) (restrictionTo *Restriction) {

	// restrictionFrom has already been copied
	if _restrictionTo, ok := mapOrigCopy[restrictionFrom]; ok {
		restrictionTo = _restrictionTo.(*Restriction)
		return
	}

	restrictionTo = new(Restriction)
	mapOrigCopy[restrictionFrom] = restrictionTo
	restrictionFrom.CopyBasicFields(restrictionTo)

	//insertion point for the staging of instances referenced by pointers
	if restrictionFrom.Annotation != nil {
		restrictionTo.Annotation = CopyBranchAnnotation(mapOrigCopy, restrictionFrom.Annotation)
	}
	if restrictionFrom.MinInclusive != nil {
		restrictionTo.MinInclusive = CopyBranchMinInclusive(mapOrigCopy, restrictionFrom.MinInclusive)
	}
	if restrictionFrom.MaxInclusive != nil {
		restrictionTo.MaxInclusive = CopyBranchMaxInclusive(mapOrigCopy, restrictionFrom.MaxInclusive)
	}
	if restrictionFrom.Pattern != nil {
		restrictionTo.Pattern = CopyBranchPattern(mapOrigCopy, restrictionFrom.Pattern)
	}
	if restrictionFrom.WhiteSpace != nil {
		restrictionTo.WhiteSpace = CopyBranchWhiteSpace(mapOrigCopy, restrictionFrom.WhiteSpace)
	}
	if restrictionFrom.MinLength != nil {
		restrictionTo.MinLength = CopyBranchMinLength(mapOrigCopy, restrictionFrom.MinLength)
	}
	if restrictionFrom.MaxLength != nil {
		restrictionTo.MaxLength = CopyBranchMaxLength(mapOrigCopy, restrictionFrom.MaxLength)
	}
	if restrictionFrom.Length != nil {
		restrictionTo.Length = CopyBranchLength(mapOrigCopy, restrictionFrom.Length)
	}
	if restrictionFrom.TotalDigit != nil {
		restrictionTo.TotalDigit = CopyBranchTotalDigit(mapOrigCopy, restrictionFrom.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restrictionFrom.Enumerations {
		restrictionTo.Enumerations = append(restrictionTo.Enumerations, CopyBranchEnumeration(mapOrigCopy, _enumeration))
	}

	return
}

func CopyBranchSchema(mapOrigCopy map[any]any, schemaFrom *Schema) (schemaTo *Schema) {

	// schemaFrom has already been copied
	if _schemaTo, ok := mapOrigCopy[schemaFrom]; ok {
		schemaTo = _schemaTo.(*Schema)
		return
	}

	schemaTo = new(Schema)
	mapOrigCopy[schemaFrom] = schemaTo
	schemaFrom.CopyBasicFields(schemaTo)

	//insertion point for the staging of instances referenced by pointers
	if schemaFrom.Annotation != nil {
		schemaTo.Annotation = CopyBranchAnnotation(mapOrigCopy, schemaFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schemaFrom.Elements {
		schemaTo.Elements = append(schemaTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _simpletype := range schemaFrom.SimpleTypes {
		schemaTo.SimpleTypes = append(schemaTo.SimpleTypes, CopyBranchSimpleType(mapOrigCopy, _simpletype))
	}
	for _, _complextype := range schemaFrom.ComplexTypes {
		schemaTo.ComplexTypes = append(schemaTo.ComplexTypes, CopyBranchComplexType(mapOrigCopy, _complextype))
	}
	for _, _attributegroup := range schemaFrom.AttributeGroups {
		schemaTo.AttributeGroups = append(schemaTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _group := range schemaFrom.Groups {
		schemaTo.Groups = append(schemaTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func CopyBranchSequence(mapOrigCopy map[any]any, sequenceFrom *Sequence) (sequenceTo *Sequence) {

	// sequenceFrom has already been copied
	if _sequenceTo, ok := mapOrigCopy[sequenceFrom]; ok {
		sequenceTo = _sequenceTo.(*Sequence)
		return
	}

	sequenceTo = new(Sequence)
	mapOrigCopy[sequenceFrom] = sequenceTo
	sequenceFrom.CopyBasicFields(sequenceTo)

	//insertion point for the staging of instances referenced by pointers
	if sequenceFrom.Annotation != nil {
		sequenceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, sequenceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequenceFrom.Sequences {
		sequenceTo.Sequences = append(sequenceTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range sequenceFrom.Alls {
		sequenceTo.Alls = append(sequenceTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range sequenceFrom.Choices {
		sequenceTo.Choices = append(sequenceTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range sequenceFrom.Groups {
		sequenceTo.Groups = append(sequenceTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range sequenceFrom.Elements {
		sequenceTo.Elements = append(sequenceTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchSimpleContent(mapOrigCopy map[any]any, simplecontentFrom *SimpleContent) (simplecontentTo *SimpleContent) {

	// simplecontentFrom has already been copied
	if _simplecontentTo, ok := mapOrigCopy[simplecontentFrom]; ok {
		simplecontentTo = _simplecontentTo.(*SimpleContent)
		return
	}

	simplecontentTo = new(SimpleContent)
	mapOrigCopy[simplecontentFrom] = simplecontentTo
	simplecontentFrom.CopyBasicFields(simplecontentTo)

	//insertion point for the staging of instances referenced by pointers
	if simplecontentFrom.Extension != nil {
		simplecontentTo.Extension = CopyBranchExtension(mapOrigCopy, simplecontentFrom.Extension)
	}
	if simplecontentFrom.Restriction != nil {
		simplecontentTo.Restriction = CopyBranchRestriction(mapOrigCopy, simplecontentFrom.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSimpleType(mapOrigCopy map[any]any, simpletypeFrom *SimpleType) (simpletypeTo *SimpleType) {

	// simpletypeFrom has already been copied
	if _simpletypeTo, ok := mapOrigCopy[simpletypeFrom]; ok {
		simpletypeTo = _simpletypeTo.(*SimpleType)
		return
	}

	simpletypeTo = new(SimpleType)
	mapOrigCopy[simpletypeFrom] = simpletypeTo
	simpletypeFrom.CopyBasicFields(simpletypeTo)

	//insertion point for the staging of instances referenced by pointers
	if simpletypeFrom.Annotation != nil {
		simpletypeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, simpletypeFrom.Annotation)
	}
	if simpletypeFrom.Restriction != nil {
		simpletypeTo.Restriction = CopyBranchRestriction(mapOrigCopy, simpletypeFrom.Restriction)
	}
	if simpletypeFrom.Union != nil {
		simpletypeTo.Union = CopyBranchUnion(mapOrigCopy, simpletypeFrom.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTotalDigit(mapOrigCopy map[any]any, totaldigitFrom *TotalDigit) (totaldigitTo *TotalDigit) {

	// totaldigitFrom has already been copied
	if _totaldigitTo, ok := mapOrigCopy[totaldigitFrom]; ok {
		totaldigitTo = _totaldigitTo.(*TotalDigit)
		return
	}

	totaldigitTo = new(TotalDigit)
	mapOrigCopy[totaldigitFrom] = totaldigitTo
	totaldigitFrom.CopyBasicFields(totaldigitTo)

	//insertion point for the staging of instances referenced by pointers
	if totaldigitFrom.Annotation != nil {
		totaldigitTo.Annotation = CopyBranchAnnotation(mapOrigCopy, totaldigitFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUnion(mapOrigCopy map[any]any, unionFrom *Union) (unionTo *Union) {

	// unionFrom has already been copied
	if _unionTo, ok := mapOrigCopy[unionFrom]; ok {
		unionTo = _unionTo.(*Union)
		return
	}

	unionTo = new(Union)
	mapOrigCopy[unionFrom] = unionTo
	unionFrom.CopyBasicFields(unionTo)

	//insertion point for the staging of instances referenced by pointers
	if unionFrom.Annotation != nil {
		unionTo.Annotation = CopyBranchAnnotation(mapOrigCopy, unionFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWhiteSpace(mapOrigCopy map[any]any, whitespaceFrom *WhiteSpace) (whitespaceTo *WhiteSpace) {

	// whitespaceFrom has already been copied
	if _whitespaceTo, ok := mapOrigCopy[whitespaceFrom]; ok {
		whitespaceTo = _whitespaceTo.(*WhiteSpace)
		return
	}

	whitespaceTo = new(WhiteSpace)
	mapOrigCopy[whitespaceFrom] = whitespaceTo
	whitespaceFrom.CopyBasicFields(whitespaceTo)

	//insertion point for the staging of instances referenced by pointers
	if whitespaceFrom.Annotation != nil {
		whitespaceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, whitespaceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *All:
		stage.UnstageBranchAll(target)

	case *Annotation:
		stage.UnstageBranchAnnotation(target)

	case *Attribute:
		stage.UnstageBranchAttribute(target)

	case *AttributeGroup:
		stage.UnstageBranchAttributeGroup(target)

	case *Choice:
		stage.UnstageBranchChoice(target)

	case *ComplexContent:
		stage.UnstageBranchComplexContent(target)

	case *ComplexType:
		stage.UnstageBranchComplexType(target)

	case *Documentation:
		stage.UnstageBranchDocumentation(target)

	case *Element:
		stage.UnstageBranchElement(target)

	case *Enumeration:
		stage.UnstageBranchEnumeration(target)

	case *Extension:
		stage.UnstageBranchExtension(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *Length:
		stage.UnstageBranchLength(target)

	case *MaxInclusive:
		stage.UnstageBranchMaxInclusive(target)

	case *MaxLength:
		stage.UnstageBranchMaxLength(target)

	case *MinInclusive:
		stage.UnstageBranchMinInclusive(target)

	case *MinLength:
		stage.UnstageBranchMinLength(target)

	case *Pattern:
		stage.UnstageBranchPattern(target)

	case *Restriction:
		stage.UnstageBranchRestriction(target)

	case *Schema:
		stage.UnstageBranchSchema(target)

	case *Sequence:
		stage.UnstageBranchSequence(target)

	case *SimpleContent:
		stage.UnstageBranchSimpleContent(target)

	case *SimpleType:
		stage.UnstageBranchSimpleType(target)

	case *TotalDigit:
		stage.UnstageBranchTotalDigit(target)

	case *Union:
		stage.UnstageBranchUnion(target)

	case *WhiteSpace:
		stage.UnstageBranchWhiteSpace(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAll(all *All) {

	// check if instance is already staged
	if !IsStaged(stage, all) {
		return
	}

	all.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		UnstageBranch(stage, all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range all.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range all.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range all.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range all.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *StageStruct) UnstageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if !IsStaged(stage, annotation) {
		return
	}

	annotation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		UnstageBranch(stage, _documentation)
	}

}

func (stage *StageStruct) UnstageBranchAttribute(attribute *Attribute) {

	// check if instance is already staged
	if !IsStaged(stage, attribute) {
		return
	}

	attribute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		UnstageBranch(stage, attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAttributeGroup(attributegroup *AttributeGroup) {

	// check if instance is already staged
	if !IsStaged(stage, attributegroup) {
		return
	}

	attributegroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		UnstageBranch(stage, attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		UnstageBranch(stage, _attribute)
	}

}

func (stage *StageStruct) UnstageBranchChoice(choice *Choice) {

	// check if instance is already staged
	if !IsStaged(stage, choice) {
		return
	}

	choice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		UnstageBranch(stage, choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range choice.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range choice.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range choice.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range choice.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *StageStruct) UnstageBranchComplexContent(complexcontent *ComplexContent) {

	// check if instance is already staged
	if !IsStaged(stage, complexcontent) {
		return
	}

	complexcontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchComplexType(complextype *ComplexType) {

	// check if instance is already staged
	if !IsStaged(stage, complextype) {
		return
	}

	complextype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		UnstageBranch(stage, complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		UnstageBranch(stage, complextype.Annotation)
	}
	if complextype.Extension != nil {
		UnstageBranch(stage, complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		UnstageBranch(stage, complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		UnstageBranch(stage, complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range complextype.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range complextype.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range complextype.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range complextype.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _attribute := range complextype.Attributes {
		UnstageBranch(stage, _attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}

}

func (stage *StageStruct) UnstageBranchDocumentation(documentation *Documentation) {

	// check if instance is already staged
	if !IsStaged(stage, documentation) {
		return
	}

	documentation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchElement(element *Element) {

	// check if instance is already staged
	if !IsStaged(stage, element) {
		return
	}

	element.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		UnstageBranch(stage, element.Annotation)
	}
	if element.SimpleType != nil {
		UnstageBranch(stage, element.SimpleType)
	}
	if element.ComplexType != nil {
		UnstageBranch(stage, element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		UnstageBranch(stage, _group)
	}

}

func (stage *StageStruct) UnstageBranchEnumeration(enumeration *Enumeration) {

	// check if instance is already staged
	if !IsStaged(stage, enumeration) {
		return
	}

	enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		UnstageBranch(stage, enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchExtension(extension *Extension) {

	// check if instance is already staged
	if !IsStaged(stage, extension) {
		return
	}

	extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range extension.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range extension.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range extension.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range extension.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _attribute := range extension.Attributes {
		UnstageBranch(stage, _attribute)
	}

}

func (stage *StageStruct) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		UnstageBranch(stage, group.Annotation)
	}
	if group.OuterElement != nil {
		UnstageBranch(stage, group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range group.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range group.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range group.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range group.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *StageStruct) UnstageBranchLength(length *Length) {

	// check if instance is already staged
	if !IsStaged(stage, length) {
		return
	}

	length.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		UnstageBranch(stage, length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMaxInclusive(maxinclusive *MaxInclusive) {

	// check if instance is already staged
	if !IsStaged(stage, maxinclusive) {
		return
	}

	maxinclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		UnstageBranch(stage, maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMaxLength(maxlength *MaxLength) {

	// check if instance is already staged
	if !IsStaged(stage, maxlength) {
		return
	}

	maxlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		UnstageBranch(stage, maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMinInclusive(mininclusive *MinInclusive) {

	// check if instance is already staged
	if !IsStaged(stage, mininclusive) {
		return
	}

	mininclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		UnstageBranch(stage, mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMinLength(minlength *MinLength) {

	// check if instance is already staged
	if !IsStaged(stage, minlength) {
		return
	}

	minlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		UnstageBranch(stage, minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPattern(pattern *Pattern) {

	// check if instance is already staged
	if !IsStaged(stage, pattern) {
		return
	}

	pattern.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		UnstageBranch(stage, pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRestriction(restriction *Restriction) {

	// check if instance is already staged
	if !IsStaged(stage, restriction) {
		return
	}

	restriction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		UnstageBranch(stage, restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		UnstageBranch(stage, restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		UnstageBranch(stage, restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		UnstageBranch(stage, restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		UnstageBranch(stage, restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		UnstageBranch(stage, restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		UnstageBranch(stage, restriction.MaxLength)
	}
	if restriction.Length != nil {
		UnstageBranch(stage, restriction.Length)
	}
	if restriction.TotalDigit != nil {
		UnstageBranch(stage, restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		UnstageBranch(stage, _enumeration)
	}

}

func (stage *StageStruct) UnstageBranchSchema(schema *Schema) {

	// check if instance is already staged
	if !IsStaged(stage, schema) {
		return
	}

	schema.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		UnstageBranch(stage, schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		UnstageBranch(stage, _simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		UnstageBranch(stage, _complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}
	for _, _group := range schema.Groups {
		UnstageBranch(stage, _group)
	}

}

func (stage *StageStruct) UnstageBranchSequence(sequence *Sequence) {

	// check if instance is already staged
	if !IsStaged(stage, sequence) {
		return
	}

	sequence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		UnstageBranch(stage, sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range sequence.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range sequence.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range sequence.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range sequence.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *StageStruct) UnstageBranchSimpleContent(simplecontent *SimpleContent) {

	// check if instance is already staged
	if !IsStaged(stage, simplecontent) {
		return
	}

	simplecontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		UnstageBranch(stage, simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		UnstageBranch(stage, simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSimpleType(simpletype *SimpleType) {

	// check if instance is already staged
	if !IsStaged(stage, simpletype) {
		return
	}

	simpletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		UnstageBranch(stage, simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		UnstageBranch(stage, simpletype.Restriction)
	}
	if simpletype.Union != nil {
		UnstageBranch(stage, simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTotalDigit(totaldigit *TotalDigit) {

	// check if instance is already staged
	if !IsStaged(stage, totaldigit) {
		return
	}

	totaldigit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		UnstageBranch(stage, totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUnion(union *Union) {

	// check if instance is already staged
	if !IsStaged(stage, union) {
		return
	}

	union.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		UnstageBranch(stage, union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWhiteSpace(whitespace *WhiteSpace) {

	// check if instance is already staged
	if !IsStaged(stage, whitespace) {
		return
	}

	whitespace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		UnstageBranch(stage, whitespace.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
