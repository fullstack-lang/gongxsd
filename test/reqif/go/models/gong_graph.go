// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ALTERNATIVE_ID:
		ok = stage.IsStagedALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		ok = stage.IsStagedATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		ok = stage.IsStagedATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		ok = stage.IsStagedATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		ok = stage.IsStagedATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		ok = stage.IsStagedATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		ok = stage.IsStagedATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		ok = stage.IsStagedATTRIBUTE_VALUE_XHTML(target)

	case *A_ALTERNATIVE_ID:
		ok = stage.IsStagedA_ALTERNATIVE_ID(target)

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(target)

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_DATE_REF(target)

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(target)

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_INTEGER_REF(target)

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_REAL_REF(target)

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_STRING_REF(target)

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		ok = stage.IsStagedA_ATTRIBUTE_DEFINITION_XHTML_REF(target)

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_BOOLEAN(target)

	case *A_ATTRIBUTE_VALUE_DATE:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_DATE(target)

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_ENUMERATION(target)

	case *A_ATTRIBUTE_VALUE_INTEGER:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_INTEGER(target)

	case *A_ATTRIBUTE_VALUE_REAL:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_REAL(target)

	case *A_ATTRIBUTE_VALUE_STRING:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_STRING(target)

	case *A_ATTRIBUTE_VALUE_XHTML:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_XHTML(target)

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		ok = stage.IsStagedA_ATTRIBUTE_VALUE_XHTML_1(target)

	case *A_CHILDREN:
		ok = stage.IsStagedA_CHILDREN(target)

	case *A_CORE_CONTENT:
		ok = stage.IsStagedA_CORE_CONTENT(target)

	case *A_DATATYPES:
		ok = stage.IsStagedA_DATATYPES(target)

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_BOOLEAN_REF(target)

	case *A_DATATYPE_DEFINITION_DATE_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_DATE_REF(target)

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_ENUMERATION_REF(target)

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_INTEGER_REF(target)

	case *A_DATATYPE_DEFINITION_REAL_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_REAL_REF(target)

	case *A_DATATYPE_DEFINITION_STRING_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_STRING_REF(target)

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		ok = stage.IsStagedA_DATATYPE_DEFINITION_XHTML_REF(target)

	case *A_EDITABLE_ATTS:
		ok = stage.IsStagedA_EDITABLE_ATTS(target)

	case *A_ENUM_VALUE_REF:
		ok = stage.IsStagedA_ENUM_VALUE_REF(target)

	case *A_OBJECT:
		ok = stage.IsStagedA_OBJECT(target)

	case *A_PROPERTIES:
		ok = stage.IsStagedA_PROPERTIES(target)

	case *A_RELATION_GROUP_TYPE_REF:
		ok = stage.IsStagedA_RELATION_GROUP_TYPE_REF(target)

	case *A_SOURCE_SPECIFICATION_1:
		ok = stage.IsStagedA_SOURCE_SPECIFICATION_1(target)

	case *A_SPECIFICATIONS:
		ok = stage.IsStagedA_SPECIFICATIONS(target)

	case *A_SPECIFICATION_TYPE_REF:
		ok = stage.IsStagedA_SPECIFICATION_TYPE_REF(target)

	case *A_SPECIFIED_VALUES:
		ok = stage.IsStagedA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		ok = stage.IsStagedA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		ok = stage.IsStagedA_SPEC_OBJECTS(target)

	case *A_SPEC_OBJECT_TYPE_REF:
		ok = stage.IsStagedA_SPEC_OBJECT_TYPE_REF(target)

	case *A_SPEC_RELATIONS:
		ok = stage.IsStagedA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATION_GROUPS:
		ok = stage.IsStagedA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_RELATION_REF:
		ok = stage.IsStagedA_SPEC_RELATION_REF(target)

	case *A_SPEC_RELATION_TYPE_REF:
		ok = stage.IsStagedA_SPEC_RELATION_TYPE_REF(target)

	case *A_SPEC_TYPES:
		ok = stage.IsStagedA_SPEC_TYPES(target)

	case *A_TARGET_1:
		ok = stage.IsStagedA_TARGET_1(target)

	case *A_THE_HEADER:
		ok = stage.IsStagedA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		ok = stage.IsStagedA_TOOL_EXTENSIONS(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		ok = stage.IsStagedDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		ok = stage.IsStagedDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		ok = stage.IsStagedDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		ok = stage.IsStagedDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		ok = stage.IsStagedDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		ok = stage.IsStagedDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		ok = stage.IsStagedDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		ok = stage.IsStagedEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		ok = stage.IsStagedENUM_VALUE(target)

	case *RELATION_GROUP:
		ok = stage.IsStagedRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		ok = stage.IsStagedRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		ok = stage.IsStagedREQ_IF(target)

	case *REQ_IF_CONTENT:
		ok = stage.IsStagedREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		ok = stage.IsStagedREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		ok = stage.IsStagedREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		ok = stage.IsStagedSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		ok = stage.IsStagedSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		ok = stage.IsStagedSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		ok = stage.IsStagedSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		ok = stage.IsStagedSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		ok = stage.IsStagedSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		ok = stage.IsStagedSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		ok = stage.IsStagedXHTML_CONTENT(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) (ok bool) {

	_, ok = stage.ALTERNATIVE_IDs[alternative_id]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]

	return
}

func (stage *StageStruct) IsStagedA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) (ok bool) {

	_, ok = stage.A_ALTERNATIVE_IDs[a_alternative_id]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[a_attribute_definition_boolean_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[a_attribute_definition_date_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[a_attribute_definition_enumeration_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[a_attribute_definition_integer_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[a_attribute_definition_real_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[a_attribute_definition_string_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[a_attribute_definition_xhtml_ref]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_BOOLEANs[a_attribute_value_boolean]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_DATEs[a_attribute_value_date]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_INTEGERs[a_attribute_value_integer]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_REALs[a_attribute_value_real]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_STRINGs[a_attribute_value_string]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_XHTMLs[a_attribute_value_xhtml]

	return
}

func (stage *StageStruct) IsStagedA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_XHTML_1s[a_attribute_value_xhtml_1]

	return
}

func (stage *StageStruct) IsStagedA_CHILDREN(a_children *A_CHILDREN) (ok bool) {

	_, ok = stage.A_CHILDRENs[a_children]

	return
}

func (stage *StageStruct) IsStagedA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) (ok bool) {

	_, ok = stage.A_CORE_CONTENTs[a_core_content]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPES(a_datatypes *A_DATATYPES) (ok bool) {

	_, ok = stage.A_DATATYPESs[a_datatypes]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[a_datatype_definition_boolean_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_DATE_REFs[a_datatype_definition_date_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[a_datatype_definition_enumeration_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_INTEGER_REFs[a_datatype_definition_integer_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_REAL_REFs[a_datatype_definition_real_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_STRING_REFs[a_datatype_definition_string_ref]

	return
}

func (stage *StageStruct) IsStagedA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_XHTML_REFs[a_datatype_definition_xhtml_ref]

	return
}

func (stage *StageStruct) IsStagedA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) (ok bool) {

	_, ok = stage.A_EDITABLE_ATTSs[a_editable_atts]

	return
}

func (stage *StageStruct) IsStagedA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) (ok bool) {

	_, ok = stage.A_ENUM_VALUE_REFs[a_enum_value_ref]

	return
}

func (stage *StageStruct) IsStagedA_OBJECT(a_object *A_OBJECT) (ok bool) {

	_, ok = stage.A_OBJECTs[a_object]

	return
}

func (stage *StageStruct) IsStagedA_PROPERTIES(a_properties *A_PROPERTIES) (ok bool) {

	_, ok = stage.A_PROPERTIESs[a_properties]

	return
}

func (stage *StageStruct) IsStagedA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) (ok bool) {

	_, ok = stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref]

	return
}

func (stage *StageStruct) IsStagedA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) (ok bool) {

	_, ok = stage.A_SOURCE_SPECIFICATION_1s[a_source_specification_1]

	return
}

func (stage *StageStruct) IsStagedA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) (ok bool) {

	_, ok = stage.A_SPECIFICATIONSs[a_specifications]

	return
}

func (stage *StageStruct) IsStagedA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) (ok bool) {

	_, ok = stage.A_SPECIFICATION_TYPE_REFs[a_specification_type_ref]

	return
}

func (stage *StageStruct) IsStagedA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) (ok bool) {

	_, ok = stage.A_SPECIFIED_VALUESs[a_specified_values]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) (ok bool) {

	_, ok = stage.A_SPEC_ATTRIBUTESs[a_spec_attributes]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) (ok bool) {

	_, ok = stage.A_SPEC_OBJECTSs[a_spec_objects]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) (ok bool) {

	_, ok = stage.A_SPEC_OBJECT_TYPE_REFs[a_spec_object_type_ref]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) (ok bool) {

	_, ok = stage.A_SPEC_RELATIONSs[a_spec_relations]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_REFs[a_spec_relation_ref]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_TYPE_REFs[a_spec_relation_type_ref]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) (ok bool) {

	_, ok = stage.A_SPEC_TYPESs[a_spec_types]

	return
}

func (stage *StageStruct) IsStagedA_TARGET_1(a_target_1 *A_TARGET_1) (ok bool) {

	_, ok = stage.A_TARGET_1s[a_target_1]

	return
}

func (stage *StageStruct) IsStagedA_THE_HEADER(a_the_header *A_THE_HEADER) (ok bool) {

	_, ok = stage.A_THE_HEADERs[a_the_header]

	return
}

func (stage *StageStruct) IsStagedA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) (ok bool) {

	_, ok = stage.A_TOOL_EXTENSIONSs[a_tool_extensions]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]

	return
}

func (stage *StageStruct) IsStagedEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) (ok bool) {

	_, ok = stage.EMBEDDED_VALUEs[embedded_value]

	return
}

func (stage *StageStruct) IsStagedENUM_VALUE(enum_value *ENUM_VALUE) (ok bool) {

	_, ok = stage.ENUM_VALUEs[enum_value]

	return
}

func (stage *StageStruct) IsStagedRELATION_GROUP(relation_group *RELATION_GROUP) (ok bool) {

	_, ok = stage.RELATION_GROUPs[relation_group]

	return
}

func (stage *StageStruct) IsStagedRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) (ok bool) {

	_, ok = stage.RELATION_GROUP_TYPEs[relation_group_type]

	return
}

func (stage *StageStruct) IsStagedREQ_IF(req_if *REQ_IF) (ok bool) {

	_, ok = stage.REQ_IFs[req_if]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) (ok bool) {

	_, ok = stage.REQ_IF_CONTENTs[req_if_content]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) (ok bool) {

	_, ok = stage.REQ_IF_HEADERs[req_if_header]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) (ok bool) {

	_, ok = stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATION(specification *SPECIFICATION) (ok bool) {

	_, ok = stage.SPECIFICATIONs[specification]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) (ok bool) {

	_, ok = stage.SPECIFICATION_TYPEs[specification_type]

	return
}

func (stage *StageStruct) IsStagedSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) (ok bool) {

	_, ok = stage.SPEC_HIERARCHYs[spec_hierarchy]

	return
}

func (stage *StageStruct) IsStagedSPEC_OBJECT(spec_object *SPEC_OBJECT) (ok bool) {

	_, ok = stage.SPEC_OBJECTs[spec_object]

	return
}

func (stage *StageStruct) IsStagedSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) (ok bool) {

	_, ok = stage.SPEC_OBJECT_TYPEs[spec_object_type]

	return
}

func (stage *StageStruct) IsStagedSPEC_RELATION(spec_relation *SPEC_RELATION) (ok bool) {

	_, ok = stage.SPEC_RELATIONs[spec_relation]

	return
}

func (stage *StageStruct) IsStagedSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) (ok bool) {

	_, ok = stage.SPEC_RELATION_TYPEs[spec_relation_type]

	return
}

