// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ALTERNATIVE_IDFormCallback(
	alternative_id *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *table.FormGroup,
) (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) {
	alternative_idFormCallback = new(ALTERNATIVE_IDFormCallback)
	alternative_idFormCallback.probe = probe
	alternative_idFormCallback.alternative_id = alternative_id
	alternative_idFormCallback.formGroup = formGroup

	alternative_idFormCallback.CreationMode = (alternative_id == nil)

	return
}

type ALTERNATIVE_IDFormCallback struct {
	alternative_id *models.ALTERNATIVE_ID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (alternative_idFormCallback *ALTERNATIVE_IDFormCallback) OnSave() {

	log.Println("ALTERNATIVE_IDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	alternative_idFormCallback.probe.formStage.Checkout()

	if alternative_idFormCallback.alternative_id == nil {
		alternative_idFormCallback.alternative_id = new(models.ALTERNATIVE_ID).Stage(alternative_idFormCallback.probe.stageOfInterest)
	}
	alternative_id_ := alternative_idFormCallback.alternative_id
	_ = alternative_id_

	for _, formDiv := range alternative_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(alternative_id_.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(alternative_id_.IDENTIFIER), formDiv)
		}
	}

	// manage the suppress operation
	if alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_id_.Unstage(alternative_idFormCallback.probe.stageOfInterest)
	}

	alternative_idFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ALTERNATIVE_ID](
		alternative_idFormCallback.probe,
	)
	alternative_idFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if alternative_idFormCallback.CreationMode || alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		alternative_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(alternative_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ALTERNATIVE_IDFormCallback(
			nil,
			alternative_idFormCallback.probe,
			newFormGroup,
		)
		alternative_id := new(models.ALTERNATIVE_ID)
		FillUpForm(alternative_id, newFormGroup, alternative_idFormCallback.probe)
		alternative_idFormCallback.probe.formStage.Commit()
	}

	fillUpTree(alternative_idFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) {
	attribute_definition_booleanFormCallback = new(ATTRIBUTE_DEFINITION_BOOLEANFormCallback)
	attribute_definition_booleanFormCallback.probe = probe
	attribute_definition_booleanFormCallback.attribute_definition_boolean = attribute_definition_boolean
	attribute_definition_booleanFormCallback.formGroup = formGroup

	attribute_definition_booleanFormCallback.CreationMode = (attribute_definition_boolean == nil)

	return
}

type ATTRIBUTE_DEFINITION_BOOLEANFormCallback struct {
	attribute_definition_boolean *models.ATTRIBUTE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_booleanFormCallback *ATTRIBUTE_DEFINITION_BOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_booleanFormCallback.probe.formStage.Checkout()

	if attribute_definition_booleanFormCallback.attribute_definition_boolean == nil {
		attribute_definition_booleanFormCallback.attribute_definition_boolean = new(models.ATTRIBUTE_DEFINITION_BOOLEAN).Stage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_definition_boolean_ := attribute_definition_booleanFormCallback.attribute_definition_boolean
	_ = attribute_definition_boolean_

	for _, formDiv := range attribute_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_boolean_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_boolean_.Unstage(attribute_definition_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](
		attribute_definition_booleanFormCallback.probe,
	)
	attribute_definition_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_booleanFormCallback.CreationMode || attribute_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_BOOLEANFormCallback(
			nil,
			attribute_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_boolean := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
		FillUpForm(attribute_definition_boolean, newFormGroup, attribute_definition_booleanFormCallback.probe)
		attribute_definition_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_booleanFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) {
	attribute_definition_dateFormCallback = new(ATTRIBUTE_DEFINITION_DATEFormCallback)
	attribute_definition_dateFormCallback.probe = probe
	attribute_definition_dateFormCallback.attribute_definition_date = attribute_definition_date
	attribute_definition_dateFormCallback.formGroup = formGroup

	attribute_definition_dateFormCallback.CreationMode = (attribute_definition_date == nil)

	return
}

type ATTRIBUTE_DEFINITION_DATEFormCallback struct {
	attribute_definition_date *models.ATTRIBUTE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_dateFormCallback *ATTRIBUTE_DEFINITION_DATEFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_dateFormCallback.probe.formStage.Checkout()

	if attribute_definition_dateFormCallback.attribute_definition_date == nil {
		attribute_definition_dateFormCallback.attribute_definition_date = new(models.ATTRIBUTE_DEFINITION_DATE).Stage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}
	attribute_definition_date_ := attribute_definition_dateFormCallback.attribute_definition_date
	_ = attribute_definition_date_

	for _, formDiv := range attribute_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_date_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_date_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_date_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_date_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_date_.Unstage(attribute_definition_dateFormCallback.probe.stageOfInterest)
	}

	attribute_definition_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](
		attribute_definition_dateFormCallback.probe,
	)
	attribute_definition_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_dateFormCallback.CreationMode || attribute_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_DATEFormCallback(
			nil,
			attribute_definition_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_date := new(models.ATTRIBUTE_DEFINITION_DATE)
		FillUpForm(attribute_definition_date, newFormGroup, attribute_definition_dateFormCallback.probe)
		attribute_definition_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_dateFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) {
	attribute_definition_enumerationFormCallback = new(ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback)
	attribute_definition_enumerationFormCallback.probe = probe
	attribute_definition_enumerationFormCallback.attribute_definition_enumeration = attribute_definition_enumeration
	attribute_definition_enumerationFormCallback.formGroup = formGroup

	attribute_definition_enumerationFormCallback.CreationMode = (attribute_definition_enumeration == nil)

	return
}

type ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback struct {
	attribute_definition_enumeration *models.ATTRIBUTE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_enumerationFormCallback *ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_definition_enumerationFormCallback.attribute_definition_enumeration == nil {
		attribute_definition_enumerationFormCallback.attribute_definition_enumeration = new(models.ATTRIBUTE_DEFINITION_ENUMERATION).Stage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_definition_enumeration_ := attribute_definition_enumerationFormCallback.attribute_definition_enumeration
	_ = attribute_definition_enumeration_

	for _, formDiv := range attribute_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.LONG_NAME), formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&(attribute_definition_enumeration_.MULTI_VALUED), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumeration_.Unstage(attribute_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](
		attribute_definition_enumerationFormCallback.probe,
	)
	attribute_definition_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_enumerationFormCallback.CreationMode || attribute_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			attribute_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_enumeration := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
		FillUpForm(attribute_definition_enumeration, newFormGroup, attribute_definition_enumerationFormCallback.probe)
		attribute_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_enumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) {
	attribute_definition_integerFormCallback = new(ATTRIBUTE_DEFINITION_INTEGERFormCallback)
	attribute_definition_integerFormCallback.probe = probe
	attribute_definition_integerFormCallback.attribute_definition_integer = attribute_definition_integer
	attribute_definition_integerFormCallback.formGroup = formGroup

	attribute_definition_integerFormCallback.CreationMode = (attribute_definition_integer == nil)

	return
}

type ATTRIBUTE_DEFINITION_INTEGERFormCallback struct {
	attribute_definition_integer *models.ATTRIBUTE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_integerFormCallback *ATTRIBUTE_DEFINITION_INTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_integerFormCallback.probe.formStage.Checkout()

	if attribute_definition_integerFormCallback.attribute_definition_integer == nil {
		attribute_definition_integerFormCallback.attribute_definition_integer = new(models.ATTRIBUTE_DEFINITION_INTEGER).Stage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}
	attribute_definition_integer_ := attribute_definition_integerFormCallback.attribute_definition_integer
	_ = attribute_definition_integer_

	for _, formDiv := range attribute_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_integer_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_integer_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_integer_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integer_.Unstage(attribute_definition_integerFormCallback.probe.stageOfInterest)
	}

	attribute_definition_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](
		attribute_definition_integerFormCallback.probe,
	)
	attribute_definition_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_integerFormCallback.CreationMode || attribute_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_INTEGERFormCallback(
			nil,
			attribute_definition_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_integer := new(models.ATTRIBUTE_DEFINITION_INTEGER)
		FillUpForm(attribute_definition_integer, newFormGroup, attribute_definition_integerFormCallback.probe)
		attribute_definition_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_integerFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) {
	attribute_definition_realFormCallback = new(ATTRIBUTE_DEFINITION_REALFormCallback)
	attribute_definition_realFormCallback.probe = probe
	attribute_definition_realFormCallback.attribute_definition_real = attribute_definition_real
	attribute_definition_realFormCallback.formGroup = formGroup

	attribute_definition_realFormCallback.CreationMode = (attribute_definition_real == nil)

	return
}

