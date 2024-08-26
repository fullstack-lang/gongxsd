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
		case "A_ALTERNATIVE_ID":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _a_alternative_id, ok := stage.A_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _a_alternative_id.Name
				}
			}
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
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_BOOLEAN_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _renamed_attribute_value_boolean_1, ok := stage.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _renamed_attribute_value_boolean_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_DATE_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _renamed_attribute_value_date_1, ok := stage.Renamed_ATTRIBUTE_VALUE_DATE_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _renamed_attribute_value_date_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_ENUMERATION_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _renamed_attribute_value_enumeration_1, ok := stage.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _renamed_attribute_value_enumeration_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_INTEGER_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _renamed_attribute_value_integer_1, ok := stage.Renamed_ATTRIBUTE_VALUE_INTEGER_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _renamed_attribute_value_integer_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_REAL_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _renamed_attribute_value_real_1, ok := stage.Renamed_ATTRIBUTE_VALUE_REAL_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _renamed_attribute_value_real_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_STRING_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _renamed_attribute_value_string_1, ok := stage.Renamed_ATTRIBUTE_VALUE_STRING_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _renamed_attribute_value_string_1.Name
				}
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _a_values_1, ok := stage.A_VALUES_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _a_values_1.Name
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _renamed_attribute_value_xhtml_1, ok := stage.Renamed_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _renamed_attribute_value_xhtml_1.Name
				}
			}
		}

	case *models.A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_boolean, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_boolean.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_date, ok := stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_date.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_enumeration, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_enumeration.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_integer, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_integer.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_real, ok := stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_real.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_string, ok := stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_string.Name
				}
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _attribute_definition_xhtml, ok := stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _attribute_definition_xhtml.Name
				}
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_boolean, ok := stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_boolean.Name
				}
			}
		case "DATATYPE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_date, ok := stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_date.Name
				}
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_enumeration, ok := stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_enumeration.Name
				}
			}
		case "DATATYPE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_integer, ok := stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_integer.Name
				}
			}
		case "DATATYPE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_real, ok := stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_real.Name
				}
			}
		case "DATATYPE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_string, ok := stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_string.Name
				}
			}
		case "DATATYPE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _datatype_definition_xhtml, ok := stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _datatype_definition_xhtml.Name
				}
			}
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _enum_value, ok := stage.ENUM_VALUE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _enum_value.Name
				}
			}
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _relation_group, ok := stage.RELATION_GROUP_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _relation_group.Name
				}
			}
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _specification, ok := stage.SPECIFICATION_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _spec_object, ok := stage.SPEC_OBJECT_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _spec_relation, ok := stage.SPEC_RELATION_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "CHILDREN":
				if _specification, ok := stage.SPECIFICATION_CHILDREN_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "CHILDREN":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_CHILDREN_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		}

	case *models.A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "CORE_CONTENT":
				if _req_if, ok := stage.REQ_IF_CORE_CONTENT_reverseMap[inst]; ok {
					res = _req_if.Name
				}
			}
		}

	case *models.A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_DATATYPES_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "EDITABLE_ATTS":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_EDITABLE_ATTS_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		}

	case *models.A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "OBJECT":
				if _spec_hierarchy, ok := stage.SPEC_HIERARCHY_OBJECT_reverseMap[inst]; ok {
					res = _spec_hierarchy.Name
				}
			}
		}

	case *models.A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "PROPERTIES":
				if _enum_value, ok := stage.ENUM_VALUE_PROPERTIES_reverseMap[inst]; ok {
					res = _enum_value.Name
				}
			}
		}

	case *models.A_SOURCE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "SOURCE":
				if _spec_relation, ok := stage.SPEC_RELATION_SOURCE_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
		}

	case *models.A_SOURCE_SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "SOURCE_SPECIFICATION":
				if _relation_group, ok := stage.RELATION_GROUP_SOURCE_SPECIFICATION_reverseMap[inst]; ok {
					res = _relation_group.Name
				}
			}
		}

	case *models.A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "SPECIFIED_VALUES":
				if _datatype_definition_enumeration, ok := stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_reverseMap[inst]; ok {
					res = _datatype_definition_enumeration.Name
				}
			}
		}

	case *models.A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				if _relation_group_type, ok := stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]; ok {
					res = _relation_group_type.Name
				}
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				if _specification_type, ok := stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]; ok {
					res = _specification_type.Name
				}
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				if _spec_object_type, ok := stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]; ok {
					res = _spec_object_type.Name
				}
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				if _spec_relation_type, ok := stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]; ok {
					res = _spec_relation_type.Name
				}
			}
		}

	case *models.A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECTS":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_OBJECTS_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_RELATIONS_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_SPEC_RELATIONS_1:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS":
				if _relation_group, ok := stage.RELATION_GROUP_SPEC_RELATIONS_reverseMap[inst]; ok {
					res = _relation_group.Name
				}
			}
		}

	case *models.A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_GROUPS":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES":
				if _req_if_content, ok := stage.REQ_IF_CONTENT_SPEC_TYPES_reverseMap[inst]; ok {
					res = _req_if_content.Name
				}
			}
		}

	case *models.A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "THE_HEADER":
				if _req_if, ok := stage.REQ_IF_THE_HEADER_reverseMap[inst]; ok {
					res = _req_if.Name
				}
			}
		}

	case *models.A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "TOOL_EXTENSIONS":
				if _req_if, ok := stage.REQ_IF_TOOL_EXTENSIONS_reverseMap[inst]; ok {
					res = _req_if.Name
				}
			}
		}

	case *models.A_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "VALUES":
				if _attribute_value_enumeration, ok := stage.ATTRIBUTE_VALUE_ENUMERATION_VALUES_reverseMap[inst]; ok {
					res = _attribute_value_enumeration.Name
				}
			}
		}

	case *models.A_VALUES_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES":
				if _specification, ok := stage.SPECIFICATION_VALUES_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES":
				if _spec_object, ok := stage.SPEC_OBJECT_VALUES_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES":
				if _spec_relation, ok := stage.SPEC_RELATION_VALUES_reverseMap[inst]; ok {
					res = _spec_relation.Name
				}
			}
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
		case "A_PROPERTIES":
			switch reverseField.Fieldname {
			case "EMBEDDED_VALUE":
				if _a_properties, ok := stage.A_PROPERTIES_EMBEDDED_VALUE_reverseMap[inst]; ok {
					res = _a_properties.Name
				}
			}
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
		case "A_CORE_CONTENT":
			switch reverseField.Fieldname {
			case "REQ_IF_CONTENT":
				if _a_core_content, ok := stage.A_CORE_CONTENT_REQ_IF_CONTENT_reverseMap[inst]; ok {
					res = _a_core_content.Name
				}
			}
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_THE_HEADER":
			switch reverseField.Fieldname {
			case "REQ_IF_HEADER":
				if _a_the_header, ok := stage.A_THE_HEADER_REQ_IF_HEADER_reverseMap[inst]; ok {
					res = _a_the_header.Name
				}
			}
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

	case *models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_boolean, ok := stage.ATTRIBUTE_VALUE_BOOLEAN_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_boolean.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_date, ok := stage.ATTRIBUTE_VALUE_DATE_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_date.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_enumeration, ok := stage.ATTRIBUTE_VALUE_ENUMERATION_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_enumeration.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_integer, ok := stage.ATTRIBUTE_VALUE_INTEGER_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_integer.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_real, ok := stage.ATTRIBUTE_VALUE_REAL_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_real.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_string, ok := stage.ATTRIBUTE_VALUE_STRING_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_string.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "DEFINITION":
				if _attribute_value_xhtml, ok := stage.ATTRIBUTE_VALUE_XHTML_DEFINITION_reverseMap[inst]; ok {
					res = _attribute_value_xhtml.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_boolean, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_boolean.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_DATE_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_date, ok := stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_date.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_enumeration, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_enumeration.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_INTEGER_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_integer, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_integer.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_REAL_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_real, ok := stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_real.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_STRING_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_string, ok := stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_string.Name
				}
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				if _attribute_definition_xhtml, ok := stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_reverseMap[inst]; ok {
					res = _attribute_definition_xhtml.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_boolean, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_boolean.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_DATE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_date, ok := stage.ATTRIBUTE_DEFINITION_DATE_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_date.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_enumeration, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_enumeration.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_integer, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_integer.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_REAL_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_real, ok := stage.ATTRIBUTE_DEFINITION_REAL_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_real.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_STRING_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_string, ok := stage.ATTRIBUTE_DEFINITION_STRING_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_string.Name
				}
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "TYPE":
				if _attribute_definition_xhtml, ok := stage.ATTRIBUTE_DEFINITION_XHTML_TYPE_reverseMap[inst]; ok {
					res = _attribute_definition_xhtml.Name
				}
			}
		}

	case *models.Renamed_RELATION_GROUP_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "TYPE":
				if _relation_group, ok := stage.RELATION_GROUP_TYPE_reverseMap[inst]; ok {
					res = _relation_group.Name
				}
			}
		}

	case *models.Renamed_SPECIFICATION_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "TYPE":
				if _specification, ok := stage.SPECIFICATION_TYPE_reverseMap[inst]; ok {
					res = _specification.Name
				}
			}
		}

	case *models.Renamed_SPEC_OBJECT_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "TYPE":
				if _spec_object, ok := stage.SPEC_OBJECT_TYPE_reverseMap[inst]; ok {
					res = _spec_object.Name
				}
			}
		}

	case *models.Renamed_SPEC_RELATION_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "TYPE":
				if _spec_relation, ok := stage.SPEC_RELATION_TYPE_reverseMap[inst]; ok {
					res = _spec_relation.Name
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
		case "A_ALTERNATIVE_ID":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.A_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[inst]
			}
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
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_BOOLEAN_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_DATE_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.Renamed_ATTRIBUTE_VALUE_DATE_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_ENUMERATION_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_INTEGER_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.Renamed_ATTRIBUTE_VALUE_INTEGER_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_REAL_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.Renamed_ATTRIBUTE_VALUE_REAL_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_STRING_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.Renamed_ATTRIBUTE_VALUE_STRING_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_VALUES_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.A_VALUES_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "Renamed_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.Renamed_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		}

	case *models.A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "DATATYPE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.ENUM_VALUE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.RELATION_GROUP_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPECIFICATION_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPEC_HIERARCHY_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPEC_OBJECT_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPEC_RELATION_ALTERNATIVE_ID_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "ALTERNATIVE_ID":
				res = stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_reverseMap[inst]
			}
		}

	case *models.A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "CHILDREN":
				res = stage.SPECIFICATION_CHILDREN_reverseMap[inst]
			}
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "CHILDREN":
				res = stage.SPEC_HIERARCHY_CHILDREN_reverseMap[inst]
			}
		}

	case *models.A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "CORE_CONTENT":
				res = stage.REQ_IF_CORE_CONTENT_reverseMap[inst]
			}
		}

	case *models.A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "DATATYPES":
				res = stage.REQ_IF_CONTENT_DATATYPES_reverseMap[inst]
			}
		}

	case *models.A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "EDITABLE_ATTS":
				res = stage.SPEC_HIERARCHY_EDITABLE_ATTS_reverseMap[inst]
			}
		}

	case *models.A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_HIERARCHY":
			switch reverseField.Fieldname {
			case "OBJECT":
				res = stage.SPEC_HIERARCHY_OBJECT_reverseMap[inst]
			}
		}

	case *models.A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		case "ENUM_VALUE":
			switch reverseField.Fieldname {
			case "PROPERTIES":
				res = stage.ENUM_VALUE_PROPERTIES_reverseMap[inst]
			}
		}

	case *models.A_SOURCE:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "SOURCE":
				res = stage.SPEC_RELATION_SOURCE_reverseMap[inst]
			}
		}

	case *models.A_SOURCE_SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "SOURCE_SPECIFICATION":
				res = stage.RELATION_GROUP_SOURCE_SPECIFICATION_reverseMap[inst]
			}
		}

	case *models.A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPECIFICATIONS":
				res = stage.REQ_IF_CONTENT_SPECIFICATIONS_reverseMap[inst]
			}
		}

	case *models.A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		case "DATATYPE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "SPECIFIED_VALUES":
				res = stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_reverseMap[inst]
			}
		}

	case *models.A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				res = stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]
			}
		case "SPECIFICATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				res = stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]
			}
		case "SPEC_OBJECT_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				res = stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]
			}
		case "SPEC_RELATION_TYPE":
			switch reverseField.Fieldname {
			case "SPEC_ATTRIBUTES":
				res = stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_reverseMap[inst]
			}
		}

	case *models.A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_OBJECTS":
				res = stage.REQ_IF_CONTENT_SPEC_OBJECTS_reverseMap[inst]
			}
		}

	case *models.A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS":
				res = stage.REQ_IF_CONTENT_SPEC_RELATIONS_reverseMap[inst]
			}
		}

	case *models.A_SPEC_RELATIONS_1:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "SPEC_RELATIONS":
				res = stage.RELATION_GROUP_SPEC_RELATIONS_reverseMap[inst]
			}
		}

	case *models.A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_GROUPS":
				res = stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_reverseMap[inst]
			}
		}

	case *models.A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF_CONTENT":
			switch reverseField.Fieldname {
			case "SPEC_TYPES":
				res = stage.REQ_IF_CONTENT_SPEC_TYPES_reverseMap[inst]
			}
		}

	case *models.A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "THE_HEADER":
				res = stage.REQ_IF_THE_HEADER_reverseMap[inst]
			}
		}

	case *models.A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		case "REQ_IF":
			switch reverseField.Fieldname {
			case "TOOL_EXTENSIONS":
				res = stage.REQ_IF_TOOL_EXTENSIONS_reverseMap[inst]
			}
		}

	case *models.A_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "VALUES":
				res = stage.ATTRIBUTE_VALUE_ENUMERATION_VALUES_reverseMap[inst]
			}
		}

	case *models.A_VALUES_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "VALUES":
				res = stage.SPECIFICATION_VALUES_reverseMap[inst]
			}
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "VALUES":
				res = stage.SPEC_OBJECT_VALUES_reverseMap[inst]
			}
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "VALUES":
				res = stage.SPEC_RELATION_VALUES_reverseMap[inst]
			}
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
		case "A_PROPERTIES":
			switch reverseField.Fieldname {
			case "EMBEDDED_VALUE":
				res = stage.A_PROPERTIES_EMBEDDED_VALUE_reverseMap[inst]
			}
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
		case "A_CORE_CONTENT":
			switch reverseField.Fieldname {
			case "REQ_IF_CONTENT":
				res = stage.A_CORE_CONTENT_REQ_IF_CONTENT_reverseMap[inst]
			}
		}

	case *models.REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_THE_HEADER":
			switch reverseField.Fieldname {
			case "REQ_IF_HEADER":
				res = stage.A_THE_HEADER_REQ_IF_HEADER_reverseMap[inst]
			}
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

	case *models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_BOOLEAN_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_DATE_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_ENUMERATION_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_INTEGER_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_REAL_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_STRING_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "DEFINITION":
				res = stage.ATTRIBUTE_VALUE_XHTML_DEFINITION_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_DATE_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_INTEGER_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_REAL_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_STRING_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "DEFAULT_VALUE":
				res = stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_BOOLEAN_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_DATE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_DATE":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_DATE_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_ENUMERATION_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_INTEGER":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_INTEGER_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_REAL_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_REAL":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_REAL_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_STRING_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_STRING":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_STRING_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "ATTRIBUTE_DEFINITION_XHTML":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.ATTRIBUTE_DEFINITION_XHTML_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_RELATION_GROUP_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "RELATION_GROUP":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.RELATION_GROUP_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_SPECIFICATION_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPECIFICATION":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.SPECIFICATION_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_SPEC_OBJECT_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_OBJECT":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.SPEC_OBJECT_TYPE_reverseMap[inst]
			}
		}

	case *models.Renamed_SPEC_RELATION_TYPE_REF_1:
		switch reverseField.GongstructName {
		// insertion point
		case "SPEC_RELATION":
			switch reverseField.Fieldname {
			case "TYPE":
				res = stage.SPEC_RELATION_TYPE_reverseMap[inst]
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
