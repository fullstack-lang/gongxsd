// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.ComplexType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "ComplexTypes":
				if _schema, ok := stage.Schema_ComplexTypes_reverseMap[inst]; ok {
					res = _schema.Name
				}
			}
		}

	case *models.Element:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "Elements":
				if _schema, ok := stage.Schema_Elements_reverseMap[inst]; ok {
					res = _schema.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Elements":
				if _sequence, ok := stage.Sequence_Elements_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *models.Enumeration:
		switch reverseField.GongstructName {
		// insertion point
		case "Restriction":
			switch reverseField.Fieldname {
			case "Enumerations":
				if _restriction, ok := stage.Restriction_Enumerations_reverseMap[inst]; ok {
					res = _restriction.Name
				}
			}
		}

	case *models.Restriction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Schema:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sequence:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SimpleType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "SimpleTypes":
				if _schema, ok := stage.Schema_SimpleTypes_reverseMap[inst]; ok {
					res = _schema.Name
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
	case *models.ComplexType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "ComplexTypes":
				res = stage.Schema_ComplexTypes_reverseMap[inst]
			}
		}

	case *models.Element:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Schema_Elements_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Sequence_Elements_reverseMap[inst]
			}
		}

	case *models.Enumeration:
		switch reverseField.GongstructName {
		// insertion point
		case "Restriction":
			switch reverseField.Fieldname {
			case "Enumerations":
				res = stage.Restriction_Enumerations_reverseMap[inst]
			}
		}

	case *models.Restriction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Schema:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sequence:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.SimpleType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "SimpleTypes":
				res = stage.Schema_SimpleTypes_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
