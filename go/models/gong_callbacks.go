// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationCreateCallback != nil {
			stage.OnAfterAnnotationCreateCallback.OnAfterCreate(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeCreateCallback != nil {
			stage.OnAfterComplexTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationCreateCallback != nil {
			stage.OnAfterDocumentationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Element:
		if stage.OnAfterElementCreateCallback != nil {
			stage.OnAfterElementCreateCallback.OnAfterCreate(stage, target)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationCreateCallback != nil {
			stage.OnAfterEnumerationCreateCallback.OnAfterCreate(stage, target)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveCreateCallback != nil {
			stage.OnAfterMaxInclusiveCreateCallback.OnAfterCreate(stage, target)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveCreateCallback != nil {
			stage.OnAfterMinInclusiveCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pattern:
		if stage.OnAfterPatternCreateCallback != nil {
			stage.OnAfterPatternCreateCallback.OnAfterCreate(stage, target)
		}
	case *Restriction:
		if stage.OnAfterRestrictionCreateCallback != nil {
			stage.OnAfterRestrictionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaCreateCallback != nil {
			stage.OnAfterSchemaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Sequence:
		if stage.OnAfterSequenceCreateCallback != nil {
			stage.OnAfterSequenceCreateCallback.OnAfterCreate(stage, target)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeCreateCallback != nil {
			stage.OnAfterSimpleTypeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Annotation:
		newTarget := any(new).(*Annotation)
		if stage.OnAfterAnnotationUpdateCallback != nil {
			stage.OnAfterAnnotationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ComplexType:
		newTarget := any(new).(*ComplexType)
		if stage.OnAfterComplexTypeUpdateCallback != nil {
			stage.OnAfterComplexTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Documentation:
		newTarget := any(new).(*Documentation)
		if stage.OnAfterDocumentationUpdateCallback != nil {
			stage.OnAfterDocumentationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Element:
		newTarget := any(new).(*Element)
		if stage.OnAfterElementUpdateCallback != nil {
			stage.OnAfterElementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Enumeration:
		newTarget := any(new).(*Enumeration)
		if stage.OnAfterEnumerationUpdateCallback != nil {
			stage.OnAfterEnumerationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MaxInclusive:
		newTarget := any(new).(*MaxInclusive)
		if stage.OnAfterMaxInclusiveUpdateCallback != nil {
			stage.OnAfterMaxInclusiveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MinInclusive:
		newTarget := any(new).(*MinInclusive)
		if stage.OnAfterMinInclusiveUpdateCallback != nil {
			stage.OnAfterMinInclusiveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pattern:
		newTarget := any(new).(*Pattern)
		if stage.OnAfterPatternUpdateCallback != nil {
			stage.OnAfterPatternUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Restriction:
		newTarget := any(new).(*Restriction)
		if stage.OnAfterRestrictionUpdateCallback != nil {
			stage.OnAfterRestrictionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Schema:
		newTarget := any(new).(*Schema)
		if stage.OnAfterSchemaUpdateCallback != nil {
			stage.OnAfterSchemaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Sequence:
		newTarget := any(new).(*Sequence)
		if stage.OnAfterSequenceUpdateCallback != nil {
			stage.OnAfterSequenceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SimpleType:
		newTarget := any(new).(*SimpleType)
		if stage.OnAfterSimpleTypeUpdateCallback != nil {
			stage.OnAfterSimpleTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationDeleteCallback != nil {
			staged := any(staged).(*Annotation)
			stage.OnAfterAnnotationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeDeleteCallback != nil {
			staged := any(staged).(*ComplexType)
			stage.OnAfterComplexTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Documentation:
		if stage.OnAfterDocumentationDeleteCallback != nil {
			staged := any(staged).(*Documentation)
			stage.OnAfterDocumentationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Element:
		if stage.OnAfterElementDeleteCallback != nil {
			staged := any(staged).(*Element)
			stage.OnAfterElementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationDeleteCallback != nil {
			staged := any(staged).(*Enumeration)
			stage.OnAfterEnumerationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveDeleteCallback != nil {
			staged := any(staged).(*MaxInclusive)
			stage.OnAfterMaxInclusiveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveDeleteCallback != nil {
			staged := any(staged).(*MinInclusive)
			stage.OnAfterMinInclusiveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pattern:
		if stage.OnAfterPatternDeleteCallback != nil {
			staged := any(staged).(*Pattern)
			stage.OnAfterPatternDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Restriction:
		if stage.OnAfterRestrictionDeleteCallback != nil {
			staged := any(staged).(*Restriction)
			stage.OnAfterRestrictionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Schema:
		if stage.OnAfterSchemaDeleteCallback != nil {
			staged := any(staged).(*Schema)
			stage.OnAfterSchemaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Sequence:
		if stage.OnAfterSequenceDeleteCallback != nil {
			staged := any(staged).(*Sequence)
			stage.OnAfterSequenceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeDeleteCallback != nil {
			staged := any(staged).(*SimpleType)
			stage.OnAfterSimpleTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationReadCallback != nil {
			stage.OnAfterAnnotationReadCallback.OnAfterRead(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeReadCallback != nil {
			stage.OnAfterComplexTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationReadCallback != nil {
			stage.OnAfterDocumentationReadCallback.OnAfterRead(stage, target)
		}
	case *Element:
		if stage.OnAfterElementReadCallback != nil {
			stage.OnAfterElementReadCallback.OnAfterRead(stage, target)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationReadCallback != nil {
			stage.OnAfterEnumerationReadCallback.OnAfterRead(stage, target)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveReadCallback != nil {
			stage.OnAfterMaxInclusiveReadCallback.OnAfterRead(stage, target)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveReadCallback != nil {
			stage.OnAfterMinInclusiveReadCallback.OnAfterRead(stage, target)
		}
	case *Pattern:
		if stage.OnAfterPatternReadCallback != nil {
			stage.OnAfterPatternReadCallback.OnAfterRead(stage, target)
		}
	case *Restriction:
		if stage.OnAfterRestrictionReadCallback != nil {
			stage.OnAfterRestrictionReadCallback.OnAfterRead(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaReadCallback != nil {
			stage.OnAfterSchemaReadCallback.OnAfterRead(stage, target)
		}
	case *Sequence:
		if stage.OnAfterSequenceReadCallback != nil {
			stage.OnAfterSequenceReadCallback.OnAfterRead(stage, target)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeReadCallback != nil {
			stage.OnAfterSimpleTypeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationUpdateCallback = any(callback).(OnAfterUpdateInterface[Annotation])
	
	case *ComplexType:
		stage.OnAfterComplexTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationUpdateCallback = any(callback).(OnAfterUpdateInterface[Documentation])
	
	case *Element:
		stage.OnAfterElementUpdateCallback = any(callback).(OnAfterUpdateInterface[Element])
	
	case *Enumeration:
		stage.OnAfterEnumerationUpdateCallback = any(callback).(OnAfterUpdateInterface[Enumeration])
	
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveUpdateCallback = any(callback).(OnAfterUpdateInterface[MaxInclusive])
	
	case *MinInclusive:
		stage.OnAfterMinInclusiveUpdateCallback = any(callback).(OnAfterUpdateInterface[MinInclusive])
	
	case *Pattern:
		stage.OnAfterPatternUpdateCallback = any(callback).(OnAfterUpdateInterface[Pattern])
	
	case *Restriction:
		stage.OnAfterRestrictionUpdateCallback = any(callback).(OnAfterUpdateInterface[Restriction])
	
	case *Schema:
		stage.OnAfterSchemaUpdateCallback = any(callback).(OnAfterUpdateInterface[Schema])
	
	case *Sequence:
		stage.OnAfterSequenceUpdateCallback = any(callback).(OnAfterUpdateInterface[Sequence])
	
	case *SimpleType:
		stage.OnAfterSimpleTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[SimpleType])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationCreateCallback = any(callback).(OnAfterCreateInterface[Annotation])
	
	case *ComplexType:
		stage.OnAfterComplexTypeCreateCallback = any(callback).(OnAfterCreateInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationCreateCallback = any(callback).(OnAfterCreateInterface[Documentation])
	
	case *Element:
		stage.OnAfterElementCreateCallback = any(callback).(OnAfterCreateInterface[Element])
	
	case *Enumeration:
		stage.OnAfterEnumerationCreateCallback = any(callback).(OnAfterCreateInterface[Enumeration])
	
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveCreateCallback = any(callback).(OnAfterCreateInterface[MaxInclusive])
	
	case *MinInclusive:
		stage.OnAfterMinInclusiveCreateCallback = any(callback).(OnAfterCreateInterface[MinInclusive])
	
	case *Pattern:
		stage.OnAfterPatternCreateCallback = any(callback).(OnAfterCreateInterface[Pattern])
	
	case *Restriction:
		stage.OnAfterRestrictionCreateCallback = any(callback).(OnAfterCreateInterface[Restriction])
	
	case *Schema:
		stage.OnAfterSchemaCreateCallback = any(callback).(OnAfterCreateInterface[Schema])
	
	case *Sequence:
		stage.OnAfterSequenceCreateCallback = any(callback).(OnAfterCreateInterface[Sequence])
	
	case *SimpleType:
		stage.OnAfterSimpleTypeCreateCallback = any(callback).(OnAfterCreateInterface[SimpleType])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationDeleteCallback = any(callback).(OnAfterDeleteInterface[Annotation])
	
	case *ComplexType:
		stage.OnAfterComplexTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationDeleteCallback = any(callback).(OnAfterDeleteInterface[Documentation])
	
	case *Element:
		stage.OnAfterElementDeleteCallback = any(callback).(OnAfterDeleteInterface[Element])
	
	case *Enumeration:
		stage.OnAfterEnumerationDeleteCallback = any(callback).(OnAfterDeleteInterface[Enumeration])
	
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveDeleteCallback = any(callback).(OnAfterDeleteInterface[MaxInclusive])
	
	case *MinInclusive:
		stage.OnAfterMinInclusiveDeleteCallback = any(callback).(OnAfterDeleteInterface[MinInclusive])
	
	case *Pattern:
		stage.OnAfterPatternDeleteCallback = any(callback).(OnAfterDeleteInterface[Pattern])
	
	case *Restriction:
		stage.OnAfterRestrictionDeleteCallback = any(callback).(OnAfterDeleteInterface[Restriction])
	
	case *Schema:
		stage.OnAfterSchemaDeleteCallback = any(callback).(OnAfterDeleteInterface[Schema])
	
	case *Sequence:
		stage.OnAfterSequenceDeleteCallback = any(callback).(OnAfterDeleteInterface[Sequence])
	
	case *SimpleType:
		stage.OnAfterSimpleTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[SimpleType])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationReadCallback = any(callback).(OnAfterReadInterface[Annotation])
	
	case *ComplexType:
		stage.OnAfterComplexTypeReadCallback = any(callback).(OnAfterReadInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationReadCallback = any(callback).(OnAfterReadInterface[Documentation])
	
	case *Element:
		stage.OnAfterElementReadCallback = any(callback).(OnAfterReadInterface[Element])
	
	case *Enumeration:
		stage.OnAfterEnumerationReadCallback = any(callback).(OnAfterReadInterface[Enumeration])
	
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveReadCallback = any(callback).(OnAfterReadInterface[MaxInclusive])
	
	case *MinInclusive:
		stage.OnAfterMinInclusiveReadCallback = any(callback).(OnAfterReadInterface[MinInclusive])
	
	case *Pattern:
		stage.OnAfterPatternReadCallback = any(callback).(OnAfterReadInterface[Pattern])
	
	case *Restriction:
		stage.OnAfterRestrictionReadCallback = any(callback).(OnAfterReadInterface[Restriction])
	
	case *Schema:
		stage.OnAfterSchemaReadCallback = any(callback).(OnAfterReadInterface[Schema])
	
	case *Sequence:
		stage.OnAfterSequenceReadCallback = any(callback).(OnAfterReadInterface[Sequence])
	
	case *SimpleType:
		stage.OnAfterSimpleTypeReadCallback = any(callback).(OnAfterReadInterface[SimpleType])
	
	}
}
