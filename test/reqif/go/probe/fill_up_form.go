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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IS_SIMPLIFIED", instanceWithInferedType.IS_SIMPLIFIED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("THE_VALUE", instanceWithInferedType.THE_VALUE, formGroup, probe)
		AssociationFieldToForm("THE_ORIGINAL_VALUE", instanceWithInferedType.THE_ORIGINAL_VALUE, formGroup, probe)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)

	case *models.A_ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)

	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)

	case *models.A_CHILDREN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SPEC_HIERARCHY", instanceWithInferedType.SPEC_HIERARCHY, formGroup, probe)

	case *models.A_CORE_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQ_IF_CONTENT", instanceWithInferedType.REQ_IF_CONTENT, formGroup, probe)

	case *models.A_DATATYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DATATYPE_DEFINITION_BOOLEAN", instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_DATE", instanceWithInferedType.DATATYPE_DEFINITION_DATE, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_ENUMERATION", instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_INTEGER", instanceWithInferedType.DATATYPE_DEFINITION_INTEGER, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_REAL", instanceWithInferedType.DATATYPE_DEFINITION_REAL, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_STRING", instanceWithInferedType.DATATYPE_DEFINITION_STRING, formGroup, probe)
		AssociationFieldToForm("DATATYPE_DEFINITION_XHTML", instanceWithInferedType.DATATYPE_DEFINITION_XHTML, formGroup, probe)

	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_DATE_REF", instanceWithInferedType.DATATYPE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_INTEGER_REF", instanceWithInferedType.DATATYPE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_REAL_REF", instanceWithInferedType.DATATYPE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_STRING_REF", instanceWithInferedType.DATATYPE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DATATYPE_DEFINITION_XHTML_REF", instanceWithInferedType.DATATYPE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

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

	case *models.A_ENUM_VALUE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ENUM_VALUE_REF", instanceWithInferedType.ENUM_VALUE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_PROPERTIES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("EMBEDDED_VALUE", instanceWithInferedType.EMBEDDED_VALUE, formGroup, probe)

	case *models.A_RELATION_GROUP_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RELATION_GROUP_TYPE_REF", instanceWithInferedType.RELATION_GROUP_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SOURCE_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SOURCE_SPECIFICATION_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATION_REF", instanceWithInferedType.SPECIFICATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SPECIFICATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SPECIFICATION", instanceWithInferedType.SPECIFICATION, formGroup, probe)

	case *models.A_SPECIFICATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPECIFICATION_TYPE_REF", instanceWithInferedType.SPECIFICATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SPECIFIED_VALUES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ENUM_VALUE", instanceWithInferedType.ENUM_VALUE, formGroup, probe)

	case *models.A_SPEC_ATTRIBUTES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationFieldToForm("ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)

	case *models.A_SPEC_OBJECTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SPEC_OBJECT", instanceWithInferedType.SPEC_OBJECT, formGroup, probe)

	case *models.A_SPEC_OBJECT_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_OBJECT_TYPE_REF", instanceWithInferedType.SPEC_OBJECT_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SPEC_RELATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SPEC_RELATION", instanceWithInferedType.SPEC_RELATION, formGroup, probe)

	case *models.A_SPEC_RELATION_GROUPS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("RELATION_GROUP", instanceWithInferedType.RELATION_GROUP, formGroup, probe)

	case *models.A_SPEC_RELATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_RELATION_REF", instanceWithInferedType.SPEC_RELATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SPEC_RELATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SPEC_RELATION_TYPE_REF", instanceWithInferedType.SPEC_RELATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.A_SPEC_TYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("RELATION_GROUP_TYPE", instanceWithInferedType.RELATION_GROUP_TYPE, formGroup, probe)
		AssociationFieldToForm("SPEC_OBJECT_TYPE", instanceWithInferedType.SPEC_OBJECT_TYPE, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATION_TYPE", instanceWithInferedType.SPEC_RELATION_TYPE, formGroup, probe)
		AssociationFieldToForm("SPECIFICATION_TYPE", instanceWithInferedType.SPECIFICATION_TYPE, formGroup, probe)

	case *models.A_THE_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQ_IF_HEADER", instanceWithInferedType.REQ_IF_HEADER, formGroup, probe)

	case *models.A_TOOL_EXTENSIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("REQ_IF_TOOL_EXTENSION", instanceWithInferedType.REQ_IF_TOOL_EXTENSION, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPECIFIED_VALUES", instanceWithInferedType.SPECIFIED_VALUES, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)

	case *models.EMBEDDED_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("KEY", instanceWithInferedType.KEY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OTHER_CONTENT", instanceWithInferedType.OTHER_CONTENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("PROPERTIES", instanceWithInferedType.PROPERTIES, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SOURCE_SPECIFICATION", instanceWithInferedType.SOURCE_SPECIFICATION, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATIONS", instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationFieldToForm("TARGET_SPECIFICATION", instanceWithInferedType.TARGET_SPECIFICATION, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)

	case *models.REQ_IF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("THE_HEADER", instanceWithInferedType.THE_HEADER, formGroup, probe)
		AssociationFieldToForm("CORE_CONTENT", instanceWithInferedType.CORE_CONTENT, formGroup, probe)
		AssociationFieldToForm("TOOL_EXTENSIONS", instanceWithInferedType.TOOL_EXTENSIONS, formGroup, probe)

	case *models.REQ_IF_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DATATYPES", instanceWithInferedType.DATATYPES, formGroup, probe)
		AssociationFieldToForm("SPEC_TYPES", instanceWithInferedType.SPEC_TYPES, formGroup, probe)
		AssociationFieldToForm("SPEC_OBJECTS", instanceWithInferedType.SPEC_OBJECTS, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATIONS", instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationFieldToForm("SPECIFICATIONS", instanceWithInferedType.SPECIFICATIONS, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATION_GROUPS", instanceWithInferedType.SPEC_RELATION_GROUPS, formGroup, probe)

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

	case *models.REQ_IF_TOOL_EXTENSION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("EDITABLE_ATTS", instanceWithInferedType.EDITABLE_ATTS, formGroup, probe)
		AssociationFieldToForm("OBJECT", instanceWithInferedType.OBJECT, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("SOURCE", instanceWithInferedType.SOURCE, formGroup, probe)
		AssociationFieldToForm("TARGET", instanceWithInferedType.TARGET, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)

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
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)

	case *models.XHTML_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
