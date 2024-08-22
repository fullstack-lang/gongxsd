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

	case *A_CHILDREN:
		ok = stage.IsStagedA_CHILDREN(target)

	case *A_CORE_CONTENT:
		ok = stage.IsStagedA_CORE_CONTENT(target)

	case *A_DATATYPES:
		ok = stage.IsStagedA_DATATYPES(target)

	case *A_DEFAULT_VALUE:
		ok = stage.IsStagedA_DEFAULT_VALUE(target)

	case *A_DEFAULT_VALUE_1:
		ok = stage.IsStagedA_DEFAULT_VALUE_1(target)

	case *A_DEFAULT_VALUE_2:
		ok = stage.IsStagedA_DEFAULT_VALUE_2(target)

	case *A_DEFAULT_VALUE_3:
		ok = stage.IsStagedA_DEFAULT_VALUE_3(target)

	case *A_DEFAULT_VALUE_4:
		ok = stage.IsStagedA_DEFAULT_VALUE_4(target)

	case *A_DEFAULT_VALUE_5:
		ok = stage.IsStagedA_DEFAULT_VALUE_5(target)

	case *A_DEFAULT_VALUE_6:
		ok = stage.IsStagedA_DEFAULT_VALUE_6(target)

	case *A_DEFINITION:
		ok = stage.IsStagedA_DEFINITION(target)

	case *A_DEFINITION_1:
		ok = stage.IsStagedA_DEFINITION_1(target)

	case *A_DEFINITION_2:
		ok = stage.IsStagedA_DEFINITION_2(target)

	case *A_DEFINITION_3:
		ok = stage.IsStagedA_DEFINITION_3(target)

	case *A_DEFINITION_4:
		ok = stage.IsStagedA_DEFINITION_4(target)

	case *A_DEFINITION_5:
		ok = stage.IsStagedA_DEFINITION_5(target)

	case *A_DEFINITION_6:
		ok = stage.IsStagedA_DEFINITION_6(target)

	case *A_EDITABLE_ATTS:
		ok = stage.IsStagedA_EDITABLE_ATTS(target)

	case *A_OBJECT:
		ok = stage.IsStagedA_OBJECT(target)

	case *A_PROPERTIES:
		ok = stage.IsStagedA_PROPERTIES(target)

	case *A_SOURCE:
		ok = stage.IsStagedA_SOURCE(target)

	case *A_SOURCE_SPECIFICATION:
		ok = stage.IsStagedA_SOURCE_SPECIFICATION(target)

	case *A_SPECIFICATIONS:
		ok = stage.IsStagedA_SPECIFICATIONS(target)

	case *A_SPECIFIED_VALUES:
		ok = stage.IsStagedA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		ok = stage.IsStagedA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		ok = stage.IsStagedA_SPEC_OBJECTS(target)

	case *A_SPEC_RELATIONS:
		ok = stage.IsStagedA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATIONS_1:
		ok = stage.IsStagedA_SPEC_RELATIONS_1(target)

	case *A_SPEC_RELATION_GROUPS:
		ok = stage.IsStagedA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_TYPES:
		ok = stage.IsStagedA_SPEC_TYPES(target)

	case *A_THE_HEADER:
		ok = stage.IsStagedA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		ok = stage.IsStagedA_TOOL_EXTENSIONS(target)

	case *A_TYPE:
		ok = stage.IsStagedA_TYPE(target)

	case *A_TYPE_1:
		ok = stage.IsStagedA_TYPE_1(target)

	case *A_TYPE_10:
		ok = stage.IsStagedA_TYPE_10(target)

	case *A_TYPE_2:
		ok = stage.IsStagedA_TYPE_2(target)

	case *A_TYPE_3:
		ok = stage.IsStagedA_TYPE_3(target)

	case *A_TYPE_4:
		ok = stage.IsStagedA_TYPE_4(target)

	case *A_TYPE_5:
		ok = stage.IsStagedA_TYPE_5(target)

	case *A_TYPE_6:
		ok = stage.IsStagedA_TYPE_6(target)

	case *A_TYPE_7:
		ok = stage.IsStagedA_TYPE_7(target)

	case *A_TYPE_8:
		ok = stage.IsStagedA_TYPE_8(target)

	case *A_TYPE_9:
		ok = stage.IsStagedA_TYPE_9(target)

	case *A_VALUES:
		ok = stage.IsStagedA_VALUES(target)

	case *A_VALUES_1:
		ok = stage.IsStagedA_VALUES_1(target)

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

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE(a_default_value *A_DEFAULT_VALUE) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUEs[a_default_value]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_1(a_default_value_1 *A_DEFAULT_VALUE_1) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_1s[a_default_value_1]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_2(a_default_value_2 *A_DEFAULT_VALUE_2) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_2s[a_default_value_2]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_3(a_default_value_3 *A_DEFAULT_VALUE_3) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_3s[a_default_value_3]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_4(a_default_value_4 *A_DEFAULT_VALUE_4) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_4s[a_default_value_4]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_5(a_default_value_5 *A_DEFAULT_VALUE_5) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_5s[a_default_value_5]

	return
}

