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
	case *models.ComplexType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXML", instanceWithInferedType.NameXML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Sequence", instanceWithInferedType.Sequence, formGroup, probe)
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

	case *models.Element:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SimpleType", instanceWithInferedType.SimpleType, formGroup, probe)
		AssociationFieldToForm("ComplexType", instanceWithInferedType.ComplexType, formGroup, probe)
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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Sequence),
					"Elements",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Sequence, *models.Element](
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

	case *models.Restriction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base", instanceWithInferedType.Base, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Enumerations", instanceWithInferedType, &instanceWithInferedType.Enumerations, formGroup, probe)

	case *models.Schema:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		AssociationSliceToForm("SimpleTypes", instanceWithInferedType, &instanceWithInferedType.SimpleTypes, formGroup, probe)
		AssociationSliceToForm("ComplexTypes", instanceWithInferedType, &instanceWithInferedType.ComplexTypes, formGroup, probe)

	case *models.Sequence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)

	case *models.SimpleType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Restriction", instanceWithInferedType.Restriction, formGroup, probe)
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

	default:
		_ = instanceWithInferedType
	}
}
