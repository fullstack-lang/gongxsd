// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "ALTERNATIVE_ID":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			nil,
			probe,
			formGroup,
		)
		alternative_id := new(models.ALTERNATIVE_ID)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(alternative_id, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_boolean := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_boolean, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_DATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_date := new(models.ATTRIBUTE_DEFINITION_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_date, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_enumeration := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_enumeration, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_integer := new(models.ATTRIBUTE_DEFINITION_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_integer, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_REAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_real := new(models.ATTRIBUTE_DEFINITION_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_real, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_STRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_string := new(models.ATTRIBUTE_DEFINITION_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_string, formGroup, probe)
	case "ATTRIBUTE_DEFINITION_XHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_definition_xhtml := new(models.ATTRIBUTE_DEFINITION_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_definition_xhtml, formGroup, probe)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_boolean := new(models.ATTRIBUTE_VALUE_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_boolean, formGroup, probe)
	case "ATTRIBUTE_VALUE_DATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_date := new(models.ATTRIBUTE_VALUE_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_date, formGroup, probe)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_enumeration := new(models.ATTRIBUTE_VALUE_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_enumeration, formGroup, probe)
	case "ATTRIBUTE_VALUE_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_integer := new(models.ATTRIBUTE_VALUE_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_integer, formGroup, probe)
	case "ATTRIBUTE_VALUE_REAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_real := new(models.ATTRIBUTE_VALUE_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_real, formGroup, probe)
	case "ATTRIBUTE_VALUE_STRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_string := new(models.ATTRIBUTE_VALUE_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_string, formGroup, probe)
	case "ATTRIBUTE_VALUE_XHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute_value_xhtml := new(models.ATTRIBUTE_VALUE_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute_value_xhtml, formGroup, probe)
	case "A_ALTERNATIVE_ID":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ALTERNATIVE_ID Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ALTERNATIVE_IDFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_alternative_id := new(models.A_ALTERNATIVE_ID)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_alternative_id, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_BOOLEAN_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_boolean_ref := new(models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_boolean_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_DATE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_DATE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_DATE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_date_ref := new(models.A_ATTRIBUTE_DEFINITION_DATE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_date_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_ENUMERATION_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_enumeration_ref := new(models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_enumeration_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_INTEGER_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_INTEGER_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_integer_ref := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_integer_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_REAL_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_REAL_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_REAL_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_real_ref := new(models.A_ATTRIBUTE_DEFINITION_REAL_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_real_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_STRING_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_STRING_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_STRING_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_string_ref := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_string_ref, formGroup, probe)
	case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_DEFINITION_XHTML_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_DEFINITION_XHTML_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_definition_xhtml_ref := new(models.A_ATTRIBUTE_DEFINITION_XHTML_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_definition_xhtml_ref, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_boolean := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_boolean, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_DATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_date := new(models.A_ATTRIBUTE_VALUE_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_date, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_enumeration := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_enumeration, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_integer := new(models.A_ATTRIBUTE_VALUE_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_integer, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_REAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_real := new(models.A_ATTRIBUTE_VALUE_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_real, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_STRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_string := new(models.A_ATTRIBUTE_VALUE_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_string, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_XHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_xhtml := new(models.A_ATTRIBUTE_VALUE_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_xhtml, formGroup, probe)
	case "A_ATTRIBUTE_VALUE_XHTML_1":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ATTRIBUTE_VALUE_XHTML_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ATTRIBUTE_VALUE_XHTML_1FormCallback(
			nil,
			probe,
			formGroup,
		)
		a_attribute_value_xhtml_1 := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_attribute_value_xhtml_1, formGroup, probe)
	case "A_CHILDREN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_CHILDREN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_CHILDRENFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_children := new(models.A_CHILDREN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_children, formGroup, probe)
	case "A_CORE_CONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_CORE_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_CORE_CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_core_content := new(models.A_CORE_CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_core_content, formGroup, probe)
	case "A_DATATYPES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPESFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatypes := new(models.A_DATATYPES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatypes, formGroup, probe)
	case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_BOOLEAN_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_BOOLEAN_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_boolean_ref := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_boolean_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_DATE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_DATE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_DATE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_date_ref := new(models.A_DATATYPE_DEFINITION_DATE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_date_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_ENUMERATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_ENUMERATION_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_enumeration_ref := new(models.A_DATATYPE_DEFINITION_ENUMERATION_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_enumeration_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_INTEGER_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_INTEGER_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_INTEGER_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_integer_ref := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_integer_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_REAL_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_REAL_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_REAL_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_real_ref := new(models.A_DATATYPE_DEFINITION_REAL_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_real_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_STRING_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_STRING_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_STRING_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_string_ref := new(models.A_DATATYPE_DEFINITION_STRING_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_string_ref, formGroup, probe)
	case "A_DATATYPE_DEFINITION_XHTML_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_DATATYPE_DEFINITION_XHTML_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_DATATYPE_DEFINITION_XHTML_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_datatype_definition_xhtml_ref := new(models.A_DATATYPE_DEFINITION_XHTML_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_datatype_definition_xhtml_ref, formGroup, probe)
	case "A_EDITABLE_ATTS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_EDITABLE_ATTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_EDITABLE_ATTSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_editable_atts := new(models.A_EDITABLE_ATTS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_editable_atts, formGroup, probe)
	case "A_ENUM_VALUE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_ENUM_VALUE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_ENUM_VALUE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_enum_value_ref := new(models.A_ENUM_VALUE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_enum_value_ref, formGroup, probe)
	case "A_OBJECT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_OBJECTFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_object := new(models.A_OBJECT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_object, formGroup, probe)
	case "A_PROPERTIES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_PROPERTIES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_PROPERTIESFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_properties := new(models.A_PROPERTIES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_properties, formGroup, probe)
	case "A_RELATION_GROUP_TYPE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_RELATION_GROUP_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_RELATION_GROUP_TYPE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_relation_group_type_ref := new(models.A_RELATION_GROUP_TYPE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_relation_group_type_ref, formGroup, probe)
	case "A_SOURCE_1":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SOURCE_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SOURCE_1FormCallback(
			nil,
			probe,
			formGroup,
		)
		a_source_1 := new(models.A_SOURCE_1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_source_1, formGroup, probe)
	case "A_SPECIFICATIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPECIFICATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFICATIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_specifications := new(models.A_SPECIFICATIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_specifications, formGroup, probe)
	case "A_SPECIFICATION_TYPE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPECIFICATION_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFICATION_TYPE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_specification_type_ref := new(models.A_SPECIFICATION_TYPE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_specification_type_ref, formGroup, probe)
	case "A_SPECIFIED_VALUES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPECIFIED_VALUES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPECIFIED_VALUESFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_specified_values := new(models.A_SPECIFIED_VALUES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_specified_values, formGroup, probe)
	case "A_SPEC_ATTRIBUTES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_ATTRIBUTES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_ATTRIBUTESFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_attributes := new(models.A_SPEC_ATTRIBUTES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_attributes, formGroup, probe)
	case "A_SPEC_OBJECTS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_OBJECTS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_OBJECTSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_objects := new(models.A_SPEC_OBJECTS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_objects, formGroup, probe)
	case "A_SPEC_OBJECT_TYPE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_OBJECT_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_OBJECT_TYPE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_object_type_ref := new(models.A_SPEC_OBJECT_TYPE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_object_type_ref, formGroup, probe)
	case "A_SPEC_RELATIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_RELATIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_relations := new(models.A_SPEC_RELATIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_relations, formGroup, probe)
	case "A_SPEC_RELATION_GROUPS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_RELATION_GROUPS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_relation_groups := new(models.A_SPEC_RELATION_GROUPS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_relation_groups, formGroup, probe)
	case "A_SPEC_RELATION_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_RELATION_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_relation_ref := new(models.A_SPEC_RELATION_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_relation_ref, formGroup, probe)
	case "A_SPEC_RELATION_TYPE_REF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_RELATION_TYPE_REF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_RELATION_TYPE_REFFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_relation_type_ref := new(models.A_SPEC_RELATION_TYPE_REF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_relation_type_ref, formGroup, probe)
	case "A_SPEC_TYPES":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_SPEC_TYPES Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_SPEC_TYPESFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_spec_types := new(models.A_SPEC_TYPES)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_spec_types, formGroup, probe)
	case "A_TARGET_SPECIFICATION_1":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_TARGET_SPECIFICATION_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TARGET_SPECIFICATION_1FormCallback(
			nil,
			probe,
			formGroup,
		)
		a_target_specification_1 := new(models.A_TARGET_SPECIFICATION_1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_target_specification_1, formGroup, probe)
	case "A_THE_HEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_THE_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_THE_HEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_the_header := new(models.A_THE_HEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_the_header, formGroup, probe)
	case "A_TOOL_EXTENSIONS":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "A_TOOL_EXTENSIONS Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_TOOL_EXTENSIONSFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_tool_extensions := new(models.A_TOOL_EXTENSIONS)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_tool_extensions, formGroup, probe)
	case "DATATYPE_DEFINITION_BOOLEAN":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_BOOLEAN Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_boolean := new(models.DATATYPE_DEFINITION_BOOLEAN)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_boolean, formGroup, probe)
	case "DATATYPE_DEFINITION_DATE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_DATE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_date := new(models.DATATYPE_DEFINITION_DATE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_date, formGroup, probe)
	case "DATATYPE_DEFINITION_ENUMERATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_ENUMERATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_enumeration := new(models.DATATYPE_DEFINITION_ENUMERATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_enumeration, formGroup, probe)
	case "DATATYPE_DEFINITION_INTEGER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_INTEGER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_integer := new(models.DATATYPE_DEFINITION_INTEGER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_integer, formGroup, probe)
	case "DATATYPE_DEFINITION_REAL":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_REAL Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_real := new(models.DATATYPE_DEFINITION_REAL)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_real, formGroup, probe)
	case "DATATYPE_DEFINITION_STRING":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_STRING Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_string := new(models.DATATYPE_DEFINITION_STRING)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_string, formGroup, probe)
	case "DATATYPE_DEFINITION_XHTML":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "DATATYPE_DEFINITION_XHTML Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			nil,
			probe,
			formGroup,
		)
		datatype_definition_xhtml := new(models.DATATYPE_DEFINITION_XHTML)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(datatype_definition_xhtml, formGroup, probe)
	case "EMBEDDED_VALUE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "EMBEDDED_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		embedded_value := new(models.EMBEDDED_VALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(embedded_value, formGroup, probe)
	case "ENUM_VALUE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ENUM_VALUE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			nil,
			probe,
			formGroup,
		)
		enum_value := new(models.ENUM_VALUE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(enum_value, formGroup, probe)
	case "RELATION_GROUP":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RELATION_GROUP Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			nil,
			probe,
			formGroup,
		)
		relation_group := new(models.RELATION_GROUP)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relation_group, formGroup, probe)
	case "RELATION_GROUP_TYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RELATION_GROUP_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		relation_group_type := new(models.RELATION_GROUP_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(relation_group_type, formGroup, probe)
	case "REQ_IF":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQ_IF Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IFFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if := new(models.REQ_IF)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if, formGroup, probe)
	case "REQ_IF_CONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQ_IF_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_content, formGroup, probe)
	case "REQ_IF_HEADER":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQ_IF_HEADER Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_header, formGroup, probe)
	case "REQ_IF_TOOL_EXTENSION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "REQ_IF_TOOL_EXTENSION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		req_if_tool_extension := new(models.REQ_IF_TOOL_EXTENSION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(req_if_tool_extension, formGroup, probe)
	case "SPECIFICATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFICATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		specification := new(models.SPECIFICATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specification, formGroup, probe)
	case "SPECIFICATION_TYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPECIFICATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		specification_type := new(models.SPECIFICATION_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(specification_type, formGroup, probe)
	case "SPEC_HIERARCHY":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPEC_HIERARCHY Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_hierarchy, formGroup, probe)
	case "SPEC_OBJECT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPEC_OBJECT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_object := new(models.SPEC_OBJECT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_object, formGroup, probe)
	case "SPEC_OBJECT_TYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPEC_OBJECT_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_object_type := new(models.SPEC_OBJECT_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_object_type, formGroup, probe)
	case "SPEC_RELATION":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPEC_RELATION Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_relation := new(models.SPEC_RELATION)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_relation, formGroup, probe)
	case "SPEC_RELATION_TYPE":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SPEC_RELATION_TYPE Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			nil,
			probe,
			formGroup,
		)
		spec_relation_type := new(models.SPEC_RELATION_TYPE)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(spec_relation_type, formGroup, probe)
	case "XHTML_CONTENT":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "XHTML_CONTENT Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			nil,
			probe,
			formGroup,
		)
		xhtml_content := new(models.XHTML_CONTENT)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xhtml_content, formGroup, probe)
	}
	formStage.Commit()
}
