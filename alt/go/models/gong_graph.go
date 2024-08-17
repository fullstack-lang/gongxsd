// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Annotation:
		ok = stage.IsStagedAnnotation(target)

	case *ComplexType:
		ok = stage.IsStagedComplexType(target)

	case *Documentation:
		ok = stage.IsStagedDocumentation(target)

	case *Schema:
		ok = stage.IsStagedSchema(target)

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

func (stage *StageStruct) IsStagedDocumentation(documentation *Documentation) (ok bool) {

	_, ok = stage.Documentations[documentation]

	return
}

func (stage *StageStruct) IsStagedSchema(schema *Schema) (ok bool) {

	_, ok = stage.Schemas[schema]

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

	case *Documentation:
		stage.StageBranchDocumentation(target)

	case *Schema:
		stage.StageBranchSchema(target)

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
	for _, _documentation := range annotation.Documentations {
		StageBranch(stage, _documentation)
	}

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
	if complextype.OuterSchema != nil {
		StageBranch(stage, complextype.OuterSchema)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	if schema.ComplexType != nil {
		StageBranch(stage, schema.ComplexType)
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

	case *Documentation:
		toT := CopyBranchDocumentation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Schema:
		toT := CopyBranchSchema(mapOrigCopy, fromT)
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
	for _, _documentation := range annotationFrom.Documentations {
		annotationTo.Documentations = append(annotationTo.Documentations, CopyBranchDocumentation(mapOrigCopy, _documentation))
	}

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
	if complextypeFrom.OuterSchema != nil {
		complextypeTo.OuterSchema = CopyBranchSchema(mapOrigCopy, complextypeFrom.OuterSchema)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	if schemaFrom.ComplexType != nil {
		schemaTo.ComplexType = CopyBranchComplexType(mapOrigCopy, schemaFrom.ComplexType)
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

	case *Documentation:
		stage.UnstageBranchDocumentation(target)

	case *Schema:
		stage.UnstageBranchSchema(target)

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
	for _, _documentation := range annotation.Documentations {
		UnstageBranch(stage, _documentation)
	}

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
	if complextype.OuterSchema != nil {
		UnstageBranch(stage, complextype.OuterSchema)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	if schema.ComplexType != nil {
		UnstageBranch(stage, schema.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
