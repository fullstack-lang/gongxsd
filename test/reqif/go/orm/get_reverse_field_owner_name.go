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
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_DATE":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_INTEGER":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_REAL":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_STRING":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_XHTML":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _a_attribute_value_boolean, ok := stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _a_attribute_value_boolean.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _a_attribute_value_date, ok := stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _a_attribute_value_date.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _a_attribute_value_enumeration, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _a_attribute_value_enumeration.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _a_attribute_value_integer, ok := stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _a_attribute_value_integer.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _a_attribute_value_real, ok := stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _a_attribute_value_real.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _a_attribute_value_string, ok := stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _a_attribute_value_string.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _a_attribute_value_xhtml, ok := stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
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

	case *models.A_SOURCE_1:
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
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_BOOLEAN":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_DATE":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_ENUMERATION":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_INTEGER":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_REAL":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_STRING":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_XHTML":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFIED_VALUES":
			switch reverseField.Fieldname {
			case "ENUM_VALUE":
				if _a_specified_values, ok := stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]; ok {
					res = _a_specified_values.Name
				}
			}
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATION_GROUPS":
			switch reverseField.Fieldname {
			case "RELATION_GROUP":
				if _a_spec_relation_groups, ok := stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]; ok {
					res = _a_spec_relation_groups.Name
				}
			}
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "RELATION_GROUP_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
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
		case "A_TOOL_EXTENSIONS":
			switch reverseField.Fieldname {
			case "REQ_IF_TOOL_EXTENSION":
				if _a_tool_extensions, ok := stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]; ok {
					res = _a_tool_extensions.Name
				}
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				if _a_specifications, ok := stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]; ok {
					res = _a_specifications.Name
				}
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "A_CHILDREN":
			switch reverseField.Fieldname {
			case "SPEC_HIERARCHY":
				if _a_children, ok := stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]; ok {
					res = _a_children.Name
				}
			}
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_OBJECTS":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT":
				if _a_spec_objects, ok := stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]; ok {
					res = _a_spec_objects.Name
				}
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATIONS":
			switch reverseField.Fieldname {
			case "SPEC_RELATION":
				if _a_spec_relations, ok := stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]; ok {
					res = _a_spec_relations.Name
				}
			}
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
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
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_DATE":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_REAL":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_STRING":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_XHTML":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
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

	case *models.A_SOURCE_1:
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
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_BOOLEAN":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_DATE":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_ENUMERATION":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_INTEGER":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_REAL":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_STRING":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_XHTML":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *models.EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFIED_VALUES":
			switch reverseField.Fieldname {
			case "ENUM_VALUE":
				res = stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]
			}
		}

	case *models.RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATION_GROUPS":
			switch reverseField.Fieldname {
			case "RELATION_GROUP":
				res = stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]
			}
		}

	case *models.RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "RELATION_GROUP_TYPE":
				res = stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]
			}
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
		case "A_TOOL_EXTENSIONS":
			switch reverseField.Fieldname {
			case "REQ_IF_TOOL_EXTENSION":
				res = stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				res = stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]
			}
		}

	case *models.SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPE":
				res = stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]
			}
		}

	case *models.SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "A_CHILDREN":
			switch reverseField.Fieldname {
			case "SPEC_HIERARCHY":
				res = stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]
			}
		}

	case *models.SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_OBJECTS":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT":
				res = stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPE":
				res = stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]
			}
		}

	case *models.SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATIONS":
			switch reverseField.Fieldname {
			case "SPEC_RELATION":
				res = stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]
			}
		}

	case *models.SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_TYPE":
				res = stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]
			}
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
