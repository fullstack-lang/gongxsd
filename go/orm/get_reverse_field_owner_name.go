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
	case *models.All:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Attribute:
		switch reverseField.GongstructName {
		// insertion point
		case "AttributeGroup":
			switch reverseField.Fieldname {
			case "Attributes":
				if _attributegroup, ok := stage.AttributeGroup_Attributes_reverseMap[inst]; ok {
					res = _attributegroup.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Attributes":
				if _complextype, ok := stage.ComplexType_Attributes_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Attributes":
				if _extension, ok := stage.Extension_Attributes_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		}

	case *models.AttributeGroup:
		switch reverseField.GongstructName {
		// insertion point
		case "AttributeGroup":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				if _attributegroup, ok := stage.AttributeGroup_AttributeGroups_reverseMap[inst]; ok {
					res = _attributegroup.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				if _complextype, ok := stage.ComplexType_AttributeGroups_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				if _schema, ok := stage.Schema_AttributeGroups_reverseMap[inst]; ok {
					res = _schema.Name
				}
			}
		}

	case *models.Choice:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ComplexContent:
		switch reverseField.GongstructName {
		// insertion point
		}

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

	case *models.Extension:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "Groups":
				if _element, ok := stage.Element_Groups_reverseMap[inst]; ok {
					res = _element.Name
				}
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "Groups":
				if _schema, ok := stage.Schema_Groups_reverseMap[inst]; ok {
					res = _schema.Name
				}
			}
		}

	case *models.Length:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MaxInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MaxLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MinInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MinLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ModelGroupElement:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _all, ok := stage.All_ModelGroupElements_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _choice, ok := stage.Choice_ModelGroupElements_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _complextype, ok := stage.ComplexType_ModelGroupElements_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _extension, ok := stage.Extension_ModelGroupElements_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _group, ok := stage.Group_ModelGroupElements_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				if _sequence, ok := stage.Sequence_ModelGroupElements_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *models.Pattern:
		switch reverseField.GongstructName {
		// insertion point
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

	case *models.SimpleContent:
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

	case *models.TotalDigit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Union:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.WhiteSpace:
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
	case *models.All:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Attribute:
		switch reverseField.GongstructName {
		// insertion point
		case "AttributeGroup":
			switch reverseField.Fieldname {
			case "Attributes":
				res = stage.AttributeGroup_Attributes_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Attributes":
				res = stage.ComplexType_Attributes_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Attributes":
				res = stage.Extension_Attributes_reverseMap[inst]
			}
		}

	case *models.AttributeGroup:
		switch reverseField.GongstructName {
		// insertion point
		case "AttributeGroup":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				res = stage.AttributeGroup_AttributeGroups_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				res = stage.ComplexType_AttributeGroups_reverseMap[inst]
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				res = stage.Schema_AttributeGroups_reverseMap[inst]
			}
		}

	case *models.Choice:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ComplexContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ComplexType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "ComplexTypes":
				res = stage.Schema_ComplexTypes_reverseMap[inst]
			}
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

	case *models.Element:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Schema_Elements_reverseMap[inst]
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

	case *models.Extension:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Element":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Element_Groups_reverseMap[inst]
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Schema_Groups_reverseMap[inst]
			}
		}

	case *models.Length:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MaxInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MaxLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MinInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.MinLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ModelGroupElement:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.All_ModelGroupElements_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.Choice_ModelGroupElements_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.ComplexType_ModelGroupElements_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.Extension_ModelGroupElements_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.Group_ModelGroupElements_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "ModelGroupElements":
				res = stage.Sequence_ModelGroupElements_reverseMap[inst]
			}
		}

	case *models.Pattern:
		switch reverseField.GongstructName {
		// insertion point
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

	case *models.SimpleContent:
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

	case *models.TotalDigit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Union:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.WhiteSpace:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
