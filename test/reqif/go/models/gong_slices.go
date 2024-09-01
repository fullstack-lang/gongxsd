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

	case *ATTRIBUTE_DEFINITION_DATE:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_REAL:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_STRING:
		// insertion point per field

	case *ATTRIBUTE_DEFINITION_XHTML:
		// insertion point per field

	case *ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point per field

	case *ATTRIBUTE_VALUE_DATE:
		// insertion point per field

	case *ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point per field

	case *ATTRIBUTE_VALUE_INTEGER:
		// insertion point per field

	case *ATTRIBUTE_VALUE_REAL:
		// insertion point per field

	case *ATTRIBUTE_VALUE_STRING:
		// insertion point per field

	case *ATTRIBUTE_VALUE_XHTML:
		// insertion point per field

	case *A_ALTERNATIVE_ID:
		// insertion point per field

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

	case *A_ATTRIBUTE_VALUE_DATE:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_INTEGER:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_REAL:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_STRING:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_XHTML:
		// insertion point per field

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		// insertion point per field

	case *A_CHILDREN:
		// insertion point per field

	case *A_CORE_CONTENT:
		// insertion point per field

	case *A_DATATYPES:
		// insertion point per field

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

	case *A_RELATION_GROUP_TYPE_REF:
		// insertion point per field

	case *A_SOURCE_SPECIFICATION_1:
		// insertion point per field

	case *A_SPECIFICATIONS:
		// insertion point per field

	case *A_SPECIFICATION_TYPE_REF:
		// insertion point per field

	case *A_SPECIFIED_VALUES:
		// insertion point per field

	case *A_SPEC_ATTRIBUTES:
		// insertion point per field

	case *A_SPEC_OBJECTS:
		// insertion point per field

	case *A_SPEC_OBJECT_TYPE_REF:
		// insertion point per field

	case *A_SPEC_RELATIONS:
		// insertion point per field

	case *A_SPEC_RELATION_GROUPS:
		// insertion point per field

	case *A_SPEC_RELATION_REF:
		// insertion point per field

	case *A_SPEC_RELATION_TYPE_REF:
		// insertion point per field

	case *A_SPEC_TYPES:
		// insertion point per field

	case *A_TARGET_1:
		// insertion point per field

	case *A_THE_HEADER:
		// insertion point per field

	case *A_TOOL_EXTENSIONS:
		// insertion point per field

	case *DATATYPE_DEFINITION_BOOLEAN:
		// insertion point per field

	case *DATATYPE_DEFINITION_DATE:
		// insertion point per field

	case *DATATYPE_DEFINITION_ENUMERATION:
		// insertion point per field

	case *DATATYPE_DEFINITION_INTEGER:
		// insertion point per field

	case *DATATYPE_DEFINITION_REAL:
		// insertion point per field

	case *DATATYPE_DEFINITION_STRING:
		// insertion point per field

	case *DATATYPE_DEFINITION_XHTML:
		// insertion point per field

	case *EMBEDDED_VALUE:
		// insertion point per field

	case *ENUM_VALUE:
		// insertion point per field

	case *RELATION_GROUP:
		// insertion point per field

	case *RELATION_GROUP_TYPE:
		// insertion point per field

	case *REQ_IF:
		// insertion point per field

	case *REQ_IF_CONTENT:
		// insertion point per field

	case *REQ_IF_HEADER:
		// insertion point per field

	case *REQ_IF_TOOL_EXTENSION:
		// insertion point per field

	case *SPECIFICATION:
		// insertion point per field

	case *SPECIFICATION_TYPE:
		// insertion point per field

	case *SPEC_HIERARCHY:
		// insertion point per field

	case *SPEC_OBJECT:
		// insertion point per field

	case *SPEC_OBJECT_TYPE:
		// insertion point per field

	case *SPEC_RELATION:
		// insertion point per field

	case *SPEC_RELATION_TYPE:
		// insertion point per field

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

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_XHTML
	// insertion point per field

	// Compute reverse map for named struct A_ALTERNATIVE_ID
	// insertion point per field

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

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_DATE
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_INTEGER
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_REAL
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_STRING
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML_1
	// insertion point per field

	// Compute reverse map for named struct A_CHILDREN
	// insertion point per field

	// Compute reverse map for named struct A_CORE_CONTENT
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPES
	// insertion point per field

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

	// Compute reverse map for named struct A_RELATION_GROUP_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_SPECIFICATION_1
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFICATIONS
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFICATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFIED_VALUES
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_ATTRIBUTES
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_OBJECTS
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_OBJECT_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATIONS
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_GROUPS
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_TYPES
	// insertion point per field

	// Compute reverse map for named struct A_TARGET_1
	// insertion point per field

	// Compute reverse map for named struct A_THE_HEADER
	// insertion point per field

	// Compute reverse map for named struct A_TOOL_EXTENSIONS
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field

	// Compute reverse map for named struct REQ_IF
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

}