func (stage *StageStruct) IsStagedXHTML_CONTENT(xhtml_content *XHTML_CONTENT) (ok bool) {

	_, ok = stage.XHTML_CONTENTs[xhtml_content]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *ALTERNATIVE_ID:
		stage.StageBranchALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.StageBranchATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		stage.StageBranchATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.StageBranchATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.StageBranchATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		stage.StageBranchATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		stage.StageBranchATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.StageBranchATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.StageBranchATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		stage.StageBranchATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.StageBranchATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		stage.StageBranchATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		stage.StageBranchATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		stage.StageBranchATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		stage.StageBranchATTRIBUTE_VALUE_XHTML(target)

	case *A_ALTERNATIVE_ID:
		stage.StageBranchA_ALTERNATIVE_ID(target)

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(target)

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(target)

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(target)

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(target)

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(target)

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(target)

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.StageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(target)

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.StageBranchA_ATTRIBUTE_VALUE_BOOLEAN(target)

	case *A_ATTRIBUTE_VALUE_DATE:
		stage.StageBranchA_ATTRIBUTE_VALUE_DATE(target)

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.StageBranchA_ATTRIBUTE_VALUE_ENUMERATION(target)

	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.StageBranchA_ATTRIBUTE_VALUE_INTEGER(target)

	case *A_ATTRIBUTE_VALUE_REAL:
		stage.StageBranchA_ATTRIBUTE_VALUE_REAL(target)

	case *A_ATTRIBUTE_VALUE_STRING:
		stage.StageBranchA_ATTRIBUTE_VALUE_STRING(target)

	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.StageBranchA_ATTRIBUTE_VALUE_XHTML(target)

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.StageBranchA_ATTRIBUTE_VALUE_XHTML_1(target)

	case *A_CHILDREN:
		stage.StageBranchA_CHILDREN(target)

	case *A_CORE_CONTENT:
		stage.StageBranchA_CORE_CONTENT(target)

	case *A_DATATYPES:
		stage.StageBranchA_DATATYPES(target)

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(target)

	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_DATE_REF(target)

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(target)

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_INTEGER_REF(target)

	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_REAL_REF(target)

	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_STRING_REF(target)

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.StageBranchA_DATATYPE_DEFINITION_XHTML_REF(target)

	case *A_EDITABLE_ATTS:
		stage.StageBranchA_EDITABLE_ATTS(target)

	case *A_ENUM_VALUE_REF:
		stage.StageBranchA_ENUM_VALUE_REF(target)

	case *A_OBJECT:
		stage.StageBranchA_OBJECT(target)

	case *A_PROPERTIES:
		stage.StageBranchA_PROPERTIES(target)

	case *A_RELATION_GROUP_TYPE_REF:
		stage.StageBranchA_RELATION_GROUP_TYPE_REF(target)

	case *A_SOURCE_SPECIFICATION_1:
		stage.StageBranchA_SOURCE_SPECIFICATION_1(target)

	case *A_SPECIFICATIONS:
		stage.StageBranchA_SPECIFICATIONS(target)

	case *A_SPECIFICATION_TYPE_REF:
		stage.StageBranchA_SPECIFICATION_TYPE_REF(target)

	case *A_SPECIFIED_VALUES:
		stage.StageBranchA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		stage.StageBranchA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		stage.StageBranchA_SPEC_OBJECTS(target)

	case *A_SPEC_OBJECT_TYPE_REF:
		stage.StageBranchA_SPEC_OBJECT_TYPE_REF(target)

	case *A_SPEC_RELATIONS:
		stage.StageBranchA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATION_GROUPS:
		stage.StageBranchA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_RELATION_REF:
		stage.StageBranchA_SPEC_RELATION_REF(target)

	case *A_SPEC_RELATION_TYPE_REF:
		stage.StageBranchA_SPEC_RELATION_TYPE_REF(target)

	case *A_SPEC_TYPES:
		stage.StageBranchA_SPEC_TYPES(target)

	case *A_TARGET_1:
		stage.StageBranchA_TARGET_1(target)

	case *A_THE_HEADER:
		stage.StageBranchA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		stage.StageBranchA_TOOL_EXTENSIONS(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.StageBranchDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		stage.StageBranchDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.StageBranchDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		stage.StageBranchDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		stage.StageBranchDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		stage.StageBranchDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		stage.StageBranchDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		stage.StageBranchEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		stage.StageBranchENUM_VALUE(target)

	case *RELATION_GROUP:
		stage.StageBranchRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		stage.StageBranchRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		stage.StageBranchREQ_IF(target)

	case *REQ_IF_CONTENT:
		stage.StageBranchREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		stage.StageBranchREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		stage.StageBranchREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		stage.StageBranchSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		stage.StageBranchSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		stage.StageBranchSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		stage.StageBranchSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		stage.StageBranchSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		stage.StageBranchSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		stage.StageBranchSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		stage.StageBranchXHTML_CONTENT(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if IsStaged(stage, alternative_id) {
		return
	}

	alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_boolean.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_boolean.ALTERNATIVE_ID)
	}
	if attribute_definition_boolean.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_boolean.DEFAULT_VALUE)
	}
	if attribute_definition_boolean.TYPE != nil {
		StageBranch(stage, attribute_definition_boolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_date.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_date.ALTERNATIVE_ID)
	}
	if attribute_definition_date.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_date.DEFAULT_VALUE)
	}
	if attribute_definition_date.TYPE != nil {
		StageBranch(stage, attribute_definition_date.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_enumeration.ALTERNATIVE_ID)
	}
	if attribute_definition_enumeration.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_enumeration.DEFAULT_VALUE)
	}
	if attribute_definition_enumeration.TYPE != nil {
		StageBranch(stage, attribute_definition_enumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integer.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_integer.ALTERNATIVE_ID)
	}
	if attribute_definition_integer.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_integer.DEFAULT_VALUE)
	}
	if attribute_definition_integer.TYPE != nil {
		StageBranch(stage, attribute_definition_integer.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_real.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_real.ALTERNATIVE_ID)
	}
	if attribute_definition_real.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_real.DEFAULT_VALUE)
	}
	if attribute_definition_real.TYPE != nil {
		StageBranch(stage, attribute_definition_real.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_string.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_string.ALTERNATIVE_ID)
	}
	if attribute_definition_string.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_string.DEFAULT_VALUE)
	}
	if attribute_definition_string.TYPE != nil {
		StageBranch(stage, attribute_definition_string.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
		StageBranch(stage, attribute_definition_xhtml.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtml.DEFAULT_VALUE != nil {
		StageBranch(stage, attribute_definition_xhtml.DEFAULT_VALUE)
	}
	if attribute_definition_xhtml.TYPE != nil {
		StageBranch(stage, attribute_definition_xhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_boolean.DEFINITION != nil {
		StageBranch(stage, attribute_value_boolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_date.DEFINITION != nil {
		StageBranch(stage, attribute_value_date.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumeration.DEFINITION != nil {
		StageBranch(stage, attribute_value_enumeration.DEFINITION)
	}
	if attribute_value_enumeration.VALUES != nil {
		StageBranch(stage, attribute_value_enumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integer.DEFINITION != nil {
		StageBranch(stage, attribute_value_integer.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_real.DEFINITION != nil {
		StageBranch(stage, attribute_value_real.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_string.DEFINITION != nil {
		StageBranch(stage, attribute_value_string.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtml.THE_VALUE != nil {
		StageBranch(stage, attribute_value_xhtml.THE_VALUE)
	}
	if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
		StageBranch(stage, attribute_value_xhtml.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtml.DEFINITION != nil {
		StageBranch(stage, attribute_value_xhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if IsStaged(stage, a_alternative_id) {
		return
	}

	a_alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_id.ALTERNATIVE_ID != nil {
		StageBranch(stage, a_alternative_id.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_boolean_ref) {
		return
	}

	a_attribute_definition_boolean_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_date_ref) {
		return
	}

	a_attribute_definition_date_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_enumeration_ref) {
		return
	}

	a_attribute_definition_enumeration_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_integer_ref) {
		return
	}

	a_attribute_definition_integer_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_real_ref) {
		return
	}

	a_attribute_definition_real_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_string_ref) {
		return
	}

	a_attribute_definition_string_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_definition_xhtml_ref) {
		return
	}

	a_attribute_definition_xhtml_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_boolean) {
		return
	}

	a_attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN != nil {
		StageBranch(stage, a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_date) {
		return
	}

	a_attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_date.ATTRIBUTE_VALUE_DATE != nil {
		StageBranch(stage, a_attribute_value_date.ATTRIBUTE_VALUE_DATE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_enumeration) {
		return
	}

	a_attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION != nil {
		StageBranch(stage, a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_integer) {
		return
	}

	a_attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER != nil {
		StageBranch(stage, a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_real) {
		return
	}

	a_attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_real.ATTRIBUTE_VALUE_REAL != nil {
		StageBranch(stage, a_attribute_value_real.ATTRIBUTE_VALUE_REAL)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_string) {
		return
	}

	a_attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_string.ATTRIBUTE_VALUE_STRING != nil {
		StageBranch(stage, a_attribute_value_string.ATTRIBUTE_VALUE_STRING)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_xhtml) {
		return
	}

	a_attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML != nil {
		StageBranch(stage, a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) {

	// check if instance is already staged
	if IsStaged(stage, a_attribute_value_xhtml_1) {
		return
	}

	a_attribute_value_xhtml_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML != nil {
		StageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if IsStaged(stage, a_children) {
		return
	}

	a_children.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_children.SPEC_HIERARCHY != nil {
		StageBranch(stage, a_children.SPEC_HIERARCHY)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, a_core_content) {
		return
	}

	a_core_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_core_content.REQ_IF_CONTENT != nil {
		StageBranch(stage, a_core_content.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if IsStaged(stage, a_datatypes) {
		return
	}

	a_datatypes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_datatypes.DATATYPE_DEFINITION_BOOLEAN != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_BOOLEAN)
	}
	if a_datatypes.DATATYPE_DEFINITION_DATE != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_DATE)
	}
	if a_datatypes.DATATYPE_DEFINITION_ENUMERATION != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_ENUMERATION)
	}
	if a_datatypes.DATATYPE_DEFINITION_INTEGER != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_INTEGER)
	}
	if a_datatypes.DATATYPE_DEFINITION_REAL != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_REAL)
	}
	if a_datatypes.DATATYPE_DEFINITION_STRING != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_STRING)
	}
	if a_datatypes.DATATYPE_DEFINITION_XHTML != nil {
		StageBranch(stage, a_datatypes.DATATYPE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_boolean_ref) {
		return
	}

	a_datatype_definition_boolean_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_date_ref) {
		return
	}

	a_datatype_definition_date_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_enumeration_ref) {
		return
	}

	a_datatype_definition_enumeration_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_integer_ref) {
		return
	}

	a_datatype_definition_integer_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_real_ref) {
		return
	}

	a_datatype_definition_real_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_string_ref) {
		return
	}

	a_datatype_definition_string_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_datatype_definition_xhtml_ref) {
		return
	}

	a_datatype_definition_xhtml_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) {

	// check if instance is already staged
	if IsStaged(stage, a_editable_atts) {
		return
	}

	a_editable_atts.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_enum_value_ref) {
		return
	}

	a_enum_value_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_OBJECT(a_object *A_OBJECT) {

	// check if instance is already staged
	if IsStaged(stage, a_object) {
		return
	}

	a_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_PROPERTIES(a_properties *A_PROPERTIES) {

	// check if instance is already staged
	if IsStaged(stage, a_properties) {
		return
	}

	a_properties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_properties.EMBEDDED_VALUE != nil {
		StageBranch(stage, a_properties.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_relation_group_type_ref) {
		return
	}

	a_relation_group_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) {

	// check if instance is already staged
	if IsStaged(stage, a_source_specification_1) {
		return
	}

	a_source_specification_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) {

	// check if instance is already staged
	if IsStaged(stage, a_specifications) {
		return
	}

	a_specifications.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_specifications.SPECIFICATION != nil {
		StageBranch(stage, a_specifications.SPECIFICATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_specification_type_ref) {
		return
	}

	a_specification_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if IsStaged(stage, a_specified_values) {
		return
	}

	a_specified_values.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_specified_values.ENUM_VALUE != nil {
		StageBranch(stage, a_specified_values.ENUM_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_attributes) {
		return
	}

	a_spec_attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_DATE != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_DATE)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_REAL != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_REAL)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_STRING != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_STRING)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML != nil {
		StageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_objects) {
		return
	}

	a_spec_objects.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_objects.SPEC_OBJECT != nil {
		StageBranch(stage, a_spec_objects.SPEC_OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_object_type_ref) {
		return
	}

	a_spec_object_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relations) {
		return
	}

	a_spec_relations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relations.SPEC_RELATION != nil {
		StageBranch(stage, a_spec_relations.SPEC_RELATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relation_groups.RELATION_GROUP != nil {
		StageBranch(stage, a_spec_relation_groups.RELATION_GROUP)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relation_ref) {
		return
	}

	a_spec_relation_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relation_type_ref) {
		return
	}

	a_spec_relation_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_types) {
		return
	}

	a_spec_types.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_types.RELATION_GROUP_TYPE != nil {
		StageBranch(stage, a_spec_types.RELATION_GROUP_TYPE)
	}
	if a_spec_types.SPEC_OBJECT_TYPE != nil {
		StageBranch(stage, a_spec_types.SPEC_OBJECT_TYPE)
	}
	if a_spec_types.SPEC_RELATION_TYPE != nil {
		StageBranch(stage, a_spec_types.SPEC_RELATION_TYPE)
	}
	if a_spec_types.SPECIFICATION_TYPE != nil {
		StageBranch(stage, a_spec_types.SPECIFICATION_TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TARGET_1(a_target_1 *A_TARGET_1) {

	// check if instance is already staged
	if IsStaged(stage, a_target_1) {
		return
	}

	a_target_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if IsStaged(stage, a_the_header) {
		return
	}

	a_the_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_the_header.REQ_IF_HEADER != nil {
		StageBranch(stage, a_the_header.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if IsStaged(stage, a_tool_extensions) {
		return
	}

	a_tool_extensions.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_tool_extensions.REQ_IF_TOOL_EXTENSION != nil {
		StageBranch(stage, a_tool_extensions.REQ_IF_TOOL_EXTENSION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_boolean.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_boolean.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_date.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_date.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_enumeration.ALTERNATIVE_ID)
	}
	if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
		StageBranch(stage, datatype_definition_enumeration.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integer.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_integer.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_real.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_real.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_string.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_string.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
		StageBranch(stage, datatype_definition_xhtml.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if IsStaged(stage, embedded_value) {
		return
	}

	embedded_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if IsStaged(stage, enum_value) {
		return
	}

	enum_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enum_value.ALTERNATIVE_ID != nil {
		StageBranch(stage, enum_value.ALTERNATIVE_ID)
	}
	if enum_value.PROPERTIES != nil {
		StageBranch(stage, enum_value.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if IsStaged(stage, relation_group) {
		return
	}

	relation_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group.ALTERNATIVE_ID != nil {
		StageBranch(stage, relation_group.ALTERNATIVE_ID)
	}
	if relation_group.SOURCE_SPECIFICATION != nil {
		StageBranch(stage, relation_group.SOURCE_SPECIFICATION)
	}
	if relation_group.SPEC_RELATIONS != nil {
		StageBranch(stage, relation_group.SPEC_RELATIONS)
	}
	if relation_group.TARGET_SPECIFICATION != nil {
		StageBranch(stage, relation_group.TARGET_SPECIFICATION)
	}
	if relation_group.TYPE != nil {
		StageBranch(stage, relation_group.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_type.ALTERNATIVE_ID != nil {
		StageBranch(stage, relation_group_type.ALTERNATIVE_ID)
	}
	if relation_group_type.SPEC_ATTRIBUTES != nil {
		StageBranch(stage, relation_group_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if IsStaged(stage, req_if) {
		return
	}

	req_if.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER != nil {
		StageBranch(stage, req_if.THE_HEADER)
	}
	if req_if.CORE_CONTENT != nil {
		StageBranch(stage, req_if.CORE_CONTENT)
	}
	if req_if.TOOL_EXTENSIONS != nil {
		StageBranch(stage, req_if.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if_content.DATATYPES != nil {
		StageBranch(stage, req_if_content.DATATYPES)
	}
	if req_if_content.SPEC_TYPES != nil {
		StageBranch(stage, req_if_content.SPEC_TYPES)
	}
	if req_if_content.SPEC_OBJECTS != nil {
		StageBranch(stage, req_if_content.SPEC_OBJECTS)
	}
	if req_if_content.SPEC_RELATIONS != nil {
		StageBranch(stage, req_if_content.SPEC_RELATIONS)
	}
	if req_if_content.SPECIFICATIONS != nil {
		StageBranch(stage, req_if_content.SPECIFICATIONS)
	}
	if req_if_content.SPEC_RELATION_GROUPS != nil {
		StageBranch(stage, req_if_content.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if IsStaged(stage, req_if_header) {
		return
	}

	req_if_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if IsStaged(stage, req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, specification) {
		return
	}

	specification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVE_ID != nil {
		StageBranch(stage, specification.ALTERNATIVE_ID)
	}
	if specification.CHILDREN != nil {
		StageBranch(stage, specification.CHILDREN)
	}
	if specification.VALUES != nil {
		StageBranch(stage, specification.VALUES)
	}
	if specification.TYPE != nil {
		StageBranch(stage, specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, specification_type) {
		return
	}

	specification_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification_type.ALTERNATIVE_ID != nil {
		StageBranch(stage, specification_type.ALTERNATIVE_ID)
	}
	if specification_type.SPEC_ATTRIBUTES != nil {
		StageBranch(stage, specification_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchy.ALTERNATIVE_ID != nil {
		StageBranch(stage, spec_hierarchy.ALTERNATIVE_ID)
	}
	if spec_hierarchy.CHILDREN != nil {
		StageBranch(stage, spec_hierarchy.CHILDREN)
	}
	if spec_hierarchy.EDITABLE_ATTS != nil {
		StageBranch(stage, spec_hierarchy.EDITABLE_ATTS)
	}
	if spec_hierarchy.OBJECT != nil {
		StageBranch(stage, spec_hierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if IsStaged(stage, spec_object) {
		return
	}

	spec_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object.ALTERNATIVE_ID != nil {
		StageBranch(stage, spec_object.ALTERNATIVE_ID)
	}
	if spec_object.VALUES != nil {
		StageBranch(stage, spec_object.VALUES)
	}
	if spec_object.TYPE != nil {
		StageBranch(stage, spec_object.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_type.ALTERNATIVE_ID != nil {
		StageBranch(stage, spec_object_type.ALTERNATIVE_ID)
	}
	if spec_object_type.SPEC_ATTRIBUTES != nil {
		StageBranch(stage, spec_object_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation.ALTERNATIVE_ID != nil {
		StageBranch(stage, spec_relation.ALTERNATIVE_ID)
	}
	if spec_relation.VALUES != nil {
		StageBranch(stage, spec_relation.VALUES)
	}
	if spec_relation.SOURCE != nil {
		StageBranch(stage, spec_relation.SOURCE)
	}
	if spec_relation.TARGET != nil {
		StageBranch(stage, spec_relation.TARGET)
	}
	if spec_relation.TYPE != nil {
		StageBranch(stage, spec_relation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_type.ALTERNATIVE_ID != nil {
		StageBranch(stage, spec_relation_type.ALTERNATIVE_ID)
	}
	if spec_relation_type.SPEC_ATTRIBUTES != nil {
		StageBranch(stage, spec_relation_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_content) {
		return
	}

	xhtml_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ALTERNATIVE_ID:
		toT := CopyBranchALTERNATIVE_ID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		toT := CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_DATE:
		toT := CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		toT := CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		toT := CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_REAL:
		toT := CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_STRING:
		toT := CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_XHTML:
		toT := CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		toT := CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_DATE:
		toT := CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		toT := CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_INTEGER:
		toT := CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_REAL:
		toT := CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_STRING:
		toT := CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_XHTML:
		toT := CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ALTERNATIVE_ID:
		toT := CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		toT := CopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		toT := CopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_DATE:
		toT := CopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		toT := CopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_INTEGER:
		toT := CopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_REAL:
		toT := CopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_STRING:
		toT := CopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_XHTML:
		toT := CopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		toT := CopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_CHILDREN:
		toT := CopyBranchA_CHILDREN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_CORE_CONTENT:
		toT := CopyBranchA_CORE_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPES:
		toT := CopyBranchA_DATATYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_DATE_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_REAL_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_STRING_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		toT := CopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_EDITABLE_ATTS:
		toT := CopyBranchA_EDITABLE_ATTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ENUM_VALUE_REF:
		toT := CopyBranchA_ENUM_VALUE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_OBJECT:
		toT := CopyBranchA_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_PROPERTIES:
		toT := CopyBranchA_PROPERTIES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_RELATION_GROUP_TYPE_REF:
		toT := CopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SOURCE_SPECIFICATION_1:
		toT := CopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFICATIONS:
		toT := CopyBranchA_SPECIFICATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFICATION_TYPE_REF:
		toT := CopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFIED_VALUES:
		toT := CopyBranchA_SPECIFIED_VALUES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_ATTRIBUTES:
		toT := CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_OBJECTS:
		toT := CopyBranchA_SPEC_OBJECTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_OBJECT_TYPE_REF:
		toT := CopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATIONS:
		toT := CopyBranchA_SPEC_RELATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_GROUPS:
		toT := CopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_REF:
		toT := CopyBranchA_SPEC_RELATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_TYPE_REF:
		toT := CopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_TYPES:
		toT := CopyBranchA_SPEC_TYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TARGET_1:
		toT := CopyBranchA_TARGET_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_THE_HEADER:
		toT := CopyBranchA_THE_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TOOL_EXTENSIONS:
		toT := CopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_BOOLEAN:
		toT := CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_DATE:
		toT := CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_ENUMERATION:
		toT := CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_INTEGER:
		toT := CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_REAL:
		toT := CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_STRING:
		toT := CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_XHTML:
		toT := CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EMBEDDED_VALUE:
		toT := CopyBranchEMBEDDED_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ENUM_VALUE:
		toT := CopyBranchENUM_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP:
		toT := CopyBranchRELATION_GROUP(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP_TYPE:
		toT := CopyBranchRELATION_GROUP_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF:
		toT := CopyBranchREQ_IF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_CONTENT:
		toT := CopyBranchREQ_IF_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_HEADER:
		toT := CopyBranchREQ_IF_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_TOOL_EXTENSION:
		toT := CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION:
		toT := CopyBranchSPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION_TYPE:
		toT := CopyBranchSPECIFICATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_HIERARCHY:
		toT := CopyBranchSPEC_HIERARCHY(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT:
		toT := CopyBranchSPEC_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT_TYPE:
		toT := CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION:
		toT := CopyBranchSPEC_RELATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION_TYPE:
		toT := CopyBranchSPEC_RELATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XHTML_CONTENT:
		toT := CopyBranchXHTML_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchALTERNATIVE_ID(mapOrigCopy map[any]any, alternative_idFrom *ALTERNATIVE_ID) (alternative_idTo *ALTERNATIVE_ID) {

	// alternative_idFrom has already been copied
	if _alternative_idTo, ok := mapOrigCopy[alternative_idFrom]; ok {
		alternative_idTo = _alternative_idTo.(*ALTERNATIVE_ID)
		return
	}

	alternative_idTo = new(ALTERNATIVE_ID)
	mapOrigCopy[alternative_idFrom] = alternative_idTo
	alternative_idFrom.CopyBasicFields(alternative_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, attribute_definition_booleanFrom *ATTRIBUTE_DEFINITION_BOOLEAN) (attribute_definition_booleanTo *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// attribute_definition_booleanFrom has already been copied
	if _attribute_definition_booleanTo, ok := mapOrigCopy[attribute_definition_booleanFrom]; ok {
		attribute_definition_booleanTo = _attribute_definition_booleanTo.(*ATTRIBUTE_DEFINITION_BOOLEAN)
		return
	}

	attribute_definition_booleanTo = new(ATTRIBUTE_DEFINITION_BOOLEAN)
	mapOrigCopy[attribute_definition_booleanFrom] = attribute_definition_booleanTo
	attribute_definition_booleanFrom.CopyBasicFields(attribute_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_booleanFrom.ALTERNATIVE_ID != nil {
		attribute_definition_booleanTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_booleanFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_booleanFrom.DEFAULT_VALUE != nil {
		attribute_definition_booleanTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, attribute_definition_booleanFrom.DEFAULT_VALUE)
	}
	if attribute_definition_booleanFrom.TYPE != nil {
		attribute_definition_booleanTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy, attribute_definition_booleanFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy map[any]any, attribute_definition_dateFrom *ATTRIBUTE_DEFINITION_DATE) (attribute_definition_dateTo *ATTRIBUTE_DEFINITION_DATE) {

	// attribute_definition_dateFrom has already been copied
	if _attribute_definition_dateTo, ok := mapOrigCopy[attribute_definition_dateFrom]; ok {
		attribute_definition_dateTo = _attribute_definition_dateTo.(*ATTRIBUTE_DEFINITION_DATE)
		return
	}

	attribute_definition_dateTo = new(ATTRIBUTE_DEFINITION_DATE)
	mapOrigCopy[attribute_definition_dateFrom] = attribute_definition_dateTo
	attribute_definition_dateFrom.CopyBasicFields(attribute_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_dateFrom.ALTERNATIVE_ID != nil {
		attribute_definition_dateTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_dateFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_dateFrom.DEFAULT_VALUE != nil {
		attribute_definition_dateTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy, attribute_definition_dateFrom.DEFAULT_VALUE)
	}
	if attribute_definition_dateFrom.TYPE != nil {
		attribute_definition_dateTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy, attribute_definition_dateFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, attribute_definition_enumerationFrom *ATTRIBUTE_DEFINITION_ENUMERATION) (attribute_definition_enumerationTo *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// attribute_definition_enumerationFrom has already been copied
	if _attribute_definition_enumerationTo, ok := mapOrigCopy[attribute_definition_enumerationFrom]; ok {
		attribute_definition_enumerationTo = _attribute_definition_enumerationTo.(*ATTRIBUTE_DEFINITION_ENUMERATION)
		return
	}

	attribute_definition_enumerationTo = new(ATTRIBUTE_DEFINITION_ENUMERATION)
	mapOrigCopy[attribute_definition_enumerationFrom] = attribute_definition_enumerationTo
	attribute_definition_enumerationFrom.CopyBasicFields(attribute_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumerationFrom.ALTERNATIVE_ID != nil {
		attribute_definition_enumerationTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_enumerationFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_enumerationFrom.DEFAULT_VALUE != nil {
		attribute_definition_enumerationTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, attribute_definition_enumerationFrom.DEFAULT_VALUE)
	}
	if attribute_definition_enumerationFrom.TYPE != nil {
		attribute_definition_enumerationTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy, attribute_definition_enumerationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy map[any]any, attribute_definition_integerFrom *ATTRIBUTE_DEFINITION_INTEGER) (attribute_definition_integerTo *ATTRIBUTE_DEFINITION_INTEGER) {

	// attribute_definition_integerFrom has already been copied
	if _attribute_definition_integerTo, ok := mapOrigCopy[attribute_definition_integerFrom]; ok {
		attribute_definition_integerTo = _attribute_definition_integerTo.(*ATTRIBUTE_DEFINITION_INTEGER)
		return
	}

	attribute_definition_integerTo = new(ATTRIBUTE_DEFINITION_INTEGER)
	mapOrigCopy[attribute_definition_integerFrom] = attribute_definition_integerTo
	attribute_definition_integerFrom.CopyBasicFields(attribute_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integerFrom.ALTERNATIVE_ID != nil {
		attribute_definition_integerTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_integerFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_integerFrom.DEFAULT_VALUE != nil {
		attribute_definition_integerTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy, attribute_definition_integerFrom.DEFAULT_VALUE)
	}
	if attribute_definition_integerFrom.TYPE != nil {
		attribute_definition_integerTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy, attribute_definition_integerFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy map[any]any, attribute_definition_realFrom *ATTRIBUTE_DEFINITION_REAL) (attribute_definition_realTo *ATTRIBUTE_DEFINITION_REAL) {

	// attribute_definition_realFrom has already been copied
	if _attribute_definition_realTo, ok := mapOrigCopy[attribute_definition_realFrom]; ok {
		attribute_definition_realTo = _attribute_definition_realTo.(*ATTRIBUTE_DEFINITION_REAL)
		return
	}

	attribute_definition_realTo = new(ATTRIBUTE_DEFINITION_REAL)
	mapOrigCopy[attribute_definition_realFrom] = attribute_definition_realTo
	attribute_definition_realFrom.CopyBasicFields(attribute_definition_realTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_realFrom.ALTERNATIVE_ID != nil {
		attribute_definition_realTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_realFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_realFrom.DEFAULT_VALUE != nil {
		attribute_definition_realTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy, attribute_definition_realFrom.DEFAULT_VALUE)
	}
	if attribute_definition_realFrom.TYPE != nil {
		attribute_definition_realTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy, attribute_definition_realFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy map[any]any, attribute_definition_stringFrom *ATTRIBUTE_DEFINITION_STRING) (attribute_definition_stringTo *ATTRIBUTE_DEFINITION_STRING) {

	// attribute_definition_stringFrom has already been copied
	if _attribute_definition_stringTo, ok := mapOrigCopy[attribute_definition_stringFrom]; ok {
		attribute_definition_stringTo = _attribute_definition_stringTo.(*ATTRIBUTE_DEFINITION_STRING)
		return
	}

	attribute_definition_stringTo = new(ATTRIBUTE_DEFINITION_STRING)
	mapOrigCopy[attribute_definition_stringFrom] = attribute_definition_stringTo
	attribute_definition_stringFrom.CopyBasicFields(attribute_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_stringFrom.ALTERNATIVE_ID != nil {
		attribute_definition_stringTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_stringFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_stringFrom.DEFAULT_VALUE != nil {
		attribute_definition_stringTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy, attribute_definition_stringFrom.DEFAULT_VALUE)
	}
	if attribute_definition_stringFrom.TYPE != nil {
		attribute_definition_stringTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy, attribute_definition_stringFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy map[any]any, attribute_definition_xhtmlFrom *ATTRIBUTE_DEFINITION_XHTML) (attribute_definition_xhtmlTo *ATTRIBUTE_DEFINITION_XHTML) {

	// attribute_definition_xhtmlFrom has already been copied
	if _attribute_definition_xhtmlTo, ok := mapOrigCopy[attribute_definition_xhtmlFrom]; ok {
		attribute_definition_xhtmlTo = _attribute_definition_xhtmlTo.(*ATTRIBUTE_DEFINITION_XHTML)
		return
	}

	attribute_definition_xhtmlTo = new(ATTRIBUTE_DEFINITION_XHTML)
	mapOrigCopy[attribute_definition_xhtmlFrom] = attribute_definition_xhtmlTo
	attribute_definition_xhtmlFrom.CopyBasicFields(attribute_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtmlFrom.ALTERNATIVE_ID != nil {
		attribute_definition_xhtmlTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_xhtmlFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtmlFrom.DEFAULT_VALUE != nil {
		attribute_definition_xhtmlTo.DEFAULT_VALUE = CopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy, attribute_definition_xhtmlFrom.DEFAULT_VALUE)
	}
	if attribute_definition_xhtmlFrom.TYPE != nil {
		attribute_definition_xhtmlTo.TYPE = CopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy, attribute_definition_xhtmlFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy map[any]any, attribute_value_booleanFrom *ATTRIBUTE_VALUE_BOOLEAN) (attribute_value_booleanTo *ATTRIBUTE_VALUE_BOOLEAN) {

	// attribute_value_booleanFrom has already been copied
	if _attribute_value_booleanTo, ok := mapOrigCopy[attribute_value_booleanFrom]; ok {
		attribute_value_booleanTo = _attribute_value_booleanTo.(*ATTRIBUTE_VALUE_BOOLEAN)
		return
	}

	attribute_value_booleanTo = new(ATTRIBUTE_VALUE_BOOLEAN)
	mapOrigCopy[attribute_value_booleanFrom] = attribute_value_booleanTo
	attribute_value_booleanFrom.CopyBasicFields(attribute_value_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_booleanFrom.DEFINITION != nil {
		attribute_value_booleanTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy, attribute_value_booleanFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy map[any]any, attribute_value_dateFrom *ATTRIBUTE_VALUE_DATE) (attribute_value_dateTo *ATTRIBUTE_VALUE_DATE) {

	// attribute_value_dateFrom has already been copied
	if _attribute_value_dateTo, ok := mapOrigCopy[attribute_value_dateFrom]; ok {
		attribute_value_dateTo = _attribute_value_dateTo.(*ATTRIBUTE_VALUE_DATE)
		return
	}

	attribute_value_dateTo = new(ATTRIBUTE_VALUE_DATE)
	mapOrigCopy[attribute_value_dateFrom] = attribute_value_dateTo
	attribute_value_dateFrom.CopyBasicFields(attribute_value_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_dateFrom.DEFINITION != nil {
		attribute_value_dateTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy, attribute_value_dateFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy map[any]any, attribute_value_enumerationFrom *ATTRIBUTE_VALUE_ENUMERATION) (attribute_value_enumerationTo *ATTRIBUTE_VALUE_ENUMERATION) {

	// attribute_value_enumerationFrom has already been copied
	if _attribute_value_enumerationTo, ok := mapOrigCopy[attribute_value_enumerationFrom]; ok {
		attribute_value_enumerationTo = _attribute_value_enumerationTo.(*ATTRIBUTE_VALUE_ENUMERATION)
		return
	}

	attribute_value_enumerationTo = new(ATTRIBUTE_VALUE_ENUMERATION)
	mapOrigCopy[attribute_value_enumerationFrom] = attribute_value_enumerationTo
	attribute_value_enumerationFrom.CopyBasicFields(attribute_value_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumerationFrom.DEFINITION != nil {
		attribute_value_enumerationTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy, attribute_value_enumerationFrom.DEFINITION)
	}
	if attribute_value_enumerationFrom.VALUES != nil {
		attribute_value_enumerationTo.VALUES = CopyBranchA_ENUM_VALUE_REF(mapOrigCopy, attribute_value_enumerationFrom.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy map[any]any, attribute_value_integerFrom *ATTRIBUTE_VALUE_INTEGER) (attribute_value_integerTo *ATTRIBUTE_VALUE_INTEGER) {

	// attribute_value_integerFrom has already been copied
	if _attribute_value_integerTo, ok := mapOrigCopy[attribute_value_integerFrom]; ok {
		attribute_value_integerTo = _attribute_value_integerTo.(*ATTRIBUTE_VALUE_INTEGER)
		return
	}

	attribute_value_integerTo = new(ATTRIBUTE_VALUE_INTEGER)
	mapOrigCopy[attribute_value_integerFrom] = attribute_value_integerTo
	attribute_value_integerFrom.CopyBasicFields(attribute_value_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integerFrom.DEFINITION != nil {
		attribute_value_integerTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy, attribute_value_integerFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy map[any]any, attribute_value_realFrom *ATTRIBUTE_VALUE_REAL) (attribute_value_realTo *ATTRIBUTE_VALUE_REAL) {

	// attribute_value_realFrom has already been copied
	if _attribute_value_realTo, ok := mapOrigCopy[attribute_value_realFrom]; ok {
		attribute_value_realTo = _attribute_value_realTo.(*ATTRIBUTE_VALUE_REAL)
		return
	}

	attribute_value_realTo = new(ATTRIBUTE_VALUE_REAL)
	mapOrigCopy[attribute_value_realFrom] = attribute_value_realTo
	attribute_value_realFrom.CopyBasicFields(attribute_value_realTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_realFrom.DEFINITION != nil {
		attribute_value_realTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy, attribute_value_realFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy map[any]any, attribute_value_stringFrom *ATTRIBUTE_VALUE_STRING) (attribute_value_stringTo *ATTRIBUTE_VALUE_STRING) {

	// attribute_value_stringFrom has already been copied
	if _attribute_value_stringTo, ok := mapOrigCopy[attribute_value_stringFrom]; ok {
		attribute_value_stringTo = _attribute_value_stringTo.(*ATTRIBUTE_VALUE_STRING)
		return
	}

	attribute_value_stringTo = new(ATTRIBUTE_VALUE_STRING)
	mapOrigCopy[attribute_value_stringFrom] = attribute_value_stringTo
	attribute_value_stringFrom.CopyBasicFields(attribute_value_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_stringFrom.DEFINITION != nil {
		attribute_value_stringTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy, attribute_value_stringFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy map[any]any, attribute_value_xhtmlFrom *ATTRIBUTE_VALUE_XHTML) (attribute_value_xhtmlTo *ATTRIBUTE_VALUE_XHTML) {

	// attribute_value_xhtmlFrom has already been copied
	if _attribute_value_xhtmlTo, ok := mapOrigCopy[attribute_value_xhtmlFrom]; ok {
		attribute_value_xhtmlTo = _attribute_value_xhtmlTo.(*ATTRIBUTE_VALUE_XHTML)
		return
	}

	attribute_value_xhtmlTo = new(ATTRIBUTE_VALUE_XHTML)
	mapOrigCopy[attribute_value_xhtmlFrom] = attribute_value_xhtmlTo
	attribute_value_xhtmlFrom.CopyBasicFields(attribute_value_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtmlFrom.THE_VALUE != nil {
		attribute_value_xhtmlTo.THE_VALUE = CopyBranchXHTML_CONTENT(mapOrigCopy, attribute_value_xhtmlFrom.THE_VALUE)
	}
	if attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE != nil {
		attribute_value_xhtmlTo.THE_ORIGINAL_VALUE = CopyBranchXHTML_CONTENT(mapOrigCopy, attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtmlFrom.DEFINITION != nil {
		attribute_value_xhtmlTo.DEFINITION = CopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy, attribute_value_xhtmlFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ALTERNATIVE_ID(mapOrigCopy map[any]any, a_alternative_idFrom *A_ALTERNATIVE_ID) (a_alternative_idTo *A_ALTERNATIVE_ID) {

	// a_alternative_idFrom has already been copied
	if _a_alternative_idTo, ok := mapOrigCopy[a_alternative_idFrom]; ok {
		a_alternative_idTo = _a_alternative_idTo.(*A_ALTERNATIVE_ID)
		return
	}

	a_alternative_idTo = new(A_ALTERNATIVE_ID)
	mapOrigCopy[a_alternative_idFrom] = a_alternative_idTo
	a_alternative_idFrom.CopyBasicFields(a_alternative_idTo)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_idFrom.ALTERNATIVE_ID != nil {
		a_alternative_idTo.ALTERNATIVE_ID = CopyBranchALTERNATIVE_ID(mapOrigCopy, a_alternative_idFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy map[any]any, a_attribute_definition_boolean_refFrom *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) (a_attribute_definition_boolean_refTo *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// a_attribute_definition_boolean_refFrom has already been copied
	if _a_attribute_definition_boolean_refTo, ok := mapOrigCopy[a_attribute_definition_boolean_refFrom]; ok {
		a_attribute_definition_boolean_refTo = _a_attribute_definition_boolean_refTo.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		return
	}

	a_attribute_definition_boolean_refTo = new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	mapOrigCopy[a_attribute_definition_boolean_refFrom] = a_attribute_definition_boolean_refTo
	a_attribute_definition_boolean_refFrom.CopyBasicFields(a_attribute_definition_boolean_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy map[any]any, a_attribute_definition_date_refFrom *A_ATTRIBUTE_DEFINITION_DATE_REF) (a_attribute_definition_date_refTo *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// a_attribute_definition_date_refFrom has already been copied
	if _a_attribute_definition_date_refTo, ok := mapOrigCopy[a_attribute_definition_date_refFrom]; ok {
		a_attribute_definition_date_refTo = _a_attribute_definition_date_refTo.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		return
	}

	a_attribute_definition_date_refTo = new(A_ATTRIBUTE_DEFINITION_DATE_REF)
	mapOrigCopy[a_attribute_definition_date_refFrom] = a_attribute_definition_date_refTo
	a_attribute_definition_date_refFrom.CopyBasicFields(a_attribute_definition_date_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy map[any]any, a_attribute_definition_enumeration_refFrom *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) (a_attribute_definition_enumeration_refTo *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// a_attribute_definition_enumeration_refFrom has already been copied
	if _a_attribute_definition_enumeration_refTo, ok := mapOrigCopy[a_attribute_definition_enumeration_refFrom]; ok {
		a_attribute_definition_enumeration_refTo = _a_attribute_definition_enumeration_refTo.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		return
	}

	a_attribute_definition_enumeration_refTo = new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	mapOrigCopy[a_attribute_definition_enumeration_refFrom] = a_attribute_definition_enumeration_refTo
	a_attribute_definition_enumeration_refFrom.CopyBasicFields(a_attribute_definition_enumeration_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy map[any]any, a_attribute_definition_integer_refFrom *A_ATTRIBUTE_DEFINITION_INTEGER_REF) (a_attribute_definition_integer_refTo *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// a_attribute_definition_integer_refFrom has already been copied
	if _a_attribute_definition_integer_refTo, ok := mapOrigCopy[a_attribute_definition_integer_refFrom]; ok {
		a_attribute_definition_integer_refTo = _a_attribute_definition_integer_refTo.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		return
	}

	a_attribute_definition_integer_refTo = new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	mapOrigCopy[a_attribute_definition_integer_refFrom] = a_attribute_definition_integer_refTo
	a_attribute_definition_integer_refFrom.CopyBasicFields(a_attribute_definition_integer_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy map[any]any, a_attribute_definition_real_refFrom *A_ATTRIBUTE_DEFINITION_REAL_REF) (a_attribute_definition_real_refTo *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// a_attribute_definition_real_refFrom has already been copied
	if _a_attribute_definition_real_refTo, ok := mapOrigCopy[a_attribute_definition_real_refFrom]; ok {
		a_attribute_definition_real_refTo = _a_attribute_definition_real_refTo.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		return
	}

	a_attribute_definition_real_refTo = new(A_ATTRIBUTE_DEFINITION_REAL_REF)
	mapOrigCopy[a_attribute_definition_real_refFrom] = a_attribute_definition_real_refTo
	a_attribute_definition_real_refFrom.CopyBasicFields(a_attribute_definition_real_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy map[any]any, a_attribute_definition_string_refFrom *A_ATTRIBUTE_DEFINITION_STRING_REF) (a_attribute_definition_string_refTo *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// a_attribute_definition_string_refFrom has already been copied
	if _a_attribute_definition_string_refTo, ok := mapOrigCopy[a_attribute_definition_string_refFrom]; ok {
		a_attribute_definition_string_refTo = _a_attribute_definition_string_refTo.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		return
	}

	a_attribute_definition_string_refTo = new(A_ATTRIBUTE_DEFINITION_STRING_REF)
	mapOrigCopy[a_attribute_definition_string_refFrom] = a_attribute_definition_string_refTo
	a_attribute_definition_string_refFrom.CopyBasicFields(a_attribute_definition_string_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy map[any]any, a_attribute_definition_xhtml_refFrom *A_ATTRIBUTE_DEFINITION_XHTML_REF) (a_attribute_definition_xhtml_refTo *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// a_attribute_definition_xhtml_refFrom has already been copied
	if _a_attribute_definition_xhtml_refTo, ok := mapOrigCopy[a_attribute_definition_xhtml_refFrom]; ok {
		a_attribute_definition_xhtml_refTo = _a_attribute_definition_xhtml_refTo.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		return
	}

	a_attribute_definition_xhtml_refTo = new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
	mapOrigCopy[a_attribute_definition_xhtml_refFrom] = a_attribute_definition_xhtml_refTo
	a_attribute_definition_xhtml_refFrom.CopyBasicFields(a_attribute_definition_xhtml_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy map[any]any, a_attribute_value_booleanFrom *A_ATTRIBUTE_VALUE_BOOLEAN) (a_attribute_value_booleanTo *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// a_attribute_value_booleanFrom has already been copied
	if _a_attribute_value_booleanTo, ok := mapOrigCopy[a_attribute_value_booleanFrom]; ok {
		a_attribute_value_booleanTo = _a_attribute_value_booleanTo.(*A_ATTRIBUTE_VALUE_BOOLEAN)
		return
	}

	a_attribute_value_booleanTo = new(A_ATTRIBUTE_VALUE_BOOLEAN)
	mapOrigCopy[a_attribute_value_booleanFrom] = a_attribute_value_booleanTo
	a_attribute_value_booleanFrom.CopyBasicFields(a_attribute_value_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_booleanFrom.ATTRIBUTE_VALUE_BOOLEAN != nil {
		a_attribute_value_booleanTo.ATTRIBUTE_VALUE_BOOLEAN = CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, a_attribute_value_booleanFrom.ATTRIBUTE_VALUE_BOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy map[any]any, a_attribute_value_dateFrom *A_ATTRIBUTE_VALUE_DATE) (a_attribute_value_dateTo *A_ATTRIBUTE_VALUE_DATE) {

	// a_attribute_value_dateFrom has already been copied
	if _a_attribute_value_dateTo, ok := mapOrigCopy[a_attribute_value_dateFrom]; ok {
		a_attribute_value_dateTo = _a_attribute_value_dateTo.(*A_ATTRIBUTE_VALUE_DATE)
		return
	}

	a_attribute_value_dateTo = new(A_ATTRIBUTE_VALUE_DATE)
	mapOrigCopy[a_attribute_value_dateFrom] = a_attribute_value_dateTo
	a_attribute_value_dateFrom.CopyBasicFields(a_attribute_value_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_dateFrom.ATTRIBUTE_VALUE_DATE != nil {
		a_attribute_value_dateTo.ATTRIBUTE_VALUE_DATE = CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, a_attribute_value_dateFrom.ATTRIBUTE_VALUE_DATE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy map[any]any, a_attribute_value_enumerationFrom *A_ATTRIBUTE_VALUE_ENUMERATION) (a_attribute_value_enumerationTo *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// a_attribute_value_enumerationFrom has already been copied
	if _a_attribute_value_enumerationTo, ok := mapOrigCopy[a_attribute_value_enumerationFrom]; ok {
		a_attribute_value_enumerationTo = _a_attribute_value_enumerationTo.(*A_ATTRIBUTE_VALUE_ENUMERATION)
		return
	}

	a_attribute_value_enumerationTo = new(A_ATTRIBUTE_VALUE_ENUMERATION)
	mapOrigCopy[a_attribute_value_enumerationFrom] = a_attribute_value_enumerationTo
	a_attribute_value_enumerationFrom.CopyBasicFields(a_attribute_value_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_enumerationFrom.ATTRIBUTE_VALUE_ENUMERATION != nil {
		a_attribute_value_enumerationTo.ATTRIBUTE_VALUE_ENUMERATION = CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, a_attribute_value_enumerationFrom.ATTRIBUTE_VALUE_ENUMERATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy map[any]any, a_attribute_value_integerFrom *A_ATTRIBUTE_VALUE_INTEGER) (a_attribute_value_integerTo *A_ATTRIBUTE_VALUE_INTEGER) {

	// a_attribute_value_integerFrom has already been copied
	if _a_attribute_value_integerTo, ok := mapOrigCopy[a_attribute_value_integerFrom]; ok {
		a_attribute_value_integerTo = _a_attribute_value_integerTo.(*A_ATTRIBUTE_VALUE_INTEGER)
		return
	}

	a_attribute_value_integerTo = new(A_ATTRIBUTE_VALUE_INTEGER)
	mapOrigCopy[a_attribute_value_integerFrom] = a_attribute_value_integerTo
	a_attribute_value_integerFrom.CopyBasicFields(a_attribute_value_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_integerFrom.ATTRIBUTE_VALUE_INTEGER != nil {
		a_attribute_value_integerTo.ATTRIBUTE_VALUE_INTEGER = CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, a_attribute_value_integerFrom.ATTRIBUTE_VALUE_INTEGER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy map[any]any, a_attribute_value_realFrom *A_ATTRIBUTE_VALUE_REAL) (a_attribute_value_realTo *A_ATTRIBUTE_VALUE_REAL) {

	// a_attribute_value_realFrom has already been copied
	if _a_attribute_value_realTo, ok := mapOrigCopy[a_attribute_value_realFrom]; ok {
		a_attribute_value_realTo = _a_attribute_value_realTo.(*A_ATTRIBUTE_VALUE_REAL)
		return
	}

	a_attribute_value_realTo = new(A_ATTRIBUTE_VALUE_REAL)
	mapOrigCopy[a_attribute_value_realFrom] = a_attribute_value_realTo
	a_attribute_value_realFrom.CopyBasicFields(a_attribute_value_realTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_realFrom.ATTRIBUTE_VALUE_REAL != nil {
		a_attribute_value_realTo.ATTRIBUTE_VALUE_REAL = CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, a_attribute_value_realFrom.ATTRIBUTE_VALUE_REAL)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy map[any]any, a_attribute_value_stringFrom *A_ATTRIBUTE_VALUE_STRING) (a_attribute_value_stringTo *A_ATTRIBUTE_VALUE_STRING) {

	// a_attribute_value_stringFrom has already been copied
	if _a_attribute_value_stringTo, ok := mapOrigCopy[a_attribute_value_stringFrom]; ok {
		a_attribute_value_stringTo = _a_attribute_value_stringTo.(*A_ATTRIBUTE_VALUE_STRING)
		return
	}

	a_attribute_value_stringTo = new(A_ATTRIBUTE_VALUE_STRING)
	mapOrigCopy[a_attribute_value_stringFrom] = a_attribute_value_stringTo
	a_attribute_value_stringFrom.CopyBasicFields(a_attribute_value_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_stringFrom.ATTRIBUTE_VALUE_STRING != nil {
		a_attribute_value_stringTo.ATTRIBUTE_VALUE_STRING = CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, a_attribute_value_stringFrom.ATTRIBUTE_VALUE_STRING)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy map[any]any, a_attribute_value_xhtmlFrom *A_ATTRIBUTE_VALUE_XHTML) (a_attribute_value_xhtmlTo *A_ATTRIBUTE_VALUE_XHTML) {

	// a_attribute_value_xhtmlFrom has already been copied
	if _a_attribute_value_xhtmlTo, ok := mapOrigCopy[a_attribute_value_xhtmlFrom]; ok {
		a_attribute_value_xhtmlTo = _a_attribute_value_xhtmlTo.(*A_ATTRIBUTE_VALUE_XHTML)
		return
	}

	a_attribute_value_xhtmlTo = new(A_ATTRIBUTE_VALUE_XHTML)
	mapOrigCopy[a_attribute_value_xhtmlFrom] = a_attribute_value_xhtmlTo
	a_attribute_value_xhtmlFrom.CopyBasicFields(a_attribute_value_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtmlFrom.ATTRIBUTE_VALUE_XHTML != nil {
		a_attribute_value_xhtmlTo.ATTRIBUTE_VALUE_XHTML = CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, a_attribute_value_xhtmlFrom.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy map[any]any, a_attribute_value_xhtml_1From *A_ATTRIBUTE_VALUE_XHTML_1) (a_attribute_value_xhtml_1To *A_ATTRIBUTE_VALUE_XHTML_1) {

	// a_attribute_value_xhtml_1From has already been copied
	if _a_attribute_value_xhtml_1To, ok := mapOrigCopy[a_attribute_value_xhtml_1From]; ok {
		a_attribute_value_xhtml_1To = _a_attribute_value_xhtml_1To.(*A_ATTRIBUTE_VALUE_XHTML_1)
		return
	}

	a_attribute_value_xhtml_1To = new(A_ATTRIBUTE_VALUE_XHTML_1)
	mapOrigCopy[a_attribute_value_xhtml_1From] = a_attribute_value_xhtml_1To
	a_attribute_value_xhtml_1From.CopyBasicFields(a_attribute_value_xhtml_1To)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_BOOLEAN != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_BOOLEAN = CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_BOOLEAN)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_DATE != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_DATE = CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_DATE)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_ENUMERATION != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_ENUMERATION = CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_ENUMERATION)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_INTEGER != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_INTEGER = CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_INTEGER)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_REAL != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_REAL = CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_REAL)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_STRING != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_STRING = CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_STRING)
	}
	if a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_XHTML != nil {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_XHTML = CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_CHILDREN(mapOrigCopy map[any]any, a_childrenFrom *A_CHILDREN) (a_childrenTo *A_CHILDREN) {

	// a_childrenFrom has already been copied
	if _a_childrenTo, ok := mapOrigCopy[a_childrenFrom]; ok {
		a_childrenTo = _a_childrenTo.(*A_CHILDREN)
		return
	}

	a_childrenTo = new(A_CHILDREN)
	mapOrigCopy[a_childrenFrom] = a_childrenTo
	a_childrenFrom.CopyBasicFields(a_childrenTo)

	//insertion point for the staging of instances referenced by pointers
	if a_childrenFrom.SPEC_HIERARCHY != nil {
		a_childrenTo.SPEC_HIERARCHY = CopyBranchSPEC_HIERARCHY(mapOrigCopy, a_childrenFrom.SPEC_HIERARCHY)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_CORE_CONTENT(mapOrigCopy map[any]any, a_core_contentFrom *A_CORE_CONTENT) (a_core_contentTo *A_CORE_CONTENT) {

	// a_core_contentFrom has already been copied
	if _a_core_contentTo, ok := mapOrigCopy[a_core_contentFrom]; ok {
		a_core_contentTo = _a_core_contentTo.(*A_CORE_CONTENT)
		return
	}

	a_core_contentTo = new(A_CORE_CONTENT)
	mapOrigCopy[a_core_contentFrom] = a_core_contentTo
	a_core_contentFrom.CopyBasicFields(a_core_contentTo)

	//insertion point for the staging of instances referenced by pointers
	if a_core_contentFrom.REQ_IF_CONTENT != nil {
		a_core_contentTo.REQ_IF_CONTENT = CopyBranchREQ_IF_CONTENT(mapOrigCopy, a_core_contentFrom.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPES(mapOrigCopy map[any]any, a_datatypesFrom *A_DATATYPES) (a_datatypesTo *A_DATATYPES) {

	// a_datatypesFrom has already been copied
	if _a_datatypesTo, ok := mapOrigCopy[a_datatypesFrom]; ok {
		a_datatypesTo = _a_datatypesTo.(*A_DATATYPES)
		return
	}

	a_datatypesTo = new(A_DATATYPES)
	mapOrigCopy[a_datatypesFrom] = a_datatypesTo
	a_datatypesFrom.CopyBasicFields(a_datatypesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_datatypesFrom.DATATYPE_DEFINITION_BOOLEAN != nil {
		a_datatypesTo.DATATYPE_DEFINITION_BOOLEAN = CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_BOOLEAN)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_DATE != nil {
		a_datatypesTo.DATATYPE_DEFINITION_DATE = CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_DATE)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_ENUMERATION != nil {
		a_datatypesTo.DATATYPE_DEFINITION_ENUMERATION = CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_ENUMERATION)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_INTEGER != nil {
		a_datatypesTo.DATATYPE_DEFINITION_INTEGER = CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_INTEGER)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_REAL != nil {
		a_datatypesTo.DATATYPE_DEFINITION_REAL = CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_REAL)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_STRING != nil {
		a_datatypesTo.DATATYPE_DEFINITION_STRING = CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_STRING)
	}
	if a_datatypesFrom.DATATYPE_DEFINITION_XHTML != nil {
		a_datatypesTo.DATATYPE_DEFINITION_XHTML = CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, a_datatypesFrom.DATATYPE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy map[any]any, a_datatype_definition_boolean_refFrom *A_DATATYPE_DEFINITION_BOOLEAN_REF) (a_datatype_definition_boolean_refTo *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// a_datatype_definition_boolean_refFrom has already been copied
	if _a_datatype_definition_boolean_refTo, ok := mapOrigCopy[a_datatype_definition_boolean_refFrom]; ok {
		a_datatype_definition_boolean_refTo = _a_datatype_definition_boolean_refTo.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		return
	}

	a_datatype_definition_boolean_refTo = new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
	mapOrigCopy[a_datatype_definition_boolean_refFrom] = a_datatype_definition_boolean_refTo
	a_datatype_definition_boolean_refFrom.CopyBasicFields(a_datatype_definition_boolean_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy map[any]any, a_datatype_definition_date_refFrom *A_DATATYPE_DEFINITION_DATE_REF) (a_datatype_definition_date_refTo *A_DATATYPE_DEFINITION_DATE_REF) {

	// a_datatype_definition_date_refFrom has already been copied
	if _a_datatype_definition_date_refTo, ok := mapOrigCopy[a_datatype_definition_date_refFrom]; ok {
		a_datatype_definition_date_refTo = _a_datatype_definition_date_refTo.(*A_DATATYPE_DEFINITION_DATE_REF)
		return
	}

	a_datatype_definition_date_refTo = new(A_DATATYPE_DEFINITION_DATE_REF)
	mapOrigCopy[a_datatype_definition_date_refFrom] = a_datatype_definition_date_refTo
	a_datatype_definition_date_refFrom.CopyBasicFields(a_datatype_definition_date_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy map[any]any, a_datatype_definition_enumeration_refFrom *A_DATATYPE_DEFINITION_ENUMERATION_REF) (a_datatype_definition_enumeration_refTo *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// a_datatype_definition_enumeration_refFrom has already been copied
	if _a_datatype_definition_enumeration_refTo, ok := mapOrigCopy[a_datatype_definition_enumeration_refFrom]; ok {
		a_datatype_definition_enumeration_refTo = _a_datatype_definition_enumeration_refTo.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		return
	}

	a_datatype_definition_enumeration_refTo = new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
	mapOrigCopy[a_datatype_definition_enumeration_refFrom] = a_datatype_definition_enumeration_refTo
	a_datatype_definition_enumeration_refFrom.CopyBasicFields(a_datatype_definition_enumeration_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy map[any]any, a_datatype_definition_integer_refFrom *A_DATATYPE_DEFINITION_INTEGER_REF) (a_datatype_definition_integer_refTo *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// a_datatype_definition_integer_refFrom has already been copied
	if _a_datatype_definition_integer_refTo, ok := mapOrigCopy[a_datatype_definition_integer_refFrom]; ok {
		a_datatype_definition_integer_refTo = _a_datatype_definition_integer_refTo.(*A_DATATYPE_DEFINITION_INTEGER_REF)
		return
	}

	a_datatype_definition_integer_refTo = new(A_DATATYPE_DEFINITION_INTEGER_REF)
	mapOrigCopy[a_datatype_definition_integer_refFrom] = a_datatype_definition_integer_refTo
	a_datatype_definition_integer_refFrom.CopyBasicFields(a_datatype_definition_integer_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy map[any]any, a_datatype_definition_real_refFrom *A_DATATYPE_DEFINITION_REAL_REF) (a_datatype_definition_real_refTo *A_DATATYPE_DEFINITION_REAL_REF) {

	// a_datatype_definition_real_refFrom has already been copied
	if _a_datatype_definition_real_refTo, ok := mapOrigCopy[a_datatype_definition_real_refFrom]; ok {
		a_datatype_definition_real_refTo = _a_datatype_definition_real_refTo.(*A_DATATYPE_DEFINITION_REAL_REF)
		return
	}

	a_datatype_definition_real_refTo = new(A_DATATYPE_DEFINITION_REAL_REF)
	mapOrigCopy[a_datatype_definition_real_refFrom] = a_datatype_definition_real_refTo
	a_datatype_definition_real_refFrom.CopyBasicFields(a_datatype_definition_real_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy map[any]any, a_datatype_definition_string_refFrom *A_DATATYPE_DEFINITION_STRING_REF) (a_datatype_definition_string_refTo *A_DATATYPE_DEFINITION_STRING_REF) {

	// a_datatype_definition_string_refFrom has already been copied
	if _a_datatype_definition_string_refTo, ok := mapOrigCopy[a_datatype_definition_string_refFrom]; ok {
		a_datatype_definition_string_refTo = _a_datatype_definition_string_refTo.(*A_DATATYPE_DEFINITION_STRING_REF)
		return
	}

	a_datatype_definition_string_refTo = new(A_DATATYPE_DEFINITION_STRING_REF)
	mapOrigCopy[a_datatype_definition_string_refFrom] = a_datatype_definition_string_refTo
	a_datatype_definition_string_refFrom.CopyBasicFields(a_datatype_definition_string_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy map[any]any, a_datatype_definition_xhtml_refFrom *A_DATATYPE_DEFINITION_XHTML_REF) (a_datatype_definition_xhtml_refTo *A_DATATYPE_DEFINITION_XHTML_REF) {

	// a_datatype_definition_xhtml_refFrom has already been copied
	if _a_datatype_definition_xhtml_refTo, ok := mapOrigCopy[a_datatype_definition_xhtml_refFrom]; ok {
		a_datatype_definition_xhtml_refTo = _a_datatype_definition_xhtml_refTo.(*A_DATATYPE_DEFINITION_XHTML_REF)
		return
	}

	a_datatype_definition_xhtml_refTo = new(A_DATATYPE_DEFINITION_XHTML_REF)
	mapOrigCopy[a_datatype_definition_xhtml_refFrom] = a_datatype_definition_xhtml_refTo
	a_datatype_definition_xhtml_refFrom.CopyBasicFields(a_datatype_definition_xhtml_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_EDITABLE_ATTS(mapOrigCopy map[any]any, a_editable_attsFrom *A_EDITABLE_ATTS) (a_editable_attsTo *A_EDITABLE_ATTS) {

	// a_editable_attsFrom has already been copied
	if _a_editable_attsTo, ok := mapOrigCopy[a_editable_attsFrom]; ok {
		a_editable_attsTo = _a_editable_attsTo.(*A_EDITABLE_ATTS)
		return
	}

	a_editable_attsTo = new(A_EDITABLE_ATTS)
	mapOrigCopy[a_editable_attsFrom] = a_editable_attsTo
	a_editable_attsFrom.CopyBasicFields(a_editable_attsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_ENUM_VALUE_REF(mapOrigCopy map[any]any, a_enum_value_refFrom *A_ENUM_VALUE_REF) (a_enum_value_refTo *A_ENUM_VALUE_REF) {

	// a_enum_value_refFrom has already been copied
	if _a_enum_value_refTo, ok := mapOrigCopy[a_enum_value_refFrom]; ok {
		a_enum_value_refTo = _a_enum_value_refTo.(*A_ENUM_VALUE_REF)
		return
	}

	a_enum_value_refTo = new(A_ENUM_VALUE_REF)
	mapOrigCopy[a_enum_value_refFrom] = a_enum_value_refTo
	a_enum_value_refFrom.CopyBasicFields(a_enum_value_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_OBJECT(mapOrigCopy map[any]any, a_objectFrom *A_OBJECT) (a_objectTo *A_OBJECT) {

	// a_objectFrom has already been copied
	if _a_objectTo, ok := mapOrigCopy[a_objectFrom]; ok {
		a_objectTo = _a_objectTo.(*A_OBJECT)
		return
	}

	a_objectTo = new(A_OBJECT)
	mapOrigCopy[a_objectFrom] = a_objectTo
	a_objectFrom.CopyBasicFields(a_objectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_PROPERTIES(mapOrigCopy map[any]any, a_propertiesFrom *A_PROPERTIES) (a_propertiesTo *A_PROPERTIES) {

	// a_propertiesFrom has already been copied
	if _a_propertiesTo, ok := mapOrigCopy[a_propertiesFrom]; ok {
		a_propertiesTo = _a_propertiesTo.(*A_PROPERTIES)
		return
	}

	a_propertiesTo = new(A_PROPERTIES)
	mapOrigCopy[a_propertiesFrom] = a_propertiesTo
	a_propertiesFrom.CopyBasicFields(a_propertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_propertiesFrom.EMBEDDED_VALUE != nil {
		a_propertiesTo.EMBEDDED_VALUE = CopyBranchEMBEDDED_VALUE(mapOrigCopy, a_propertiesFrom.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy map[any]any, a_relation_group_type_refFrom *A_RELATION_GROUP_TYPE_REF) (a_relation_group_type_refTo *A_RELATION_GROUP_TYPE_REF) {

	// a_relation_group_type_refFrom has already been copied
	if _a_relation_group_type_refTo, ok := mapOrigCopy[a_relation_group_type_refFrom]; ok {
		a_relation_group_type_refTo = _a_relation_group_type_refTo.(*A_RELATION_GROUP_TYPE_REF)
		return
	}

	a_relation_group_type_refTo = new(A_RELATION_GROUP_TYPE_REF)
	mapOrigCopy[a_relation_group_type_refFrom] = a_relation_group_type_refTo
	a_relation_group_type_refFrom.CopyBasicFields(a_relation_group_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy map[any]any, a_source_specification_1From *A_SOURCE_SPECIFICATION_1) (a_source_specification_1To *A_SOURCE_SPECIFICATION_1) {

	// a_source_specification_1From has already been copied
	if _a_source_specification_1To, ok := mapOrigCopy[a_source_specification_1From]; ok {
		a_source_specification_1To = _a_source_specification_1To.(*A_SOURCE_SPECIFICATION_1)
		return
	}

	a_source_specification_1To = new(A_SOURCE_SPECIFICATION_1)
	mapOrigCopy[a_source_specification_1From] = a_source_specification_1To
	a_source_specification_1From.CopyBasicFields(a_source_specification_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPECIFICATIONS(mapOrigCopy map[any]any, a_specificationsFrom *A_SPECIFICATIONS) (a_specificationsTo *A_SPECIFICATIONS) {

	// a_specificationsFrom has already been copied
	if _a_specificationsTo, ok := mapOrigCopy[a_specificationsFrom]; ok {
		a_specificationsTo = _a_specificationsTo.(*A_SPECIFICATIONS)
		return
	}

	a_specificationsTo = new(A_SPECIFICATIONS)
	mapOrigCopy[a_specificationsFrom] = a_specificationsTo
	a_specificationsFrom.CopyBasicFields(a_specificationsTo)

	//insertion point for the staging of instances referenced by pointers
	if a_specificationsFrom.SPECIFICATION != nil {
		a_specificationsTo.SPECIFICATION = CopyBranchSPECIFICATION(mapOrigCopy, a_specificationsFrom.SPECIFICATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy map[any]any, a_specification_type_refFrom *A_SPECIFICATION_TYPE_REF) (a_specification_type_refTo *A_SPECIFICATION_TYPE_REF) {

	// a_specification_type_refFrom has already been copied
	if _a_specification_type_refTo, ok := mapOrigCopy[a_specification_type_refFrom]; ok {
		a_specification_type_refTo = _a_specification_type_refTo.(*A_SPECIFICATION_TYPE_REF)
		return
	}

	a_specification_type_refTo = new(A_SPECIFICATION_TYPE_REF)
	mapOrigCopy[a_specification_type_refFrom] = a_specification_type_refTo
	a_specification_type_refFrom.CopyBasicFields(a_specification_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPECIFIED_VALUES(mapOrigCopy map[any]any, a_specified_valuesFrom *A_SPECIFIED_VALUES) (a_specified_valuesTo *A_SPECIFIED_VALUES) {

	// a_specified_valuesFrom has already been copied
	if _a_specified_valuesTo, ok := mapOrigCopy[a_specified_valuesFrom]; ok {
		a_specified_valuesTo = _a_specified_valuesTo.(*A_SPECIFIED_VALUES)
		return
	}

	a_specified_valuesTo = new(A_SPECIFIED_VALUES)
	mapOrigCopy[a_specified_valuesFrom] = a_specified_valuesTo
	a_specified_valuesFrom.CopyBasicFields(a_specified_valuesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_specified_valuesFrom.ENUM_VALUE != nil {
		a_specified_valuesTo.ENUM_VALUE = CopyBranchENUM_VALUE(mapOrigCopy, a_specified_valuesFrom.ENUM_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy map[any]any, a_spec_attributesFrom *A_SPEC_ATTRIBUTES) (a_spec_attributesTo *A_SPEC_ATTRIBUTES) {

	// a_spec_attributesFrom has already been copied
	if _a_spec_attributesTo, ok := mapOrigCopy[a_spec_attributesFrom]; ok {
		a_spec_attributesTo = _a_spec_attributesTo.(*A_SPEC_ATTRIBUTES)
		return
	}

	a_spec_attributesTo = new(A_SPEC_ATTRIBUTES)
	mapOrigCopy[a_spec_attributesFrom] = a_spec_attributesTo
	a_spec_attributesFrom.CopyBasicFields(a_spec_attributesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_BOOLEAN = CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_BOOLEAN)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_DATE != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_DATE = CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_DATE)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_ENUMERATION = CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_ENUMERATION)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_INTEGER != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_INTEGER = CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_INTEGER)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_REAL != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_REAL = CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_REAL)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_STRING != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_STRING = CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_STRING)
	}
	if a_spec_attributesFrom.ATTRIBUTE_DEFINITION_XHTML != nil {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_XHTML = CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, a_spec_attributesFrom.ATTRIBUTE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_OBJECTS(mapOrigCopy map[any]any, a_spec_objectsFrom *A_SPEC_OBJECTS) (a_spec_objectsTo *A_SPEC_OBJECTS) {

	// a_spec_objectsFrom has already been copied
	if _a_spec_objectsTo, ok := mapOrigCopy[a_spec_objectsFrom]; ok {
		a_spec_objectsTo = _a_spec_objectsTo.(*A_SPEC_OBJECTS)
		return
	}

	a_spec_objectsTo = new(A_SPEC_OBJECTS)
	mapOrigCopy[a_spec_objectsFrom] = a_spec_objectsTo
	a_spec_objectsFrom.CopyBasicFields(a_spec_objectsTo)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_objectsFrom.SPEC_OBJECT != nil {
		a_spec_objectsTo.SPEC_OBJECT = CopyBranchSPEC_OBJECT(mapOrigCopy, a_spec_objectsFrom.SPEC_OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy map[any]any, a_spec_object_type_refFrom *A_SPEC_OBJECT_TYPE_REF) (a_spec_object_type_refTo *A_SPEC_OBJECT_TYPE_REF) {

	// a_spec_object_type_refFrom has already been copied
	if _a_spec_object_type_refTo, ok := mapOrigCopy[a_spec_object_type_refFrom]; ok {
		a_spec_object_type_refTo = _a_spec_object_type_refTo.(*A_SPEC_OBJECT_TYPE_REF)
		return
	}

	a_spec_object_type_refTo = new(A_SPEC_OBJECT_TYPE_REF)
	mapOrigCopy[a_spec_object_type_refFrom] = a_spec_object_type_refTo
	a_spec_object_type_refFrom.CopyBasicFields(a_spec_object_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_RELATIONS(mapOrigCopy map[any]any, a_spec_relationsFrom *A_SPEC_RELATIONS) (a_spec_relationsTo *A_SPEC_RELATIONS) {

	// a_spec_relationsFrom has already been copied
	if _a_spec_relationsTo, ok := mapOrigCopy[a_spec_relationsFrom]; ok {
		a_spec_relationsTo = _a_spec_relationsTo.(*A_SPEC_RELATIONS)
		return
	}

	a_spec_relationsTo = new(A_SPEC_RELATIONS)
	mapOrigCopy[a_spec_relationsFrom] = a_spec_relationsTo
	a_spec_relationsFrom.CopyBasicFields(a_spec_relationsTo)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relationsFrom.SPEC_RELATION != nil {
		a_spec_relationsTo.SPEC_RELATION = CopyBranchSPEC_RELATION(mapOrigCopy, a_spec_relationsFrom.SPEC_RELATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy map[any]any, a_spec_relation_groupsFrom *A_SPEC_RELATION_GROUPS) (a_spec_relation_groupsTo *A_SPEC_RELATION_GROUPS) {

	// a_spec_relation_groupsFrom has already been copied
	if _a_spec_relation_groupsTo, ok := mapOrigCopy[a_spec_relation_groupsFrom]; ok {
		a_spec_relation_groupsTo = _a_spec_relation_groupsTo.(*A_SPEC_RELATION_GROUPS)
		return
	}

	a_spec_relation_groupsTo = new(A_SPEC_RELATION_GROUPS)
	mapOrigCopy[a_spec_relation_groupsFrom] = a_spec_relation_groupsTo
	a_spec_relation_groupsFrom.CopyBasicFields(a_spec_relation_groupsTo)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relation_groupsFrom.RELATION_GROUP != nil {
		a_spec_relation_groupsTo.RELATION_GROUP = CopyBranchRELATION_GROUP(mapOrigCopy, a_spec_relation_groupsFrom.RELATION_GROUP)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_RELATION_REF(mapOrigCopy map[any]any, a_spec_relation_refFrom *A_SPEC_RELATION_REF) (a_spec_relation_refTo *A_SPEC_RELATION_REF) {

	// a_spec_relation_refFrom has already been copied
	if _a_spec_relation_refTo, ok := mapOrigCopy[a_spec_relation_refFrom]; ok {
		a_spec_relation_refTo = _a_spec_relation_refTo.(*A_SPEC_RELATION_REF)
		return
	}

	a_spec_relation_refTo = new(A_SPEC_RELATION_REF)
	mapOrigCopy[a_spec_relation_refFrom] = a_spec_relation_refTo
	a_spec_relation_refFrom.CopyBasicFields(a_spec_relation_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy map[any]any, a_spec_relation_type_refFrom *A_SPEC_RELATION_TYPE_REF) (a_spec_relation_type_refTo *A_SPEC_RELATION_TYPE_REF) {

	// a_spec_relation_type_refFrom has already been copied
	if _a_spec_relation_type_refTo, ok := mapOrigCopy[a_spec_relation_type_refFrom]; ok {
		a_spec_relation_type_refTo = _a_spec_relation_type_refTo.(*A_SPEC_RELATION_TYPE_REF)
		return
	}

	a_spec_relation_type_refTo = new(A_SPEC_RELATION_TYPE_REF)
	mapOrigCopy[a_spec_relation_type_refFrom] = a_spec_relation_type_refTo
	a_spec_relation_type_refFrom.CopyBasicFields(a_spec_relation_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_TYPES(mapOrigCopy map[any]any, a_spec_typesFrom *A_SPEC_TYPES) (a_spec_typesTo *A_SPEC_TYPES) {

	// a_spec_typesFrom has already been copied
	if _a_spec_typesTo, ok := mapOrigCopy[a_spec_typesFrom]; ok {
		a_spec_typesTo = _a_spec_typesTo.(*A_SPEC_TYPES)
		return
	}

	a_spec_typesTo = new(A_SPEC_TYPES)
	mapOrigCopy[a_spec_typesFrom] = a_spec_typesTo
	a_spec_typesFrom.CopyBasicFields(a_spec_typesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_typesFrom.RELATION_GROUP_TYPE != nil {
		a_spec_typesTo.RELATION_GROUP_TYPE = CopyBranchRELATION_GROUP_TYPE(mapOrigCopy, a_spec_typesFrom.RELATION_GROUP_TYPE)
	}
	if a_spec_typesFrom.SPEC_OBJECT_TYPE != nil {
		a_spec_typesTo.SPEC_OBJECT_TYPE = CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, a_spec_typesFrom.SPEC_OBJECT_TYPE)
	}
	if a_spec_typesFrom.SPEC_RELATION_TYPE != nil {
		a_spec_typesTo.SPEC_RELATION_TYPE = CopyBranchSPEC_RELATION_TYPE(mapOrigCopy, a_spec_typesFrom.SPEC_RELATION_TYPE)
	}
	if a_spec_typesFrom.SPECIFICATION_TYPE != nil {
		a_spec_typesTo.SPECIFICATION_TYPE = CopyBranchSPECIFICATION_TYPE(mapOrigCopy, a_spec_typesFrom.SPECIFICATION_TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TARGET_1(mapOrigCopy map[any]any, a_target_1From *A_TARGET_1) (a_target_1To *A_TARGET_1) {

	// a_target_1From has already been copied
	if _a_target_1To, ok := mapOrigCopy[a_target_1From]; ok {
		a_target_1To = _a_target_1To.(*A_TARGET_1)
		return
	}

	a_target_1To = new(A_TARGET_1)
	mapOrigCopy[a_target_1From] = a_target_1To
	a_target_1From.CopyBasicFields(a_target_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_THE_HEADER(mapOrigCopy map[any]any, a_the_headerFrom *A_THE_HEADER) (a_the_headerTo *A_THE_HEADER) {

	// a_the_headerFrom has already been copied
	if _a_the_headerTo, ok := mapOrigCopy[a_the_headerFrom]; ok {
		a_the_headerTo = _a_the_headerTo.(*A_THE_HEADER)
		return
	}

	a_the_headerTo = new(A_THE_HEADER)
	mapOrigCopy[a_the_headerFrom] = a_the_headerTo
	a_the_headerFrom.CopyBasicFields(a_the_headerTo)

	//insertion point for the staging of instances referenced by pointers
	if a_the_headerFrom.REQ_IF_HEADER != nil {
		a_the_headerTo.REQ_IF_HEADER = CopyBranchREQ_IF_HEADER(mapOrigCopy, a_the_headerFrom.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TOOL_EXTENSIONS(mapOrigCopy map[any]any, a_tool_extensionsFrom *A_TOOL_EXTENSIONS) (a_tool_extensionsTo *A_TOOL_EXTENSIONS) {

	// a_tool_extensionsFrom has already been copied
	if _a_tool_extensionsTo, ok := mapOrigCopy[a_tool_extensionsFrom]; ok {
		a_tool_extensionsTo = _a_tool_extensionsTo.(*A_TOOL_EXTENSIONS)
		return
	}

	a_tool_extensionsTo = new(A_TOOL_EXTENSIONS)
	mapOrigCopy[a_tool_extensionsFrom] = a_tool_extensionsTo
	a_tool_extensionsFrom.CopyBasicFields(a_tool_extensionsTo)

	//insertion point for the staging of instances referenced by pointers
	if a_tool_extensionsFrom.REQ_IF_TOOL_EXTENSION != nil {
		a_tool_extensionsTo.REQ_IF_TOOL_EXTENSION = CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, a_tool_extensionsFrom.REQ_IF_TOOL_EXTENSION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, datatype_definition_booleanFrom *DATATYPE_DEFINITION_BOOLEAN) (datatype_definition_booleanTo *DATATYPE_DEFINITION_BOOLEAN) {

	// datatype_definition_booleanFrom has already been copied
	if _datatype_definition_booleanTo, ok := mapOrigCopy[datatype_definition_booleanFrom]; ok {
		datatype_definition_booleanTo = _datatype_definition_booleanTo.(*DATATYPE_DEFINITION_BOOLEAN)
		return
	}

	datatype_definition_booleanTo = new(DATATYPE_DEFINITION_BOOLEAN)
	mapOrigCopy[datatype_definition_booleanFrom] = datatype_definition_booleanTo
	datatype_definition_booleanFrom.CopyBasicFields(datatype_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_booleanFrom.ALTERNATIVE_ID != nil {
		datatype_definition_booleanTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_booleanFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy map[any]any, datatype_definition_dateFrom *DATATYPE_DEFINITION_DATE) (datatype_definition_dateTo *DATATYPE_DEFINITION_DATE) {

	// datatype_definition_dateFrom has already been copied
	if _datatype_definition_dateTo, ok := mapOrigCopy[datatype_definition_dateFrom]; ok {
		datatype_definition_dateTo = _datatype_definition_dateTo.(*DATATYPE_DEFINITION_DATE)
		return
	}

	datatype_definition_dateTo = new(DATATYPE_DEFINITION_DATE)
	mapOrigCopy[datatype_definition_dateFrom] = datatype_definition_dateTo
	datatype_definition_dateFrom.CopyBasicFields(datatype_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_dateFrom.ALTERNATIVE_ID != nil {
		datatype_definition_dateTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_dateFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, datatype_definition_enumerationFrom *DATATYPE_DEFINITION_ENUMERATION) (datatype_definition_enumerationTo *DATATYPE_DEFINITION_ENUMERATION) {

	// datatype_definition_enumerationFrom has already been copied
	if _datatype_definition_enumerationTo, ok := mapOrigCopy[datatype_definition_enumerationFrom]; ok {
		datatype_definition_enumerationTo = _datatype_definition_enumerationTo.(*DATATYPE_DEFINITION_ENUMERATION)
		return
	}

	datatype_definition_enumerationTo = new(DATATYPE_DEFINITION_ENUMERATION)
	mapOrigCopy[datatype_definition_enumerationFrom] = datatype_definition_enumerationTo
	datatype_definition_enumerationFrom.CopyBasicFields(datatype_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumerationFrom.ALTERNATIVE_ID != nil {
		datatype_definition_enumerationTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_enumerationFrom.ALTERNATIVE_ID)
	}
	if datatype_definition_enumerationFrom.SPECIFIED_VALUES != nil {
		datatype_definition_enumerationTo.SPECIFIED_VALUES = CopyBranchA_SPECIFIED_VALUES(mapOrigCopy, datatype_definition_enumerationFrom.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy map[any]any, datatype_definition_integerFrom *DATATYPE_DEFINITION_INTEGER) (datatype_definition_integerTo *DATATYPE_DEFINITION_INTEGER) {

	// datatype_definition_integerFrom has already been copied
	if _datatype_definition_integerTo, ok := mapOrigCopy[datatype_definition_integerFrom]; ok {
		datatype_definition_integerTo = _datatype_definition_integerTo.(*DATATYPE_DEFINITION_INTEGER)
		return
	}

	datatype_definition_integerTo = new(DATATYPE_DEFINITION_INTEGER)
	mapOrigCopy[datatype_definition_integerFrom] = datatype_definition_integerTo
	datatype_definition_integerFrom.CopyBasicFields(datatype_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integerFrom.ALTERNATIVE_ID != nil {
		datatype_definition_integerTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_integerFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy map[any]any, datatype_definition_realFrom *DATATYPE_DEFINITION_REAL) (datatype_definition_realTo *DATATYPE_DEFINITION_REAL) {

	// datatype_definition_realFrom has already been copied
	if _datatype_definition_realTo, ok := mapOrigCopy[datatype_definition_realFrom]; ok {
		datatype_definition_realTo = _datatype_definition_realTo.(*DATATYPE_DEFINITION_REAL)
		return
	}

	datatype_definition_realTo = new(DATATYPE_DEFINITION_REAL)
	mapOrigCopy[datatype_definition_realFrom] = datatype_definition_realTo
	datatype_definition_realFrom.CopyBasicFields(datatype_definition_realTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_realFrom.ALTERNATIVE_ID != nil {
		datatype_definition_realTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_realFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy map[any]any, datatype_definition_stringFrom *DATATYPE_DEFINITION_STRING) (datatype_definition_stringTo *DATATYPE_DEFINITION_STRING) {

	// datatype_definition_stringFrom has already been copied
	if _datatype_definition_stringTo, ok := mapOrigCopy[datatype_definition_stringFrom]; ok {
		datatype_definition_stringTo = _datatype_definition_stringTo.(*DATATYPE_DEFINITION_STRING)
		return
	}

	datatype_definition_stringTo = new(DATATYPE_DEFINITION_STRING)
	mapOrigCopy[datatype_definition_stringFrom] = datatype_definition_stringTo
	datatype_definition_stringFrom.CopyBasicFields(datatype_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_stringFrom.ALTERNATIVE_ID != nil {
		datatype_definition_stringTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_stringFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy map[any]any, datatype_definition_xhtmlFrom *DATATYPE_DEFINITION_XHTML) (datatype_definition_xhtmlTo *DATATYPE_DEFINITION_XHTML) {

	// datatype_definition_xhtmlFrom has already been copied
	if _datatype_definition_xhtmlTo, ok := mapOrigCopy[datatype_definition_xhtmlFrom]; ok {
		datatype_definition_xhtmlTo = _datatype_definition_xhtmlTo.(*DATATYPE_DEFINITION_XHTML)
		return
	}

	datatype_definition_xhtmlTo = new(DATATYPE_DEFINITION_XHTML)
	mapOrigCopy[datatype_definition_xhtmlFrom] = datatype_definition_xhtmlTo
	datatype_definition_xhtmlFrom.CopyBasicFields(datatype_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtmlFrom.ALTERNATIVE_ID != nil {
		datatype_definition_xhtmlTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_xhtmlFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEMBEDDED_VALUE(mapOrigCopy map[any]any, embedded_valueFrom *EMBEDDED_VALUE) (embedded_valueTo *EMBEDDED_VALUE) {

	// embedded_valueFrom has already been copied
	if _embedded_valueTo, ok := mapOrigCopy[embedded_valueFrom]; ok {
		embedded_valueTo = _embedded_valueTo.(*EMBEDDED_VALUE)
		return
	}

	embedded_valueTo = new(EMBEDDED_VALUE)
	mapOrigCopy[embedded_valueFrom] = embedded_valueTo
	embedded_valueFrom.CopyBasicFields(embedded_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchENUM_VALUE(mapOrigCopy map[any]any, enum_valueFrom *ENUM_VALUE) (enum_valueTo *ENUM_VALUE) {

	// enum_valueFrom has already been copied
	if _enum_valueTo, ok := mapOrigCopy[enum_valueFrom]; ok {
		enum_valueTo = _enum_valueTo.(*ENUM_VALUE)
		return
	}

	enum_valueTo = new(ENUM_VALUE)
	mapOrigCopy[enum_valueFrom] = enum_valueTo
	enum_valueFrom.CopyBasicFields(enum_valueTo)

	//insertion point for the staging of instances referenced by pointers
	if enum_valueFrom.ALTERNATIVE_ID != nil {
		enum_valueTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, enum_valueFrom.ALTERNATIVE_ID)
	}
	if enum_valueFrom.PROPERTIES != nil {
		enum_valueTo.PROPERTIES = CopyBranchA_PROPERTIES(mapOrigCopy, enum_valueFrom.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRELATION_GROUP(mapOrigCopy map[any]any, relation_groupFrom *RELATION_GROUP) (relation_groupTo *RELATION_GROUP) {

	// relation_groupFrom has already been copied
	if _relation_groupTo, ok := mapOrigCopy[relation_groupFrom]; ok {
		relation_groupTo = _relation_groupTo.(*RELATION_GROUP)
		return
	}

	relation_groupTo = new(RELATION_GROUP)
	mapOrigCopy[relation_groupFrom] = relation_groupTo
	relation_groupFrom.CopyBasicFields(relation_groupTo)

	//insertion point for the staging of instances referenced by pointers
	if relation_groupFrom.ALTERNATIVE_ID != nil {
		relation_groupTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, relation_groupFrom.ALTERNATIVE_ID)
	}
	if relation_groupFrom.SOURCE_SPECIFICATION != nil {
		relation_groupTo.SOURCE_SPECIFICATION = CopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, relation_groupFrom.SOURCE_SPECIFICATION)
	}
	if relation_groupFrom.SPEC_RELATIONS != nil {
		relation_groupTo.SPEC_RELATIONS = CopyBranchA_SPEC_RELATION_REF(mapOrigCopy, relation_groupFrom.SPEC_RELATIONS)
	}
	if relation_groupFrom.TARGET_SPECIFICATION != nil {
		relation_groupTo.TARGET_SPECIFICATION = CopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, relation_groupFrom.TARGET_SPECIFICATION)
	}
	if relation_groupFrom.TYPE != nil {
		relation_groupTo.TYPE = CopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy, relation_groupFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRELATION_GROUP_TYPE(mapOrigCopy map[any]any, relation_group_typeFrom *RELATION_GROUP_TYPE) (relation_group_typeTo *RELATION_GROUP_TYPE) {

	// relation_group_typeFrom has already been copied
	if _relation_group_typeTo, ok := mapOrigCopy[relation_group_typeFrom]; ok {
		relation_group_typeTo = _relation_group_typeTo.(*RELATION_GROUP_TYPE)
		return
	}

	relation_group_typeTo = new(RELATION_GROUP_TYPE)
	mapOrigCopy[relation_group_typeFrom] = relation_group_typeTo
	relation_group_typeFrom.CopyBasicFields(relation_group_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_typeFrom.ALTERNATIVE_ID != nil {
		relation_group_typeTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, relation_group_typeFrom.ALTERNATIVE_ID)
	}
	if relation_group_typeFrom.SPEC_ATTRIBUTES != nil {
		relation_group_typeTo.SPEC_ATTRIBUTES = CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, relation_group_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQ_IF(mapOrigCopy map[any]any, req_ifFrom *REQ_IF) (req_ifTo *REQ_IF) {

	// req_ifFrom has already been copied
	if _req_ifTo, ok := mapOrigCopy[req_ifFrom]; ok {
		req_ifTo = _req_ifTo.(*REQ_IF)
		return
	}

	req_ifTo = new(REQ_IF)
	mapOrigCopy[req_ifFrom] = req_ifTo
	req_ifFrom.CopyBasicFields(req_ifTo)

	//insertion point for the staging of instances referenced by pointers
	if req_ifFrom.THE_HEADER != nil {
		req_ifTo.THE_HEADER = CopyBranchA_THE_HEADER(mapOrigCopy, req_ifFrom.THE_HEADER)
	}
	if req_ifFrom.CORE_CONTENT != nil {
		req_ifTo.CORE_CONTENT = CopyBranchA_CORE_CONTENT(mapOrigCopy, req_ifFrom.CORE_CONTENT)
	}
	if req_ifFrom.TOOL_EXTENSIONS != nil {
		req_ifTo.TOOL_EXTENSIONS = CopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, req_ifFrom.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQ_IF_CONTENT(mapOrigCopy map[any]any, req_if_contentFrom *REQ_IF_CONTENT) (req_if_contentTo *REQ_IF_CONTENT) {

	// req_if_contentFrom has already been copied
	if _req_if_contentTo, ok := mapOrigCopy[req_if_contentFrom]; ok {
		req_if_contentTo = _req_if_contentTo.(*REQ_IF_CONTENT)
		return
	}

	req_if_contentTo = new(REQ_IF_CONTENT)
	mapOrigCopy[req_if_contentFrom] = req_if_contentTo
	req_if_contentFrom.CopyBasicFields(req_if_contentTo)

	//insertion point for the staging of instances referenced by pointers
	if req_if_contentFrom.DATATYPES != nil {
		req_if_contentTo.DATATYPES = CopyBranchA_DATATYPES(mapOrigCopy, req_if_contentFrom.DATATYPES)
	}
	if req_if_contentFrom.SPEC_TYPES != nil {
		req_if_contentTo.SPEC_TYPES = CopyBranchA_SPEC_TYPES(mapOrigCopy, req_if_contentFrom.SPEC_TYPES)
	}
	if req_if_contentFrom.SPEC_OBJECTS != nil {
		req_if_contentTo.SPEC_OBJECTS = CopyBranchA_SPEC_OBJECTS(mapOrigCopy, req_if_contentFrom.SPEC_OBJECTS)
	}
	if req_if_contentFrom.SPEC_RELATIONS != nil {
		req_if_contentTo.SPEC_RELATIONS = CopyBranchA_SPEC_RELATIONS(mapOrigCopy, req_if_contentFrom.SPEC_RELATIONS)
	}
	if req_if_contentFrom.SPECIFICATIONS != nil {
		req_if_contentTo.SPECIFICATIONS = CopyBranchA_SPECIFICATIONS(mapOrigCopy, req_if_contentFrom.SPECIFICATIONS)
	}
	if req_if_contentFrom.SPEC_RELATION_GROUPS != nil {
		req_if_contentTo.SPEC_RELATION_GROUPS = CopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, req_if_contentFrom.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQ_IF_HEADER(mapOrigCopy map[any]any, req_if_headerFrom *REQ_IF_HEADER) (req_if_headerTo *REQ_IF_HEADER) {

	// req_if_headerFrom has already been copied
	if _req_if_headerTo, ok := mapOrigCopy[req_if_headerFrom]; ok {
		req_if_headerTo = _req_if_headerTo.(*REQ_IF_HEADER)
		return
	}

	req_if_headerTo = new(REQ_IF_HEADER)
	mapOrigCopy[req_if_headerFrom] = req_if_headerTo
	req_if_headerFrom.CopyBasicFields(req_if_headerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy map[any]any, req_if_tool_extensionFrom *REQ_IF_TOOL_EXTENSION) (req_if_tool_extensionTo *REQ_IF_TOOL_EXTENSION) {

	// req_if_tool_extensionFrom has already been copied
	if _req_if_tool_extensionTo, ok := mapOrigCopy[req_if_tool_extensionFrom]; ok {
		req_if_tool_extensionTo = _req_if_tool_extensionTo.(*REQ_IF_TOOL_EXTENSION)
		return
	}

	req_if_tool_extensionTo = new(REQ_IF_TOOL_EXTENSION)
	mapOrigCopy[req_if_tool_extensionFrom] = req_if_tool_extensionTo
	req_if_tool_extensionFrom.CopyBasicFields(req_if_tool_extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFICATION(mapOrigCopy map[any]any, specificationFrom *SPECIFICATION) (specificationTo *SPECIFICATION) {

	// specificationFrom has already been copied
	if _specificationTo, ok := mapOrigCopy[specificationFrom]; ok {
		specificationTo = _specificationTo.(*SPECIFICATION)
		return
	}

	specificationTo = new(SPECIFICATION)
	mapOrigCopy[specificationFrom] = specificationTo
	specificationFrom.CopyBasicFields(specificationTo)

	//insertion point for the staging of instances referenced by pointers
	if specificationFrom.ALTERNATIVE_ID != nil {
		specificationTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, specificationFrom.ALTERNATIVE_ID)
	}
	if specificationFrom.CHILDREN != nil {
		specificationTo.CHILDREN = CopyBranchA_CHILDREN(mapOrigCopy, specificationFrom.CHILDREN)
	}
	if specificationFrom.VALUES != nil {
		specificationTo.VALUES = CopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, specificationFrom.VALUES)
	}
	if specificationFrom.TYPE != nil {
		specificationTo.TYPE = CopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy, specificationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFICATION_TYPE(mapOrigCopy map[any]any, specification_typeFrom *SPECIFICATION_TYPE) (specification_typeTo *SPECIFICATION_TYPE) {

	// specification_typeFrom has already been copied
	if _specification_typeTo, ok := mapOrigCopy[specification_typeFrom]; ok {
		specification_typeTo = _specification_typeTo.(*SPECIFICATION_TYPE)
		return
	}

	specification_typeTo = new(SPECIFICATION_TYPE)
	mapOrigCopy[specification_typeFrom] = specification_typeTo
	specification_typeFrom.CopyBasicFields(specification_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if specification_typeFrom.ALTERNATIVE_ID != nil {
		specification_typeTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, specification_typeFrom.ALTERNATIVE_ID)
	}
	if specification_typeFrom.SPEC_ATTRIBUTES != nil {
		specification_typeTo.SPEC_ATTRIBUTES = CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, specification_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPEC_HIERARCHY(mapOrigCopy map[any]any, spec_hierarchyFrom *SPEC_HIERARCHY) (spec_hierarchyTo *SPEC_HIERARCHY) {

	// spec_hierarchyFrom has already been copied
	if _spec_hierarchyTo, ok := mapOrigCopy[spec_hierarchyFrom]; ok {
		spec_hierarchyTo = _spec_hierarchyTo.(*SPEC_HIERARCHY)
		return
	}

	spec_hierarchyTo = new(SPEC_HIERARCHY)
	mapOrigCopy[spec_hierarchyFrom] = spec_hierarchyTo
	spec_hierarchyFrom.CopyBasicFields(spec_hierarchyTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchyFrom.ALTERNATIVE_ID != nil {
		spec_hierarchyTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_hierarchyFrom.ALTERNATIVE_ID)
	}
	if spec_hierarchyFrom.CHILDREN != nil {
		spec_hierarchyTo.CHILDREN = CopyBranchA_CHILDREN(mapOrigCopy, spec_hierarchyFrom.CHILDREN)
	}
	if spec_hierarchyFrom.EDITABLE_ATTS != nil {
		spec_hierarchyTo.EDITABLE_ATTS = CopyBranchA_EDITABLE_ATTS(mapOrigCopy, spec_hierarchyFrom.EDITABLE_ATTS)
	}
	if spec_hierarchyFrom.OBJECT != nil {
		spec_hierarchyTo.OBJECT = CopyBranchA_OBJECT(mapOrigCopy, spec_hierarchyFrom.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPEC_OBJECT(mapOrigCopy map[any]any, spec_objectFrom *SPEC_OBJECT) (spec_objectTo *SPEC_OBJECT) {

	// spec_objectFrom has already been copied
	if _spec_objectTo, ok := mapOrigCopy[spec_objectFrom]; ok {
		spec_objectTo = _spec_objectTo.(*SPEC_OBJECT)
		return
	}

	spec_objectTo = new(SPEC_OBJECT)
	mapOrigCopy[spec_objectFrom] = spec_objectTo
	spec_objectFrom.CopyBasicFields(spec_objectTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_objectFrom.ALTERNATIVE_ID != nil {
		spec_objectTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_objectFrom.ALTERNATIVE_ID)
	}
	if spec_objectFrom.VALUES != nil {
		spec_objectTo.VALUES = CopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, spec_objectFrom.VALUES)
	}
	if spec_objectFrom.TYPE != nil {
		spec_objectTo.TYPE = CopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy, spec_objectFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy map[any]any, spec_object_typeFrom *SPEC_OBJECT_TYPE) (spec_object_typeTo *SPEC_OBJECT_TYPE) {

	// spec_object_typeFrom has already been copied
	if _spec_object_typeTo, ok := mapOrigCopy[spec_object_typeFrom]; ok {
		spec_object_typeTo = _spec_object_typeTo.(*SPEC_OBJECT_TYPE)
		return
	}

	spec_object_typeTo = new(SPEC_OBJECT_TYPE)
	mapOrigCopy[spec_object_typeFrom] = spec_object_typeTo
	spec_object_typeFrom.CopyBasicFields(spec_object_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_typeFrom.ALTERNATIVE_ID != nil {
		spec_object_typeTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_object_typeFrom.ALTERNATIVE_ID)
	}
	if spec_object_typeFrom.SPEC_ATTRIBUTES != nil {
		spec_object_typeTo.SPEC_ATTRIBUTES = CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, spec_object_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPEC_RELATION(mapOrigCopy map[any]any, spec_relationFrom *SPEC_RELATION) (spec_relationTo *SPEC_RELATION) {

	// spec_relationFrom has already been copied
	if _spec_relationTo, ok := mapOrigCopy[spec_relationFrom]; ok {
		spec_relationTo = _spec_relationTo.(*SPEC_RELATION)
		return
	}

	spec_relationTo = new(SPEC_RELATION)
	mapOrigCopy[spec_relationFrom] = spec_relationTo
	spec_relationFrom.CopyBasicFields(spec_relationTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_relationFrom.ALTERNATIVE_ID != nil {
		spec_relationTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_relationFrom.ALTERNATIVE_ID)
	}
	if spec_relationFrom.VALUES != nil {
		spec_relationTo.VALUES = CopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, spec_relationFrom.VALUES)
	}
	if spec_relationFrom.SOURCE != nil {
		spec_relationTo.SOURCE = CopyBranchA_TARGET_1(mapOrigCopy, spec_relationFrom.SOURCE)
	}
	if spec_relationFrom.TARGET != nil {
		spec_relationTo.TARGET = CopyBranchA_TARGET_1(mapOrigCopy, spec_relationFrom.TARGET)
	}
	if spec_relationFrom.TYPE != nil {
		spec_relationTo.TYPE = CopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy, spec_relationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPEC_RELATION_TYPE(mapOrigCopy map[any]any, spec_relation_typeFrom *SPEC_RELATION_TYPE) (spec_relation_typeTo *SPEC_RELATION_TYPE) {

	// spec_relation_typeFrom has already been copied
	if _spec_relation_typeTo, ok := mapOrigCopy[spec_relation_typeFrom]; ok {
		spec_relation_typeTo = _spec_relation_typeTo.(*SPEC_RELATION_TYPE)
		return
	}

	spec_relation_typeTo = new(SPEC_RELATION_TYPE)
	mapOrigCopy[spec_relation_typeFrom] = spec_relation_typeTo
	spec_relation_typeFrom.CopyBasicFields(spec_relation_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_typeFrom.ALTERNATIVE_ID != nil {
		spec_relation_typeTo.ALTERNATIVE_ID = CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_relation_typeFrom.ALTERNATIVE_ID)
	}
	if spec_relation_typeFrom.SPEC_ATTRIBUTES != nil {
		spec_relation_typeTo.SPEC_ATTRIBUTES = CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, spec_relation_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXHTML_CONTENT(mapOrigCopy map[any]any, xhtml_contentFrom *XHTML_CONTENT) (xhtml_contentTo *XHTML_CONTENT) {

	// xhtml_contentFrom has already been copied
	if _xhtml_contentTo, ok := mapOrigCopy[xhtml_contentFrom]; ok {
		xhtml_contentTo = _xhtml_contentTo.(*XHTML_CONTENT)
		return
	}

	xhtml_contentTo = new(XHTML_CONTENT)
	mapOrigCopy[xhtml_contentFrom] = xhtml_contentTo
	xhtml_contentFrom.CopyBasicFields(xhtml_contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *ALTERNATIVE_ID:
		stage.UnstageBranchALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		stage.UnstageBranchATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.UnstageBranchATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		stage.UnstageBranchATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		stage.UnstageBranchATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.UnstageBranchATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.UnstageBranchATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		stage.UnstageBranchATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.UnstageBranchATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		stage.UnstageBranchATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		stage.UnstageBranchATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		stage.UnstageBranchATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		stage.UnstageBranchATTRIBUTE_VALUE_XHTML(target)

	case *A_ALTERNATIVE_ID:
		stage.UnstageBranchA_ALTERNATIVE_ID(target)

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(target)

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(target)

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(target)

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(target)

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(target)

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(target)

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.UnstageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(target)

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_BOOLEAN(target)

	case *A_ATTRIBUTE_VALUE_DATE:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_DATE(target)

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_ENUMERATION(target)

	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_INTEGER(target)

	case *A_ATTRIBUTE_VALUE_REAL:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_REAL(target)

	case *A_ATTRIBUTE_VALUE_STRING:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_STRING(target)

	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_XHTML(target)

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.UnstageBranchA_ATTRIBUTE_VALUE_XHTML_1(target)

	case *A_CHILDREN:
		stage.UnstageBranchA_CHILDREN(target)

	case *A_CORE_CONTENT:
		stage.UnstageBranchA_CORE_CONTENT(target)

	case *A_DATATYPES:
		stage.UnstageBranchA_DATATYPES(target)

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(target)

	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_DATE_REF(target)

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(target)

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_INTEGER_REF(target)

	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_REAL_REF(target)

	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_STRING_REF(target)

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.UnstageBranchA_DATATYPE_DEFINITION_XHTML_REF(target)

	case *A_EDITABLE_ATTS:
		stage.UnstageBranchA_EDITABLE_ATTS(target)

	case *A_ENUM_VALUE_REF:
		stage.UnstageBranchA_ENUM_VALUE_REF(target)

	case *A_OBJECT:
		stage.UnstageBranchA_OBJECT(target)

	case *A_PROPERTIES:
		stage.UnstageBranchA_PROPERTIES(target)

	case *A_RELATION_GROUP_TYPE_REF:
		stage.UnstageBranchA_RELATION_GROUP_TYPE_REF(target)

	case *A_SOURCE_SPECIFICATION_1:
		stage.UnstageBranchA_SOURCE_SPECIFICATION_1(target)

	case *A_SPECIFICATIONS:
		stage.UnstageBranchA_SPECIFICATIONS(target)

	case *A_SPECIFICATION_TYPE_REF:
		stage.UnstageBranchA_SPECIFICATION_TYPE_REF(target)

	case *A_SPECIFIED_VALUES:
		stage.UnstageBranchA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		stage.UnstageBranchA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		stage.UnstageBranchA_SPEC_OBJECTS(target)

	case *A_SPEC_OBJECT_TYPE_REF:
		stage.UnstageBranchA_SPEC_OBJECT_TYPE_REF(target)

	case *A_SPEC_RELATIONS:
		stage.UnstageBranchA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATION_GROUPS:
		stage.UnstageBranchA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_RELATION_REF:
		stage.UnstageBranchA_SPEC_RELATION_REF(target)

	case *A_SPEC_RELATION_TYPE_REF:
		stage.UnstageBranchA_SPEC_RELATION_TYPE_REF(target)

	case *A_SPEC_TYPES:
		stage.UnstageBranchA_SPEC_TYPES(target)

	case *A_TARGET_1:
		stage.UnstageBranchA_TARGET_1(target)

	case *A_THE_HEADER:
		stage.UnstageBranchA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		stage.UnstageBranchA_TOOL_EXTENSIONS(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.UnstageBranchDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		stage.UnstageBranchDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.UnstageBranchDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		stage.UnstageBranchDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		stage.UnstageBranchDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		stage.UnstageBranchDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		stage.UnstageBranchDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		stage.UnstageBranchEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		stage.UnstageBranchENUM_VALUE(target)

	case *RELATION_GROUP:
		stage.UnstageBranchRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		stage.UnstageBranchRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		stage.UnstageBranchREQ_IF(target)

	case *REQ_IF_CONTENT:
		stage.UnstageBranchREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		stage.UnstageBranchREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		stage.UnstageBranchREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		stage.UnstageBranchSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		stage.UnstageBranchSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		stage.UnstageBranchSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		stage.UnstageBranchSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		stage.UnstageBranchSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		stage.UnstageBranchSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		stage.UnstageBranchSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		stage.UnstageBranchXHTML_CONTENT(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if !IsStaged(stage, alternative_id) {
		return
	}

	alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_boolean.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_boolean.ALTERNATIVE_ID)
	}
	if attribute_definition_boolean.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_boolean.DEFAULT_VALUE)
	}
	if attribute_definition_boolean.TYPE != nil {
		UnstageBranch(stage, attribute_definition_boolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_date.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_date.ALTERNATIVE_ID)
	}
	if attribute_definition_date.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_date.DEFAULT_VALUE)
	}
	if attribute_definition_date.TYPE != nil {
		UnstageBranch(stage, attribute_definition_date.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_enumeration.ALTERNATIVE_ID)
	}
	if attribute_definition_enumeration.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_enumeration.DEFAULT_VALUE)
	}
	if attribute_definition_enumeration.TYPE != nil {
		UnstageBranch(stage, attribute_definition_enumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integer.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_integer.ALTERNATIVE_ID)
	}
	if attribute_definition_integer.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_integer.DEFAULT_VALUE)
	}
	if attribute_definition_integer.TYPE != nil {
		UnstageBranch(stage, attribute_definition_integer.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_real.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_real.ALTERNATIVE_ID)
	}
	if attribute_definition_real.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_real.DEFAULT_VALUE)
	}
	if attribute_definition_real.TYPE != nil {
		UnstageBranch(stage, attribute_definition_real.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_string.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_string.ALTERNATIVE_ID)
	}
	if attribute_definition_string.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_string.DEFAULT_VALUE)
	}
	if attribute_definition_string.TYPE != nil {
		UnstageBranch(stage, attribute_definition_string.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, attribute_definition_xhtml.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtml.DEFAULT_VALUE != nil {
		UnstageBranch(stage, attribute_definition_xhtml.DEFAULT_VALUE)
	}
	if attribute_definition_xhtml.TYPE != nil {
		UnstageBranch(stage, attribute_definition_xhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_boolean.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_boolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_date.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_date.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumeration.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_enumeration.DEFINITION)
	}
	if attribute_value_enumeration.VALUES != nil {
		UnstageBranch(stage, attribute_value_enumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integer.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_integer.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_real.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_real.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_string.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_string.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtml.THE_VALUE != nil {
		UnstageBranch(stage, attribute_value_xhtml.THE_VALUE)
	}
	if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
		UnstageBranch(stage, attribute_value_xhtml.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtml.DEFINITION != nil {
		UnstageBranch(stage, attribute_value_xhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if !IsStaged(stage, a_alternative_id) {
		return
	}

	a_alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_id.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, a_alternative_id.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_boolean_ref) {
		return
	}

	a_attribute_definition_boolean_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_date_ref) {
		return
	}

	a_attribute_definition_date_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_enumeration_ref) {
		return
	}

	a_attribute_definition_enumeration_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_integer_ref) {
		return
	}

	a_attribute_definition_integer_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_real_ref) {
		return
	}

	a_attribute_definition_real_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_string_ref) {
		return
	}

	a_attribute_definition_string_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_definition_xhtml_ref) {
		return
	}

	a_attribute_definition_xhtml_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_boolean) {
		return
	}

	a_attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN != nil {
		UnstageBranch(stage, a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_date) {
		return
	}

	a_attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_date.ATTRIBUTE_VALUE_DATE != nil {
		UnstageBranch(stage, a_attribute_value_date.ATTRIBUTE_VALUE_DATE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_enumeration) {
		return
	}

	a_attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION != nil {
		UnstageBranch(stage, a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_integer) {
		return
	}

	a_attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER != nil {
		UnstageBranch(stage, a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_real) {
		return
	}

	a_attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_real.ATTRIBUTE_VALUE_REAL != nil {
		UnstageBranch(stage, a_attribute_value_real.ATTRIBUTE_VALUE_REAL)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_string) {
		return
	}

	a_attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_string.ATTRIBUTE_VALUE_STRING != nil {
		UnstageBranch(stage, a_attribute_value_string.ATTRIBUTE_VALUE_STRING)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_xhtml) {
		return
	}

	a_attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML != nil {
		UnstageBranch(stage, a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_attribute_value_xhtml_1) {
		return
	}

	a_attribute_value_xhtml_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING)
	}
	if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML != nil {
		UnstageBranch(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if !IsStaged(stage, a_children) {
		return
	}

	a_children.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_children.SPEC_HIERARCHY != nil {
		UnstageBranch(stage, a_children.SPEC_HIERARCHY)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, a_core_content) {
		return
	}

	a_core_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_core_content.REQ_IF_CONTENT != nil {
		UnstageBranch(stage, a_core_content.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatypes) {
		return
	}

	a_datatypes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_datatypes.DATATYPE_DEFINITION_BOOLEAN != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_BOOLEAN)
	}
	if a_datatypes.DATATYPE_DEFINITION_DATE != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_DATE)
	}
	if a_datatypes.DATATYPE_DEFINITION_ENUMERATION != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_ENUMERATION)
	}
	if a_datatypes.DATATYPE_DEFINITION_INTEGER != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_INTEGER)
	}
	if a_datatypes.DATATYPE_DEFINITION_REAL != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_REAL)
	}
	if a_datatypes.DATATYPE_DEFINITION_STRING != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_STRING)
	}
	if a_datatypes.DATATYPE_DEFINITION_XHTML != nil {
		UnstageBranch(stage, a_datatypes.DATATYPE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_boolean_ref) {
		return
	}

	a_datatype_definition_boolean_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_date_ref) {
		return
	}

	a_datatype_definition_date_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_enumeration_ref) {
		return
	}

	a_datatype_definition_enumeration_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_integer_ref) {
		return
	}

	a_datatype_definition_integer_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_real_ref) {
		return
	}

	a_datatype_definition_real_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_string_ref) {
		return
	}

	a_datatype_definition_string_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatype_definition_xhtml_ref) {
		return
	}

	a_datatype_definition_xhtml_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) {

	// check if instance is already staged
	if !IsStaged(stage, a_editable_atts) {
		return
	}

	a_editable_atts.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_enum_value_ref) {
		return
	}

	a_enum_value_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_OBJECT(a_object *A_OBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, a_object) {
		return
	}

	a_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_PROPERTIES(a_properties *A_PROPERTIES) {

	// check if instance is already staged
	if !IsStaged(stage, a_properties) {
		return
	}

	a_properties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_properties.EMBEDDED_VALUE != nil {
		UnstageBranch(stage, a_properties.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_relation_group_type_ref) {
		return
	}

	a_relation_group_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_source_specification_1) {
		return
	}

	a_source_specification_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) {

	// check if instance is already staged
	if !IsStaged(stage, a_specifications) {
		return
	}

	a_specifications.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_specifications.SPECIFICATION != nil {
		UnstageBranch(stage, a_specifications.SPECIFICATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_specification_type_ref) {
		return
	}

	a_specification_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if !IsStaged(stage, a_specified_values) {
		return
	}

	a_specified_values.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_specified_values.ENUM_VALUE != nil {
		UnstageBranch(stage, a_specified_values.ENUM_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_attributes) {
		return
	}

	a_spec_attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_DATE != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_DATE)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_REAL != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_REAL)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_STRING != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_STRING)
	}
	if a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML != nil {
		UnstageBranch(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_objects) {
		return
	}

	a_spec_objects.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_objects.SPEC_OBJECT != nil {
		UnstageBranch(stage, a_spec_objects.SPEC_OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_object_type_ref) {
		return
	}

	a_spec_object_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relations) {
		return
	}

	a_spec_relations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relations.SPEC_RELATION != nil {
		UnstageBranch(stage, a_spec_relations.SPEC_RELATION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_relation_groups.RELATION_GROUP != nil {
		UnstageBranch(stage, a_spec_relation_groups.RELATION_GROUP)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relation_ref) {
		return
	}

	a_spec_relation_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relation_type_ref) {
		return
	}

	a_spec_relation_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_types) {
		return
	}

	a_spec_types.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_spec_types.RELATION_GROUP_TYPE != nil {
		UnstageBranch(stage, a_spec_types.RELATION_GROUP_TYPE)
	}
	if a_spec_types.SPEC_OBJECT_TYPE != nil {
		UnstageBranch(stage, a_spec_types.SPEC_OBJECT_TYPE)
	}
	if a_spec_types.SPEC_RELATION_TYPE != nil {
		UnstageBranch(stage, a_spec_types.SPEC_RELATION_TYPE)
	}
	if a_spec_types.SPECIFICATION_TYPE != nil {
		UnstageBranch(stage, a_spec_types.SPECIFICATION_TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TARGET_1(a_target_1 *A_TARGET_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_target_1) {
		return
	}

	a_target_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if !IsStaged(stage, a_the_header) {
		return
	}

	a_the_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_the_header.REQ_IF_HEADER != nil {
		UnstageBranch(stage, a_the_header.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if !IsStaged(stage, a_tool_extensions) {
		return
	}

	a_tool_extensions.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_tool_extensions.REQ_IF_TOOL_EXTENSION != nil {
		UnstageBranch(stage, a_tool_extensions.REQ_IF_TOOL_EXTENSION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_boolean.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_boolean.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_date.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_date.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_enumeration.ALTERNATIVE_ID)
	}
	if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
		UnstageBranch(stage, datatype_definition_enumeration.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integer.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_integer.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_real.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_real.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_string.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_string.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, datatype_definition_xhtml.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if !IsStaged(stage, embedded_value) {
		return
	}

	embedded_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if !IsStaged(stage, enum_value) {
		return
	}

	enum_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enum_value.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, enum_value.ALTERNATIVE_ID)
	}
	if enum_value.PROPERTIES != nil {
		UnstageBranch(stage, enum_value.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group) {
		return
	}

	relation_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, relation_group.ALTERNATIVE_ID)
	}
	if relation_group.SOURCE_SPECIFICATION != nil {
		UnstageBranch(stage, relation_group.SOURCE_SPECIFICATION)
	}
	if relation_group.SPEC_RELATIONS != nil {
		UnstageBranch(stage, relation_group.SPEC_RELATIONS)
	}
	if relation_group.TARGET_SPECIFICATION != nil {
		UnstageBranch(stage, relation_group.TARGET_SPECIFICATION)
	}
	if relation_group.TYPE != nil {
		UnstageBranch(stage, relation_group.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_type.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, relation_group_type.ALTERNATIVE_ID)
	}
	if relation_group_type.SPEC_ATTRIBUTES != nil {
		UnstageBranch(stage, relation_group_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if !IsStaged(stage, req_if) {
		return
	}

	req_if.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER != nil {
		UnstageBranch(stage, req_if.THE_HEADER)
	}
	if req_if.CORE_CONTENT != nil {
		UnstageBranch(stage, req_if.CORE_CONTENT)
	}
	if req_if.TOOL_EXTENSIONS != nil {
		UnstageBranch(stage, req_if.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if_content.DATATYPES != nil {
		UnstageBranch(stage, req_if_content.DATATYPES)
	}
	if req_if_content.SPEC_TYPES != nil {
		UnstageBranch(stage, req_if_content.SPEC_TYPES)
	}
	if req_if_content.SPEC_OBJECTS != nil {
		UnstageBranch(stage, req_if_content.SPEC_OBJECTS)
	}
	if req_if_content.SPEC_RELATIONS != nil {
		UnstageBranch(stage, req_if_content.SPEC_RELATIONS)
	}
	if req_if_content.SPECIFICATIONS != nil {
		UnstageBranch(stage, req_if_content.SPECIFICATIONS)
	}
	if req_if_content.SPEC_RELATION_GROUPS != nil {
		UnstageBranch(stage, req_if_content.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_header) {
		return
	}

	req_if_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, specification) {
		return
	}

	specification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, specification.ALTERNATIVE_ID)
	}
	if specification.CHILDREN != nil {
		UnstageBranch(stage, specification.CHILDREN)
	}
	if specification.VALUES != nil {
		UnstageBranch(stage, specification.VALUES)
	}
	if specification.TYPE != nil {
		UnstageBranch(stage, specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specification_type) {
		return
	}

	specification_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification_type.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, specification_type.ALTERNATIVE_ID)
	}
	if specification_type.SPEC_ATTRIBUTES != nil {
		UnstageBranch(stage, specification_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if !IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchy.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, spec_hierarchy.ALTERNATIVE_ID)
	}
	if spec_hierarchy.CHILDREN != nil {
		UnstageBranch(stage, spec_hierarchy.CHILDREN)
	}
	if spec_hierarchy.EDITABLE_ATTS != nil {
		UnstageBranch(stage, spec_hierarchy.EDITABLE_ATTS)
	}
	if spec_hierarchy.OBJECT != nil {
		UnstageBranch(stage, spec_hierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object) {
		return
	}

	spec_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, spec_object.ALTERNATIVE_ID)
	}
	if spec_object.VALUES != nil {
		UnstageBranch(stage, spec_object.VALUES)
	}
	if spec_object.TYPE != nil {
		UnstageBranch(stage, spec_object.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_type.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, spec_object_type.ALTERNATIVE_ID)
	}
	if spec_object_type.SPEC_ATTRIBUTES != nil {
		UnstageBranch(stage, spec_object_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, spec_relation.ALTERNATIVE_ID)
	}
	if spec_relation.VALUES != nil {
		UnstageBranch(stage, spec_relation.VALUES)
	}
	if spec_relation.SOURCE != nil {
		UnstageBranch(stage, spec_relation.SOURCE)
	}
	if spec_relation.TARGET != nil {
		UnstageBranch(stage, spec_relation.TARGET)
	}
	if spec_relation.TYPE != nil {
		UnstageBranch(stage, spec_relation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_type.ALTERNATIVE_ID != nil {
		UnstageBranch(stage, spec_relation_type.ALTERNATIVE_ID)
	}
	if spec_relation_type.SPEC_ATTRIBUTES != nil {
		UnstageBranch(stage, spec_relation_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_content) {
		return
	}

	xhtml_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