type ATTRIBUTE_DEFINITION_REALFormCallback struct {
	attribute_definition_real *models.ATTRIBUTE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_realFormCallback *ATTRIBUTE_DEFINITION_REALFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_realFormCallback.probe.formStage.Checkout()

	if attribute_definition_realFormCallback.attribute_definition_real == nil {
		attribute_definition_realFormCallback.attribute_definition_real = new(models.ATTRIBUTE_DEFINITION_REAL).Stage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}
	attribute_definition_real_ := attribute_definition_realFormCallback.attribute_definition_real
	_ = attribute_definition_real_

	for _, formDiv := range attribute_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_real_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_real_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_real_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_real_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_real_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_real_.Unstage(attribute_definition_realFormCallback.probe.stageOfInterest)
	}

	attribute_definition_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](
		attribute_definition_realFormCallback.probe,
	)
	attribute_definition_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_realFormCallback.CreationMode || attribute_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_REALFormCallback(
			nil,
			attribute_definition_realFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_real := new(models.ATTRIBUTE_DEFINITION_REAL)
		FillUpForm(attribute_definition_real, newFormGroup, attribute_definition_realFormCallback.probe)
		attribute_definition_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_realFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) {
	attribute_definition_stringFormCallback = new(ATTRIBUTE_DEFINITION_STRINGFormCallback)
	attribute_definition_stringFormCallback.probe = probe
	attribute_definition_stringFormCallback.attribute_definition_string = attribute_definition_string
	attribute_definition_stringFormCallback.formGroup = formGroup

	attribute_definition_stringFormCallback.CreationMode = (attribute_definition_string == nil)

	return
}

type ATTRIBUTE_DEFINITION_STRINGFormCallback struct {
	attribute_definition_string *models.ATTRIBUTE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_stringFormCallback *ATTRIBUTE_DEFINITION_STRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_stringFormCallback.probe.formStage.Checkout()

	if attribute_definition_stringFormCallback.attribute_definition_string == nil {
		attribute_definition_stringFormCallback.attribute_definition_string = new(models.ATTRIBUTE_DEFINITION_STRING).Stage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}
	attribute_definition_string_ := attribute_definition_stringFormCallback.attribute_definition_string
	_ = attribute_definition_string_

	for _, formDiv := range attribute_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_string_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_string_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_string_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_string_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_string_.Unstage(attribute_definition_stringFormCallback.probe.stageOfInterest)
	}

	attribute_definition_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](
		attribute_definition_stringFormCallback.probe,
	)
	attribute_definition_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_stringFormCallback.CreationMode || attribute_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_STRINGFormCallback(
			nil,
			attribute_definition_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_string := new(models.ATTRIBUTE_DEFINITION_STRING)
		FillUpForm(attribute_definition_string, newFormGroup, attribute_definition_stringFormCallback.probe)
		attribute_definition_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_stringFormCallback.probe)
}
func __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) {
	attribute_definition_xhtmlFormCallback = new(ATTRIBUTE_DEFINITION_XHTMLFormCallback)
	attribute_definition_xhtmlFormCallback.probe = probe
	attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = attribute_definition_xhtml
	attribute_definition_xhtmlFormCallback.formGroup = formGroup

	attribute_definition_xhtmlFormCallback.CreationMode = (attribute_definition_xhtml == nil)

	return
}

type ATTRIBUTE_DEFINITION_XHTMLFormCallback struct {
	attribute_definition_xhtml *models.ATTRIBUTE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_definition_xhtmlFormCallback *ATTRIBUTE_DEFINITION_XHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_definition_xhtmlFormCallback.attribute_definition_xhtml == nil {
		attribute_definition_xhtmlFormCallback.attribute_definition_xhtml = new(models.ATTRIBUTE_DEFINITION_XHTML).Stage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_definition_xhtml_ := attribute_definition_xhtmlFormCallback.attribute_definition_xhtml
	_ = attribute_definition_xhtml_

	for _, formDiv := range attribute_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.IS_EDITABLE), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(attribute_definition_xhtml_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtml_.Unstage(attribute_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](
		attribute_definition_xhtmlFormCallback.probe,
	)
	attribute_definition_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_definition_xhtmlFormCallback.CreationMode || attribute_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_DEFINITION_XHTMLFormCallback(
			nil,
			attribute_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_definition_xhtml := new(models.ATTRIBUTE_DEFINITION_XHTML)
		FillUpForm(attribute_definition_xhtml, newFormGroup, attribute_definition_xhtmlFormCallback.probe)
		attribute_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_definition_xhtmlFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) {
	attribute_value_booleanFormCallback = new(ATTRIBUTE_VALUE_BOOLEANFormCallback)
	attribute_value_booleanFormCallback.probe = probe
	attribute_value_booleanFormCallback.attribute_value_boolean = attribute_value_boolean
	attribute_value_booleanFormCallback.formGroup = formGroup

	attribute_value_booleanFormCallback.CreationMode = (attribute_value_boolean == nil)

	return
}

type ATTRIBUTE_VALUE_BOOLEANFormCallback struct {
	attribute_value_boolean *models.ATTRIBUTE_VALUE_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_booleanFormCallback *ATTRIBUTE_VALUE_BOOLEANFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_booleanFormCallback.probe.formStage.Checkout()

	if attribute_value_booleanFormCallback.attribute_value_boolean == nil {
		attribute_value_booleanFormCallback.attribute_value_boolean = new(models.ATTRIBUTE_VALUE_BOOLEAN).Stage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}
	attribute_value_boolean_ := attribute_value_booleanFormCallback.attribute_value_boolean
	_ = attribute_value_boolean_

	for _, formDiv := range attribute_value_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_boolean_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_boolean_.THE_VALUE), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_boolean_.Unstage(attribute_value_booleanFormCallback.probe.stageOfInterest)
	}

	attribute_value_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](
		attribute_value_booleanFormCallback.probe,
	)
	attribute_value_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_booleanFormCallback.CreationMode || attribute_value_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_BOOLEANFormCallback(
			nil,
			attribute_value_booleanFormCallback.probe,
			newFormGroup,
		)
		attribute_value_boolean := new(models.ATTRIBUTE_VALUE_BOOLEAN)
		FillUpForm(attribute_value_boolean, newFormGroup, attribute_value_booleanFormCallback.probe)
		attribute_value_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_booleanFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) {
	attribute_value_dateFormCallback = new(ATTRIBUTE_VALUE_DATEFormCallback)
	attribute_value_dateFormCallback.probe = probe
	attribute_value_dateFormCallback.attribute_value_date = attribute_value_date
	attribute_value_dateFormCallback.formGroup = formGroup

	attribute_value_dateFormCallback.CreationMode = (attribute_value_date == nil)

	return
}

type ATTRIBUTE_VALUE_DATEFormCallback struct {
	attribute_value_date *models.ATTRIBUTE_VALUE_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_dateFormCallback *ATTRIBUTE_VALUE_DATEFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_dateFormCallback.probe.formStage.Checkout()

	if attribute_value_dateFormCallback.attribute_value_date == nil {
		attribute_value_dateFormCallback.attribute_value_date = new(models.ATTRIBUTE_VALUE_DATE).Stage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}
	attribute_value_date_ := attribute_value_dateFormCallback.attribute_value_date
	_ = attribute_value_date_

	for _, formDiv := range attribute_value_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_date_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_date_.THE_VALUE), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_date_.Unstage(attribute_value_dateFormCallback.probe.stageOfInterest)
	}

	attribute_value_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_DATE](
		attribute_value_dateFormCallback.probe,
	)
	attribute_value_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_dateFormCallback.CreationMode || attribute_value_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_DATEFormCallback(
			nil,
			attribute_value_dateFormCallback.probe,
			newFormGroup,
		)
		attribute_value_date := new(models.ATTRIBUTE_VALUE_DATE)
		FillUpForm(attribute_value_date, newFormGroup, attribute_value_dateFormCallback.probe)
		attribute_value_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_dateFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) {
	attribute_value_enumerationFormCallback = new(ATTRIBUTE_VALUE_ENUMERATIONFormCallback)
	attribute_value_enumerationFormCallback.probe = probe
	attribute_value_enumerationFormCallback.attribute_value_enumeration = attribute_value_enumeration
	attribute_value_enumerationFormCallback.formGroup = formGroup

	attribute_value_enumerationFormCallback.CreationMode = (attribute_value_enumeration == nil)

	return
}

