// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/alt/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Documentation:
		switch reverseField.GongstructName {
		// insertion point
		case "Annotation":
			switch reverseField.Fieldname {
			case "Documentations":
				if _annotation, ok := stage.Annotation_Documentations_reverseMap[inst]; ok {
					res = _annotation.Name
				}
			}
		}

	case *models.Schema:
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
	case *models.Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Documentation:
		switch reverseField.GongstructName {
		// insertion point
		case "Annotation":
			switch reverseField.Fieldname {
			case "Documentations":
				res = stage.Annotation_Documentations_reverseMap[inst]
			}
		}

	case *models.Schema:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
