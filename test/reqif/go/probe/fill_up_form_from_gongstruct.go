// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ALTERNATIVE_ID:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ALTERNATIVE_IDFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_DATE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_INTEGER_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_REAL_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_STRING_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_DEFINITION_XHTML_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ATTRIBUTE_VALUE_XHTML_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTML_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_CHILDREN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_CHILDREN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_CHILDRENFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_CORE_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_CORE_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_CORE_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_BOOLEAN_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_DATE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_DATE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_ENUMERATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_INTEGER_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_INTEGER_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_REAL_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_REAL_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_STRING_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_STRING_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DATATYPE_DEFINITION_XHTML_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_XHTML_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_EDITABLE_ATTS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_EDITABLE_ATTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_EDITABLE_ATTSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_ENUM_VALUE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_ENUM_VALUE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ENUM_VALUE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_OBJECT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_OBJECTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_PROPERTIES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_PROPERTIES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_PROPERTIESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_RELATION_GROUP_TYPE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_RELATION_GROUP_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_RELATION_GROUP_TYPE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SOURCE_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SOURCE_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SOURCE_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SOURCE_SPECIFICATION_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SOURCE_SPECIFICATION_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SOURCE_SPECIFICATION_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPECIFICATIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPECIFICATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFICATIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPECIFICATION_TYPE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPECIFICATION_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFICATION_TYPE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPECIFIED_VALUES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPECIFIED_VALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFIED_VALUESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_ATTRIBUTES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_ATTRIBUTES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_ATTRIBUTESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_OBJECTS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_OBJECTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_OBJECTSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_OBJECT_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_OBJECT_TYPE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_RELATIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_RELATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_RELATION_GROUPS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_RELATION_GROUPS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_RELATION_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_RELATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_RELATION_TYPE_REF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_RELATION_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_TYPE_REFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SPEC_TYPES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_TYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_TYPESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_THE_HEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_THE_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_THE_HEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TOOL_EXTENSIONS:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TOOL_EXTENSIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TOOL_EXTENSIONSFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_DATE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_REAL:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_STRING:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DATATYPE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.EMBEDDED_VALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "EMBEDDED_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ENUM_VALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ENUM_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATION_GROUP:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATION_GROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RELATION_GROUP_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RELATION_GROUP_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IFFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_HEADER:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "REQ_IF_TOOL_EXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPECIFICATION_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPECIFICATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_HIERARCHY:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_HIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_OBJECT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_OBJECT_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_OBJECT_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_RELATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_RELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SPEC_RELATION_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SPEC_RELATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.XHTML_CONTENT:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "XHTML_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
