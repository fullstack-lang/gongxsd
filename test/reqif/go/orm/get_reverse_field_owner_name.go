// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ENUM_VALUE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_RELATION_GROUP_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SOURCE_SPECIFICATION_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFICATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_OBJECT_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_TARGET_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_ENUM_VALUE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_RELATION_GROUP_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SOURCE_SPECIFICATION_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFICATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_OBJECT_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_RELATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_TARGET_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