type ATTRIBUTE_VALUE_ENUMERATIONFormCallback struct {
	attribute_value_enumeration *models.ATTRIBUTE_VALUE_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_enumerationFormCallback *ATTRIBUTE_VALUE_ENUMERATIONFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_enumerationFormCallback.probe.formStage.Checkout()

	if attribute_value_enumerationFormCallback.attribute_value_enumeration == nil {
		attribute_value_enumerationFormCallback.attribute_value_enumeration = new(models.ATTRIBUTE_VALUE_ENUMERATION).Stage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}
	attribute_value_enumeration_ := attribute_value_enumerationFormCallback.attribute_value_enumeration
	_ = attribute_value_enumeration_

	for _, formDiv := range attribute_value_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_enumeration_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumeration_.Unstage(attribute_value_enumerationFormCallback.probe.stageOfInterest)
	}

	attribute_value_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](
		attribute_value_enumerationFormCallback.probe,
	)
	attribute_value_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_enumerationFormCallback.CreationMode || attribute_value_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_ENUMERATIONFormCallback(
			nil,
			attribute_value_enumerationFormCallback.probe,
			newFormGroup,
		)
		attribute_value_enumeration := new(models.ATTRIBUTE_VALUE_ENUMERATION)
		FillUpForm(attribute_value_enumeration, newFormGroup, attribute_value_enumerationFormCallback.probe)
		attribute_value_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_enumerationFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) {
	attribute_value_integerFormCallback = new(ATTRIBUTE_VALUE_INTEGERFormCallback)
	attribute_value_integerFormCallback.probe = probe
	attribute_value_integerFormCallback.attribute_value_integer = attribute_value_integer
	attribute_value_integerFormCallback.formGroup = formGroup

	attribute_value_integerFormCallback.CreationMode = (attribute_value_integer == nil)

	return
}

type ATTRIBUTE_VALUE_INTEGERFormCallback struct {
	attribute_value_integer *models.ATTRIBUTE_VALUE_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_integerFormCallback *ATTRIBUTE_VALUE_INTEGERFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_integerFormCallback.probe.formStage.Checkout()

	if attribute_value_integerFormCallback.attribute_value_integer == nil {
		attribute_value_integerFormCallback.attribute_value_integer = new(models.ATTRIBUTE_VALUE_INTEGER).Stage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}
	attribute_value_integer_ := attribute_value_integerFormCallback.attribute_value_integer
	_ = attribute_value_integer_

	for _, formDiv := range attribute_value_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_integer_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_integer_.THE_VALUE), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integer_.Unstage(attribute_value_integerFormCallback.probe.stageOfInterest)
	}

	attribute_value_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](
		attribute_value_integerFormCallback.probe,
	)
	attribute_value_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_integerFormCallback.CreationMode || attribute_value_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_INTEGERFormCallback(
			nil,
			attribute_value_integerFormCallback.probe,
			newFormGroup,
		)
		attribute_value_integer := new(models.ATTRIBUTE_VALUE_INTEGER)
		FillUpForm(attribute_value_integer, newFormGroup, attribute_value_integerFormCallback.probe)
		attribute_value_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_integerFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) {
	attribute_value_realFormCallback = new(ATTRIBUTE_VALUE_REALFormCallback)
	attribute_value_realFormCallback.probe = probe
	attribute_value_realFormCallback.attribute_value_real = attribute_value_real
	attribute_value_realFormCallback.formGroup = formGroup

	attribute_value_realFormCallback.CreationMode = (attribute_value_real == nil)

	return
}

type ATTRIBUTE_VALUE_REALFormCallback struct {
	attribute_value_real *models.ATTRIBUTE_VALUE_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_realFormCallback *ATTRIBUTE_VALUE_REALFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_realFormCallback.probe.formStage.Checkout()

	if attribute_value_realFormCallback.attribute_value_real == nil {
		attribute_value_realFormCallback.attribute_value_real = new(models.ATTRIBUTE_VALUE_REAL).Stage(attribute_value_realFormCallback.probe.stageOfInterest)
	}
	attribute_value_real_ := attribute_value_realFormCallback.attribute_value_real
	_ = attribute_value_real_

	for _, formDiv := range attribute_value_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_real_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_real_.THE_VALUE), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_real_.Unstage(attribute_value_realFormCallback.probe.stageOfInterest)
	}

	attribute_value_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_REAL](
		attribute_value_realFormCallback.probe,
	)
	attribute_value_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_realFormCallback.CreationMode || attribute_value_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_REALFormCallback(
			nil,
			attribute_value_realFormCallback.probe,
			newFormGroup,
		)
		attribute_value_real := new(models.ATTRIBUTE_VALUE_REAL)
		FillUpForm(attribute_value_real, newFormGroup, attribute_value_realFormCallback.probe)
		attribute_value_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_realFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) {
	attribute_value_stringFormCallback = new(ATTRIBUTE_VALUE_STRINGFormCallback)
	attribute_value_stringFormCallback.probe = probe
	attribute_value_stringFormCallback.attribute_value_string = attribute_value_string
	attribute_value_stringFormCallback.formGroup = formGroup

	attribute_value_stringFormCallback.CreationMode = (attribute_value_string == nil)

	return
}

type ATTRIBUTE_VALUE_STRINGFormCallback struct {
	attribute_value_string *models.ATTRIBUTE_VALUE_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_stringFormCallback *ATTRIBUTE_VALUE_STRINGFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_stringFormCallback.probe.formStage.Checkout()

	if attribute_value_stringFormCallback.attribute_value_string == nil {
		attribute_value_stringFormCallback.attribute_value_string = new(models.ATTRIBUTE_VALUE_STRING).Stage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}
	attribute_value_string_ := attribute_value_stringFormCallback.attribute_value_string
	_ = attribute_value_string_

	for _, formDiv := range attribute_value_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_string_.Name), formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&(attribute_value_string_.THE_VALUE), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_string_.Unstage(attribute_value_stringFormCallback.probe.stageOfInterest)
	}

	attribute_value_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_STRING](
		attribute_value_stringFormCallback.probe,
	)
	attribute_value_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_stringFormCallback.CreationMode || attribute_value_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_STRINGFormCallback(
			nil,
			attribute_value_stringFormCallback.probe,
			newFormGroup,
		)
		attribute_value_string := new(models.ATTRIBUTE_VALUE_STRING)
		FillUpForm(attribute_value_string, newFormGroup, attribute_value_stringFormCallback.probe)
		attribute_value_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_stringFormCallback.probe)
}
func __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) {
	attribute_value_xhtmlFormCallback = new(ATTRIBUTE_VALUE_XHTMLFormCallback)
	attribute_value_xhtmlFormCallback.probe = probe
	attribute_value_xhtmlFormCallback.attribute_value_xhtml = attribute_value_xhtml
	attribute_value_xhtmlFormCallback.formGroup = formGroup

	attribute_value_xhtmlFormCallback.CreationMode = (attribute_value_xhtml == nil)

	return
}

type ATTRIBUTE_VALUE_XHTMLFormCallback struct {
	attribute_value_xhtml *models.ATTRIBUTE_VALUE_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attribute_value_xhtmlFormCallback *ATTRIBUTE_VALUE_XHTMLFormCallback) OnSave() {

	log.Println("ATTRIBUTE_VALUE_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attribute_value_xhtmlFormCallback.probe.formStage.Checkout()

	if attribute_value_xhtmlFormCallback.attribute_value_xhtml == nil {
		attribute_value_xhtmlFormCallback.attribute_value_xhtml = new(models.ATTRIBUTE_VALUE_XHTML).Stage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}
	attribute_value_xhtml_ := attribute_value_xhtmlFormCallback.attribute_value_xhtml
	_ = attribute_value_xhtml_

	for _, formDiv := range attribute_value_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.Name), formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&(attribute_value_xhtml_.IS_SIMPLIFIED), formDiv)
		}
	}

	// manage the suppress operation
	if attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtml_.Unstage(attribute_value_xhtmlFormCallback.probe.stageOfInterest)
	}

	attribute_value_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ATTRIBUTE_VALUE_XHTML](
		attribute_value_xhtmlFormCallback.probe,
	)
	attribute_value_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attribute_value_xhtmlFormCallback.CreationMode || attribute_value_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_value_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attribute_value_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ATTRIBUTE_VALUE_XHTMLFormCallback(
			nil,
			attribute_value_xhtmlFormCallback.probe,
			newFormGroup,
		)
		attribute_value_xhtml := new(models.ATTRIBUTE_VALUE_XHTML)
		FillUpForm(attribute_value_xhtml, newFormGroup, attribute_value_xhtmlFormCallback.probe)
		attribute_value_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attribute_value_xhtmlFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) {
	datatype_definition_booleanFormCallback = new(DATATYPE_DEFINITION_BOOLEANFormCallback)
	datatype_definition_booleanFormCallback.probe = probe
	datatype_definition_booleanFormCallback.datatype_definition_boolean = datatype_definition_boolean
	datatype_definition_booleanFormCallback.formGroup = formGroup

	datatype_definition_booleanFormCallback.CreationMode = (datatype_definition_boolean == nil)

	return
}

