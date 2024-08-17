// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Annotation:
		// insertion point per field
		if fieldName == "Documentations" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Annotation) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Annotation)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Documentations).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Documentations = _inferedTypeInstance.Documentations[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Documentations =
								append(_inferedTypeInstance.Documentations, any(fieldInstance).(*Documentation))
						}
					}
				}
			}
		}

	case *ComplexType:
		// insertion point per field

	case *Documentation:
		// insertion point per field

	case *Schema:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Annotation
	// insertion point per field
	clear(stage.Annotation_Documentations_reverseMap)
	stage.Annotation_Documentations_reverseMap = make(map[*Documentation]*Annotation)
	for annotation := range stage.Annotations {
		_ = annotation
		for _, _documentation := range annotation.Documentations {
			stage.Annotation_Documentations_reverseMap[_documentation] = annotation
		}
	}

	// Compute reverse map for named struct ComplexType
	// insertion point per field

	// Compute reverse map for named struct Documentation
	// insertion point per field

	// Compute reverse map for named struct Schema
	// insertion point per field

}
