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

	case *Element:
		// insertion point per field

	case *Enumeration:
		// insertion point per field

	case *MaxInclusive:
		// insertion point per field

	case *MinInclusive:
		// insertion point per field

	case *Pattern:
		// insertion point per field

	case *Restriction:
		// insertion point per field
		if fieldName == "Enumerations" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Restriction) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Restriction)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Enumerations).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Enumerations = _inferedTypeInstance.Enumerations[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Enumerations =
								append(_inferedTypeInstance.Enumerations, any(fieldInstance).(*Enumeration))
						}
					}
				}
			}
		}

	case *Schema:
		// insertion point per field
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Schema) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Schema)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Elements).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Elements = _inferedTypeInstance.Elements[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Elements =
								append(_inferedTypeInstance.Elements, any(fieldInstance).(*Element))
						}
					}
				}
			}
		}
		if fieldName == "SimpleTypes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Schema) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Schema)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SimpleTypes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SimpleTypes = _inferedTypeInstance.SimpleTypes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SimpleTypes =
								append(_inferedTypeInstance.SimpleTypes, any(fieldInstance).(*SimpleType))
						}
					}
				}
			}
		}
		if fieldName == "ComplexTypes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Schema) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Schema)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ComplexTypes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ComplexTypes = _inferedTypeInstance.ComplexTypes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ComplexTypes =
								append(_inferedTypeInstance.ComplexTypes, any(fieldInstance).(*ComplexType))
						}
					}
				}
			}
		}

	case *Sequence:
		// insertion point per field
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Sequence) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Sequence)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Elements).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Elements = _inferedTypeInstance.Elements[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Elements =
								append(_inferedTypeInstance.Elements, any(fieldInstance).(*Element))
						}
					}
				}
			}
		}

	case *SimpleType:
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

	// Compute reverse map for named struct Element
	// insertion point per field

	// Compute reverse map for named struct Enumeration
	// insertion point per field

	// Compute reverse map for named struct MaxInclusive
	// insertion point per field

	// Compute reverse map for named struct MinInclusive
	// insertion point per field

	// Compute reverse map for named struct Pattern
	// insertion point per field

	// Compute reverse map for named struct Restriction
	// insertion point per field
	clear(stage.Restriction_Enumerations_reverseMap)
	stage.Restriction_Enumerations_reverseMap = make(map[*Enumeration]*Restriction)
	for restriction := range stage.Restrictions {
		_ = restriction
		for _, _enumeration := range restriction.Enumerations {
			stage.Restriction_Enumerations_reverseMap[_enumeration] = restriction
		}
	}

	// Compute reverse map for named struct Schema
	// insertion point per field
	clear(stage.Schema_Elements_reverseMap)
	stage.Schema_Elements_reverseMap = make(map[*Element]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _element := range schema.Elements {
			stage.Schema_Elements_reverseMap[_element] = schema
		}
	}
	clear(stage.Schema_SimpleTypes_reverseMap)
	stage.Schema_SimpleTypes_reverseMap = make(map[*SimpleType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _simpletype := range schema.SimpleTypes {
			stage.Schema_SimpleTypes_reverseMap[_simpletype] = schema
		}
	}
	clear(stage.Schema_ComplexTypes_reverseMap)
	stage.Schema_ComplexTypes_reverseMap = make(map[*ComplexType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _complextype := range schema.ComplexTypes {
			stage.Schema_ComplexTypes_reverseMap[_complextype] = schema
		}
	}

	// Compute reverse map for named struct Sequence
	// insertion point per field
	clear(stage.Sequence_Elements_reverseMap)
	stage.Sequence_Elements_reverseMap = make(map[*Element]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _element := range sequence.Elements {
			stage.Sequence_Elements_reverseMap[_element] = sequence
		}
	}

	// Compute reverse map for named struct SimpleType
	// insertion point per field

}