type DATATYPE_DEFINITION_BOOLEANFormCallback struct {
	datatype_definition_boolean *models.DATATYPE_DEFINITION_BOOLEAN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_booleanFormCallback *DATATYPE_DEFINITION_BOOLEANFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_BOOLEANFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_booleanFormCallback.probe.formStage.Checkout()

	if datatype_definition_booleanFormCallback.datatype_definition_boolean == nil {
		datatype_definition_booleanFormCallback.datatype_definition_boolean = new(models.DATATYPE_DEFINITION_BOOLEAN).Stage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}
	datatype_definition_boolean_ := datatype_definition_booleanFormCallback.datatype_definition_boolean
	_ = datatype_definition_boolean_

	for _, formDiv := range datatype_definition_booleanFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_boolean_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_boolean_.Unstage(datatype_definition_booleanFormCallback.probe.stageOfInterest)
	}

	datatype_definition_booleanFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](
		datatype_definition_booleanFormCallback.probe,
	)
	datatype_definition_booleanFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_booleanFormCallback.CreationMode || datatype_definition_booleanFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_booleanFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_booleanFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_BOOLEANFormCallback(
			nil,
			datatype_definition_booleanFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_boolean := new(models.DATATYPE_DEFINITION_BOOLEAN)
		FillUpForm(datatype_definition_boolean, newFormGroup, datatype_definition_booleanFormCallback.probe)
		datatype_definition_booleanFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_booleanFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) {
	datatype_definition_dateFormCallback = new(DATATYPE_DEFINITION_DATEFormCallback)
	datatype_definition_dateFormCallback.probe = probe
	datatype_definition_dateFormCallback.datatype_definition_date = datatype_definition_date
	datatype_definition_dateFormCallback.formGroup = formGroup

	datatype_definition_dateFormCallback.CreationMode = (datatype_definition_date == nil)

	return
}

type DATATYPE_DEFINITION_DATEFormCallback struct {
	datatype_definition_date *models.DATATYPE_DEFINITION_DATE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_dateFormCallback *DATATYPE_DEFINITION_DATEFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_DATEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_dateFormCallback.probe.formStage.Checkout()

	if datatype_definition_dateFormCallback.datatype_definition_date == nil {
		datatype_definition_dateFormCallback.datatype_definition_date = new(models.DATATYPE_DEFINITION_DATE).Stage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}
	datatype_definition_date_ := datatype_definition_dateFormCallback.datatype_definition_date
	_ = datatype_definition_date_

	for _, formDiv := range datatype_definition_dateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_date_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_date_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_date_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_date_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_date_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_date_.Unstage(datatype_definition_dateFormCallback.probe.stageOfInterest)
	}

	datatype_definition_dateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_DATE](
		datatype_definition_dateFormCallback.probe,
	)
	datatype_definition_dateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_dateFormCallback.CreationMode || datatype_definition_dateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_dateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_dateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_DATEFormCallback(
			nil,
			datatype_definition_dateFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_date := new(models.DATATYPE_DEFINITION_DATE)
		FillUpForm(datatype_definition_date, newFormGroup, datatype_definition_dateFormCallback.probe)
		datatype_definition_dateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_dateFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) {
	datatype_definition_enumerationFormCallback = new(DATATYPE_DEFINITION_ENUMERATIONFormCallback)
	datatype_definition_enumerationFormCallback.probe = probe
	datatype_definition_enumerationFormCallback.datatype_definition_enumeration = datatype_definition_enumeration
	datatype_definition_enumerationFormCallback.formGroup = formGroup

	datatype_definition_enumerationFormCallback.CreationMode = (datatype_definition_enumeration == nil)

	return
}

type DATATYPE_DEFINITION_ENUMERATIONFormCallback struct {
	datatype_definition_enumeration *models.DATATYPE_DEFINITION_ENUMERATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_enumerationFormCallback *DATATYPE_DEFINITION_ENUMERATIONFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_ENUMERATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_enumerationFormCallback.probe.formStage.Checkout()

	if datatype_definition_enumerationFormCallback.datatype_definition_enumeration == nil {
		datatype_definition_enumerationFormCallback.datatype_definition_enumeration = new(models.DATATYPE_DEFINITION_ENUMERATION).Stage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}
	datatype_definition_enumeration_ := datatype_definition_enumerationFormCallback.datatype_definition_enumeration
	_ = datatype_definition_enumeration_

	for _, formDiv := range datatype_definition_enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_enumeration_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumeration_.Unstage(datatype_definition_enumerationFormCallback.probe.stageOfInterest)
	}

	datatype_definition_enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](
		datatype_definition_enumerationFormCallback.probe,
	)
	datatype_definition_enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_enumerationFormCallback.CreationMode || datatype_definition_enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_ENUMERATIONFormCallback(
			nil,
			datatype_definition_enumerationFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_enumeration := new(models.DATATYPE_DEFINITION_ENUMERATION)
		FillUpForm(datatype_definition_enumeration, newFormGroup, datatype_definition_enumerationFormCallback.probe)
		datatype_definition_enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_enumerationFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) {
	datatype_definition_integerFormCallback = new(DATATYPE_DEFINITION_INTEGERFormCallback)
	datatype_definition_integerFormCallback.probe = probe
	datatype_definition_integerFormCallback.datatype_definition_integer = datatype_definition_integer
	datatype_definition_integerFormCallback.formGroup = formGroup

	datatype_definition_integerFormCallback.CreationMode = (datatype_definition_integer == nil)

	return
}

type DATATYPE_DEFINITION_INTEGERFormCallback struct {
	datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_integerFormCallback *DATATYPE_DEFINITION_INTEGERFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_INTEGERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_integerFormCallback.probe.formStage.Checkout()

	if datatype_definition_integerFormCallback.datatype_definition_integer == nil {
		datatype_definition_integerFormCallback.datatype_definition_integer = new(models.DATATYPE_DEFINITION_INTEGER).Stage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}
	datatype_definition_integer_ := datatype_definition_integerFormCallback.datatype_definition_integer
	_ = datatype_definition_integer_

	for _, formDiv := range datatype_definition_integerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_integer_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_integer_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_integer_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_integer_.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(datatype_definition_integer_.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(datatype_definition_integer_.MIN), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integer_.Unstage(datatype_definition_integerFormCallback.probe.stageOfInterest)
	}

	datatype_definition_integerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_INTEGER](
		datatype_definition_integerFormCallback.probe,
	)
	datatype_definition_integerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_integerFormCallback.CreationMode || datatype_definition_integerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_integerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_integerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_INTEGERFormCallback(
			nil,
			datatype_definition_integerFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_integer := new(models.DATATYPE_DEFINITION_INTEGER)
		FillUpForm(datatype_definition_integer, newFormGroup, datatype_definition_integerFormCallback.probe)
		datatype_definition_integerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_integerFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_REALFormCallback(
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) {
	datatype_definition_realFormCallback = new(DATATYPE_DEFINITION_REALFormCallback)
	datatype_definition_realFormCallback.probe = probe
	datatype_definition_realFormCallback.datatype_definition_real = datatype_definition_real
	datatype_definition_realFormCallback.formGroup = formGroup

	datatype_definition_realFormCallback.CreationMode = (datatype_definition_real == nil)

	return
}

type DATATYPE_DEFINITION_REALFormCallback struct {
	datatype_definition_real *models.DATATYPE_DEFINITION_REAL

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_realFormCallback *DATATYPE_DEFINITION_REALFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_REALFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_realFormCallback.probe.formStage.Checkout()

	if datatype_definition_realFormCallback.datatype_definition_real == nil {
		datatype_definition_realFormCallback.datatype_definition_real = new(models.DATATYPE_DEFINITION_REAL).Stage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}
	datatype_definition_real_ := datatype_definition_realFormCallback.datatype_definition_real
	_ = datatype_definition_real_

	for _, formDiv := range datatype_definition_realFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_real_.Name), formDiv)
		case "ACCURACY":
			FormDivBasicFieldToField(&(datatype_definition_real_.ACCURACY), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_real_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_real_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_real_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_real_.LONG_NAME), formDiv)
		case "MAX":
			FormDivBasicFieldToField(&(datatype_definition_real_.MAX), formDiv)
		case "MIN":
			FormDivBasicFieldToField(&(datatype_definition_real_.MIN), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_real_.Unstage(datatype_definition_realFormCallback.probe.stageOfInterest)
	}

	datatype_definition_realFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_REAL](
		datatype_definition_realFormCallback.probe,
	)
	datatype_definition_realFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_realFormCallback.CreationMode || datatype_definition_realFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_realFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_realFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_REALFormCallback(
			nil,
			datatype_definition_realFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_real := new(models.DATATYPE_DEFINITION_REAL)
		FillUpForm(datatype_definition_real, newFormGroup, datatype_definition_realFormCallback.probe)
		datatype_definition_realFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_realFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) {
	datatype_definition_stringFormCallback = new(DATATYPE_DEFINITION_STRINGFormCallback)
	datatype_definition_stringFormCallback.probe = probe
	datatype_definition_stringFormCallback.datatype_definition_string = datatype_definition_string
	datatype_definition_stringFormCallback.formGroup = formGroup

	datatype_definition_stringFormCallback.CreationMode = (datatype_definition_string == nil)

	return
}

