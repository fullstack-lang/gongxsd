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

	A_ALTERNATIVE_IDAPIs []*A_ALTERNATIVE_IDAPI

	A_CHILDRENAPIs []*A_CHILDRENAPI

	A_CORE_CONTENTAPIs []*A_CORE_CONTENTAPI

	A_DATATYPESAPIs []*A_DATATYPESAPI

	A_DEFAULT_VALUEAPIs []*A_DEFAULT_VALUEAPI

	A_DEFAULT_VALUE_1APIs []*A_DEFAULT_VALUE_1API

	A_DEFAULT_VALUE_2APIs []*A_DEFAULT_VALUE_2API

	A_DEFAULT_VALUE_3APIs []*A_DEFAULT_VALUE_3API

	A_DEFAULT_VALUE_4APIs []*A_DEFAULT_VALUE_4API

	A_DEFAULT_VALUE_5APIs []*A_DEFAULT_VALUE_5API

	A_DEFAULT_VALUE_6APIs []*A_DEFAULT_VALUE_6API

	A_DEFINITIONAPIs []*A_DEFINITIONAPI

	A_DEFINITION_1APIs []*A_DEFINITION_1API

	A_DEFINITION_2APIs []*A_DEFINITION_2API

	A_DEFINITION_3APIs []*A_DEFINITION_3API

	A_DEFINITION_4APIs []*A_DEFINITION_4API

	A_DEFINITION_5APIs []*A_DEFINITION_5API

	A_DEFINITION_6APIs []*A_DEFINITION_6API

	A_EDITABLE_ATTSAPIs []*A_EDITABLE_ATTSAPI

	A_OBJECTAPIs []*A_OBJECTAPI

	A_PROPERTIESAPIs []*A_PROPERTIESAPI

	A_SOURCEAPIs []*A_SOURCEAPI

	A_SOURCE_SPECIFICATIONAPIs []*A_SOURCE_SPECIFICATIONAPI

	A_SPECIFICATIONSAPIs []*A_SPECIFICATIONSAPI

	A_SPECIFIED_VALUESAPIs []*A_SPECIFIED_VALUESAPI

	A_SPEC_ATTRIBUTESAPIs []*A_SPEC_ATTRIBUTESAPI

	A_SPEC_OBJECTSAPIs []*A_SPEC_OBJECTSAPI

	A_SPEC_RELATIONSAPIs []*A_SPEC_RELATIONSAPI

	A_SPEC_RELATIONS_1APIs []*A_SPEC_RELATIONS_1API

	A_SPEC_RELATION_GROUPSAPIs []*A_SPEC_RELATION_GROUPSAPI

	A_SPEC_TYPESAPIs []*A_SPEC_TYPESAPI

	A_THE_HEADERAPIs []*A_THE_HEADERAPI

	A_TOOL_EXTENSIONSAPIs []*A_TOOL_EXTENSIONSAPI

	A_TYPEAPIs []*A_TYPEAPI

	A_TYPE_1APIs []*A_TYPE_1API

	A_TYPE_10APIs []*A_TYPE_10API

	A_TYPE_2APIs []*A_TYPE_2API

	A_TYPE_3APIs []*A_TYPE_3API

	A_TYPE_4APIs []*A_TYPE_4API

	A_TYPE_5APIs []*A_TYPE_5API

	A_TYPE_6APIs []*A_TYPE_6API

	A_TYPE_7APIs []*A_TYPE_7API

	A_TYPE_8APIs []*A_TYPE_8API

	A_TYPE_9APIs []*A_TYPE_9API

	A_VALUESAPIs []*A_VALUESAPI

	A_VALUES_1APIs []*A_VALUES_1API

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

	for _, a_alternative_idDB := range backRepo.BackRepoA_ALTERNATIVE_ID.Map_A_ALTERNATIVE_IDDBID_A_ALTERNATIVE_IDDB {

		var a_alternative_idAPI A_ALTERNATIVE_IDAPI
		a_alternative_idAPI.ID = a_alternative_idDB.ID
		a_alternative_idAPI.A_ALTERNATIVE_IDPointersEncoding = a_alternative_idDB.A_ALTERNATIVE_IDPointersEncoding
		a_alternative_idDB.CopyBasicFieldsToA_ALTERNATIVE_ID_WOP(&a_alternative_idAPI.A_ALTERNATIVE_ID_WOP)

		backRepoData.A_ALTERNATIVE_IDAPIs = append(backRepoData.A_ALTERNATIVE_IDAPIs, &a_alternative_idAPI)
	}

	for _, a_childrenDB := range backRepo.BackRepoA_CHILDREN.Map_A_CHILDRENDBID_A_CHILDRENDB {

		var a_childrenAPI A_CHILDRENAPI
		a_childrenAPI.ID = a_childrenDB.ID
		a_childrenAPI.A_CHILDRENPointersEncoding = a_childrenDB.A_CHILDRENPointersEncoding
		a_childrenDB.CopyBasicFieldsToA_CHILDREN_WOP(&a_childrenAPI.A_CHILDREN_WOP)

		backRepoData.A_CHILDRENAPIs = append(backRepoData.A_CHILDRENAPIs, &a_childrenAPI)
	}

	for _, a_core_contentDB := range backRepo.BackRepoA_CORE_CONTENT.Map_A_CORE_CONTENTDBID_A_CORE_CONTENTDB {

		var a_core_contentAPI A_CORE_CONTENTAPI
		a_core_contentAPI.ID = a_core_contentDB.ID
		a_core_contentAPI.A_CORE_CONTENTPointersEncoding = a_core_contentDB.A_CORE_CONTENTPointersEncoding
		a_core_contentDB.CopyBasicFieldsToA_CORE_CONTENT_WOP(&a_core_contentAPI.A_CORE_CONTENT_WOP)

		backRepoData.A_CORE_CONTENTAPIs = append(backRepoData.A_CORE_CONTENTAPIs, &a_core_contentAPI)
	}

	for _, a_datatypesDB := range backRepo.BackRepoA_DATATYPES.Map_A_DATATYPESDBID_A_DATATYPESDB {

		var a_datatypesAPI A_DATATYPESAPI
		a_datatypesAPI.ID = a_datatypesDB.ID
		a_datatypesAPI.A_DATATYPESPointersEncoding = a_datatypesDB.A_DATATYPESPointersEncoding
		a_datatypesDB.CopyBasicFieldsToA_DATATYPES_WOP(&a_datatypesAPI.A_DATATYPES_WOP)

		backRepoData.A_DATATYPESAPIs = append(backRepoData.A_DATATYPESAPIs, &a_datatypesAPI)
	}

	for _, a_default_valueDB := range backRepo.BackRepoA_DEFAULT_VALUE.Map_A_DEFAULT_VALUEDBID_A_DEFAULT_VALUEDB {

		var a_default_valueAPI A_DEFAULT_VALUEAPI
		a_default_valueAPI.ID = a_default_valueDB.ID
		a_default_valueAPI.A_DEFAULT_VALUEPointersEncoding = a_default_valueDB.A_DEFAULT_VALUEPointersEncoding
		a_default_valueDB.CopyBasicFieldsToA_DEFAULT_VALUE_WOP(&a_default_valueAPI.A_DEFAULT_VALUE_WOP)

		backRepoData.A_DEFAULT_VALUEAPIs = append(backRepoData.A_DEFAULT_VALUEAPIs, &a_default_valueAPI)
	}

	for _, a_default_value_1DB := range backRepo.BackRepoA_DEFAULT_VALUE_1.Map_A_DEFAULT_VALUE_1DBID_A_DEFAULT_VALUE_1DB {

		var a_default_value_1API A_DEFAULT_VALUE_1API
		a_default_value_1API.ID = a_default_value_1DB.ID
		a_default_value_1API.A_DEFAULT_VALUE_1PointersEncoding = a_default_value_1DB.A_DEFAULT_VALUE_1PointersEncoding
		a_default_value_1DB.CopyBasicFieldsToA_DEFAULT_VALUE_1_WOP(&a_default_value_1API.A_DEFAULT_VALUE_1_WOP)

		backRepoData.A_DEFAULT_VALUE_1APIs = append(backRepoData.A_DEFAULT_VALUE_1APIs, &a_default_value_1API)
	}

	for _, a_default_value_2DB := range backRepo.BackRepoA_DEFAULT_VALUE_2.Map_A_DEFAULT_VALUE_2DBID_A_DEFAULT_VALUE_2DB {

		var a_default_value_2API A_DEFAULT_VALUE_2API
		a_default_value_2API.ID = a_default_value_2DB.ID
		a_default_value_2API.A_DEFAULT_VALUE_2PointersEncoding = a_default_value_2DB.A_DEFAULT_VALUE_2PointersEncoding
		a_default_value_2DB.CopyBasicFieldsToA_DEFAULT_VALUE_2_WOP(&a_default_value_2API.A_DEFAULT_VALUE_2_WOP)

		backRepoData.A_DEFAULT_VALUE_2APIs = append(backRepoData.A_DEFAULT_VALUE_2APIs, &a_default_value_2API)
	}

	for _, a_default_value_3DB := range backRepo.BackRepoA_DEFAULT_VALUE_3.Map_A_DEFAULT_VALUE_3DBID_A_DEFAULT_VALUE_3DB {

		var a_default_value_3API A_DEFAULT_VALUE_3API
		a_default_value_3API.ID = a_default_value_3DB.ID
		a_default_value_3API.A_DEFAULT_VALUE_3PointersEncoding = a_default_value_3DB.A_DEFAULT_VALUE_3PointersEncoding
		a_default_value_3DB.CopyBasicFieldsToA_DEFAULT_VALUE_3_WOP(&a_default_value_3API.A_DEFAULT_VALUE_3_WOP)

		backRepoData.A_DEFAULT_VALUE_3APIs = append(backRepoData.A_DEFAULT_VALUE_3APIs, &a_default_value_3API)
	}

	for _, a_default_value_4DB := range backRepo.BackRepoA_DEFAULT_VALUE_4.Map_A_DEFAULT_VALUE_4DBID_A_DEFAULT_VALUE_4DB {

		var a_default_value_4API A_DEFAULT_VALUE_4API
		a_default_value_4API.ID = a_default_value_4DB.ID
		a_default_value_4API.A_DEFAULT_VALUE_4PointersEncoding = a_default_value_4DB.A_DEFAULT_VALUE_4PointersEncoding
		a_default_value_4DB.CopyBasicFieldsToA_DEFAULT_VALUE_4_WOP(&a_default_value_4API.A_DEFAULT_VALUE_4_WOP)

		backRepoData.A_DEFAULT_VALUE_4APIs = append(backRepoData.A_DEFAULT_VALUE_4APIs, &a_default_value_4API)
	}

	for _, a_default_value_5DB := range backRepo.BackRepoA_DEFAULT_VALUE_5.Map_A_DEFAULT_VALUE_5DBID_A_DEFAULT_VALUE_5DB {

		var a_default_value_5API A_DEFAULT_VALUE_5API
		a_default_value_5API.ID = a_default_value_5DB.ID
		a_default_value_5API.A_DEFAULT_VALUE_5PointersEncoding = a_default_value_5DB.A_DEFAULT_VALUE_5PointersEncoding
		a_default_value_5DB.CopyBasicFieldsToA_DEFAULT_VALUE_5_WOP(&a_default_value_5API.A_DEFAULT_VALUE_5_WOP)

		backRepoData.A_DEFAULT_VALUE_5APIs = append(backRepoData.A_DEFAULT_VALUE_5APIs, &a_default_value_5API)
	}

	for _, a_default_value_6DB := range backRepo.BackRepoA_DEFAULT_VALUE_6.Map_A_DEFAULT_VALUE_6DBID_A_DEFAULT_VALUE_6DB {

		var a_default_value_6API A_DEFAULT_VALUE_6API
		a_default_value_6API.ID = a_default_value_6DB.ID
		a_default_value_6API.A_DEFAULT_VALUE_6PointersEncoding = a_default_value_6DB.A_DEFAULT_VALUE_6PointersEncoding
		a_default_value_6DB.CopyBasicFieldsToA_DEFAULT_VALUE_6_WOP(&a_default_value_6API.A_DEFAULT_VALUE_6_WOP)

		backRepoData.A_DEFAULT_VALUE_6APIs = append(backRepoData.A_DEFAULT_VALUE_6APIs, &a_default_value_6API)
	}

	for _, a_definitionDB := range backRepo.BackRepoA_DEFINITION.Map_A_DEFINITIONDBID_A_DEFINITIONDB {

		var a_definitionAPI A_DEFINITIONAPI
		a_definitionAPI.ID = a_definitionDB.ID
		a_definitionAPI.A_DEFINITIONPointersEncoding = a_definitionDB.A_DEFINITIONPointersEncoding
		a_definitionDB.CopyBasicFieldsToA_DEFINITION_WOP(&a_definitionAPI.A_DEFINITION_WOP)

		backRepoData.A_DEFINITIONAPIs = append(backRepoData.A_DEFINITIONAPIs, &a_definitionAPI)
	}

	for _, a_definition_1DB := range backRepo.BackRepoA_DEFINITION_1.Map_A_DEFINITION_1DBID_A_DEFINITION_1DB {

		var a_definition_1API A_DEFINITION_1API
		a_definition_1API.ID = a_definition_1DB.ID
		a_definition_1API.A_DEFINITION_1PointersEncoding = a_definition_1DB.A_DEFINITION_1PointersEncoding
		a_definition_1DB.CopyBasicFieldsToA_DEFINITION_1_WOP(&a_definition_1API.A_DEFINITION_1_WOP)

		backRepoData.A_DEFINITION_1APIs = append(backRepoData.A_DEFINITION_1APIs, &a_definition_1API)
	}

	for _, a_definition_2DB := range backRepo.BackRepoA_DEFINITION_2.Map_A_DEFINITION_2DBID_A_DEFINITION_2DB {

		var a_definition_2API A_DEFINITION_2API
		a_definition_2API.ID = a_definition_2DB.ID
		a_definition_2API.A_DEFINITION_2PointersEncoding = a_definition_2DB.A_DEFINITION_2PointersEncoding
		a_definition_2DB.CopyBasicFieldsToA_DEFINITION_2_WOP(&a_definition_2API.A_DEFINITION_2_WOP)

		backRepoData.A_DEFINITION_2APIs = append(backRepoData.A_DEFINITION_2APIs, &a_definition_2API)
	}

	for _, a_definition_3DB := range backRepo.BackRepoA_DEFINITION_3.Map_A_DEFINITION_3DBID_A_DEFINITION_3DB {

		var a_definition_3API A_DEFINITION_3API
		a_definition_3API.ID = a_definition_3DB.ID
		a_definition_3API.A_DEFINITION_3PointersEncoding = a_definition_3DB.A_DEFINITION_3PointersEncoding
		a_definition_3DB.CopyBasicFieldsToA_DEFINITION_3_WOP(&a_definition_3API.A_DEFINITION_3_WOP)

		backRepoData.A_DEFINITION_3APIs = append(backRepoData.A_DEFINITION_3APIs, &a_definition_3API)
	}

	for _, a_definition_4DB := range backRepo.BackRepoA_DEFINITION_4.Map_A_DEFINITION_4DBID_A_DEFINITION_4DB {

		var a_definition_4API A_DEFINITION_4API
		a_definition_4API.ID = a_definition_4DB.ID
		a_definition_4API.A_DEFINITION_4PointersEncoding = a_definition_4DB.A_DEFINITION_4PointersEncoding
		a_definition_4DB.CopyBasicFieldsToA_DEFINITION_4_WOP(&a_definition_4API.A_DEFINITION_4_WOP)

		backRepoData.A_DEFINITION_4APIs = append(backRepoData.A_DEFINITION_4APIs, &a_definition_4API)
	}

	for _, a_definition_5DB := range backRepo.BackRepoA_DEFINITION_5.Map_A_DEFINITION_5DBID_A_DEFINITION_5DB {

		var a_definition_5API A_DEFINITION_5API
		a_definition_5API.ID = a_definition_5DB.ID
		a_definition_5API.A_DEFINITION_5PointersEncoding = a_definition_5DB.A_DEFINITION_5PointersEncoding
		a_definition_5DB.CopyBasicFieldsToA_DEFINITION_5_WOP(&a_definition_5API.A_DEFINITION_5_WOP)

		backRepoData.A_DEFINITION_5APIs = append(backRepoData.A_DEFINITION_5APIs, &a_definition_5API)
	}

	for _, a_definition_6DB := range backRepo.BackRepoA_DEFINITION_6.Map_A_DEFINITION_6DBID_A_DEFINITION_6DB {

		var a_definition_6API A_DEFINITION_6API
		a_definition_6API.ID = a_definition_6DB.ID
		a_definition_6API.A_DEFINITION_6PointersEncoding = a_definition_6DB.A_DEFINITION_6PointersEncoding
		a_definition_6DB.CopyBasicFieldsToA_DEFINITION_6_WOP(&a_definition_6API.A_DEFINITION_6_WOP)

		backRepoData.A_DEFINITION_6APIs = append(backRepoData.A_DEFINITION_6APIs, &a_definition_6API)
	}

	for _, a_editable_attsDB := range backRepo.BackRepoA_EDITABLE_ATTS.Map_A_EDITABLE_ATTSDBID_A_EDITABLE_ATTSDB {

		var a_editable_attsAPI A_EDITABLE_ATTSAPI
		a_editable_attsAPI.ID = a_editable_attsDB.ID
		a_editable_attsAPI.A_EDITABLE_ATTSPointersEncoding = a_editable_attsDB.A_EDITABLE_ATTSPointersEncoding
		a_editable_attsDB.CopyBasicFieldsToA_EDITABLE_ATTS_WOP(&a_editable_attsAPI.A_EDITABLE_ATTS_WOP)

		backRepoData.A_EDITABLE_ATTSAPIs = append(backRepoData.A_EDITABLE_ATTSAPIs, &a_editable_attsAPI)
	}

	for _, a_objectDB := range backRepo.BackRepoA_OBJECT.Map_A_OBJECTDBID_A_OBJECTDB {

		var a_objectAPI A_OBJECTAPI
		a_objectAPI.ID = a_objectDB.ID
		a_objectAPI.A_OBJECTPointersEncoding = a_objectDB.A_OBJECTPointersEncoding
		a_objectDB.CopyBasicFieldsToA_OBJECT_WOP(&a_objectAPI.A_OBJECT_WOP)

		backRepoData.A_OBJECTAPIs = append(backRepoData.A_OBJECTAPIs, &a_objectAPI)
	}

	for _, a_propertiesDB := range backRepo.BackRepoA_PROPERTIES.Map_A_PROPERTIESDBID_A_PROPERTIESDB {

		var a_propertiesAPI A_PROPERTIESAPI
		a_propertiesAPI.ID = a_propertiesDB.ID
		a_propertiesAPI.A_PROPERTIESPointersEncoding = a_propertiesDB.A_PROPERTIESPointersEncoding
		a_propertiesDB.CopyBasicFieldsToA_PROPERTIES_WOP(&a_propertiesAPI.A_PROPERTIES_WOP)

		backRepoData.A_PROPERTIESAPIs = append(backRepoData.A_PROPERTIESAPIs, &a_propertiesAPI)
	}

	for _, a_sourceDB := range backRepo.BackRepoA_SOURCE.Map_A_SOURCEDBID_A_SOURCEDB {

		var a_sourceAPI A_SOURCEAPI
		a_sourceAPI.ID = a_sourceDB.ID
		a_sourceAPI.A_SOURCEPointersEncoding = a_sourceDB.A_SOURCEPointersEncoding
		a_sourceDB.CopyBasicFieldsToA_SOURCE_WOP(&a_sourceAPI.A_SOURCE_WOP)

		backRepoData.A_SOURCEAPIs = append(backRepoData.A_SOURCEAPIs, &a_sourceAPI)
	}

	for _, a_source_specificationDB := range backRepo.BackRepoA_SOURCE_SPECIFICATION.Map_A_SOURCE_SPECIFICATIONDBID_A_SOURCE_SPECIFICATIONDB {

		var a_source_specificationAPI A_SOURCE_SPECIFICATIONAPI
		a_source_specificationAPI.ID = a_source_specificationDB.ID
		a_source_specificationAPI.A_SOURCE_SPECIFICATIONPointersEncoding = a_source_specificationDB.A_SOURCE_SPECIFICATIONPointersEncoding
		a_source_specificationDB.CopyBasicFieldsToA_SOURCE_SPECIFICATION_WOP(&a_source_specificationAPI.A_SOURCE_SPECIFICATION_WOP)

		backRepoData.A_SOURCE_SPECIFICATIONAPIs = append(backRepoData.A_SOURCE_SPECIFICATIONAPIs, &a_source_specificationAPI)
	}

	for _, a_specificationsDB := range backRepo.BackRepoA_SPECIFICATIONS.Map_A_SPECIFICATIONSDBID_A_SPECIFICATIONSDB {

		var a_specificationsAPI A_SPECIFICATIONSAPI
		a_specificationsAPI.ID = a_specificationsDB.ID
		a_specificationsAPI.A_SPECIFICATIONSPointersEncoding = a_specificationsDB.A_SPECIFICATIONSPointersEncoding
		a_specificationsDB.CopyBasicFieldsToA_SPECIFICATIONS_WOP(&a_specificationsAPI.A_SPECIFICATIONS_WOP)

		backRepoData.A_SPECIFICATIONSAPIs = append(backRepoData.A_SPECIFICATIONSAPIs, &a_specificationsAPI)
	}

	for _, a_specified_valuesDB := range backRepo.BackRepoA_SPECIFIED_VALUES.Map_A_SPECIFIED_VALUESDBID_A_SPECIFIED_VALUESDB {

		var a_specified_valuesAPI A_SPECIFIED_VALUESAPI
		a_specified_valuesAPI.ID = a_specified_valuesDB.ID
		a_specified_valuesAPI.A_SPECIFIED_VALUESPointersEncoding = a_specified_valuesDB.A_SPECIFIED_VALUESPointersEncoding
		a_specified_valuesDB.CopyBasicFieldsToA_SPECIFIED_VALUES_WOP(&a_specified_valuesAPI.A_SPECIFIED_VALUES_WOP)

		backRepoData.A_SPECIFIED_VALUESAPIs = append(backRepoData.A_SPECIFIED_VALUESAPIs, &a_specified_valuesAPI)
	}

	for _, a_spec_attributesDB := range backRepo.BackRepoA_SPEC_ATTRIBUTES.Map_A_SPEC_ATTRIBUTESDBID_A_SPEC_ATTRIBUTESDB {

		var a_spec_attributesAPI A_SPEC_ATTRIBUTESAPI
		a_spec_attributesAPI.ID = a_spec_attributesDB.ID
		a_spec_attributesAPI.A_SPEC_ATTRIBUTESPointersEncoding = a_spec_attributesDB.A_SPEC_ATTRIBUTESPointersEncoding
		a_spec_attributesDB.CopyBasicFieldsToA_SPEC_ATTRIBUTES_WOP(&a_spec_attributesAPI.A_SPEC_ATTRIBUTES_WOP)

		backRepoData.A_SPEC_ATTRIBUTESAPIs = append(backRepoData.A_SPEC_ATTRIBUTESAPIs, &a_spec_attributesAPI)
	}

	for _, a_spec_objectsDB := range backRepo.BackRepoA_SPEC_OBJECTS.Map_A_SPEC_OBJECTSDBID_A_SPEC_OBJECTSDB {

		var a_spec_objectsAPI A_SPEC_OBJECTSAPI
		a_spec_objectsAPI.ID = a_spec_objectsDB.ID
		a_spec_objectsAPI.A_SPEC_OBJECTSPointersEncoding = a_spec_objectsDB.A_SPEC_OBJECTSPointersEncoding
		a_spec_objectsDB.CopyBasicFieldsToA_SPEC_OBJECTS_WOP(&a_spec_objectsAPI.A_SPEC_OBJECTS_WOP)

		backRepoData.A_SPEC_OBJECTSAPIs = append(backRepoData.A_SPEC_OBJECTSAPIs, &a_spec_objectsAPI)
	}

	for _, a_spec_relationsDB := range backRepo.BackRepoA_SPEC_RELATIONS.Map_A_SPEC_RELATIONSDBID_A_SPEC_RELATIONSDB {

		var a_spec_relationsAPI A_SPEC_RELATIONSAPI
		a_spec_relationsAPI.ID = a_spec_relationsDB.ID
		a_spec_relationsAPI.A_SPEC_RELATIONSPointersEncoding = a_spec_relationsDB.A_SPEC_RELATIONSPointersEncoding
		a_spec_relationsDB.CopyBasicFieldsToA_SPEC_RELATIONS_WOP(&a_spec_relationsAPI.A_SPEC_RELATIONS_WOP)

		backRepoData.A_SPEC_RELATIONSAPIs = append(backRepoData.A_SPEC_RELATIONSAPIs, &a_spec_relationsAPI)
	}

	for _, a_spec_relations_1DB := range backRepo.BackRepoA_SPEC_RELATIONS_1.Map_A_SPEC_RELATIONS_1DBID_A_SPEC_RELATIONS_1DB {

		var a_spec_relations_1API A_SPEC_RELATIONS_1API
		a_spec_relations_1API.ID = a_spec_relations_1DB.ID
		a_spec_relations_1API.A_SPEC_RELATIONS_1PointersEncoding = a_spec_relations_1DB.A_SPEC_RELATIONS_1PointersEncoding
		a_spec_relations_1DB.CopyBasicFieldsToA_SPEC_RELATIONS_1_WOP(&a_spec_relations_1API.A_SPEC_RELATIONS_1_WOP)

		backRepoData.A_SPEC_RELATIONS_1APIs = append(backRepoData.A_SPEC_RELATIONS_1APIs, &a_spec_relations_1API)
	}

	for _, a_spec_relation_groupsDB := range backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSDB {

		var a_spec_relation_groupsAPI A_SPEC_RELATION_GROUPSAPI
		a_spec_relation_groupsAPI.ID = a_spec_relation_groupsDB.ID
		a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPSPointersEncoding = a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding
		a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS_WOP(&a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPS_WOP)

		backRepoData.A_SPEC_RELATION_GROUPSAPIs = append(backRepoData.A_SPEC_RELATION_GROUPSAPIs, &a_spec_relation_groupsAPI)
	}

	for _, a_spec_typesDB := range backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESDB {

		var a_spec_typesAPI A_SPEC_TYPESAPI
		a_spec_typesAPI.ID = a_spec_typesDB.ID
		a_spec_typesAPI.A_SPEC_TYPESPointersEncoding = a_spec_typesDB.A_SPEC_TYPESPointersEncoding
		a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES_WOP(&a_spec_typesAPI.A_SPEC_TYPES_WOP)

		backRepoData.A_SPEC_TYPESAPIs = append(backRepoData.A_SPEC_TYPESAPIs, &a_spec_typesAPI)
	}

	for _, a_the_headerDB := range backRepo.BackRepoA_THE_HEADER.Map_A_THE_HEADERDBID_A_THE_HEADERDB {

		var a_the_headerAPI A_THE_HEADERAPI
		a_the_headerAPI.ID = a_the_headerDB.ID
		a_the_headerAPI.A_THE_HEADERPointersEncoding = a_the_headerDB.A_THE_HEADERPointersEncoding
		a_the_headerDB.CopyBasicFieldsToA_THE_HEADER_WOP(&a_the_headerAPI.A_THE_HEADER_WOP)

		backRepoData.A_THE_HEADERAPIs = append(backRepoData.A_THE_HEADERAPIs, &a_the_headerAPI)
	}

	for _, a_tool_extensionsDB := range backRepo.BackRepoA_TOOL_EXTENSIONS.Map_A_TOOL_EXTENSIONSDBID_A_TOOL_EXTENSIONSDB {

		var a_tool_extensionsAPI A_TOOL_EXTENSIONSAPI
		a_tool_extensionsAPI.ID = a_tool_extensionsDB.ID
		a_tool_extensionsAPI.A_TOOL_EXTENSIONSPointersEncoding = a_tool_extensionsDB.A_TOOL_EXTENSIONSPointersEncoding
		a_tool_extensionsDB.CopyBasicFieldsToA_TOOL_EXTENSIONS_WOP(&a_tool_extensionsAPI.A_TOOL_EXTENSIONS_WOP)

		backRepoData.A_TOOL_EXTENSIONSAPIs = append(backRepoData.A_TOOL_EXTENSIONSAPIs, &a_tool_extensionsAPI)
	}

	for _, a_typeDB := range backRepo.BackRepoA_TYPE.Map_A_TYPEDBID_A_TYPEDB {

		var a_typeAPI A_TYPEAPI
		a_typeAPI.ID = a_typeDB.ID
		a_typeAPI.A_TYPEPointersEncoding = a_typeDB.A_TYPEPointersEncoding
		a_typeDB.CopyBasicFieldsToA_TYPE_WOP(&a_typeAPI.A_TYPE_WOP)

		backRepoData.A_TYPEAPIs = append(backRepoData.A_TYPEAPIs, &a_typeAPI)
	}

	for _, a_type_1DB := range backRepo.BackRepoA_TYPE_1.Map_A_TYPE_1DBID_A_TYPE_1DB {

		var a_type_1API A_TYPE_1API
		a_type_1API.ID = a_type_1DB.ID
		a_type_1API.A_TYPE_1PointersEncoding = a_type_1DB.A_TYPE_1PointersEncoding
		a_type_1DB.CopyBasicFieldsToA_TYPE_1_WOP(&a_type_1API.A_TYPE_1_WOP)

		backRepoData.A_TYPE_1APIs = append(backRepoData.A_TYPE_1APIs, &a_type_1API)
	}

	for _, a_type_10DB := range backRepo.BackRepoA_TYPE_10.Map_A_TYPE_10DBID_A_TYPE_10DB {

		var a_type_10API A_TYPE_10API
		a_type_10API.ID = a_type_10DB.ID
		a_type_10API.A_TYPE_10PointersEncoding = a_type_10DB.A_TYPE_10PointersEncoding
		a_type_10DB.CopyBasicFieldsToA_TYPE_10_WOP(&a_type_10API.A_TYPE_10_WOP)

		backRepoData.A_TYPE_10APIs = append(backRepoData.A_TYPE_10APIs, &a_type_10API)
	}

	for _, a_type_2DB := range backRepo.BackRepoA_TYPE_2.Map_A_TYPE_2DBID_A_TYPE_2DB {

		var a_type_2API A_TYPE_2API
		a_type_2API.ID = a_type_2DB.ID
		a_type_2API.A_TYPE_2PointersEncoding = a_type_2DB.A_TYPE_2PointersEncoding
		a_type_2DB.CopyBasicFieldsToA_TYPE_2_WOP(&a_type_2API.A_TYPE_2_WOP)

		backRepoData.A_TYPE_2APIs = append(backRepoData.A_TYPE_2APIs, &a_type_2API)
	}

	for _, a_type_3DB := range backRepo.BackRepoA_TYPE_3.Map_A_TYPE_3DBID_A_TYPE_3DB {

		var a_type_3API A_TYPE_3API
		a_type_3API.ID = a_type_3DB.ID
		a_type_3API.A_TYPE_3PointersEncoding = a_type_3DB.A_TYPE_3PointersEncoding
		a_type_3DB.CopyBasicFieldsToA_TYPE_3_WOP(&a_type_3API.A_TYPE_3_WOP)

		backRepoData.A_TYPE_3APIs = append(backRepoData.A_TYPE_3APIs, &a_type_3API)
	}

	for _, a_type_4DB := range backRepo.BackRepoA_TYPE_4.Map_A_TYPE_4DBID_A_TYPE_4DB {

		var a_type_4API A_TYPE_4API
		a_type_4API.ID = a_type_4DB.ID
		a_type_4API.A_TYPE_4PointersEncoding = a_type_4DB.A_TYPE_4PointersEncoding
		a_type_4DB.CopyBasicFieldsToA_TYPE_4_WOP(&a_type_4API.A_TYPE_4_WOP)

		backRepoData.A_TYPE_4APIs = append(backRepoData.A_TYPE_4APIs, &a_type_4API)
	}

	for _, a_type_5DB := range backRepo.BackRepoA_TYPE_5.Map_A_TYPE_5DBID_A_TYPE_5DB {

		var a_type_5API A_TYPE_5API
		a_type_5API.ID = a_type_5DB.ID
		a_type_5API.A_TYPE_5PointersEncoding = a_type_5DB.A_TYPE_5PointersEncoding
		a_type_5DB.CopyBasicFieldsToA_TYPE_5_WOP(&a_type_5API.A_TYPE_5_WOP)

		backRepoData.A_TYPE_5APIs = append(backRepoData.A_TYPE_5APIs, &a_type_5API)
	}

	for _, a_type_6DB := range backRepo.BackRepoA_TYPE_6.Map_A_TYPE_6DBID_A_TYPE_6DB {

		var a_type_6API A_TYPE_6API
		a_type_6API.ID = a_type_6DB.ID
		a_type_6API.A_TYPE_6PointersEncoding = a_type_6DB.A_TYPE_6PointersEncoding
		a_type_6DB.CopyBasicFieldsToA_TYPE_6_WOP(&a_type_6API.A_TYPE_6_WOP)

		backRepoData.A_TYPE_6APIs = append(backRepoData.A_TYPE_6APIs, &a_type_6API)
	}

	for _, a_type_7DB := range backRepo.BackRepoA_TYPE_7.Map_A_TYPE_7DBID_A_TYPE_7DB {

		var a_type_7API A_TYPE_7API
		a_type_7API.ID = a_type_7DB.ID
		a_type_7API.A_TYPE_7PointersEncoding = a_type_7DB.A_TYPE_7PointersEncoding
		a_type_7DB.CopyBasicFieldsToA_TYPE_7_WOP(&a_type_7API.A_TYPE_7_WOP)

		backRepoData.A_TYPE_7APIs = append(backRepoData.A_TYPE_7APIs, &a_type_7API)
	}

	for _, a_type_8DB := range backRepo.BackRepoA_TYPE_8.Map_A_TYPE_8DBID_A_TYPE_8DB {

		var a_type_8API A_TYPE_8API
		a_type_8API.ID = a_type_8DB.ID
		a_type_8API.A_TYPE_8PointersEncoding = a_type_8DB.A_TYPE_8PointersEncoding
		a_type_8DB.CopyBasicFieldsToA_TYPE_8_WOP(&a_type_8API.A_TYPE_8_WOP)

		backRepoData.A_TYPE_8APIs = append(backRepoData.A_TYPE_8APIs, &a_type_8API)
	}

	for _, a_type_9DB := range backRepo.BackRepoA_TYPE_9.Map_A_TYPE_9DBID_A_TYPE_9DB {

		var a_type_9API A_TYPE_9API
		a_type_9API.ID = a_type_9DB.ID
		a_type_9API.A_TYPE_9PointersEncoding = a_type_9DB.A_TYPE_9PointersEncoding
		a_type_9DB.CopyBasicFieldsToA_TYPE_9_WOP(&a_type_9API.A_TYPE_9_WOP)

		backRepoData.A_TYPE_9APIs = append(backRepoData.A_TYPE_9APIs, &a_type_9API)
	}

	for _, a_valuesDB := range backRepo.BackRepoA_VALUES.Map_A_VALUESDBID_A_VALUESDB {

		var a_valuesAPI A_VALUESAPI
		a_valuesAPI.ID = a_valuesDB.ID
		a_valuesAPI.A_VALUESPointersEncoding = a_valuesDB.A_VALUESPointersEncoding
		a_valuesDB.CopyBasicFieldsToA_VALUES_WOP(&a_valuesAPI.A_VALUES_WOP)

		backRepoData.A_VALUESAPIs = append(backRepoData.A_VALUESAPIs, &a_valuesAPI)
	}

	for _, a_values_1DB := range backRepo.BackRepoA_VALUES_1.Map_A_VALUES_1DBID_A_VALUES_1DB {

		var a_values_1API A_VALUES_1API
		a_values_1API.ID = a_values_1DB.ID
		a_values_1API.A_VALUES_1PointersEncoding = a_values_1DB.A_VALUES_1PointersEncoding
		a_values_1DB.CopyBasicFieldsToA_VALUES_1_WOP(&a_values_1API.A_VALUES_1_WOP)

		backRepoData.A_VALUES_1APIs = append(backRepoData.A_VALUES_1APIs, &a_values_1API)
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
