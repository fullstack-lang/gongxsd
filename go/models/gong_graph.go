// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Annotation:
		ok = stage.IsStagedAnnotation(target)

	case *ComplexType:
		ok = stage.IsStagedComplexType(target)

	case *Element:
		ok = stage.IsStagedElement(target)

	case *Enumeration:
		ok = stage.IsStagedEnumeration(target)

	case *MaxInclusive:
		ok = stage.IsStagedMaxInclusive(target)

	case *MinInclusive:
		ok = stage.IsStagedMinInclusive(target)

	case *Pattern:
		ok = stage.IsStagedPattern(target)

	case *Restriction:
		ok = stage.IsStagedRestriction(target)

	case *Schema:
		ok = stage.IsStagedSchema(target)

	case *Sequence:
		ok = stage.IsStagedSequence(target)

	case *SimpleType:
		ok = stage.IsStagedSimpleType(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAnnotation(annotation *Annotation) (ok bool) {

	_, ok = stage.Annotations[annotation]

	return
}

func (stage *StageStruct) IsStagedComplexType(complextype *ComplexType) (ok bool) {

	_, ok = stage.ComplexTypes[complextype]

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

func (stage *StageStruct) IsStagedMaxInclusive(maxinclusive *MaxInclusive) (ok bool) {

	_, ok = stage.MaxInclusives[maxinclusive]

	return
}

func (stage *StageStruct) IsStagedMinInclusive(mininclusive *MinInclusive) (ok bool) {

	_, ok = stage.MinInclusives[mininclusive]

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

func (stage *StageStruct) IsStagedSimpleType(simpletype *SimpleType) (ok bool) {

	_, ok = stage.SimpleTypes[simpletype]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Annotation:
		stage.StageBranchAnnotation(target)

	case *ComplexType:
		stage.StageBranchComplexType(target)

	case *Element:
		stage.StageBranchElement(target)

	case *Enumeration:
		stage.StageBranchEnumeration(target)

	case *MaxInclusive:
		stage.StageBranchMaxInclusive(target)

	case *MinInclusive:
		stage.StageBranchMinInclusive(target)

	case *Pattern:
		stage.StageBranchPattern(target)

	case *Restriction:
		stage.StageBranchRestriction(target)

	case *Schema:
		stage.StageBranchSchema(target)

	case *Sequence:
		stage.StageBranchSequence(target)

	case *SimpleType:
		stage.StageBranchSimpleType(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if IsStaged(stage, annotation) {
		return
	}

	annotation.Stage(stage)

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
	if complextype.Annotation != nil {
		StageBranch(stage, complextype.Annotation)
	}
	if complextype.Sequence != nil {
		StageBranch(stage, complextype.Sequence)
	}

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
	for _, _element := range sequence.Elements {
		StageBranch(stage, _element)
	}

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
	case *Annotation:
		toT := CopyBranchAnnotation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexType:
		toT := CopyBranchComplexType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Element:
		toT := CopyBranchElement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Enumeration:
		toT := CopyBranchEnumeration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxInclusive:
		toT := CopyBranchMaxInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinInclusive:
		toT := CopyBranchMinInclusive(mapOrigCopy, fromT)
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

	case *SimpleType:
		toT := CopyBranchSimpleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	if complextypeFrom.Annotation != nil {
		complextypeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, complextypeFrom.Annotation)
	}
	if complextypeFrom.Sequence != nil {
		complextypeTo.Sequence = CopyBranchSequence(mapOrigCopy, complextypeFrom.Sequence)
	}

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
	for _, _element := range sequenceFrom.Elements {
		sequenceTo.Elements = append(sequenceTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

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
	case *Annotation:
		stage.UnstageBranchAnnotation(target)

	case *ComplexType:
		stage.UnstageBranchComplexType(target)

	case *Element:
		stage.UnstageBranchElement(target)

	case *Enumeration:
		stage.UnstageBranchEnumeration(target)

	case *MaxInclusive:
		stage.UnstageBranchMaxInclusive(target)

	case *MinInclusive:
		stage.UnstageBranchMinInclusive(target)

	case *Pattern:
		stage.UnstageBranchPattern(target)

	case *Restriction:
		stage.UnstageBranchRestriction(target)

	case *Schema:
		stage.UnstageBranchSchema(target)

	case *Sequence:
		stage.UnstageBranchSequence(target)

	case *SimpleType:
		stage.UnstageBranchSimpleType(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if !IsStaged(stage, annotation) {
		return
	}

	annotation.Unstage(stage)

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
	if complextype.Annotation != nil {
		UnstageBranch(stage, complextype.Annotation)
	}
	if complextype.Sequence != nil {
		UnstageBranch(stage, complextype.Sequence)
	}

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
	for _, _element := range sequence.Elements {
		UnstageBranch(stage, _element)
	}

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

	//insertion point for the staging of instances referenced by slice of pointers

}