type DATATYPE_DEFINITION_STRINGFormCallback struct {
	datatype_definition_string *models.DATATYPE_DEFINITION_STRING

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_stringFormCallback *DATATYPE_DEFINITION_STRINGFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_STRINGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_stringFormCallback.probe.formStage.Checkout()

	if datatype_definition_stringFormCallback.datatype_definition_string == nil {
		datatype_definition_stringFormCallback.datatype_definition_string = new(models.DATATYPE_DEFINITION_STRING).Stage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}
	datatype_definition_string_ := datatype_definition_stringFormCallback.datatype_definition_string
	_ = datatype_definition_string_

	for _, formDiv := range datatype_definition_stringFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_string_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_string_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_string_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_string_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_string_.LONG_NAME), formDiv)
		case "MAX_LENGTH":
			FormDivBasicFieldToField(&(datatype_definition_string_.MAX_LENGTH), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_string_.Unstage(datatype_definition_stringFormCallback.probe.stageOfInterest)
	}

	datatype_definition_stringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_STRING](
		datatype_definition_stringFormCallback.probe,
	)
	datatype_definition_stringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_stringFormCallback.CreationMode || datatype_definition_stringFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_stringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_stringFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_STRINGFormCallback(
			nil,
			datatype_definition_stringFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_string := new(models.DATATYPE_DEFINITION_STRING)
		FillUpForm(datatype_definition_string, newFormGroup, datatype_definition_stringFormCallback.probe)
		datatype_definition_stringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_stringFormCallback.probe)
}
func __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *table.FormGroup,
) (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) {
	datatype_definition_xhtmlFormCallback = new(DATATYPE_DEFINITION_XHTMLFormCallback)
	datatype_definition_xhtmlFormCallback.probe = probe
	datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = datatype_definition_xhtml
	datatype_definition_xhtmlFormCallback.formGroup = formGroup

	datatype_definition_xhtmlFormCallback.CreationMode = (datatype_definition_xhtml == nil)

	return
}

type DATATYPE_DEFINITION_XHTMLFormCallback struct {
	datatype_definition_xhtml *models.DATATYPE_DEFINITION_XHTML

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (datatype_definition_xhtmlFormCallback *DATATYPE_DEFINITION_XHTMLFormCallback) OnSave() {

	log.Println("DATATYPE_DEFINITION_XHTMLFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datatype_definition_xhtmlFormCallback.probe.formStage.Checkout()

	if datatype_definition_xhtmlFormCallback.datatype_definition_xhtml == nil {
		datatype_definition_xhtmlFormCallback.datatype_definition_xhtml = new(models.DATATYPE_DEFINITION_XHTML).Stage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}
	datatype_definition_xhtml_ := datatype_definition_xhtmlFormCallback.datatype_definition_xhtml
	_ = datatype_definition_xhtml_

	for _, formDiv := range datatype_definition_xhtmlFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(datatype_definition_xhtml_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtml_.Unstage(datatype_definition_xhtmlFormCallback.probe.stageOfInterest)
	}

	datatype_definition_xhtmlFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DATATYPE_DEFINITION_XHTML](
		datatype_definition_xhtmlFormCallback.probe,
	)
	datatype_definition_xhtmlFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if datatype_definition_xhtmlFormCallback.CreationMode || datatype_definition_xhtmlFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datatype_definition_xhtmlFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(datatype_definition_xhtmlFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DATATYPE_DEFINITION_XHTMLFormCallback(
			nil,
			datatype_definition_xhtmlFormCallback.probe,
			newFormGroup,
		)
		datatype_definition_xhtml := new(models.DATATYPE_DEFINITION_XHTML)
		FillUpForm(datatype_definition_xhtml, newFormGroup, datatype_definition_xhtmlFormCallback.probe)
		datatype_definition_xhtmlFormCallback.probe.formStage.Commit()
	}

	fillUpTree(datatype_definition_xhtmlFormCallback.probe)
}
func __gong__New__EMBEDDED_VALUEFormCallback(
	embedded_value *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) {
	embedded_valueFormCallback = new(EMBEDDED_VALUEFormCallback)
	embedded_valueFormCallback.probe = probe
	embedded_valueFormCallback.embedded_value = embedded_value
	embedded_valueFormCallback.formGroup = formGroup

	embedded_valueFormCallback.CreationMode = (embedded_value == nil)

	return
}

type EMBEDDED_VALUEFormCallback struct {
	embedded_value *models.EMBEDDED_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (embedded_valueFormCallback *EMBEDDED_VALUEFormCallback) OnSave() {

	log.Println("EMBEDDED_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	embedded_valueFormCallback.probe.formStage.Checkout()

	if embedded_valueFormCallback.embedded_value == nil {
		embedded_valueFormCallback.embedded_value = new(models.EMBEDDED_VALUE).Stage(embedded_valueFormCallback.probe.stageOfInterest)
	}
	embedded_value_ := embedded_valueFormCallback.embedded_value
	_ = embedded_value_

	for _, formDiv := range embedded_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(embedded_value_.Name), formDiv)
		case "KEY":
			FormDivBasicFieldToField(&(embedded_value_.KEY), formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&(embedded_value_.OTHER_CONTENT), formDiv)
		}
	}

	// manage the suppress operation
	if embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_value_.Unstage(embedded_valueFormCallback.probe.stageOfInterest)
	}

	embedded_valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.EMBEDDED_VALUE](
		embedded_valueFormCallback.probe,
	)
	embedded_valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if embedded_valueFormCallback.CreationMode || embedded_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		embedded_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(embedded_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EMBEDDED_VALUEFormCallback(
			nil,
			embedded_valueFormCallback.probe,
			newFormGroup,
		)
		embedded_value := new(models.EMBEDDED_VALUE)
		FillUpForm(embedded_value, newFormGroup, embedded_valueFormCallback.probe)
		embedded_valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(embedded_valueFormCallback.probe)
}
func __gong__New__ENUM_VALUEFormCallback(
	enum_value *models.ENUM_VALUE,
	probe *Probe,
	formGroup *table.FormGroup,
) (enum_valueFormCallback *ENUM_VALUEFormCallback) {
	enum_valueFormCallback = new(ENUM_VALUEFormCallback)
	enum_valueFormCallback.probe = probe
	enum_valueFormCallback.enum_value = enum_value
	enum_valueFormCallback.formGroup = formGroup

	enum_valueFormCallback.CreationMode = (enum_value == nil)

	return
}