func (stage *StageStruct) IsStagedA_DEFAULT_VALUE_6(a_default_value_6 *A_DEFAULT_VALUE_6) (ok bool) {

	_, ok = stage.A_DEFAULT_VALUE_6s[a_default_value_6]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION(a_definition *A_DEFINITION) (ok bool) {

	_, ok = stage.A_DEFINITIONs[a_definition]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_1(a_definition_1 *A_DEFINITION_1) (ok bool) {

	_, ok = stage.A_DEFINITION_1s[a_definition_1]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_2(a_definition_2 *A_DEFINITION_2) (ok bool) {

	_, ok = stage.A_DEFINITION_2s[a_definition_2]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_3(a_definition_3 *A_DEFINITION_3) (ok bool) {

	_, ok = stage.A_DEFINITION_3s[a_definition_3]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_4(a_definition_4 *A_DEFINITION_4) (ok bool) {

	_, ok = stage.A_DEFINITION_4s[a_definition_4]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_5(a_definition_5 *A_DEFINITION_5) (ok bool) {

	_, ok = stage.A_DEFINITION_5s[a_definition_5]

	return
}

func (stage *StageStruct) IsStagedA_DEFINITION_6(a_definition_6 *A_DEFINITION_6) (ok bool) {

	_, ok = stage.A_DEFINITION_6s[a_definition_6]

	return
}

func (stage *StageStruct) IsStagedA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) (ok bool) {

	_, ok = stage.A_EDITABLE_ATTSs[a_editable_atts]

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

func (stage *StageStruct) IsStagedA_SOURCE(a_source *A_SOURCE) (ok bool) {

	_, ok = stage.A_SOURCEs[a_source]

	return
}

func (stage *StageStruct) IsStagedA_SOURCE_SPECIFICATION(a_source_specification *A_SOURCE_SPECIFICATION) (ok bool) {

	_, ok = stage.A_SOURCE_SPECIFICATIONs[a_source_specification]

	return
}

func (stage *StageStruct) IsStagedA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) (ok bool) {

	_, ok = stage.A_SPECIFICATIONSs[a_specifications]

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

func (stage *StageStruct) IsStagedA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) (ok bool) {

	_, ok = stage.A_SPEC_RELATIONSs[a_spec_relations]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATIONS_1(a_spec_relations_1 *A_SPEC_RELATIONS_1) (ok bool) {

	_, ok = stage.A_SPEC_RELATIONS_1s[a_spec_relations_1]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]

	return
}

func (stage *StageStruct) IsStagedA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) (ok bool) {

	_, ok = stage.A_SPEC_TYPESs[a_spec_types]

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

func (stage *StageStruct) IsStagedA_TYPE(a_type *A_TYPE) (ok bool) {

	_, ok = stage.A_TYPEs[a_type]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_1(a_type_1 *A_TYPE_1) (ok bool) {

	_, ok = stage.A_TYPE_1s[a_type_1]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_10(a_type_10 *A_TYPE_10) (ok bool) {

	_, ok = stage.A_TYPE_10s[a_type_10]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_2(a_type_2 *A_TYPE_2) (ok bool) {

	_, ok = stage.A_TYPE_2s[a_type_2]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_3(a_type_3 *A_TYPE_3) (ok bool) {

	_, ok = stage.A_TYPE_3s[a_type_3]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_4(a_type_4 *A_TYPE_4) (ok bool) {

	_, ok = stage.A_TYPE_4s[a_type_4]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_5(a_type_5 *A_TYPE_5) (ok bool) {

	_, ok = stage.A_TYPE_5s[a_type_5]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_6(a_type_6 *A_TYPE_6) (ok bool) {

	_, ok = stage.A_TYPE_6s[a_type_6]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_7(a_type_7 *A_TYPE_7) (ok bool) {

	_, ok = stage.A_TYPE_7s[a_type_7]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_8(a_type_8 *A_TYPE_8) (ok bool) {

	_, ok = stage.A_TYPE_8s[a_type_8]

	return
}

func (stage *StageStruct) IsStagedA_TYPE_9(a_type_9 *A_TYPE_9) (ok bool) {

	_, ok = stage.A_TYPE_9s[a_type_9]

	return
}

func (stage *StageStruct) IsStagedA_VALUES(a_values *A_VALUES) (ok bool) {

	_, ok = stage.A_VALUESs[a_values]

	return
}

func (stage *StageStruct) IsStagedA_VALUES_1(a_values_1 *A_VALUES_1) (ok bool) {

	_, ok = stage.A_VALUES_1s[a_values_1]

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

	case *A_CHILDREN:
		stage.StageBranchA_CHILDREN(target)

	case *A_CORE_CONTENT:
		stage.StageBranchA_CORE_CONTENT(target)

	case *A_DATATYPES:
		stage.StageBranchA_DATATYPES(target)

	case *A_DEFAULT_VALUE:
		stage.StageBranchA_DEFAULT_VALUE(target)

	case *A_DEFAULT_VALUE_1:
		stage.StageBranchA_DEFAULT_VALUE_1(target)

	case *A_DEFAULT_VALUE_2:
		stage.StageBranchA_DEFAULT_VALUE_2(target)

	case *A_DEFAULT_VALUE_3:
		stage.StageBranchA_DEFAULT_VALUE_3(target)

	case *A_DEFAULT_VALUE_4:
		stage.StageBranchA_DEFAULT_VALUE_4(target)

	case *A_DEFAULT_VALUE_5:
		stage.StageBranchA_DEFAULT_VALUE_5(target)

	case *A_DEFAULT_VALUE_6:
		stage.StageBranchA_DEFAULT_VALUE_6(target)

	case *A_DEFINITION:
		stage.StageBranchA_DEFINITION(target)

	case *A_DEFINITION_1:
		stage.StageBranchA_DEFINITION_1(target)

	case *A_DEFINITION_2:
		stage.StageBranchA_DEFINITION_2(target)

	case *A_DEFINITION_3:
		stage.StageBranchA_DEFINITION_3(target)

	case *A_DEFINITION_4:
		stage.StageBranchA_DEFINITION_4(target)

	case *A_DEFINITION_5:
		stage.StageBranchA_DEFINITION_5(target)

	case *A_DEFINITION_6:
		stage.StageBranchA_DEFINITION_6(target)

	case *A_EDITABLE_ATTS:
		stage.StageBranchA_EDITABLE_ATTS(target)

	case *A_OBJECT:
		stage.StageBranchA_OBJECT(target)

	case *A_PROPERTIES:
		stage.StageBranchA_PROPERTIES(target)

	case *A_SOURCE:
		stage.StageBranchA_SOURCE(target)

	case *A_SOURCE_SPECIFICATION:
		stage.StageBranchA_SOURCE_SPECIFICATION(target)

	case *A_SPECIFICATIONS:
		stage.StageBranchA_SPECIFICATIONS(target)

	case *A_SPECIFIED_VALUES:
		stage.StageBranchA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		stage.StageBranchA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		stage.StageBranchA_SPEC_OBJECTS(target)

	case *A_SPEC_RELATIONS:
		stage.StageBranchA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATIONS_1:
		stage.StageBranchA_SPEC_RELATIONS_1(target)

	case *A_SPEC_RELATION_GROUPS:
		stage.StageBranchA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_TYPES:
		stage.StageBranchA_SPEC_TYPES(target)

	case *A_THE_HEADER:
		stage.StageBranchA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		stage.StageBranchA_TOOL_EXTENSIONS(target)

	case *A_TYPE:
		stage.StageBranchA_TYPE(target)

	case *A_TYPE_1:
		stage.StageBranchA_TYPE_1(target)

	case *A_TYPE_10:
		stage.StageBranchA_TYPE_10(target)

	case *A_TYPE_2:
		stage.StageBranchA_TYPE_2(target)

	case *A_TYPE_3:
		stage.StageBranchA_TYPE_3(target)

	case *A_TYPE_4:
		stage.StageBranchA_TYPE_4(target)

	case *A_TYPE_5:
		stage.StageBranchA_TYPE_5(target)

	case *A_TYPE_6:
		stage.StageBranchA_TYPE_6(target)

	case *A_TYPE_7:
		stage.StageBranchA_TYPE_7(target)

	case *A_TYPE_8:
		stage.StageBranchA_TYPE_8(target)

	case *A_TYPE_9:
		stage.StageBranchA_TYPE_9(target)

	case *A_VALUES:
		stage.StageBranchA_VALUES(target)

	case *A_VALUES_1:
		stage.StageBranchA_VALUES_1(target)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_1 := range attribute_definition_boolean.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_1)
	}
	for _, _a_type_7 := range attribute_definition_boolean.TYPE {
		StageBranch(stage, _a_type_7)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_date.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_3 := range attribute_definition_date.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_3)
	}
	for _, _a_type := range attribute_definition_date.TYPE {
		StageBranch(stage, _a_type)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_5 := range attribute_definition_enumeration.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_5)
	}
	for _, _a_type_9 := range attribute_definition_enumeration.TYPE {
		StageBranch(stage, _a_type_9)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_integer.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_6 := range attribute_definition_integer.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_6)
	}
	for _, _a_type_5 := range attribute_definition_integer.TYPE {
		StageBranch(stage, _a_type_5)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_real.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_4 := range attribute_definition_real.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_4)
	}
	for _, _a_type_3 := range attribute_definition_real.TYPE {
		StageBranch(stage, _a_type_3)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_string.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value := range attribute_definition_string.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value)
	}
	for _, _a_type_8 := range attribute_definition_string.TYPE {
		StageBranch(stage, _a_type_8)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_2 := range attribute_definition_xhtml.DEFAULT_VALUE {
		StageBranch(stage, _a_default_value_2)
	}
	for _, _a_type_6 := range attribute_definition_xhtml.TYPE {
		StageBranch(stage, _a_type_6)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_4 := range attribute_value_boolean.DEFINITION {
		StageBranch(stage, _a_definition_4)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_2 := range attribute_value_date.DEFINITION {
		StageBranch(stage, _a_definition_2)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_5 := range attribute_value_enumeration.DEFINITION {
		StageBranch(stage, _a_definition_5)
	}
	for _, _a_values := range attribute_value_enumeration.VALUES {
		StageBranch(stage, _a_values)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition := range attribute_value_integer.DEFINITION {
		StageBranch(stage, _a_definition)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_6 := range attribute_value_real.DEFINITION {
		StageBranch(stage, _a_definition_6)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_3 := range attribute_value_string.DEFINITION {
		StageBranch(stage, _a_definition_3)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
		StageBranch(stage, _xhtml_content)
	}
	for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
		StageBranch(stage, _xhtml_content)
	}
	for _, _a_definition_1 := range attribute_value_xhtml.DEFINITION {
		StageBranch(stage, _a_definition_1)
	}

}

func (stage *StageStruct) StageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if IsStaged(stage, a_alternative_id) {
		return
	}

	a_alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range a_alternative_id.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if IsStaged(stage, a_children) {
		return
	}

	a_children.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
		StageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) StageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, a_core_content) {
		return
	}

	a_core_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_content := range a_core_content.REQ_IF_CONTENT {
		StageBranch(stage, _req_if_content)
	}

}

func (stage *StageStruct) StageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if IsStaged(stage, a_datatypes) {
		return
	}

	a_datatypes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
		StageBranch(stage, _datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
		StageBranch(stage, _datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
		StageBranch(stage, _datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
		StageBranch(stage, _datatype_definition_integer)
	}
	for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
		StageBranch(stage, _datatype_definition_real)
	}
	for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
		StageBranch(stage, _datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
		StageBranch(stage, _datatype_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE(a_default_value *A_DEFAULT_VALUE) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value) {
		return
	}

	a_default_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_default_value.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_1(a_default_value_1 *A_DEFAULT_VALUE_1) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_1) {
		return
	}

	a_default_value_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_default_value_1.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_2(a_default_value_2 *A_DEFAULT_VALUE_2) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_2) {
		return
	}

	a_default_value_2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_default_value_2.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_3(a_default_value_3 *A_DEFAULT_VALUE_3) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_3) {
		return
	}

	a_default_value_3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_default_value_3.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_4(a_default_value_4 *A_DEFAULT_VALUE_4) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_4) {
		return
	}

	a_default_value_4.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_default_value_4.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_5(a_default_value_5 *A_DEFAULT_VALUE_5) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_5) {
		return
	}

	a_default_value_5.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_default_value_5.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}

}

func (stage *StageStruct) StageBranchA_DEFAULT_VALUE_6(a_default_value_6 *A_DEFAULT_VALUE_6) {

	// check if instance is already staged
	if IsStaged(stage, a_default_value_6) {
		return
	}

	a_default_value_6.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_default_value_6.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}

}

func (stage *StageStruct) StageBranchA_DEFINITION(a_definition *A_DEFINITION) {

	// check if instance is already staged
	if IsStaged(stage, a_definition) {
		return
	}

	a_definition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_1(a_definition_1 *A_DEFINITION_1) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_1) {
		return
	}

	a_definition_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_2(a_definition_2 *A_DEFINITION_2) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_2) {
		return
	}

	a_definition_2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_3(a_definition_3 *A_DEFINITION_3) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_3) {
		return
	}

	a_definition_3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_4(a_definition_4 *A_DEFINITION_4) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_4) {
		return
	}

	a_definition_4.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_5(a_definition_5 *A_DEFINITION_5) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_5) {
		return
	}

	a_definition_5.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_DEFINITION_6(a_definition_6 *A_DEFINITION_6) {

	// check if instance is already staged
	if IsStaged(stage, a_definition_6) {
		return
	}

	a_definition_6.Stage(stage)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _embedded_value := range a_properties.EMBEDDED_VALUE {
		StageBranch(stage, _embedded_value)
	}

}

func (stage *StageStruct) StageBranchA_SOURCE(a_source *A_SOURCE) {

	// check if instance is already staged
	if IsStaged(stage, a_source) {
		return
	}

	a_source.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SOURCE_SPECIFICATION(a_source_specification *A_SOURCE_SPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, a_source_specification) {
		return
	}

	a_source_specification.Stage(stage)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specifications.SPECIFICATION {
		StageBranch(stage, _specification)
	}

}

func (stage *StageStruct) StageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if IsStaged(stage, a_specified_values) {
		return
	}

	a_specified_values.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_values.ENUM_VALUE {
		StageBranch(stage, _enum_value)
	}

}

func (stage *StageStruct) StageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_attributes) {
		return
	}

	a_spec_attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
		StageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
		StageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
		StageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
		StageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
		StageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
		StageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
		StageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_objects) {
		return
	}

	a_spec_objects.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
		StageBranch(stage, _spec_object)
	}

}

func (stage *StageStruct) StageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relations) {
		return
	}

	a_spec_relations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_SPEC_RELATIONS_1(a_spec_relations_1 *A_SPEC_RELATIONS_1) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relations_1) {
		return
	}

	a_spec_relations_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relations_1.SPEC_RELATION {
		StageBranch(stage, _spec_relation)
	}

}

func (stage *StageStruct) StageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
		StageBranch(stage, _relation_group)
	}

}

func (stage *StageStruct) StageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if IsStaged(stage, a_spec_types) {
		return
	}

	a_spec_types.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
		StageBranch(stage, _relation_group_type)
	}
	for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
		StageBranch(stage, _spec_object_type)
	}
	for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
		StageBranch(stage, _spec_relation_type)
	}
	for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
		StageBranch(stage, _specification_type)
	}

}

func (stage *StageStruct) StageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if IsStaged(stage, a_the_header) {
		return
	}

	a_the_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_header := range a_the_header.REQ_IF_HEADER {
		StageBranch(stage, _req_if_header)
	}

}

func (stage *StageStruct) StageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if IsStaged(stage, a_tool_extensions) {
		return
	}

	a_tool_extensions.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
		StageBranch(stage, _req_if_tool_extension)
	}

}

