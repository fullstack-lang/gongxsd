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
	case *ALTERNATIVE_ID:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_BOOLEAN_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_DATE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_DATE_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_ENUMERATION_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_INTEGER_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_REAL:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_REAL_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_STRING:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_STRING_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_XHTML:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "DEFAULT_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFAULT_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFAULT_VALUE = _inferedTypeInstance.DEFAULT_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFAULT_VALUE =
								append(_inferedTypeInstance.DEFAULT_VALUE, any(fieldInstance).(*A_ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_DATATYPE_DEFINITION_XHTML_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_DATE:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_DATE_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF))
						}
					}
				}
			}
		}
		if fieldName == "VALUES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES = _inferedTypeInstance.VALUES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES =
								append(_inferedTypeInstance.VALUES, any(fieldInstance).(*A_ENUM_VALUE_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_INTEGER:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_INTEGER_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_REAL:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_REAL_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_STRING:
		// insertion point per field
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_STRING_REF))
						}
					}
				}
			}
		}

	case *ATTRIBUTE_VALUE_XHTML:
		// insertion point per field
		if fieldName == "THE_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.THE_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.THE_VALUE = _inferedTypeInstance.THE_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.THE_VALUE =
								append(_inferedTypeInstance.THE_VALUE, any(fieldInstance).(*XHTML_CONTENT))
						}
					}
				}
			}
		}
		if fieldName == "THE_ORIGINAL_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.THE_ORIGINAL_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.THE_ORIGINAL_VALUE = _inferedTypeInstance.THE_ORIGINAL_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.THE_ORIGINAL_VALUE =
								append(_inferedTypeInstance.THE_ORIGINAL_VALUE, any(fieldInstance).(*XHTML_CONTENT))
						}
					}
				}
			}
		}
		if fieldName == "DEFINITION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DEFINITION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DEFINITION = _inferedTypeInstance.DEFINITION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DEFINITION =
								append(_inferedTypeInstance.DEFINITION, any(fieldInstance).(*A_ATTRIBUTE_DEFINITION_XHTML_REF))
						}
					}
				}
			}
		}

	case *A_ALTERNATIVE_ID:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ALTERNATIVE_ID) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ALTERNATIVE_ID)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		// insertion point per field

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_DATE:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_INTEGER:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_REAL:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_STRING:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_XHTML:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		// insertion point per field
		if fieldName == "ATTRIBUTE_VALUE_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN = _inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_VALUE_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_DATE = _inferedTypeInstance.ATTRIBUTE_VALUE_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_DATE =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_DATE, any(fieldInstance).(*ATTRIBUTE_VALUE_DATE))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION = _inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_VALUE_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER = _inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_INTEGER, any(fieldInstance).(*ATTRIBUTE_VALUE_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_REAL = _inferedTypeInstance.ATTRIBUTE_VALUE_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_REAL =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_REAL, any(fieldInstance).(*ATTRIBUTE_VALUE_REAL))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_STRING = _inferedTypeInstance.ATTRIBUTE_VALUE_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_STRING =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_STRING, any(fieldInstance).(*ATTRIBUTE_VALUE_STRING))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_VALUE_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_ATTRIBUTE_VALUE_XHTML_1)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML = _inferedTypeInstance.ATTRIBUTE_VALUE_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML =
								append(_inferedTypeInstance.ATTRIBUTE_VALUE_XHTML, any(fieldInstance).(*ATTRIBUTE_VALUE_XHTML))
						}
					}
				}
			}
		}

	case *A_CHILDREN:
		// insertion point per field
		if fieldName == "SPEC_HIERARCHY" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_CHILDREN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_CHILDREN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_HIERARCHY).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_HIERARCHY = _inferedTypeInstance.SPEC_HIERARCHY[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_HIERARCHY =
								append(_inferedTypeInstance.SPEC_HIERARCHY, any(fieldInstance).(*SPEC_HIERARCHY))
						}
					}
				}
			}
		}

	case *A_CORE_CONTENT:
		// insertion point per field
		if fieldName == "REQ_IF_CONTENT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_CORE_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_CORE_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.REQ_IF_CONTENT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.REQ_IF_CONTENT = _inferedTypeInstance.REQ_IF_CONTENT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.REQ_IF_CONTENT =
								append(_inferedTypeInstance.REQ_IF_CONTENT, any(fieldInstance).(*REQ_IF_CONTENT))
						}
					}
				}
			}
		}

	case *A_DATATYPES:
		// insertion point per field
		if fieldName == "DATATYPE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_BOOLEAN = _inferedTypeInstance.DATATYPE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_BOOLEAN, any(fieldInstance).(*DATATYPE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_DATE = _inferedTypeInstance.DATATYPE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_DATE =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_DATE, any(fieldInstance).(*DATATYPE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_ENUMERATION = _inferedTypeInstance.DATATYPE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_ENUMERATION, any(fieldInstance).(*DATATYPE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_INTEGER = _inferedTypeInstance.DATATYPE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_INTEGER, any(fieldInstance).(*DATATYPE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_REAL = _inferedTypeInstance.DATATYPE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_REAL =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_REAL, any(fieldInstance).(*DATATYPE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_STRING = _inferedTypeInstance.DATATYPE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_STRING =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_STRING, any(fieldInstance).(*DATATYPE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "DATATYPE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_DATATYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_DATATYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPE_DEFINITION_XHTML = _inferedTypeInstance.DATATYPE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPE_DEFINITION_XHTML =
								append(_inferedTypeInstance.DATATYPE_DEFINITION_XHTML, any(fieldInstance).(*DATATYPE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_DATE_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_REAL_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_STRING_REF:
		// insertion point per field

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		// insertion point per field

	case *A_EDITABLE_ATTS:
		// insertion point per field

	case *A_ENUM_VALUE_REF:
		// insertion point per field

	case *A_OBJECT:
		// insertion point per field

	case *A_PROPERTIES:
		// insertion point per field
		if fieldName == "EMBEDDED_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_PROPERTIES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_PROPERTIES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.EMBEDDED_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.EMBEDDED_VALUE = _inferedTypeInstance.EMBEDDED_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.EMBEDDED_VALUE =
								append(_inferedTypeInstance.EMBEDDED_VALUE, any(fieldInstance).(*EMBEDDED_VALUE))
						}
					}
				}
			}
		}

	case *A_RELATION_GROUP_TYPE_REF:
		// insertion point per field

	case *A_SOURCE:
		// insertion point per field

	case *A_SOURCE_SPECIFICATION:
		// insertion point per field

	case *A_SPECIFICATIONS:
		// insertion point per field
		if fieldName == "SPECIFICATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPECIFICATIONS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPECIFICATIONS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATION = _inferedTypeInstance.SPECIFICATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATION =
								append(_inferedTypeInstance.SPECIFICATION, any(fieldInstance).(*SPECIFICATION))
						}
					}
				}
			}
		}

	case *A_SPECIFICATION_TYPE_REF:
		// insertion point per field

	case *A_SPECIFIED_VALUES:
		// insertion point per field
		if fieldName == "ENUM_VALUE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPECIFIED_VALUES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPECIFIED_VALUES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ENUM_VALUE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ENUM_VALUE = _inferedTypeInstance.ENUM_VALUE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ENUM_VALUE =
								append(_inferedTypeInstance.ENUM_VALUE, any(fieldInstance).(*ENUM_VALUE))
						}
					}
				}
			}
		}

	case *A_SPEC_ATTRIBUTES:
		// insertion point per field
		if fieldName == "ATTRIBUTE_DEFINITION_BOOLEAN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_BOOLEAN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_BOOLEAN = _inferedTypeInstance.ATTRIBUTE_DEFINITION_BOOLEAN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_BOOLEAN =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_BOOLEAN, any(fieldInstance).(*ATTRIBUTE_DEFINITION_BOOLEAN))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_DATE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_DATE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_DATE = _inferedTypeInstance.ATTRIBUTE_DEFINITION_DATE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_DATE =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_DATE, any(fieldInstance).(*ATTRIBUTE_DEFINITION_DATE))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_ENUMERATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_ENUMERATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_ENUMERATION = _inferedTypeInstance.ATTRIBUTE_DEFINITION_ENUMERATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_ENUMERATION =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_ENUMERATION, any(fieldInstance).(*ATTRIBUTE_DEFINITION_ENUMERATION))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_INTEGER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_INTEGER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_INTEGER = _inferedTypeInstance.ATTRIBUTE_DEFINITION_INTEGER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_INTEGER =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_INTEGER, any(fieldInstance).(*ATTRIBUTE_DEFINITION_INTEGER))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_REAL" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_REAL).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_REAL = _inferedTypeInstance.ATTRIBUTE_DEFINITION_REAL[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_REAL =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_REAL, any(fieldInstance).(*ATTRIBUTE_DEFINITION_REAL))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_STRING" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_STRING).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_STRING = _inferedTypeInstance.ATTRIBUTE_DEFINITION_STRING[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_STRING =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_STRING, any(fieldInstance).(*ATTRIBUTE_DEFINITION_STRING))
						}
					}
				}
			}
		}
		if fieldName == "ATTRIBUTE_DEFINITION_XHTML" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_ATTRIBUTES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_ATTRIBUTES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ATTRIBUTE_DEFINITION_XHTML).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ATTRIBUTE_DEFINITION_XHTML = _inferedTypeInstance.ATTRIBUTE_DEFINITION_XHTML[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ATTRIBUTE_DEFINITION_XHTML =
								append(_inferedTypeInstance.ATTRIBUTE_DEFINITION_XHTML, any(fieldInstance).(*ATTRIBUTE_DEFINITION_XHTML))
						}
					}
				}
			}
		}

	case *A_SPEC_OBJECTS:
		// insertion point per field
		if fieldName == "SPEC_OBJECT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_OBJECTS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_OBJECTS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_OBJECT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_OBJECT = _inferedTypeInstance.SPEC_OBJECT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_OBJECT =
								append(_inferedTypeInstance.SPEC_OBJECT, any(fieldInstance).(*SPEC_OBJECT))
						}
					}
				}
			}
		}

	case *A_SPEC_OBJECT_TYPE_REF:
		// insertion point per field

	case *A_SPEC_RELATIONS:
		// insertion point per field
		if fieldName == "SPEC_RELATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_RELATIONS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_RELATIONS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATION = _inferedTypeInstance.SPEC_RELATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATION =
								append(_inferedTypeInstance.SPEC_RELATION, any(fieldInstance).(*SPEC_RELATION))
						}
					}
				}
			}
		}

	case *A_SPEC_RELATION_GROUPS:
		// insertion point per field
		if fieldName == "RELATION_GROUP" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_RELATION_GROUPS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_RELATION_GROUPS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RELATION_GROUP).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RELATION_GROUP = _inferedTypeInstance.RELATION_GROUP[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RELATION_GROUP =
								append(_inferedTypeInstance.RELATION_GROUP, any(fieldInstance).(*RELATION_GROUP))
						}
					}
				}
			}
		}

	case *A_SPEC_RELATION_REF:
		// insertion point per field

	case *A_SPEC_RELATION_TYPE_REF:
		// insertion point per field

	case *A_SPEC_TYPES:
		// insertion point per field
		if fieldName == "RELATION_GROUP_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_TYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_TYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RELATION_GROUP_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RELATION_GROUP_TYPE = _inferedTypeInstance.RELATION_GROUP_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RELATION_GROUP_TYPE =
								append(_inferedTypeInstance.RELATION_GROUP_TYPE, any(fieldInstance).(*RELATION_GROUP_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_OBJECT_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_TYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_TYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_OBJECT_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_OBJECT_TYPE = _inferedTypeInstance.SPEC_OBJECT_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_OBJECT_TYPE =
								append(_inferedTypeInstance.SPEC_OBJECT_TYPE, any(fieldInstance).(*SPEC_OBJECT_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATION_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_TYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_TYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATION_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATION_TYPE = _inferedTypeInstance.SPEC_RELATION_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATION_TYPE =
								append(_inferedTypeInstance.SPEC_RELATION_TYPE, any(fieldInstance).(*SPEC_RELATION_TYPE))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFICATION_TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_SPEC_TYPES) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_SPEC_TYPES)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATION_TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATION_TYPE = _inferedTypeInstance.SPECIFICATION_TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATION_TYPE =
								append(_inferedTypeInstance.SPECIFICATION_TYPE, any(fieldInstance).(*SPECIFICATION_TYPE))
						}
					}
				}
			}
		}

	case *A_THE_HEADER:
		// insertion point per field
		if fieldName == "REQ_IF_HEADER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_THE_HEADER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_THE_HEADER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.REQ_IF_HEADER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.REQ_IF_HEADER = _inferedTypeInstance.REQ_IF_HEADER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.REQ_IF_HEADER =
								append(_inferedTypeInstance.REQ_IF_HEADER, any(fieldInstance).(*REQ_IF_HEADER))
						}
					}
				}
			}
		}

	case *A_TOOL_EXTENSIONS:
		// insertion point per field
		if fieldName == "REQ_IF_TOOL_EXTENSION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*A_TOOL_EXTENSIONS) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*A_TOOL_EXTENSIONS)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.REQ_IF_TOOL_EXTENSION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.REQ_IF_TOOL_EXTENSION = _inferedTypeInstance.REQ_IF_TOOL_EXTENSION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.REQ_IF_TOOL_EXTENSION =
								append(_inferedTypeInstance.REQ_IF_TOOL_EXTENSION, any(fieldInstance).(*REQ_IF_TOOL_EXTENSION))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_BOOLEAN:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_BOOLEAN) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_BOOLEAN)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_DATE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_DATE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_DATE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_ENUMERATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFIED_VALUES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_ENUMERATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_ENUMERATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFIED_VALUES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFIED_VALUES = _inferedTypeInstance.SPECIFIED_VALUES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFIED_VALUES =
								append(_inferedTypeInstance.SPECIFIED_VALUES, any(fieldInstance).(*A_SPECIFIED_VALUES))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_INTEGER:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_INTEGER) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_INTEGER)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_REAL:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_REAL) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_REAL)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_STRING:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_STRING) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_STRING)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *DATATYPE_DEFINITION_XHTML:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DATATYPE_DEFINITION_XHTML) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DATATYPE_DEFINITION_XHTML)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}

	case *EMBEDDED_VALUE:
		// insertion point per field

	case *ENUM_VALUE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ENUM_VALUE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ENUM_VALUE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "PROPERTIES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*ENUM_VALUE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*ENUM_VALUE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.PROPERTIES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.PROPERTIES = _inferedTypeInstance.PROPERTIES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.PROPERTIES =
								append(_inferedTypeInstance.PROPERTIES, any(fieldInstance).(*A_PROPERTIES))
						}
					}
				}
			}
		}

	case *RELATION_GROUP:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SOURCE_SPECIFICATION" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SOURCE_SPECIFICATION).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SOURCE_SPECIFICATION = _inferedTypeInstance.SOURCE_SPECIFICATION[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SOURCE_SPECIFICATION =
								append(_inferedTypeInstance.SOURCE_SPECIFICATION, any(fieldInstance).(*A_SOURCE_SPECIFICATION))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATIONS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATIONS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATIONS = _inferedTypeInstance.SPEC_RELATIONS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATIONS =
								append(_inferedTypeInstance.SPEC_RELATIONS, any(fieldInstance).(*A_SPEC_RELATION_REF))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_RELATION_GROUP_TYPE_REF))
						}
					}
				}
			}
		}

	case *RELATION_GROUP_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*RELATION_GROUP_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*RELATION_GROUP_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES = _inferedTypeInstance.SPEC_ATTRIBUTES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES, any(fieldInstance).(*A_SPEC_ATTRIBUTES))
						}
					}
				}
			}
		}

	case *REQ_IF:
		// insertion point per field
		if fieldName == "THE_HEADER" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.THE_HEADER).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.THE_HEADER = _inferedTypeInstance.THE_HEADER[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.THE_HEADER =
								append(_inferedTypeInstance.THE_HEADER, any(fieldInstance).(*A_THE_HEADER))
						}
					}
				}
			}
		}
		if fieldName == "CORE_CONTENT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CORE_CONTENT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CORE_CONTENT = _inferedTypeInstance.CORE_CONTENT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CORE_CONTENT =
								append(_inferedTypeInstance.CORE_CONTENT, any(fieldInstance).(*A_CORE_CONTENT))
						}
					}
				}
			}
		}
		if fieldName == "TOOL_EXTENSIONS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TOOL_EXTENSIONS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TOOL_EXTENSIONS = _inferedTypeInstance.TOOL_EXTENSIONS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TOOL_EXTENSIONS =
								append(_inferedTypeInstance.TOOL_EXTENSIONS, any(fieldInstance).(*A_TOOL_EXTENSIONS))
						}
					}
				}
			}
		}

	case *REQ_IF_CONTENT:
		// insertion point per field
		if fieldName == "DATATYPES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DATATYPES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DATATYPES = _inferedTypeInstance.DATATYPES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DATATYPES =
								append(_inferedTypeInstance.DATATYPES, any(fieldInstance).(*A_DATATYPES))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_TYPES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_TYPES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_TYPES = _inferedTypeInstance.SPEC_TYPES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_TYPES =
								append(_inferedTypeInstance.SPEC_TYPES, any(fieldInstance).(*A_SPEC_TYPES))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_OBJECTS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_OBJECTS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_OBJECTS = _inferedTypeInstance.SPEC_OBJECTS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_OBJECTS =
								append(_inferedTypeInstance.SPEC_OBJECTS, any(fieldInstance).(*A_SPEC_OBJECTS))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATIONS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATIONS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATIONS = _inferedTypeInstance.SPEC_RELATIONS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATIONS =
								append(_inferedTypeInstance.SPEC_RELATIONS, any(fieldInstance).(*A_SPEC_RELATIONS))
						}
					}
				}
			}
		}
		if fieldName == "SPECIFICATIONS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPECIFICATIONS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPECIFICATIONS = _inferedTypeInstance.SPECIFICATIONS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPECIFICATIONS =
								append(_inferedTypeInstance.SPECIFICATIONS, any(fieldInstance).(*A_SPECIFICATIONS))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_RELATION_GROUPS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*REQ_IF_CONTENT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*REQ_IF_CONTENT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_RELATION_GROUPS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_RELATION_GROUPS = _inferedTypeInstance.SPEC_RELATION_GROUPS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_RELATION_GROUPS =
								append(_inferedTypeInstance.SPEC_RELATION_GROUPS, any(fieldInstance).(*A_SPEC_RELATION_GROUPS))
						}
					}
				}
			}
		}

	case *REQ_IF_HEADER:
		// insertion point per field

	case *REQ_IF_TOOL_EXTENSION:
		// insertion point per field

	case *SPECIFICATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "CHILDREN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CHILDREN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CHILDREN = _inferedTypeInstance.CHILDREN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CHILDREN =
								append(_inferedTypeInstance.CHILDREN, any(fieldInstance).(*A_CHILDREN))
						}
					}
				}
			}
		}
		if fieldName == "VALUES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES = _inferedTypeInstance.VALUES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES =
								append(_inferedTypeInstance.VALUES, any(fieldInstance).(*A_ATTRIBUTE_VALUE_XHTML_1))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_SPECIFICATION_TYPE_REF))
						}
					}
				}
			}
		}

	case *SPECIFICATION_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPECIFICATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPECIFICATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES = _inferedTypeInstance.SPEC_ATTRIBUTES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES, any(fieldInstance).(*A_SPEC_ATTRIBUTES))
						}
					}
				}
			}
		}

	case *SPEC_HIERARCHY:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "CHILDREN" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CHILDREN).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CHILDREN = _inferedTypeInstance.CHILDREN[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CHILDREN =
								append(_inferedTypeInstance.CHILDREN, any(fieldInstance).(*A_CHILDREN))
						}
					}
				}
			}
		}
		if fieldName == "EDITABLE_ATTS" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.EDITABLE_ATTS).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.EDITABLE_ATTS = _inferedTypeInstance.EDITABLE_ATTS[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.EDITABLE_ATTS =
								append(_inferedTypeInstance.EDITABLE_ATTS, any(fieldInstance).(*A_EDITABLE_ATTS))
						}
					}
				}
			}
		}
		if fieldName == "OBJECT" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_HIERARCHY) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_HIERARCHY)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.OBJECT).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.OBJECT = _inferedTypeInstance.OBJECT[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.OBJECT =
								append(_inferedTypeInstance.OBJECT, any(fieldInstance).(*A_OBJECT))
						}
					}
				}
			}
		}

	case *SPEC_OBJECT:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "VALUES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES = _inferedTypeInstance.VALUES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES =
								append(_inferedTypeInstance.VALUES, any(fieldInstance).(*A_ATTRIBUTE_VALUE_XHTML_1))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_SPEC_OBJECT_TYPE_REF))
						}
					}
				}
			}
		}

	case *SPEC_OBJECT_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_OBJECT_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_OBJECT_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES = _inferedTypeInstance.SPEC_ATTRIBUTES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES, any(fieldInstance).(*A_SPEC_ATTRIBUTES))
						}
					}
				}
			}
		}

	case *SPEC_RELATION:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "VALUES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.VALUES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.VALUES = _inferedTypeInstance.VALUES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.VALUES =
								append(_inferedTypeInstance.VALUES, any(fieldInstance).(*A_ATTRIBUTE_VALUE_XHTML_1))
						}
					}
				}
			}
		}
		if fieldName == "SOURCE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SOURCE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SOURCE = _inferedTypeInstance.SOURCE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SOURCE =
								append(_inferedTypeInstance.SOURCE, any(fieldInstance).(*A_SOURCE))
						}
					}
				}
			}
		}
		if fieldName == "TYPE" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TYPE).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TYPE = _inferedTypeInstance.TYPE[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TYPE =
								append(_inferedTypeInstance.TYPE, any(fieldInstance).(*A_SPEC_RELATION_TYPE_REF))
						}
					}
				}
			}
		}

	case *SPEC_RELATION_TYPE:
		// insertion point per field
		if fieldName == "ALTERNATIVE_ID" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.ALTERNATIVE_ID).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.ALTERNATIVE_ID = _inferedTypeInstance.ALTERNATIVE_ID[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.ALTERNATIVE_ID =
								append(_inferedTypeInstance.ALTERNATIVE_ID, any(fieldInstance).(*A_ALTERNATIVE_ID))
						}
					}
				}
			}
		}
		if fieldName == "SPEC_ATTRIBUTES" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*SPEC_RELATION_TYPE) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*SPEC_RELATION_TYPE)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SPEC_ATTRIBUTES).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SPEC_ATTRIBUTES = _inferedTypeInstance.SPEC_ATTRIBUTES[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SPEC_ATTRIBUTES =
								append(_inferedTypeInstance.SPEC_ATTRIBUTES, any(fieldInstance).(*A_SPEC_ATTRIBUTES))
						}
					}
				}
			}
		}

	case *XHTML_CONTENT:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _a_alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_boolean
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _a_attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_reverseMap[_a_attribute_value_boolean] = attribute_definition_boolean
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _a_datatype_definition_boolean_ref := range attribute_definition_boolean.TYPE {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_TYPE_reverseMap[_a_datatype_definition_boolean_ref] = attribute_definition_boolean
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _a_alternative_id := range attribute_definition_date.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_date
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _a_attribute_value_date := range attribute_definition_date.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_reverseMap[_a_attribute_value_date] = attribute_definition_date
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_DATE_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_DATE_REF]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _a_datatype_definition_date_ref := range attribute_definition_date.TYPE {
			stage.ATTRIBUTE_DEFINITION_DATE_TYPE_reverseMap[_a_datatype_definition_date_ref] = attribute_definition_date
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _a_alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_enumeration
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _a_attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_reverseMap[_a_attribute_value_enumeration] = attribute_definition_enumeration
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _a_datatype_definition_enumeration_ref := range attribute_definition_enumeration.TYPE {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_TYPE_reverseMap[_a_datatype_definition_enumeration_ref] = attribute_definition_enumeration
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _a_alternative_id := range attribute_definition_integer.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_integer
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _a_attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_reverseMap[_a_attribute_value_integer] = attribute_definition_integer
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _a_datatype_definition_integer_ref := range attribute_definition_integer.TYPE {
			stage.ATTRIBUTE_DEFINITION_INTEGER_TYPE_reverseMap[_a_datatype_definition_integer_ref] = attribute_definition_integer
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _a_alternative_id := range attribute_definition_real.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_real
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _a_attribute_value_real := range attribute_definition_real.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_reverseMap[_a_attribute_value_real] = attribute_definition_real
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_REAL_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_REAL_REF]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _a_datatype_definition_real_ref := range attribute_definition_real.TYPE {
			stage.ATTRIBUTE_DEFINITION_REAL_TYPE_reverseMap[_a_datatype_definition_real_ref] = attribute_definition_real
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _a_alternative_id := range attribute_definition_string.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_string
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _a_attribute_value_string := range attribute_definition_string.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_reverseMap[_a_attribute_value_string] = attribute_definition_string
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_STRING_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_STRING_REF]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _a_datatype_definition_string_ref := range attribute_definition_string.TYPE {
			stage.ATTRIBUTE_DEFINITION_STRING_TYPE_reverseMap[_a_datatype_definition_string_ref] = attribute_definition_string
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _a_alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = attribute_definition_xhtml
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_reverseMap = make(map[*A_ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _a_attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE {
			stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_reverseMap[_a_attribute_value_xhtml] = attribute_definition_xhtml
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_TYPE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_TYPE_reverseMap = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _a_datatype_definition_xhtml_ref := range attribute_definition_xhtml.TYPE {
			stage.ATTRIBUTE_DEFINITION_XHTML_TYPE_reverseMap[_a_datatype_definition_xhtml_ref] = attribute_definition_xhtml
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_BOOLEAN_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_BOOLEAN_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]*ATTRIBUTE_VALUE_BOOLEAN)
	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		_ = attribute_value_boolean
		for _, _a_attribute_definition_boolean_ref := range attribute_value_boolean.DEFINITION {
			stage.ATTRIBUTE_VALUE_BOOLEAN_DEFINITION_reverseMap[_a_attribute_definition_boolean_ref] = attribute_value_boolean
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_DATE
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_DATE_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_DATE_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]*ATTRIBUTE_VALUE_DATE)
	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		_ = attribute_value_date
		for _, _a_attribute_definition_date_ref := range attribute_value_date.DEFINITION {
			stage.ATTRIBUTE_VALUE_DATE_DEFINITION_reverseMap[_a_attribute_definition_date_ref] = attribute_value_date
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_ENUMERATION_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_ENUMERATION_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]*ATTRIBUTE_VALUE_ENUMERATION)
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = attribute_value_enumeration
		for _, _a_attribute_definition_enumeration_ref := range attribute_value_enumeration.DEFINITION {
			stage.ATTRIBUTE_VALUE_ENUMERATION_DEFINITION_reverseMap[_a_attribute_definition_enumeration_ref] = attribute_value_enumeration
		}
	}
	clear(stage.ATTRIBUTE_VALUE_ENUMERATION_VALUES_reverseMap)
	stage.ATTRIBUTE_VALUE_ENUMERATION_VALUES_reverseMap = make(map[*A_ENUM_VALUE_REF]*ATTRIBUTE_VALUE_ENUMERATION)
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = attribute_value_enumeration
		for _, _a_enum_value_ref := range attribute_value_enumeration.VALUES {
			stage.ATTRIBUTE_VALUE_ENUMERATION_VALUES_reverseMap[_a_enum_value_ref] = attribute_value_enumeration
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_INTEGER
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_INTEGER_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_INTEGER_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]*ATTRIBUTE_VALUE_INTEGER)
	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		_ = attribute_value_integer
		for _, _a_attribute_definition_integer_ref := range attribute_value_integer.DEFINITION {
			stage.ATTRIBUTE_VALUE_INTEGER_DEFINITION_reverseMap[_a_attribute_definition_integer_ref] = attribute_value_integer
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_REAL
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_REAL_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_REAL_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]*ATTRIBUTE_VALUE_REAL)
	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		_ = attribute_value_real
		for _, _a_attribute_definition_real_ref := range attribute_value_real.DEFINITION {
			stage.ATTRIBUTE_VALUE_REAL_DEFINITION_reverseMap[_a_attribute_definition_real_ref] = attribute_value_real
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_STRING
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_STRING_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_STRING_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]*ATTRIBUTE_VALUE_STRING)
	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		_ = attribute_value_string
		for _, _a_attribute_definition_string_ref := range attribute_value_string.DEFINITION {
			stage.ATTRIBUTE_VALUE_STRING_DEFINITION_reverseMap[_a_attribute_definition_string_ref] = attribute_value_string
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}
	clear(stage.ATTRIBUTE_VALUE_XHTML_DEFINITION_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_DEFINITION_reverseMap = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _a_attribute_definition_xhtml_ref := range attribute_value_xhtml.DEFINITION {
			stage.ATTRIBUTE_VALUE_XHTML_DEFINITION_reverseMap[_a_attribute_definition_xhtml_ref] = attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct A_ALTERNATIVE_ID
	// insertion point per field
	clear(stage.A_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.A_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*A_ALTERNATIVE_ID)
	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		_ = a_alternative_id
		for _, _alternative_id := range a_alternative_id.ALTERNATIVE_ID {
			stage.A_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = a_alternative_id
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_XHTML_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		_ = a_attribute_value_boolean
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_boolean
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_DATE
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		_ = a_attribute_value_date
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_date
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = a_attribute_value_enumeration
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_enumeration
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_INTEGER
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		_ = a_attribute_value_integer
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_integer
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_REAL
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		_ = a_attribute_value_real
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_real
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_STRING
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		_ = a_attribute_value_string
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_string
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		_ = a_attribute_value_xhtml
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML_1
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml_1
		}
	}

	// Compute reverse map for named struct A_CHILDREN
	// insertion point per field
	clear(stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap)
	stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*A_CHILDREN)
	for a_children := range stage.A_CHILDRENs {
		_ = a_children
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = a_children
		}
	}

	// Compute reverse map for named struct A_CORE_CONTENT
	// insertion point per field
	clear(stage.A_CORE_CONTENT_REQ_IF_CONTENT_reverseMap)
	stage.A_CORE_CONTENT_REQ_IF_CONTENT_reverseMap = make(map[*REQ_IF_CONTENT]*A_CORE_CONTENT)
	for a_core_content := range stage.A_CORE_CONTENTs {
		_ = a_core_content
		for _, _req_if_content := range a_core_content.REQ_IF_CONTENT {
			stage.A_CORE_CONTENT_REQ_IF_CONTENT_reverseMap[_req_if_content] = a_core_content
		}
	}

	// Compute reverse map for named struct A_DATATYPES
	// insertion point per field
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap = make(map[*DATATYPE_DEFINITION_XHTML]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[_datatype_definition_xhtml] = a_datatypes
		}
	}

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_XHTML_REF
	// insertion point per field

	// Compute reverse map for named struct A_EDITABLE_ATTS
	// insertion point per field

	// Compute reverse map for named struct A_ENUM_VALUE_REF
	// insertion point per field

	// Compute reverse map for named struct A_OBJECT
	// insertion point per field

	// Compute reverse map for named struct A_PROPERTIES
	// insertion point per field
	clear(stage.A_PROPERTIES_EMBEDDED_VALUE_reverseMap)
	stage.A_PROPERTIES_EMBEDDED_VALUE_reverseMap = make(map[*EMBEDDED_VALUE]*A_PROPERTIES)
	for a_properties := range stage.A_PROPERTIESs {
		_ = a_properties
		for _, _embedded_value := range a_properties.EMBEDDED_VALUE {
			stage.A_PROPERTIES_EMBEDDED_VALUE_reverseMap[_embedded_value] = a_properties
		}
	}

	// Compute reverse map for named struct A_RELATION_GROUP_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFICATIONS
	// insertion point per field
	clear(stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap)
	stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*A_SPECIFICATIONS)
	for a_specifications := range stage.A_SPECIFICATIONSs {
		_ = a_specifications
		for _, _specification := range a_specifications.SPECIFICATION {
			stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = a_specifications
		}
	}

	// Compute reverse map for named struct A_SPECIFICATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFIED_VALUES
	// insertion point per field
	clear(stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap)
	stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*A_SPECIFIED_VALUES)
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		_ = a_specified_values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = a_specified_values
		}
	}

	// Compute reverse map for named struct A_SPEC_ATTRIBUTES
	// insertion point per field
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = a_spec_attributes
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECTS
	// insertion point per field
	clear(stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap)
	stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap = make(map[*SPEC_OBJECT]*A_SPEC_OBJECTS)
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		_ = a_spec_objects
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[_spec_object] = a_spec_objects
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECT_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATIONS
	// insertion point per field
	clear(stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap)
	stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*A_SPEC_RELATIONS)
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		_ = a_spec_relations
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = a_spec_relations
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_GROUPS
	// insertion point per field
	clear(stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap)
	stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap = make(map[*RELATION_GROUP]*A_SPEC_RELATION_GROUPS)
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		_ = a_spec_relation_groups
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[_relation_group] = a_spec_relation_groups
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_TYPES
	// insertion point per field
	clear(stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap)
	stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap)
	stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap)
	stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap)
	stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap = make(map[*SPECIFICATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[_specification_type] = a_spec_types
		}
	}

	// Compute reverse map for named struct A_THE_HEADER
	// insertion point per field
	clear(stage.A_THE_HEADER_REQ_IF_HEADER_reverseMap)
	stage.A_THE_HEADER_REQ_IF_HEADER_reverseMap = make(map[*REQ_IF_HEADER]*A_THE_HEADER)
	for a_the_header := range stage.A_THE_HEADERs {
		_ = a_the_header
		for _, _req_if_header := range a_the_header.REQ_IF_HEADER {
			stage.A_THE_HEADER_REQ_IF_HEADER_reverseMap[_req_if_header] = a_the_header
		}
	}

	// Compute reverse map for named struct A_TOOL_EXTENSIONS
	// insertion point per field
	clear(stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap)
	stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*A_TOOL_EXTENSIONS)
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		_ = a_tool_extensions
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = a_tool_extensions
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_BOOLEAN)
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		_ = datatype_definition_boolean
		for _, _a_alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_boolean
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_DATE)
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		_ = datatype_definition_date
		for _, _a_alternative_id := range datatype_definition_date.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_date
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _a_alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_enumeration
		}
	}
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_reverseMap = make(map[*A_SPECIFIED_VALUES]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _a_specified_values := range datatype_definition_enumeration.SPECIFIED_VALUES {
			stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_reverseMap[_a_specified_values] = datatype_definition_enumeration
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_INTEGER)
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		_ = datatype_definition_integer
		for _, _a_alternative_id := range datatype_definition_integer.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_integer
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_REAL)
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		_ = datatype_definition_real
		for _, _a_alternative_id := range datatype_definition_real.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_real
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_STRING)
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		_ = datatype_definition_string
		for _, _a_alternative_id := range datatype_definition_string.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_string
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*DATATYPE_DEFINITION_XHTML)
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		_ = datatype_definition_xhtml
		for _, _a_alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = datatype_definition_xhtml
		}
	}

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field
	clear(stage.ENUM_VALUE_ALTERNATIVE_ID_reverseMap)
	stage.ENUM_VALUE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _a_alternative_id := range enum_value.ALTERNATIVE_ID {
			stage.ENUM_VALUE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = enum_value
		}
	}
	clear(stage.ENUM_VALUE_PROPERTIES_reverseMap)
	stage.ENUM_VALUE_PROPERTIES_reverseMap = make(map[*A_PROPERTIES]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _a_properties := range enum_value.PROPERTIES {
			stage.ENUM_VALUE_PROPERTIES_reverseMap[_a_properties] = enum_value
		}
	}

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field
	clear(stage.RELATION_GROUP_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _a_alternative_id := range relation_group.ALTERNATIVE_ID {
			stage.RELATION_GROUP_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = relation_group
		}
	}
	clear(stage.RELATION_GROUP_SOURCE_SPECIFICATION_reverseMap)
	stage.RELATION_GROUP_SOURCE_SPECIFICATION_reverseMap = make(map[*A_SOURCE_SPECIFICATION]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _a_source_specification := range relation_group.SOURCE_SPECIFICATION {
			stage.RELATION_GROUP_SOURCE_SPECIFICATION_reverseMap[_a_source_specification] = relation_group
		}
	}
	clear(stage.RELATION_GROUP_SPEC_RELATIONS_reverseMap)
	stage.RELATION_GROUP_SPEC_RELATIONS_reverseMap = make(map[*A_SPEC_RELATION_REF]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _a_spec_relation_ref := range relation_group.SPEC_RELATIONS {
			stage.RELATION_GROUP_SPEC_RELATIONS_reverseMap[_a_spec_relation_ref] = relation_group
		}
	}
	clear(stage.RELATION_GROUP_TYPE_reverseMap)
	stage.RELATION_GROUP_TYPE_reverseMap = make(map[*A_RELATION_GROUP_TYPE_REF]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _a_relation_group_type_ref := range relation_group.TYPE {
			stage.RELATION_GROUP_TYPE_reverseMap[_a_relation_group_type_ref] = relation_group
		}
	}

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field
	clear(stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _a_alternative_id := range relation_group_type.ALTERNATIVE_ID {
			stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_reverseMap = make(map[*A_SPEC_ATTRIBUTES]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _a_spec_attributes := range relation_group_type.SPEC_ATTRIBUTES {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_reverseMap[_a_spec_attributes] = relation_group_type
		}
	}

	// Compute reverse map for named struct REQ_IF
	// insertion point per field
	clear(stage.REQ_IF_THE_HEADER_reverseMap)
	stage.REQ_IF_THE_HEADER_reverseMap = make(map[*A_THE_HEADER]*REQ_IF)
	for req_if := range stage.REQ_IFs {
		_ = req_if
		for _, _a_the_header := range req_if.THE_HEADER {
			stage.REQ_IF_THE_HEADER_reverseMap[_a_the_header] = req_if
		}
	}
	clear(stage.REQ_IF_CORE_CONTENT_reverseMap)
	stage.REQ_IF_CORE_CONTENT_reverseMap = make(map[*A_CORE_CONTENT]*REQ_IF)
	for req_if := range stage.REQ_IFs {
		_ = req_if
		for _, _a_core_content := range req_if.CORE_CONTENT {
			stage.REQ_IF_CORE_CONTENT_reverseMap[_a_core_content] = req_if
		}
	}
	clear(stage.REQ_IF_TOOL_EXTENSIONS_reverseMap)
	stage.REQ_IF_TOOL_EXTENSIONS_reverseMap = make(map[*A_TOOL_EXTENSIONS]*REQ_IF)
	for req_if := range stage.REQ_IFs {
		_ = req_if
		for _, _a_tool_extensions := range req_if.TOOL_EXTENSIONS {
			stage.REQ_IF_TOOL_EXTENSIONS_reverseMap[_a_tool_extensions] = req_if
		}
	}

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field
	clear(stage.REQ_IF_CONTENT_DATATYPES_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_reverseMap = make(map[*A_DATATYPES]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_datatypes := range req_if_content.DATATYPES {
			stage.REQ_IF_CONTENT_DATATYPES_reverseMap[_a_datatypes] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_reverseMap = make(map[*A_SPEC_TYPES]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_spec_types := range req_if_content.SPEC_TYPES {
			stage.REQ_IF_CONTENT_SPEC_TYPES_reverseMap[_a_spec_types] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_OBJECTS_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_OBJECTS_reverseMap = make(map[*A_SPEC_OBJECTS]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_spec_objects := range req_if_content.SPEC_OBJECTS {
			stage.REQ_IF_CONTENT_SPEC_OBJECTS_reverseMap[_a_spec_objects] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATIONS_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATIONS_reverseMap = make(map[*A_SPEC_RELATIONS]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_spec_relations := range req_if_content.SPEC_RELATIONS {
			stage.REQ_IF_CONTENT_SPEC_RELATIONS_reverseMap[_a_spec_relations] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap)
	stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap = make(map[*A_SPECIFICATIONS]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_specifications := range req_if_content.SPECIFICATIONS {
			stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[_a_specifications] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_reverseMap = make(map[*A_SPEC_RELATION_GROUPS]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _a_spec_relation_groups := range req_if_content.SPEC_RELATION_GROUPS {
			stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_reverseMap[_a_spec_relation_groups] = req_if_content
		}
	}

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field
	clear(stage.SPECIFICATION_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _a_alternative_id := range specification.ALTERNATIVE_ID {
			stage.SPECIFICATION_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = specification
		}
	}
	clear(stage.SPECIFICATION_CHILDREN_reverseMap)
	stage.SPECIFICATION_CHILDREN_reverseMap = make(map[*A_CHILDREN]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _a_children := range specification.CHILDREN {
			stage.SPECIFICATION_CHILDREN_reverseMap[_a_children] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_reverseMap)
	stage.SPECIFICATION_VALUES_reverseMap = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _a_attribute_value_xhtml_1 := range specification.VALUES {
			stage.SPECIFICATION_VALUES_reverseMap[_a_attribute_value_xhtml_1] = specification
		}
	}
	clear(stage.SPECIFICATION_TYPE_reverseMap)
	stage.SPECIFICATION_TYPE_reverseMap = make(map[*A_SPECIFICATION_TYPE_REF]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _a_specification_type_ref := range specification.TYPE {
			stage.SPECIFICATION_TYPE_reverseMap[_a_specification_type_ref] = specification
		}
	}

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field
	clear(stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _a_alternative_id := range specification_type.ALTERNATIVE_ID {
			stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_reverseMap = make(map[*A_SPEC_ATTRIBUTES]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _a_spec_attributes := range specification_type.SPEC_ATTRIBUTES {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_reverseMap[_a_spec_attributes] = specification_type
		}
	}

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field
	clear(stage.SPEC_HIERARCHY_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_HIERARCHY_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _a_alternative_id := range spec_hierarchy.ALTERNATIVE_ID {
			stage.SPEC_HIERARCHY_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = spec_hierarchy
		}
	}
	clear(stage.SPEC_HIERARCHY_CHILDREN_reverseMap)
	stage.SPEC_HIERARCHY_CHILDREN_reverseMap = make(map[*A_CHILDREN]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _a_children := range spec_hierarchy.CHILDREN {
			stage.SPEC_HIERARCHY_CHILDREN_reverseMap[_a_children] = spec_hierarchy
		}
	}
	clear(stage.SPEC_HIERARCHY_EDITABLE_ATTS_reverseMap)
	stage.SPEC_HIERARCHY_EDITABLE_ATTS_reverseMap = make(map[*A_EDITABLE_ATTS]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _a_editable_atts := range spec_hierarchy.EDITABLE_ATTS {
			stage.SPEC_HIERARCHY_EDITABLE_ATTS_reverseMap[_a_editable_atts] = spec_hierarchy
		}
	}
	clear(stage.SPEC_HIERARCHY_OBJECT_reverseMap)
	stage.SPEC_HIERARCHY_OBJECT_reverseMap = make(map[*A_OBJECT]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _a_object := range spec_hierarchy.OBJECT {
			stage.SPEC_HIERARCHY_OBJECT_reverseMap[_a_object] = spec_hierarchy
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field
	clear(stage.SPEC_OBJECT_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _a_alternative_id := range spec_object.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_reverseMap)
	stage.SPEC_OBJECT_VALUES_reverseMap = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _a_attribute_value_xhtml_1 := range spec_object.VALUES {
			stage.SPEC_OBJECT_VALUES_reverseMap[_a_attribute_value_xhtml_1] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_reverseMap)
	stage.SPEC_OBJECT_TYPE_reverseMap = make(map[*A_SPEC_OBJECT_TYPE_REF]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _a_spec_object_type_ref := range spec_object.TYPE {
			stage.SPEC_OBJECT_TYPE_reverseMap[_a_spec_object_type_ref] = spec_object
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field
	clear(stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _a_alternative_id := range spec_object_type.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_reverseMap = make(map[*A_SPEC_ATTRIBUTES]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _a_spec_attributes := range spec_object_type.SPEC_ATTRIBUTES {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_reverseMap[_a_spec_attributes] = spec_object_type
		}
	}

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field
	clear(stage.SPEC_RELATION_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _a_alternative_id := range spec_relation.ALTERNATIVE_ID {
			stage.SPEC_RELATION_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_reverseMap)
	stage.SPEC_RELATION_VALUES_reverseMap = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _a_attribute_value_xhtml_1 := range spec_relation.VALUES {
			stage.SPEC_RELATION_VALUES_reverseMap[_a_attribute_value_xhtml_1] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_SOURCE_reverseMap)
	stage.SPEC_RELATION_SOURCE_reverseMap = make(map[*A_SOURCE]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _a_source := range spec_relation.SOURCE {
			stage.SPEC_RELATION_SOURCE_reverseMap[_a_source] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_TYPE_reverseMap)
	stage.SPEC_RELATION_TYPE_reverseMap = make(map[*A_SPEC_RELATION_TYPE_REF]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _a_spec_relation_type_ref := range spec_relation.TYPE {
			stage.SPEC_RELATION_TYPE_reverseMap[_a_spec_relation_type_ref] = spec_relation
		}
	}

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field
	clear(stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_reverseMap = make(map[*A_ALTERNATIVE_ID]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _a_alternative_id := range spec_relation_type.ALTERNATIVE_ID {
			stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_reverseMap[_a_alternative_id] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_reverseMap = make(map[*A_SPEC_ATTRIBUTES]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _a_spec_attributes := range spec_relation_type.SPEC_ATTRIBUTES {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_reverseMap[_a_spec_attributes] = spec_relation_type
		}
	}

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

}
