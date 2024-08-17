// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/alt/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Annotation:
		annotationInstance := any(concreteInstance).(*models.Annotation)
		ret2 := backRepo.BackRepoAnnotation.GetAnnotationDBFromAnnotationPtr(annotationInstance)
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
	case *models.Schema:
		schemaInstance := any(concreteInstance).(*models.Schema)
		ret2 := backRepo.BackRepoSchema.GetSchemaDBFromSchemaPtr(schemaInstance)
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
	case *models.Annotation:
		tmp := GetInstanceDBFromInstance[models.Annotation, AnnotationDB](
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
	case *models.Schema:
		tmp := GetInstanceDBFromInstance[models.Schema, SchemaDB](
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
	case *models.Annotation:
		tmp := GetInstanceDBFromInstance[models.Annotation, AnnotationDB](
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
	case *models.Schema:
		tmp := GetInstanceDBFromInstance[models.Schema, SchemaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