func (stage *StageStruct) StageBranchA_TYPE(a_type *A_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, a_type) {
		return
	}

	a_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_1(a_type_1 *A_TYPE_1) {

	// check if instance is already staged
	if IsStaged(stage, a_type_1) {
		return
	}

	a_type_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_10(a_type_10 *A_TYPE_10) {

	// check if instance is already staged
	if IsStaged(stage, a_type_10) {
		return
	}

	a_type_10.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_2(a_type_2 *A_TYPE_2) {

	// check if instance is already staged
	if IsStaged(stage, a_type_2) {
		return
	}

	a_type_2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_3(a_type_3 *A_TYPE_3) {

	// check if instance is already staged
	if IsStaged(stage, a_type_3) {
		return
	}

	a_type_3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_4(a_type_4 *A_TYPE_4) {

	// check if instance is already staged
	if IsStaged(stage, a_type_4) {
		return
	}

	a_type_4.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_5(a_type_5 *A_TYPE_5) {

	// check if instance is already staged
	if IsStaged(stage, a_type_5) {
		return
	}

	a_type_5.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_6(a_type_6 *A_TYPE_6) {

	// check if instance is already staged
	if IsStaged(stage, a_type_6) {
		return
	}

	a_type_6.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_7(a_type_7 *A_TYPE_7) {

	// check if instance is already staged
	if IsStaged(stage, a_type_7) {
		return
	}

	a_type_7.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_8(a_type_8 *A_TYPE_8) {

	// check if instance is already staged
	if IsStaged(stage, a_type_8) {
		return
	}

	a_type_8.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_TYPE_9(a_type_9 *A_TYPE_9) {

	// check if instance is already staged
	if IsStaged(stage, a_type_9) {
		return
	}

	a_type_9.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_VALUES(a_values *A_VALUES) {

	// check if instance is already staged
	if IsStaged(stage, a_values) {
		return
	}

	a_values.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchA_VALUES_1(a_values_1 *A_VALUES_1) {

	// check if instance is already staged
	if IsStaged(stage, a_values_1) {
		return
	}

	a_values_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_values_1.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range a_values_1.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range a_values_1.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range a_values_1.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range a_values_1.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range a_values_1.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range a_values_1.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_date.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_specified_values := range datatype_definition_enumeration.SPECIFIED_VALUES {
		StageBranch(stage, _a_specified_values)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_integer.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_real.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_string.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range enum_value.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_properties := range enum_value.PROPERTIES {
		StageBranch(stage, _a_properties)
	}

}

func (stage *StageStruct) StageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if IsStaged(stage, relation_group) {
		return
	}

	relation_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_group.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_source_specification := range relation_group.SOURCE_SPECIFICATION {
		StageBranch(stage, _a_source_specification)
	}
	for _, _a_spec_relations := range relation_group.SPEC_RELATIONS {
		StageBranch(stage, _a_spec_relations)
	}
	for _, _a_type_1 := range relation_group.TYPE {
		StageBranch(stage, _a_type_1)
	}

}

func (stage *StageStruct) StageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_group_type.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range relation_group_type.SPEC_ATTRIBUTES {
		StageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) StageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if IsStaged(stage, req_if) {
		return
	}

	req_if.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_the_header := range req_if.THE_HEADER {
		StageBranch(stage, _a_the_header)
	}
	for _, _a_core_content := range req_if.CORE_CONTENT {
		StageBranch(stage, _a_core_content)
	}
	for _, _a_tool_extensions := range req_if.TOOL_EXTENSIONS {
		StageBranch(stage, _a_tool_extensions)
	}

}

func (stage *StageStruct) StageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_datatypes := range req_if_content.DATATYPES {
		StageBranch(stage, _a_datatypes)
	}
	for _, _a_spec_types := range req_if_content.SPEC_TYPES {
		StageBranch(stage, _a_spec_types)
	}
	for _, _a_spec_objects := range req_if_content.SPEC_OBJECTS {
		StageBranch(stage, _a_spec_objects)
	}
	for _, _a_spec_relations_1 := range req_if_content.SPEC_RELATIONS {
		StageBranch(stage, _a_spec_relations_1)
	}
	for _, _a_specifications := range req_if_content.SPECIFICATIONS {
		StageBranch(stage, _a_specifications)
	}
	for _, _a_spec_relation_groups := range req_if_content.SPEC_RELATION_GROUPS {
		StageBranch(stage, _a_spec_relation_groups)
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specification.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_children := range specification.CHILDREN {
		StageBranch(stage, _a_children)
	}
	for _, _a_values_1 := range specification.VALUES {
		StageBranch(stage, _a_values_1)
	}
	for _, _a_type_10 := range specification.TYPE {
		StageBranch(stage, _a_type_10)
	}

}

func (stage *StageStruct) StageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, specification_type) {
		return
	}

	specification_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specification_type.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range specification_type.SPEC_ATTRIBUTES {
		StageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) StageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_hierarchy.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_children := range spec_hierarchy.CHILDREN {
		StageBranch(stage, _a_children)
	}
	for _, _a_editable_atts := range spec_hierarchy.EDITABLE_ATTS {
		StageBranch(stage, _a_editable_atts)
	}
	for _, _a_object := range spec_hierarchy.OBJECT {
		StageBranch(stage, _a_object)
	}

}

func (stage *StageStruct) StageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if IsStaged(stage, spec_object) {
		return
	}

	spec_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_object.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_values_1 := range spec_object.VALUES {
		StageBranch(stage, _a_values_1)
	}
	for _, _a_type_2 := range spec_object.TYPE {
		StageBranch(stage, _a_type_2)
	}

}

func (stage *StageStruct) StageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_object_type.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range spec_object_type.SPEC_ATTRIBUTES {
		StageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) StageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relation.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_values_1 := range spec_relation.VALUES {
		StageBranch(stage, _a_values_1)
	}
	for _, _a_source := range spec_relation.SOURCE {
		StageBranch(stage, _a_source)
	}
	for _, _a_type_4 := range spec_relation.TYPE {
		StageBranch(stage, _a_type_4)
	}

}

func (stage *StageStruct) StageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relation_type.ALTERNATIVE_ID {
		StageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range spec_relation_type.SPEC_ATTRIBUTES {
		StageBranch(stage, _a_spec_attributes)
	}

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

	case *A_CHILDREN:
		toT := CopyBranchA_CHILDREN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_CORE_CONTENT:
		toT := CopyBranchA_CORE_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPES:
		toT := CopyBranchA_DATATYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE:
		toT := CopyBranchA_DEFAULT_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_1:
		toT := CopyBranchA_DEFAULT_VALUE_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_2:
		toT := CopyBranchA_DEFAULT_VALUE_2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_3:
		toT := CopyBranchA_DEFAULT_VALUE_3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_4:
		toT := CopyBranchA_DEFAULT_VALUE_4(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_5:
		toT := CopyBranchA_DEFAULT_VALUE_5(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFAULT_VALUE_6:
		toT := CopyBranchA_DEFAULT_VALUE_6(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION:
		toT := CopyBranchA_DEFINITION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_1:
		toT := CopyBranchA_DEFINITION_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_2:
		toT := CopyBranchA_DEFINITION_2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_3:
		toT := CopyBranchA_DEFINITION_3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_4:
		toT := CopyBranchA_DEFINITION_4(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_5:
		toT := CopyBranchA_DEFINITION_5(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DEFINITION_6:
		toT := CopyBranchA_DEFINITION_6(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_EDITABLE_ATTS:
		toT := CopyBranchA_EDITABLE_ATTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_OBJECT:
		toT := CopyBranchA_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_PROPERTIES:
		toT := CopyBranchA_PROPERTIES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SOURCE:
		toT := CopyBranchA_SOURCE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SOURCE_SPECIFICATION:
		toT := CopyBranchA_SOURCE_SPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFICATIONS:
		toT := CopyBranchA_SPECIFICATIONS(mapOrigCopy, fromT)
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

	case *A_SPEC_RELATIONS:
		toT := CopyBranchA_SPEC_RELATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATIONS_1:
		toT := CopyBranchA_SPEC_RELATIONS_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_GROUPS:
		toT := CopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_TYPES:
		toT := CopyBranchA_SPEC_TYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_THE_HEADER:
		toT := CopyBranchA_THE_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TOOL_EXTENSIONS:
		toT := CopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE:
		toT := CopyBranchA_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_1:
		toT := CopyBranchA_TYPE_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_10:
		toT := CopyBranchA_TYPE_10(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_2:
		toT := CopyBranchA_TYPE_2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_3:
		toT := CopyBranchA_TYPE_3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_4:
		toT := CopyBranchA_TYPE_4(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_5:
		toT := CopyBranchA_TYPE_5(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_6:
		toT := CopyBranchA_TYPE_6(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_7:
		toT := CopyBranchA_TYPE_7(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_8:
		toT := CopyBranchA_TYPE_8(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TYPE_9:
		toT := CopyBranchA_TYPE_9(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_VALUES:
		toT := CopyBranchA_VALUES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_VALUES_1:
		toT := CopyBranchA_VALUES_1(mapOrigCopy, fromT)
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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_booleanFrom.ALTERNATIVE_ID {
		attribute_definition_booleanTo.ALTERNATIVE_ID = append(attribute_definition_booleanTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_1 := range attribute_definition_booleanFrom.DEFAULT_VALUE {
		attribute_definition_booleanTo.DEFAULT_VALUE = append(attribute_definition_booleanTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_1(mapOrigCopy, _a_default_value_1))
	}
	for _, _a_type_7 := range attribute_definition_booleanFrom.TYPE {
		attribute_definition_booleanTo.TYPE = append(attribute_definition_booleanTo.TYPE, CopyBranchA_TYPE_7(mapOrigCopy, _a_type_7))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_dateFrom.ALTERNATIVE_ID {
		attribute_definition_dateTo.ALTERNATIVE_ID = append(attribute_definition_dateTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_3 := range attribute_definition_dateFrom.DEFAULT_VALUE {
		attribute_definition_dateTo.DEFAULT_VALUE = append(attribute_definition_dateTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_3(mapOrigCopy, _a_default_value_3))
	}
	for _, _a_type := range attribute_definition_dateFrom.TYPE {
		attribute_definition_dateTo.TYPE = append(attribute_definition_dateTo.TYPE, CopyBranchA_TYPE(mapOrigCopy, _a_type))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_enumerationFrom.ALTERNATIVE_ID {
		attribute_definition_enumerationTo.ALTERNATIVE_ID = append(attribute_definition_enumerationTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_5 := range attribute_definition_enumerationFrom.DEFAULT_VALUE {
		attribute_definition_enumerationTo.DEFAULT_VALUE = append(attribute_definition_enumerationTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_5(mapOrigCopy, _a_default_value_5))
	}
	for _, _a_type_9 := range attribute_definition_enumerationFrom.TYPE {
		attribute_definition_enumerationTo.TYPE = append(attribute_definition_enumerationTo.TYPE, CopyBranchA_TYPE_9(mapOrigCopy, _a_type_9))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_integerFrom.ALTERNATIVE_ID {
		attribute_definition_integerTo.ALTERNATIVE_ID = append(attribute_definition_integerTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_6 := range attribute_definition_integerFrom.DEFAULT_VALUE {
		attribute_definition_integerTo.DEFAULT_VALUE = append(attribute_definition_integerTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_6(mapOrigCopy, _a_default_value_6))
	}
	for _, _a_type_5 := range attribute_definition_integerFrom.TYPE {
		attribute_definition_integerTo.TYPE = append(attribute_definition_integerTo.TYPE, CopyBranchA_TYPE_5(mapOrigCopy, _a_type_5))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_realFrom.ALTERNATIVE_ID {
		attribute_definition_realTo.ALTERNATIVE_ID = append(attribute_definition_realTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_4 := range attribute_definition_realFrom.DEFAULT_VALUE {
		attribute_definition_realTo.DEFAULT_VALUE = append(attribute_definition_realTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_4(mapOrigCopy, _a_default_value_4))
	}
	for _, _a_type_3 := range attribute_definition_realFrom.TYPE {
		attribute_definition_realTo.TYPE = append(attribute_definition_realTo.TYPE, CopyBranchA_TYPE_3(mapOrigCopy, _a_type_3))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_stringFrom.ALTERNATIVE_ID {
		attribute_definition_stringTo.ALTERNATIVE_ID = append(attribute_definition_stringTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value := range attribute_definition_stringFrom.DEFAULT_VALUE {
		attribute_definition_stringTo.DEFAULT_VALUE = append(attribute_definition_stringTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE(mapOrigCopy, _a_default_value))
	}
	for _, _a_type_8 := range attribute_definition_stringFrom.TYPE {
		attribute_definition_stringTo.TYPE = append(attribute_definition_stringTo.TYPE, CopyBranchA_TYPE_8(mapOrigCopy, _a_type_8))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_xhtmlFrom.ALTERNATIVE_ID {
		attribute_definition_xhtmlTo.ALTERNATIVE_ID = append(attribute_definition_xhtmlTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_default_value_2 := range attribute_definition_xhtmlFrom.DEFAULT_VALUE {
		attribute_definition_xhtmlTo.DEFAULT_VALUE = append(attribute_definition_xhtmlTo.DEFAULT_VALUE, CopyBranchA_DEFAULT_VALUE_2(mapOrigCopy, _a_default_value_2))
	}
	for _, _a_type_6 := range attribute_definition_xhtmlFrom.TYPE {
		attribute_definition_xhtmlTo.TYPE = append(attribute_definition_xhtmlTo.TYPE, CopyBranchA_TYPE_6(mapOrigCopy, _a_type_6))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_4 := range attribute_value_booleanFrom.DEFINITION {
		attribute_value_booleanTo.DEFINITION = append(attribute_value_booleanTo.DEFINITION, CopyBranchA_DEFINITION_4(mapOrigCopy, _a_definition_4))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_2 := range attribute_value_dateFrom.DEFINITION {
		attribute_value_dateTo.DEFINITION = append(attribute_value_dateTo.DEFINITION, CopyBranchA_DEFINITION_2(mapOrigCopy, _a_definition_2))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_5 := range attribute_value_enumerationFrom.DEFINITION {
		attribute_value_enumerationTo.DEFINITION = append(attribute_value_enumerationTo.DEFINITION, CopyBranchA_DEFINITION_5(mapOrigCopy, _a_definition_5))
	}
	for _, _a_values := range attribute_value_enumerationFrom.VALUES {
		attribute_value_enumerationTo.VALUES = append(attribute_value_enumerationTo.VALUES, CopyBranchA_VALUES(mapOrigCopy, _a_values))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition := range attribute_value_integerFrom.DEFINITION {
		attribute_value_integerTo.DEFINITION = append(attribute_value_integerTo.DEFINITION, CopyBranchA_DEFINITION(mapOrigCopy, _a_definition))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_6 := range attribute_value_realFrom.DEFINITION {
		attribute_value_realTo.DEFINITION = append(attribute_value_realTo.DEFINITION, CopyBranchA_DEFINITION_6(mapOrigCopy, _a_definition_6))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_3 := range attribute_value_stringFrom.DEFINITION {
		attribute_value_stringTo.DEFINITION = append(attribute_value_stringTo.DEFINITION, CopyBranchA_DEFINITION_3(mapOrigCopy, _a_definition_3))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtmlFrom.THE_VALUE {
		attribute_value_xhtmlTo.THE_VALUE = append(attribute_value_xhtmlTo.THE_VALUE, CopyBranchXHTML_CONTENT(mapOrigCopy, _xhtml_content))
	}
	for _, _xhtml_content := range attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE {
		attribute_value_xhtmlTo.THE_ORIGINAL_VALUE = append(attribute_value_xhtmlTo.THE_ORIGINAL_VALUE, CopyBranchXHTML_CONTENT(mapOrigCopy, _xhtml_content))
	}
	for _, _a_definition_1 := range attribute_value_xhtmlFrom.DEFINITION {
		attribute_value_xhtmlTo.DEFINITION = append(attribute_value_xhtmlTo.DEFINITION, CopyBranchA_DEFINITION_1(mapOrigCopy, _a_definition_1))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range a_alternative_idFrom.ALTERNATIVE_ID {
		a_alternative_idTo.ALTERNATIVE_ID = append(a_alternative_idTo.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_childrenFrom.SPEC_HIERARCHY {
		a_childrenTo.SPEC_HIERARCHY = append(a_childrenTo.SPEC_HIERARCHY, CopyBranchSPEC_HIERARCHY(mapOrigCopy, _spec_hierarchy))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_content := range a_core_contentFrom.REQ_IF_CONTENT {
		a_core_contentTo.REQ_IF_CONTENT = append(a_core_contentTo.REQ_IF_CONTENT, CopyBranchREQ_IF_CONTENT(mapOrigCopy, _req_if_content))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypesFrom.DATATYPE_DEFINITION_BOOLEAN {
		a_datatypesTo.DATATYPE_DEFINITION_BOOLEAN = append(a_datatypesTo.DATATYPE_DEFINITION_BOOLEAN, CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, _datatype_definition_boolean))
	}
	for _, _datatype_definition_date := range a_datatypesFrom.DATATYPE_DEFINITION_DATE {
		a_datatypesTo.DATATYPE_DEFINITION_DATE = append(a_datatypesTo.DATATYPE_DEFINITION_DATE, CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, _datatype_definition_date))
	}
	for _, _datatype_definition_enumeration := range a_datatypesFrom.DATATYPE_DEFINITION_ENUMERATION {
		a_datatypesTo.DATATYPE_DEFINITION_ENUMERATION = append(a_datatypesTo.DATATYPE_DEFINITION_ENUMERATION, CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, _datatype_definition_enumeration))
	}
	for _, _datatype_definition_integer := range a_datatypesFrom.DATATYPE_DEFINITION_INTEGER {
		a_datatypesTo.DATATYPE_DEFINITION_INTEGER = append(a_datatypesTo.DATATYPE_DEFINITION_INTEGER, CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, _datatype_definition_integer))
	}
	for _, _datatype_definition_real := range a_datatypesFrom.DATATYPE_DEFINITION_REAL {
		a_datatypesTo.DATATYPE_DEFINITION_REAL = append(a_datatypesTo.DATATYPE_DEFINITION_REAL, CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, _datatype_definition_real))
	}
	for _, _datatype_definition_string := range a_datatypesFrom.DATATYPE_DEFINITION_STRING {
		a_datatypesTo.DATATYPE_DEFINITION_STRING = append(a_datatypesTo.DATATYPE_DEFINITION_STRING, CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, _datatype_definition_string))
	}
	for _, _datatype_definition_xhtml := range a_datatypesFrom.DATATYPE_DEFINITION_XHTML {
		a_datatypesTo.DATATYPE_DEFINITION_XHTML = append(a_datatypesTo.DATATYPE_DEFINITION_XHTML, CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, _datatype_definition_xhtml))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE(mapOrigCopy map[any]any, a_default_valueFrom *A_DEFAULT_VALUE) (a_default_valueTo *A_DEFAULT_VALUE) {

	// a_default_valueFrom has already been copied
	if _a_default_valueTo, ok := mapOrigCopy[a_default_valueFrom]; ok {
		a_default_valueTo = _a_default_valueTo.(*A_DEFAULT_VALUE)
		return
	}

	a_default_valueTo = new(A_DEFAULT_VALUE)
	mapOrigCopy[a_default_valueFrom] = a_default_valueTo
	a_default_valueFrom.CopyBasicFields(a_default_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_default_valueFrom.ATTRIBUTE_VALUE_STRING {
		a_default_valueTo.ATTRIBUTE_VALUE_STRING = append(a_default_valueTo.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_1(mapOrigCopy map[any]any, a_default_value_1From *A_DEFAULT_VALUE_1) (a_default_value_1To *A_DEFAULT_VALUE_1) {

	// a_default_value_1From has already been copied
	if _a_default_value_1To, ok := mapOrigCopy[a_default_value_1From]; ok {
		a_default_value_1To = _a_default_value_1To.(*A_DEFAULT_VALUE_1)
		return
	}

	a_default_value_1To = new(A_DEFAULT_VALUE_1)
	mapOrigCopy[a_default_value_1From] = a_default_value_1To
	a_default_value_1From.CopyBasicFields(a_default_value_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_default_value_1From.ATTRIBUTE_VALUE_BOOLEAN {
		a_default_value_1To.ATTRIBUTE_VALUE_BOOLEAN = append(a_default_value_1To.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_2(mapOrigCopy map[any]any, a_default_value_2From *A_DEFAULT_VALUE_2) (a_default_value_2To *A_DEFAULT_VALUE_2) {

	// a_default_value_2From has already been copied
	if _a_default_value_2To, ok := mapOrigCopy[a_default_value_2From]; ok {
		a_default_value_2To = _a_default_value_2To.(*A_DEFAULT_VALUE_2)
		return
	}

	a_default_value_2To = new(A_DEFAULT_VALUE_2)
	mapOrigCopy[a_default_value_2From] = a_default_value_2To
	a_default_value_2From.CopyBasicFields(a_default_value_2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_default_value_2From.ATTRIBUTE_VALUE_XHTML {
		a_default_value_2To.ATTRIBUTE_VALUE_XHTML = append(a_default_value_2To.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_3(mapOrigCopy map[any]any, a_default_value_3From *A_DEFAULT_VALUE_3) (a_default_value_3To *A_DEFAULT_VALUE_3) {

	// a_default_value_3From has already been copied
	if _a_default_value_3To, ok := mapOrigCopy[a_default_value_3From]; ok {
		a_default_value_3To = _a_default_value_3To.(*A_DEFAULT_VALUE_3)
		return
	}

	a_default_value_3To = new(A_DEFAULT_VALUE_3)
	mapOrigCopy[a_default_value_3From] = a_default_value_3To
	a_default_value_3From.CopyBasicFields(a_default_value_3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_default_value_3From.ATTRIBUTE_VALUE_DATE {
		a_default_value_3To.ATTRIBUTE_VALUE_DATE = append(a_default_value_3To.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_4(mapOrigCopy map[any]any, a_default_value_4From *A_DEFAULT_VALUE_4) (a_default_value_4To *A_DEFAULT_VALUE_4) {

	// a_default_value_4From has already been copied
	if _a_default_value_4To, ok := mapOrigCopy[a_default_value_4From]; ok {
		a_default_value_4To = _a_default_value_4To.(*A_DEFAULT_VALUE_4)
		return
	}

	a_default_value_4To = new(A_DEFAULT_VALUE_4)
	mapOrigCopy[a_default_value_4From] = a_default_value_4To
	a_default_value_4From.CopyBasicFields(a_default_value_4To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_default_value_4From.ATTRIBUTE_VALUE_REAL {
		a_default_value_4To.ATTRIBUTE_VALUE_REAL = append(a_default_value_4To.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_5(mapOrigCopy map[any]any, a_default_value_5From *A_DEFAULT_VALUE_5) (a_default_value_5To *A_DEFAULT_VALUE_5) {

	// a_default_value_5From has already been copied
	if _a_default_value_5To, ok := mapOrigCopy[a_default_value_5From]; ok {
		a_default_value_5To = _a_default_value_5To.(*A_DEFAULT_VALUE_5)
		return
	}

	a_default_value_5To = new(A_DEFAULT_VALUE_5)
	mapOrigCopy[a_default_value_5From] = a_default_value_5To
	a_default_value_5From.CopyBasicFields(a_default_value_5To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_default_value_5From.ATTRIBUTE_VALUE_ENUMERATION {
		a_default_value_5To.ATTRIBUTE_VALUE_ENUMERATION = append(a_default_value_5To.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}

	return
}

func CopyBranchA_DEFAULT_VALUE_6(mapOrigCopy map[any]any, a_default_value_6From *A_DEFAULT_VALUE_6) (a_default_value_6To *A_DEFAULT_VALUE_6) {

	// a_default_value_6From has already been copied
	if _a_default_value_6To, ok := mapOrigCopy[a_default_value_6From]; ok {
		a_default_value_6To = _a_default_value_6To.(*A_DEFAULT_VALUE_6)
		return
	}

	a_default_value_6To = new(A_DEFAULT_VALUE_6)
	mapOrigCopy[a_default_value_6From] = a_default_value_6To
	a_default_value_6From.CopyBasicFields(a_default_value_6To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_default_value_6From.ATTRIBUTE_VALUE_INTEGER {
		a_default_value_6To.ATTRIBUTE_VALUE_INTEGER = append(a_default_value_6To.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}

	return
}

func CopyBranchA_DEFINITION(mapOrigCopy map[any]any, a_definitionFrom *A_DEFINITION) (a_definitionTo *A_DEFINITION) {

	// a_definitionFrom has already been copied
	if _a_definitionTo, ok := mapOrigCopy[a_definitionFrom]; ok {
		a_definitionTo = _a_definitionTo.(*A_DEFINITION)
		return
	}

	a_definitionTo = new(A_DEFINITION)
	mapOrigCopy[a_definitionFrom] = a_definitionTo
	a_definitionFrom.CopyBasicFields(a_definitionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_1(mapOrigCopy map[any]any, a_definition_1From *A_DEFINITION_1) (a_definition_1To *A_DEFINITION_1) {

	// a_definition_1From has already been copied
	if _a_definition_1To, ok := mapOrigCopy[a_definition_1From]; ok {
		a_definition_1To = _a_definition_1To.(*A_DEFINITION_1)
		return
	}

	a_definition_1To = new(A_DEFINITION_1)
	mapOrigCopy[a_definition_1From] = a_definition_1To
	a_definition_1From.CopyBasicFields(a_definition_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_2(mapOrigCopy map[any]any, a_definition_2From *A_DEFINITION_2) (a_definition_2To *A_DEFINITION_2) {

	// a_definition_2From has already been copied
	if _a_definition_2To, ok := mapOrigCopy[a_definition_2From]; ok {
		a_definition_2To = _a_definition_2To.(*A_DEFINITION_2)
		return
	}

	a_definition_2To = new(A_DEFINITION_2)
	mapOrigCopy[a_definition_2From] = a_definition_2To
	a_definition_2From.CopyBasicFields(a_definition_2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_3(mapOrigCopy map[any]any, a_definition_3From *A_DEFINITION_3) (a_definition_3To *A_DEFINITION_3) {

	// a_definition_3From has already been copied
	if _a_definition_3To, ok := mapOrigCopy[a_definition_3From]; ok {
		a_definition_3To = _a_definition_3To.(*A_DEFINITION_3)
		return
	}

	a_definition_3To = new(A_DEFINITION_3)
	mapOrigCopy[a_definition_3From] = a_definition_3To
	a_definition_3From.CopyBasicFields(a_definition_3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_4(mapOrigCopy map[any]any, a_definition_4From *A_DEFINITION_4) (a_definition_4To *A_DEFINITION_4) {

	// a_definition_4From has already been copied
	if _a_definition_4To, ok := mapOrigCopy[a_definition_4From]; ok {
		a_definition_4To = _a_definition_4To.(*A_DEFINITION_4)
		return
	}

	a_definition_4To = new(A_DEFINITION_4)
	mapOrigCopy[a_definition_4From] = a_definition_4To
	a_definition_4From.CopyBasicFields(a_definition_4To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_5(mapOrigCopy map[any]any, a_definition_5From *A_DEFINITION_5) (a_definition_5To *A_DEFINITION_5) {

	// a_definition_5From has already been copied
	if _a_definition_5To, ok := mapOrigCopy[a_definition_5From]; ok {
		a_definition_5To = _a_definition_5To.(*A_DEFINITION_5)
		return
	}

	a_definition_5To = new(A_DEFINITION_5)
	mapOrigCopy[a_definition_5From] = a_definition_5To
	a_definition_5From.CopyBasicFields(a_definition_5To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_DEFINITION_6(mapOrigCopy map[any]any, a_definition_6From *A_DEFINITION_6) (a_definition_6To *A_DEFINITION_6) {

	// a_definition_6From has already been copied
	if _a_definition_6To, ok := mapOrigCopy[a_definition_6From]; ok {
		a_definition_6To = _a_definition_6To.(*A_DEFINITION_6)
		return
	}

	a_definition_6To = new(A_DEFINITION_6)
	mapOrigCopy[a_definition_6From] = a_definition_6To
	a_definition_6From.CopyBasicFields(a_definition_6To)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _embedded_value := range a_propertiesFrom.EMBEDDED_VALUE {
		a_propertiesTo.EMBEDDED_VALUE = append(a_propertiesTo.EMBEDDED_VALUE, CopyBranchEMBEDDED_VALUE(mapOrigCopy, _embedded_value))
	}

	return
}

func CopyBranchA_SOURCE(mapOrigCopy map[any]any, a_sourceFrom *A_SOURCE) (a_sourceTo *A_SOURCE) {

	// a_sourceFrom has already been copied
	if _a_sourceTo, ok := mapOrigCopy[a_sourceFrom]; ok {
		a_sourceTo = _a_sourceTo.(*A_SOURCE)
		return
	}

	a_sourceTo = new(A_SOURCE)
	mapOrigCopy[a_sourceFrom] = a_sourceTo
	a_sourceFrom.CopyBasicFields(a_sourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SOURCE_SPECIFICATION(mapOrigCopy map[any]any, a_source_specificationFrom *A_SOURCE_SPECIFICATION) (a_source_specificationTo *A_SOURCE_SPECIFICATION) {

	// a_source_specificationFrom has already been copied
	if _a_source_specificationTo, ok := mapOrigCopy[a_source_specificationFrom]; ok {
		a_source_specificationTo = _a_source_specificationTo.(*A_SOURCE_SPECIFICATION)
		return
	}

	a_source_specificationTo = new(A_SOURCE_SPECIFICATION)
	mapOrigCopy[a_source_specificationFrom] = a_source_specificationTo
	a_source_specificationFrom.CopyBasicFields(a_source_specificationTo)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specificationsFrom.SPECIFICATION {
		a_specificationsTo.SPECIFICATION = append(a_specificationsTo.SPECIFICATION, CopyBranchSPECIFICATION(mapOrigCopy, _specification))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_valuesFrom.ENUM_VALUE {
		a_specified_valuesTo.ENUM_VALUE = append(a_specified_valuesTo.ENUM_VALUE, CopyBranchENUM_VALUE(mapOrigCopy, _enum_value))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_BOOLEAN {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_BOOLEAN = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_BOOLEAN, CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_DATE {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_DATE = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_DATE, CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_ENUMERATION {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_ENUMERATION = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_ENUMERATION, CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_INTEGER {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_INTEGER = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_INTEGER, CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_REAL {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_REAL = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_REAL, CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_STRING {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_STRING = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_STRING, CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_XHTML {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_XHTML = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_XHTML, CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objectsFrom.SPEC_OBJECT {
		a_spec_objectsTo.SPEC_OBJECT = append(a_spec_objectsTo.SPEC_OBJECT, CopyBranchSPEC_OBJECT(mapOrigCopy, _spec_object))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_SPEC_RELATIONS_1(mapOrigCopy map[any]any, a_spec_relations_1From *A_SPEC_RELATIONS_1) (a_spec_relations_1To *A_SPEC_RELATIONS_1) {

	// a_spec_relations_1From has already been copied
	if _a_spec_relations_1To, ok := mapOrigCopy[a_spec_relations_1From]; ok {
		a_spec_relations_1To = _a_spec_relations_1To.(*A_SPEC_RELATIONS_1)
		return
	}

	a_spec_relations_1To = new(A_SPEC_RELATIONS_1)
	mapOrigCopy[a_spec_relations_1From] = a_spec_relations_1To
	a_spec_relations_1From.CopyBasicFields(a_spec_relations_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relations_1From.SPEC_RELATION {
		a_spec_relations_1To.SPEC_RELATION = append(a_spec_relations_1To.SPEC_RELATION, CopyBranchSPEC_RELATION(mapOrigCopy, _spec_relation))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groupsFrom.RELATION_GROUP {
		a_spec_relation_groupsTo.RELATION_GROUP = append(a_spec_relation_groupsTo.RELATION_GROUP, CopyBranchRELATION_GROUP(mapOrigCopy, _relation_group))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_typesFrom.RELATION_GROUP_TYPE {
		a_spec_typesTo.RELATION_GROUP_TYPE = append(a_spec_typesTo.RELATION_GROUP_TYPE, CopyBranchRELATION_GROUP_TYPE(mapOrigCopy, _relation_group_type))
	}
	for _, _spec_object_type := range a_spec_typesFrom.SPEC_OBJECT_TYPE {
		a_spec_typesTo.SPEC_OBJECT_TYPE = append(a_spec_typesTo.SPEC_OBJECT_TYPE, CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, _spec_object_type))
	}
	for _, _spec_relation_type := range a_spec_typesFrom.SPEC_RELATION_TYPE {
		a_spec_typesTo.SPEC_RELATION_TYPE = append(a_spec_typesTo.SPEC_RELATION_TYPE, CopyBranchSPEC_RELATION_TYPE(mapOrigCopy, _spec_relation_type))
	}
	for _, _specification_type := range a_spec_typesFrom.SPECIFICATION_TYPE {
		a_spec_typesTo.SPECIFICATION_TYPE = append(a_spec_typesTo.SPECIFICATION_TYPE, CopyBranchSPECIFICATION_TYPE(mapOrigCopy, _specification_type))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_header := range a_the_headerFrom.REQ_IF_HEADER {
		a_the_headerTo.REQ_IF_HEADER = append(a_the_headerTo.REQ_IF_HEADER, CopyBranchREQ_IF_HEADER(mapOrigCopy, _req_if_header))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensionsFrom.REQ_IF_TOOL_EXTENSION {
		a_tool_extensionsTo.REQ_IF_TOOL_EXTENSION = append(a_tool_extensionsTo.REQ_IF_TOOL_EXTENSION, CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, _req_if_tool_extension))
	}

	return
}

func CopyBranchA_TYPE(mapOrigCopy map[any]any, a_typeFrom *A_TYPE) (a_typeTo *A_TYPE) {

	// a_typeFrom has already been copied
	if _a_typeTo, ok := mapOrigCopy[a_typeFrom]; ok {
		a_typeTo = _a_typeTo.(*A_TYPE)
		return
	}

	a_typeTo = new(A_TYPE)
	mapOrigCopy[a_typeFrom] = a_typeTo
	a_typeFrom.CopyBasicFields(a_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_1(mapOrigCopy map[any]any, a_type_1From *A_TYPE_1) (a_type_1To *A_TYPE_1) {

	// a_type_1From has already been copied
	if _a_type_1To, ok := mapOrigCopy[a_type_1From]; ok {
		a_type_1To = _a_type_1To.(*A_TYPE_1)
		return
	}

	a_type_1To = new(A_TYPE_1)
	mapOrigCopy[a_type_1From] = a_type_1To
	a_type_1From.CopyBasicFields(a_type_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_10(mapOrigCopy map[any]any, a_type_10From *A_TYPE_10) (a_type_10To *A_TYPE_10) {

	// a_type_10From has already been copied
	if _a_type_10To, ok := mapOrigCopy[a_type_10From]; ok {
		a_type_10To = _a_type_10To.(*A_TYPE_10)
		return
	}

	a_type_10To = new(A_TYPE_10)
	mapOrigCopy[a_type_10From] = a_type_10To
	a_type_10From.CopyBasicFields(a_type_10To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_2(mapOrigCopy map[any]any, a_type_2From *A_TYPE_2) (a_type_2To *A_TYPE_2) {

	// a_type_2From has already been copied
	if _a_type_2To, ok := mapOrigCopy[a_type_2From]; ok {
		a_type_2To = _a_type_2To.(*A_TYPE_2)
		return
	}

	a_type_2To = new(A_TYPE_2)
	mapOrigCopy[a_type_2From] = a_type_2To
	a_type_2From.CopyBasicFields(a_type_2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_3(mapOrigCopy map[any]any, a_type_3From *A_TYPE_3) (a_type_3To *A_TYPE_3) {

	// a_type_3From has already been copied
	if _a_type_3To, ok := mapOrigCopy[a_type_3From]; ok {
		a_type_3To = _a_type_3To.(*A_TYPE_3)
		return
	}

	a_type_3To = new(A_TYPE_3)
	mapOrigCopy[a_type_3From] = a_type_3To
	a_type_3From.CopyBasicFields(a_type_3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_4(mapOrigCopy map[any]any, a_type_4From *A_TYPE_4) (a_type_4To *A_TYPE_4) {

	// a_type_4From has already been copied
	if _a_type_4To, ok := mapOrigCopy[a_type_4From]; ok {
		a_type_4To = _a_type_4To.(*A_TYPE_4)
		return
	}

	a_type_4To = new(A_TYPE_4)
	mapOrigCopy[a_type_4From] = a_type_4To
	a_type_4From.CopyBasicFields(a_type_4To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_5(mapOrigCopy map[any]any, a_type_5From *A_TYPE_5) (a_type_5To *A_TYPE_5) {

	// a_type_5From has already been copied
	if _a_type_5To, ok := mapOrigCopy[a_type_5From]; ok {
		a_type_5To = _a_type_5To.(*A_TYPE_5)
		return
	}

	a_type_5To = new(A_TYPE_5)
	mapOrigCopy[a_type_5From] = a_type_5To
	a_type_5From.CopyBasicFields(a_type_5To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_6(mapOrigCopy map[any]any, a_type_6From *A_TYPE_6) (a_type_6To *A_TYPE_6) {

	// a_type_6From has already been copied
	if _a_type_6To, ok := mapOrigCopy[a_type_6From]; ok {
		a_type_6To = _a_type_6To.(*A_TYPE_6)
		return
	}

	a_type_6To = new(A_TYPE_6)
	mapOrigCopy[a_type_6From] = a_type_6To
	a_type_6From.CopyBasicFields(a_type_6To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_7(mapOrigCopy map[any]any, a_type_7From *A_TYPE_7) (a_type_7To *A_TYPE_7) {

	// a_type_7From has already been copied
	if _a_type_7To, ok := mapOrigCopy[a_type_7From]; ok {
		a_type_7To = _a_type_7To.(*A_TYPE_7)
		return
	}

	a_type_7To = new(A_TYPE_7)
	mapOrigCopy[a_type_7From] = a_type_7To
	a_type_7From.CopyBasicFields(a_type_7To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_8(mapOrigCopy map[any]any, a_type_8From *A_TYPE_8) (a_type_8To *A_TYPE_8) {

	// a_type_8From has already been copied
	if _a_type_8To, ok := mapOrigCopy[a_type_8From]; ok {
		a_type_8To = _a_type_8To.(*A_TYPE_8)
		return
	}

	a_type_8To = new(A_TYPE_8)
	mapOrigCopy[a_type_8From] = a_type_8To
	a_type_8From.CopyBasicFields(a_type_8To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_TYPE_9(mapOrigCopy map[any]any, a_type_9From *A_TYPE_9) (a_type_9To *A_TYPE_9) {

	// a_type_9From has already been copied
	if _a_type_9To, ok := mapOrigCopy[a_type_9From]; ok {
		a_type_9To = _a_type_9To.(*A_TYPE_9)
		return
	}

	a_type_9To = new(A_TYPE_9)
	mapOrigCopy[a_type_9From] = a_type_9To
	a_type_9From.CopyBasicFields(a_type_9To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_VALUES(mapOrigCopy map[any]any, a_valuesFrom *A_VALUES) (a_valuesTo *A_VALUES) {

	// a_valuesFrom has already been copied
	if _a_valuesTo, ok := mapOrigCopy[a_valuesFrom]; ok {
		a_valuesTo = _a_valuesTo.(*A_VALUES)
		return
	}

	a_valuesTo = new(A_VALUES)
	mapOrigCopy[a_valuesFrom] = a_valuesTo
	a_valuesFrom.CopyBasicFields(a_valuesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchA_VALUES_1(mapOrigCopy map[any]any, a_values_1From *A_VALUES_1) (a_values_1To *A_VALUES_1) {

	// a_values_1From has already been copied
	if _a_values_1To, ok := mapOrigCopy[a_values_1From]; ok {
		a_values_1To = _a_values_1To.(*A_VALUES_1)
		return
	}

	a_values_1To = new(A_VALUES_1)
	mapOrigCopy[a_values_1From] = a_values_1To
	a_values_1From.CopyBasicFields(a_values_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_values_1From.ATTRIBUTE_VALUE_BOOLEAN {
		a_values_1To.ATTRIBUTE_VALUE_BOOLEAN = append(a_values_1To.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}
	for _, _attribute_value_date := range a_values_1From.ATTRIBUTE_VALUE_DATE {
		a_values_1To.ATTRIBUTE_VALUE_DATE = append(a_values_1To.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}
	for _, _attribute_value_enumeration := range a_values_1From.ATTRIBUTE_VALUE_ENUMERATION {
		a_values_1To.ATTRIBUTE_VALUE_ENUMERATION = append(a_values_1To.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _attribute_value_integer := range a_values_1From.ATTRIBUTE_VALUE_INTEGER {
		a_values_1To.ATTRIBUTE_VALUE_INTEGER = append(a_values_1To.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}
	for _, _attribute_value_real := range a_values_1From.ATTRIBUTE_VALUE_REAL {
		a_values_1To.ATTRIBUTE_VALUE_REAL = append(a_values_1To.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}
	for _, _attribute_value_string := range a_values_1From.ATTRIBUTE_VALUE_STRING {
		a_values_1To.ATTRIBUTE_VALUE_STRING = append(a_values_1To.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}
	for _, _attribute_value_xhtml := range a_values_1From.ATTRIBUTE_VALUE_XHTML {
		a_values_1To.ATTRIBUTE_VALUE_XHTML = append(a_values_1To.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_booleanFrom.ALTERNATIVE_ID {
		datatype_definition_booleanTo.ALTERNATIVE_ID = append(datatype_definition_booleanTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_dateFrom.ALTERNATIVE_ID {
		datatype_definition_dateTo.ALTERNATIVE_ID = append(datatype_definition_dateTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_enumerationFrom.ALTERNATIVE_ID {
		datatype_definition_enumerationTo.ALTERNATIVE_ID = append(datatype_definition_enumerationTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_specified_values := range datatype_definition_enumerationFrom.SPECIFIED_VALUES {
		datatype_definition_enumerationTo.SPECIFIED_VALUES = append(datatype_definition_enumerationTo.SPECIFIED_VALUES, CopyBranchA_SPECIFIED_VALUES(mapOrigCopy, _a_specified_values))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_integerFrom.ALTERNATIVE_ID {
		datatype_definition_integerTo.ALTERNATIVE_ID = append(datatype_definition_integerTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_realFrom.ALTERNATIVE_ID {
		datatype_definition_realTo.ALTERNATIVE_ID = append(datatype_definition_realTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_stringFrom.ALTERNATIVE_ID {
		datatype_definition_stringTo.ALTERNATIVE_ID = append(datatype_definition_stringTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_xhtmlFrom.ALTERNATIVE_ID {
		datatype_definition_xhtmlTo.ALTERNATIVE_ID = append(datatype_definition_xhtmlTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range enum_valueFrom.ALTERNATIVE_ID {
		enum_valueTo.ALTERNATIVE_ID = append(enum_valueTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_properties := range enum_valueFrom.PROPERTIES {
		enum_valueTo.PROPERTIES = append(enum_valueTo.PROPERTIES, CopyBranchA_PROPERTIES(mapOrigCopy, _a_properties))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_groupFrom.ALTERNATIVE_ID {
		relation_groupTo.ALTERNATIVE_ID = append(relation_groupTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_source_specification := range relation_groupFrom.SOURCE_SPECIFICATION {
		relation_groupTo.SOURCE_SPECIFICATION = append(relation_groupTo.SOURCE_SPECIFICATION, CopyBranchA_SOURCE_SPECIFICATION(mapOrigCopy, _a_source_specification))
	}
	for _, _a_spec_relations := range relation_groupFrom.SPEC_RELATIONS {
		relation_groupTo.SPEC_RELATIONS = append(relation_groupTo.SPEC_RELATIONS, CopyBranchA_SPEC_RELATIONS(mapOrigCopy, _a_spec_relations))
	}
	for _, _a_type_1 := range relation_groupFrom.TYPE {
		relation_groupTo.TYPE = append(relation_groupTo.TYPE, CopyBranchA_TYPE_1(mapOrigCopy, _a_type_1))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_group_typeFrom.ALTERNATIVE_ID {
		relation_group_typeTo.ALTERNATIVE_ID = append(relation_group_typeTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_spec_attributes := range relation_group_typeFrom.SPEC_ATTRIBUTES {
		relation_group_typeTo.SPEC_ATTRIBUTES = append(relation_group_typeTo.SPEC_ATTRIBUTES, CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, _a_spec_attributes))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_the_header := range req_ifFrom.THE_HEADER {
		req_ifTo.THE_HEADER = append(req_ifTo.THE_HEADER, CopyBranchA_THE_HEADER(mapOrigCopy, _a_the_header))
	}
	for _, _a_core_content := range req_ifFrom.CORE_CONTENT {
		req_ifTo.CORE_CONTENT = append(req_ifTo.CORE_CONTENT, CopyBranchA_CORE_CONTENT(mapOrigCopy, _a_core_content))
	}
	for _, _a_tool_extensions := range req_ifFrom.TOOL_EXTENSIONS {
		req_ifTo.TOOL_EXTENSIONS = append(req_ifTo.TOOL_EXTENSIONS, CopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, _a_tool_extensions))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_datatypes := range req_if_contentFrom.DATATYPES {
		req_if_contentTo.DATATYPES = append(req_if_contentTo.DATATYPES, CopyBranchA_DATATYPES(mapOrigCopy, _a_datatypes))
	}
	for _, _a_spec_types := range req_if_contentFrom.SPEC_TYPES {
		req_if_contentTo.SPEC_TYPES = append(req_if_contentTo.SPEC_TYPES, CopyBranchA_SPEC_TYPES(mapOrigCopy, _a_spec_types))
	}
	for _, _a_spec_objects := range req_if_contentFrom.SPEC_OBJECTS {
		req_if_contentTo.SPEC_OBJECTS = append(req_if_contentTo.SPEC_OBJECTS, CopyBranchA_SPEC_OBJECTS(mapOrigCopy, _a_spec_objects))
	}
	for _, _a_spec_relations_1 := range req_if_contentFrom.SPEC_RELATIONS {
		req_if_contentTo.SPEC_RELATIONS = append(req_if_contentTo.SPEC_RELATIONS, CopyBranchA_SPEC_RELATIONS_1(mapOrigCopy, _a_spec_relations_1))
	}
	for _, _a_specifications := range req_if_contentFrom.SPECIFICATIONS {
		req_if_contentTo.SPECIFICATIONS = append(req_if_contentTo.SPECIFICATIONS, CopyBranchA_SPECIFICATIONS(mapOrigCopy, _a_specifications))
	}
	for _, _a_spec_relation_groups := range req_if_contentFrom.SPEC_RELATION_GROUPS {
		req_if_contentTo.SPEC_RELATION_GROUPS = append(req_if_contentTo.SPEC_RELATION_GROUPS, CopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, _a_spec_relation_groups))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specificationFrom.ALTERNATIVE_ID {
		specificationTo.ALTERNATIVE_ID = append(specificationTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_children := range specificationFrom.CHILDREN {
		specificationTo.CHILDREN = append(specificationTo.CHILDREN, CopyBranchA_CHILDREN(mapOrigCopy, _a_children))
	}
	for _, _a_values_1 := range specificationFrom.VALUES {
		specificationTo.VALUES = append(specificationTo.VALUES, CopyBranchA_VALUES_1(mapOrigCopy, _a_values_1))
	}
	for _, _a_type_10 := range specificationFrom.TYPE {
		specificationTo.TYPE = append(specificationTo.TYPE, CopyBranchA_TYPE_10(mapOrigCopy, _a_type_10))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specification_typeFrom.ALTERNATIVE_ID {
		specification_typeTo.ALTERNATIVE_ID = append(specification_typeTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_spec_attributes := range specification_typeFrom.SPEC_ATTRIBUTES {
		specification_typeTo.SPEC_ATTRIBUTES = append(specification_typeTo.SPEC_ATTRIBUTES, CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, _a_spec_attributes))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_hierarchyFrom.ALTERNATIVE_ID {
		spec_hierarchyTo.ALTERNATIVE_ID = append(spec_hierarchyTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_children := range spec_hierarchyFrom.CHILDREN {
		spec_hierarchyTo.CHILDREN = append(spec_hierarchyTo.CHILDREN, CopyBranchA_CHILDREN(mapOrigCopy, _a_children))
	}
	for _, _a_editable_atts := range spec_hierarchyFrom.EDITABLE_ATTS {
		spec_hierarchyTo.EDITABLE_ATTS = append(spec_hierarchyTo.EDITABLE_ATTS, CopyBranchA_EDITABLE_ATTS(mapOrigCopy, _a_editable_atts))
	}
	for _, _a_object := range spec_hierarchyFrom.OBJECT {
		spec_hierarchyTo.OBJECT = append(spec_hierarchyTo.OBJECT, CopyBranchA_OBJECT(mapOrigCopy, _a_object))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_objectFrom.ALTERNATIVE_ID {
		spec_objectTo.ALTERNATIVE_ID = append(spec_objectTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_values_1 := range spec_objectFrom.VALUES {
		spec_objectTo.VALUES = append(spec_objectTo.VALUES, CopyBranchA_VALUES_1(mapOrigCopy, _a_values_1))
	}
	for _, _a_type_2 := range spec_objectFrom.TYPE {
		spec_objectTo.TYPE = append(spec_objectTo.TYPE, CopyBranchA_TYPE_2(mapOrigCopy, _a_type_2))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_object_typeFrom.ALTERNATIVE_ID {
		spec_object_typeTo.ALTERNATIVE_ID = append(spec_object_typeTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_spec_attributes := range spec_object_typeFrom.SPEC_ATTRIBUTES {
		spec_object_typeTo.SPEC_ATTRIBUTES = append(spec_object_typeTo.SPEC_ATTRIBUTES, CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, _a_spec_attributes))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relationFrom.ALTERNATIVE_ID {
		spec_relationTo.ALTERNATIVE_ID = append(spec_relationTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_values_1 := range spec_relationFrom.VALUES {
		spec_relationTo.VALUES = append(spec_relationTo.VALUES, CopyBranchA_VALUES_1(mapOrigCopy, _a_values_1))
	}
	for _, _a_source := range spec_relationFrom.SOURCE {
		spec_relationTo.SOURCE = append(spec_relationTo.SOURCE, CopyBranchA_SOURCE(mapOrigCopy, _a_source))
	}
	for _, _a_type_4 := range spec_relationFrom.TYPE {
		spec_relationTo.TYPE = append(spec_relationTo.TYPE, CopyBranchA_TYPE_4(mapOrigCopy, _a_type_4))
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relation_typeFrom.ALTERNATIVE_ID {
		spec_relation_typeTo.ALTERNATIVE_ID = append(spec_relation_typeTo.ALTERNATIVE_ID, CopyBranchA_ALTERNATIVE_ID(mapOrigCopy, _a_alternative_id))
	}
	for _, _a_spec_attributes := range spec_relation_typeFrom.SPEC_ATTRIBUTES {
		spec_relation_typeTo.SPEC_ATTRIBUTES = append(spec_relation_typeTo.SPEC_ATTRIBUTES, CopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, _a_spec_attributes))
	}

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

	case *A_CHILDREN:
		stage.UnstageBranchA_CHILDREN(target)

	case *A_CORE_CONTENT:
		stage.UnstageBranchA_CORE_CONTENT(target)

	case *A_DATATYPES:
		stage.UnstageBranchA_DATATYPES(target)

	case *A_DEFAULT_VALUE:
		stage.UnstageBranchA_DEFAULT_VALUE(target)

	case *A_DEFAULT_VALUE_1:
		stage.UnstageBranchA_DEFAULT_VALUE_1(target)

	case *A_DEFAULT_VALUE_2:
		stage.UnstageBranchA_DEFAULT_VALUE_2(target)

	case *A_DEFAULT_VALUE_3:
		stage.UnstageBranchA_DEFAULT_VALUE_3(target)

	case *A_DEFAULT_VALUE_4:
		stage.UnstageBranchA_DEFAULT_VALUE_4(target)

	case *A_DEFAULT_VALUE_5:
		stage.UnstageBranchA_DEFAULT_VALUE_5(target)

	case *A_DEFAULT_VALUE_6:
		stage.UnstageBranchA_DEFAULT_VALUE_6(target)

	case *A_DEFINITION:
		stage.UnstageBranchA_DEFINITION(target)

	case *A_DEFINITION_1:
		stage.UnstageBranchA_DEFINITION_1(target)

	case *A_DEFINITION_2:
		stage.UnstageBranchA_DEFINITION_2(target)

	case *A_DEFINITION_3:
		stage.UnstageBranchA_DEFINITION_3(target)

	case *A_DEFINITION_4:
		stage.UnstageBranchA_DEFINITION_4(target)

	case *A_DEFINITION_5:
		stage.UnstageBranchA_DEFINITION_5(target)

	case *A_DEFINITION_6:
		stage.UnstageBranchA_DEFINITION_6(target)

	case *A_EDITABLE_ATTS:
		stage.UnstageBranchA_EDITABLE_ATTS(target)

	case *A_OBJECT:
		stage.UnstageBranchA_OBJECT(target)

	case *A_PROPERTIES:
		stage.UnstageBranchA_PROPERTIES(target)

	case *A_SOURCE:
		stage.UnstageBranchA_SOURCE(target)

	case *A_SOURCE_SPECIFICATION:
		stage.UnstageBranchA_SOURCE_SPECIFICATION(target)

	case *A_SPECIFICATIONS:
		stage.UnstageBranchA_SPECIFICATIONS(target)

	case *A_SPECIFIED_VALUES:
		stage.UnstageBranchA_SPECIFIED_VALUES(target)

	case *A_SPEC_ATTRIBUTES:
		stage.UnstageBranchA_SPEC_ATTRIBUTES(target)

	case *A_SPEC_OBJECTS:
		stage.UnstageBranchA_SPEC_OBJECTS(target)

	case *A_SPEC_RELATIONS:
		stage.UnstageBranchA_SPEC_RELATIONS(target)

	case *A_SPEC_RELATIONS_1:
		stage.UnstageBranchA_SPEC_RELATIONS_1(target)

	case *A_SPEC_RELATION_GROUPS:
		stage.UnstageBranchA_SPEC_RELATION_GROUPS(target)

	case *A_SPEC_TYPES:
		stage.UnstageBranchA_SPEC_TYPES(target)

	case *A_THE_HEADER:
		stage.UnstageBranchA_THE_HEADER(target)

	case *A_TOOL_EXTENSIONS:
		stage.UnstageBranchA_TOOL_EXTENSIONS(target)

	case *A_TYPE:
		stage.UnstageBranchA_TYPE(target)

	case *A_TYPE_1:
		stage.UnstageBranchA_TYPE_1(target)

	case *A_TYPE_10:
		stage.UnstageBranchA_TYPE_10(target)

	case *A_TYPE_2:
		stage.UnstageBranchA_TYPE_2(target)

	case *A_TYPE_3:
		stage.UnstageBranchA_TYPE_3(target)

	case *A_TYPE_4:
		stage.UnstageBranchA_TYPE_4(target)

	case *A_TYPE_5:
		stage.UnstageBranchA_TYPE_5(target)

	case *A_TYPE_6:
		stage.UnstageBranchA_TYPE_6(target)

	case *A_TYPE_7:
		stage.UnstageBranchA_TYPE_7(target)

	case *A_TYPE_8:
		stage.UnstageBranchA_TYPE_8(target)

	case *A_TYPE_9:
		stage.UnstageBranchA_TYPE_9(target)

	case *A_VALUES:
		stage.UnstageBranchA_VALUES(target)

	case *A_VALUES_1:
		stage.UnstageBranchA_VALUES_1(target)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_1 := range attribute_definition_boolean.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_1)
	}
	for _, _a_type_7 := range attribute_definition_boolean.TYPE {
		UnstageBranch(stage, _a_type_7)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_date.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_3 := range attribute_definition_date.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_3)
	}
	for _, _a_type := range attribute_definition_date.TYPE {
		UnstageBranch(stage, _a_type)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_5 := range attribute_definition_enumeration.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_5)
	}
	for _, _a_type_9 := range attribute_definition_enumeration.TYPE {
		UnstageBranch(stage, _a_type_9)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_integer.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_6 := range attribute_definition_integer.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_6)
	}
	for _, _a_type_5 := range attribute_definition_integer.TYPE {
		UnstageBranch(stage, _a_type_5)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_real.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_4 := range attribute_definition_real.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_4)
	}
	for _, _a_type_3 := range attribute_definition_real.TYPE {
		UnstageBranch(stage, _a_type_3)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_string.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value := range attribute_definition_string.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value)
	}
	for _, _a_type_8 := range attribute_definition_string.TYPE {
		UnstageBranch(stage, _a_type_8)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_default_value_2 := range attribute_definition_xhtml.DEFAULT_VALUE {
		UnstageBranch(stage, _a_default_value_2)
	}
	for _, _a_type_6 := range attribute_definition_xhtml.TYPE {
		UnstageBranch(stage, _a_type_6)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_4 := range attribute_value_boolean.DEFINITION {
		UnstageBranch(stage, _a_definition_4)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_2 := range attribute_value_date.DEFINITION {
		UnstageBranch(stage, _a_definition_2)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_5 := range attribute_value_enumeration.DEFINITION {
		UnstageBranch(stage, _a_definition_5)
	}
	for _, _a_values := range attribute_value_enumeration.VALUES {
		UnstageBranch(stage, _a_values)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition := range attribute_value_integer.DEFINITION {
		UnstageBranch(stage, _a_definition)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_6 := range attribute_value_real.DEFINITION {
		UnstageBranch(stage, _a_definition_6)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_definition_3 := range attribute_value_string.DEFINITION {
		UnstageBranch(stage, _a_definition_3)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
		UnstageBranch(stage, _xhtml_content)
	}
	for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
		UnstageBranch(stage, _xhtml_content)
	}
	for _, _a_definition_1 := range attribute_value_xhtml.DEFINITION {
		UnstageBranch(stage, _a_definition_1)
	}

}

func (stage *StageStruct) UnstageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if !IsStaged(stage, a_alternative_id) {
		return
	}

	a_alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range a_alternative_id.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if !IsStaged(stage, a_children) {
		return
	}

	a_children.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
		UnstageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) UnstageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, a_core_content) {
		return
	}

	a_core_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_content := range a_core_content.REQ_IF_CONTENT {
		UnstageBranch(stage, _req_if_content)
	}

}

func (stage *StageStruct) UnstageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if !IsStaged(stage, a_datatypes) {
		return
	}

	a_datatypes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
		UnstageBranch(stage, _datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
		UnstageBranch(stage, _datatype_definition_integer)
	}
	for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
		UnstageBranch(stage, _datatype_definition_real)
	}
	for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
		UnstageBranch(stage, _datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
		UnstageBranch(stage, _datatype_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE(a_default_value *A_DEFAULT_VALUE) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value) {
		return
	}

	a_default_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_default_value.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_1(a_default_value_1 *A_DEFAULT_VALUE_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_1) {
		return
	}

	a_default_value_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_default_value_1.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_2(a_default_value_2 *A_DEFAULT_VALUE_2) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_2) {
		return
	}

	a_default_value_2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_default_value_2.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_3(a_default_value_3 *A_DEFAULT_VALUE_3) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_3) {
		return
	}

	a_default_value_3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_default_value_3.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_4(a_default_value_4 *A_DEFAULT_VALUE_4) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_4) {
		return
	}

	a_default_value_4.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_default_value_4.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_5(a_default_value_5 *A_DEFAULT_VALUE_5) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_5) {
		return
	}

	a_default_value_5.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_default_value_5.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFAULT_VALUE_6(a_default_value_6 *A_DEFAULT_VALUE_6) {

	// check if instance is already staged
	if !IsStaged(stage, a_default_value_6) {
		return
	}

	a_default_value_6.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_default_value_6.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}

}

func (stage *StageStruct) UnstageBranchA_DEFINITION(a_definition *A_DEFINITION) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition) {
		return
	}

	a_definition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_1(a_definition_1 *A_DEFINITION_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_1) {
		return
	}

	a_definition_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_2(a_definition_2 *A_DEFINITION_2) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_2) {
		return
	}

	a_definition_2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_3(a_definition_3 *A_DEFINITION_3) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_3) {
		return
	}

	a_definition_3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_4(a_definition_4 *A_DEFINITION_4) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_4) {
		return
	}

	a_definition_4.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_5(a_definition_5 *A_DEFINITION_5) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_5) {
		return
	}

	a_definition_5.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_DEFINITION_6(a_definition_6 *A_DEFINITION_6) {

	// check if instance is already staged
	if !IsStaged(stage, a_definition_6) {
		return
	}

	a_definition_6.Unstage(stage)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _embedded_value := range a_properties.EMBEDDED_VALUE {
		UnstageBranch(stage, _embedded_value)
	}

}

func (stage *StageStruct) UnstageBranchA_SOURCE(a_source *A_SOURCE) {

	// check if instance is already staged
	if !IsStaged(stage, a_source) {
		return
	}

	a_source.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SOURCE_SPECIFICATION(a_source_specification *A_SOURCE_SPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, a_source_specification) {
		return
	}

	a_source_specification.Unstage(stage)

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specifications.SPECIFICATION {
		UnstageBranch(stage, _specification)
	}

}

func (stage *StageStruct) UnstageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if !IsStaged(stage, a_specified_values) {
		return
	}

	a_specified_values.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_values.ENUM_VALUE {
		UnstageBranch(stage, _enum_value)
	}

}

func (stage *StageStruct) UnstageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_attributes) {
		return
	}

	a_spec_attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
		UnstageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
		UnstageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
		UnstageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
		UnstageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
		UnstageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_objects) {
		return
	}

	a_spec_objects.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
		UnstageBranch(stage, _spec_object)
	}

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relations) {
		return
	}

	a_spec_relations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATIONS_1(a_spec_relations_1 *A_SPEC_RELATIONS_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relations_1) {
		return
	}

	a_spec_relations_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relations_1.SPEC_RELATION {
		UnstageBranch(stage, _spec_relation)
	}

}

func (stage *StageStruct) UnstageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
		UnstageBranch(stage, _relation_group)
	}

}

func (stage *StageStruct) UnstageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if !IsStaged(stage, a_spec_types) {
		return
	}

	a_spec_types.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
		UnstageBranch(stage, _relation_group_type)
	}
	for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
		UnstageBranch(stage, _spec_object_type)
	}
	for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
		UnstageBranch(stage, _spec_relation_type)
	}
	for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
		UnstageBranch(stage, _specification_type)
	}

}

func (stage *StageStruct) UnstageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if !IsStaged(stage, a_the_header) {
		return
	}

	a_the_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_header := range a_the_header.REQ_IF_HEADER {
		UnstageBranch(stage, _req_if_header)
	}

}

func (stage *StageStruct) UnstageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if !IsStaged(stage, a_tool_extensions) {
		return
	}

	a_tool_extensions.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
		UnstageBranch(stage, _req_if_tool_extension)
	}

}

func (stage *StageStruct) UnstageBranchA_TYPE(a_type *A_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, a_type) {
		return
	}

	a_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_1(a_type_1 *A_TYPE_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_1) {
		return
	}

	a_type_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_10(a_type_10 *A_TYPE_10) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_10) {
		return
	}

	a_type_10.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_2(a_type_2 *A_TYPE_2) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_2) {
		return
	}

	a_type_2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_3(a_type_3 *A_TYPE_3) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_3) {
		return
	}

	a_type_3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_4(a_type_4 *A_TYPE_4) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_4) {
		return
	}

	a_type_4.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_5(a_type_5 *A_TYPE_5) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_5) {
		return
	}

	a_type_5.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_6(a_type_6 *A_TYPE_6) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_6) {
		return
	}

	a_type_6.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_7(a_type_7 *A_TYPE_7) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_7) {
		return
	}

	a_type_7.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_8(a_type_8 *A_TYPE_8) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_8) {
		return
	}

	a_type_8.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_TYPE_9(a_type_9 *A_TYPE_9) {

	// check if instance is already staged
	if !IsStaged(stage, a_type_9) {
		return
	}

	a_type_9.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_VALUES(a_values *A_VALUES) {

	// check if instance is already staged
	if !IsStaged(stage, a_values) {
		return
	}

	a_values.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchA_VALUES_1(a_values_1 *A_VALUES_1) {

	// check if instance is already staged
	if !IsStaged(stage, a_values_1) {
		return
	}

	a_values_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_values_1.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range a_values_1.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range a_values_1.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range a_values_1.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range a_values_1.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range a_values_1.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range a_values_1.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_date.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_specified_values := range datatype_definition_enumeration.SPECIFIED_VALUES {
		UnstageBranch(stage, _a_specified_values)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_integer.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_real.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_string.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range enum_value.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_properties := range enum_value.PROPERTIES {
		UnstageBranch(stage, _a_properties)
	}

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group) {
		return
	}

	relation_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_group.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_source_specification := range relation_group.SOURCE_SPECIFICATION {
		UnstageBranch(stage, _a_source_specification)
	}
	for _, _a_spec_relations := range relation_group.SPEC_RELATIONS {
		UnstageBranch(stage, _a_spec_relations)
	}
	for _, _a_type_1 := range relation_group.TYPE {
		UnstageBranch(stage, _a_type_1)
	}

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range relation_group_type.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range relation_group_type.SPEC_ATTRIBUTES {
		UnstageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) UnstageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if !IsStaged(stage, req_if) {
		return
	}

	req_if.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_the_header := range req_if.THE_HEADER {
		UnstageBranch(stage, _a_the_header)
	}
	for _, _a_core_content := range req_if.CORE_CONTENT {
		UnstageBranch(stage, _a_core_content)
	}
	for _, _a_tool_extensions := range req_if.TOOL_EXTENSIONS {
		UnstageBranch(stage, _a_tool_extensions)
	}

}

func (stage *StageStruct) UnstageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_datatypes := range req_if_content.DATATYPES {
		UnstageBranch(stage, _a_datatypes)
	}
	for _, _a_spec_types := range req_if_content.SPEC_TYPES {
		UnstageBranch(stage, _a_spec_types)
	}
	for _, _a_spec_objects := range req_if_content.SPEC_OBJECTS {
		UnstageBranch(stage, _a_spec_objects)
	}
	for _, _a_spec_relations_1 := range req_if_content.SPEC_RELATIONS {
		UnstageBranch(stage, _a_spec_relations_1)
	}
	for _, _a_specifications := range req_if_content.SPECIFICATIONS {
		UnstageBranch(stage, _a_specifications)
	}
	for _, _a_spec_relation_groups := range req_if_content.SPEC_RELATION_GROUPS {
		UnstageBranch(stage, _a_spec_relation_groups)
	}

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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specification.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_children := range specification.CHILDREN {
		UnstageBranch(stage, _a_children)
	}
	for _, _a_values_1 := range specification.VALUES {
		UnstageBranch(stage, _a_values_1)
	}
	for _, _a_type_10 := range specification.TYPE {
		UnstageBranch(stage, _a_type_10)
	}

}

func (stage *StageStruct) UnstageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specification_type) {
		return
	}

	specification_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range specification_type.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range specification_type.SPEC_ATTRIBUTES {
		UnstageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if !IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_hierarchy.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_children := range spec_hierarchy.CHILDREN {
		UnstageBranch(stage, _a_children)
	}
	for _, _a_editable_atts := range spec_hierarchy.EDITABLE_ATTS {
		UnstageBranch(stage, _a_editable_atts)
	}
	for _, _a_object := range spec_hierarchy.OBJECT {
		UnstageBranch(stage, _a_object)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object) {
		return
	}

	spec_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_object.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_values_1 := range spec_object.VALUES {
		UnstageBranch(stage, _a_values_1)
	}
	for _, _a_type_2 := range spec_object.TYPE {
		UnstageBranch(stage, _a_type_2)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_object_type.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range spec_object_type.SPEC_ATTRIBUTES {
		UnstageBranch(stage, _a_spec_attributes)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relation.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_values_1 := range spec_relation.VALUES {
		UnstageBranch(stage, _a_values_1)
	}
	for _, _a_source := range spec_relation.SOURCE {
		UnstageBranch(stage, _a_source)
	}
	for _, _a_type_4 := range spec_relation.TYPE {
		UnstageBranch(stage, _a_type_4)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_alternative_id := range spec_relation_type.ALTERNATIVE_ID {
		UnstageBranch(stage, _a_alternative_id)
	}
	for _, _a_spec_attributes := range spec_relation_type.SPEC_ATTRIBUTES {
		UnstageBranch(stage, _a_spec_attributes)
	}

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
