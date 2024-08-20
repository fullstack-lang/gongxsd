// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ALTERNATIVE_IDAPIs []*ALTERNATIVE_IDAPI

	ATTRIBUTE_DEFINITION_BOOLEANAPIs []*ATTRIBUTE_DEFINITION_BOOLEANAPI

	ATTRIBUTE_DEFINITION_DATEAPIs []*ATTRIBUTE_DEFINITION_DATEAPI

	ATTRIBUTE_DEFINITION_ENUMERATIONAPIs []*ATTRIBUTE_DEFINITION_ENUMERATIONAPI

	ATTRIBUTE_DEFINITION_INTEGERAPIs []*ATTRIBUTE_DEFINITION_INTEGERAPI

	ATTRIBUTE_DEFINITION_REALAPIs []*ATTRIBUTE_DEFINITION_REALAPI

	ATTRIBUTE_DEFINITION_STRINGAPIs []*ATTRIBUTE_DEFINITION_STRINGAPI

	ATTRIBUTE_DEFINITION_XHTMLAPIs []*ATTRIBUTE_DEFINITION_XHTMLAPI

	ATTRIBUTE_VALUE_BOOLEANAPIs []*ATTRIBUTE_VALUE_BOOLEANAPI

	ATTRIBUTE_VALUE_DATEAPIs []*ATTRIBUTE_VALUE_DATEAPI

	ATTRIBUTE_VALUE_ENUMERATIONAPIs []*ATTRIBUTE_VALUE_ENUMERATIONAPI

	ATTRIBUTE_VALUE_INTEGERAPIs []*ATTRIBUTE_VALUE_INTEGERAPI

	ATTRIBUTE_VALUE_REALAPIs []*ATTRIBUTE_VALUE_REALAPI

	ATTRIBUTE_VALUE_STRINGAPIs []*ATTRIBUTE_VALUE_STRINGAPI

	ATTRIBUTE_VALUE_XHTMLAPIs []*ATTRIBUTE_VALUE_XHTMLAPI

	DATATYPE_DEFINITION_BOOLEANAPIs []*DATATYPE_DEFINITION_BOOLEANAPI

	DATATYPE_DEFINITION_DATEAPIs []*DATATYPE_DEFINITION_DATEAPI

	DATATYPE_DEFINITION_ENUMERATIONAPIs []*DATATYPE_DEFINITION_ENUMERATIONAPI

	DATATYPE_DEFINITION_INTEGERAPIs []*DATATYPE_DEFINITION_INTEGERAPI

	DATATYPE_DEFINITION_REALAPIs []*DATATYPE_DEFINITION_REALAPI

	DATATYPE_DEFINITION_STRINGAPIs []*DATATYPE_DEFINITION_STRINGAPI

	DATATYPE_DEFINITION_XHTMLAPIs []*DATATYPE_DEFINITION_XHTMLAPI

	EMBEDDED_VALUEAPIs []*EMBEDDED_VALUEAPI

	ENUM_VALUEAPIs []*ENUM_VALUEAPI

	RELATION_GROUPAPIs []*RELATION_GROUPAPI

	RELATION_GROUP_TYPEAPIs []*RELATION_GROUP_TYPEAPI

	REQ_IFAPIs []*REQ_IFAPI

	REQ_IF_CONTENTAPIs []*REQ_IF_CONTENTAPI

	REQ_IF_HEADERAPIs []*REQ_IF_HEADERAPI

	REQ_IF_TOOL_EXTENSIONAPIs []*REQ_IF_TOOL_EXTENSIONAPI

	SPECIFICATIONAPIs []*SPECIFICATIONAPI

	SPECIFICATION_TYPEAPIs []*SPECIFICATION_TYPEAPI

	SPEC_HIERARCHYAPIs []*SPEC_HIERARCHYAPI

	SPEC_OBJECTAPIs []*SPEC_OBJECTAPI

	SPEC_OBJECT_TYPEAPIs []*SPEC_OBJECT_TYPEAPI

	SPEC_RELATIONAPIs []*SPEC_RELATIONAPI

	SPEC_RELATION_TYPEAPIs []*SPEC_RELATION_TYPEAPI

	XHTML_CONTENTAPIs []*XHTML_CONTENTAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, alternative_idDB := range backRepo.BackRepoALTERNATIVE_ID.Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDDB {

		var alternative_idAPI ALTERNATIVE_IDAPI
		alternative_idAPI.ID = alternative_idDB.ID
		alternative_idAPI.ALTERNATIVE_IDPointersEncoding = alternative_idDB.ALTERNATIVE_IDPointersEncoding
		alternative_idDB.CopyBasicFieldsToALTERNATIVE_ID_WOP(&alternative_idAPI.ALTERNATIVE_ID_WOP)

		backRepoData.ALTERNATIVE_IDAPIs = append(backRepoData.ALTERNATIVE_IDAPIs, &alternative_idAPI)
	}

	for _, attribute_definition_booleanDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB {

		var attribute_definition_booleanAPI ATTRIBUTE_DEFINITION_BOOLEANAPI
		attribute_definition_booleanAPI.ID = attribute_definition_booleanDB.ID
		attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
		attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP(&attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_BOOLEANAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_BOOLEANAPIs, &attribute_definition_booleanAPI)
	}

	for _, attribute_definition_dateDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEDB {

		var attribute_definition_dateAPI ATTRIBUTE_DEFINITION_DATEAPI
		attribute_definition_dateAPI.ID = attribute_definition_dateDB.ID
		attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATEPointersEncoding = attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding
		attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE_WOP(&attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATE_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_DATEAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_DATEAPIs, &attribute_definition_dateAPI)
	}

	for _, attribute_definition_enumerationDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONDB {

		var attribute_definition_enumerationAPI ATTRIBUTE_DEFINITION_ENUMERATIONAPI
		attribute_definition_enumerationAPI.ID = attribute_definition_enumerationDB.ID
		attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding = attribute_definition_enumerationDB.ATTRIBUTE_DEFINITION_ENUMERATIONPointersEncoding
		attribute_definition_enumerationDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_ENUMERATION_WOP(&attribute_definition_enumerationAPI.ATTRIBUTE_DEFINITION_ENUMERATION_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_ENUMERATIONAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_ENUMERATIONAPIs, &attribute_definition_enumerationAPI)
	}

	for _, attribute_definition_integerDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB {

		var attribute_definition_integerAPI ATTRIBUTE_DEFINITION_INTEGERAPI
		attribute_definition_integerAPI.ID = attribute_definition_integerDB.ID
		attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
		attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP(&attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGER_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_INTEGERAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_INTEGERAPIs, &attribute_definition_integerAPI)
	}

	for _, attribute_definition_realDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALDB {

		var attribute_definition_realAPI ATTRIBUTE_DEFINITION_REALAPI
		attribute_definition_realAPI.ID = attribute_definition_realDB.ID
		attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REALPointersEncoding = attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding
		attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL_WOP(&attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REAL_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_REALAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_REALAPIs, &attribute_definition_realAPI)
	}

	for _, attribute_definition_stringDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB {

		var attribute_definition_stringAPI ATTRIBUTE_DEFINITION_STRINGAPI
		attribute_definition_stringAPI.ID = attribute_definition_stringDB.ID
		attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding
		attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP(&attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRING_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_STRINGAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_STRINGAPIs, &attribute_definition_stringAPI)
	}

	for _, attribute_definition_xhtmlDB := range backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLDB {

		var attribute_definition_xhtmlAPI ATTRIBUTE_DEFINITION_XHTMLAPI
		attribute_definition_xhtmlAPI.ID = attribute_definition_xhtmlDB.ID
		attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding
		attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML_WOP(&attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTML_WOP)

		backRepoData.ATTRIBUTE_DEFINITION_XHTMLAPIs = append(backRepoData.ATTRIBUTE_DEFINITION_XHTMLAPIs, &attribute_definition_xhtmlAPI)
	}

	for _, attribute_value_booleanDB := range backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANDB {

		var attribute_value_booleanAPI ATTRIBUTE_VALUE_BOOLEANAPI
		attribute_value_booleanAPI.ID = attribute_value_booleanDB.ID
		attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding
		attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN_WOP(&attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEAN_WOP)

		backRepoData.ATTRIBUTE_VALUE_BOOLEANAPIs = append(backRepoData.ATTRIBUTE_VALUE_BOOLEANAPIs, &attribute_value_booleanAPI)
	}

	for _, attribute_value_dateDB := range backRepo.BackRepoATTRIBUTE_VALUE_DATE.Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEDB {

		var attribute_value_dateAPI ATTRIBUTE_VALUE_DATEAPI
		attribute_value_dateAPI.ID = attribute_value_dateDB.ID
		attribute_value_dateAPI.ATTRIBUTE_VALUE_DATEPointersEncoding = attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding
		attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE_WOP(&attribute_value_dateAPI.ATTRIBUTE_VALUE_DATE_WOP)

		backRepoData.ATTRIBUTE_VALUE_DATEAPIs = append(backRepoData.ATTRIBUTE_VALUE_DATEAPIs, &attribute_value_dateAPI)
	}

	for _, attribute_value_enumerationDB := range backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONDB {

		var attribute_value_enumerationAPI ATTRIBUTE_VALUE_ENUMERATIONAPI
		attribute_value_enumerationAPI.ID = attribute_value_enumerationDB.ID
		attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
		attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION_WOP(&attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATION_WOP)

		backRepoData.ATTRIBUTE_VALUE_ENUMERATIONAPIs = append(backRepoData.ATTRIBUTE_VALUE_ENUMERATIONAPIs, &attribute_value_enumerationAPI)
	}

	for _, attribute_value_integerDB := range backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERDB {

		var attribute_value_integerAPI ATTRIBUTE_VALUE_INTEGERAPI
		attribute_value_integerAPI.ID = attribute_value_integerDB.ID
		attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGERPointersEncoding = attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding
		attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER_WOP(&attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGER_WOP)

		backRepoData.ATTRIBUTE_VALUE_INTEGERAPIs = append(backRepoData.ATTRIBUTE_VALUE_INTEGERAPIs, &attribute_value_integerAPI)
	}

	for _, attribute_value_realDB := range backRepo.BackRepoATTRIBUTE_VALUE_REAL.Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALDB {

		var attribute_value_realAPI ATTRIBUTE_VALUE_REALAPI
		attribute_value_realAPI.ID = attribute_value_realDB.ID
		attribute_value_realAPI.ATTRIBUTE_VALUE_REALPointersEncoding = attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding
		attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL_WOP(&attribute_value_realAPI.ATTRIBUTE_VALUE_REAL_WOP)

		backRepoData.ATTRIBUTE_VALUE_REALAPIs = append(backRepoData.ATTRIBUTE_VALUE_REALAPIs, &attribute_value_realAPI)
	}

	for _, attribute_value_stringDB := range backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGDB {

		var attribute_value_stringAPI ATTRIBUTE_VALUE_STRINGAPI
		attribute_value_stringAPI.ID = attribute_value_stringDB.ID
		attribute_value_stringAPI.ATTRIBUTE_VALUE_STRINGPointersEncoding = attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding
		attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING_WOP(&attribute_value_stringAPI.ATTRIBUTE_VALUE_STRING_WOP)

		backRepoData.ATTRIBUTE_VALUE_STRINGAPIs = append(backRepoData.ATTRIBUTE_VALUE_STRINGAPIs, &attribute_value_stringAPI)
	}

	for _, attribute_value_xhtmlDB := range backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB {

		var attribute_value_xhtmlAPI ATTRIBUTE_VALUE_XHTMLAPI
		attribute_value_xhtmlAPI.ID = attribute_value_xhtmlDB.ID
		attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTMLPointersEncoding = attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding
		attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP(&attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTML_WOP)

		backRepoData.ATTRIBUTE_VALUE_XHTMLAPIs = append(backRepoData.ATTRIBUTE_VALUE_XHTMLAPIs, &attribute_value_xhtmlAPI)
	}

	for _, datatype_definition_booleanDB := range backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANDB {

		var datatype_definition_booleanAPI DATATYPE_DEFINITION_BOOLEANAPI
		datatype_definition_booleanAPI.ID = datatype_definition_booleanDB.ID
		datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEANPointersEncoding = datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding
		datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN_WOP(&datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEAN_WOP)

		backRepoData.DATATYPE_DEFINITION_BOOLEANAPIs = append(backRepoData.DATATYPE_DEFINITION_BOOLEANAPIs, &datatype_definition_booleanAPI)
	}

	for _, datatype_definition_dateDB := range backRepo.BackRepoDATATYPE_DEFINITION_DATE.Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEDB {

		var datatype_definition_dateAPI DATATYPE_DEFINITION_DATEAPI
		datatype_definition_dateAPI.ID = datatype_definition_dateDB.ID
		datatype_definition_dateAPI.DATATYPE_DEFINITION_DATEPointersEncoding = datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding
		datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE_WOP(&datatype_definition_dateAPI.DATATYPE_DEFINITION_DATE_WOP)

		backRepoData.DATATYPE_DEFINITION_DATEAPIs = append(backRepoData.DATATYPE_DEFINITION_DATEAPIs, &datatype_definition_dateAPI)
	}

	for _, datatype_definition_enumerationDB := range backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONDB {

		var datatype_definition_enumerationAPI DATATYPE_DEFINITION_ENUMERATIONAPI
		datatype_definition_enumerationAPI.ID = datatype_definition_enumerationDB.ID
		datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding
		datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION_WOP(&datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATION_WOP)

		backRepoData.DATATYPE_DEFINITION_ENUMERATIONAPIs = append(backRepoData.DATATYPE_DEFINITION_ENUMERATIONAPIs, &datatype_definition_enumerationAPI)
	}

	for _, datatype_definition_integerDB := range backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {

		var datatype_definition_integerAPI DATATYPE_DEFINITION_INTEGERAPI
		datatype_definition_integerAPI.ID = datatype_definition_integerDB.ID
		datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGERPointersEncoding = datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding
		datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP(&datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGER_WOP)

		backRepoData.DATATYPE_DEFINITION_INTEGERAPIs = append(backRepoData.DATATYPE_DEFINITION_INTEGERAPIs, &datatype_definition_integerAPI)
	}

	for _, datatype_definition_realDB := range backRepo.BackRepoDATATYPE_DEFINITION_REAL.Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALDB {

		var datatype_definition_realAPI DATATYPE_DEFINITION_REALAPI
		datatype_definition_realAPI.ID = datatype_definition_realDB.ID
		datatype_definition_realAPI.DATATYPE_DEFINITION_REALPointersEncoding = datatype_definition_realDB.DATATYPE_DEFINITION_REALPointersEncoding
		datatype_definition_realDB.CopyBasicFieldsToDATATYPE_DEFINITION_REAL_WOP(&datatype_definition_realAPI.DATATYPE_DEFINITION_REAL_WOP)

		backRepoData.DATATYPE_DEFINITION_REALAPIs = append(backRepoData.DATATYPE_DEFINITION_REALAPIs, &datatype_definition_realAPI)
	}

	for _, datatype_definition_stringDB := range backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {

		var datatype_definition_stringAPI DATATYPE_DEFINITION_STRINGAPI
		datatype_definition_stringAPI.ID = datatype_definition_stringDB.ID
		datatype_definition_stringAPI.DATATYPE_DEFINITION_STRINGPointersEncoding = datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding
		datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP(&datatype_definition_stringAPI.DATATYPE_DEFINITION_STRING_WOP)

		backRepoData.DATATYPE_DEFINITION_STRINGAPIs = append(backRepoData.DATATYPE_DEFINITION_STRINGAPIs, &datatype_definition_stringAPI)
	}

	for _, datatype_definition_xhtmlDB := range backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLDB {

		var datatype_definition_xhtmlAPI DATATYPE_DEFINITION_XHTMLAPI
		datatype_definition_xhtmlAPI.ID = datatype_definition_xhtmlDB.ID
		datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTMLPointersEncoding = datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding
		datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML_WOP(&datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTML_WOP)

		backRepoData.DATATYPE_DEFINITION_XHTMLAPIs = append(backRepoData.DATATYPE_DEFINITION_XHTMLAPIs, &datatype_definition_xhtmlAPI)
	}

	for _, embedded_valueDB := range backRepo.BackRepoEMBEDDED_VALUE.Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEDB {

		var embedded_valueAPI EMBEDDED_VALUEAPI
		embedded_valueAPI.ID = embedded_valueDB.ID
		embedded_valueAPI.EMBEDDED_VALUEPointersEncoding = embedded_valueDB.EMBEDDED_VALUEPointersEncoding
		embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE_WOP(&embedded_valueAPI.EMBEDDED_VALUE_WOP)

		backRepoData.EMBEDDED_VALUEAPIs = append(backRepoData.EMBEDDED_VALUEAPIs, &embedded_valueAPI)
	}

	for _, enum_valueDB := range backRepo.BackRepoENUM_VALUE.Map_ENUM_VALUEDBID_ENUM_VALUEDB {

		var enum_valueAPI ENUM_VALUEAPI
		enum_valueAPI.ID = enum_valueDB.ID
		enum_valueAPI.ENUM_VALUEPointersEncoding = enum_valueDB.ENUM_VALUEPointersEncoding
		enum_valueDB.CopyBasicFieldsToENUM_VALUE_WOP(&enum_valueAPI.ENUM_VALUE_WOP)

		backRepoData.ENUM_VALUEAPIs = append(backRepoData.ENUM_VALUEAPIs, &enum_valueAPI)
	}

	for _, relation_groupDB := range backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPDB {

		var relation_groupAPI RELATION_GROUPAPI
		relation_groupAPI.ID = relation_groupDB.ID
		relation_groupAPI.RELATION_GROUPPointersEncoding = relation_groupDB.RELATION_GROUPPointersEncoding
		relation_groupDB.CopyBasicFieldsToRELATION_GROUP_WOP(&relation_groupAPI.RELATION_GROUP_WOP)

		backRepoData.RELATION_GROUPAPIs = append(backRepoData.RELATION_GROUPAPIs, &relation_groupAPI)
	}

	for _, relation_group_typeDB := range backRepo.BackRepoRELATION_GROUP_TYPE.Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEDB {

		var relation_group_typeAPI RELATION_GROUP_TYPEAPI
		relation_group_typeAPI.ID = relation_group_typeDB.ID
		relation_group_typeAPI.RELATION_GROUP_TYPEPointersEncoding = relation_group_typeDB.RELATION_GROUP_TYPEPointersEncoding
		relation_group_typeDB.CopyBasicFieldsToRELATION_GROUP_TYPE_WOP(&relation_group_typeAPI.RELATION_GROUP_TYPE_WOP)

		backRepoData.RELATION_GROUP_TYPEAPIs = append(backRepoData.RELATION_GROUP_TYPEAPIs, &relation_group_typeAPI)
	}

	for _, req_ifDB := range backRepo.BackRepoREQ_IF.Map_REQ_IFDBID_REQ_IFDB {

		var req_ifAPI REQ_IFAPI
		req_ifAPI.ID = req_ifDB.ID
		req_ifAPI.REQ_IFPointersEncoding = req_ifDB.REQ_IFPointersEncoding
		req_ifDB.CopyBasicFieldsToREQ_IF_WOP(&req_ifAPI.REQ_IF_WOP)

		backRepoData.REQ_IFAPIs = append(backRepoData.REQ_IFAPIs, &req_ifAPI)
	}

	for _, req_if_contentDB := range backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {

		var req_if_contentAPI REQ_IF_CONTENTAPI
		req_if_contentAPI.ID = req_if_contentDB.ID
		req_if_contentAPI.REQ_IF_CONTENTPointersEncoding = req_if_contentDB.REQ_IF_CONTENTPointersEncoding
		req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT_WOP(&req_if_contentAPI.REQ_IF_CONTENT_WOP)

		backRepoData.REQ_IF_CONTENTAPIs = append(backRepoData.REQ_IF_CONTENTAPIs, &req_if_contentAPI)
	}

	for _, req_if_headerDB := range backRepo.BackRepoREQ_IF_HEADER.Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB {

		var req_if_headerAPI REQ_IF_HEADERAPI
		req_if_headerAPI.ID = req_if_headerDB.ID
		req_if_headerAPI.REQ_IF_HEADERPointersEncoding = req_if_headerDB.REQ_IF_HEADERPointersEncoding
		req_if_headerDB.CopyBasicFieldsToREQ_IF_HEADER_WOP(&req_if_headerAPI.REQ_IF_HEADER_WOP)

		backRepoData.REQ_IF_HEADERAPIs = append(backRepoData.REQ_IF_HEADERAPIs, &req_if_headerAPI)
	}

	for _, req_if_tool_extensionDB := range backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB {

		var req_if_tool_extensionAPI REQ_IF_TOOL_EXTENSIONAPI
		req_if_tool_extensionAPI.ID = req_if_tool_extensionDB.ID
		req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSIONPointersEncoding = req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding
		req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP(&req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSION_WOP)

		backRepoData.REQ_IF_TOOL_EXTENSIONAPIs = append(backRepoData.REQ_IF_TOOL_EXTENSIONAPIs, &req_if_tool_extensionAPI)
	}

	for _, specificationDB := range backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONDB {

		var specificationAPI SPECIFICATIONAPI
		specificationAPI.ID = specificationDB.ID
		specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
		specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)

		backRepoData.SPECIFICATIONAPIs = append(backRepoData.SPECIFICATIONAPIs, &specificationAPI)
	}

	for _, specification_typeDB := range backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEDB {

		var specification_typeAPI SPECIFICATION_TYPEAPI
		specification_typeAPI.ID = specification_typeDB.ID
		specification_typeAPI.SPECIFICATION_TYPEPointersEncoding = specification_typeDB.SPECIFICATION_TYPEPointersEncoding
		specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE_WOP(&specification_typeAPI.SPECIFICATION_TYPE_WOP)

		backRepoData.SPECIFICATION_TYPEAPIs = append(backRepoData.SPECIFICATION_TYPEAPIs, &specification_typeAPI)
	}

	for _, spec_hierarchyDB := range backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB {

		var spec_hierarchyAPI SPEC_HIERARCHYAPI
		spec_hierarchyAPI.ID = spec_hierarchyDB.ID
		spec_hierarchyAPI.SPEC_HIERARCHYPointersEncoding = spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding
		spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY_WOP(&spec_hierarchyAPI.SPEC_HIERARCHY_WOP)

		backRepoData.SPEC_HIERARCHYAPIs = append(backRepoData.SPEC_HIERARCHYAPIs, &spec_hierarchyAPI)
	}

	for _, spec_objectDB := range backRepo.BackRepoSPEC_OBJECT.Map_SPEC_OBJECTDBID_SPEC_OBJECTDB {

		var spec_objectAPI SPEC_OBJECTAPI
		spec_objectAPI.ID = spec_objectDB.ID
		spec_objectAPI.SPEC_OBJECTPointersEncoding = spec_objectDB.SPEC_OBJECTPointersEncoding
		spec_objectDB.CopyBasicFieldsToSPEC_OBJECT_WOP(&spec_objectAPI.SPEC_OBJECT_WOP)

		backRepoData.SPEC_OBJECTAPIs = append(backRepoData.SPEC_OBJECTAPIs, &spec_objectAPI)
	}

	for _, spec_object_typeDB := range backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEDB {

		var spec_object_typeAPI SPEC_OBJECT_TYPEAPI
		spec_object_typeAPI.ID = spec_object_typeDB.ID
		spec_object_typeAPI.SPEC_OBJECT_TYPEPointersEncoding = spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding
		spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE_WOP(&spec_object_typeAPI.SPEC_OBJECT_TYPE_WOP)

		backRepoData.SPEC_OBJECT_TYPEAPIs = append(backRepoData.SPEC_OBJECT_TYPEAPIs, &spec_object_typeAPI)
	}

	for _, spec_relationDB := range backRepo.BackRepoSPEC_RELATION.Map_SPEC_RELATIONDBID_SPEC_RELATIONDB {

		var spec_relationAPI SPEC_RELATIONAPI
		spec_relationAPI.ID = spec_relationDB.ID
		spec_relationAPI.SPEC_RELATIONPointersEncoding = spec_relationDB.SPEC_RELATIONPointersEncoding
		spec_relationDB.CopyBasicFieldsToSPEC_RELATION_WOP(&spec_relationAPI.SPEC_RELATION_WOP)

		backRepoData.SPEC_RELATIONAPIs = append(backRepoData.SPEC_RELATIONAPIs, &spec_relationAPI)
	}

	for _, spec_relation_typeDB := range backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {

		var spec_relation_typeAPI SPEC_RELATION_TYPEAPI
		spec_relation_typeAPI.ID = spec_relation_typeDB.ID
		spec_relation_typeAPI.SPEC_RELATION_TYPEPointersEncoding = spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding
		spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE_WOP(&spec_relation_typeAPI.SPEC_RELATION_TYPE_WOP)

		backRepoData.SPEC_RELATION_TYPEAPIs = append(backRepoData.SPEC_RELATION_TYPEAPIs, &spec_relation_typeAPI)
	}

	for _, xhtml_contentDB := range backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTDB {

		var xhtml_contentAPI XHTML_CONTENTAPI
		xhtml_contentAPI.ID = xhtml_contentDB.ID
		xhtml_contentAPI.XHTML_CONTENTPointersEncoding = xhtml_contentDB.XHTML_CONTENTPointersEncoding
		xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT_WOP(&xhtml_contentAPI.XHTML_CONTENT_WOP)

		backRepoData.XHTML_CONTENTAPIs = append(backRepoData.XHTML_CONTENTAPIs, &xhtml_contentAPI)
	}

}
