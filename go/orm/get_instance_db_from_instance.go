// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.All:
		allInstance := any(concreteInstance).(*models.All)
		ret2 := backRepo.BackRepoAll.GetAllDBFromAllPtr(allInstance)
		ret = any(ret2).(*T2)
	case *models.Annotation:
		annotationInstance := any(concreteInstance).(*models.Annotation)
		ret2 := backRepo.BackRepoAnnotation.GetAnnotationDBFromAnnotationPtr(annotationInstance)
		ret = any(ret2).(*T2)
	case *models.Attribute:
		attributeInstance := any(concreteInstance).(*models.Attribute)
		ret2 := backRepo.BackRepoAttribute.GetAttributeDBFromAttributePtr(attributeInstance)
		ret = any(ret2).(*T2)
	case *models.AttributeGroup:
		attributegroupInstance := any(concreteInstance).(*models.AttributeGroup)
		ret2 := backRepo.BackRepoAttributeGroup.GetAttributeGroupDBFromAttributeGroupPtr(attributegroupInstance)
		ret = any(ret2).(*T2)
	case *models.Choice:
		choiceInstance := any(concreteInstance).(*models.Choice)
		ret2 := backRepo.BackRepoChoice.GetChoiceDBFromChoicePtr(choiceInstance)
		ret = any(ret2).(*T2)
	case *models.ComplexContent:
		complexcontentInstance := any(concreteInstance).(*models.ComplexContent)
		ret2 := backRepo.BackRepoComplexContent.GetComplexContentDBFromComplexContentPtr(complexcontentInstance)
		ret = any(ret2).(*T2)
	case *models.ComplexType:
		complextypeInstance := any(concreteInstance).(*models.ComplexType)
		ret2 := backRepo.BackRepoComplexType.GetComplexTypeDBFromComplexTypePtr(complextypeInstance)
		ret = any(ret2).(*T2)
	case *models.Documentation:
		documentationInstance := any(concreteInstance).(*models.Documentation)
		ret2 := backRepo.BackRepoDocumentation.GetDocumentationDBFromDocumentationPtr(documentationInstance)
		ret = any(ret2).(*T2)
	case *models.Element:
		elementInstance := any(concreteInstance).(*models.Element)
		ret2 := backRepo.BackRepoElement.GetElementDBFromElementPtr(elementInstance)
		ret = any(ret2).(*T2)
	case *models.Enumeration:
		enumerationInstance := any(concreteInstance).(*models.Enumeration)
		ret2 := backRepo.BackRepoEnumeration.GetEnumerationDBFromEnumerationPtr(enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.Extension:
		extensionInstance := any(concreteInstance).(*models.Extension)
		ret2 := backRepo.BackRepoExtension.GetExtensionDBFromExtensionPtr(extensionInstance)
		ret = any(ret2).(*T2)
	case *models.Group:
		groupInstance := any(concreteInstance).(*models.Group)
		ret2 := backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupInstance)
		ret = any(ret2).(*T2)
	case *models.Length:
		lengthInstance := any(concreteInstance).(*models.Length)
		ret2 := backRepo.BackRepoLength.GetLengthDBFromLengthPtr(lengthInstance)
		ret = any(ret2).(*T2)
	case *models.MaxInclusive:
		maxinclusiveInstance := any(concreteInstance).(*models.MaxInclusive)
		ret2 := backRepo.BackRepoMaxInclusive.GetMaxInclusiveDBFromMaxInclusivePtr(maxinclusiveInstance)
		ret = any(ret2).(*T2)
	case *models.MaxLength:
		maxlengthInstance := any(concreteInstance).(*models.MaxLength)
		ret2 := backRepo.BackRepoMaxLength.GetMaxLengthDBFromMaxLengthPtr(maxlengthInstance)
		ret = any(ret2).(*T2)
	case *models.MinInclusive:
		mininclusiveInstance := any(concreteInstance).(*models.MinInclusive)
		ret2 := backRepo.BackRepoMinInclusive.GetMinInclusiveDBFromMinInclusivePtr(mininclusiveInstance)
		ret = any(ret2).(*T2)
	case *models.MinLength:
		minlengthInstance := any(concreteInstance).(*models.MinLength)
		ret2 := backRepo.BackRepoMinLength.GetMinLengthDBFromMinLengthPtr(minlengthInstance)
		ret = any(ret2).(*T2)
	case *models.Pattern:
		patternInstance := any(concreteInstance).(*models.Pattern)
		ret2 := backRepo.BackRepoPattern.GetPatternDBFromPatternPtr(patternInstance)
		ret = any(ret2).(*T2)
	case *models.Restriction:
		restrictionInstance := any(concreteInstance).(*models.Restriction)
		ret2 := backRepo.BackRepoRestriction.GetRestrictionDBFromRestrictionPtr(restrictionInstance)
		ret = any(ret2).(*T2)
	case *models.Schema:
		schemaInstance := any(concreteInstance).(*models.Schema)
		ret2 := backRepo.BackRepoSchema.GetSchemaDBFromSchemaPtr(schemaInstance)
		ret = any(ret2).(*T2)
	case *models.Sequence:
		sequenceInstance := any(concreteInstance).(*models.Sequence)
		ret2 := backRepo.BackRepoSequence.GetSequenceDBFromSequencePtr(sequenceInstance)
		ret = any(ret2).(*T2)
	case *models.SimpleContent:
		simplecontentInstance := any(concreteInstance).(*models.SimpleContent)
		ret2 := backRepo.BackRepoSimpleContent.GetSimpleContentDBFromSimpleContentPtr(simplecontentInstance)
		ret = any(ret2).(*T2)
	case *models.SimpleType:
		simpletypeInstance := any(concreteInstance).(*models.SimpleType)
		ret2 := backRepo.BackRepoSimpleType.GetSimpleTypeDBFromSimpleTypePtr(simpletypeInstance)
		ret = any(ret2).(*T2)
	case *models.TotalDigit:
		totaldigitInstance := any(concreteInstance).(*models.TotalDigit)
		ret2 := backRepo.BackRepoTotalDigit.GetTotalDigitDBFromTotalDigitPtr(totaldigitInstance)
		ret = any(ret2).(*T2)
	case *models.Union:
		unionInstance := any(concreteInstance).(*models.Union)
		ret2 := backRepo.BackRepoUnion.GetUnionDBFromUnionPtr(unionInstance)
		ret = any(ret2).(*T2)
	case *models.WhiteSpace:
		whitespaceInstance := any(concreteInstance).(*models.WhiteSpace)
		ret2 := backRepo.BackRepoWhiteSpace.GetWhiteSpaceDBFromWhiteSpacePtr(whitespaceInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.All:
		tmp := GetInstanceDBFromInstance[models.All, AllDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Annotation:
		tmp := GetInstanceDBFromInstance[models.Annotation, AnnotationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Attribute:
		tmp := GetInstanceDBFromInstance[models.Attribute, AttributeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AttributeGroup:
		tmp := GetInstanceDBFromInstance[models.AttributeGroup, AttributeGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Choice:
		tmp := GetInstanceDBFromInstance[models.Choice, ChoiceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ComplexContent:
		tmp := GetInstanceDBFromInstance[models.ComplexContent, ComplexContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ComplexType:
		tmp := GetInstanceDBFromInstance[models.ComplexType, ComplexTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Documentation:
		tmp := GetInstanceDBFromInstance[models.Documentation, DocumentationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Element:
		tmp := GetInstanceDBFromInstance[models.Element, ElementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Enumeration:
		tmp := GetInstanceDBFromInstance[models.Enumeration, EnumerationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Extension:
		tmp := GetInstanceDBFromInstance[models.Extension, ExtensionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Length:
		tmp := GetInstanceDBFromInstance[models.Length, LengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MaxInclusive:
		tmp := GetInstanceDBFromInstance[models.MaxInclusive, MaxInclusiveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MaxLength:
		tmp := GetInstanceDBFromInstance[models.MaxLength, MaxLengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MinInclusive:
		tmp := GetInstanceDBFromInstance[models.MinInclusive, MinInclusiveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MinLength:
		tmp := GetInstanceDBFromInstance[models.MinLength, MinLengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pattern:
		tmp := GetInstanceDBFromInstance[models.Pattern, PatternDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Restriction:
		tmp := GetInstanceDBFromInstance[models.Restriction, RestrictionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Schema:
		tmp := GetInstanceDBFromInstance[models.Schema, SchemaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sequence:
		tmp := GetInstanceDBFromInstance[models.Sequence, SequenceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SimpleContent:
		tmp := GetInstanceDBFromInstance[models.SimpleContent, SimpleContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SimpleType:
		tmp := GetInstanceDBFromInstance[models.SimpleType, SimpleTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TotalDigit:
		tmp := GetInstanceDBFromInstance[models.TotalDigit, TotalDigitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Union:
		tmp := GetInstanceDBFromInstance[models.Union, UnionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.WhiteSpace:
		tmp := GetInstanceDBFromInstance[models.WhiteSpace, WhiteSpaceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.All:
		tmp := GetInstanceDBFromInstance[models.All, AllDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Annotation:
		tmp := GetInstanceDBFromInstance[models.Annotation, AnnotationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Attribute:
		tmp := GetInstanceDBFromInstance[models.Attribute, AttributeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AttributeGroup:
		tmp := GetInstanceDBFromInstance[models.AttributeGroup, AttributeGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Choice:
		tmp := GetInstanceDBFromInstance[models.Choice, ChoiceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ComplexContent:
		tmp := GetInstanceDBFromInstance[models.ComplexContent, ComplexContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ComplexType:
		tmp := GetInstanceDBFromInstance[models.ComplexType, ComplexTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Documentation:
		tmp := GetInstanceDBFromInstance[models.Documentation, DocumentationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Element:
		tmp := GetInstanceDBFromInstance[models.Element, ElementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Enumeration:
		tmp := GetInstanceDBFromInstance[models.Enumeration, EnumerationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Extension:
		tmp := GetInstanceDBFromInstance[models.Extension, ExtensionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Length:
		tmp := GetInstanceDBFromInstance[models.Length, LengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MaxInclusive:
		tmp := GetInstanceDBFromInstance[models.MaxInclusive, MaxInclusiveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MaxLength:
		tmp := GetInstanceDBFromInstance[models.MaxLength, MaxLengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MinInclusive:
		tmp := GetInstanceDBFromInstance[models.MinInclusive, MinInclusiveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.MinLength:
		tmp := GetInstanceDBFromInstance[models.MinLength, MinLengthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pattern:
		tmp := GetInstanceDBFromInstance[models.Pattern, PatternDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Restriction:
		tmp := GetInstanceDBFromInstance[models.Restriction, RestrictionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Schema:
		tmp := GetInstanceDBFromInstance[models.Schema, SchemaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sequence:
		tmp := GetInstanceDBFromInstance[models.Sequence, SequenceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SimpleContent:
		tmp := GetInstanceDBFromInstance[models.SimpleContent, SimpleContentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SimpleType:
		tmp := GetInstanceDBFromInstance[models.SimpleType, SimpleTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TotalDigit:
		tmp := GetInstanceDBFromInstance[models.TotalDigit, TotalDigitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Union:
		tmp := GetInstanceDBFromInstance[models.Union, UnionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.WhiteSpace:
		tmp := GetInstanceDBFromInstance[models.WhiteSpace, WhiteSpaceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
