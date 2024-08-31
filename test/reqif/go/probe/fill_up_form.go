// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ALTERNATIVE_ID"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ALTERNATIVE_ID),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ALTERNATIVE_ID, *models.ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
					nil,
					"ATTRIBUTE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_DATE](
					nil,
					"ATTRIBUTE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MULTI_VALUED", instanceWithInferedType.MULTI_VALUED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
					nil,
					"ATTRIBUTE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_INTEGER](
					nil,
					"ATTRIBUTE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_REAL](
					nil,
					"ATTRIBUTE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_STRING](
					nil,
					"ATTRIBUTE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("DEFAULT_VALUE", instanceWithInferedType, &instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES),
					"ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_XHTML](
					nil,
					"ATTRIBUTE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_BOOLEAN"
			rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_BOOLEAN),
					"ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_BOOLEAN, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"ATTRIBUTE_VALUE_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_DATE"
			rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_DATE),
					"ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_DATE, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_DATE](
					nil,
					"ATTRIBUTE_VALUE_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		AssociationSliceToForm("VALUES", instanceWithInferedType, &instanceWithInferedType.VALUES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_ENUMERATION"
			rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_ENUMERATION),
					"ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_ENUMERATION, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"ATTRIBUTE_VALUE_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_INTEGER"
			rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_INTEGER),
					"ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_INTEGER, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_INTEGER](
					nil,
					"ATTRIBUTE_VALUE_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_REAL"
			rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_REAL),
					"ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_REAL, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_REAL](
					nil,
					"ATTRIBUTE_VALUE_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_STRING"
			rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_STRING),
					"ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_STRING, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_STRING](
					nil,
					"ATTRIBUTE_VALUE_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_SIMPLIFIED", instanceWithInferedType.IS_SIMPLIFIED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("THE_VALUE", instanceWithInferedType, &instanceWithInferedType.THE_VALUE, formGroup, probe)
		AssociationSliceToForm("THE_ORIGINAL_VALUE", instanceWithInferedType, &instanceWithInferedType.THE_ORIGINAL_VALUE, formGroup, probe)
		AssociationSliceToForm("DEFINITION", instanceWithInferedType, &instanceWithInferedType.DEFINITION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML),
					"ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_ATTRIBUTE_VALUE_XHTML_1),
					"ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_XHTML](
					nil,
					"ATTRIBUTE_VALUE_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_BOOLEAN, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_DATE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_ENUMERATION, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_INTEGER, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_REAL, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_STRING, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_XHTML, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_BOOLEAN),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_BOOLEAN, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_DATE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_DATE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_ENUMERATION, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_INTEGER),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_INTEGER, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_REAL),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_REAL, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_STRING),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_STRING, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_XHTML),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_XHTML, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ENUM_VALUE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ENUM_VALUE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.A_ALTERNATIVE_ID](
					nil,
					"ALTERNATIVE_ID",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_BOOLEAN"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_BOOLEAN),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_BOOLEAN, *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_DATE"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_DATE),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_DATE, *models.A_ATTRIBUTE_DEFINITION_DATE_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_ENUMERATION"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_ENUMERATION),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_ENUMERATION, *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_INTEGER"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_INTEGER),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_INTEGER, *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_REAL"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_REAL),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_REAL, *models.A_ATTRIBUTE_DEFINITION_REAL_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_STRING"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_STRING),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_STRING, *models.A_ATTRIBUTE_DEFINITION_STRING_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML),
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_XHTML, *models.A_ATTRIBUTE_DEFINITION_XHTML_REF](
					nil,
					"DEFINITION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_BOOLEAN, *models.A_ATTRIBUTE_VALUE_BOOLEAN](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_DATE, *models.A_ATTRIBUTE_VALUE_DATE](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_ENUMERATION, *models.A_ATTRIBUTE_VALUE_ENUMERATION](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_INTEGER, *models.A_ATTRIBUTE_VALUE_INTEGER](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_REAL, *models.A_ATTRIBUTE_VALUE_REAL](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_STRING, *models.A_ATTRIBUTE_VALUE_STRING](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML),
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_XHTML, *models.A_ATTRIBUTE_VALUE_XHTML](
					nil,
					"DEFAULT_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.A_ATTRIBUTE_VALUE_XHTML_1](
					nil,
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.A_ATTRIBUTE_VALUE_XHTML_1](
					nil,
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.A_ATTRIBUTE_VALUE_XHTML_1](
					nil,
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_CHILDREN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPEC_HIERARCHY", instanceWithInferedType, &instanceWithInferedType.SPEC_HIERARCHY, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "CHILDREN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"CHILDREN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.A_CHILDREN](
					nil,
					"CHILDREN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "CHILDREN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"CHILDREN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.A_CHILDREN](
					nil,
					"CHILDREN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_CORE_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("REQ_IF_CONTENT", instanceWithInferedType, &instanceWithInferedType.REQ_IF_CONTENT, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "CORE_CONTENT"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF),
					"CORE_CONTENT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF, *models.A_CORE_CONTENT](
					nil,
					"CORE_CONTENT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DATATYPE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_XHTML, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"DATATYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_DATATYPES](
					nil,
					"DATATYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_BOOLEAN, *models.A_DATATYPE_DEFINITION_BOOLEAN_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_DATE_REF", instanceWithInferedType.DATATYPE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_DATE, *models.A_DATATYPE_DEFINITION_DATE_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_ENUMERATION, *models.A_DATATYPE_DEFINITION_ENUMERATION_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_INTEGER_REF", instanceWithInferedType.DATATYPE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_INTEGER, *models.A_DATATYPE_DEFINITION_INTEGER_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_REAL_REF", instanceWithInferedType.DATATYPE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_REAL, *models.A_DATATYPE_DEFINITION_REAL_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_STRING_REF", instanceWithInferedType.DATATYPE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_STRING, *models.A_DATATYPE_DEFINITION_STRING_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_XHTML_REF", instanceWithInferedType.DATATYPE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_DEFINITION_XHTML, *models.A_DATATYPE_DEFINITION_XHTML_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_EDITABLE_ATTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "EDITABLE_ATTS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"EDITABLE_ATTS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.A_EDITABLE_ATTS](
					nil,
					"EDITABLE_ATTS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_ENUM_VALUE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ENUM_VALUE_REF", instanceWithInferedType.ENUM_VALUE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_ENUMERATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_ENUMERATION),
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_ENUMERATION, *models.A_ENUM_VALUE_REF](
					nil,
					"VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "OBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_HIERARCHY),
					"OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_HIERARCHY, *models.A_OBJECT](
					nil,
					"OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_PROPERTIES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("EMBEDDED_VALUE", instanceWithInferedType, &instanceWithInferedType.EMBEDDED_VALUE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "PROPERTIES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ENUM_VALUE),
					"PROPERTIES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ENUM_VALUE, *models.A_PROPERTIES](
					nil,
					"PROPERTIES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_RELATION_GROUP_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RELATION_GROUP_TYPE_REF", instanceWithInferedType.RELATION_GROUP_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.A_RELATION_GROUP_TYPE_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPECIFICATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.SPECIFICATION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPECIFICATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPECIFICATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_SPECIFICATIONS](
					nil,
					"SPECIFICATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPECIFICATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATION_TYPE_REF", instanceWithInferedType.SPECIFICATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION, *models.A_SPECIFICATION_TYPE_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPECIFIED_VALUES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ENUM_VALUE", instanceWithInferedType, &instanceWithInferedType.ENUM_VALUE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "SPECIFIED_VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION),
					"SPECIFIED_VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DATATYPE_DEFINITION_ENUMERATION, *models.A_SPECIFIED_VALUES](
					nil,
					"SPECIFIED_VALUES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_ATTRIBUTES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP_TYPE),
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP_TYPE, *models.A_SPEC_ATTRIBUTES](
					nil,
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPECIFICATION_TYPE),
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPECIFICATION_TYPE, *models.A_SPEC_ATTRIBUTES](
					nil,
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT_TYPE),
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT_TYPE, *models.A_SPEC_ATTRIBUTES](
					nil,
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION_TYPE),
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION_TYPE, *models.A_SPEC_ATTRIBUTES](
					nil,
					"SPEC_ATTRIBUTES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_OBJECTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPEC_OBJECT", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECT, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_OBJECTS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_OBJECTS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_SPEC_OBJECTS](
					nil,
					"SPEC_OBJECTS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_OBJECT_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_TYPE_REF", instanceWithInferedType.SPEC_OBJECT_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_OBJECT),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_OBJECT, *models.A_SPEC_OBJECT_TYPE_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_RELATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SPEC_RELATION", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_RELATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_SPEC_RELATIONS](
					nil,
					"SPEC_RELATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_RELATION_GROUPS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RELATION_GROUP", instanceWithInferedType, &instanceWithInferedType.RELATION_GROUP, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATION_GROUPS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_RELATION_GROUPS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_SPEC_RELATION_GROUPS](
					nil,
					"SPEC_RELATION_GROUPS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_RELATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_RELATION_REF", instanceWithInferedType.SPEC_RELATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "SPEC_RELATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"SPEC_RELATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.A_SPEC_RELATION_REF](
					nil,
					"SPEC_RELATIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_RELATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_RELATION_TYPE_REF", instanceWithInferedType.SPEC_RELATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.A_SPEC_RELATION_TYPE_REF](
					nil,
					"TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_SPEC_TYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RELATION_GROUP_TYPE", instanceWithInferedType, &instanceWithInferedType.RELATION_GROUP_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_OBJECT_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECT_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION_TYPE, formGroup, probe)
		AssociationSliceToForm("SPECIFICATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPECIFICATION_TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF_CONTENT),
					"SPEC_TYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF_CONTENT, *models.A_SPEC_TYPES](
					nil,
					"SPEC_TYPES",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_TARGET_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "SOURCE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"SOURCE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.A_TARGET_1](
					nil,
					"SOURCE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "TARGET"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.SPEC_RELATION),
					"TARGET",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.SPEC_RELATION, *models.A_TARGET_1](
					nil,
					"TARGET",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_TARGET_SPECIFICATION_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATION_REF", instanceWithInferedType.SPECIFICATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "SOURCE_SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"SOURCE_SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.A_TARGET_SPECIFICATION_1](
					nil,
					"SOURCE_SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "TARGET_SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.RELATION_GROUP),
					"TARGET_SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.RELATION_GROUP, *models.A_TARGET_SPECIFICATION_1](
					nil,
					"TARGET_SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_THE_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("REQ_IF_HEADER", instanceWithInferedType, &instanceWithInferedType.REQ_IF_HEADER, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "THE_HEADER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF),
					"THE_HEADER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF, *models.A_THE_HEADER](
					nil,
					"THE_HEADER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.A_TOOL_EXTENSIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("REQ_IF_TOOL_EXTENSION", instanceWithInferedType, &instanceWithInferedType.REQ_IF_TOOL_EXTENSION, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "TOOL_EXTENSIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.REQ_IF),
					"TOOL_EXTENSIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.REQ_IF, *models.A_TOOL_EXTENSIONS](
					nil,
					"TOOL_EXTENSIONS",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_BOOLEAN](
					nil,
					"DATATYPE_DEFINITION_BOOLEAN",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_DATE](
					nil,
					"DATATYPE_DEFINITION_DATE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPECIFIED_VALUES", instanceWithInferedType, &instanceWithInferedType.SPECIFIED_VALUES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_ENUMERATION](
					nil,
					"DATATYPE_DEFINITION_ENUMERATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAX", instanceWithInferedType.MAX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MIN", instanceWithInferedType.MIN, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_INTEGER](
					nil,
					"DATATYPE_DEFINITION_INTEGER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ACCURACY", instanceWithInferedType.ACCURACY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAX", instanceWithInferedType.MAX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MIN", instanceWithInferedType.MIN, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_REAL](
					nil,
					"DATATYPE_DEFINITION_REAL",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MAX_LENGTH", instanceWithInferedType.MAX_LENGTH, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_STRING](
					nil,
					"DATATYPE_DEFINITION_STRING",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_DATATYPES),
					"DATATYPE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_XHTML](
					nil,
					"DATATYPE_DEFINITION_XHTML",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.EMBEDDED_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("KEY", instanceWithInferedType.KEY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OTHER_CONTENT", instanceWithInferedType.OTHER_CONTENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_PROPERTIES"
			rf.Fieldname = "EMBEDDED_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_PROPERTIES),
					"EMBEDDED_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_PROPERTIES, *models.EMBEDDED_VALUE](
					nil,
					"EMBEDDED_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ENUM_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("PROPERTIES", instanceWithInferedType, &instanceWithInferedType.PROPERTIES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPECIFIED_VALUES"
			rf.Fieldname = "ENUM_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPECIFIED_VALUES),
					"ENUM_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPECIFIED_VALUES, *models.ENUM_VALUE](
					nil,
					"ENUM_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RELATION_GROUP:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SOURCE_SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.SOURCE_SPECIFICATION, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATIONS", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationSliceToForm("TARGET_SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.TARGET_SPECIFICATION, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_RELATION_GROUPS"
			rf.Fieldname = "RELATION_GROUP"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_RELATION_GROUPS),
					"RELATION_GROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_RELATION_GROUPS, *models.RELATION_GROUP](
					nil,
					"RELATION_GROUP",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.RELATION_GROUP_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "RELATION_GROUP_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_TYPES),
					"RELATION_GROUP_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_TYPES, *models.RELATION_GROUP_TYPE](
					nil,
					"RELATION_GROUP_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQ_IF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("THE_HEADER", instanceWithInferedType, &instanceWithInferedType.THE_HEADER, formGroup, probe)
		AssociationSliceToForm("CORE_CONTENT", instanceWithInferedType, &instanceWithInferedType.CORE_CONTENT, formGroup, probe)
		AssociationSliceToForm("TOOL_EXTENSIONS", instanceWithInferedType, &instanceWithInferedType.TOOL_EXTENSIONS, formGroup, probe)

	case *models.REQ_IF_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DATATYPES", instanceWithInferedType, &instanceWithInferedType.DATATYPES, formGroup, probe)
		AssociationSliceToForm("SPEC_TYPES", instanceWithInferedType, &instanceWithInferedType.SPEC_TYPES, formGroup, probe)
		AssociationSliceToForm("SPEC_OBJECTS", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECTS, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATIONS", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationSliceToForm("SPECIFICATIONS", instanceWithInferedType, &instanceWithInferedType.SPECIFICATIONS, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATION_GROUPS", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION_GROUPS, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_CORE_CONTENT"
			rf.Fieldname = "REQ_IF_CONTENT"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_CORE_CONTENT),
					"REQ_IF_CONTENT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_CORE_CONTENT, *models.REQ_IF_CONTENT](
					nil,
					"REQ_IF_CONTENT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQ_IF_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("COMMENT", instanceWithInferedType.COMMENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CREATION_TIME", instanceWithInferedType.CREATION_TIME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REPOSITORY_ID", instanceWithInferedType.REPOSITORY_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_TOOL_ID", instanceWithInferedType.REQ_IF_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("REQ_IF_VERSION", instanceWithInferedType.REQ_IF_VERSION, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SOURCE_TOOL_ID", instanceWithInferedType.SOURCE_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TITLE", instanceWithInferedType.TITLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_THE_HEADER"
			rf.Fieldname = "REQ_IF_HEADER"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_THE_HEADER),
					"REQ_IF_HEADER",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_THE_HEADER, *models.REQ_IF_HEADER](
					nil,
					"REQ_IF_HEADER",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.REQ_IF_TOOL_EXTENSION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_TOOL_EXTENSIONS"
			rf.Fieldname = "REQ_IF_TOOL_EXTENSION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_TOOL_EXTENSIONS),
					"REQ_IF_TOOL_EXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_TOOL_EXTENSIONS, *models.REQ_IF_TOOL_EXTENSION](
					nil,
					"REQ_IF_TOOL_EXTENSION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("CHILDREN", instanceWithInferedType, &instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationSliceToForm("VALUES", instanceWithInferedType, &instanceWithInferedType.VALUES, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPECIFICATIONS"
			rf.Fieldname = "SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPECIFICATIONS),
					"SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPECIFICATIONS, *models.SPECIFICATION](
					nil,
					"SPECIFICATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPECIFICATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPECIFICATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_TYPES),
					"SPECIFICATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_TYPES, *models.SPECIFICATION_TYPE](
					nil,
					"SPECIFICATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_HIERARCHY:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_TABLE_INTERNAL", instanceWithInferedType.IS_TABLE_INTERNAL, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("CHILDREN", instanceWithInferedType, &instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationSliceToForm("EDITABLE_ATTS", instanceWithInferedType, &instanceWithInferedType.EDITABLE_ATTS, formGroup, probe)
		AssociationSliceToForm("OBJECT", instanceWithInferedType, &instanceWithInferedType.OBJECT, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_CHILDREN"
			rf.Fieldname = "SPEC_HIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_CHILDREN),
					"SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_CHILDREN, *models.SPEC_HIERARCHY](
					nil,
					"SPEC_HIERARCHY",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("VALUES", instanceWithInferedType, &instanceWithInferedType.VALUES, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_OBJECTS"
			rf.Fieldname = "SPEC_OBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_OBJECTS),
					"SPEC_OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_OBJECTS, *models.SPEC_OBJECT](
					nil,
					"SPEC_OBJECT",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_OBJECT_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPEC_OBJECT_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_TYPES),
					"SPEC_OBJECT_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_TYPES, *models.SPEC_OBJECT_TYPE](
					nil,
					"SPEC_OBJECT_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_RELATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("VALUES", instanceWithInferedType, &instanceWithInferedType.VALUES, formGroup, probe)
		AssociationSliceToForm("SOURCE", instanceWithInferedType, &instanceWithInferedType.SOURCE, formGroup, probe)
		AssociationSliceToForm("TARGET", instanceWithInferedType, &instanceWithInferedType.TARGET, formGroup, probe)
		AssociationSliceToForm("TYPE", instanceWithInferedType, &instanceWithInferedType.TYPE, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_RELATIONS"
			rf.Fieldname = "SPEC_RELATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_RELATIONS),
					"SPEC_RELATION",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_RELATIONS, *models.SPEC_RELATION](
					nil,
					"SPEC_RELATION",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SPEC_RELATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ALTERNATIVE_ID", instanceWithInferedType, &instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationSliceToForm("SPEC_ATTRIBUTES", instanceWithInferedType, &instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPEC_RELATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.A_SPEC_TYPES),
					"SPEC_RELATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.A_SPEC_TYPES, *models.SPEC_RELATION_TYPE](
					nil,
					"SPEC_RELATION_TYPE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.XHTML_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML),
					"THE_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_XHTML, *models.XHTML_CONTENT](
					nil,
					"THE_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_ORIGINAL_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML),
					"THE_ORIGINAL_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.ATTRIBUTE_VALUE_XHTML, *models.XHTML_CONTENT](
					nil,
					"THE_ORIGINAL_VALUE",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