type ENUM_VALUEFormCallback struct {
	enum_value *models.ENUM_VALUE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enum_valueFormCallback *ENUM_VALUEFormCallback) OnSave() {

	log.Println("ENUM_VALUEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enum_valueFormCallback.probe.formStage.Checkout()

	if enum_valueFormCallback.enum_value == nil {
		enum_valueFormCallback.enum_value = new(models.ENUM_VALUE).Stage(enum_valueFormCallback.probe.stageOfInterest)
	}
	enum_value_ := enum_valueFormCallback.enum_value
	_ = enum_value_

	for _, formDiv := range enum_valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enum_value_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(enum_value_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(enum_value_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(enum_value_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(enum_value_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_value_.Unstage(enum_valueFormCallback.probe.stageOfInterest)
	}

	enum_valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ENUM_VALUE](
		enum_valueFormCallback.probe,
	)
	enum_valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if enum_valueFormCallback.CreationMode || enum_valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enum_valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(enum_valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ENUM_VALUEFormCallback(
			nil,
			enum_valueFormCallback.probe,
			newFormGroup,
		)
		enum_value := new(models.ENUM_VALUE)
		FillUpForm(enum_value, newFormGroup, enum_valueFormCallback.probe)
		enum_valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(enum_valueFormCallback.probe)
}
func __gong__New__RELATION_GROUPFormCallback(
	relation_group *models.RELATION_GROUP,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_groupFormCallback *RELATION_GROUPFormCallback) {
	relation_groupFormCallback = new(RELATION_GROUPFormCallback)
	relation_groupFormCallback.probe = probe
	relation_groupFormCallback.relation_group = relation_group
	relation_groupFormCallback.formGroup = formGroup

	relation_groupFormCallback.CreationMode = (relation_group == nil)

	return
}

type RELATION_GROUPFormCallback struct {
	relation_group *models.RELATION_GROUP

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_groupFormCallback *RELATION_GROUPFormCallback) OnSave() {

	log.Println("RELATION_GROUPFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_groupFormCallback.probe.formStage.Checkout()

	if relation_groupFormCallback.relation_group == nil {
		relation_groupFormCallback.relation_group = new(models.RELATION_GROUP).Stage(relation_groupFormCallback.probe.stageOfInterest)
	}
	relation_group_ := relation_groupFormCallback.relation_group
	_ = relation_group_

	for _, formDiv := range relation_groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(relation_group_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_.Unstage(relation_groupFormCallback.probe.stageOfInterest)
	}

	relation_groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATION_GROUP](
		relation_groupFormCallback.probe,
	)
	relation_groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relation_groupFormCallback.CreationMode || relation_groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relation_groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUPFormCallback(
			nil,
			relation_groupFormCallback.probe,
			newFormGroup,
		)
		relation_group := new(models.RELATION_GROUP)
		FillUpForm(relation_group, newFormGroup, relation_groupFormCallback.probe)
		relation_groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relation_groupFormCallback.probe)
}
func __gong__New__RELATION_GROUP_TYPEFormCallback(
	relation_group_type *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) {
	relation_group_typeFormCallback = new(RELATION_GROUP_TYPEFormCallback)
	relation_group_typeFormCallback.probe = probe
	relation_group_typeFormCallback.relation_group_type = relation_group_type
	relation_group_typeFormCallback.formGroup = formGroup

	relation_group_typeFormCallback.CreationMode = (relation_group_type == nil)

	return
}

type RELATION_GROUP_TYPEFormCallback struct {
	relation_group_type *models.RELATION_GROUP_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (relation_group_typeFormCallback *RELATION_GROUP_TYPEFormCallback) OnSave() {

	log.Println("RELATION_GROUP_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	relation_group_typeFormCallback.probe.formStage.Checkout()

	if relation_group_typeFormCallback.relation_group_type == nil {
		relation_group_typeFormCallback.relation_group_type = new(models.RELATION_GROUP_TYPE).Stage(relation_group_typeFormCallback.probe.stageOfInterest)
	}
	relation_group_type_ := relation_group_typeFormCallback.relation_group_type
	_ = relation_group_type_

	for _, formDiv := range relation_group_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(relation_group_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(relation_group_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(relation_group_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(relation_group_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(relation_group_type_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_type_.Unstage(relation_group_typeFormCallback.probe.stageOfInterest)
	}

	relation_group_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RELATION_GROUP_TYPE](
		relation_group_typeFormCallback.probe,
	)
	relation_group_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if relation_group_typeFormCallback.CreationMode || relation_group_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		relation_group_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(relation_group_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RELATION_GROUP_TYPEFormCallback(
			nil,
			relation_group_typeFormCallback.probe,
			newFormGroup,
		)
		relation_group_type := new(models.RELATION_GROUP_TYPE)
		FillUpForm(relation_group_type, newFormGroup, relation_group_typeFormCallback.probe)
		relation_group_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(relation_group_typeFormCallback.probe)
}
func __gong__New__REQ_IFFormCallback(
	req_if *models.REQ_IF,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_ifFormCallback *REQ_IFFormCallback) {
	req_ifFormCallback = new(REQ_IFFormCallback)
	req_ifFormCallback.probe = probe
	req_ifFormCallback.req_if = req_if
	req_ifFormCallback.formGroup = formGroup

	req_ifFormCallback.CreationMode = (req_if == nil)

	return
}

type REQ_IFFormCallback struct {
	req_if *models.REQ_IF

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_ifFormCallback *REQ_IFFormCallback) OnSave() {

	log.Println("REQ_IFFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_ifFormCallback.probe.formStage.Checkout()

	if req_ifFormCallback.req_if == nil {
		req_ifFormCallback.req_if = new(models.REQ_IF).Stage(req_ifFormCallback.probe.stageOfInterest)
	}
	req_if_ := req_ifFormCallback.req_if
	_ = req_if_

	for _, formDiv := range req_ifFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_.Name), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(req_if_.Lang), formDiv)
		}
	}

	// manage the suppress operation
	if req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_.Unstage(req_ifFormCallback.probe.stageOfInterest)
	}

	req_ifFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF](
		req_ifFormCallback.probe,
	)
	req_ifFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_ifFormCallback.CreationMode || req_ifFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_ifFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_ifFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IFFormCallback(
			nil,
			req_ifFormCallback.probe,
			newFormGroup,
		)
		req_if := new(models.REQ_IF)
		FillUpForm(req_if, newFormGroup, req_ifFormCallback.probe)
		req_ifFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_ifFormCallback.probe)
}
func __gong__New__REQ_IF_CONTENTFormCallback(
	req_if_content *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) {
	req_if_contentFormCallback = new(REQ_IF_CONTENTFormCallback)
	req_if_contentFormCallback.probe = probe
	req_if_contentFormCallback.req_if_content = req_if_content
	req_if_contentFormCallback.formGroup = formGroup

	req_if_contentFormCallback.CreationMode = (req_if_content == nil)

	return
}

type REQ_IF_CONTENTFormCallback struct {
	req_if_content *models.REQ_IF_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_contentFormCallback *REQ_IF_CONTENTFormCallback) OnSave() {

	log.Println("REQ_IF_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_contentFormCallback.probe.formStage.Checkout()

	if req_if_contentFormCallback.req_if_content == nil {
		req_if_contentFormCallback.req_if_content = new(models.REQ_IF_CONTENT).Stage(req_if_contentFormCallback.probe.stageOfInterest)
	}
	req_if_content_ := req_if_contentFormCallback.req_if_content
	_ = req_if_content_

	for _, formDiv := range req_if_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_content_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_content_.Unstage(req_if_contentFormCallback.probe.stageOfInterest)
	}

	req_if_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_CONTENT](
		req_if_contentFormCallback.probe,
	)
	req_if_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_contentFormCallback.CreationMode || req_if_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_CONTENTFormCallback(
			nil,
			req_if_contentFormCallback.probe,
			newFormGroup,
		)
		req_if_content := new(models.REQ_IF_CONTENT)
		FillUpForm(req_if_content, newFormGroup, req_if_contentFormCallback.probe)
		req_if_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_contentFormCallback.probe)
}
func __gong__New__REQ_IF_HEADERFormCallback(
	req_if_header *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) {
	req_if_headerFormCallback = new(REQ_IF_HEADERFormCallback)
	req_if_headerFormCallback.probe = probe
	req_if_headerFormCallback.req_if_header = req_if_header
	req_if_headerFormCallback.formGroup = formGroup

	req_if_headerFormCallback.CreationMode = (req_if_header == nil)

	return
}

type REQ_IF_HEADERFormCallback struct {
	req_if_header *models.REQ_IF_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_headerFormCallback *REQ_IF_HEADERFormCallback) OnSave() {

	log.Println("REQ_IF_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_headerFormCallback.probe.formStage.Checkout()

	if req_if_headerFormCallback.req_if_header == nil {
		req_if_headerFormCallback.req_if_header = new(models.REQ_IF_HEADER).Stage(req_if_headerFormCallback.probe.stageOfInterest)
	}
	req_if_header_ := req_if_headerFormCallback.req_if_header
	_ = req_if_header_

	for _, formDiv := range req_if_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_header_.Name), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(req_if_header_.IDENTIFIER), formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&(req_if_header_.COMMENT), formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&(req_if_header_.CREATION_TIME), formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&(req_if_header_.REPOSITORY_ID), formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_TOOL_ID), formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&(req_if_header_.REQ_IF_VERSION), formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&(req_if_header_.SOURCE_TOOL_ID), formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&(req_if_header_.TITLE), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_header_.Unstage(req_if_headerFormCallback.probe.stageOfInterest)
	}

	req_if_headerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_HEADER](
		req_if_headerFormCallback.probe,
	)
	req_if_headerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_headerFormCallback.CreationMode || req_if_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_HEADERFormCallback(
			nil,
			req_if_headerFormCallback.probe,
			newFormGroup,
		)
		req_if_header := new(models.REQ_IF_HEADER)
		FillUpForm(req_if_header, newFormGroup, req_if_headerFormCallback.probe)
		req_if_headerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_headerFormCallback.probe)
}
func __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *table.FormGroup,
) (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) {
	req_if_tool_extensionFormCallback = new(REQ_IF_TOOL_EXTENSIONFormCallback)
	req_if_tool_extensionFormCallback.probe = probe
	req_if_tool_extensionFormCallback.req_if_tool_extension = req_if_tool_extension
	req_if_tool_extensionFormCallback.formGroup = formGroup

	req_if_tool_extensionFormCallback.CreationMode = (req_if_tool_extension == nil)

	return
}

