// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *All:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Alls":
				if _all, ok := stage.All_Alls_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Alls":
				if _choice, ok := stage.Choice_Alls_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Alls":
				if _complextype, ok := stage.ComplexType_Alls_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Alls":
				if _extension, ok := stage.Extension_Alls_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Alls":
				if _group, ok := stage.Group_Alls_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Alls":
				if _sequence, ok := stage.Sequence_Alls_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Attribute:
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

	case *AttributeGroup:
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
		case "Extension":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				if _extension, ok := stage.Extension_AttributeGroups_reverseMap[inst]; ok {
					res = _extension.Name
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

	case *Choice:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Choices":
				if _all, ok := stage.All_Choices_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Choices":
				if _choice, ok := stage.Choice_Choices_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Choices":
				if _complextype, ok := stage.ComplexType_Choices_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Choices":
				if _extension, ok := stage.Extension_Choices_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Choices":
				if _group, ok := stage.Group_Choices_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Choices":
				if _sequence, ok := stage.Sequence_Choices_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *ComplexContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ComplexType:
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

	case *Documentation:
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

	case *Element:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Elements":
				if _all, ok := stage.All_Elements_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Elements":
				if _choice, ok := stage.Choice_Elements_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Elements":
				if _complextype, ok := stage.ComplexType_Elements_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Elements":
				if _extension, ok := stage.Extension_Elements_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Elements":
				if _group, ok := stage.Group_Elements_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
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

	case *Enumeration:
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

	case *Extension:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Group:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Groups":
				if _all, ok := stage.All_Groups_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Groups":
				if _choice, ok := stage.Choice_Groups_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Groups":
				if _complextype, ok := stage.ComplexType_Groups_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Element":
			switch reverseField.Fieldname {
			case "Groups":
				if _element, ok := stage.Element_Groups_reverseMap[inst]; ok {
					res = _element.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Groups":
				if _extension, ok := stage.Extension_Groups_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Groups":
				if _group, ok := stage.Group_Groups_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "Groups":
				if _schema, ok := stage.Schema_Groups_reverseMap[inst]; ok {
					res = _schema.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Groups":
				if _sequence, ok := stage.Sequence_Groups_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *Length:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MaxInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MaxLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MinInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MinLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Pattern:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Restriction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Schema:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Sequence:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Sequences":
				if _all, ok := stage.All_Sequences_reverseMap[inst]; ok {
					res = _all.Name
				}
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Sequences":
				if _choice, ok := stage.Choice_Sequences_reverseMap[inst]; ok {
					res = _choice.Name
				}
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Sequences":
				if _complextype, ok := stage.ComplexType_Sequences_reverseMap[inst]; ok {
					res = _complextype.Name
				}
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Sequences":
				if _extension, ok := stage.Extension_Sequences_reverseMap[inst]; ok {
					res = _extension.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Sequences":
				if _group, ok := stage.Group_Sequences_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Sequences":
				if _sequence, ok := stage.Sequence_Sequences_reverseMap[inst]; ok {
					res = _sequence.Name
				}
			}
		}

	case *SimpleContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SimpleType:
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

	case *TotalDigit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Union:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *WhiteSpace:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *All:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.All_Alls_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.Choice_Alls_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.ComplexType_Alls_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.Extension_Alls_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.Group_Alls_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Alls":
				res = stage.Sequence_Alls_reverseMap[inst]
			}
		}

	case *Annotation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Attribute:
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

	case *AttributeGroup:
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
		case "Extension":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				res = stage.Extension_AttributeGroups_reverseMap[inst]
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "AttributeGroups":
				res = stage.Schema_AttributeGroups_reverseMap[inst]
			}
		}

	case *Choice:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.All_Choices_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.Choice_Choices_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.ComplexType_Choices_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.Extension_Choices_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.Group_Choices_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Choices":
				res = stage.Sequence_Choices_reverseMap[inst]
			}
		}

	case *ComplexContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ComplexType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "ComplexTypes":
				res = stage.Schema_ComplexTypes_reverseMap[inst]
			}
		}

	case *Documentation:
		switch reverseField.GongstructName {
		// insertion point
		case "Annotation":
			switch reverseField.Fieldname {
			case "Documentations":
				res = stage.Annotation_Documentations_reverseMap[inst]
			}
		}

	case *Element:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.All_Elements_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Choice_Elements_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.ComplexType_Elements_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Extension_Elements_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Elements":
				res = stage.Group_Elements_reverseMap[inst]
			}
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

	case *Enumeration:
		switch reverseField.GongstructName {
		// insertion point
		case "Restriction":
			switch reverseField.Fieldname {
			case "Enumerations":
				res = stage.Restriction_Enumerations_reverseMap[inst]
			}
		}

	case *Extension:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Group:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.All_Groups_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Choice_Groups_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.ComplexType_Groups_reverseMap[inst]
			}
		case "Element":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Element_Groups_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Extension_Groups_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Group_Groups_reverseMap[inst]
			}
		case "Schema":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Schema_Groups_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Sequence_Groups_reverseMap[inst]
			}
		}

	case *Length:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MaxInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MaxLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MinInclusive:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *MinLength:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Pattern:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Restriction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Schema:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Sequence:
		switch reverseField.GongstructName {
		// insertion point
		case "All":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.All_Sequences_reverseMap[inst]
			}
		case "Choice":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.Choice_Sequences_reverseMap[inst]
			}
		case "ComplexType":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.ComplexType_Sequences_reverseMap[inst]
			}
		case "Extension":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.Extension_Sequences_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.Group_Sequences_reverseMap[inst]
			}
		case "Sequence":
			switch reverseField.Fieldname {
			case "Sequences":
				res = stage.Sequence_Sequences_reverseMap[inst]
			}
		}

	case *SimpleContent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SimpleType:
		switch reverseField.GongstructName {
		// insertion point
		case "Schema":
			switch reverseField.Fieldname {
			case "SimpleTypes":
				res = stage.Schema_SimpleTypes_reverseMap[inst]
			}
		}

	case *TotalDigit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Union:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *WhiteSpace:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
