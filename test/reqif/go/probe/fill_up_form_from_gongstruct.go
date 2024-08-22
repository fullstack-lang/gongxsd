// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

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
	case *models.A_DEFAULT_VALUE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_2:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_3:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_3FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_4:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_4 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_4FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_5:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_5 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_5FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFAULT_VALUE_6:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFAULT_VALUE_6 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFAULT_VALUE_6FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITIONFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_2:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_3:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_3FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_4:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_4 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_4FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_5:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_5 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_5FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_DEFINITION_6:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_DEFINITION_6 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DEFINITION_6FormCallback(
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
	case *models.A_SOURCE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SOURCE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SOURCEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_SOURCE_SPECIFICATION:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SOURCE_SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SOURCE_SPECIFICATIONFormCallback(
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
	case *models.A_SPEC_RELATIONS_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_SPEC_RELATIONS_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATIONS_1FormCallback(
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
	case *models.A_TYPE:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPEFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_10:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_10 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_10FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_2:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_2 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_2FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_3:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_3 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_3FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_4:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_4 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_4FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_5:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_5 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_5FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_6:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_6 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_6FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_7:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_7 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_7FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_8:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_8 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_8FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_TYPE_9:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_TYPE_9 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TYPE_9FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_VALUES:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_VALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_VALUESFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_VALUES_1:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "A_VALUES_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_VALUES_1FormCallback(
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