type REQ_IF_TOOL_EXTENSIONFormCallback struct {
	req_if_tool_extension *models.REQ_IF_TOOL_EXTENSION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (req_if_tool_extensionFormCallback *REQ_IF_TOOL_EXTENSIONFormCallback) OnSave() {

	log.Println("REQ_IF_TOOL_EXTENSIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	req_if_tool_extensionFormCallback.probe.formStage.Checkout()

	if req_if_tool_extensionFormCallback.req_if_tool_extension == nil {
		req_if_tool_extensionFormCallback.req_if_tool_extension = new(models.REQ_IF_TOOL_EXTENSION).Stage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}
	req_if_tool_extension_ := req_if_tool_extensionFormCallback.req_if_tool_extension
	_ = req_if_tool_extension_

	for _, formDiv := range req_if_tool_extensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(req_if_tool_extension_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extension_.Unstage(req_if_tool_extensionFormCallback.probe.stageOfInterest)
	}

	req_if_tool_extensionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.REQ_IF_TOOL_EXTENSION](
		req_if_tool_extensionFormCallback.probe,
	)
	req_if_tool_extensionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if req_if_tool_extensionFormCallback.CreationMode || req_if_tool_extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		req_if_tool_extensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(req_if_tool_extensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__REQ_IF_TOOL_EXTENSIONFormCallback(
			nil,
			req_if_tool_extensionFormCallback.probe,
			newFormGroup,
		)
		req_if_tool_extension := new(models.REQ_IF_TOOL_EXTENSION)
		FillUpForm(req_if_tool_extension, newFormGroup, req_if_tool_extensionFormCallback.probe)
		req_if_tool_extensionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(req_if_tool_extensionFormCallback.probe)
}
func __gong__New__SPECIFICATIONFormCallback(
	specification *models.SPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (specificationFormCallback *SPECIFICATIONFormCallback) {
	specificationFormCallback = new(SPECIFICATIONFormCallback)
	specificationFormCallback.probe = probe
	specificationFormCallback.specification = specification
	specificationFormCallback.formGroup = formGroup

	specificationFormCallback.CreationMode = (specification == nil)

	return
}

type SPECIFICATIONFormCallback struct {
	specification *models.SPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specificationFormCallback *SPECIFICATIONFormCallback) OnSave() {

	log.Println("SPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specificationFormCallback.probe.formStage.Checkout()

	if specificationFormCallback.specification == nil {
		specificationFormCallback.specification = new(models.SPECIFICATION).Stage(specificationFormCallback.probe.stageOfInterest)
	}
	specification_ := specificationFormCallback.specification
	_ = specification_

	for _, formDiv := range specificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(specification_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_.Unstage(specificationFormCallback.probe.stageOfInterest)
	}

	specificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATION](
		specificationFormCallback.probe,
	)
	specificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specificationFormCallback.CreationMode || specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATIONFormCallback(
			nil,
			specificationFormCallback.probe,
			newFormGroup,
		)
		specification := new(models.SPECIFICATION)
		FillUpForm(specification, newFormGroup, specificationFormCallback.probe)
		specificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specificationFormCallback.probe)
}
func __gong__New__SPECIFICATION_TYPEFormCallback(
	specification_type *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) {
	specification_typeFormCallback = new(SPECIFICATION_TYPEFormCallback)
	specification_typeFormCallback.probe = probe
	specification_typeFormCallback.specification_type = specification_type
	specification_typeFormCallback.formGroup = formGroup

	specification_typeFormCallback.CreationMode = (specification_type == nil)

	return
}

type SPECIFICATION_TYPEFormCallback struct {
	specification_type *models.SPECIFICATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (specification_typeFormCallback *SPECIFICATION_TYPEFormCallback) OnSave() {

	log.Println("SPECIFICATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	specification_typeFormCallback.probe.formStage.Checkout()

	if specification_typeFormCallback.specification_type == nil {
		specification_typeFormCallback.specification_type = new(models.SPECIFICATION_TYPE).Stage(specification_typeFormCallback.probe.stageOfInterest)
	}
	specification_type_ := specification_typeFormCallback.specification_type
	_ = specification_type_

	for _, formDiv := range specification_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(specification_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(specification_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(specification_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(specification_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(specification_type_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_type_.Unstage(specification_typeFormCallback.probe.stageOfInterest)
	}

	specification_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPECIFICATION_TYPE](
		specification_typeFormCallback.probe,
	)
	specification_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if specification_typeFormCallback.CreationMode || specification_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		specification_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(specification_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPECIFICATION_TYPEFormCallback(
			nil,
			specification_typeFormCallback.probe,
			newFormGroup,
		)
		specification_type := new(models.SPECIFICATION_TYPE)
		FillUpForm(specification_type, newFormGroup, specification_typeFormCallback.probe)
		specification_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(specification_typeFormCallback.probe)
}
func __gong__New__SPEC_HIERARCHYFormCallback(
	spec_hierarchy *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) {
	spec_hierarchyFormCallback = new(SPEC_HIERARCHYFormCallback)
	spec_hierarchyFormCallback.probe = probe
	spec_hierarchyFormCallback.spec_hierarchy = spec_hierarchy
	spec_hierarchyFormCallback.formGroup = formGroup

	spec_hierarchyFormCallback.CreationMode = (spec_hierarchy == nil)

	return
}

type SPEC_HIERARCHYFormCallback struct {
	spec_hierarchy *models.SPEC_HIERARCHY

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_hierarchyFormCallback *SPEC_HIERARCHYFormCallback) OnSave() {

	log.Println("SPEC_HIERARCHYFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_hierarchyFormCallback.probe.formStage.Checkout()

	if spec_hierarchyFormCallback.spec_hierarchy == nil {
		spec_hierarchyFormCallback.spec_hierarchy = new(models.SPEC_HIERARCHY).Stage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}
	spec_hierarchy_ := spec_hierarchyFormCallback.spec_hierarchy
	_ = spec_hierarchy_

	for _, formDiv := range spec_hierarchyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_hierarchy_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_hierarchy_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_hierarchy_.IDENTIFIER), formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_EDITABLE), formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&(spec_hierarchy_.IS_TABLE_INTERNAL), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_hierarchy_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_hierarchy_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchy_.Unstage(spec_hierarchyFormCallback.probe.stageOfInterest)
	}

	spec_hierarchyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_HIERARCHY](
		spec_hierarchyFormCallback.probe,
	)
	spec_hierarchyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_hierarchyFormCallback.CreationMode || spec_hierarchyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_hierarchyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_hierarchyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_HIERARCHYFormCallback(
			nil,
			spec_hierarchyFormCallback.probe,
			newFormGroup,
		)
		spec_hierarchy := new(models.SPEC_HIERARCHY)
		FillUpForm(spec_hierarchy, newFormGroup, spec_hierarchyFormCallback.probe)
		spec_hierarchyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_hierarchyFormCallback.probe)
}
func __gong__New__SPEC_OBJECTFormCallback(
	spec_object *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_objectFormCallback *SPEC_OBJECTFormCallback) {
	spec_objectFormCallback = new(SPEC_OBJECTFormCallback)
	spec_objectFormCallback.probe = probe
	spec_objectFormCallback.spec_object = spec_object
	spec_objectFormCallback.formGroup = formGroup

	spec_objectFormCallback.CreationMode = (spec_object == nil)

	return
}

type SPEC_OBJECTFormCallback struct {
	spec_object *models.SPEC_OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_objectFormCallback *SPEC_OBJECTFormCallback) OnSave() {

	log.Println("SPEC_OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_objectFormCallback.probe.formStage.Checkout()

	if spec_objectFormCallback.spec_object == nil {
		spec_objectFormCallback.spec_object = new(models.SPEC_OBJECT).Stage(spec_objectFormCallback.probe.stageOfInterest)
	}
	spec_object_ := spec_objectFormCallback.spec_object
	_ = spec_object_

	for _, formDiv := range spec_objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_object_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_.Unstage(spec_objectFormCallback.probe.stageOfInterest)
	}

	spec_objectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_OBJECT](
		spec_objectFormCallback.probe,
	)
	spec_objectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_objectFormCallback.CreationMode || spec_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECTFormCallback(
			nil,
			spec_objectFormCallback.probe,
			newFormGroup,
		)
		spec_object := new(models.SPEC_OBJECT)
		FillUpForm(spec_object, newFormGroup, spec_objectFormCallback.probe)
		spec_objectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_objectFormCallback.probe)
}
func __gong__New__SPEC_OBJECT_TYPEFormCallback(
	spec_object_type *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) {
	spec_object_typeFormCallback = new(SPEC_OBJECT_TYPEFormCallback)
	spec_object_typeFormCallback.probe = probe
	spec_object_typeFormCallback.spec_object_type = spec_object_type
	spec_object_typeFormCallback.formGroup = formGroup

	spec_object_typeFormCallback.CreationMode = (spec_object_type == nil)

	return
}

type SPEC_OBJECT_TYPEFormCallback struct {
	spec_object_type *models.SPEC_OBJECT_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_object_typeFormCallback *SPEC_OBJECT_TYPEFormCallback) OnSave() {

	log.Println("SPEC_OBJECT_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_object_typeFormCallback.probe.formStage.Checkout()

	if spec_object_typeFormCallback.spec_object_type == nil {
		spec_object_typeFormCallback.spec_object_type = new(models.SPEC_OBJECT_TYPE).Stage(spec_object_typeFormCallback.probe.stageOfInterest)
	}
	spec_object_type_ := spec_object_typeFormCallback.spec_object_type
	_ = spec_object_type_

	for _, formDiv := range spec_object_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_object_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_object_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_object_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_object_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_object_type_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_type_.Unstage(spec_object_typeFormCallback.probe.stageOfInterest)
	}

	spec_object_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_OBJECT_TYPE](
		spec_object_typeFormCallback.probe,
	)
	spec_object_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_object_typeFormCallback.CreationMode || spec_object_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_object_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_object_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_OBJECT_TYPEFormCallback(
			nil,
			spec_object_typeFormCallback.probe,
			newFormGroup,
		)
		spec_object_type := new(models.SPEC_OBJECT_TYPE)
		FillUpForm(spec_object_type, newFormGroup, spec_object_typeFormCallback.probe)
		spec_object_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_object_typeFormCallback.probe)
}
func __gong__New__SPEC_RELATIONFormCallback(
	spec_relation *models.SPEC_RELATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relationFormCallback *SPEC_RELATIONFormCallback) {
	spec_relationFormCallback = new(SPEC_RELATIONFormCallback)
	spec_relationFormCallback.probe = probe
	spec_relationFormCallback.spec_relation = spec_relation
	spec_relationFormCallback.formGroup = formGroup

	spec_relationFormCallback.CreationMode = (spec_relation == nil)

	return
}

