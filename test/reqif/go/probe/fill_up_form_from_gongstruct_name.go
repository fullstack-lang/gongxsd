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
