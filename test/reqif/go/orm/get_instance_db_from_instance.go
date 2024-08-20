// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		alternative_idInstance := any(concreteInstance).(*models.ALTERNATIVE_ID)
		ret2 := backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		attribute_definition_booleanInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetATTRIBUTE_DEFINITION_BOOLEANDBFromATTRIBUTE_DEFINITION_BOOLEANPtr(attribute_definition_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		attribute_definition_dateInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_DATE)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetATTRIBUTE_DEFINITION_DATEDBFromATTRIBUTE_DEFINITION_DATEPtr(attribute_definition_dateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		attribute_definition_enumerationInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetATTRIBUTE_DEFINITION_ENUMERATIONDBFromATTRIBUTE_DEFINITION_ENUMERATIONPtr(attribute_definition_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		attribute_definition_integerInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_INTEGER)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetATTRIBUTE_DEFINITION_INTEGERDBFromATTRIBUTE_DEFINITION_INTEGERPtr(attribute_definition_integerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		attribute_definition_realInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_REAL)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetATTRIBUTE_DEFINITION_REALDBFromATTRIBUTE_DEFINITION_REALPtr(attribute_definition_realInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		attribute_definition_stringInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_STRING)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetATTRIBUTE_DEFINITION_STRINGDBFromATTRIBUTE_DEFINITION_STRINGPtr(attribute_definition_stringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		attribute_definition_xhtmlInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_XHTML)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetATTRIBUTE_DEFINITION_XHTMLDBFromATTRIBUTE_DEFINITION_XHTMLPtr(attribute_definition_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		attribute_value_booleanInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_BOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetATTRIBUTE_VALUE_BOOLEANDBFromATTRIBUTE_VALUE_BOOLEANPtr(attribute_value_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_DATE:
		attribute_value_dateInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_DATE)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetATTRIBUTE_VALUE_DATEDBFromATTRIBUTE_VALUE_DATEPtr(attribute_value_dateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		attribute_value_enumerationInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_ENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetATTRIBUTE_VALUE_ENUMERATIONDBFromATTRIBUTE_VALUE_ENUMERATIONPtr(attribute_value_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		attribute_value_integerInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_INTEGER)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetATTRIBUTE_VALUE_INTEGERDBFromATTRIBUTE_VALUE_INTEGERPtr(attribute_value_integerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_REAL:
		attribute_value_realInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_REAL)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetATTRIBUTE_VALUE_REALDBFromATTRIBUTE_VALUE_REALPtr(attribute_value_realInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_STRING:
		attribute_value_stringInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_STRING)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetATTRIBUTE_VALUE_STRINGDBFromATTRIBUTE_VALUE_STRINGPtr(attribute_value_stringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_XHTML:
		attribute_value_xhtmlInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_XHTML)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetATTRIBUTE_VALUE_XHTMLDBFromATTRIBUTE_VALUE_XHTMLPtr(attribute_value_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		datatype_definition_booleanInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_BOOLEAN)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDATATYPE_DEFINITION_BOOLEANDBFromDATATYPE_DEFINITION_BOOLEANPtr(datatype_definition_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_DATE:
		datatype_definition_dateInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_DATE)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDATATYPE_DEFINITION_DATEDBFromDATATYPE_DEFINITION_DATEPtr(datatype_definition_dateInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		datatype_definition_enumerationInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_ENUMERATION)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDATATYPE_DEFINITION_ENUMERATIONDBFromDATATYPE_DEFINITION_ENUMERATIONPtr(datatype_definition_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_INTEGER:
		datatype_definition_integerInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_INTEGER)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDATATYPE_DEFINITION_INTEGERDBFromDATATYPE_DEFINITION_INTEGERPtr(datatype_definition_integerInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_REAL:
		datatype_definition_realInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_REAL)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDATATYPE_DEFINITION_REALDBFromDATATYPE_DEFINITION_REALPtr(datatype_definition_realInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_STRING:
		datatype_definition_stringInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_STRING)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDATATYPE_DEFINITION_STRINGDBFromDATATYPE_DEFINITION_STRINGPtr(datatype_definition_stringInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_XHTML:
		datatype_definition_xhtmlInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_XHTML)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDATATYPE_DEFINITION_XHTMLDBFromDATATYPE_DEFINITION_XHTMLPtr(datatype_definition_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.EMBEDDED_VALUE:
		embedded_valueInstance := any(concreteInstance).(*models.EMBEDDED_VALUE)
		ret2 := backRepo.BackRepoEMBEDDED_VALUE.GetEMBEDDED_VALUEDBFromEMBEDDED_VALUEPtr(embedded_valueInstance)
		ret = any(ret2).(*T2)
	case *models.ENUM_VALUE:
		enum_valueInstance := any(concreteInstance).(*models.ENUM_VALUE)
		ret2 := backRepo.BackRepoENUM_VALUE.GetENUM_VALUEDBFromENUM_VALUEPtr(enum_valueInstance)
		ret = any(ret2).(*T2)
	case *models.RELATION_GROUP:
		relation_groupInstance := any(concreteInstance).(*models.RELATION_GROUP)
		ret2 := backRepo.BackRepoRELATION_GROUP.GetRELATION_GROUPDBFromRELATION_GROUPPtr(relation_groupInstance)
		ret = any(ret2).(*T2)
	case *models.RELATION_GROUP_TYPE:
		relation_group_typeInstance := any(concreteInstance).(*models.RELATION_GROUP_TYPE)
		ret2 := backRepo.BackRepoRELATION_GROUP_TYPE.GetRELATION_GROUP_TYPEDBFromRELATION_GROUP_TYPEPtr(relation_group_typeInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF:
		req_ifInstance := any(concreteInstance).(*models.REQ_IF)
		ret2 := backRepo.BackRepoREQ_IF.GetREQ_IFDBFromREQ_IFPtr(req_ifInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_CONTENT:
		req_if_contentInstance := any(concreteInstance).(*models.REQ_IF_CONTENT)
		ret2 := backRepo.BackRepoREQ_IF_CONTENT.GetREQ_IF_CONTENTDBFromREQ_IF_CONTENTPtr(req_if_contentInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_HEADER:
		req_if_headerInstance := any(concreteInstance).(*models.REQ_IF_HEADER)
		ret2 := backRepo.BackRepoREQ_IF_HEADER.GetREQ_IF_HEADERDBFromREQ_IF_HEADERPtr(req_if_headerInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_TOOL_EXTENSION:
		req_if_tool_extensionInstance := any(concreteInstance).(*models.REQ_IF_TOOL_EXTENSION)
		ret2 := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetREQ_IF_TOOL_EXTENSIONDBFromREQ_IF_TOOL_EXTENSIONPtr(req_if_tool_extensionInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATION:
		specificationInstance := any(concreteInstance).(*models.SPECIFICATION)
		ret2 := backRepo.BackRepoSPECIFICATION.GetSPECIFICATIONDBFromSPECIFICATIONPtr(specificationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATION_TYPE:
		specification_typeInstance := any(concreteInstance).(*models.SPECIFICATION_TYPE)
		ret2 := backRepo.BackRepoSPECIFICATION_TYPE.GetSPECIFICATION_TYPEDBFromSPECIFICATION_TYPEPtr(specification_typeInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_HIERARCHY:
		spec_hierarchyInstance := any(concreteInstance).(*models.SPEC_HIERARCHY)
		ret2 := backRepo.BackRepoSPEC_HIERARCHY.GetSPEC_HIERARCHYDBFromSPEC_HIERARCHYPtr(spec_hierarchyInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_OBJECT:
		spec_objectInstance := any(concreteInstance).(*models.SPEC_OBJECT)
		ret2 := backRepo.BackRepoSPEC_OBJECT.GetSPEC_OBJECTDBFromSPEC_OBJECTPtr(spec_objectInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_OBJECT_TYPE:
		spec_object_typeInstance := any(concreteInstance).(*models.SPEC_OBJECT_TYPE)
		ret2 := backRepo.BackRepoSPEC_OBJECT_TYPE.GetSPEC_OBJECT_TYPEDBFromSPEC_OBJECT_TYPEPtr(spec_object_typeInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_RELATION:
		spec_relationInstance := any(concreteInstance).(*models.SPEC_RELATION)
		ret2 := backRepo.BackRepoSPEC_RELATION.GetSPEC_RELATIONDBFromSPEC_RELATIONPtr(spec_relationInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_RELATION_TYPE:
		spec_relation_typeInstance := any(concreteInstance).(*models.SPEC_RELATION_TYPE)
		ret2 := backRepo.BackRepoSPEC_RELATION_TYPE.GetSPEC_RELATION_TYPEDBFromSPEC_RELATION_TYPEPtr(spec_relation_typeInstance)
		ret = any(ret2).(*T2)
	case *models.XHTML_CONTENT:
		xhtml_contentInstance := any(concreteInstance).(*models.XHTML_CONTENT)
		ret2 := backRepo.BackRepoXHTML_CONTENT.GetXHTML_CONTENTDBFromXHTML_CONTENTPtr(xhtml_contentInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVE_ID, ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_BOOLEAN, ATTRIBUTE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_DATE, ATTRIBUTE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_ENUMERATION, ATTRIBUTE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_INTEGER, ATTRIBUTE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_REAL, ATTRIBUTE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_STRING, ATTRIBUTE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_XHTML, ATTRIBUTE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_BOOLEAN, ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_DATE, ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_ENUMERATION, ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_INTEGER, ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_REAL, ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_STRING, ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_XHTML, ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_BOOLEAN, DATATYPE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_DATE, DATATYPE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_ENUMERATION, DATATYPE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_INTEGER, DATATYPE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_REAL, DATATYPE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_STRING, DATATYPE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_XHTML, DATATYPE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDED_VALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDED_VALUE, EMBEDDED_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUM_VALUE:
		tmp := GetInstanceDBFromInstance[models.ENUM_VALUE, ENUM_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP, RELATION_GROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP_TYPE:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP_TYPE, RELATION_GROUP_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF:
		tmp := GetInstanceDBFromInstance[models.REQ_IF, REQ_IFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_CONTENT:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_CONTENT, REQ_IF_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_HEADER:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_HEADER, REQ_IF_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_TOOL_EXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_TOOL_EXTENSION, REQ_IF_TOOL_EXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION_TYPE, SPECIFICATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_HIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPEC_HIERARCHY, SPEC_HIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT, SPEC_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT_TYPE, SPEC_OBJECT_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION, SPEC_RELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION_TYPE, SPEC_RELATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVE_ID, ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_BOOLEAN, ATTRIBUTE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_DATE, ATTRIBUTE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_ENUMERATION, ATTRIBUTE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_INTEGER, ATTRIBUTE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_REAL, ATTRIBUTE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_STRING, ATTRIBUTE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_XHTML, ATTRIBUTE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_BOOLEAN, ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_DATE, ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_ENUMERATION, ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_INTEGER, ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_REAL, ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_STRING, ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_XHTML, ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_BOOLEAN, DATATYPE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_DATE, DATATYPE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_ENUMERATION, DATATYPE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_INTEGER, DATATYPE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_REAL, DATATYPE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_STRING, DATATYPE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_XHTML, DATATYPE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDED_VALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDED_VALUE, EMBEDDED_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUM_VALUE:
		tmp := GetInstanceDBFromInstance[models.ENUM_VALUE, ENUM_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP, RELATION_GROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP_TYPE:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP_TYPE, RELATION_GROUP_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF:
		tmp := GetInstanceDBFromInstance[models.REQ_IF, REQ_IFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_CONTENT:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_CONTENT, REQ_IF_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_HEADER:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_HEADER, REQ_IF_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_TOOL_EXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_TOOL_EXTENSION, REQ_IF_TOOL_EXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION_TYPE, SPECIFICATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_HIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPEC_HIERARCHY, SPEC_HIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT, SPEC_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT_TYPE, SPEC_OBJECT_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION, SPEC_RELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION_TYPE, SPEC_RELATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
