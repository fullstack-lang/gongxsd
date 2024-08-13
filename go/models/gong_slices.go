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
	case *All:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*All) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*All)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*All) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*All)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*All) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*All)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*All) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*All)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*All) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*All)
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

	case *Attribute:
		// insertion point per field

	case *AttributeGroup:
		// insertion point per field
		if fieldName == "AttributeGroups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*AttributeGroup) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*AttributeGroup)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.AttributeGroups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.AttributeGroups = _inferedTypeInstance.AttributeGroups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.AttributeGroups =
								append(_inferedTypeInstance.AttributeGroups, any(fieldInstance).(*AttributeGroup))
						}
					}
				}
			}
		}
		if fieldName == "Attributes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*AttributeGroup) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*AttributeGroup)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Attributes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Attributes = _inferedTypeInstance.Attributes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Attributes =
								append(_inferedTypeInstance.Attributes, any(fieldInstance).(*Attribute))
						}
					}
				}
			}
		}

	case *Choice:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Choice) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Choice)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Choice) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Choice)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Choice) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Choice)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Choice) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Choice)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Choice) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Choice)
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

	case *ComplexContent:
		// insertion point per field

	case *ComplexType:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
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
		if fieldName == "Attributes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Attributes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Attributes = _inferedTypeInstance.Attributes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Attributes =
								append(_inferedTypeInstance.Attributes, any(fieldInstance).(*Attribute))
						}
					}
				}
			}
		}
		if fieldName == "AttributeGroups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ComplexType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ComplexType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.AttributeGroups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.AttributeGroups = _inferedTypeInstance.AttributeGroups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.AttributeGroups =
								append(_inferedTypeInstance.AttributeGroups, any(fieldInstance).(*AttributeGroup))
						}
					}
				}
			}
		}

	case *Documentation:
		// insertion point per field

	case *Element:
		// insertion point per field
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Element) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Element)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}

	case *Enumeration:
		// insertion point per field

	case *Extension:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Extension) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Extension)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Extension) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Extension)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Extension) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Extension)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Extension) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Extension)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Extension) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Extension)
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

	case *Group:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Elements" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
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

	case *Length:
		// insertion point per field

	case *MaxInclusive:
		// insertion point per field

	case *MaxLength:
		// insertion point per field

	case *MinInclusive:
		// insertion point per field

	case *MinLength:
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
		if fieldName == "AttributeGroups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Schema) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Schema)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.AttributeGroups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.AttributeGroups = _inferedTypeInstance.AttributeGroups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.AttributeGroups =
								append(_inferedTypeInstance.AttributeGroups, any(fieldInstance).(*AttributeGroup))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Schema) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Schema)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}

	case *Sequence:
		// insertion point per field
		if fieldName == "Sequences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Sequence) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Sequence)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sequences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sequences = _inferedTypeInstance.Sequences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sequences =
								append(_inferedTypeInstance.Sequences, any(fieldInstance).(*Sequence))
						}
					}
				}
			}
		}
		if fieldName == "Alls" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Sequence) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Sequence)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Alls).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Alls = _inferedTypeInstance.Alls[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Alls =
								append(_inferedTypeInstance.Alls, any(fieldInstance).(*All))
						}
					}
				}
			}
		}
		if fieldName == "Choices" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Sequence) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Sequence)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Choices).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Choices = _inferedTypeInstance.Choices[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Choices =
								append(_inferedTypeInstance.Choices, any(fieldInstance).(*Choice))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Sequence) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Sequence)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
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

	case *SimpleContent:
		// insertion point per field

	case *SimpleType:
		// insertion point per field

	case *TotalDigit:
		// insertion point per field

	case *Union:
		// insertion point per field

	case *WhiteSpace:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct All
	// insertion point per field
	clear(stage.All_Sequences_reverseMap)
	stage.All_Sequences_reverseMap = make(map[*Sequence]*All)
	for all := range stage.Alls {
		_ = all
		for _, _sequence := range all.Sequences {
			stage.All_Sequences_reverseMap[_sequence] = all
		}
	}
	clear(stage.All_Alls_reverseMap)
	stage.All_Alls_reverseMap = make(map[*All]*All)
	for all := range stage.Alls {
		_ = all
		for _, _all := range all.Alls {
			stage.All_Alls_reverseMap[_all] = all
		}
	}
	clear(stage.All_Choices_reverseMap)
	stage.All_Choices_reverseMap = make(map[*Choice]*All)
	for all := range stage.Alls {
		_ = all
		for _, _choice := range all.Choices {
			stage.All_Choices_reverseMap[_choice] = all
		}
	}
	clear(stage.All_Groups_reverseMap)
	stage.All_Groups_reverseMap = make(map[*Group]*All)
	for all := range stage.Alls {
		_ = all
		for _, _group := range all.Groups {
			stage.All_Groups_reverseMap[_group] = all
		}
	}
	clear(stage.All_Elements_reverseMap)
	stage.All_Elements_reverseMap = make(map[*Element]*All)
	for all := range stage.Alls {
		_ = all
		for _, _element := range all.Elements {
			stage.All_Elements_reverseMap[_element] = all
		}
	}

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

	// Compute reverse map for named struct Attribute
	// insertion point per field

	// Compute reverse map for named struct AttributeGroup
	// insertion point per field
	clear(stage.AttributeGroup_AttributeGroups_reverseMap)
	stage.AttributeGroup_AttributeGroups_reverseMap = make(map[*AttributeGroup]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attributegroup := range attributegroup.AttributeGroups {
			stage.AttributeGroup_AttributeGroups_reverseMap[_attributegroup] = attributegroup
		}
	}
	clear(stage.AttributeGroup_Attributes_reverseMap)
	stage.AttributeGroup_Attributes_reverseMap = make(map[*Attribute]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attribute := range attributegroup.Attributes {
			stage.AttributeGroup_Attributes_reverseMap[_attribute] = attributegroup
		}
	}

	// Compute reverse map for named struct Choice
	// insertion point per field
	clear(stage.Choice_Sequences_reverseMap)
	stage.Choice_Sequences_reverseMap = make(map[*Sequence]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _sequence := range choice.Sequences {
			stage.Choice_Sequences_reverseMap[_sequence] = choice
		}
	}
	clear(stage.Choice_Alls_reverseMap)
	stage.Choice_Alls_reverseMap = make(map[*All]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _all := range choice.Alls {
			stage.Choice_Alls_reverseMap[_all] = choice
		}
	}
	clear(stage.Choice_Choices_reverseMap)
	stage.Choice_Choices_reverseMap = make(map[*Choice]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _choice := range choice.Choices {
			stage.Choice_Choices_reverseMap[_choice] = choice
		}
	}
	clear(stage.Choice_Groups_reverseMap)
	stage.Choice_Groups_reverseMap = make(map[*Group]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _group := range choice.Groups {
			stage.Choice_Groups_reverseMap[_group] = choice
		}
	}
	clear(stage.Choice_Elements_reverseMap)
	stage.Choice_Elements_reverseMap = make(map[*Element]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _element := range choice.Elements {
			stage.Choice_Elements_reverseMap[_element] = choice
		}
	}

	// Compute reverse map for named struct ComplexContent
	// insertion point per field

	// Compute reverse map for named struct ComplexType
	// insertion point per field
	clear(stage.ComplexType_Sequences_reverseMap)
	stage.ComplexType_Sequences_reverseMap = make(map[*Sequence]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _sequence := range complextype.Sequences {
			stage.ComplexType_Sequences_reverseMap[_sequence] = complextype
		}
	}
	clear(stage.ComplexType_Alls_reverseMap)
	stage.ComplexType_Alls_reverseMap = make(map[*All]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _all := range complextype.Alls {
			stage.ComplexType_Alls_reverseMap[_all] = complextype
		}
	}
	clear(stage.ComplexType_Choices_reverseMap)
	stage.ComplexType_Choices_reverseMap = make(map[*Choice]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _choice := range complextype.Choices {
			stage.ComplexType_Choices_reverseMap[_choice] = complextype
		}
	}
	clear(stage.ComplexType_Groups_reverseMap)
	stage.ComplexType_Groups_reverseMap = make(map[*Group]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _group := range complextype.Groups {
			stage.ComplexType_Groups_reverseMap[_group] = complextype
		}
	}
	clear(stage.ComplexType_Elements_reverseMap)
	stage.ComplexType_Elements_reverseMap = make(map[*Element]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _element := range complextype.Elements {
			stage.ComplexType_Elements_reverseMap[_element] = complextype
		}
	}
	clear(stage.ComplexType_Attributes_reverseMap)
	stage.ComplexType_Attributes_reverseMap = make(map[*Attribute]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attribute := range complextype.Attributes {
			stage.ComplexType_Attributes_reverseMap[_attribute] = complextype
		}
	}
	clear(stage.ComplexType_AttributeGroups_reverseMap)
	stage.ComplexType_AttributeGroups_reverseMap = make(map[*AttributeGroup]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attributegroup := range complextype.AttributeGroups {
			stage.ComplexType_AttributeGroups_reverseMap[_attributegroup] = complextype
		}
	}

	// Compute reverse map for named struct Documentation
	// insertion point per field

	// Compute reverse map for named struct Element
	// insertion point per field
	clear(stage.Element_Groups_reverseMap)
	stage.Element_Groups_reverseMap = make(map[*Group]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _group := range element.Groups {
			stage.Element_Groups_reverseMap[_group] = element
		}
	}

	// Compute reverse map for named struct Enumeration
	// insertion point per field

	// Compute reverse map for named struct Extension
	// insertion point per field
	clear(stage.Extension_Sequences_reverseMap)
	stage.Extension_Sequences_reverseMap = make(map[*Sequence]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _sequence := range extension.Sequences {
			stage.Extension_Sequences_reverseMap[_sequence] = extension
		}
	}
	clear(stage.Extension_Alls_reverseMap)
	stage.Extension_Alls_reverseMap = make(map[*All]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _all := range extension.Alls {
			stage.Extension_Alls_reverseMap[_all] = extension
		}
	}
	clear(stage.Extension_Choices_reverseMap)
	stage.Extension_Choices_reverseMap = make(map[*Choice]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _choice := range extension.Choices {
			stage.Extension_Choices_reverseMap[_choice] = extension
		}
	}
	clear(stage.Extension_Groups_reverseMap)
	stage.Extension_Groups_reverseMap = make(map[*Group]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _group := range extension.Groups {
			stage.Extension_Groups_reverseMap[_group] = extension
		}
	}
	clear(stage.Extension_Elements_reverseMap)
	stage.Extension_Elements_reverseMap = make(map[*Element]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _element := range extension.Elements {
			stage.Extension_Elements_reverseMap[_element] = extension
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	clear(stage.Group_Sequences_reverseMap)
	stage.Group_Sequences_reverseMap = make(map[*Sequence]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _sequence := range group.Sequences {
			stage.Group_Sequences_reverseMap[_sequence] = group
		}
	}
	clear(stage.Group_Alls_reverseMap)
	stage.Group_Alls_reverseMap = make(map[*All]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _all := range group.Alls {
			stage.Group_Alls_reverseMap[_all] = group
		}
	}
	clear(stage.Group_Choices_reverseMap)
	stage.Group_Choices_reverseMap = make(map[*Choice]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _choice := range group.Choices {
			stage.Group_Choices_reverseMap[_choice] = group
		}
	}
	clear(stage.Group_Groups_reverseMap)
	stage.Group_Groups_reverseMap = make(map[*Group]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _group := range group.Groups {
			stage.Group_Groups_reverseMap[_group] = group
		}
	}
	clear(stage.Group_Elements_reverseMap)
	stage.Group_Elements_reverseMap = make(map[*Element]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _element := range group.Elements {
			stage.Group_Elements_reverseMap[_element] = group
		}
	}

	// Compute reverse map for named struct Length
	// insertion point per field

	// Compute reverse map for named struct MaxInclusive
	// insertion point per field

	// Compute reverse map for named struct MaxLength
	// insertion point per field

	// Compute reverse map for named struct MinInclusive
	// insertion point per field

	// Compute reverse map for named struct MinLength
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
	clear(stage.Schema_AttributeGroups_reverseMap)
	stage.Schema_AttributeGroups_reverseMap = make(map[*AttributeGroup]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _attributegroup := range schema.AttributeGroups {
			stage.Schema_AttributeGroups_reverseMap[_attributegroup] = schema
		}
	}
	clear(stage.Schema_Groups_reverseMap)
	stage.Schema_Groups_reverseMap = make(map[*Group]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _group := range schema.Groups {
			stage.Schema_Groups_reverseMap[_group] = schema
		}
	}

	// Compute reverse map for named struct Sequence
	// insertion point per field
	clear(stage.Sequence_Sequences_reverseMap)
	stage.Sequence_Sequences_reverseMap = make(map[*Sequence]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _sequence := range sequence.Sequences {
			stage.Sequence_Sequences_reverseMap[_sequence] = sequence
		}
	}
	clear(stage.Sequence_Alls_reverseMap)
	stage.Sequence_Alls_reverseMap = make(map[*All]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _all := range sequence.Alls {
			stage.Sequence_Alls_reverseMap[_all] = sequence
		}
	}
	clear(stage.Sequence_Choices_reverseMap)
	stage.Sequence_Choices_reverseMap = make(map[*Choice]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _choice := range sequence.Choices {
			stage.Sequence_Choices_reverseMap[_choice] = sequence
		}
	}
	clear(stage.Sequence_Groups_reverseMap)
	stage.Sequence_Groups_reverseMap = make(map[*Group]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _group := range sequence.Groups {
			stage.Sequence_Groups_reverseMap[_group] = sequence
		}
	}
	clear(stage.Sequence_Elements_reverseMap)
	stage.Sequence_Elements_reverseMap = make(map[*Element]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _element := range sequence.Elements {
			stage.Sequence_Elements_reverseMap[_element] = sequence
		}
	}

	// Compute reverse map for named struct SimpleContent
	// insertion point per field

	// Compute reverse map for named struct SimpleType
	// insertion point per field

	// Compute reverse map for named struct TotalDigit
	// insertion point per field

	// Compute reverse map for named struct Union
	// insertion point per field

	// Compute reverse map for named struct WhiteSpace
	// insertion point per field

}