type SPEC_RELATIONFormCallback struct {
	spec_relation *models.SPEC_RELATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relationFormCallback *SPEC_RELATIONFormCallback) OnSave() {

	log.Println("SPEC_RELATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relationFormCallback.probe.formStage.Checkout()

	if spec_relationFormCallback.spec_relation == nil {
		spec_relationFormCallback.spec_relation = new(models.SPEC_RELATION).Stage(spec_relationFormCallback.probe.stageOfInterest)
	}
	spec_relation_ := spec_relationFormCallback.spec_relation
	_ = spec_relation_

	for _, formDiv := range spec_relationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_relation_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_.Unstage(spec_relationFormCallback.probe.stageOfInterest)
	}

	spec_relationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_RELATION](
		spec_relationFormCallback.probe,
	)
	spec_relationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_relationFormCallback.CreationMode || spec_relationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_relationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATIONFormCallback(
			nil,
			spec_relationFormCallback.probe,
			newFormGroup,
		)
		spec_relation := new(models.SPEC_RELATION)
		FillUpForm(spec_relation, newFormGroup, spec_relationFormCallback.probe)
		spec_relationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_relationFormCallback.probe)
}
func __gong__New__SPEC_RELATION_TYPEFormCallback(
	spec_relation_type *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *table.FormGroup,
) (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) {
	spec_relation_typeFormCallback = new(SPEC_RELATION_TYPEFormCallback)
	spec_relation_typeFormCallback.probe = probe
	spec_relation_typeFormCallback.spec_relation_type = spec_relation_type
	spec_relation_typeFormCallback.formGroup = formGroup

	spec_relation_typeFormCallback.CreationMode = (spec_relation_type == nil)

	return
}

type SPEC_RELATION_TYPEFormCallback struct {
	spec_relation_type *models.SPEC_RELATION_TYPE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (spec_relation_typeFormCallback *SPEC_RELATION_TYPEFormCallback) OnSave() {

	log.Println("SPEC_RELATION_TYPEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	spec_relation_typeFormCallback.probe.formStage.Checkout()

	if spec_relation_typeFormCallback.spec_relation_type == nil {
		spec_relation_typeFormCallback.spec_relation_type = new(models.SPEC_RELATION_TYPE).Stage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}
	spec_relation_type_ := spec_relation_typeFormCallback.spec_relation_type
	_ = spec_relation_type_

	for _, formDiv := range spec_relation_typeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(spec_relation_type_.Name), formDiv)
		case "DESC":
			FormDivBasicFieldToField(&(spec_relation_type_.DESC), formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&(spec_relation_type_.IDENTIFIER), formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&(spec_relation_type_.LAST_CHANGE), formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&(spec_relation_type_.LONG_NAME), formDiv)
		}
	}

	// manage the suppress operation
	if spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_type_.Unstage(spec_relation_typeFormCallback.probe.stageOfInterest)
	}

	spec_relation_typeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SPEC_RELATION_TYPE](
		spec_relation_typeFormCallback.probe,
	)
	spec_relation_typeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if spec_relation_typeFormCallback.CreationMode || spec_relation_typeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		spec_relation_typeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(spec_relation_typeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SPEC_RELATION_TYPEFormCallback(
			nil,
			spec_relation_typeFormCallback.probe,
			newFormGroup,
		)
		spec_relation_type := new(models.SPEC_RELATION_TYPE)
		FillUpForm(spec_relation_type, newFormGroup, spec_relation_typeFormCallback.probe)
		spec_relation_typeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(spec_relation_typeFormCallback.probe)
}
func __gong__New__XHTML_CONTENTFormCallback(
	xhtml_content *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) {
	xhtml_contentFormCallback = new(XHTML_CONTENTFormCallback)
	xhtml_contentFormCallback.probe = probe
	xhtml_contentFormCallback.xhtml_content = xhtml_content
	xhtml_contentFormCallback.formGroup = formGroup

	xhtml_contentFormCallback.CreationMode = (xhtml_content == nil)

	return
}

type XHTML_CONTENTFormCallback struct {
	xhtml_content *models.XHTML_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xhtml_contentFormCallback *XHTML_CONTENTFormCallback) OnSave() {

	log.Println("XHTML_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xhtml_contentFormCallback.probe.formStage.Checkout()

	if xhtml_contentFormCallback.xhtml_content == nil {
		xhtml_contentFormCallback.xhtml_content = new(models.XHTML_CONTENT).Stage(xhtml_contentFormCallback.probe.stageOfInterest)
	}
	xhtml_content_ := xhtml_contentFormCallback.xhtml_content
	_ = xhtml_content_

	for _, formDiv := range xhtml_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xhtml_content_.Name), formDiv)
		case "ATTRIBUTE_VALUE_XHTML:THE_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_XHTMLOwner *models.ATTRIBUTE_VALUE_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xhtml_contentFormCallback.probe.stageOfInterest,
				xhtml_contentFormCallback.probe.backRepoOfInterest,
				xhtml_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
					pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](xhtml_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_XHTMLOwner := _attribute_value_xhtml // we have a match
						if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
							if newATTRIBUTE_VALUE_XHTMLOwner != pastATTRIBUTE_VALUE_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
								pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, idx, idx+1)
								newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
							}
						} else {
							newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_VALUE, xhtml_content_)
						}
					}
				}
			}
		case "ATTRIBUTE_VALUE_XHTML:THE_ORIGINAL_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_XHTMLOwner *models.ATTRIBUTE_VALUE_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "THE_ORIGINAL_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xhtml_contentFormCallback.probe.stageOfInterest,
				xhtml_contentFormCallback.probe.backRepoOfInterest,
				xhtml_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
					pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](xhtml_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_XHTMLOwner := _attribute_value_xhtml // we have a match
						if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
							if newATTRIBUTE_VALUE_XHTMLOwner != pastATTRIBUTE_VALUE_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
								pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, idx, idx+1)
								newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
							}
						} else {
							newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE = append(newATTRIBUTE_VALUE_XHTMLOwner.THE_ORIGINAL_VALUE, xhtml_content_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_content_.Unstage(xhtml_contentFormCallback.probe.stageOfInterest)
	}

	xhtml_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XHTML_CONTENT](
		xhtml_contentFormCallback.probe,
	)
	xhtml_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xhtml_contentFormCallback.CreationMode || xhtml_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xhtml_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(xhtml_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XHTML_CONTENTFormCallback(
			nil,
			xhtml_contentFormCallback.probe,
			newFormGroup,
		)
		xhtml_content := new(models.XHTML_CONTENT)
		FillUpForm(xhtml_content, newFormGroup, xhtml_contentFormCallback.probe)
		xhtml_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xhtml_contentFormCallback.probe)
}
