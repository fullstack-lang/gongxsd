// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.All:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)

	case *models.Annotation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Documentations", instanceWithInferedType, &instanceWithInferedType.Documentations, formGroup, probe)

	case *models.Attribute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default", instanceWithInferedType.Default, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Use", instanceWithInferedType.Use, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Form", instanceWithInferedType.Form, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fixed", instanceWithInferedType.Fixed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TargetNamespace", instanceWithInferedType.TargetNamespace, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SimpleType", instanceWithInferedType.SimpleType, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDXSD", instanceWithInferedType.IDXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AttributeGroup"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AttributeGroup),
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AttributeGroup, *models.Attribute](
					nil,
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ComplexType),
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ComplexType, *models.Attribute](
					nil,
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Extension),
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Extension, *models.Attribute](
					nil,
					"Attributes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.AttributeGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AttributeGroup"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.AttributeGroup),
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.AttributeGroup, *models.AttributeGroup](
					nil,
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ComplexType),
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ComplexType, *models.AttributeGroup](
					nil,
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Schema),
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Schema, *models.AttributeGroup](
					nil,
					"AttributeGroups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Choice:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)

	case *models.ComplexContent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.ComplexType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAnonymous", instanceWithInferedType.IsAnonymous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("OuterElement", instanceWithInferedType.OuterElement, formGroup, probe)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)
		AssociationFieldToForm("Extension", instanceWithInferedType.Extension, formGroup, probe)
		AssociationFieldToForm("SimpleContent", instanceWithInferedType.SimpleContent, formGroup, probe)
		AssociationFieldToForm("ComplexContent", instanceWithInferedType.ComplexContent, formGroup, probe)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "ComplexTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Schema),
					"ComplexTypes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Schema, *models.ComplexType](
					nil,
					"ComplexTypes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Documentation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Annotation"
			rf.Fieldname = "Documentations"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Annotation),
					"Documentations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Annotation, *models.Documentation](
					nil,
					"Documentations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Element:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Default", instanceWithInferedType.Default, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fixed", instanceWithInferedType.Fixed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Nillable", instanceWithInferedType.Nillable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Abstract", instanceWithInferedType.Abstract, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Form", instanceWithInferedType.Form, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Block", instanceWithInferedType.Block, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Final", instanceWithInferedType.Final, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SimpleType", instanceWithInferedType.SimpleType, formGroup, probe)
		AssociationFieldToForm("ComplexType", instanceWithInferedType.ComplexType, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Schema),
					"Elements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Schema, *models.Element](
					nil,
					"Elements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Enumeration:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Restriction"
			rf.Fieldname = "Enumerations"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Restriction),
					"Enumerations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Restriction, *models.Enumeration](
					nil,
					"Enumerations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Extension:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)
		BasicFieldtoForm("Base", instanceWithInferedType.Base, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAnonymous", instanceWithInferedType.IsAnonymous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("OuterElement", instanceWithInferedType.OuterElement, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Element),
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Element, *models.Group](
					nil,
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Schema),
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Schema, *models.Group](
					nil,
					"Groups",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Length:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MaxInclusive:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MaxLength:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MinInclusive:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.MinLength:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.ModelGroupElement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Sequences", instanceWithInferedType.Sequences, formGroup, probe)
		AssociationFieldToForm("Alls", instanceWithInferedType.Alls, formGroup, probe)
		AssociationFieldToForm("Choices", instanceWithInferedType.Choices, formGroup, probe)
		AssociationFieldToForm("Groups", instanceWithInferedType.Groups, formGroup, probe)
		AssociationFieldToForm("Elements", instanceWithInferedType.Elements, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.All),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.All, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Choice),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Choice, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ComplexType),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ComplexType, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Extension),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Extension, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Group),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Group, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "ModelGroupElements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sequence),
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sequence, *models.ModelGroupElement](
					nil,
					"ModelGroupElements",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Pattern:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Restriction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Base", instanceWithInferedType.Base, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Enumerations", instanceWithInferedType, &instanceWithInferedType.Enumerations, formGroup, probe)
		AssociationFieldToForm("MinInclusive", instanceWithInferedType.MinInclusive, formGroup, probe)
		AssociationFieldToForm("MaxInclusive", instanceWithInferedType.MaxInclusive, formGroup, probe)
		AssociationFieldToForm("Pattern", instanceWithInferedType.Pattern, formGroup, probe)
		AssociationFieldToForm("WhiteSpace", instanceWithInferedType.WhiteSpace, formGroup, probe)
		AssociationFieldToForm("MinLength", instanceWithInferedType.MinLength, formGroup, probe)
		AssociationFieldToForm("MaxLength", instanceWithInferedType.MaxLength, formGroup, probe)
		AssociationFieldToForm("Length", instanceWithInferedType.Length, formGroup, probe)
		AssociationFieldToForm("TotalDigit", instanceWithInferedType.TotalDigit, formGroup, probe)

	case *models.Schema:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Xs", instanceWithInferedType.Xs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		AssociationSliceToForm("SimpleTypes", instanceWithInferedType, &instanceWithInferedType.SimpleTypes, formGroup, probe)
		AssociationSliceToForm("ComplexTypes", instanceWithInferedType, &instanceWithInferedType.ComplexTypes, formGroup, probe)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)

	case *models.Sequence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ModelGroupElements", instanceWithInferedType, &instanceWithInferedType.ModelGroupElements, formGroup, probe)

	case *models.SimpleContent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Extension", instanceWithInferedType.Extension, formGroup, probe)
		AssociationFieldToForm("Restriction", instanceWithInferedType.Restriction, formGroup, probe)

	case *models.SimpleType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Restriction", instanceWithInferedType.Restriction, formGroup, probe)
		AssociationFieldToForm("Union", instanceWithInferedType.Union, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "SimpleTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Schema),
					"SimpleTypes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Schema, *models.SimpleType](
					nil,
					"SimpleTypes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TotalDigit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Union:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("MemberTypes", instanceWithInferedType.MemberTypes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.WhiteSpace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
