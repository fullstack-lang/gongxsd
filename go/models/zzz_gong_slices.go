// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct All
	// insertion point per field
	stage.All_Sequences_reverseMap = make(map[*Sequence]*All)
	for all := range stage.Alls {
		_ = all
		for _, _sequence := range all.Sequences {
			stage.All_Sequences_reverseMap[_sequence] = all
		}
	}
	stage.All_Alls_reverseMap = make(map[*All]*All)
	for all := range stage.Alls {
		_ = all
		for _, _all := range all.Alls {
			stage.All_Alls_reverseMap[_all] = all
		}
	}
	stage.All_Choices_reverseMap = make(map[*Choice]*All)
	for all := range stage.Alls {
		_ = all
		for _, _choice := range all.Choices {
			stage.All_Choices_reverseMap[_choice] = all
		}
	}
	stage.All_Groups_reverseMap = make(map[*Group]*All)
	for all := range stage.Alls {
		_ = all
		for _, _group := range all.Groups {
			stage.All_Groups_reverseMap[_group] = all
		}
	}
	stage.All_Elements_reverseMap = make(map[*Element]*All)
	for all := range stage.Alls {
		_ = all
		for _, _element := range all.Elements {
			stage.All_Elements_reverseMap[_element] = all
		}
	}

	// Compute reverse map for named struct Annotation
	// insertion point per field
	stage.Annotation_Documentations_reverseMap = make(map[*Documentation]*Annotation)
	for annotation := range stage.Annotations {
		_ = annotation
		for _, _documentation := range annotation.Documentations {
			stage.Annotation_Documentations_reverseMap[_documentation] = annotation
		}
	}

	// Compute reverse map for named struct Attribute
	// insertion point per field

	// Compute reverse map for named struct AttributeGroup
	// insertion point per field
	stage.AttributeGroup_AttributeGroups_reverseMap = make(map[*AttributeGroup]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attributegroup := range attributegroup.AttributeGroups {
			stage.AttributeGroup_AttributeGroups_reverseMap[_attributegroup] = attributegroup
		}
	}
	stage.AttributeGroup_Attributes_reverseMap = make(map[*Attribute]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attribute := range attributegroup.Attributes {
			stage.AttributeGroup_Attributes_reverseMap[_attribute] = attributegroup
		}
	}

	// Compute reverse map for named struct Choice
	// insertion point per field
	stage.Choice_Sequences_reverseMap = make(map[*Sequence]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _sequence := range choice.Sequences {
			stage.Choice_Sequences_reverseMap[_sequence] = choice
		}
	}
	stage.Choice_Alls_reverseMap = make(map[*All]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _all := range choice.Alls {
			stage.Choice_Alls_reverseMap[_all] = choice
		}
	}
	stage.Choice_Choices_reverseMap = make(map[*Choice]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _choice := range choice.Choices {
			stage.Choice_Choices_reverseMap[_choice] = choice
		}
	}
	stage.Choice_Groups_reverseMap = make(map[*Group]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _group := range choice.Groups {
			stage.Choice_Groups_reverseMap[_group] = choice
		}
	}
	stage.Choice_Elements_reverseMap = make(map[*Element]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _element := range choice.Elements {
			stage.Choice_Elements_reverseMap[_element] = choice
		}
	}

	// Compute reverse map for named struct ComplexContent
	// insertion point per field

	// Compute reverse map for named struct ComplexType
	// insertion point per field
	stage.ComplexType_Sequences_reverseMap = make(map[*Sequence]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _sequence := range complextype.Sequences {
			stage.ComplexType_Sequences_reverseMap[_sequence] = complextype
		}
	}
	stage.ComplexType_Alls_reverseMap = make(map[*All]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _all := range complextype.Alls {
			stage.ComplexType_Alls_reverseMap[_all] = complextype
		}
	}
	stage.ComplexType_Choices_reverseMap = make(map[*Choice]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _choice := range complextype.Choices {
			stage.ComplexType_Choices_reverseMap[_choice] = complextype
		}
	}
	stage.ComplexType_Groups_reverseMap = make(map[*Group]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _group := range complextype.Groups {
			stage.ComplexType_Groups_reverseMap[_group] = complextype
		}
	}
	stage.ComplexType_Elements_reverseMap = make(map[*Element]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _element := range complextype.Elements {
			stage.ComplexType_Elements_reverseMap[_element] = complextype
		}
	}
	stage.ComplexType_Attributes_reverseMap = make(map[*Attribute]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attribute := range complextype.Attributes {
			stage.ComplexType_Attributes_reverseMap[_attribute] = complextype
		}
	}
	stage.ComplexType_AttributeGroups_reverseMap = make(map[*AttributeGroup]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attributegroup := range complextype.AttributeGroups {
			stage.ComplexType_AttributeGroups_reverseMap[_attributegroup] = complextype
		}
	}

	// Compute reverse map for named struct Documentation
	// insertion point per field

	// Compute reverse map for named struct Element
	// insertion point per field
	stage.Element_Groups_reverseMap = make(map[*Group]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _group := range element.Groups {
			stage.Element_Groups_reverseMap[_group] = element
		}
	}

	// Compute reverse map for named struct Enumeration
	// insertion point per field

	// Compute reverse map for named struct Extension
	// insertion point per field
	stage.Extension_Sequences_reverseMap = make(map[*Sequence]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _sequence := range extension.Sequences {
			stage.Extension_Sequences_reverseMap[_sequence] = extension
		}
	}
	stage.Extension_Alls_reverseMap = make(map[*All]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _all := range extension.Alls {
			stage.Extension_Alls_reverseMap[_all] = extension
		}
	}
	stage.Extension_Choices_reverseMap = make(map[*Choice]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _choice := range extension.Choices {
			stage.Extension_Choices_reverseMap[_choice] = extension
		}
	}
	stage.Extension_Groups_reverseMap = make(map[*Group]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _group := range extension.Groups {
			stage.Extension_Groups_reverseMap[_group] = extension
		}
	}
	stage.Extension_Elements_reverseMap = make(map[*Element]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _element := range extension.Elements {
			stage.Extension_Elements_reverseMap[_element] = extension
		}
	}
	stage.Extension_Attributes_reverseMap = make(map[*Attribute]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _attribute := range extension.Attributes {
			stage.Extension_Attributes_reverseMap[_attribute] = extension
		}
	}
	stage.Extension_AttributeGroups_reverseMap = make(map[*AttributeGroup]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _attributegroup := range extension.AttributeGroups {
			stage.Extension_AttributeGroups_reverseMap[_attributegroup] = extension
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Sequences_reverseMap = make(map[*Sequence]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _sequence := range group.Sequences {
			stage.Group_Sequences_reverseMap[_sequence] = group
		}
	}
	stage.Group_Alls_reverseMap = make(map[*All]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _all := range group.Alls {
			stage.Group_Alls_reverseMap[_all] = group
		}
	}
	stage.Group_Choices_reverseMap = make(map[*Choice]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _choice := range group.Choices {
			stage.Group_Choices_reverseMap[_choice] = group
		}
	}
	stage.Group_Groups_reverseMap = make(map[*Group]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _group := range group.Groups {
			stage.Group_Groups_reverseMap[_group] = group
		}
	}
	stage.Group_Elements_reverseMap = make(map[*Element]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _element := range group.Elements {
			stage.Group_Elements_reverseMap[_element] = group
		}
	}

	// Compute reverse map for named struct Length
	// insertion point per field

	// Compute reverse map for named struct MaxInclusive
	// insertion point per field

	// Compute reverse map for named struct MaxLength
	// insertion point per field

	// Compute reverse map for named struct MinInclusive
	// insertion point per field

	// Compute reverse map for named struct MinLength
	// insertion point per field

	// Compute reverse map for named struct Pattern
	// insertion point per field

	// Compute reverse map for named struct Restriction
	// insertion point per field
	stage.Restriction_Enumerations_reverseMap = make(map[*Enumeration]*Restriction)
	for restriction := range stage.Restrictions {
		_ = restriction
		for _, _enumeration := range restriction.Enumerations {
			stage.Restriction_Enumerations_reverseMap[_enumeration] = restriction
		}
	}

	// Compute reverse map for named struct Schema
	// insertion point per field
	stage.Schema_Elements_reverseMap = make(map[*Element]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _element := range schema.Elements {
			stage.Schema_Elements_reverseMap[_element] = schema
		}
	}
	stage.Schema_SimpleTypes_reverseMap = make(map[*SimpleType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _simpletype := range schema.SimpleTypes {
			stage.Schema_SimpleTypes_reverseMap[_simpletype] = schema
		}
	}
	stage.Schema_ComplexTypes_reverseMap = make(map[*ComplexType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _complextype := range schema.ComplexTypes {
			stage.Schema_ComplexTypes_reverseMap[_complextype] = schema
		}
	}
	stage.Schema_AttributeGroups_reverseMap = make(map[*AttributeGroup]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _attributegroup := range schema.AttributeGroups {
			stage.Schema_AttributeGroups_reverseMap[_attributegroup] = schema
		}
	}
	stage.Schema_Groups_reverseMap = make(map[*Group]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _group := range schema.Groups {
			stage.Schema_Groups_reverseMap[_group] = schema
		}
	}

	// Compute reverse map for named struct Sequence
	// insertion point per field
	stage.Sequence_Sequences_reverseMap = make(map[*Sequence]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _sequence := range sequence.Sequences {
			stage.Sequence_Sequences_reverseMap[_sequence] = sequence
		}
	}
	stage.Sequence_Alls_reverseMap = make(map[*All]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _all := range sequence.Alls {
			stage.Sequence_Alls_reverseMap[_all] = sequence
		}
	}
	stage.Sequence_Choices_reverseMap = make(map[*Choice]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _choice := range sequence.Choices {
			stage.Sequence_Choices_reverseMap[_choice] = sequence
		}
	}
	stage.Sequence_Groups_reverseMap = make(map[*Group]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _group := range sequence.Groups {
			stage.Sequence_Groups_reverseMap[_group] = sequence
		}
	}
	stage.Sequence_Elements_reverseMap = make(map[*Element]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _element := range sequence.Elements {
			stage.Sequence_Elements_reverseMap[_element] = sequence
		}
	}

	// Compute reverse map for named struct SimpleContent
	// insertion point per field

	// Compute reverse map for named struct SimpleType
	// insertion point per field

	// Compute reverse map for named struct TotalDigit
	// insertion point per field

	// Compute reverse map for named struct Union
	// insertion point per field

	// Compute reverse map for named struct WhiteSpace
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Alls {
		res = append(res, instance)
	}

	for instance := range stage.Annotations {
		res = append(res, instance)
	}

	for instance := range stage.Attributes {
		res = append(res, instance)
	}

	for instance := range stage.AttributeGroups {
		res = append(res, instance)
	}

	for instance := range stage.Choices {
		res = append(res, instance)
	}

	for instance := range stage.ComplexContents {
		res = append(res, instance)
	}

	for instance := range stage.ComplexTypes {
		res = append(res, instance)
	}

	for instance := range stage.Documentations {
		res = append(res, instance)
	}

	for instance := range stage.Elements {
		res = append(res, instance)
	}

	for instance := range stage.Enumerations {
		res = append(res, instance)
	}

	for instance := range stage.Extensions {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Lengths {
		res = append(res, instance)
	}

	for instance := range stage.MaxInclusives {
		res = append(res, instance)
	}

	for instance := range stage.MaxLengths {
		res = append(res, instance)
	}

	for instance := range stage.MinInclusives {
		res = append(res, instance)
	}

	for instance := range stage.MinLengths {
		res = append(res, instance)
	}

	for instance := range stage.Patterns {
		res = append(res, instance)
	}

	for instance := range stage.Restrictions {
		res = append(res, instance)
	}

	for instance := range stage.Schemas {
		res = append(res, instance)
	}

	for instance := range stage.Sequences {
		res = append(res, instance)
	}

	for instance := range stage.SimpleContents {
		res = append(res, instance)
	}

	for instance := range stage.SimpleTypes {
		res = append(res, instance)
	}

	for instance := range stage.TotalDigits {
		res = append(res, instance)
	}

	for instance := range stage.Unions {
		res = append(res, instance)
	}

	for instance := range stage.WhiteSpaces {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (all *All) GongCopy() GongstructIF {
	newInstance := *all
	return &newInstance
}

func (annotation *Annotation) GongCopy() GongstructIF {
	newInstance := *annotation
	return &newInstance
}

func (attribute *Attribute) GongCopy() GongstructIF {
	newInstance := *attribute
	return &newInstance
}

func (attributegroup *AttributeGroup) GongCopy() GongstructIF {
	newInstance := *attributegroup
	return &newInstance
}

func (choice *Choice) GongCopy() GongstructIF {
	newInstance := *choice
	return &newInstance
}

func (complexcontent *ComplexContent) GongCopy() GongstructIF {
	newInstance := *complexcontent
	return &newInstance
}

func (complextype *ComplexType) GongCopy() GongstructIF {
	newInstance := *complextype
	return &newInstance
}

func (documentation *Documentation) GongCopy() GongstructIF {
	newInstance := *documentation
	return &newInstance
}

func (element *Element) GongCopy() GongstructIF {
	newInstance := *element
	return &newInstance
}

func (enumeration *Enumeration) GongCopy() GongstructIF {
	newInstance := *enumeration
	return &newInstance
}

func (extension *Extension) GongCopy() GongstructIF {
	newInstance := *extension
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := *group
	return &newInstance
}

func (length *Length) GongCopy() GongstructIF {
	newInstance := *length
	return &newInstance
}

func (maxinclusive *MaxInclusive) GongCopy() GongstructIF {
	newInstance := *maxinclusive
	return &newInstance
}

func (maxlength *MaxLength) GongCopy() GongstructIF {
	newInstance := *maxlength
	return &newInstance
}

func (mininclusive *MinInclusive) GongCopy() GongstructIF {
	newInstance := *mininclusive
	return &newInstance
}

func (minlength *MinLength) GongCopy() GongstructIF {
	newInstance := *minlength
	return &newInstance
}

func (pattern *Pattern) GongCopy() GongstructIF {
	newInstance := *pattern
	return &newInstance
}

func (restriction *Restriction) GongCopy() GongstructIF {
	newInstance := *restriction
	return &newInstance
}

func (schema *Schema) GongCopy() GongstructIF {
	newInstance := *schema
	return &newInstance
}

func (sequence *Sequence) GongCopy() GongstructIF {
	newInstance := *sequence
	return &newInstance
}

func (simplecontent *SimpleContent) GongCopy() GongstructIF {
	newInstance := *simplecontent
	return &newInstance
}

func (simpletype *SimpleType) GongCopy() GongstructIF {
	newInstance := *simpletype
	return &newInstance
}

func (totaldigit *TotalDigit) GongCopy() GongstructIF {
	newInstance := *totaldigit
	return &newInstance
}

func (union *Union) GongCopy() GongstructIF {
	newInstance := *union
	return &newInstance
}

func (whitespace *WhiteSpace) GongCopy() GongstructIF {
	newInstance := *whitespace
	return &newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var alls_newInstances []*All
	var alls_deletedInstances []*All

	// parse all staged instances and check if they have a reference
	for all := range stage.Alls {
		if ref, ok := stage.Alls_reference[all]; !ok {
			alls_newInstances = append(alls_newInstances, all)
			newInstancesSlice = append(newInstancesSlice, all.GongMarshallIdentifier(stage))
			if stage.Alls_referenceOrder == nil {
				stage.Alls_referenceOrder = make(map[*All]uint)
			}
			stage.Alls_referenceOrder[all] = stage.AllMap_Staged_Order[all]
			newInstancesReverseSlice = append(newInstancesReverseSlice, all.GongMarshallUnstaging(stage))
			delete(stage.Alls_referenceOrder, all)
			fieldInitializers, pointersInitializations := all.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AllMap_Staged_Order[ref] = stage.AllMap_Staged_Order[all]
			diffs := all.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, all)
			delete(stage.AllMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", all.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Alls_reference {
		if _, ok := stage.Alls[ref]; !ok {
			alls_deletedInstances = append(alls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(alls_newInstances)
	lenDeletedInstances += len(alls_deletedInstances)
	var annotations_newInstances []*Annotation
	var annotations_deletedInstances []*Annotation

	// parse all staged instances and check if they have a reference
	for annotation := range stage.Annotations {
		if ref, ok := stage.Annotations_reference[annotation]; !ok {
			annotations_newInstances = append(annotations_newInstances, annotation)
			newInstancesSlice = append(newInstancesSlice, annotation.GongMarshallIdentifier(stage))
			if stage.Annotations_referenceOrder == nil {
				stage.Annotations_referenceOrder = make(map[*Annotation]uint)
			}
			stage.Annotations_referenceOrder[annotation] = stage.AnnotationMap_Staged_Order[annotation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, annotation.GongMarshallUnstaging(stage))
			delete(stage.Annotations_referenceOrder, annotation)
			fieldInitializers, pointersInitializations := annotation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AnnotationMap_Staged_Order[ref] = stage.AnnotationMap_Staged_Order[annotation]
			diffs := annotation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, annotation)
			delete(stage.AnnotationMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", annotation.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Annotations_reference {
		if _, ok := stage.Annotations[ref]; !ok {
			annotations_deletedInstances = append(annotations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(annotations_newInstances)
	lenDeletedInstances += len(annotations_deletedInstances)
	var attributes_newInstances []*Attribute
	var attributes_deletedInstances []*Attribute

	// parse all staged instances and check if they have a reference
	for attribute := range stage.Attributes {
		if ref, ok := stage.Attributes_reference[attribute]; !ok {
			attributes_newInstances = append(attributes_newInstances, attribute)
			newInstancesSlice = append(newInstancesSlice, attribute.GongMarshallIdentifier(stage))
			if stage.Attributes_referenceOrder == nil {
				stage.Attributes_referenceOrder = make(map[*Attribute]uint)
			}
			stage.Attributes_referenceOrder[attribute] = stage.AttributeMap_Staged_Order[attribute]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute.GongMarshallUnstaging(stage))
			delete(stage.Attributes_referenceOrder, attribute)
			fieldInitializers, pointersInitializations := attribute.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AttributeMap_Staged_Order[ref] = stage.AttributeMap_Staged_Order[attribute]
			diffs := attribute.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute)
			delete(stage.AttributeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Attributes_reference {
		if _, ok := stage.Attributes[ref]; !ok {
			attributes_deletedInstances = append(attributes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attributes_newInstances)
	lenDeletedInstances += len(attributes_deletedInstances)
	var attributegroups_newInstances []*AttributeGroup
	var attributegroups_deletedInstances []*AttributeGroup

	// parse all staged instances and check if they have a reference
	for attributegroup := range stage.AttributeGroups {
		if ref, ok := stage.AttributeGroups_reference[attributegroup]; !ok {
			attributegroups_newInstances = append(attributegroups_newInstances, attributegroup)
			newInstancesSlice = append(newInstancesSlice, attributegroup.GongMarshallIdentifier(stage))
			if stage.AttributeGroups_referenceOrder == nil {
				stage.AttributeGroups_referenceOrder = make(map[*AttributeGroup]uint)
			}
			stage.AttributeGroups_referenceOrder[attributegroup] = stage.AttributeGroupMap_Staged_Order[attributegroup]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attributegroup.GongMarshallUnstaging(stage))
			delete(stage.AttributeGroups_referenceOrder, attributegroup)
			fieldInitializers, pointersInitializations := attributegroup.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AttributeGroupMap_Staged_Order[ref] = stage.AttributeGroupMap_Staged_Order[attributegroup]
			diffs := attributegroup.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attributegroup)
			delete(stage.AttributeGroupMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attributegroup.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.AttributeGroups_reference {
		if _, ok := stage.AttributeGroups[ref]; !ok {
			attributegroups_deletedInstances = append(attributegroups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attributegroups_newInstances)
	lenDeletedInstances += len(attributegroups_deletedInstances)
	var choices_newInstances []*Choice
	var choices_deletedInstances []*Choice

	// parse all staged instances and check if they have a reference
	for choice := range stage.Choices {
		if ref, ok := stage.Choices_reference[choice]; !ok {
			choices_newInstances = append(choices_newInstances, choice)
			newInstancesSlice = append(newInstancesSlice, choice.GongMarshallIdentifier(stage))
			if stage.Choices_referenceOrder == nil {
				stage.Choices_referenceOrder = make(map[*Choice]uint)
			}
			stage.Choices_referenceOrder[choice] = stage.ChoiceMap_Staged_Order[choice]
			newInstancesReverseSlice = append(newInstancesReverseSlice, choice.GongMarshallUnstaging(stage))
			delete(stage.Choices_referenceOrder, choice)
			fieldInitializers, pointersInitializations := choice.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ChoiceMap_Staged_Order[ref] = stage.ChoiceMap_Staged_Order[choice]
			diffs := choice.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, choice)
			delete(stage.ChoiceMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", choice.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Choices_reference {
		if _, ok := stage.Choices[ref]; !ok {
			choices_deletedInstances = append(choices_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(choices_newInstances)
	lenDeletedInstances += len(choices_deletedInstances)
	var complexcontents_newInstances []*ComplexContent
	var complexcontents_deletedInstances []*ComplexContent

	// parse all staged instances and check if they have a reference
	for complexcontent := range stage.ComplexContents {
		if ref, ok := stage.ComplexContents_reference[complexcontent]; !ok {
			complexcontents_newInstances = append(complexcontents_newInstances, complexcontent)
			newInstancesSlice = append(newInstancesSlice, complexcontent.GongMarshallIdentifier(stage))
			if stage.ComplexContents_referenceOrder == nil {
				stage.ComplexContents_referenceOrder = make(map[*ComplexContent]uint)
			}
			stage.ComplexContents_referenceOrder[complexcontent] = stage.ComplexContentMap_Staged_Order[complexcontent]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complexcontent.GongMarshallUnstaging(stage))
			delete(stage.ComplexContents_referenceOrder, complexcontent)
			fieldInitializers, pointersInitializations := complexcontent.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ComplexContentMap_Staged_Order[ref] = stage.ComplexContentMap_Staged_Order[complexcontent]
			diffs := complexcontent.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complexcontent)
			delete(stage.ComplexContentMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", complexcontent.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ComplexContents_reference {
		if _, ok := stage.ComplexContents[ref]; !ok {
			complexcontents_deletedInstances = append(complexcontents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complexcontents_newInstances)
	lenDeletedInstances += len(complexcontents_deletedInstances)
	var complextypes_newInstances []*ComplexType
	var complextypes_deletedInstances []*ComplexType

	// parse all staged instances and check if they have a reference
	for complextype := range stage.ComplexTypes {
		if ref, ok := stage.ComplexTypes_reference[complextype]; !ok {
			complextypes_newInstances = append(complextypes_newInstances, complextype)
			newInstancesSlice = append(newInstancesSlice, complextype.GongMarshallIdentifier(stage))
			if stage.ComplexTypes_referenceOrder == nil {
				stage.ComplexTypes_referenceOrder = make(map[*ComplexType]uint)
			}
			stage.ComplexTypes_referenceOrder[complextype] = stage.ComplexTypeMap_Staged_Order[complextype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complextype.GongMarshallUnstaging(stage))
			delete(stage.ComplexTypes_referenceOrder, complextype)
			fieldInitializers, pointersInitializations := complextype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ComplexTypeMap_Staged_Order[ref] = stage.ComplexTypeMap_Staged_Order[complextype]
			diffs := complextype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complextype)
			delete(stage.ComplexTypeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", complextype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ComplexTypes_reference {
		if _, ok := stage.ComplexTypes[ref]; !ok {
			complextypes_deletedInstances = append(complextypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complextypes_newInstances)
	lenDeletedInstances += len(complextypes_deletedInstances)
	var documentations_newInstances []*Documentation
	var documentations_deletedInstances []*Documentation

	// parse all staged instances and check if they have a reference
	for documentation := range stage.Documentations {
		if ref, ok := stage.Documentations_reference[documentation]; !ok {
			documentations_newInstances = append(documentations_newInstances, documentation)
			newInstancesSlice = append(newInstancesSlice, documentation.GongMarshallIdentifier(stage))
			if stage.Documentations_referenceOrder == nil {
				stage.Documentations_referenceOrder = make(map[*Documentation]uint)
			}
			stage.Documentations_referenceOrder[documentation] = stage.DocumentationMap_Staged_Order[documentation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, documentation.GongMarshallUnstaging(stage))
			delete(stage.Documentations_referenceOrder, documentation)
			fieldInitializers, pointersInitializations := documentation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DocumentationMap_Staged_Order[ref] = stage.DocumentationMap_Staged_Order[documentation]
			diffs := documentation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, documentation)
			delete(stage.DocumentationMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", documentation.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Documentations_reference {
		if _, ok := stage.Documentations[ref]; !ok {
			documentations_deletedInstances = append(documentations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(documentations_newInstances)
	lenDeletedInstances += len(documentations_deletedInstances)
	var elements_newInstances []*Element
	var elements_deletedInstances []*Element

	// parse all staged instances and check if they have a reference
	for element := range stage.Elements {
		if ref, ok := stage.Elements_reference[element]; !ok {
			elements_newInstances = append(elements_newInstances, element)
			newInstancesSlice = append(newInstancesSlice, element.GongMarshallIdentifier(stage))
			if stage.Elements_referenceOrder == nil {
				stage.Elements_referenceOrder = make(map[*Element]uint)
			}
			stage.Elements_referenceOrder[element] = stage.ElementMap_Staged_Order[element]
			newInstancesReverseSlice = append(newInstancesReverseSlice, element.GongMarshallUnstaging(stage))
			delete(stage.Elements_referenceOrder, element)
			fieldInitializers, pointersInitializations := element.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ElementMap_Staged_Order[ref] = stage.ElementMap_Staged_Order[element]
			diffs := element.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, element)
			delete(stage.ElementMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", element.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Elements_reference {
		if _, ok := stage.Elements[ref]; !ok {
			elements_deletedInstances = append(elements_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(elements_newInstances)
	lenDeletedInstances += len(elements_deletedInstances)
	var enumerations_newInstances []*Enumeration
	var enumerations_deletedInstances []*Enumeration

	// parse all staged instances and check if they have a reference
	for enumeration := range stage.Enumerations {
		if ref, ok := stage.Enumerations_reference[enumeration]; !ok {
			enumerations_newInstances = append(enumerations_newInstances, enumeration)
			newInstancesSlice = append(newInstancesSlice, enumeration.GongMarshallIdentifier(stage))
			if stage.Enumerations_referenceOrder == nil {
				stage.Enumerations_referenceOrder = make(map[*Enumeration]uint)
			}
			stage.Enumerations_referenceOrder[enumeration] = stage.EnumerationMap_Staged_Order[enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, enumeration.GongMarshallUnstaging(stage))
			delete(stage.Enumerations_referenceOrder, enumeration)
			fieldInitializers, pointersInitializations := enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EnumerationMap_Staged_Order[ref] = stage.EnumerationMap_Staged_Order[enumeration]
			diffs := enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, enumeration)
			delete(stage.EnumerationMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", enumeration.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Enumerations_reference {
		if _, ok := stage.Enumerations[ref]; !ok {
			enumerations_deletedInstances = append(enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(enumerations_newInstances)
	lenDeletedInstances += len(enumerations_deletedInstances)
	var extensions_newInstances []*Extension
	var extensions_deletedInstances []*Extension

	// parse all staged instances and check if they have a reference
	for extension := range stage.Extensions {
		if ref, ok := stage.Extensions_reference[extension]; !ok {
			extensions_newInstances = append(extensions_newInstances, extension)
			newInstancesSlice = append(newInstancesSlice, extension.GongMarshallIdentifier(stage))
			if stage.Extensions_referenceOrder == nil {
				stage.Extensions_referenceOrder = make(map[*Extension]uint)
			}
			stage.Extensions_referenceOrder[extension] = stage.ExtensionMap_Staged_Order[extension]
			newInstancesReverseSlice = append(newInstancesReverseSlice, extension.GongMarshallUnstaging(stage))
			delete(stage.Extensions_referenceOrder, extension)
			fieldInitializers, pointersInitializations := extension.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExtensionMap_Staged_Order[ref] = stage.ExtensionMap_Staged_Order[extension]
			diffs := extension.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, extension)
			delete(stage.ExtensionMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", extension.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Extensions_reference {
		if _, ok := stage.Extensions[ref]; !ok {
			extensions_deletedInstances = append(extensions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(extensions_newInstances)
	lenDeletedInstances += len(extensions_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			newInstancesSlice = append(newInstancesSlice, group.GongMarshallIdentifier(stage))
			if stage.Groups_referenceOrder == nil {
				stage.Groups_referenceOrder = make(map[*Group]uint)
			}
			stage.Groups_referenceOrder[group] = stage.GroupMap_Staged_Order[group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, group.GongMarshallUnstaging(stage))
			delete(stage.Groups_referenceOrder, group)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GroupMap_Staged_Order[ref] = stage.GroupMap_Staged_Order[group]
			diffs := group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, group)
			delete(stage.GroupMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", group.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Groups_reference {
		if _, ok := stage.Groups[ref]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var lengths_newInstances []*Length
	var lengths_deletedInstances []*Length

	// parse all staged instances and check if they have a reference
	for length := range stage.Lengths {
		if ref, ok := stage.Lengths_reference[length]; !ok {
			lengths_newInstances = append(lengths_newInstances, length)
			newInstancesSlice = append(newInstancesSlice, length.GongMarshallIdentifier(stage))
			if stage.Lengths_referenceOrder == nil {
				stage.Lengths_referenceOrder = make(map[*Length]uint)
			}
			stage.Lengths_referenceOrder[length] = stage.LengthMap_Staged_Order[length]
			newInstancesReverseSlice = append(newInstancesReverseSlice, length.GongMarshallUnstaging(stage))
			delete(stage.Lengths_referenceOrder, length)
			fieldInitializers, pointersInitializations := length.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LengthMap_Staged_Order[ref] = stage.LengthMap_Staged_Order[length]
			diffs := length.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, length)
			delete(stage.LengthMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", length.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Lengths_reference {
		if _, ok := stage.Lengths[ref]; !ok {
			lengths_deletedInstances = append(lengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(lengths_newInstances)
	lenDeletedInstances += len(lengths_deletedInstances)
	var maxinclusives_newInstances []*MaxInclusive
	var maxinclusives_deletedInstances []*MaxInclusive

	// parse all staged instances and check if they have a reference
	for maxinclusive := range stage.MaxInclusives {
		if ref, ok := stage.MaxInclusives_reference[maxinclusive]; !ok {
			maxinclusives_newInstances = append(maxinclusives_newInstances, maxinclusive)
			newInstancesSlice = append(newInstancesSlice, maxinclusive.GongMarshallIdentifier(stage))
			if stage.MaxInclusives_referenceOrder == nil {
				stage.MaxInclusives_referenceOrder = make(map[*MaxInclusive]uint)
			}
			stage.MaxInclusives_referenceOrder[maxinclusive] = stage.MaxInclusiveMap_Staged_Order[maxinclusive]
			newInstancesReverseSlice = append(newInstancesReverseSlice, maxinclusive.GongMarshallUnstaging(stage))
			delete(stage.MaxInclusives_referenceOrder, maxinclusive)
			fieldInitializers, pointersInitializations := maxinclusive.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MaxInclusiveMap_Staged_Order[ref] = stage.MaxInclusiveMap_Staged_Order[maxinclusive]
			diffs := maxinclusive.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, maxinclusive)
			delete(stage.MaxInclusiveMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", maxinclusive.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.MaxInclusives_reference {
		if _, ok := stage.MaxInclusives[ref]; !ok {
			maxinclusives_deletedInstances = append(maxinclusives_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(maxinclusives_newInstances)
	lenDeletedInstances += len(maxinclusives_deletedInstances)
	var maxlengths_newInstances []*MaxLength
	var maxlengths_deletedInstances []*MaxLength

	// parse all staged instances and check if they have a reference
	for maxlength := range stage.MaxLengths {
		if ref, ok := stage.MaxLengths_reference[maxlength]; !ok {
			maxlengths_newInstances = append(maxlengths_newInstances, maxlength)
			newInstancesSlice = append(newInstancesSlice, maxlength.GongMarshallIdentifier(stage))
			if stage.MaxLengths_referenceOrder == nil {
				stage.MaxLengths_referenceOrder = make(map[*MaxLength]uint)
			}
			stage.MaxLengths_referenceOrder[maxlength] = stage.MaxLengthMap_Staged_Order[maxlength]
			newInstancesReverseSlice = append(newInstancesReverseSlice, maxlength.GongMarshallUnstaging(stage))
			delete(stage.MaxLengths_referenceOrder, maxlength)
			fieldInitializers, pointersInitializations := maxlength.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MaxLengthMap_Staged_Order[ref] = stage.MaxLengthMap_Staged_Order[maxlength]
			diffs := maxlength.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, maxlength)
			delete(stage.MaxLengthMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", maxlength.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.MaxLengths_reference {
		if _, ok := stage.MaxLengths[ref]; !ok {
			maxlengths_deletedInstances = append(maxlengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(maxlengths_newInstances)
	lenDeletedInstances += len(maxlengths_deletedInstances)
	var mininclusives_newInstances []*MinInclusive
	var mininclusives_deletedInstances []*MinInclusive

	// parse all staged instances and check if they have a reference
	for mininclusive := range stage.MinInclusives {
		if ref, ok := stage.MinInclusives_reference[mininclusive]; !ok {
			mininclusives_newInstances = append(mininclusives_newInstances, mininclusive)
			newInstancesSlice = append(newInstancesSlice, mininclusive.GongMarshallIdentifier(stage))
			if stage.MinInclusives_referenceOrder == nil {
				stage.MinInclusives_referenceOrder = make(map[*MinInclusive]uint)
			}
			stage.MinInclusives_referenceOrder[mininclusive] = stage.MinInclusiveMap_Staged_Order[mininclusive]
			newInstancesReverseSlice = append(newInstancesReverseSlice, mininclusive.GongMarshallUnstaging(stage))
			delete(stage.MinInclusives_referenceOrder, mininclusive)
			fieldInitializers, pointersInitializations := mininclusive.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MinInclusiveMap_Staged_Order[ref] = stage.MinInclusiveMap_Staged_Order[mininclusive]
			diffs := mininclusive.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, mininclusive)
			delete(stage.MinInclusiveMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", mininclusive.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.MinInclusives_reference {
		if _, ok := stage.MinInclusives[ref]; !ok {
			mininclusives_deletedInstances = append(mininclusives_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(mininclusives_newInstances)
	lenDeletedInstances += len(mininclusives_deletedInstances)
	var minlengths_newInstances []*MinLength
	var minlengths_deletedInstances []*MinLength

	// parse all staged instances and check if they have a reference
	for minlength := range stage.MinLengths {
		if ref, ok := stage.MinLengths_reference[minlength]; !ok {
			minlengths_newInstances = append(minlengths_newInstances, minlength)
			newInstancesSlice = append(newInstancesSlice, minlength.GongMarshallIdentifier(stage))
			if stage.MinLengths_referenceOrder == nil {
				stage.MinLengths_referenceOrder = make(map[*MinLength]uint)
			}
			stage.MinLengths_referenceOrder[minlength] = stage.MinLengthMap_Staged_Order[minlength]
			newInstancesReverseSlice = append(newInstancesReverseSlice, minlength.GongMarshallUnstaging(stage))
			delete(stage.MinLengths_referenceOrder, minlength)
			fieldInitializers, pointersInitializations := minlength.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MinLengthMap_Staged_Order[ref] = stage.MinLengthMap_Staged_Order[minlength]
			diffs := minlength.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, minlength)
			delete(stage.MinLengthMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", minlength.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.MinLengths_reference {
		if _, ok := stage.MinLengths[ref]; !ok {
			minlengths_deletedInstances = append(minlengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(minlengths_newInstances)
	lenDeletedInstances += len(minlengths_deletedInstances)
	var patterns_newInstances []*Pattern
	var patterns_deletedInstances []*Pattern

	// parse all staged instances and check if they have a reference
	for pattern := range stage.Patterns {
		if ref, ok := stage.Patterns_reference[pattern]; !ok {
			patterns_newInstances = append(patterns_newInstances, pattern)
			newInstancesSlice = append(newInstancesSlice, pattern.GongMarshallIdentifier(stage))
			if stage.Patterns_referenceOrder == nil {
				stage.Patterns_referenceOrder = make(map[*Pattern]uint)
			}
			stage.Patterns_referenceOrder[pattern] = stage.PatternMap_Staged_Order[pattern]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pattern.GongMarshallUnstaging(stage))
			delete(stage.Patterns_referenceOrder, pattern)
			fieldInitializers, pointersInitializations := pattern.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PatternMap_Staged_Order[ref] = stage.PatternMap_Staged_Order[pattern]
			diffs := pattern.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pattern)
			delete(stage.PatternMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", pattern.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Patterns_reference {
		if _, ok := stage.Patterns[ref]; !ok {
			patterns_deletedInstances = append(patterns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(patterns_newInstances)
	lenDeletedInstances += len(patterns_deletedInstances)
	var restrictions_newInstances []*Restriction
	var restrictions_deletedInstances []*Restriction

	// parse all staged instances and check if they have a reference
	for restriction := range stage.Restrictions {
		if ref, ok := stage.Restrictions_reference[restriction]; !ok {
			restrictions_newInstances = append(restrictions_newInstances, restriction)
			newInstancesSlice = append(newInstancesSlice, restriction.GongMarshallIdentifier(stage))
			if stage.Restrictions_referenceOrder == nil {
				stage.Restrictions_referenceOrder = make(map[*Restriction]uint)
			}
			stage.Restrictions_referenceOrder[restriction] = stage.RestrictionMap_Staged_Order[restriction]
			newInstancesReverseSlice = append(newInstancesReverseSlice, restriction.GongMarshallUnstaging(stage))
			delete(stage.Restrictions_referenceOrder, restriction)
			fieldInitializers, pointersInitializations := restriction.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RestrictionMap_Staged_Order[ref] = stage.RestrictionMap_Staged_Order[restriction]
			diffs := restriction.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, restriction)
			delete(stage.RestrictionMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", restriction.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Restrictions_reference {
		if _, ok := stage.Restrictions[ref]; !ok {
			restrictions_deletedInstances = append(restrictions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(restrictions_newInstances)
	lenDeletedInstances += len(restrictions_deletedInstances)
	var schemas_newInstances []*Schema
	var schemas_deletedInstances []*Schema

	// parse all staged instances and check if they have a reference
	for schema := range stage.Schemas {
		if ref, ok := stage.Schemas_reference[schema]; !ok {
			schemas_newInstances = append(schemas_newInstances, schema)
			newInstancesSlice = append(newInstancesSlice, schema.GongMarshallIdentifier(stage))
			if stage.Schemas_referenceOrder == nil {
				stage.Schemas_referenceOrder = make(map[*Schema]uint)
			}
			stage.Schemas_referenceOrder[schema] = stage.SchemaMap_Staged_Order[schema]
			newInstancesReverseSlice = append(newInstancesReverseSlice, schema.GongMarshallUnstaging(stage))
			delete(stage.Schemas_referenceOrder, schema)
			fieldInitializers, pointersInitializations := schema.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SchemaMap_Staged_Order[ref] = stage.SchemaMap_Staged_Order[schema]
			diffs := schema.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, schema)
			delete(stage.SchemaMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", schema.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Schemas_reference {
		if _, ok := stage.Schemas[ref]; !ok {
			schemas_deletedInstances = append(schemas_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(schemas_newInstances)
	lenDeletedInstances += len(schemas_deletedInstances)
	var sequences_newInstances []*Sequence
	var sequences_deletedInstances []*Sequence

	// parse all staged instances and check if they have a reference
	for sequence := range stage.Sequences {
		if ref, ok := stage.Sequences_reference[sequence]; !ok {
			sequences_newInstances = append(sequences_newInstances, sequence)
			newInstancesSlice = append(newInstancesSlice, sequence.GongMarshallIdentifier(stage))
			if stage.Sequences_referenceOrder == nil {
				stage.Sequences_referenceOrder = make(map[*Sequence]uint)
			}
			stage.Sequences_referenceOrder[sequence] = stage.SequenceMap_Staged_Order[sequence]
			newInstancesReverseSlice = append(newInstancesReverseSlice, sequence.GongMarshallUnstaging(stage))
			delete(stage.Sequences_referenceOrder, sequence)
			fieldInitializers, pointersInitializations := sequence.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SequenceMap_Staged_Order[ref] = stage.SequenceMap_Staged_Order[sequence]
			diffs := sequence.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, sequence)
			delete(stage.SequenceMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", sequence.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Sequences_reference {
		if _, ok := stage.Sequences[ref]; !ok {
			sequences_deletedInstances = append(sequences_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(sequences_newInstances)
	lenDeletedInstances += len(sequences_deletedInstances)
	var simplecontents_newInstances []*SimpleContent
	var simplecontents_deletedInstances []*SimpleContent

	// parse all staged instances and check if they have a reference
	for simplecontent := range stage.SimpleContents {
		if ref, ok := stage.SimpleContents_reference[simplecontent]; !ok {
			simplecontents_newInstances = append(simplecontents_newInstances, simplecontent)
			newInstancesSlice = append(newInstancesSlice, simplecontent.GongMarshallIdentifier(stage))
			if stage.SimpleContents_referenceOrder == nil {
				stage.SimpleContents_referenceOrder = make(map[*SimpleContent]uint)
			}
			stage.SimpleContents_referenceOrder[simplecontent] = stage.SimpleContentMap_Staged_Order[simplecontent]
			newInstancesReverseSlice = append(newInstancesReverseSlice, simplecontent.GongMarshallUnstaging(stage))
			delete(stage.SimpleContents_referenceOrder, simplecontent)
			fieldInitializers, pointersInitializations := simplecontent.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SimpleContentMap_Staged_Order[ref] = stage.SimpleContentMap_Staged_Order[simplecontent]
			diffs := simplecontent.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, simplecontent)
			delete(stage.SimpleContentMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", simplecontent.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SimpleContents_reference {
		if _, ok := stage.SimpleContents[ref]; !ok {
			simplecontents_deletedInstances = append(simplecontents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(simplecontents_newInstances)
	lenDeletedInstances += len(simplecontents_deletedInstances)
	var simpletypes_newInstances []*SimpleType
	var simpletypes_deletedInstances []*SimpleType

	// parse all staged instances and check if they have a reference
	for simpletype := range stage.SimpleTypes {
		if ref, ok := stage.SimpleTypes_reference[simpletype]; !ok {
			simpletypes_newInstances = append(simpletypes_newInstances, simpletype)
			newInstancesSlice = append(newInstancesSlice, simpletype.GongMarshallIdentifier(stage))
			if stage.SimpleTypes_referenceOrder == nil {
				stage.SimpleTypes_referenceOrder = make(map[*SimpleType]uint)
			}
			stage.SimpleTypes_referenceOrder[simpletype] = stage.SimpleTypeMap_Staged_Order[simpletype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, simpletype.GongMarshallUnstaging(stage))
			delete(stage.SimpleTypes_referenceOrder, simpletype)
			fieldInitializers, pointersInitializations := simpletype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SimpleTypeMap_Staged_Order[ref] = stage.SimpleTypeMap_Staged_Order[simpletype]
			diffs := simpletype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, simpletype)
			delete(stage.SimpleTypeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", simpletype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SimpleTypes_reference {
		if _, ok := stage.SimpleTypes[ref]; !ok {
			simpletypes_deletedInstances = append(simpletypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(simpletypes_newInstances)
	lenDeletedInstances += len(simpletypes_deletedInstances)
	var totaldigits_newInstances []*TotalDigit
	var totaldigits_deletedInstances []*TotalDigit

	// parse all staged instances and check if they have a reference
	for totaldigit := range stage.TotalDigits {
		if ref, ok := stage.TotalDigits_reference[totaldigit]; !ok {
			totaldigits_newInstances = append(totaldigits_newInstances, totaldigit)
			newInstancesSlice = append(newInstancesSlice, totaldigit.GongMarshallIdentifier(stage))
			if stage.TotalDigits_referenceOrder == nil {
				stage.TotalDigits_referenceOrder = make(map[*TotalDigit]uint)
			}
			stage.TotalDigits_referenceOrder[totaldigit] = stage.TotalDigitMap_Staged_Order[totaldigit]
			newInstancesReverseSlice = append(newInstancesReverseSlice, totaldigit.GongMarshallUnstaging(stage))
			delete(stage.TotalDigits_referenceOrder, totaldigit)
			fieldInitializers, pointersInitializations := totaldigit.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TotalDigitMap_Staged_Order[ref] = stage.TotalDigitMap_Staged_Order[totaldigit]
			diffs := totaldigit.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, totaldigit)
			delete(stage.TotalDigitMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", totaldigit.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.TotalDigits_reference {
		if _, ok := stage.TotalDigits[ref]; !ok {
			totaldigits_deletedInstances = append(totaldigits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(totaldigits_newInstances)
	lenDeletedInstances += len(totaldigits_deletedInstances)
	var unions_newInstances []*Union
	var unions_deletedInstances []*Union

	// parse all staged instances and check if they have a reference
	for union := range stage.Unions {
		if ref, ok := stage.Unions_reference[union]; !ok {
			unions_newInstances = append(unions_newInstances, union)
			newInstancesSlice = append(newInstancesSlice, union.GongMarshallIdentifier(stage))
			if stage.Unions_referenceOrder == nil {
				stage.Unions_referenceOrder = make(map[*Union]uint)
			}
			stage.Unions_referenceOrder[union] = stage.UnionMap_Staged_Order[union]
			newInstancesReverseSlice = append(newInstancesReverseSlice, union.GongMarshallUnstaging(stage))
			delete(stage.Unions_referenceOrder, union)
			fieldInitializers, pointersInitializations := union.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.UnionMap_Staged_Order[ref] = stage.UnionMap_Staged_Order[union]
			diffs := union.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, union)
			delete(stage.UnionMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", union.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Unions_reference {
		if _, ok := stage.Unions[ref]; !ok {
			unions_deletedInstances = append(unions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(unions_newInstances)
	lenDeletedInstances += len(unions_deletedInstances)
	var whitespaces_newInstances []*WhiteSpace
	var whitespaces_deletedInstances []*WhiteSpace

	// parse all staged instances and check if they have a reference
	for whitespace := range stage.WhiteSpaces {
		if ref, ok := stage.WhiteSpaces_reference[whitespace]; !ok {
			whitespaces_newInstances = append(whitespaces_newInstances, whitespace)
			newInstancesSlice = append(newInstancesSlice, whitespace.GongMarshallIdentifier(stage))
			if stage.WhiteSpaces_referenceOrder == nil {
				stage.WhiteSpaces_referenceOrder = make(map[*WhiteSpace]uint)
			}
			stage.WhiteSpaces_referenceOrder[whitespace] = stage.WhiteSpaceMap_Staged_Order[whitespace]
			newInstancesReverseSlice = append(newInstancesReverseSlice, whitespace.GongMarshallUnstaging(stage))
			delete(stage.WhiteSpaces_referenceOrder, whitespace)
			fieldInitializers, pointersInitializations := whitespace.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.WhiteSpaceMap_Staged_Order[ref] = stage.WhiteSpaceMap_Staged_Order[whitespace]
			diffs := whitespace.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, whitespace)
			delete(stage.WhiteSpaceMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", whitespace.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.WhiteSpaces_reference {
		if _, ok := stage.WhiteSpaces[ref]; !ok {
			whitespaces_deletedInstances = append(whitespaces_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(whitespaces_newInstances)
	lenDeletedInstances += len(whitespaces_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.Alls_reference = make(map[*All]*All)
	stage.Alls_referenceOrder = make(map[*All]uint) // diff Unstage needs the reference order
	for instance := range stage.Alls {
		stage.Alls_reference[instance] = instance.GongCopy().(*All)
		stage.Alls_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Annotations_reference = make(map[*Annotation]*Annotation)
	stage.Annotations_referenceOrder = make(map[*Annotation]uint) // diff Unstage needs the reference order
	for instance := range stage.Annotations {
		stage.Annotations_reference[instance] = instance.GongCopy().(*Annotation)
		stage.Annotations_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Attributes_reference = make(map[*Attribute]*Attribute)
	stage.Attributes_referenceOrder = make(map[*Attribute]uint) // diff Unstage needs the reference order
	for instance := range stage.Attributes {
		stage.Attributes_reference[instance] = instance.GongCopy().(*Attribute)
		stage.Attributes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.AttributeGroups_reference = make(map[*AttributeGroup]*AttributeGroup)
	stage.AttributeGroups_referenceOrder = make(map[*AttributeGroup]uint) // diff Unstage needs the reference order
	for instance := range stage.AttributeGroups {
		stage.AttributeGroups_reference[instance] = instance.GongCopy().(*AttributeGroup)
		stage.AttributeGroups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Choices_reference = make(map[*Choice]*Choice)
	stage.Choices_referenceOrder = make(map[*Choice]uint) // diff Unstage needs the reference order
	for instance := range stage.Choices {
		stage.Choices_reference[instance] = instance.GongCopy().(*Choice)
		stage.Choices_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ComplexContents_reference = make(map[*ComplexContent]*ComplexContent)
	stage.ComplexContents_referenceOrder = make(map[*ComplexContent]uint) // diff Unstage needs the reference order
	for instance := range stage.ComplexContents {
		stage.ComplexContents_reference[instance] = instance.GongCopy().(*ComplexContent)
		stage.ComplexContents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ComplexTypes_reference = make(map[*ComplexType]*ComplexType)
	stage.ComplexTypes_referenceOrder = make(map[*ComplexType]uint) // diff Unstage needs the reference order
	for instance := range stage.ComplexTypes {
		stage.ComplexTypes_reference[instance] = instance.GongCopy().(*ComplexType)
		stage.ComplexTypes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Documentations_reference = make(map[*Documentation]*Documentation)
	stage.Documentations_referenceOrder = make(map[*Documentation]uint) // diff Unstage needs the reference order
	for instance := range stage.Documentations {
		stage.Documentations_reference[instance] = instance.GongCopy().(*Documentation)
		stage.Documentations_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Elements_reference = make(map[*Element]*Element)
	stage.Elements_referenceOrder = make(map[*Element]uint) // diff Unstage needs the reference order
	for instance := range stage.Elements {
		stage.Elements_reference[instance] = instance.GongCopy().(*Element)
		stage.Elements_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Enumerations_reference = make(map[*Enumeration]*Enumeration)
	stage.Enumerations_referenceOrder = make(map[*Enumeration]uint) // diff Unstage needs the reference order
	for instance := range stage.Enumerations {
		stage.Enumerations_reference[instance] = instance.GongCopy().(*Enumeration)
		stage.Enumerations_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Extensions_reference = make(map[*Extension]*Extension)
	stage.Extensions_referenceOrder = make(map[*Extension]uint) // diff Unstage needs the reference order
	for instance := range stage.Extensions {
		stage.Extensions_reference[instance] = instance.GongCopy().(*Extension)
		stage.Extensions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
		stage.Groups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Lengths_reference = make(map[*Length]*Length)
	stage.Lengths_referenceOrder = make(map[*Length]uint) // diff Unstage needs the reference order
	for instance := range stage.Lengths {
		stage.Lengths_reference[instance] = instance.GongCopy().(*Length)
		stage.Lengths_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MaxInclusives_reference = make(map[*MaxInclusive]*MaxInclusive)
	stage.MaxInclusives_referenceOrder = make(map[*MaxInclusive]uint) // diff Unstage needs the reference order
	for instance := range stage.MaxInclusives {
		stage.MaxInclusives_reference[instance] = instance.GongCopy().(*MaxInclusive)
		stage.MaxInclusives_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MaxLengths_reference = make(map[*MaxLength]*MaxLength)
	stage.MaxLengths_referenceOrder = make(map[*MaxLength]uint) // diff Unstage needs the reference order
	for instance := range stage.MaxLengths {
		stage.MaxLengths_reference[instance] = instance.GongCopy().(*MaxLength)
		stage.MaxLengths_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MinInclusives_reference = make(map[*MinInclusive]*MinInclusive)
	stage.MinInclusives_referenceOrder = make(map[*MinInclusive]uint) // diff Unstage needs the reference order
	for instance := range stage.MinInclusives {
		stage.MinInclusives_reference[instance] = instance.GongCopy().(*MinInclusive)
		stage.MinInclusives_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MinLengths_reference = make(map[*MinLength]*MinLength)
	stage.MinLengths_referenceOrder = make(map[*MinLength]uint) // diff Unstage needs the reference order
	for instance := range stage.MinLengths {
		stage.MinLengths_reference[instance] = instance.GongCopy().(*MinLength)
		stage.MinLengths_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Patterns_reference = make(map[*Pattern]*Pattern)
	stage.Patterns_referenceOrder = make(map[*Pattern]uint) // diff Unstage needs the reference order
	for instance := range stage.Patterns {
		stage.Patterns_reference[instance] = instance.GongCopy().(*Pattern)
		stage.Patterns_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Restrictions_reference = make(map[*Restriction]*Restriction)
	stage.Restrictions_referenceOrder = make(map[*Restriction]uint) // diff Unstage needs the reference order
	for instance := range stage.Restrictions {
		stage.Restrictions_reference[instance] = instance.GongCopy().(*Restriction)
		stage.Restrictions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Schemas_reference = make(map[*Schema]*Schema)
	stage.Schemas_referenceOrder = make(map[*Schema]uint) // diff Unstage needs the reference order
	for instance := range stage.Schemas {
		stage.Schemas_reference[instance] = instance.GongCopy().(*Schema)
		stage.Schemas_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Sequences_reference = make(map[*Sequence]*Sequence)
	stage.Sequences_referenceOrder = make(map[*Sequence]uint) // diff Unstage needs the reference order
	for instance := range stage.Sequences {
		stage.Sequences_reference[instance] = instance.GongCopy().(*Sequence)
		stage.Sequences_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SimpleContents_reference = make(map[*SimpleContent]*SimpleContent)
	stage.SimpleContents_referenceOrder = make(map[*SimpleContent]uint) // diff Unstage needs the reference order
	for instance := range stage.SimpleContents {
		stage.SimpleContents_reference[instance] = instance.GongCopy().(*SimpleContent)
		stage.SimpleContents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SimpleTypes_reference = make(map[*SimpleType]*SimpleType)
	stage.SimpleTypes_referenceOrder = make(map[*SimpleType]uint) // diff Unstage needs the reference order
	for instance := range stage.SimpleTypes {
		stage.SimpleTypes_reference[instance] = instance.GongCopy().(*SimpleType)
		stage.SimpleTypes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TotalDigits_reference = make(map[*TotalDigit]*TotalDigit)
	stage.TotalDigits_referenceOrder = make(map[*TotalDigit]uint) // diff Unstage needs the reference order
	for instance := range stage.TotalDigits {
		stage.TotalDigits_reference[instance] = instance.GongCopy().(*TotalDigit)
		stage.TotalDigits_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Unions_reference = make(map[*Union]*Union)
	stage.Unions_referenceOrder = make(map[*Union]uint) // diff Unstage needs the reference order
	for instance := range stage.Unions {
		stage.Unions_reference[instance] = instance.GongCopy().(*Union)
		stage.Unions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.WhiteSpaces_reference = make(map[*WhiteSpace]*WhiteSpace)
	stage.WhiteSpaces_referenceOrder = make(map[*WhiteSpace]uint) // diff Unstage needs the reference order
	for instance := range stage.WhiteSpaces {
		stage.WhiteSpaces_reference[instance] = instance.GongCopy().(*WhiteSpace)
		stage.WhiteSpaces_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (all *All) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AllMap_Staged_Order[all]; ok {
		return order
	}
	return stage.Alls_referenceOrder[all]
}

func (all *All) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Alls_referenceOrder[all]
}

func (annotation *Annotation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AnnotationMap_Staged_Order[annotation]; ok {
		return order
	}
	return stage.Annotations_referenceOrder[annotation]
}

func (annotation *Annotation) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Annotations_referenceOrder[annotation]
}

func (attribute *Attribute) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AttributeMap_Staged_Order[attribute]; ok {
		return order
	}
	return stage.Attributes_referenceOrder[attribute]
}

func (attribute *Attribute) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Attributes_referenceOrder[attribute]
}

func (attributegroup *AttributeGroup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AttributeGroupMap_Staged_Order[attributegroup]; ok {
		return order
	}
	return stage.AttributeGroups_referenceOrder[attributegroup]
}

func (attributegroup *AttributeGroup) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AttributeGroups_referenceOrder[attributegroup]
}

func (choice *Choice) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ChoiceMap_Staged_Order[choice]; ok {
		return order
	}
	return stage.Choices_referenceOrder[choice]
}

func (choice *Choice) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Choices_referenceOrder[choice]
}

func (complexcontent *ComplexContent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ComplexContentMap_Staged_Order[complexcontent]; ok {
		return order
	}
	return stage.ComplexContents_referenceOrder[complexcontent]
}

func (complexcontent *ComplexContent) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ComplexContents_referenceOrder[complexcontent]
}

func (complextype *ComplexType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ComplexTypeMap_Staged_Order[complextype]; ok {
		return order
	}
	return stage.ComplexTypes_referenceOrder[complextype]
}

func (complextype *ComplexType) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ComplexTypes_referenceOrder[complextype]
}

func (documentation *Documentation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DocumentationMap_Staged_Order[documentation]; ok {
		return order
	}
	return stage.Documentations_referenceOrder[documentation]
}

func (documentation *Documentation) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Documentations_referenceOrder[documentation]
}

func (element *Element) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ElementMap_Staged_Order[element]; ok {
		return order
	}
	return stage.Elements_referenceOrder[element]
}

func (element *Element) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Elements_referenceOrder[element]
}

func (enumeration *Enumeration) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EnumerationMap_Staged_Order[enumeration]; ok {
		return order
	}
	return stage.Enumerations_referenceOrder[enumeration]
}

func (enumeration *Enumeration) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Enumerations_referenceOrder[enumeration]
}

func (extension *Extension) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExtensionMap_Staged_Order[extension]; ok {
		return order
	}
	return stage.Extensions_referenceOrder[extension]
}

func (extension *Extension) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Extensions_referenceOrder[extension]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GroupMap_Staged_Order[group]; ok {
		return order
	}
	return stage.Groups_referenceOrder[group]
}

func (group *Group) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Groups_referenceOrder[group]
}

func (length *Length) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LengthMap_Staged_Order[length]; ok {
		return order
	}
	return stage.Lengths_referenceOrder[length]
}

func (length *Length) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Lengths_referenceOrder[length]
}

func (maxinclusive *MaxInclusive) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MaxInclusiveMap_Staged_Order[maxinclusive]; ok {
		return order
	}
	return stage.MaxInclusives_referenceOrder[maxinclusive]
}

func (maxinclusive *MaxInclusive) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MaxInclusives_referenceOrder[maxinclusive]
}

func (maxlength *MaxLength) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MaxLengthMap_Staged_Order[maxlength]; ok {
		return order
	}
	return stage.MaxLengths_referenceOrder[maxlength]
}

func (maxlength *MaxLength) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MaxLengths_referenceOrder[maxlength]
}

func (mininclusive *MinInclusive) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MinInclusiveMap_Staged_Order[mininclusive]; ok {
		return order
	}
	return stage.MinInclusives_referenceOrder[mininclusive]
}

func (mininclusive *MinInclusive) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MinInclusives_referenceOrder[mininclusive]
}

func (minlength *MinLength) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MinLengthMap_Staged_Order[minlength]; ok {
		return order
	}
	return stage.MinLengths_referenceOrder[minlength]
}

func (minlength *MinLength) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MinLengths_referenceOrder[minlength]
}

func (pattern *Pattern) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PatternMap_Staged_Order[pattern]; ok {
		return order
	}
	return stage.Patterns_referenceOrder[pattern]
}

func (pattern *Pattern) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Patterns_referenceOrder[pattern]
}

func (restriction *Restriction) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RestrictionMap_Staged_Order[restriction]; ok {
		return order
	}
	return stage.Restrictions_referenceOrder[restriction]
}

func (restriction *Restriction) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Restrictions_referenceOrder[restriction]
}

func (schema *Schema) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SchemaMap_Staged_Order[schema]; ok {
		return order
	}
	return stage.Schemas_referenceOrder[schema]
}

func (schema *Schema) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Schemas_referenceOrder[schema]
}

func (sequence *Sequence) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SequenceMap_Staged_Order[sequence]; ok {
		return order
	}
	return stage.Sequences_referenceOrder[sequence]
}

func (sequence *Sequence) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Sequences_referenceOrder[sequence]
}

func (simplecontent *SimpleContent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SimpleContentMap_Staged_Order[simplecontent]; ok {
		return order
	}
	return stage.SimpleContents_referenceOrder[simplecontent]
}

func (simplecontent *SimpleContent) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SimpleContents_referenceOrder[simplecontent]
}

func (simpletype *SimpleType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SimpleTypeMap_Staged_Order[simpletype]; ok {
		return order
	}
	return stage.SimpleTypes_referenceOrder[simpletype]
}

func (simpletype *SimpleType) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SimpleTypes_referenceOrder[simpletype]
}

func (totaldigit *TotalDigit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TotalDigitMap_Staged_Order[totaldigit]; ok {
		return order
	}
	return stage.TotalDigits_referenceOrder[totaldigit]
}

func (totaldigit *TotalDigit) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TotalDigits_referenceOrder[totaldigit]
}

func (union *Union) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.UnionMap_Staged_Order[union]; ok {
		return order
	}
	return stage.Unions_referenceOrder[union]
}

func (union *Union) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Unions_referenceOrder[union]
}

func (whitespace *WhiteSpace) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.WhiteSpaceMap_Staged_Order[whitespace]; ok {
		return order
	}
	return stage.WhiteSpaces_referenceOrder[whitespace]
}

func (whitespace *WhiteSpace) GongGetReferenceOrder(stage *Stage) uint {
	return stage.WhiteSpaces_referenceOrder[whitespace]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (all *All) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", all.GongGetGongstructName(), all.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (all *All) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", all.GongGetGongstructName(), all.GongGetReferenceOrder(stage))
}

func (annotation *Annotation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", annotation.GongGetGongstructName(), annotation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (annotation *Annotation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", annotation.GongGetGongstructName(), annotation.GongGetReferenceOrder(stage))
}

func (attribute *Attribute) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute.GongGetGongstructName(), attribute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute *Attribute) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute.GongGetGongstructName(), attribute.GongGetReferenceOrder(stage))
}

func (attributegroup *AttributeGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributegroup.GongGetGongstructName(), attributegroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributegroup *AttributeGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributegroup.GongGetGongstructName(), attributegroup.GongGetReferenceOrder(stage))
}

func (choice *Choice) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", choice.GongGetGongstructName(), choice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (choice *Choice) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", choice.GongGetGongstructName(), choice.GongGetReferenceOrder(stage))
}

func (complexcontent *ComplexContent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexcontent.GongGetGongstructName(), complexcontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexcontent *ComplexContent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexcontent.GongGetGongstructName(), complexcontent.GongGetReferenceOrder(stage))
}

func (complextype *ComplexType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complextype.GongGetGongstructName(), complextype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complextype *ComplexType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complextype.GongGetGongstructName(), complextype.GongGetReferenceOrder(stage))
}

func (documentation *Documentation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentation.GongGetGongstructName(), documentation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (documentation *Documentation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentation.GongGetGongstructName(), documentation.GongGetReferenceOrder(stage))
}

func (element *Element) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", element.GongGetGongstructName(), element.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (element *Element) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", element.GongGetGongstructName(), element.GongGetReferenceOrder(stage))
}

func (enumeration *Enumeration) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enumeration.GongGetGongstructName(), enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enumeration *Enumeration) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enumeration.GongGetGongstructName(), enumeration.GongGetReferenceOrder(stage))
}

func (extension *Extension) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extension.GongGetGongstructName(), extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extension *Extension) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extension.GongGetGongstructName(), extension.GongGetReferenceOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetReferenceOrder(stage))
}

func (length *Length) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", length.GongGetGongstructName(), length.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (length *Length) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", length.GongGetGongstructName(), length.GongGetReferenceOrder(stage))
}

func (maxinclusive *MaxInclusive) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxinclusive.GongGetGongstructName(), maxinclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxinclusive *MaxInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxinclusive.GongGetGongstructName(), maxinclusive.GongGetReferenceOrder(stage))
}

func (maxlength *MaxLength) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxlength.GongGetGongstructName(), maxlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxlength *MaxLength) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxlength.GongGetGongstructName(), maxlength.GongGetReferenceOrder(stage))
}

func (mininclusive *MinInclusive) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mininclusive.GongGetGongstructName(), mininclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mininclusive *MinInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mininclusive.GongGetGongstructName(), mininclusive.GongGetReferenceOrder(stage))
}

func (minlength *MinLength) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", minlength.GongGetGongstructName(), minlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (minlength *MinLength) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", minlength.GongGetGongstructName(), minlength.GongGetReferenceOrder(stage))
}

func (pattern *Pattern) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pattern.GongGetGongstructName(), pattern.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pattern *Pattern) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pattern.GongGetGongstructName(), pattern.GongGetReferenceOrder(stage))
}

func (restriction *Restriction) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", restriction.GongGetGongstructName(), restriction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (restriction *Restriction) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", restriction.GongGetGongstructName(), restriction.GongGetReferenceOrder(stage))
}

func (schema *Schema) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", schema.GongGetGongstructName(), schema.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (schema *Schema) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", schema.GongGetGongstructName(), schema.GongGetReferenceOrder(stage))
}

func (sequence *Sequence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sequence.GongGetGongstructName(), sequence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sequence *Sequence) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sequence.GongGetGongstructName(), sequence.GongGetReferenceOrder(stage))
}

func (simplecontent *SimpleContent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simplecontent.GongGetGongstructName(), simplecontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simplecontent *SimpleContent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simplecontent.GongGetGongstructName(), simplecontent.GongGetReferenceOrder(stage))
}

func (simpletype *SimpleType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simpletype.GongGetGongstructName(), simpletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simpletype *SimpleType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simpletype.GongGetGongstructName(), simpletype.GongGetReferenceOrder(stage))
}

func (totaldigit *TotalDigit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", totaldigit.GongGetGongstructName(), totaldigit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (totaldigit *TotalDigit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", totaldigit.GongGetGongstructName(), totaldigit.GongGetReferenceOrder(stage))
}

func (union *Union) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", union.GongGetGongstructName(), union.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (union *Union) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", union.GongGetGongstructName(), union.GongGetReferenceOrder(stage))
}

func (whitespace *WhiteSpace) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", whitespace.GongGetGongstructName(), whitespace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (whitespace *WhiteSpace) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", whitespace.GongGetGongstructName(), whitespace.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (all *All) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", all.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "All")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", all.Name)
	return
}

func (annotation *Annotation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", annotation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Annotation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", annotation.Name)
	return
}

func (attribute *Attribute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attribute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute.Name)
	return
}

func (attributegroup *AttributeGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributegroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributegroup.Name)
	return
}

func (choice *Choice) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", choice.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Choice")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", choice.Name)
	return
}

func (complexcontent *ComplexContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexcontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complexcontent.Name)
	return
}

func (complextype *ComplexType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complextype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complextype.Name)
	return
}

func (documentation *Documentation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Documentation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", documentation.Name)
	return
}

func (element *Element) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", element.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", element.Name)
	return
}

func (enumeration *Enumeration) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Enumeration")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enumeration.Name)
	return
}

func (extension *Extension) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extension.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extension")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", extension.Name)
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}

func (length *Length) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", length.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Length")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", length.Name)
	return
}

func (maxinclusive *MaxInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxinclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", maxinclusive.Name)
	return
}

func (maxlength *MaxLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", maxlength.Name)
	return
}

func (mininclusive *MinInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mininclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mininclusive.Name)
	return
}

func (minlength *MinLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", minlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", minlength.Name)
	return
}

func (pattern *Pattern) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pattern.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pattern")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pattern.Name)
	return
}

func (restriction *Restriction) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", restriction.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Restriction")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", restriction.Name)
	return
}

func (schema *Schema) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", schema.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Schema")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", schema.Name)
	return
}

func (sequence *Sequence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sequence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sequence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sequence.Name)
	return
}

func (simplecontent *SimpleContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simplecontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", simplecontent.Name)
	return
}

func (simpletype *SimpleType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simpletype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", simpletype.Name)
	return
}

func (totaldigit *TotalDigit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", totaldigit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TotalDigit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", totaldigit.Name)
	return
}

func (union *Union) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", union.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Union")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", union.Name)
	return
}

func (whitespace *WhiteSpace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", whitespace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "WhiteSpace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", whitespace.Name)
	return
}

// insertion point for unstaging
func (all *All) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", all.GongGetReferenceIdentifier(stage))
	return
}

func (annotation *Annotation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", annotation.GongGetReferenceIdentifier(stage))
	return
}

func (attribute *Attribute) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute.GongGetReferenceIdentifier(stage))
	return
}

func (attributegroup *AttributeGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributegroup.GongGetReferenceIdentifier(stage))
	return
}

func (choice *Choice) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", choice.GongGetReferenceIdentifier(stage))
	return
}

func (complexcontent *ComplexContent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexcontent.GongGetReferenceIdentifier(stage))
	return
}

func (complextype *ComplexType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complextype.GongGetReferenceIdentifier(stage))
	return
}

func (documentation *Documentation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentation.GongGetReferenceIdentifier(stage))
	return
}

func (element *Element) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", element.GongGetReferenceIdentifier(stage))
	return
}

func (enumeration *Enumeration) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (extension *Extension) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extension.GongGetReferenceIdentifier(stage))
	return
}

func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}

func (length *Length) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", length.GongGetReferenceIdentifier(stage))
	return
}

func (maxinclusive *MaxInclusive) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxinclusive.GongGetReferenceIdentifier(stage))
	return
}

func (maxlength *MaxLength) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxlength.GongGetReferenceIdentifier(stage))
	return
}

func (mininclusive *MinInclusive) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mininclusive.GongGetReferenceIdentifier(stage))
	return
}

func (minlength *MinLength) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", minlength.GongGetReferenceIdentifier(stage))
	return
}

func (pattern *Pattern) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pattern.GongGetReferenceIdentifier(stage))
	return
}

func (restriction *Restriction) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", restriction.GongGetReferenceIdentifier(stage))
	return
}

func (schema *Schema) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", schema.GongGetReferenceIdentifier(stage))
	return
}

func (sequence *Sequence) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sequence.GongGetReferenceIdentifier(stage))
	return
}

func (simplecontent *SimpleContent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simplecontent.GongGetReferenceIdentifier(stage))
	return
}

func (simpletype *SimpleType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simpletype.GongGetReferenceIdentifier(stage))
	return
}

func (totaldigit *TotalDigit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", totaldigit.GongGetReferenceIdentifier(stage))
	return
}

func (union *Union) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", union.GongGetReferenceIdentifier(stage))
	return
}

func (whitespace *WhiteSpace) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", whitespace.GongGetReferenceIdentifier(stage))
	return
}

// end of template
