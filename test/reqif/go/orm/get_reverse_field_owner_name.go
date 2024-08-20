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
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "THE_VALUE":
				if _attribute_value_xhtml, ok := stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[inst]; ok {
					res = _attribute_value_xhtml.Name
				}
			case "THE_ORIGINAL_VALUE":
				if _attribute_value_xhtml, ok := stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[inst]; ok {
					res = _attribute_value_xhtml.Name
				}
			}
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
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "THE_VALUE":
				res = stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[inst]
			case "THE_ORIGINAL_VALUE":
				res = stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
