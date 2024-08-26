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
		case "A_ALTERNATIVE_ID:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastA_ALTERNATIVE_IDOwner *models.A_ALTERNATIVE_ID
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_ALTERNATIVE_ID"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				alternative_idFormCallback.probe.stageOfInterest,
				alternative_idFormCallback.probe.backRepoOfInterest,
				alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_ALTERNATIVE_IDOwner = reverseFieldOwner.(*models.A_ALTERNATIVE_ID)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_ALTERNATIVE_IDOwner != nil {
					idx := slices.Index(pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, alternative_id_)
					pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID = slices.Delete(pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_alternative_id := range *models.GetGongstructInstancesSet[models.A_ALTERNATIVE_ID](alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_alternative_id.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_ALTERNATIVE_IDOwner := _a_alternative_id // we have a match
						if pastA_ALTERNATIVE_IDOwner != nil {
							if newA_ALTERNATIVE_IDOwner != pastA_ALTERNATIVE_IDOwner {
								idx := slices.Index(pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, alternative_id_)
								pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID = slices.Delete(pastA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, idx, idx+1)
								newA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID = append(newA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, alternative_id_)
							}
						} else {
							newA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID = append(newA_ALTERNATIVE_IDOwner.ALTERNATIVE_ID, alternative_id_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_booleanFormCallback.probe.stageOfInterest,
				attribute_definition_booleanFormCallback.probe.backRepoOfInterest,
				attribute_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_BOOLEAN, attribute_definition_boolean_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_dateFormCallback.probe.stageOfInterest,
				attribute_definition_dateFormCallback.probe.backRepoOfInterest,
				attribute_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_DATE, attribute_definition_date_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_enumerationFormCallback.probe.stageOfInterest,
				attribute_definition_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_ENUMERATION, attribute_definition_enumeration_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_integerFormCallback.probe.stageOfInterest,
				attribute_definition_integerFormCallback.probe.backRepoOfInterest,
				attribute_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_INTEGER, attribute_definition_integer_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_realFormCallback.probe.stageOfInterest,
				attribute_definition_realFormCallback.probe.backRepoOfInterest,
				attribute_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_REAL, attribute_definition_real_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_stringFormCallback.probe.stageOfInterest,
				attribute_definition_stringFormCallback.probe.backRepoOfInterest,
				attribute_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_STRING, attribute_definition_string_)
						}
					}
				}
			}
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
		case "A_SPEC_ATTRIBUTES:ATTRIBUTE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_ATTRIBUTESOwner *models.A_SPEC_ATTRIBUTES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_ATTRIBUTES"
			rf.Fieldname = "ATTRIBUTE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_definition_xhtmlFormCallback.probe.stageOfInterest,
				attribute_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_ATTRIBUTESOwner = reverseFieldOwner.(*models.A_SPEC_ATTRIBUTES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_ATTRIBUTESOwner != nil {
					idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
					pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_attributes := range *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](attribute_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_attributes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_ATTRIBUTESOwner := _a_spec_attributes // we have a match
						if pastA_SPEC_ATTRIBUTESOwner != nil {
							if newA_SPEC_ATTRIBUTESOwner != pastA_SPEC_ATTRIBUTESOwner {
								idx := slices.Index(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
								pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML = slices.Delete(pastA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, idx, idx+1)
								newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
							}
						} else {
							newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML = append(newA_SPEC_ATTRIBUTESOwner.ATTRIBUTE_DEFINITION_XHTML, attribute_definition_xhtml_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_BOOLEAN_1:ATTRIBUTE_VALUE_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner *models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_BOOLEAN_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_booleanFormCallback.probe.stageOfInterest,
				attribute_value_booleanFormCallback.probe.backRepoOfInterest,
				attribute_value_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
					pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_boolean_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1](attribute_value_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_boolean_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner := _renamed_attribute_value_boolean_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner != pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
								pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN = append(newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN = append(newRenamed_ATTRIBUTE_VALUE_BOOLEAN_1Owner.ATTRIBUTE_VALUE_BOOLEAN, attribute_value_boolean_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_DATE_1:ATTRIBUTE_VALUE_DATE":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner *models.Renamed_ATTRIBUTE_VALUE_DATE_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_DATE_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_dateFormCallback.probe.stageOfInterest,
				attribute_value_dateFormCallback.probe.backRepoOfInterest,
				attribute_value_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_DATE_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
					pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_date_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_DATE_1](attribute_value_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_date_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_DATE_1Owner := _renamed_attribute_value_date_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_DATE_1Owner != pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
								pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE = append(newRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE = append(newRenamed_ATTRIBUTE_VALUE_DATE_1Owner.ATTRIBUTE_VALUE_DATE, attribute_value_date_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_ENUMERATION_1:ATTRIBUTE_VALUE_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner *models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_ENUMERATION_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_enumerationFormCallback.probe.stageOfInterest,
				attribute_value_enumerationFormCallback.probe.backRepoOfInterest,
				attribute_value_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
					pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_enumeration_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1](attribute_value_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_enumeration_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner := _renamed_attribute_value_enumeration_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner != pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
								pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION = append(newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION = append(newRenamed_ATTRIBUTE_VALUE_ENUMERATION_1Owner.ATTRIBUTE_VALUE_ENUMERATION, attribute_value_enumeration_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_INTEGER_1:ATTRIBUTE_VALUE_INTEGER":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner *models.Renamed_ATTRIBUTE_VALUE_INTEGER_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_INTEGER_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_integerFormCallback.probe.stageOfInterest,
				attribute_value_integerFormCallback.probe.backRepoOfInterest,
				attribute_value_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_INTEGER_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
					pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_integer_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_INTEGER_1](attribute_value_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_integer_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner := _renamed_attribute_value_integer_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner != pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
								pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER = append(newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER = append(newRenamed_ATTRIBUTE_VALUE_INTEGER_1Owner.ATTRIBUTE_VALUE_INTEGER, attribute_value_integer_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_REAL_1:ATTRIBUTE_VALUE_REAL":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner *models.Renamed_ATTRIBUTE_VALUE_REAL_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_REAL_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_realFormCallback.probe.stageOfInterest,
				attribute_value_realFormCallback.probe.backRepoOfInterest,
				attribute_value_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_REAL_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
					pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_real_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_REAL_1](attribute_value_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_real_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_REAL_1Owner := _renamed_attribute_value_real_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_REAL_1Owner != pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
								pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL = append(newRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL = append(newRenamed_ATTRIBUTE_VALUE_REAL_1Owner.ATTRIBUTE_VALUE_REAL, attribute_value_real_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_STRING_1:ATTRIBUTE_VALUE_STRING":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner *models.Renamed_ATTRIBUTE_VALUE_STRING_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_STRING_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_stringFormCallback.probe.stageOfInterest,
				attribute_value_stringFormCallback.probe.backRepoOfInterest,
				attribute_value_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_STRING_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
					pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_string_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_STRING_1](attribute_value_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_string_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_STRING_1Owner := _renamed_attribute_value_string_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_STRING_1Owner != pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
								pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING = append(newRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING = append(newRenamed_ATTRIBUTE_VALUE_STRING_1Owner.ATTRIBUTE_VALUE_STRING, attribute_value_string_)
						}
					}
				}
			}
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
		case "A_VALUES_1:ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastA_VALUES_1Owner *models.A_VALUES_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_VALUES_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_VALUES_1Owner = reverseFieldOwner.(*models.A_VALUES_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_VALUES_1Owner != nil {
					idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_values_1 := range *models.GetGongstructInstancesSet[models.A_VALUES_1](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_values_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_VALUES_1Owner := _a_values_1 // we have a match
						if pastA_VALUES_1Owner != nil {
							if newA_VALUES_1Owner != pastA_VALUES_1Owner {
								idx := slices.Index(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML = append(newA_VALUES_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
		case "Renamed_ATTRIBUTE_VALUE_XHTML_1:ATTRIBUTE_VALUE_XHTML":
			// we need to retrieve the field owner before the change
			var pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner *models.Renamed_ATTRIBUTE_VALUE_XHTML_1
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Renamed_ATTRIBUTE_VALUE_XHTML_1"
			rf.Fieldname = "ATTRIBUTE_VALUE_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attribute_value_xhtmlFormCallback.probe.stageOfInterest,
				attribute_value_xhtmlFormCallback.probe.backRepoOfInterest,
				attribute_value_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner = reverseFieldOwner.(*models.Renamed_ATTRIBUTE_VALUE_XHTML_1)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner != nil {
					idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
					pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _renamed_attribute_value_xhtml_1 := range *models.GetGongstructInstancesSet[models.Renamed_ATTRIBUTE_VALUE_XHTML_1](attribute_value_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _renamed_attribute_value_xhtml_1.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner := _renamed_attribute_value_xhtml_1 // we have a match
						if pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner != nil {
							if newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner != pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner {
								idx := slices.Index(pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
								pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML = slices.Delete(pastRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, idx, idx+1)
								newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML = append(newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
							}
						} else {
							newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML = append(newRenamed_ATTRIBUTE_VALUE_XHTML_1Owner.ATTRIBUTE_VALUE_XHTML, attribute_value_xhtml_)
						}
					}
				}
			}
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
func __gong__New__A_ALTERNATIVE_IDFormCallback(
	a_alternative_id *models.A_ALTERNATIVE_ID,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_alternative_idFormCallback *A_ALTERNATIVE_IDFormCallback) {
	a_alternative_idFormCallback = new(A_ALTERNATIVE_IDFormCallback)
	a_alternative_idFormCallback.probe = probe
	a_alternative_idFormCallback.a_alternative_id = a_alternative_id
	a_alternative_idFormCallback.formGroup = formGroup

	a_alternative_idFormCallback.CreationMode = (a_alternative_id == nil)

	return
}

type A_ALTERNATIVE_IDFormCallback struct {
	a_alternative_id *models.A_ALTERNATIVE_ID

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_alternative_idFormCallback *A_ALTERNATIVE_IDFormCallback) OnSave() {

	log.Println("A_ALTERNATIVE_IDFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_alternative_idFormCallback.probe.formStage.Checkout()

	if a_alternative_idFormCallback.a_alternative_id == nil {
		a_alternative_idFormCallback.a_alternative_id = new(models.A_ALTERNATIVE_ID).Stage(a_alternative_idFormCallback.probe.stageOfInterest)
	}
	a_alternative_id_ := a_alternative_idFormCallback.a_alternative_id
	_ = a_alternative_id_

	for _, formDiv := range a_alternative_idFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_alternative_id_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_BOOLEANOwner *models.ATTRIBUTE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_BOOLEANOwner := _attribute_definition_boolean // we have a match
						if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
							if newATTRIBUTE_DEFINITION_BOOLEANOwner != pastATTRIBUTE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_DATE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_DATEOwner *models.ATTRIBUTE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_DATEOwner := _attribute_definition_date // we have a match
						if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
							if newATTRIBUTE_DEFINITION_DATEOwner != pastATTRIBUTE_DEFINITION_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_ENUMERATIONOwner *models.ATTRIBUTE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_ENUMERATIONOwner := _attribute_definition_enumeration // we have a match
						if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
							if newATTRIBUTE_DEFINITION_ENUMERATIONOwner != pastATTRIBUTE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_INTEGER:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_INTEGEROwner *models.ATTRIBUTE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_INTEGEROwner := _attribute_definition_integer // we have a match
						if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
							if newATTRIBUTE_DEFINITION_INTEGEROwner != pastATTRIBUTE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_REAL:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_REALOwner *models.ATTRIBUTE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_REALOwner := _attribute_definition_real // we have a match
						if pastATTRIBUTE_DEFINITION_REALOwner != nil {
							if newATTRIBUTE_DEFINITION_REALOwner != pastATTRIBUTE_DEFINITION_REALOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_STRING:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_STRINGOwner *models.ATTRIBUTE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_STRINGOwner := _attribute_definition_string // we have a match
						if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
							if newATTRIBUTE_DEFINITION_STRINGOwner != pastATTRIBUTE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ATTRIBUTE_DEFINITION_XHTML:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_XHTMLOwner *models.ATTRIBUTE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_XHTMLOwner := _attribute_definition_xhtml // we have a match
						if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
							if newATTRIBUTE_DEFINITION_XHTMLOwner != pastATTRIBUTE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, idx, idx+1)
								newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = append(newATTRIBUTE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_BOOLEAN:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_BOOLEANOwner *models.DATATYPE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_BOOLEAN"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_boolean := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_BOOLEAN](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_BOOLEANOwner := _datatype_definition_boolean // we have a match
						if pastDATATYPE_DEFINITION_BOOLEANOwner != nil {
							if newDATATYPE_DEFINITION_BOOLEANOwner != pastDATATYPE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_BOOLEANOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_DATE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_DATEOwner *models.DATATYPE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_DATE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_date := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_DATE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_DATEOwner := _datatype_definition_date // we have a match
						if pastDATATYPE_DEFINITION_DATEOwner != nil {
							if newDATATYPE_DEFINITION_DATEOwner != pastDATATYPE_DEFINITION_DATEOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_DATEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_ENUMERATION:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_ENUMERATIONOwner *models.DATATYPE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_enumeration := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_ENUMERATIONOwner := _datatype_definition_enumeration // we have a match
						if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
							if newDATATYPE_DEFINITION_ENUMERATIONOwner != pastDATATYPE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_INTEGER:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_INTEGEROwner *models.DATATYPE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_INTEGER"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_integer := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_INTEGER](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_INTEGEROwner := _datatype_definition_integer // we have a match
						if pastDATATYPE_DEFINITION_INTEGEROwner != nil {
							if newDATATYPE_DEFINITION_INTEGEROwner != pastDATATYPE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_INTEGEROwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_REAL:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_REALOwner *models.DATATYPE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_REAL"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_REALOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_real := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_REAL](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_REALOwner := _datatype_definition_real // we have a match
						if pastDATATYPE_DEFINITION_REALOwner != nil {
							if newDATATYPE_DEFINITION_REALOwner != pastDATATYPE_DEFINITION_REALOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_REALOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_STRING:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_STRINGOwner *models.DATATYPE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_STRING"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_string := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_STRING](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_STRINGOwner := _datatype_definition_string // we have a match
						if pastDATATYPE_DEFINITION_STRINGOwner != nil {
							if newDATATYPE_DEFINITION_STRINGOwner != pastDATATYPE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_STRINGOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "DATATYPE_DEFINITION_XHTML:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_XHTMLOwner *models.DATATYPE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_XHTML"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_xhtml := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_XHTMLOwner := _datatype_definition_xhtml // we have a match
						if pastDATATYPE_DEFINITION_XHTMLOwner != nil {
							if newDATATYPE_DEFINITION_XHTMLOwner != pastDATATYPE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = slices.Delete(pastDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, idx, idx+1)
								newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID = append(newDATATYPE_DEFINITION_XHTMLOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "ENUM_VALUE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastENUM_VALUEOwner *models.ENUM_VALUE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastENUM_VALUEOwner = reverseFieldOwner.(*models.ENUM_VALUE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastENUM_VALUEOwner != nil {
					idx := slices.Index(pastENUM_VALUEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastENUM_VALUEOwner.ALTERNATIVE_ID = slices.Delete(pastENUM_VALUEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _enum_value := range *models.GetGongstructInstancesSet[models.ENUM_VALUE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _enum_value.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newENUM_VALUEOwner := _enum_value // we have a match
						if pastENUM_VALUEOwner != nil {
							if newENUM_VALUEOwner != pastENUM_VALUEOwner {
								idx := slices.Index(pastENUM_VALUEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastENUM_VALUEOwner.ALTERNATIVE_ID = slices.Delete(pastENUM_VALUEOwner.ALTERNATIVE_ID, idx, idx+1)
								newENUM_VALUEOwner.ALTERNATIVE_ID = append(newENUM_VALUEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newENUM_VALUEOwner.ALTERNATIVE_ID = append(newENUM_VALUEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "RELATION_GROUP:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUPOwner *models.RELATION_GROUP
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUPOwner = reverseFieldOwner.(*models.RELATION_GROUP)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUPOwner != nil {
					idx := slices.Index(pastRELATION_GROUPOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastRELATION_GROUPOwner.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUPOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group := range *models.GetGongstructInstancesSet[models.RELATION_GROUP](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUPOwner := _relation_group // we have a match
						if pastRELATION_GROUPOwner != nil {
							if newRELATION_GROUPOwner != pastRELATION_GROUPOwner {
								idx := slices.Index(pastRELATION_GROUPOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastRELATION_GROUPOwner.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUPOwner.ALTERNATIVE_ID, idx, idx+1)
								newRELATION_GROUPOwner.ALTERNATIVE_ID = append(newRELATION_GROUPOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newRELATION_GROUPOwner.ALTERNATIVE_ID = append(newRELATION_GROUPOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "RELATION_GROUP_TYPE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID = append(newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID = append(newRELATION_GROUP_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPECIFICATION:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPECIFICATIONOwner.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATIONOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPECIFICATIONOwner.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATIONOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPECIFICATIONOwner.ALTERNATIVE_ID = append(newSPECIFICATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPECIFICATIONOwner.ALTERNATIVE_ID = append(newSPECIFICATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID = append(newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID = append(newSPECIFICATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPEC_HIERARCHY:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_HIERARCHYOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_HIERARCHYOwner.ALTERNATIVE_ID = append(newSPEC_HIERARCHYOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPEC_HIERARCHYOwner.ALTERNATIVE_ID = append(newSPEC_HIERARCHYOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPEC_OBJECT:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPEC_OBJECTOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECTOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPEC_OBJECTOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECTOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_OBJECTOwner.ALTERNATIVE_ID = append(newSPEC_OBJECTOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPEC_OBJECTOwner.ALTERNATIVE_ID = append(newSPEC_OBJECTOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID = append(newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID = append(newSPEC_OBJECT_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPEC_RELATION:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPEC_RELATIONOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATIONOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPEC_RELATIONOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATIONOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_RELATIONOwner.ALTERNATIVE_ID = append(newSPEC_RELATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPEC_RELATIONOwner.ALTERNATIVE_ID = append(newSPEC_RELATIONOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:ALTERNATIVE_ID":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "ALTERNATIVE_ID"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_alternative_idFormCallback.probe.stageOfInterest,
				a_alternative_idFormCallback.probe.backRepoOfInterest,
				a_alternative_id_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
					pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](a_alternative_idFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
								pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID = slices.Delete(pastSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID = append(newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID = append(newSPEC_RELATION_TYPEOwner.ALTERNATIVE_ID, a_alternative_id_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_alternative_id_.Unstage(a_alternative_idFormCallback.probe.stageOfInterest)
	}

	a_alternative_idFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_ALTERNATIVE_ID](
		a_alternative_idFormCallback.probe,
	)
	a_alternative_idFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_alternative_idFormCallback.CreationMode || a_alternative_idFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_alternative_idFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_alternative_idFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_ALTERNATIVE_IDFormCallback(
			nil,
			a_alternative_idFormCallback.probe,
			newFormGroup,
		)
		a_alternative_id := new(models.A_ALTERNATIVE_ID)
		FillUpForm(a_alternative_id, newFormGroup, a_alternative_idFormCallback.probe)
		a_alternative_idFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_alternative_idFormCallback.probe)
}
func __gong__New__A_CHILDRENFormCallback(
	a_children *models.A_CHILDREN,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_childrenFormCallback *A_CHILDRENFormCallback) {
	a_childrenFormCallback = new(A_CHILDRENFormCallback)
	a_childrenFormCallback.probe = probe
	a_childrenFormCallback.a_children = a_children
	a_childrenFormCallback.formGroup = formGroup

	a_childrenFormCallback.CreationMode = (a_children == nil)

	return
}

type A_CHILDRENFormCallback struct {
	a_children *models.A_CHILDREN

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_childrenFormCallback *A_CHILDRENFormCallback) OnSave() {

	log.Println("A_CHILDRENFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_childrenFormCallback.probe.formStage.Checkout()

	if a_childrenFormCallback.a_children == nil {
		a_childrenFormCallback.a_children = new(models.A_CHILDREN).Stage(a_childrenFormCallback.probe.stageOfInterest)
	}
	a_children_ := a_childrenFormCallback.a_children
	_ = a_children_

	for _, formDiv := range a_childrenFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_children_.Name), formDiv)
		case "SPECIFICATION:CHILDREN":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "CHILDREN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_childrenFormCallback.probe.stageOfInterest,
				a_childrenFormCallback.probe.backRepoOfInterest,
				a_children_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.CHILDREN, a_children_)
					pastSPECIFICATIONOwner.CHILDREN = slices.Delete(pastSPECIFICATIONOwner.CHILDREN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](a_childrenFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.CHILDREN, a_children_)
								pastSPECIFICATIONOwner.CHILDREN = slices.Delete(pastSPECIFICATIONOwner.CHILDREN, idx, idx+1)
								newSPECIFICATIONOwner.CHILDREN = append(newSPECIFICATIONOwner.CHILDREN, a_children_)
							}
						} else {
							newSPECIFICATIONOwner.CHILDREN = append(newSPECIFICATIONOwner.CHILDREN, a_children_)
						}
					}
				}
			}
		case "SPEC_HIERARCHY:CHILDREN":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "CHILDREN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_childrenFormCallback.probe.stageOfInterest,
				a_childrenFormCallback.probe.backRepoOfInterest,
				a_children_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.CHILDREN, a_children_)
					pastSPEC_HIERARCHYOwner.CHILDREN = slices.Delete(pastSPEC_HIERARCHYOwner.CHILDREN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](a_childrenFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.CHILDREN, a_children_)
								pastSPEC_HIERARCHYOwner.CHILDREN = slices.Delete(pastSPEC_HIERARCHYOwner.CHILDREN, idx, idx+1)
								newSPEC_HIERARCHYOwner.CHILDREN = append(newSPEC_HIERARCHYOwner.CHILDREN, a_children_)
							}
						} else {
							newSPEC_HIERARCHYOwner.CHILDREN = append(newSPEC_HIERARCHYOwner.CHILDREN, a_children_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_children_.Unstage(a_childrenFormCallback.probe.stageOfInterest)
	}

	a_childrenFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_CHILDREN](
		a_childrenFormCallback.probe,
	)
	a_childrenFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_childrenFormCallback.CreationMode || a_childrenFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_childrenFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_childrenFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_CHILDRENFormCallback(
			nil,
			a_childrenFormCallback.probe,
			newFormGroup,
		)
		a_children := new(models.A_CHILDREN)
		FillUpForm(a_children, newFormGroup, a_childrenFormCallback.probe)
		a_childrenFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_childrenFormCallback.probe)
}
func __gong__New__A_CORE_CONTENTFormCallback(
	a_core_content *models.A_CORE_CONTENT,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_core_contentFormCallback *A_CORE_CONTENTFormCallback) {
	a_core_contentFormCallback = new(A_CORE_CONTENTFormCallback)
	a_core_contentFormCallback.probe = probe
	a_core_contentFormCallback.a_core_content = a_core_content
	a_core_contentFormCallback.formGroup = formGroup

	a_core_contentFormCallback.CreationMode = (a_core_content == nil)

	return
}

type A_CORE_CONTENTFormCallback struct {
	a_core_content *models.A_CORE_CONTENT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_core_contentFormCallback *A_CORE_CONTENTFormCallback) OnSave() {

	log.Println("A_CORE_CONTENTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_core_contentFormCallback.probe.formStage.Checkout()

	if a_core_contentFormCallback.a_core_content == nil {
		a_core_contentFormCallback.a_core_content = new(models.A_CORE_CONTENT).Stage(a_core_contentFormCallback.probe.stageOfInterest)
	}
	a_core_content_ := a_core_contentFormCallback.a_core_content
	_ = a_core_content_

	for _, formDiv := range a_core_contentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_core_content_.Name), formDiv)
		case "REQ_IF:CORE_CONTENT":
			// we need to retrieve the field owner before the change
			var pastREQ_IFOwner *models.REQ_IF
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "CORE_CONTENT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_core_contentFormCallback.probe.stageOfInterest,
				a_core_contentFormCallback.probe.backRepoOfInterest,
				a_core_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IFOwner = reverseFieldOwner.(*models.REQ_IF)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IFOwner != nil {
					idx := slices.Index(pastREQ_IFOwner.CORE_CONTENT, a_core_content_)
					pastREQ_IFOwner.CORE_CONTENT = slices.Delete(pastREQ_IFOwner.CORE_CONTENT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if := range *models.GetGongstructInstancesSet[models.REQ_IF](a_core_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IFOwner := _req_if // we have a match
						if pastREQ_IFOwner != nil {
							if newREQ_IFOwner != pastREQ_IFOwner {
								idx := slices.Index(pastREQ_IFOwner.CORE_CONTENT, a_core_content_)
								pastREQ_IFOwner.CORE_CONTENT = slices.Delete(pastREQ_IFOwner.CORE_CONTENT, idx, idx+1)
								newREQ_IFOwner.CORE_CONTENT = append(newREQ_IFOwner.CORE_CONTENT, a_core_content_)
							}
						} else {
							newREQ_IFOwner.CORE_CONTENT = append(newREQ_IFOwner.CORE_CONTENT, a_core_content_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_core_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_core_content_.Unstage(a_core_contentFormCallback.probe.stageOfInterest)
	}

	a_core_contentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_CORE_CONTENT](
		a_core_contentFormCallback.probe,
	)
	a_core_contentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_core_contentFormCallback.CreationMode || a_core_contentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_core_contentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_core_contentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_CORE_CONTENTFormCallback(
			nil,
			a_core_contentFormCallback.probe,
			newFormGroup,
		)
		a_core_content := new(models.A_CORE_CONTENT)
		FillUpForm(a_core_content, newFormGroup, a_core_contentFormCallback.probe)
		a_core_contentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_core_contentFormCallback.probe)
}
func __gong__New__A_DATATYPESFormCallback(
	a_datatypes *models.A_DATATYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_datatypesFormCallback *A_DATATYPESFormCallback) {
	a_datatypesFormCallback = new(A_DATATYPESFormCallback)
	a_datatypesFormCallback.probe = probe
	a_datatypesFormCallback.a_datatypes = a_datatypes
	a_datatypesFormCallback.formGroup = formGroup

	a_datatypesFormCallback.CreationMode = (a_datatypes == nil)

	return
}

type A_DATATYPESFormCallback struct {
	a_datatypes *models.A_DATATYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_datatypesFormCallback *A_DATATYPESFormCallback) OnSave() {

	log.Println("A_DATATYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_datatypesFormCallback.probe.formStage.Checkout()

	if a_datatypesFormCallback.a_datatypes == nil {
		a_datatypesFormCallback.a_datatypes = new(models.A_DATATYPES).Stage(a_datatypesFormCallback.probe.stageOfInterest)
	}
	a_datatypes_ := a_datatypesFormCallback.a_datatypes
	_ = a_datatypes_

	for _, formDiv := range a_datatypesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_datatypes_.Name), formDiv)
		case "REQ_IF_CONTENT:DATATYPES":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "DATATYPES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_datatypesFormCallback.probe.stageOfInterest,
				a_datatypesFormCallback.probe.backRepoOfInterest,
				a_datatypes_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES, a_datatypes_)
					pastREQ_IF_CONTENTOwner.DATATYPES = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_datatypesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.DATATYPES, a_datatypes_)
								pastREQ_IF_CONTENTOwner.DATATYPES = slices.Delete(pastREQ_IF_CONTENTOwner.DATATYPES, idx, idx+1)
								newREQ_IF_CONTENTOwner.DATATYPES = append(newREQ_IF_CONTENTOwner.DATATYPES, a_datatypes_)
							}
						} else {
							newREQ_IF_CONTENTOwner.DATATYPES = append(newREQ_IF_CONTENTOwner.DATATYPES, a_datatypes_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatypes_.Unstage(a_datatypesFormCallback.probe.stageOfInterest)
	}

	a_datatypesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_DATATYPES](
		a_datatypesFormCallback.probe,
	)
	a_datatypesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_datatypesFormCallback.CreationMode || a_datatypesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_datatypesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_datatypesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_DATATYPESFormCallback(
			nil,
			a_datatypesFormCallback.probe,
			newFormGroup,
		)
		a_datatypes := new(models.A_DATATYPES)
		FillUpForm(a_datatypes, newFormGroup, a_datatypesFormCallback.probe)
		a_datatypesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_datatypesFormCallback.probe)
}
func __gong__New__A_EDITABLE_ATTSFormCallback(
	a_editable_atts *models.A_EDITABLE_ATTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_editable_attsFormCallback *A_EDITABLE_ATTSFormCallback) {
	a_editable_attsFormCallback = new(A_EDITABLE_ATTSFormCallback)
	a_editable_attsFormCallback.probe = probe
	a_editable_attsFormCallback.a_editable_atts = a_editable_atts
	a_editable_attsFormCallback.formGroup = formGroup

	a_editable_attsFormCallback.CreationMode = (a_editable_atts == nil)

	return
}

type A_EDITABLE_ATTSFormCallback struct {
	a_editable_atts *models.A_EDITABLE_ATTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_editable_attsFormCallback *A_EDITABLE_ATTSFormCallback) OnSave() {

	log.Println("A_EDITABLE_ATTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_editable_attsFormCallback.probe.formStage.Checkout()

	if a_editable_attsFormCallback.a_editable_atts == nil {
		a_editable_attsFormCallback.a_editable_atts = new(models.A_EDITABLE_ATTS).Stage(a_editable_attsFormCallback.probe.stageOfInterest)
	}
	a_editable_atts_ := a_editable_attsFormCallback.a_editable_atts
	_ = a_editable_atts_

	for _, formDiv := range a_editable_attsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_editable_atts_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(a_editable_atts_.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		case "SPEC_HIERARCHY:EDITABLE_ATTS":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "EDITABLE_ATTS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_editable_attsFormCallback.probe.stageOfInterest,
				a_editable_attsFormCallback.probe.backRepoOfInterest,
				a_editable_atts_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.EDITABLE_ATTS, a_editable_atts_)
					pastSPEC_HIERARCHYOwner.EDITABLE_ATTS = slices.Delete(pastSPEC_HIERARCHYOwner.EDITABLE_ATTS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](a_editable_attsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.EDITABLE_ATTS, a_editable_atts_)
								pastSPEC_HIERARCHYOwner.EDITABLE_ATTS = slices.Delete(pastSPEC_HIERARCHYOwner.EDITABLE_ATTS, idx, idx+1)
								newSPEC_HIERARCHYOwner.EDITABLE_ATTS = append(newSPEC_HIERARCHYOwner.EDITABLE_ATTS, a_editable_atts_)
							}
						} else {
							newSPEC_HIERARCHYOwner.EDITABLE_ATTS = append(newSPEC_HIERARCHYOwner.EDITABLE_ATTS, a_editable_atts_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_editable_attsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_editable_atts_.Unstage(a_editable_attsFormCallback.probe.stageOfInterest)
	}

	a_editable_attsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_EDITABLE_ATTS](
		a_editable_attsFormCallback.probe,
	)
	a_editable_attsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_editable_attsFormCallback.CreationMode || a_editable_attsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_editable_attsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_editable_attsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_EDITABLE_ATTSFormCallback(
			nil,
			a_editable_attsFormCallback.probe,
			newFormGroup,
		)
		a_editable_atts := new(models.A_EDITABLE_ATTS)
		FillUpForm(a_editable_atts, newFormGroup, a_editable_attsFormCallback.probe)
		a_editable_attsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_editable_attsFormCallback.probe)
}
func __gong__New__A_OBJECTFormCallback(
	a_object *models.A_OBJECT,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_objectFormCallback *A_OBJECTFormCallback) {
	a_objectFormCallback = new(A_OBJECTFormCallback)
	a_objectFormCallback.probe = probe
	a_objectFormCallback.a_object = a_object
	a_objectFormCallback.formGroup = formGroup

	a_objectFormCallback.CreationMode = (a_object == nil)

	return
}

type A_OBJECTFormCallback struct {
	a_object *models.A_OBJECT

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_objectFormCallback *A_OBJECTFormCallback) OnSave() {

	log.Println("A_OBJECTFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_objectFormCallback.probe.formStage.Checkout()

	if a_objectFormCallback.a_object == nil {
		a_objectFormCallback.a_object = new(models.A_OBJECT).Stage(a_objectFormCallback.probe.stageOfInterest)
	}
	a_object_ := a_objectFormCallback.a_object
	_ = a_object_

	for _, formDiv := range a_objectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_object_.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(a_object_.SPEC_OBJECT_REF), formDiv)
		case "SPEC_HIERARCHY:OBJECT":
			// we need to retrieve the field owner before the change
			var pastSPEC_HIERARCHYOwner *models.SPEC_HIERARCHY
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_HIERARCHY"
			rf.Fieldname = "OBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_objectFormCallback.probe.stageOfInterest,
				a_objectFormCallback.probe.backRepoOfInterest,
				a_object_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_HIERARCHYOwner = reverseFieldOwner.(*models.SPEC_HIERARCHY)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_HIERARCHYOwner != nil {
					idx := slices.Index(pastSPEC_HIERARCHYOwner.OBJECT, a_object_)
					pastSPEC_HIERARCHYOwner.OBJECT = slices.Delete(pastSPEC_HIERARCHYOwner.OBJECT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_hierarchy := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](a_objectFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_hierarchy.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_HIERARCHYOwner := _spec_hierarchy // we have a match
						if pastSPEC_HIERARCHYOwner != nil {
							if newSPEC_HIERARCHYOwner != pastSPEC_HIERARCHYOwner {
								idx := slices.Index(pastSPEC_HIERARCHYOwner.OBJECT, a_object_)
								pastSPEC_HIERARCHYOwner.OBJECT = slices.Delete(pastSPEC_HIERARCHYOwner.OBJECT, idx, idx+1)
								newSPEC_HIERARCHYOwner.OBJECT = append(newSPEC_HIERARCHYOwner.OBJECT, a_object_)
							}
						} else {
							newSPEC_HIERARCHYOwner.OBJECT = append(newSPEC_HIERARCHYOwner.OBJECT, a_object_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_object_.Unstage(a_objectFormCallback.probe.stageOfInterest)
	}

	a_objectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_OBJECT](
		a_objectFormCallback.probe,
	)
	a_objectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_objectFormCallback.CreationMode || a_objectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_objectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_objectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_OBJECTFormCallback(
			nil,
			a_objectFormCallback.probe,
			newFormGroup,
		)
		a_object := new(models.A_OBJECT)
		FillUpForm(a_object, newFormGroup, a_objectFormCallback.probe)
		a_objectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_objectFormCallback.probe)
}
func __gong__New__A_PROPERTIESFormCallback(
	a_properties *models.A_PROPERTIES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_propertiesFormCallback *A_PROPERTIESFormCallback) {
	a_propertiesFormCallback = new(A_PROPERTIESFormCallback)
	a_propertiesFormCallback.probe = probe
	a_propertiesFormCallback.a_properties = a_properties
	a_propertiesFormCallback.formGroup = formGroup

	a_propertiesFormCallback.CreationMode = (a_properties == nil)

	return
}

type A_PROPERTIESFormCallback struct {
	a_properties *models.A_PROPERTIES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_propertiesFormCallback *A_PROPERTIESFormCallback) OnSave() {

	log.Println("A_PROPERTIESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_propertiesFormCallback.probe.formStage.Checkout()

	if a_propertiesFormCallback.a_properties == nil {
		a_propertiesFormCallback.a_properties = new(models.A_PROPERTIES).Stage(a_propertiesFormCallback.probe.stageOfInterest)
	}
	a_properties_ := a_propertiesFormCallback.a_properties
	_ = a_properties_

	for _, formDiv := range a_propertiesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_properties_.Name), formDiv)
		case "ENUM_VALUE:PROPERTIES":
			// we need to retrieve the field owner before the change
			var pastENUM_VALUEOwner *models.ENUM_VALUE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ENUM_VALUE"
			rf.Fieldname = "PROPERTIES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_propertiesFormCallback.probe.stageOfInterest,
				a_propertiesFormCallback.probe.backRepoOfInterest,
				a_properties_,
				&rf)

			if reverseFieldOwner != nil {
				pastENUM_VALUEOwner = reverseFieldOwner.(*models.ENUM_VALUE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastENUM_VALUEOwner != nil {
					idx := slices.Index(pastENUM_VALUEOwner.PROPERTIES, a_properties_)
					pastENUM_VALUEOwner.PROPERTIES = slices.Delete(pastENUM_VALUEOwner.PROPERTIES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _enum_value := range *models.GetGongstructInstancesSet[models.ENUM_VALUE](a_propertiesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _enum_value.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newENUM_VALUEOwner := _enum_value // we have a match
						if pastENUM_VALUEOwner != nil {
							if newENUM_VALUEOwner != pastENUM_VALUEOwner {
								idx := slices.Index(pastENUM_VALUEOwner.PROPERTIES, a_properties_)
								pastENUM_VALUEOwner.PROPERTIES = slices.Delete(pastENUM_VALUEOwner.PROPERTIES, idx, idx+1)
								newENUM_VALUEOwner.PROPERTIES = append(newENUM_VALUEOwner.PROPERTIES, a_properties_)
							}
						} else {
							newENUM_VALUEOwner.PROPERTIES = append(newENUM_VALUEOwner.PROPERTIES, a_properties_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_properties_.Unstage(a_propertiesFormCallback.probe.stageOfInterest)
	}

	a_propertiesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_PROPERTIES](
		a_propertiesFormCallback.probe,
	)
	a_propertiesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_propertiesFormCallback.CreationMode || a_propertiesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_propertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_propertiesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_PROPERTIESFormCallback(
			nil,
			a_propertiesFormCallback.probe,
			newFormGroup,
		)
		a_properties := new(models.A_PROPERTIES)
		FillUpForm(a_properties, newFormGroup, a_propertiesFormCallback.probe)
		a_propertiesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_propertiesFormCallback.probe)
}
func __gong__New__A_SOURCEFormCallback(
	a_source *models.A_SOURCE,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_sourceFormCallback *A_SOURCEFormCallback) {
	a_sourceFormCallback = new(A_SOURCEFormCallback)
	a_sourceFormCallback.probe = probe
	a_sourceFormCallback.a_source = a_source
	a_sourceFormCallback.formGroup = formGroup

	a_sourceFormCallback.CreationMode = (a_source == nil)

	return
}

type A_SOURCEFormCallback struct {
	a_source *models.A_SOURCE

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_sourceFormCallback *A_SOURCEFormCallback) OnSave() {

	log.Println("A_SOURCEFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_sourceFormCallback.probe.formStage.Checkout()

	if a_sourceFormCallback.a_source == nil {
		a_sourceFormCallback.a_source = new(models.A_SOURCE).Stage(a_sourceFormCallback.probe.stageOfInterest)
	}
	a_source_ := a_sourceFormCallback.a_source
	_ = a_source_

	for _, formDiv := range a_sourceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_source_.Name), formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&(a_source_.SPEC_OBJECT_REF), formDiv)
		case "SPEC_RELATION:SOURCE":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "SOURCE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_sourceFormCallback.probe.stageOfInterest,
				a_sourceFormCallback.probe.backRepoOfInterest,
				a_source_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.SOURCE, a_source_)
					pastSPEC_RELATIONOwner.SOURCE = slices.Delete(pastSPEC_RELATIONOwner.SOURCE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](a_sourceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.SOURCE, a_source_)
								pastSPEC_RELATIONOwner.SOURCE = slices.Delete(pastSPEC_RELATIONOwner.SOURCE, idx, idx+1)
								newSPEC_RELATIONOwner.SOURCE = append(newSPEC_RELATIONOwner.SOURCE, a_source_)
							}
						} else {
							newSPEC_RELATIONOwner.SOURCE = append(newSPEC_RELATIONOwner.SOURCE, a_source_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_sourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_.Unstage(a_sourceFormCallback.probe.stageOfInterest)
	}

	a_sourceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SOURCE](
		a_sourceFormCallback.probe,
	)
	a_sourceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_sourceFormCallback.CreationMode || a_sourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_sourceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_sourceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SOURCEFormCallback(
			nil,
			a_sourceFormCallback.probe,
			newFormGroup,
		)
		a_source := new(models.A_SOURCE)
		FillUpForm(a_source, newFormGroup, a_sourceFormCallback.probe)
		a_sourceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_sourceFormCallback.probe)
}
func __gong__New__A_SOURCE_SPECIFICATIONFormCallback(
	a_source_specification *models.A_SOURCE_SPECIFICATION,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_source_specificationFormCallback *A_SOURCE_SPECIFICATIONFormCallback) {
	a_source_specificationFormCallback = new(A_SOURCE_SPECIFICATIONFormCallback)
	a_source_specificationFormCallback.probe = probe
	a_source_specificationFormCallback.a_source_specification = a_source_specification
	a_source_specificationFormCallback.formGroup = formGroup

	a_source_specificationFormCallback.CreationMode = (a_source_specification == nil)

	return
}

type A_SOURCE_SPECIFICATIONFormCallback struct {
	a_source_specification *models.A_SOURCE_SPECIFICATION

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_source_specificationFormCallback *A_SOURCE_SPECIFICATIONFormCallback) OnSave() {

	log.Println("A_SOURCE_SPECIFICATIONFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_source_specificationFormCallback.probe.formStage.Checkout()

	if a_source_specificationFormCallback.a_source_specification == nil {
		a_source_specificationFormCallback.a_source_specification = new(models.A_SOURCE_SPECIFICATION).Stage(a_source_specificationFormCallback.probe.stageOfInterest)
	}
	a_source_specification_ := a_source_specificationFormCallback.a_source_specification
	_ = a_source_specification_

	for _, formDiv := range a_source_specificationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_source_specification_.Name), formDiv)
		case "SPECIFICATION_REF":
			FormDivBasicFieldToField(&(a_source_specification_.SPECIFICATION_REF), formDiv)
		case "RELATION_GROUP:SOURCE_SPECIFICATION":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUPOwner *models.RELATION_GROUP
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "SOURCE_SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_source_specificationFormCallback.probe.stageOfInterest,
				a_source_specificationFormCallback.probe.backRepoOfInterest,
				a_source_specification_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUPOwner = reverseFieldOwner.(*models.RELATION_GROUP)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUPOwner != nil {
					idx := slices.Index(pastRELATION_GROUPOwner.SOURCE_SPECIFICATION, a_source_specification_)
					pastRELATION_GROUPOwner.SOURCE_SPECIFICATION = slices.Delete(pastRELATION_GROUPOwner.SOURCE_SPECIFICATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group := range *models.GetGongstructInstancesSet[models.RELATION_GROUP](a_source_specificationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUPOwner := _relation_group // we have a match
						if pastRELATION_GROUPOwner != nil {
							if newRELATION_GROUPOwner != pastRELATION_GROUPOwner {
								idx := slices.Index(pastRELATION_GROUPOwner.SOURCE_SPECIFICATION, a_source_specification_)
								pastRELATION_GROUPOwner.SOURCE_SPECIFICATION = slices.Delete(pastRELATION_GROUPOwner.SOURCE_SPECIFICATION, idx, idx+1)
								newRELATION_GROUPOwner.SOURCE_SPECIFICATION = append(newRELATION_GROUPOwner.SOURCE_SPECIFICATION, a_source_specification_)
							}
						} else {
							newRELATION_GROUPOwner.SOURCE_SPECIFICATION = append(newRELATION_GROUPOwner.SOURCE_SPECIFICATION, a_source_specification_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_source_specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_specification_.Unstage(a_source_specificationFormCallback.probe.stageOfInterest)
	}

	a_source_specificationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SOURCE_SPECIFICATION](
		a_source_specificationFormCallback.probe,
	)
	a_source_specificationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_source_specificationFormCallback.CreationMode || a_source_specificationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_source_specificationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_source_specificationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SOURCE_SPECIFICATIONFormCallback(
			nil,
			a_source_specificationFormCallback.probe,
			newFormGroup,
		)
		a_source_specification := new(models.A_SOURCE_SPECIFICATION)
		FillUpForm(a_source_specification, newFormGroup, a_source_specificationFormCallback.probe)
		a_source_specificationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_source_specificationFormCallback.probe)
}
func __gong__New__A_SPECIFICATIONSFormCallback(
	a_specifications *models.A_SPECIFICATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_specificationsFormCallback *A_SPECIFICATIONSFormCallback) {
	a_specificationsFormCallback = new(A_SPECIFICATIONSFormCallback)
	a_specificationsFormCallback.probe = probe
	a_specificationsFormCallback.a_specifications = a_specifications
	a_specificationsFormCallback.formGroup = formGroup

	a_specificationsFormCallback.CreationMode = (a_specifications == nil)

	return
}

type A_SPECIFICATIONSFormCallback struct {
	a_specifications *models.A_SPECIFICATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_specificationsFormCallback *A_SPECIFICATIONSFormCallback) OnSave() {

	log.Println("A_SPECIFICATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_specificationsFormCallback.probe.formStage.Checkout()

	if a_specificationsFormCallback.a_specifications == nil {
		a_specificationsFormCallback.a_specifications = new(models.A_SPECIFICATIONS).Stage(a_specificationsFormCallback.probe.stageOfInterest)
	}
	a_specifications_ := a_specificationsFormCallback.a_specifications
	_ = a_specifications_

	for _, formDiv := range a_specificationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_specifications_.Name), formDiv)
		case "REQ_IF_CONTENT:SPECIFICATIONS":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPECIFICATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_specificationsFormCallback.probe.stageOfInterest,
				a_specificationsFormCallback.probe.backRepoOfInterest,
				a_specifications_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPECIFICATIONS, a_specifications_)
					pastREQ_IF_CONTENTOwner.SPECIFICATIONS = slices.Delete(pastREQ_IF_CONTENTOwner.SPECIFICATIONS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_specificationsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPECIFICATIONS, a_specifications_)
								pastREQ_IF_CONTENTOwner.SPECIFICATIONS = slices.Delete(pastREQ_IF_CONTENTOwner.SPECIFICATIONS, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPECIFICATIONS = append(newREQ_IF_CONTENTOwner.SPECIFICATIONS, a_specifications_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPECIFICATIONS = append(newREQ_IF_CONTENTOwner.SPECIFICATIONS, a_specifications_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specifications_.Unstage(a_specificationsFormCallback.probe.stageOfInterest)
	}

	a_specificationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPECIFICATIONS](
		a_specificationsFormCallback.probe,
	)
	a_specificationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_specificationsFormCallback.CreationMode || a_specificationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specificationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_specificationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPECIFICATIONSFormCallback(
			nil,
			a_specificationsFormCallback.probe,
			newFormGroup,
		)
		a_specifications := new(models.A_SPECIFICATIONS)
		FillUpForm(a_specifications, newFormGroup, a_specificationsFormCallback.probe)
		a_specificationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_specificationsFormCallback.probe)
}
func __gong__New__A_SPECIFIED_VALUESFormCallback(
	a_specified_values *models.A_SPECIFIED_VALUES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_specified_valuesFormCallback *A_SPECIFIED_VALUESFormCallback) {
	a_specified_valuesFormCallback = new(A_SPECIFIED_VALUESFormCallback)
	a_specified_valuesFormCallback.probe = probe
	a_specified_valuesFormCallback.a_specified_values = a_specified_values
	a_specified_valuesFormCallback.formGroup = formGroup

	a_specified_valuesFormCallback.CreationMode = (a_specified_values == nil)

	return
}

type A_SPECIFIED_VALUESFormCallback struct {
	a_specified_values *models.A_SPECIFIED_VALUES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_specified_valuesFormCallback *A_SPECIFIED_VALUESFormCallback) OnSave() {

	log.Println("A_SPECIFIED_VALUESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_specified_valuesFormCallback.probe.formStage.Checkout()

	if a_specified_valuesFormCallback.a_specified_values == nil {
		a_specified_valuesFormCallback.a_specified_values = new(models.A_SPECIFIED_VALUES).Stage(a_specified_valuesFormCallback.probe.stageOfInterest)
	}
	a_specified_values_ := a_specified_valuesFormCallback.a_specified_values
	_ = a_specified_values_

	for _, formDiv := range a_specified_valuesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_specified_values_.Name), formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION:SPECIFIED_VALUES":
			// we need to retrieve the field owner before the change
			var pastDATATYPE_DEFINITION_ENUMERATIONOwner *models.DATATYPE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DATATYPE_DEFINITION_ENUMERATION"
			rf.Fieldname = "SPECIFIED_VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_specified_valuesFormCallback.probe.stageOfInterest,
				a_specified_valuesFormCallback.probe.backRepoOfInterest,
				a_specified_values_,
				&rf)

			if reverseFieldOwner != nil {
				pastDATATYPE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.DATATYPE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, a_specified_values_)
					pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _datatype_definition_enumeration := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](a_specified_valuesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _datatype_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDATATYPE_DEFINITION_ENUMERATIONOwner := _datatype_definition_enumeration // we have a match
						if pastDATATYPE_DEFINITION_ENUMERATIONOwner != nil {
							if newDATATYPE_DEFINITION_ENUMERATIONOwner != pastDATATYPE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, a_specified_values_)
								pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES = slices.Delete(pastDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, idx, idx+1)
								newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, a_specified_values_)
							}
						} else {
							newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES = append(newDATATYPE_DEFINITION_ENUMERATIONOwner.SPECIFIED_VALUES, a_specified_values_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_specified_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specified_values_.Unstage(a_specified_valuesFormCallback.probe.stageOfInterest)
	}

	a_specified_valuesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPECIFIED_VALUES](
		a_specified_valuesFormCallback.probe,
	)
	a_specified_valuesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_specified_valuesFormCallback.CreationMode || a_specified_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_specified_valuesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_specified_valuesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPECIFIED_VALUESFormCallback(
			nil,
			a_specified_valuesFormCallback.probe,
			newFormGroup,
		)
		a_specified_values := new(models.A_SPECIFIED_VALUES)
		FillUpForm(a_specified_values, newFormGroup, a_specified_valuesFormCallback.probe)
		a_specified_valuesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_specified_valuesFormCallback.probe)
}
func __gong__New__A_SPEC_ATTRIBUTESFormCallback(
	a_spec_attributes *models.A_SPEC_ATTRIBUTES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_attributesFormCallback *A_SPEC_ATTRIBUTESFormCallback) {
	a_spec_attributesFormCallback = new(A_SPEC_ATTRIBUTESFormCallback)
	a_spec_attributesFormCallback.probe = probe
	a_spec_attributesFormCallback.a_spec_attributes = a_spec_attributes
	a_spec_attributesFormCallback.formGroup = formGroup

	a_spec_attributesFormCallback.CreationMode = (a_spec_attributes == nil)

	return
}

type A_SPEC_ATTRIBUTESFormCallback struct {
	a_spec_attributes *models.A_SPEC_ATTRIBUTES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_attributesFormCallback *A_SPEC_ATTRIBUTESFormCallback) OnSave() {

	log.Println("A_SPEC_ATTRIBUTESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_attributesFormCallback.probe.formStage.Checkout()

	if a_spec_attributesFormCallback.a_spec_attributes == nil {
		a_spec_attributesFormCallback.a_spec_attributes = new(models.A_SPEC_ATTRIBUTES).Stage(a_spec_attributesFormCallback.probe.stageOfInterest)
	}
	a_spec_attributes_ := a_spec_attributesFormCallback.a_spec_attributes
	_ = a_spec_attributes_

	for _, formDiv := range a_spec_attributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_attributes_.Name), formDiv)
		case "RELATION_GROUP_TYPE:SPEC_ATTRIBUTES":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUP_TYPEOwner *models.RELATION_GROUP_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_attributesFormCallback.probe.stageOfInterest,
				a_spec_attributesFormCallback.probe.backRepoOfInterest,
				a_spec_attributes_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUP_TYPEOwner = reverseFieldOwner.(*models.RELATION_GROUP_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUP_TYPEOwner != nil {
					idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
					pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group_type := range *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](a_spec_attributesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUP_TYPEOwner := _relation_group_type // we have a match
						if pastRELATION_GROUP_TYPEOwner != nil {
							if newRELATION_GROUP_TYPEOwner != pastRELATION_GROUP_TYPEOwner {
								idx := slices.Index(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
								pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
								newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
							}
						} else {
							newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES = append(newRELATION_GROUP_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
						}
					}
				}
			}
		case "SPECIFICATION_TYPE:SPEC_ATTRIBUTES":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATION_TYPEOwner *models.SPECIFICATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_attributesFormCallback.probe.stageOfInterest,
				a_spec_attributesFormCallback.probe.backRepoOfInterest,
				a_spec_attributes_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATION_TYPEOwner = reverseFieldOwner.(*models.SPECIFICATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATION_TYPEOwner != nil {
					idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
					pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification_type := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](a_spec_attributesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATION_TYPEOwner := _specification_type // we have a match
						if pastSPECIFICATION_TYPEOwner != nil {
							if newSPECIFICATION_TYPEOwner != pastSPECIFICATION_TYPEOwner {
								idx := slices.Index(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
								pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
								newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
							}
						} else {
							newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES = append(newSPECIFICATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
						}
					}
				}
			}
		case "SPEC_OBJECT_TYPE:SPEC_ATTRIBUTES":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECT_TYPEOwner *models.SPEC_OBJECT_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_attributesFormCallback.probe.stageOfInterest,
				a_spec_attributesFormCallback.probe.backRepoOfInterest,
				a_spec_attributes_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECT_TYPEOwner = reverseFieldOwner.(*models.SPEC_OBJECT_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECT_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
					pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object_type := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](a_spec_attributesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECT_TYPEOwner := _spec_object_type // we have a match
						if pastSPEC_OBJECT_TYPEOwner != nil {
							if newSPEC_OBJECT_TYPEOwner != pastSPEC_OBJECT_TYPEOwner {
								idx := slices.Index(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
								pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
								newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
							}
						} else {
							newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES = append(newSPEC_OBJECT_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
						}
					}
				}
			}
		case "SPEC_RELATION_TYPE:SPEC_ATTRIBUTES":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATION_TYPEOwner *models.SPEC_RELATION_TYPE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION_TYPE"
			rf.Fieldname = "SPEC_ATTRIBUTES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_attributesFormCallback.probe.stageOfInterest,
				a_spec_attributesFormCallback.probe.backRepoOfInterest,
				a_spec_attributes_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATION_TYPEOwner = reverseFieldOwner.(*models.SPEC_RELATION_TYPE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATION_TYPEOwner != nil {
					idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
					pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation_type := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](a_spec_attributesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation_type.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATION_TYPEOwner := _spec_relation_type // we have a match
						if pastSPEC_RELATION_TYPEOwner != nil {
							if newSPEC_RELATION_TYPEOwner != pastSPEC_RELATION_TYPEOwner {
								idx := slices.Index(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
								pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES = slices.Delete(pastSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, idx, idx+1)
								newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
							}
						} else {
							newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES = append(newSPEC_RELATION_TYPEOwner.SPEC_ATTRIBUTES, a_spec_attributes_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_attributes_.Unstage(a_spec_attributesFormCallback.probe.stageOfInterest)
	}

	a_spec_attributesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_ATTRIBUTES](
		a_spec_attributesFormCallback.probe,
	)
	a_spec_attributesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_attributesFormCallback.CreationMode || a_spec_attributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_attributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_attributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_ATTRIBUTESFormCallback(
			nil,
			a_spec_attributesFormCallback.probe,
			newFormGroup,
		)
		a_spec_attributes := new(models.A_SPEC_ATTRIBUTES)
		FillUpForm(a_spec_attributes, newFormGroup, a_spec_attributesFormCallback.probe)
		a_spec_attributesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_attributesFormCallback.probe)
}
func __gong__New__A_SPEC_OBJECTSFormCallback(
	a_spec_objects *models.A_SPEC_OBJECTS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_objectsFormCallback *A_SPEC_OBJECTSFormCallback) {
	a_spec_objectsFormCallback = new(A_SPEC_OBJECTSFormCallback)
	a_spec_objectsFormCallback.probe = probe
	a_spec_objectsFormCallback.a_spec_objects = a_spec_objects
	a_spec_objectsFormCallback.formGroup = formGroup

	a_spec_objectsFormCallback.CreationMode = (a_spec_objects == nil)

	return
}

type A_SPEC_OBJECTSFormCallback struct {
	a_spec_objects *models.A_SPEC_OBJECTS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_objectsFormCallback *A_SPEC_OBJECTSFormCallback) OnSave() {

	log.Println("A_SPEC_OBJECTSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_objectsFormCallback.probe.formStage.Checkout()

	if a_spec_objectsFormCallback.a_spec_objects == nil {
		a_spec_objectsFormCallback.a_spec_objects = new(models.A_SPEC_OBJECTS).Stage(a_spec_objectsFormCallback.probe.stageOfInterest)
	}
	a_spec_objects_ := a_spec_objectsFormCallback.a_spec_objects
	_ = a_spec_objects_

	for _, formDiv := range a_spec_objectsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_objects_.Name), formDiv)
		case "REQ_IF_CONTENT:SPEC_OBJECTS":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_OBJECTS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_objectsFormCallback.probe.stageOfInterest,
				a_spec_objectsFormCallback.probe.backRepoOfInterest,
				a_spec_objects_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS, a_spec_objects_)
					pastREQ_IF_CONTENTOwner.SPEC_OBJECTS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_spec_objectsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS, a_spec_objects_)
								pastREQ_IF_CONTENTOwner.SPEC_OBJECTS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_OBJECTS, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_OBJECTS = append(newREQ_IF_CONTENTOwner.SPEC_OBJECTS, a_spec_objects_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_OBJECTS = append(newREQ_IF_CONTENTOwner.SPEC_OBJECTS, a_spec_objects_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_objectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_objects_.Unstage(a_spec_objectsFormCallback.probe.stageOfInterest)
	}

	a_spec_objectsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_OBJECTS](
		a_spec_objectsFormCallback.probe,
	)
	a_spec_objectsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_objectsFormCallback.CreationMode || a_spec_objectsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_objectsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_objectsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_OBJECTSFormCallback(
			nil,
			a_spec_objectsFormCallback.probe,
			newFormGroup,
		)
		a_spec_objects := new(models.A_SPEC_OBJECTS)
		FillUpForm(a_spec_objects, newFormGroup, a_spec_objectsFormCallback.probe)
		a_spec_objectsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_objectsFormCallback.probe)
}
func __gong__New__A_SPEC_RELATIONSFormCallback(
	a_spec_relations *models.A_SPEC_RELATIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relationsFormCallback *A_SPEC_RELATIONSFormCallback) {
	a_spec_relationsFormCallback = new(A_SPEC_RELATIONSFormCallback)
	a_spec_relationsFormCallback.probe = probe
	a_spec_relationsFormCallback.a_spec_relations = a_spec_relations
	a_spec_relationsFormCallback.formGroup = formGroup

	a_spec_relationsFormCallback.CreationMode = (a_spec_relations == nil)

	return
}

type A_SPEC_RELATIONSFormCallback struct {
	a_spec_relations *models.A_SPEC_RELATIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relationsFormCallback *A_SPEC_RELATIONSFormCallback) OnSave() {

	log.Println("A_SPEC_RELATIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relationsFormCallback.probe.formStage.Checkout()

	if a_spec_relationsFormCallback.a_spec_relations == nil {
		a_spec_relationsFormCallback.a_spec_relations = new(models.A_SPEC_RELATIONS).Stage(a_spec_relationsFormCallback.probe.stageOfInterest)
	}
	a_spec_relations_ := a_spec_relationsFormCallback.a_spec_relations
	_ = a_spec_relations_

	for _, formDiv := range a_spec_relationsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relations_.Name), formDiv)
		case "REQ_IF_CONTENT:SPEC_RELATIONS":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_relationsFormCallback.probe.stageOfInterest,
				a_spec_relationsFormCallback.probe.backRepoOfInterest,
				a_spec_relations_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS, a_spec_relations_)
					pastREQ_IF_CONTENTOwner.SPEC_RELATIONS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_spec_relationsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS, a_spec_relations_)
								pastREQ_IF_CONTENTOwner.SPEC_RELATIONS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATIONS, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_RELATIONS = append(newREQ_IF_CONTENTOwner.SPEC_RELATIONS, a_spec_relations_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_RELATIONS = append(newREQ_IF_CONTENTOwner.SPEC_RELATIONS, a_spec_relations_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_relationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relations_.Unstage(a_spec_relationsFormCallback.probe.stageOfInterest)
	}

	a_spec_relationsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_RELATIONS](
		a_spec_relationsFormCallback.probe,
	)
	a_spec_relationsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_relationsFormCallback.CreationMode || a_spec_relationsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relationsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_relationsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATIONSFormCallback(
			nil,
			a_spec_relationsFormCallback.probe,
			newFormGroup,
		)
		a_spec_relations := new(models.A_SPEC_RELATIONS)
		FillUpForm(a_spec_relations, newFormGroup, a_spec_relationsFormCallback.probe)
		a_spec_relationsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_relationsFormCallback.probe)
}
func __gong__New__A_SPEC_RELATIONS_1FormCallback(
	a_spec_relations_1 *models.A_SPEC_RELATIONS_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relations_1FormCallback *A_SPEC_RELATIONS_1FormCallback) {
	a_spec_relations_1FormCallback = new(A_SPEC_RELATIONS_1FormCallback)
	a_spec_relations_1FormCallback.probe = probe
	a_spec_relations_1FormCallback.a_spec_relations_1 = a_spec_relations_1
	a_spec_relations_1FormCallback.formGroup = formGroup

	a_spec_relations_1FormCallback.CreationMode = (a_spec_relations_1 == nil)

	return
}

type A_SPEC_RELATIONS_1FormCallback struct {
	a_spec_relations_1 *models.A_SPEC_RELATIONS_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relations_1FormCallback *A_SPEC_RELATIONS_1FormCallback) OnSave() {

	log.Println("A_SPEC_RELATIONS_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relations_1FormCallback.probe.formStage.Checkout()

	if a_spec_relations_1FormCallback.a_spec_relations_1 == nil {
		a_spec_relations_1FormCallback.a_spec_relations_1 = new(models.A_SPEC_RELATIONS_1).Stage(a_spec_relations_1FormCallback.probe.stageOfInterest)
	}
	a_spec_relations_1_ := a_spec_relations_1FormCallback.a_spec_relations_1
	_ = a_spec_relations_1_

	for _, formDiv := range a_spec_relations_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relations_1_.Name), formDiv)
		case "SPEC_RELATION_REF":
			FormDivBasicFieldToField(&(a_spec_relations_1_.SPEC_RELATION_REF), formDiv)
		case "RELATION_GROUP:SPEC_RELATIONS":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUPOwner *models.RELATION_GROUP
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "SPEC_RELATIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_relations_1FormCallback.probe.stageOfInterest,
				a_spec_relations_1FormCallback.probe.backRepoOfInterest,
				a_spec_relations_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUPOwner = reverseFieldOwner.(*models.RELATION_GROUP)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUPOwner != nil {
					idx := slices.Index(pastRELATION_GROUPOwner.SPEC_RELATIONS, a_spec_relations_1_)
					pastRELATION_GROUPOwner.SPEC_RELATIONS = slices.Delete(pastRELATION_GROUPOwner.SPEC_RELATIONS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group := range *models.GetGongstructInstancesSet[models.RELATION_GROUP](a_spec_relations_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUPOwner := _relation_group // we have a match
						if pastRELATION_GROUPOwner != nil {
							if newRELATION_GROUPOwner != pastRELATION_GROUPOwner {
								idx := slices.Index(pastRELATION_GROUPOwner.SPEC_RELATIONS, a_spec_relations_1_)
								pastRELATION_GROUPOwner.SPEC_RELATIONS = slices.Delete(pastRELATION_GROUPOwner.SPEC_RELATIONS, idx, idx+1)
								newRELATION_GROUPOwner.SPEC_RELATIONS = append(newRELATION_GROUPOwner.SPEC_RELATIONS, a_spec_relations_1_)
							}
						} else {
							newRELATION_GROUPOwner.SPEC_RELATIONS = append(newRELATION_GROUPOwner.SPEC_RELATIONS, a_spec_relations_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_relations_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relations_1_.Unstage(a_spec_relations_1FormCallback.probe.stageOfInterest)
	}

	a_spec_relations_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_RELATIONS_1](
		a_spec_relations_1FormCallback.probe,
	)
	a_spec_relations_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_relations_1FormCallback.CreationMode || a_spec_relations_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relations_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_relations_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATIONS_1FormCallback(
			nil,
			a_spec_relations_1FormCallback.probe,
			newFormGroup,
		)
		a_spec_relations_1 := new(models.A_SPEC_RELATIONS_1)
		FillUpForm(a_spec_relations_1, newFormGroup, a_spec_relations_1FormCallback.probe)
		a_spec_relations_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_relations_1FormCallback.probe)
}
func __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
	a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_relation_groupsFormCallback *A_SPEC_RELATION_GROUPSFormCallback) {
	a_spec_relation_groupsFormCallback = new(A_SPEC_RELATION_GROUPSFormCallback)
	a_spec_relation_groupsFormCallback.probe = probe
	a_spec_relation_groupsFormCallback.a_spec_relation_groups = a_spec_relation_groups
	a_spec_relation_groupsFormCallback.formGroup = formGroup

	a_spec_relation_groupsFormCallback.CreationMode = (a_spec_relation_groups == nil)

	return
}

type A_SPEC_RELATION_GROUPSFormCallback struct {
	a_spec_relation_groups *models.A_SPEC_RELATION_GROUPS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_relation_groupsFormCallback *A_SPEC_RELATION_GROUPSFormCallback) OnSave() {

	log.Println("A_SPEC_RELATION_GROUPSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_relation_groupsFormCallback.probe.formStage.Checkout()

	if a_spec_relation_groupsFormCallback.a_spec_relation_groups == nil {
		a_spec_relation_groupsFormCallback.a_spec_relation_groups = new(models.A_SPEC_RELATION_GROUPS).Stage(a_spec_relation_groupsFormCallback.probe.stageOfInterest)
	}
	a_spec_relation_groups_ := a_spec_relation_groupsFormCallback.a_spec_relation_groups
	_ = a_spec_relation_groups_

	for _, formDiv := range a_spec_relation_groupsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_relation_groups_.Name), formDiv)
		case "REQ_IF_CONTENT:SPEC_RELATION_GROUPS":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_RELATION_GROUPS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_relation_groupsFormCallback.probe.stageOfInterest,
				a_spec_relation_groupsFormCallback.probe.backRepoOfInterest,
				a_spec_relation_groups_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, a_spec_relation_groups_)
					pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_spec_relation_groupsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, a_spec_relation_groups_)
								pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS = append(newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, a_spec_relation_groups_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS = append(newREQ_IF_CONTENTOwner.SPEC_RELATION_GROUPS, a_spec_relation_groups_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_relation_groupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_groups_.Unstage(a_spec_relation_groupsFormCallback.probe.stageOfInterest)
	}

	a_spec_relation_groupsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_RELATION_GROUPS](
		a_spec_relation_groupsFormCallback.probe,
	)
	a_spec_relation_groupsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_relation_groupsFormCallback.CreationMode || a_spec_relation_groupsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_relation_groupsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_relation_groupsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_RELATION_GROUPSFormCallback(
			nil,
			a_spec_relation_groupsFormCallback.probe,
			newFormGroup,
		)
		a_spec_relation_groups := new(models.A_SPEC_RELATION_GROUPS)
		FillUpForm(a_spec_relation_groups, newFormGroup, a_spec_relation_groupsFormCallback.probe)
		a_spec_relation_groupsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_relation_groupsFormCallback.probe)
}
func __gong__New__A_SPEC_TYPESFormCallback(
	a_spec_types *models.A_SPEC_TYPES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_spec_typesFormCallback *A_SPEC_TYPESFormCallback) {
	a_spec_typesFormCallback = new(A_SPEC_TYPESFormCallback)
	a_spec_typesFormCallback.probe = probe
	a_spec_typesFormCallback.a_spec_types = a_spec_types
	a_spec_typesFormCallback.formGroup = formGroup

	a_spec_typesFormCallback.CreationMode = (a_spec_types == nil)

	return
}

type A_SPEC_TYPESFormCallback struct {
	a_spec_types *models.A_SPEC_TYPES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_spec_typesFormCallback *A_SPEC_TYPESFormCallback) OnSave() {

	log.Println("A_SPEC_TYPESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_spec_typesFormCallback.probe.formStage.Checkout()

	if a_spec_typesFormCallback.a_spec_types == nil {
		a_spec_typesFormCallback.a_spec_types = new(models.A_SPEC_TYPES).Stage(a_spec_typesFormCallback.probe.stageOfInterest)
	}
	a_spec_types_ := a_spec_typesFormCallback.a_spec_types
	_ = a_spec_types_

	for _, formDiv := range a_spec_typesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_spec_types_.Name), formDiv)
		case "REQ_IF_CONTENT:SPEC_TYPES":
			// we need to retrieve the field owner before the change
			var pastREQ_IF_CONTENTOwner *models.REQ_IF_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF_CONTENT"
			rf.Fieldname = "SPEC_TYPES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_spec_typesFormCallback.probe.stageOfInterest,
				a_spec_typesFormCallback.probe.backRepoOfInterest,
				a_spec_types_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IF_CONTENTOwner = reverseFieldOwner.(*models.REQ_IF_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IF_CONTENTOwner != nil {
					idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES, a_spec_types_)
					pastREQ_IF_CONTENTOwner.SPEC_TYPES = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if_content := range *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](a_spec_typesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IF_CONTENTOwner := _req_if_content // we have a match
						if pastREQ_IF_CONTENTOwner != nil {
							if newREQ_IF_CONTENTOwner != pastREQ_IF_CONTENTOwner {
								idx := slices.Index(pastREQ_IF_CONTENTOwner.SPEC_TYPES, a_spec_types_)
								pastREQ_IF_CONTENTOwner.SPEC_TYPES = slices.Delete(pastREQ_IF_CONTENTOwner.SPEC_TYPES, idx, idx+1)
								newREQ_IF_CONTENTOwner.SPEC_TYPES = append(newREQ_IF_CONTENTOwner.SPEC_TYPES, a_spec_types_)
							}
						} else {
							newREQ_IF_CONTENTOwner.SPEC_TYPES = append(newREQ_IF_CONTENTOwner.SPEC_TYPES, a_spec_types_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_spec_typesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_types_.Unstage(a_spec_typesFormCallback.probe.stageOfInterest)
	}

	a_spec_typesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_SPEC_TYPES](
		a_spec_typesFormCallback.probe,
	)
	a_spec_typesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_spec_typesFormCallback.CreationMode || a_spec_typesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_spec_typesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_spec_typesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_SPEC_TYPESFormCallback(
			nil,
			a_spec_typesFormCallback.probe,
			newFormGroup,
		)
		a_spec_types := new(models.A_SPEC_TYPES)
		FillUpForm(a_spec_types, newFormGroup, a_spec_typesFormCallback.probe)
		a_spec_typesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_spec_typesFormCallback.probe)
}
func __gong__New__A_THE_HEADERFormCallback(
	a_the_header *models.A_THE_HEADER,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_the_headerFormCallback *A_THE_HEADERFormCallback) {
	a_the_headerFormCallback = new(A_THE_HEADERFormCallback)
	a_the_headerFormCallback.probe = probe
	a_the_headerFormCallback.a_the_header = a_the_header
	a_the_headerFormCallback.formGroup = formGroup

	a_the_headerFormCallback.CreationMode = (a_the_header == nil)

	return
}

type A_THE_HEADERFormCallback struct {
	a_the_header *models.A_THE_HEADER

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_the_headerFormCallback *A_THE_HEADERFormCallback) OnSave() {

	log.Println("A_THE_HEADERFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_the_headerFormCallback.probe.formStage.Checkout()

	if a_the_headerFormCallback.a_the_header == nil {
		a_the_headerFormCallback.a_the_header = new(models.A_THE_HEADER).Stage(a_the_headerFormCallback.probe.stageOfInterest)
	}
	a_the_header_ := a_the_headerFormCallback.a_the_header
	_ = a_the_header_

	for _, formDiv := range a_the_headerFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_the_header_.Name), formDiv)
		case "REQ_IF:THE_HEADER":
			// we need to retrieve the field owner before the change
			var pastREQ_IFOwner *models.REQ_IF
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "THE_HEADER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_the_headerFormCallback.probe.stageOfInterest,
				a_the_headerFormCallback.probe.backRepoOfInterest,
				a_the_header_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IFOwner = reverseFieldOwner.(*models.REQ_IF)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IFOwner != nil {
					idx := slices.Index(pastREQ_IFOwner.THE_HEADER, a_the_header_)
					pastREQ_IFOwner.THE_HEADER = slices.Delete(pastREQ_IFOwner.THE_HEADER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if := range *models.GetGongstructInstancesSet[models.REQ_IF](a_the_headerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IFOwner := _req_if // we have a match
						if pastREQ_IFOwner != nil {
							if newREQ_IFOwner != pastREQ_IFOwner {
								idx := slices.Index(pastREQ_IFOwner.THE_HEADER, a_the_header_)
								pastREQ_IFOwner.THE_HEADER = slices.Delete(pastREQ_IFOwner.THE_HEADER, idx, idx+1)
								newREQ_IFOwner.THE_HEADER = append(newREQ_IFOwner.THE_HEADER, a_the_header_)
							}
						} else {
							newREQ_IFOwner.THE_HEADER = append(newREQ_IFOwner.THE_HEADER, a_the_header_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_the_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_the_header_.Unstage(a_the_headerFormCallback.probe.stageOfInterest)
	}

	a_the_headerFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_THE_HEADER](
		a_the_headerFormCallback.probe,
	)
	a_the_headerFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_the_headerFormCallback.CreationMode || a_the_headerFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_the_headerFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_the_headerFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_THE_HEADERFormCallback(
			nil,
			a_the_headerFormCallback.probe,
			newFormGroup,
		)
		a_the_header := new(models.A_THE_HEADER)
		FillUpForm(a_the_header, newFormGroup, a_the_headerFormCallback.probe)
		a_the_headerFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_the_headerFormCallback.probe)
}
func __gong__New__A_TOOL_EXTENSIONSFormCallback(
	a_tool_extensions *models.A_TOOL_EXTENSIONS,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_tool_extensionsFormCallback *A_TOOL_EXTENSIONSFormCallback) {
	a_tool_extensionsFormCallback = new(A_TOOL_EXTENSIONSFormCallback)
	a_tool_extensionsFormCallback.probe = probe
	a_tool_extensionsFormCallback.a_tool_extensions = a_tool_extensions
	a_tool_extensionsFormCallback.formGroup = formGroup

	a_tool_extensionsFormCallback.CreationMode = (a_tool_extensions == nil)

	return
}

type A_TOOL_EXTENSIONSFormCallback struct {
	a_tool_extensions *models.A_TOOL_EXTENSIONS

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_tool_extensionsFormCallback *A_TOOL_EXTENSIONSFormCallback) OnSave() {

	log.Println("A_TOOL_EXTENSIONSFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_tool_extensionsFormCallback.probe.formStage.Checkout()

	if a_tool_extensionsFormCallback.a_tool_extensions == nil {
		a_tool_extensionsFormCallback.a_tool_extensions = new(models.A_TOOL_EXTENSIONS).Stage(a_tool_extensionsFormCallback.probe.stageOfInterest)
	}
	a_tool_extensions_ := a_tool_extensionsFormCallback.a_tool_extensions
	_ = a_tool_extensions_

	for _, formDiv := range a_tool_extensionsFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_tool_extensions_.Name), formDiv)
		case "REQ_IF:TOOL_EXTENSIONS":
			// we need to retrieve the field owner before the change
			var pastREQ_IFOwner *models.REQ_IF
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "REQ_IF"
			rf.Fieldname = "TOOL_EXTENSIONS"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_tool_extensionsFormCallback.probe.stageOfInterest,
				a_tool_extensionsFormCallback.probe.backRepoOfInterest,
				a_tool_extensions_,
				&rf)

			if reverseFieldOwner != nil {
				pastREQ_IFOwner = reverseFieldOwner.(*models.REQ_IF)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastREQ_IFOwner != nil {
					idx := slices.Index(pastREQ_IFOwner.TOOL_EXTENSIONS, a_tool_extensions_)
					pastREQ_IFOwner.TOOL_EXTENSIONS = slices.Delete(pastREQ_IFOwner.TOOL_EXTENSIONS, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _req_if := range *models.GetGongstructInstancesSet[models.REQ_IF](a_tool_extensionsFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _req_if.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newREQ_IFOwner := _req_if // we have a match
						if pastREQ_IFOwner != nil {
							if newREQ_IFOwner != pastREQ_IFOwner {
								idx := slices.Index(pastREQ_IFOwner.TOOL_EXTENSIONS, a_tool_extensions_)
								pastREQ_IFOwner.TOOL_EXTENSIONS = slices.Delete(pastREQ_IFOwner.TOOL_EXTENSIONS, idx, idx+1)
								newREQ_IFOwner.TOOL_EXTENSIONS = append(newREQ_IFOwner.TOOL_EXTENSIONS, a_tool_extensions_)
							}
						} else {
							newREQ_IFOwner.TOOL_EXTENSIONS = append(newREQ_IFOwner.TOOL_EXTENSIONS, a_tool_extensions_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_tool_extensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_tool_extensions_.Unstage(a_tool_extensionsFormCallback.probe.stageOfInterest)
	}

	a_tool_extensionsFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_TOOL_EXTENSIONS](
		a_tool_extensionsFormCallback.probe,
	)
	a_tool_extensionsFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_tool_extensionsFormCallback.CreationMode || a_tool_extensionsFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_tool_extensionsFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_tool_extensionsFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_TOOL_EXTENSIONSFormCallback(
			nil,
			a_tool_extensionsFormCallback.probe,
			newFormGroup,
		)
		a_tool_extensions := new(models.A_TOOL_EXTENSIONS)
		FillUpForm(a_tool_extensions, newFormGroup, a_tool_extensionsFormCallback.probe)
		a_tool_extensionsFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_tool_extensionsFormCallback.probe)
}
func __gong__New__A_VALUESFormCallback(
	a_values *models.A_VALUES,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_valuesFormCallback *A_VALUESFormCallback) {
	a_valuesFormCallback = new(A_VALUESFormCallback)
	a_valuesFormCallback.probe = probe
	a_valuesFormCallback.a_values = a_values
	a_valuesFormCallback.formGroup = formGroup

	a_valuesFormCallback.CreationMode = (a_values == nil)

	return
}

type A_VALUESFormCallback struct {
	a_values *models.A_VALUES

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_valuesFormCallback *A_VALUESFormCallback) OnSave() {

	log.Println("A_VALUESFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_valuesFormCallback.probe.formStage.Checkout()

	if a_valuesFormCallback.a_values == nil {
		a_valuesFormCallback.a_values = new(models.A_VALUES).Stage(a_valuesFormCallback.probe.stageOfInterest)
	}
	a_values_ := a_valuesFormCallback.a_values
	_ = a_values_

	for _, formDiv := range a_valuesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_values_.Name), formDiv)
		case "ENUM_VALUE_REF":
			FormDivBasicFieldToField(&(a_values_.ENUM_VALUE_REF), formDiv)
		case "ATTRIBUTE_VALUE_ENUMERATION:VALUES":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_ENUMERATIONOwner *models.ATTRIBUTE_VALUE_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_ENUMERATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_valuesFormCallback.probe.stageOfInterest,
				a_valuesFormCallback.probe.backRepoOfInterest,
				a_values_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, a_values_)
					pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES = slices.Delete(pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](a_valuesFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_ENUMERATIONOwner := _attribute_value_enumeration // we have a match
						if pastATTRIBUTE_VALUE_ENUMERATIONOwner != nil {
							if newATTRIBUTE_VALUE_ENUMERATIONOwner != pastATTRIBUTE_VALUE_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, a_values_)
								pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES = slices.Delete(pastATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, idx, idx+1)
								newATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES = append(newATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, a_values_)
							}
						} else {
							newATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES = append(newATTRIBUTE_VALUE_ENUMERATIONOwner.VALUES, a_values_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_values_.Unstage(a_valuesFormCallback.probe.stageOfInterest)
	}

	a_valuesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_VALUES](
		a_valuesFormCallback.probe,
	)
	a_valuesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_valuesFormCallback.CreationMode || a_valuesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_valuesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_valuesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_VALUESFormCallback(
			nil,
			a_valuesFormCallback.probe,
			newFormGroup,
		)
		a_values := new(models.A_VALUES)
		FillUpForm(a_values, newFormGroup, a_valuesFormCallback.probe)
		a_valuesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_valuesFormCallback.probe)
}
func __gong__New__A_VALUES_1FormCallback(
	a_values_1 *models.A_VALUES_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_values_1FormCallback *A_VALUES_1FormCallback) {
	a_values_1FormCallback = new(A_VALUES_1FormCallback)
	a_values_1FormCallback.probe = probe
	a_values_1FormCallback.a_values_1 = a_values_1
	a_values_1FormCallback.formGroup = formGroup

	a_values_1FormCallback.CreationMode = (a_values_1 == nil)

	return
}

type A_VALUES_1FormCallback struct {
	a_values_1 *models.A_VALUES_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_values_1FormCallback *A_VALUES_1FormCallback) OnSave() {

	log.Println("A_VALUES_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_values_1FormCallback.probe.formStage.Checkout()

	if a_values_1FormCallback.a_values_1 == nil {
		a_values_1FormCallback.a_values_1 = new(models.A_VALUES_1).Stage(a_values_1FormCallback.probe.stageOfInterest)
	}
	a_values_1_ := a_values_1FormCallback.a_values_1
	_ = a_values_1_

	for _, formDiv := range a_values_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_values_1_.Name), formDiv)
		case "SPECIFICATION:VALUES":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_values_1FormCallback.probe.stageOfInterest,
				a_values_1FormCallback.probe.backRepoOfInterest,
				a_values_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.VALUES, a_values_1_)
					pastSPECIFICATIONOwner.VALUES = slices.Delete(pastSPECIFICATIONOwner.VALUES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](a_values_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.VALUES, a_values_1_)
								pastSPECIFICATIONOwner.VALUES = slices.Delete(pastSPECIFICATIONOwner.VALUES, idx, idx+1)
								newSPECIFICATIONOwner.VALUES = append(newSPECIFICATIONOwner.VALUES, a_values_1_)
							}
						} else {
							newSPECIFICATIONOwner.VALUES = append(newSPECIFICATIONOwner.VALUES, a_values_1_)
						}
					}
				}
			}
		case "SPEC_OBJECT:VALUES":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_values_1FormCallback.probe.stageOfInterest,
				a_values_1FormCallback.probe.backRepoOfInterest,
				a_values_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.VALUES, a_values_1_)
					pastSPEC_OBJECTOwner.VALUES = slices.Delete(pastSPEC_OBJECTOwner.VALUES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](a_values_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.VALUES, a_values_1_)
								pastSPEC_OBJECTOwner.VALUES = slices.Delete(pastSPEC_OBJECTOwner.VALUES, idx, idx+1)
								newSPEC_OBJECTOwner.VALUES = append(newSPEC_OBJECTOwner.VALUES, a_values_1_)
							}
						} else {
							newSPEC_OBJECTOwner.VALUES = append(newSPEC_OBJECTOwner.VALUES, a_values_1_)
						}
					}
				}
			}
		case "SPEC_RELATION:VALUES":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "VALUES"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				a_values_1FormCallback.probe.stageOfInterest,
				a_values_1FormCallback.probe.backRepoOfInterest,
				a_values_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.VALUES, a_values_1_)
					pastSPEC_RELATIONOwner.VALUES = slices.Delete(pastSPEC_RELATIONOwner.VALUES, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](a_values_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.VALUES, a_values_1_)
								pastSPEC_RELATIONOwner.VALUES = slices.Delete(pastSPEC_RELATIONOwner.VALUES, idx, idx+1)
								newSPEC_RELATIONOwner.VALUES = append(newSPEC_RELATIONOwner.VALUES, a_values_1_)
							}
						} else {
							newSPEC_RELATIONOwner.VALUES = append(newSPEC_RELATIONOwner.VALUES, a_values_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if a_values_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_values_1_.Unstage(a_values_1FormCallback.probe.stageOfInterest)
	}

	a_values_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.A_VALUES_1](
		a_values_1FormCallback.probe,
	)
	a_values_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if a_values_1FormCallback.CreationMode || a_values_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_values_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(a_values_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_VALUES_1FormCallback(
			nil,
			a_values_1FormCallback.probe,
			newFormGroup,
		)
		a_values_1 := new(models.A_VALUES_1)
		FillUpForm(a_values_1, newFormGroup, a_values_1FormCallback.probe)
		a_values_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(a_values_1FormCallback.probe)
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
		case "A_DATATYPES:DATATYPE_DEFINITION_BOOLEAN":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_BOOLEAN"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_booleanFormCallback.probe.stageOfInterest,
				datatype_definition_booleanFormCallback.probe.backRepoOfInterest,
				datatype_definition_boolean_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_booleanFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_BOOLEAN, datatype_definition_boolean_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_DATE":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_DATE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_dateFormCallback.probe.stageOfInterest,
				datatype_definition_dateFormCallback.probe.backRepoOfInterest,
				datatype_definition_date_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_dateFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_DATE = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_DATE = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_DATE, datatype_definition_date_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_ENUMERATION":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_ENUMERATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_enumerationFormCallback.probe.stageOfInterest,
				datatype_definition_enumerationFormCallback.probe.backRepoOfInterest,
				datatype_definition_enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_ENUMERATION, datatype_definition_enumeration_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_INTEGER":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_INTEGER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_integerFormCallback.probe.stageOfInterest,
				datatype_definition_integerFormCallback.probe.backRepoOfInterest,
				datatype_definition_integer_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_integerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_INTEGER, datatype_definition_integer_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_REAL":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_REAL"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_realFormCallback.probe.stageOfInterest,
				datatype_definition_realFormCallback.probe.backRepoOfInterest,
				datatype_definition_real_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_realFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_REAL = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_REAL = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_REAL, datatype_definition_real_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_STRING":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_STRING"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_stringFormCallback.probe.stageOfInterest,
				datatype_definition_stringFormCallback.probe.backRepoOfInterest,
				datatype_definition_string_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_stringFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_STRING = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_STRING = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_STRING, datatype_definition_string_)
						}
					}
				}
			}
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
		case "A_DATATYPES:DATATYPE_DEFINITION_XHTML":
			// we need to retrieve the field owner before the change
			var pastA_DATATYPESOwner *models.A_DATATYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_DATATYPES"
			rf.Fieldname = "DATATYPE_DEFINITION_XHTML"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				datatype_definition_xhtmlFormCallback.probe.stageOfInterest,
				datatype_definition_xhtmlFormCallback.probe.backRepoOfInterest,
				datatype_definition_xhtml_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_DATATYPESOwner = reverseFieldOwner.(*models.A_DATATYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_DATATYPESOwner != nil {
					idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
					pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_datatypes := range *models.GetGongstructInstancesSet[models.A_DATATYPES](datatype_definition_xhtmlFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_datatypes.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_DATATYPESOwner := _a_datatypes // we have a match
						if pastA_DATATYPESOwner != nil {
							if newA_DATATYPESOwner != pastA_DATATYPESOwner {
								idx := slices.Index(pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
								pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML = slices.Delete(pastA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, idx, idx+1)
								newA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
							}
						} else {
							newA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML = append(newA_DATATYPESOwner.DATATYPE_DEFINITION_XHTML, datatype_definition_xhtml_)
						}
					}
				}
			}
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
		case "A_PROPERTIES:EMBEDDED_VALUE":
			// we need to retrieve the field owner before the change
			var pastA_PROPERTIESOwner *models.A_PROPERTIES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_PROPERTIES"
			rf.Fieldname = "EMBEDDED_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				embedded_valueFormCallback.probe.stageOfInterest,
				embedded_valueFormCallback.probe.backRepoOfInterest,
				embedded_value_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_PROPERTIESOwner = reverseFieldOwner.(*models.A_PROPERTIES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_PROPERTIESOwner != nil {
					idx := slices.Index(pastA_PROPERTIESOwner.EMBEDDED_VALUE, embedded_value_)
					pastA_PROPERTIESOwner.EMBEDDED_VALUE = slices.Delete(pastA_PROPERTIESOwner.EMBEDDED_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_properties := range *models.GetGongstructInstancesSet[models.A_PROPERTIES](embedded_valueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_properties.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_PROPERTIESOwner := _a_properties // we have a match
						if pastA_PROPERTIESOwner != nil {
							if newA_PROPERTIESOwner != pastA_PROPERTIESOwner {
								idx := slices.Index(pastA_PROPERTIESOwner.EMBEDDED_VALUE, embedded_value_)
								pastA_PROPERTIESOwner.EMBEDDED_VALUE = slices.Delete(pastA_PROPERTIESOwner.EMBEDDED_VALUE, idx, idx+1)
								newA_PROPERTIESOwner.EMBEDDED_VALUE = append(newA_PROPERTIESOwner.EMBEDDED_VALUE, embedded_value_)
							}
						} else {
							newA_PROPERTIESOwner.EMBEDDED_VALUE = append(newA_PROPERTIESOwner.EMBEDDED_VALUE, embedded_value_)
						}
					}
				}
			}
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
		case "A_SPECIFIED_VALUES:ENUM_VALUE":
			// we need to retrieve the field owner before the change
			var pastA_SPECIFIED_VALUESOwner *models.A_SPECIFIED_VALUES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPECIFIED_VALUES"
			rf.Fieldname = "ENUM_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				enum_valueFormCallback.probe.stageOfInterest,
				enum_valueFormCallback.probe.backRepoOfInterest,
				enum_value_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPECIFIED_VALUESOwner = reverseFieldOwner.(*models.A_SPECIFIED_VALUES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPECIFIED_VALUESOwner != nil {
					idx := slices.Index(pastA_SPECIFIED_VALUESOwner.ENUM_VALUE, enum_value_)
					pastA_SPECIFIED_VALUESOwner.ENUM_VALUE = slices.Delete(pastA_SPECIFIED_VALUESOwner.ENUM_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_specified_values := range *models.GetGongstructInstancesSet[models.A_SPECIFIED_VALUES](enum_valueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_specified_values.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPECIFIED_VALUESOwner := _a_specified_values // we have a match
						if pastA_SPECIFIED_VALUESOwner != nil {
							if newA_SPECIFIED_VALUESOwner != pastA_SPECIFIED_VALUESOwner {
								idx := slices.Index(pastA_SPECIFIED_VALUESOwner.ENUM_VALUE, enum_value_)
								pastA_SPECIFIED_VALUESOwner.ENUM_VALUE = slices.Delete(pastA_SPECIFIED_VALUESOwner.ENUM_VALUE, idx, idx+1)
								newA_SPECIFIED_VALUESOwner.ENUM_VALUE = append(newA_SPECIFIED_VALUESOwner.ENUM_VALUE, enum_value_)
							}
						} else {
							newA_SPECIFIED_VALUESOwner.ENUM_VALUE = append(newA_SPECIFIED_VALUESOwner.ENUM_VALUE, enum_value_)
						}
					}
				}
			}
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
		case "A_SPEC_RELATION_GROUPS:RELATION_GROUP":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_RELATION_GROUPSOwner *models.A_SPEC_RELATION_GROUPS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_RELATION_GROUPS"
			rf.Fieldname = "RELATION_GROUP"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relation_groupFormCallback.probe.stageOfInterest,
				relation_groupFormCallback.probe.backRepoOfInterest,
				relation_group_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_RELATION_GROUPSOwner = reverseFieldOwner.(*models.A_SPEC_RELATION_GROUPS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_RELATION_GROUPSOwner != nil {
					idx := slices.Index(pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, relation_group_)
					pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP = slices.Delete(pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_relation_groups := range *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_GROUPS](relation_groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_relation_groups.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_RELATION_GROUPSOwner := _a_spec_relation_groups // we have a match
						if pastA_SPEC_RELATION_GROUPSOwner != nil {
							if newA_SPEC_RELATION_GROUPSOwner != pastA_SPEC_RELATION_GROUPSOwner {
								idx := slices.Index(pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, relation_group_)
								pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP = slices.Delete(pastA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, idx, idx+1)
								newA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP = append(newA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, relation_group_)
							}
						} else {
							newA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP = append(newA_SPEC_RELATION_GROUPSOwner.RELATION_GROUP, relation_group_)
						}
					}
				}
			}
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
		case "A_SPEC_TYPES:RELATION_GROUP_TYPE":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_TYPESOwner *models.A_SPEC_TYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "RELATION_GROUP_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				relation_group_typeFormCallback.probe.stageOfInterest,
				relation_group_typeFormCallback.probe.backRepoOfInterest,
				relation_group_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_TYPESOwner = reverseFieldOwner.(*models.A_SPEC_TYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_TYPESOwner != nil {
					idx := slices.Index(pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, relation_group_type_)
					pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](relation_group_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_types.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_TYPESOwner := _a_spec_types // we have a match
						if pastA_SPEC_TYPESOwner != nil {
							if newA_SPEC_TYPESOwner != pastA_SPEC_TYPESOwner {
								idx := slices.Index(pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, relation_group_type_)
								pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, idx, idx+1)
								newA_SPEC_TYPESOwner.RELATION_GROUP_TYPE = append(newA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, relation_group_type_)
							}
						} else {
							newA_SPEC_TYPESOwner.RELATION_GROUP_TYPE = append(newA_SPEC_TYPESOwner.RELATION_GROUP_TYPE, relation_group_type_)
						}
					}
				}
			}
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
		case "A_CORE_CONTENT:REQ_IF_CONTENT":
			// we need to retrieve the field owner before the change
			var pastA_CORE_CONTENTOwner *models.A_CORE_CONTENT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_CORE_CONTENT"
			rf.Fieldname = "REQ_IF_CONTENT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				req_if_contentFormCallback.probe.stageOfInterest,
				req_if_contentFormCallback.probe.backRepoOfInterest,
				req_if_content_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_CORE_CONTENTOwner = reverseFieldOwner.(*models.A_CORE_CONTENT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_CORE_CONTENTOwner != nil {
					idx := slices.Index(pastA_CORE_CONTENTOwner.REQ_IF_CONTENT, req_if_content_)
					pastA_CORE_CONTENTOwner.REQ_IF_CONTENT = slices.Delete(pastA_CORE_CONTENTOwner.REQ_IF_CONTENT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_core_content := range *models.GetGongstructInstancesSet[models.A_CORE_CONTENT](req_if_contentFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_core_content.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_CORE_CONTENTOwner := _a_core_content // we have a match
						if pastA_CORE_CONTENTOwner != nil {
							if newA_CORE_CONTENTOwner != pastA_CORE_CONTENTOwner {
								idx := slices.Index(pastA_CORE_CONTENTOwner.REQ_IF_CONTENT, req_if_content_)
								pastA_CORE_CONTENTOwner.REQ_IF_CONTENT = slices.Delete(pastA_CORE_CONTENTOwner.REQ_IF_CONTENT, idx, idx+1)
								newA_CORE_CONTENTOwner.REQ_IF_CONTENT = append(newA_CORE_CONTENTOwner.REQ_IF_CONTENT, req_if_content_)
							}
						} else {
							newA_CORE_CONTENTOwner.REQ_IF_CONTENT = append(newA_CORE_CONTENTOwner.REQ_IF_CONTENT, req_if_content_)
						}
					}
				}
			}
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
		case "A_THE_HEADER:REQ_IF_HEADER":
			// we need to retrieve the field owner before the change
			var pastA_THE_HEADEROwner *models.A_THE_HEADER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_THE_HEADER"
			rf.Fieldname = "REQ_IF_HEADER"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				req_if_headerFormCallback.probe.stageOfInterest,
				req_if_headerFormCallback.probe.backRepoOfInterest,
				req_if_header_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_THE_HEADEROwner = reverseFieldOwner.(*models.A_THE_HEADER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_THE_HEADEROwner != nil {
					idx := slices.Index(pastA_THE_HEADEROwner.REQ_IF_HEADER, req_if_header_)
					pastA_THE_HEADEROwner.REQ_IF_HEADER = slices.Delete(pastA_THE_HEADEROwner.REQ_IF_HEADER, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_the_header := range *models.GetGongstructInstancesSet[models.A_THE_HEADER](req_if_headerFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_the_header.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_THE_HEADEROwner := _a_the_header // we have a match
						if pastA_THE_HEADEROwner != nil {
							if newA_THE_HEADEROwner != pastA_THE_HEADEROwner {
								idx := slices.Index(pastA_THE_HEADEROwner.REQ_IF_HEADER, req_if_header_)
								pastA_THE_HEADEROwner.REQ_IF_HEADER = slices.Delete(pastA_THE_HEADEROwner.REQ_IF_HEADER, idx, idx+1)
								newA_THE_HEADEROwner.REQ_IF_HEADER = append(newA_THE_HEADEROwner.REQ_IF_HEADER, req_if_header_)
							}
						} else {
							newA_THE_HEADEROwner.REQ_IF_HEADER = append(newA_THE_HEADEROwner.REQ_IF_HEADER, req_if_header_)
						}
					}
				}
			}
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
		case "A_TOOL_EXTENSIONS:REQ_IF_TOOL_EXTENSION":
			// we need to retrieve the field owner before the change
			var pastA_TOOL_EXTENSIONSOwner *models.A_TOOL_EXTENSIONS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_TOOL_EXTENSIONS"
			rf.Fieldname = "REQ_IF_TOOL_EXTENSION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				req_if_tool_extensionFormCallback.probe.stageOfInterest,
				req_if_tool_extensionFormCallback.probe.backRepoOfInterest,
				req_if_tool_extension_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_TOOL_EXTENSIONSOwner = reverseFieldOwner.(*models.A_TOOL_EXTENSIONS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_TOOL_EXTENSIONSOwner != nil {
					idx := slices.Index(pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
					pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION = slices.Delete(pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_tool_extensions := range *models.GetGongstructInstancesSet[models.A_TOOL_EXTENSIONS](req_if_tool_extensionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_tool_extensions.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_TOOL_EXTENSIONSOwner := _a_tool_extensions // we have a match
						if pastA_TOOL_EXTENSIONSOwner != nil {
							if newA_TOOL_EXTENSIONSOwner != pastA_TOOL_EXTENSIONSOwner {
								idx := slices.Index(pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
								pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION = slices.Delete(pastA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, idx, idx+1)
								newA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION = append(newA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
							}
						} else {
							newA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION = append(newA_TOOL_EXTENSIONSOwner.REQ_IF_TOOL_EXTENSION, req_if_tool_extension_)
						}
					}
				}
			}
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
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback(
	renamed_attribute_definition_boolean_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_boolean_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback) {
	renamed_attribute_definition_boolean_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback)
	renamed_attribute_definition_boolean_ref_1FormCallback.probe = probe
	renamed_attribute_definition_boolean_ref_1FormCallback.renamed_attribute_definition_boolean_ref_1 = renamed_attribute_definition_boolean_ref_1
	renamed_attribute_definition_boolean_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_boolean_ref_1FormCallback.CreationMode = (renamed_attribute_definition_boolean_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback struct {
	renamed_attribute_definition_boolean_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_boolean_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_boolean_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_boolean_ref_1FormCallback.renamed_attribute_definition_boolean_ref_1 == nil {
		renamed_attribute_definition_boolean_ref_1FormCallback.renamed_attribute_definition_boolean_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1).Stage(renamed_attribute_definition_boolean_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_boolean_ref_1_ := renamed_attribute_definition_boolean_ref_1FormCallback.renamed_attribute_definition_boolean_ref_1
	_ = renamed_attribute_definition_boolean_ref_1_

	for _, formDiv := range renamed_attribute_definition_boolean_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_boolean_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_boolean_ref_1_.ATTRIBUTE_DEFINITION_BOOLEAN_REF), formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_BOOLEANOwner *models.ATTRIBUTE_VALUE_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_BOOLEAN"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_boolean_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_boolean_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_boolean_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, renamed_attribute_definition_boolean_ref_1_)
					pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_BOOLEAN](renamed_attribute_definition_boolean_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_BOOLEANOwner := _attribute_value_boolean // we have a match
						if pastATTRIBUTE_VALUE_BOOLEANOwner != nil {
							if newATTRIBUTE_VALUE_BOOLEANOwner != pastATTRIBUTE_VALUE_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, renamed_attribute_definition_boolean_ref_1_)
								pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION = append(newATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, renamed_attribute_definition_boolean_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION = append(newATTRIBUTE_VALUE_BOOLEANOwner.DEFINITION, renamed_attribute_definition_boolean_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_boolean_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_boolean_ref_1_.Unstage(renamed_attribute_definition_boolean_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_boolean_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1](
		renamed_attribute_definition_boolean_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_boolean_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_boolean_ref_1FormCallback.CreationMode || renamed_attribute_definition_boolean_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_boolean_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_boolean_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1FormCallback(
			nil,
			renamed_attribute_definition_boolean_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_boolean_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_BOOLEAN_REF_1)
		FillUpForm(renamed_attribute_definition_boolean_ref_1, newFormGroup, renamed_attribute_definition_boolean_ref_1FormCallback.probe)
		renamed_attribute_definition_boolean_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_boolean_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback(
	renamed_attribute_definition_date_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_date_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback) {
	renamed_attribute_definition_date_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback)
	renamed_attribute_definition_date_ref_1FormCallback.probe = probe
	renamed_attribute_definition_date_ref_1FormCallback.renamed_attribute_definition_date_ref_1 = renamed_attribute_definition_date_ref_1
	renamed_attribute_definition_date_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_date_ref_1FormCallback.CreationMode = (renamed_attribute_definition_date_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback struct {
	renamed_attribute_definition_date_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_date_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_date_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_date_ref_1FormCallback.renamed_attribute_definition_date_ref_1 == nil {
		renamed_attribute_definition_date_ref_1FormCallback.renamed_attribute_definition_date_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1).Stage(renamed_attribute_definition_date_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_date_ref_1_ := renamed_attribute_definition_date_ref_1FormCallback.renamed_attribute_definition_date_ref_1
	_ = renamed_attribute_definition_date_ref_1_

	for _, formDiv := range renamed_attribute_definition_date_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_date_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_date_ref_1_.ATTRIBUTE_DEFINITION_DATE_REF), formDiv)
		case "ATTRIBUTE_VALUE_DATE:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_DATEOwner *models.ATTRIBUTE_VALUE_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_DATE"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_date_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_date_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_date_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_DATEOwner.DEFINITION, renamed_attribute_definition_date_ref_1_)
					pastATTRIBUTE_VALUE_DATEOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_DATEOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_DATE](renamed_attribute_definition_date_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_DATEOwner := _attribute_value_date // we have a match
						if pastATTRIBUTE_VALUE_DATEOwner != nil {
							if newATTRIBUTE_VALUE_DATEOwner != pastATTRIBUTE_VALUE_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_DATEOwner.DEFINITION, renamed_attribute_definition_date_ref_1_)
								pastATTRIBUTE_VALUE_DATEOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_DATEOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_DATEOwner.DEFINITION = append(newATTRIBUTE_VALUE_DATEOwner.DEFINITION, renamed_attribute_definition_date_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_DATEOwner.DEFINITION = append(newATTRIBUTE_VALUE_DATEOwner.DEFINITION, renamed_attribute_definition_date_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_date_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_date_ref_1_.Unstage(renamed_attribute_definition_date_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_date_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1](
		renamed_attribute_definition_date_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_date_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_date_ref_1FormCallback.CreationMode || renamed_attribute_definition_date_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_date_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_date_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1FormCallback(
			nil,
			renamed_attribute_definition_date_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_date_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_DATE_REF_1)
		FillUpForm(renamed_attribute_definition_date_ref_1, newFormGroup, renamed_attribute_definition_date_ref_1FormCallback.probe)
		renamed_attribute_definition_date_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_date_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback(
	renamed_attribute_definition_enumeration_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_enumeration_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback) {
	renamed_attribute_definition_enumeration_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback)
	renamed_attribute_definition_enumeration_ref_1FormCallback.probe = probe
	renamed_attribute_definition_enumeration_ref_1FormCallback.renamed_attribute_definition_enumeration_ref_1 = renamed_attribute_definition_enumeration_ref_1
	renamed_attribute_definition_enumeration_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_enumeration_ref_1FormCallback.CreationMode = (renamed_attribute_definition_enumeration_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback struct {
	renamed_attribute_definition_enumeration_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_enumeration_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_enumeration_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_enumeration_ref_1FormCallback.renamed_attribute_definition_enumeration_ref_1 == nil {
		renamed_attribute_definition_enumeration_ref_1FormCallback.renamed_attribute_definition_enumeration_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1).Stage(renamed_attribute_definition_enumeration_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_enumeration_ref_1_ := renamed_attribute_definition_enumeration_ref_1FormCallback.renamed_attribute_definition_enumeration_ref_1
	_ = renamed_attribute_definition_enumeration_ref_1_

	for _, formDiv := range renamed_attribute_definition_enumeration_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_enumeration_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_enumeration_ref_1_.ATTRIBUTE_DEFINITION_ENUMERATION_REF), formDiv)
		case "ATTRIBUTE_VALUE_ENUMERATION:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_ENUMERATIONOwner *models.ATTRIBUTE_VALUE_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_ENUMERATION"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_enumeration_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_enumeration_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_enumeration_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, renamed_attribute_definition_enumeration_ref_1_)
					pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](renamed_attribute_definition_enumeration_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_ENUMERATIONOwner := _attribute_value_enumeration // we have a match
						if pastATTRIBUTE_VALUE_ENUMERATIONOwner != nil {
							if newATTRIBUTE_VALUE_ENUMERATIONOwner != pastATTRIBUTE_VALUE_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, renamed_attribute_definition_enumeration_ref_1_)
								pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION = append(newATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, renamed_attribute_definition_enumeration_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION = append(newATTRIBUTE_VALUE_ENUMERATIONOwner.DEFINITION, renamed_attribute_definition_enumeration_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_enumeration_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_enumeration_ref_1_.Unstage(renamed_attribute_definition_enumeration_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_enumeration_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1](
		renamed_attribute_definition_enumeration_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_enumeration_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_enumeration_ref_1FormCallback.CreationMode || renamed_attribute_definition_enumeration_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_enumeration_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_enumeration_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1FormCallback(
			nil,
			renamed_attribute_definition_enumeration_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_enumeration_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_ENUMERATION_REF_1)
		FillUpForm(renamed_attribute_definition_enumeration_ref_1, newFormGroup, renamed_attribute_definition_enumeration_ref_1FormCallback.probe)
		renamed_attribute_definition_enumeration_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_enumeration_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback(
	renamed_attribute_definition_integer_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_integer_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback) {
	renamed_attribute_definition_integer_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback)
	renamed_attribute_definition_integer_ref_1FormCallback.probe = probe
	renamed_attribute_definition_integer_ref_1FormCallback.renamed_attribute_definition_integer_ref_1 = renamed_attribute_definition_integer_ref_1
	renamed_attribute_definition_integer_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_integer_ref_1FormCallback.CreationMode = (renamed_attribute_definition_integer_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback struct {
	renamed_attribute_definition_integer_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_integer_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_integer_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_integer_ref_1FormCallback.renamed_attribute_definition_integer_ref_1 == nil {
		renamed_attribute_definition_integer_ref_1FormCallback.renamed_attribute_definition_integer_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1).Stage(renamed_attribute_definition_integer_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_integer_ref_1_ := renamed_attribute_definition_integer_ref_1FormCallback.renamed_attribute_definition_integer_ref_1
	_ = renamed_attribute_definition_integer_ref_1_

	for _, formDiv := range renamed_attribute_definition_integer_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_integer_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_integer_ref_1_.ATTRIBUTE_DEFINITION_INTEGER_REF), formDiv)
		case "ATTRIBUTE_VALUE_INTEGER:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_INTEGEROwner *models.ATTRIBUTE_VALUE_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_INTEGER"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_integer_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_integer_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_integer_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, renamed_attribute_definition_integer_ref_1_)
					pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_INTEGER](renamed_attribute_definition_integer_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_INTEGEROwner := _attribute_value_integer // we have a match
						if pastATTRIBUTE_VALUE_INTEGEROwner != nil {
							if newATTRIBUTE_VALUE_INTEGEROwner != pastATTRIBUTE_VALUE_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, renamed_attribute_definition_integer_ref_1_)
								pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_INTEGEROwner.DEFINITION = append(newATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, renamed_attribute_definition_integer_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_INTEGEROwner.DEFINITION = append(newATTRIBUTE_VALUE_INTEGEROwner.DEFINITION, renamed_attribute_definition_integer_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_integer_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_integer_ref_1_.Unstage(renamed_attribute_definition_integer_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_integer_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1](
		renamed_attribute_definition_integer_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_integer_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_integer_ref_1FormCallback.CreationMode || renamed_attribute_definition_integer_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_integer_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_integer_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1FormCallback(
			nil,
			renamed_attribute_definition_integer_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_integer_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_INTEGER_REF_1)
		FillUpForm(renamed_attribute_definition_integer_ref_1, newFormGroup, renamed_attribute_definition_integer_ref_1FormCallback.probe)
		renamed_attribute_definition_integer_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_integer_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback(
	renamed_attribute_definition_real_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_real_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback) {
	renamed_attribute_definition_real_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback)
	renamed_attribute_definition_real_ref_1FormCallback.probe = probe
	renamed_attribute_definition_real_ref_1FormCallback.renamed_attribute_definition_real_ref_1 = renamed_attribute_definition_real_ref_1
	renamed_attribute_definition_real_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_real_ref_1FormCallback.CreationMode = (renamed_attribute_definition_real_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback struct {
	renamed_attribute_definition_real_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_real_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_real_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_real_ref_1FormCallback.renamed_attribute_definition_real_ref_1 == nil {
		renamed_attribute_definition_real_ref_1FormCallback.renamed_attribute_definition_real_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1).Stage(renamed_attribute_definition_real_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_real_ref_1_ := renamed_attribute_definition_real_ref_1FormCallback.renamed_attribute_definition_real_ref_1
	_ = renamed_attribute_definition_real_ref_1_

	for _, formDiv := range renamed_attribute_definition_real_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_real_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_real_ref_1_.ATTRIBUTE_DEFINITION_REAL_REF), formDiv)
		case "ATTRIBUTE_VALUE_REAL:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_REALOwner *models.ATTRIBUTE_VALUE_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_REAL"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_real_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_real_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_real_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_REALOwner.DEFINITION, renamed_attribute_definition_real_ref_1_)
					pastATTRIBUTE_VALUE_REALOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_REALOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_REAL](renamed_attribute_definition_real_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_REALOwner := _attribute_value_real // we have a match
						if pastATTRIBUTE_VALUE_REALOwner != nil {
							if newATTRIBUTE_VALUE_REALOwner != pastATTRIBUTE_VALUE_REALOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_REALOwner.DEFINITION, renamed_attribute_definition_real_ref_1_)
								pastATTRIBUTE_VALUE_REALOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_REALOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_REALOwner.DEFINITION = append(newATTRIBUTE_VALUE_REALOwner.DEFINITION, renamed_attribute_definition_real_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_REALOwner.DEFINITION = append(newATTRIBUTE_VALUE_REALOwner.DEFINITION, renamed_attribute_definition_real_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_real_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_real_ref_1_.Unstage(renamed_attribute_definition_real_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_real_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1](
		renamed_attribute_definition_real_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_real_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_real_ref_1FormCallback.CreationMode || renamed_attribute_definition_real_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_real_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_real_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1FormCallback(
			nil,
			renamed_attribute_definition_real_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_real_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_REAL_REF_1)
		FillUpForm(renamed_attribute_definition_real_ref_1, newFormGroup, renamed_attribute_definition_real_ref_1FormCallback.probe)
		renamed_attribute_definition_real_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_real_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback(
	renamed_attribute_definition_string_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_string_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback) {
	renamed_attribute_definition_string_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback)
	renamed_attribute_definition_string_ref_1FormCallback.probe = probe
	renamed_attribute_definition_string_ref_1FormCallback.renamed_attribute_definition_string_ref_1 = renamed_attribute_definition_string_ref_1
	renamed_attribute_definition_string_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_string_ref_1FormCallback.CreationMode = (renamed_attribute_definition_string_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback struct {
	renamed_attribute_definition_string_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_string_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_string_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_string_ref_1FormCallback.renamed_attribute_definition_string_ref_1 == nil {
		renamed_attribute_definition_string_ref_1FormCallback.renamed_attribute_definition_string_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1).Stage(renamed_attribute_definition_string_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_string_ref_1_ := renamed_attribute_definition_string_ref_1FormCallback.renamed_attribute_definition_string_ref_1
	_ = renamed_attribute_definition_string_ref_1_

	for _, formDiv := range renamed_attribute_definition_string_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_string_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_string_ref_1_.ATTRIBUTE_DEFINITION_STRING_REF), formDiv)
		case "ATTRIBUTE_VALUE_STRING:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_STRINGOwner *models.ATTRIBUTE_VALUE_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_STRING"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_string_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_string_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_string_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION, renamed_attribute_definition_string_ref_1_)
					pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_STRING](renamed_attribute_definition_string_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_STRINGOwner := _attribute_value_string // we have a match
						if pastATTRIBUTE_VALUE_STRINGOwner != nil {
							if newATTRIBUTE_VALUE_STRINGOwner != pastATTRIBUTE_VALUE_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION, renamed_attribute_definition_string_ref_1_)
								pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_STRINGOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_STRINGOwner.DEFINITION = append(newATTRIBUTE_VALUE_STRINGOwner.DEFINITION, renamed_attribute_definition_string_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_STRINGOwner.DEFINITION = append(newATTRIBUTE_VALUE_STRINGOwner.DEFINITION, renamed_attribute_definition_string_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_string_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_string_ref_1_.Unstage(renamed_attribute_definition_string_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_string_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1](
		renamed_attribute_definition_string_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_string_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_string_ref_1FormCallback.CreationMode || renamed_attribute_definition_string_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_string_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_string_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1FormCallback(
			nil,
			renamed_attribute_definition_string_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_string_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_STRING_REF_1)
		FillUpForm(renamed_attribute_definition_string_ref_1, newFormGroup, renamed_attribute_definition_string_ref_1FormCallback.probe)
		renamed_attribute_definition_string_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_string_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback(
	renamed_attribute_definition_xhtml_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_definition_xhtml_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback) {
	renamed_attribute_definition_xhtml_ref_1FormCallback = new(Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback)
	renamed_attribute_definition_xhtml_ref_1FormCallback.probe = probe
	renamed_attribute_definition_xhtml_ref_1FormCallback.renamed_attribute_definition_xhtml_ref_1 = renamed_attribute_definition_xhtml_ref_1
	renamed_attribute_definition_xhtml_ref_1FormCallback.formGroup = formGroup

	renamed_attribute_definition_xhtml_ref_1FormCallback.CreationMode = (renamed_attribute_definition_xhtml_ref_1 == nil)

	return
}

type Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback struct {
	renamed_attribute_definition_xhtml_ref_1 *models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_definition_xhtml_ref_1FormCallback *Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_definition_xhtml_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_definition_xhtml_ref_1FormCallback.renamed_attribute_definition_xhtml_ref_1 == nil {
		renamed_attribute_definition_xhtml_ref_1FormCallback.renamed_attribute_definition_xhtml_ref_1 = new(models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1).Stage(renamed_attribute_definition_xhtml_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_definition_xhtml_ref_1_ := renamed_attribute_definition_xhtml_ref_1FormCallback.renamed_attribute_definition_xhtml_ref_1
	_ = renamed_attribute_definition_xhtml_ref_1_

	for _, formDiv := range renamed_attribute_definition_xhtml_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_definition_xhtml_ref_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(renamed_attribute_definition_xhtml_ref_1_.ATTRIBUTE_DEFINITION_XHTML_REF), formDiv)
		case "ATTRIBUTE_VALUE_XHTML:DEFINITION":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_VALUE_XHTMLOwner *models.ATTRIBUTE_VALUE_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
			rf.Fieldname = "DEFINITION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_definition_xhtml_ref_1FormCallback.probe.stageOfInterest,
				renamed_attribute_definition_xhtml_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_definition_xhtml_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_VALUE_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_VALUE_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, renamed_attribute_definition_xhtml_ref_1_)
					pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_value_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](renamed_attribute_definition_xhtml_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_value_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_VALUE_XHTMLOwner := _attribute_value_xhtml // we have a match
						if pastATTRIBUTE_VALUE_XHTMLOwner != nil {
							if newATTRIBUTE_VALUE_XHTMLOwner != pastATTRIBUTE_VALUE_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, renamed_attribute_definition_xhtml_ref_1_)
								pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION = slices.Delete(pastATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, idx, idx+1)
								newATTRIBUTE_VALUE_XHTMLOwner.DEFINITION = append(newATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, renamed_attribute_definition_xhtml_ref_1_)
							}
						} else {
							newATTRIBUTE_VALUE_XHTMLOwner.DEFINITION = append(newATTRIBUTE_VALUE_XHTMLOwner.DEFINITION, renamed_attribute_definition_xhtml_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_definition_xhtml_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_xhtml_ref_1_.Unstage(renamed_attribute_definition_xhtml_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_definition_xhtml_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1](
		renamed_attribute_definition_xhtml_ref_1FormCallback.probe,
	)
	renamed_attribute_definition_xhtml_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_definition_xhtml_ref_1FormCallback.CreationMode || renamed_attribute_definition_xhtml_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_definition_xhtml_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_definition_xhtml_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1FormCallback(
			nil,
			renamed_attribute_definition_xhtml_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_definition_xhtml_ref_1 := new(models.Renamed_ATTRIBUTE_DEFINITION_XHTML_REF_1)
		FillUpForm(renamed_attribute_definition_xhtml_ref_1, newFormGroup, renamed_attribute_definition_xhtml_ref_1FormCallback.probe)
		renamed_attribute_definition_xhtml_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_definition_xhtml_ref_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback(
	renamed_attribute_value_boolean_1 *models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_boolean_1FormCallback *Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback) {
	renamed_attribute_value_boolean_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback)
	renamed_attribute_value_boolean_1FormCallback.probe = probe
	renamed_attribute_value_boolean_1FormCallback.renamed_attribute_value_boolean_1 = renamed_attribute_value_boolean_1
	renamed_attribute_value_boolean_1FormCallback.formGroup = formGroup

	renamed_attribute_value_boolean_1FormCallback.CreationMode = (renamed_attribute_value_boolean_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback struct {
	renamed_attribute_value_boolean_1 *models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_boolean_1FormCallback *Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_boolean_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_boolean_1FormCallback.renamed_attribute_value_boolean_1 == nil {
		renamed_attribute_value_boolean_1FormCallback.renamed_attribute_value_boolean_1 = new(models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1).Stage(renamed_attribute_value_boolean_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_boolean_1_ := renamed_attribute_value_boolean_1FormCallback.renamed_attribute_value_boolean_1
	_ = renamed_attribute_value_boolean_1_

	for _, formDiv := range renamed_attribute_value_boolean_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_boolean_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_BOOLEANOwner *models.ATTRIBUTE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_boolean_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_boolean_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_boolean_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, renamed_attribute_value_boolean_1_)
					pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](renamed_attribute_value_boolean_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_BOOLEANOwner := _attribute_definition_boolean // we have a match
						if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
							if newATTRIBUTE_DEFINITION_BOOLEANOwner != pastATTRIBUTE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, renamed_attribute_value_boolean_1_)
								pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, renamed_attribute_value_boolean_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.DEFAULT_VALUE, renamed_attribute_value_boolean_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_boolean_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_boolean_1_.Unstage(renamed_attribute_value_boolean_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_boolean_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1](
		renamed_attribute_value_boolean_1FormCallback.probe,
	)
	renamed_attribute_value_boolean_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_boolean_1FormCallback.CreationMode || renamed_attribute_value_boolean_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_boolean_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_boolean_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_BOOLEAN_1FormCallback(
			nil,
			renamed_attribute_value_boolean_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_boolean_1 := new(models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1)
		FillUpForm(renamed_attribute_value_boolean_1, newFormGroup, renamed_attribute_value_boolean_1FormCallback.probe)
		renamed_attribute_value_boolean_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_boolean_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback(
	renamed_attribute_value_date_1 *models.Renamed_ATTRIBUTE_VALUE_DATE_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_date_1FormCallback *Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback) {
	renamed_attribute_value_date_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback)
	renamed_attribute_value_date_1FormCallback.probe = probe
	renamed_attribute_value_date_1FormCallback.renamed_attribute_value_date_1 = renamed_attribute_value_date_1
	renamed_attribute_value_date_1FormCallback.formGroup = formGroup

	renamed_attribute_value_date_1FormCallback.CreationMode = (renamed_attribute_value_date_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback struct {
	renamed_attribute_value_date_1 *models.Renamed_ATTRIBUTE_VALUE_DATE_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_date_1FormCallback *Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_date_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_date_1FormCallback.renamed_attribute_value_date_1 == nil {
		renamed_attribute_value_date_1FormCallback.renamed_attribute_value_date_1 = new(models.Renamed_ATTRIBUTE_VALUE_DATE_1).Stage(renamed_attribute_value_date_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_date_1_ := renamed_attribute_value_date_1FormCallback.renamed_attribute_value_date_1
	_ = renamed_attribute_value_date_1_

	for _, formDiv := range renamed_attribute_value_date_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_date_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_DATEOwner *models.ATTRIBUTE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_date_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_date_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_date_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, renamed_attribute_value_date_1_)
					pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](renamed_attribute_value_date_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_DATEOwner := _attribute_definition_date // we have a match
						if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
							if newATTRIBUTE_DEFINITION_DATEOwner != pastATTRIBUTE_DEFINITION_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, renamed_attribute_value_date_1_)
								pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, renamed_attribute_value_date_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_DATEOwner.DEFAULT_VALUE, renamed_attribute_value_date_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_date_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_date_1_.Unstage(renamed_attribute_value_date_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_date_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_DATE_1](
		renamed_attribute_value_date_1FormCallback.probe,
	)
	renamed_attribute_value_date_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_date_1FormCallback.CreationMode || renamed_attribute_value_date_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_date_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_date_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_DATE_1FormCallback(
			nil,
			renamed_attribute_value_date_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_date_1 := new(models.Renamed_ATTRIBUTE_VALUE_DATE_1)
		FillUpForm(renamed_attribute_value_date_1, newFormGroup, renamed_attribute_value_date_1FormCallback.probe)
		renamed_attribute_value_date_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_date_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback(
	renamed_attribute_value_enumeration_1 *models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_enumeration_1FormCallback *Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback) {
	renamed_attribute_value_enumeration_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback)
	renamed_attribute_value_enumeration_1FormCallback.probe = probe
	renamed_attribute_value_enumeration_1FormCallback.renamed_attribute_value_enumeration_1 = renamed_attribute_value_enumeration_1
	renamed_attribute_value_enumeration_1FormCallback.formGroup = formGroup

	renamed_attribute_value_enumeration_1FormCallback.CreationMode = (renamed_attribute_value_enumeration_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback struct {
	renamed_attribute_value_enumeration_1 *models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_enumeration_1FormCallback *Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_enumeration_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_enumeration_1FormCallback.renamed_attribute_value_enumeration_1 == nil {
		renamed_attribute_value_enumeration_1FormCallback.renamed_attribute_value_enumeration_1 = new(models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1).Stage(renamed_attribute_value_enumeration_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_enumeration_1_ := renamed_attribute_value_enumeration_1FormCallback.renamed_attribute_value_enumeration_1
	_ = renamed_attribute_value_enumeration_1_

	for _, formDiv := range renamed_attribute_value_enumeration_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_enumeration_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_ENUMERATIONOwner *models.ATTRIBUTE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_enumeration_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_enumeration_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_enumeration_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, renamed_attribute_value_enumeration_1_)
					pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](renamed_attribute_value_enumeration_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_ENUMERATIONOwner := _attribute_definition_enumeration // we have a match
						if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
							if newATTRIBUTE_DEFINITION_ENUMERATIONOwner != pastATTRIBUTE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, renamed_attribute_value_enumeration_1_)
								pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, renamed_attribute_value_enumeration_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.DEFAULT_VALUE, renamed_attribute_value_enumeration_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_enumeration_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_enumeration_1_.Unstage(renamed_attribute_value_enumeration_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_enumeration_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1](
		renamed_attribute_value_enumeration_1FormCallback.probe,
	)
	renamed_attribute_value_enumeration_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_enumeration_1FormCallback.CreationMode || renamed_attribute_value_enumeration_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_enumeration_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_enumeration_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_ENUMERATION_1FormCallback(
			nil,
			renamed_attribute_value_enumeration_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_enumeration_1 := new(models.Renamed_ATTRIBUTE_VALUE_ENUMERATION_1)
		FillUpForm(renamed_attribute_value_enumeration_1, newFormGroup, renamed_attribute_value_enumeration_1FormCallback.probe)
		renamed_attribute_value_enumeration_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_enumeration_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback(
	renamed_attribute_value_integer_1 *models.Renamed_ATTRIBUTE_VALUE_INTEGER_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_integer_1FormCallback *Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback) {
	renamed_attribute_value_integer_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback)
	renamed_attribute_value_integer_1FormCallback.probe = probe
	renamed_attribute_value_integer_1FormCallback.renamed_attribute_value_integer_1 = renamed_attribute_value_integer_1
	renamed_attribute_value_integer_1FormCallback.formGroup = formGroup

	renamed_attribute_value_integer_1FormCallback.CreationMode = (renamed_attribute_value_integer_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback struct {
	renamed_attribute_value_integer_1 *models.Renamed_ATTRIBUTE_VALUE_INTEGER_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_integer_1FormCallback *Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_integer_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_integer_1FormCallback.renamed_attribute_value_integer_1 == nil {
		renamed_attribute_value_integer_1FormCallback.renamed_attribute_value_integer_1 = new(models.Renamed_ATTRIBUTE_VALUE_INTEGER_1).Stage(renamed_attribute_value_integer_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_integer_1_ := renamed_attribute_value_integer_1FormCallback.renamed_attribute_value_integer_1
	_ = renamed_attribute_value_integer_1_

	for _, formDiv := range renamed_attribute_value_integer_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_integer_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_INTEGEROwner *models.ATTRIBUTE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_integer_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_integer_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_integer_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, renamed_attribute_value_integer_1_)
					pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](renamed_attribute_value_integer_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_INTEGEROwner := _attribute_definition_integer // we have a match
						if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
							if newATTRIBUTE_DEFINITION_INTEGEROwner != pastATTRIBUTE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, renamed_attribute_value_integer_1_)
								pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, renamed_attribute_value_integer_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_INTEGEROwner.DEFAULT_VALUE, renamed_attribute_value_integer_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_integer_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_integer_1_.Unstage(renamed_attribute_value_integer_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_integer_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_INTEGER_1](
		renamed_attribute_value_integer_1FormCallback.probe,
	)
	renamed_attribute_value_integer_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_integer_1FormCallback.CreationMode || renamed_attribute_value_integer_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_integer_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_integer_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_INTEGER_1FormCallback(
			nil,
			renamed_attribute_value_integer_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_integer_1 := new(models.Renamed_ATTRIBUTE_VALUE_INTEGER_1)
		FillUpForm(renamed_attribute_value_integer_1, newFormGroup, renamed_attribute_value_integer_1FormCallback.probe)
		renamed_attribute_value_integer_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_integer_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback(
	renamed_attribute_value_real_1 *models.Renamed_ATTRIBUTE_VALUE_REAL_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_real_1FormCallback *Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback) {
	renamed_attribute_value_real_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback)
	renamed_attribute_value_real_1FormCallback.probe = probe
	renamed_attribute_value_real_1FormCallback.renamed_attribute_value_real_1 = renamed_attribute_value_real_1
	renamed_attribute_value_real_1FormCallback.formGroup = formGroup

	renamed_attribute_value_real_1FormCallback.CreationMode = (renamed_attribute_value_real_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback struct {
	renamed_attribute_value_real_1 *models.Renamed_ATTRIBUTE_VALUE_REAL_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_real_1FormCallback *Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_real_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_real_1FormCallback.renamed_attribute_value_real_1 == nil {
		renamed_attribute_value_real_1FormCallback.renamed_attribute_value_real_1 = new(models.Renamed_ATTRIBUTE_VALUE_REAL_1).Stage(renamed_attribute_value_real_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_real_1_ := renamed_attribute_value_real_1FormCallback.renamed_attribute_value_real_1
	_ = renamed_attribute_value_real_1_

	for _, formDiv := range renamed_attribute_value_real_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_real_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_REALOwner *models.ATTRIBUTE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_real_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_real_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_real_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, renamed_attribute_value_real_1_)
					pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](renamed_attribute_value_real_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_REALOwner := _attribute_definition_real // we have a match
						if pastATTRIBUTE_DEFINITION_REALOwner != nil {
							if newATTRIBUTE_DEFINITION_REALOwner != pastATTRIBUTE_DEFINITION_REALOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, renamed_attribute_value_real_1_)
								pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, renamed_attribute_value_real_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_REALOwner.DEFAULT_VALUE, renamed_attribute_value_real_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_real_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_real_1_.Unstage(renamed_attribute_value_real_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_real_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_REAL_1](
		renamed_attribute_value_real_1FormCallback.probe,
	)
	renamed_attribute_value_real_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_real_1FormCallback.CreationMode || renamed_attribute_value_real_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_real_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_real_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_REAL_1FormCallback(
			nil,
			renamed_attribute_value_real_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_real_1 := new(models.Renamed_ATTRIBUTE_VALUE_REAL_1)
		FillUpForm(renamed_attribute_value_real_1, newFormGroup, renamed_attribute_value_real_1FormCallback.probe)
		renamed_attribute_value_real_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_real_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback(
	renamed_attribute_value_string_1 *models.Renamed_ATTRIBUTE_VALUE_STRING_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_string_1FormCallback *Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback) {
	renamed_attribute_value_string_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback)
	renamed_attribute_value_string_1FormCallback.probe = probe
	renamed_attribute_value_string_1FormCallback.renamed_attribute_value_string_1 = renamed_attribute_value_string_1
	renamed_attribute_value_string_1FormCallback.formGroup = formGroup

	renamed_attribute_value_string_1FormCallback.CreationMode = (renamed_attribute_value_string_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback struct {
	renamed_attribute_value_string_1 *models.Renamed_ATTRIBUTE_VALUE_STRING_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_string_1FormCallback *Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_string_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_string_1FormCallback.renamed_attribute_value_string_1 == nil {
		renamed_attribute_value_string_1FormCallback.renamed_attribute_value_string_1 = new(models.Renamed_ATTRIBUTE_VALUE_STRING_1).Stage(renamed_attribute_value_string_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_string_1_ := renamed_attribute_value_string_1FormCallback.renamed_attribute_value_string_1
	_ = renamed_attribute_value_string_1_

	for _, formDiv := range renamed_attribute_value_string_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_string_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_STRINGOwner *models.ATTRIBUTE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_string_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_string_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_string_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, renamed_attribute_value_string_1_)
					pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](renamed_attribute_value_string_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_STRINGOwner := _attribute_definition_string // we have a match
						if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
							if newATTRIBUTE_DEFINITION_STRINGOwner != pastATTRIBUTE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, renamed_attribute_value_string_1_)
								pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, renamed_attribute_value_string_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_STRINGOwner.DEFAULT_VALUE, renamed_attribute_value_string_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_string_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_string_1_.Unstage(renamed_attribute_value_string_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_string_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_STRING_1](
		renamed_attribute_value_string_1FormCallback.probe,
	)
	renamed_attribute_value_string_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_string_1FormCallback.CreationMode || renamed_attribute_value_string_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_string_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_string_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_STRING_1FormCallback(
			nil,
			renamed_attribute_value_string_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_string_1 := new(models.Renamed_ATTRIBUTE_VALUE_STRING_1)
		FillUpForm(renamed_attribute_value_string_1, newFormGroup, renamed_attribute_value_string_1FormCallback.probe)
		renamed_attribute_value_string_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_string_1FormCallback.probe)
}
func __gong__New__Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback(
	renamed_attribute_value_xhtml_1 *models.Renamed_ATTRIBUTE_VALUE_XHTML_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_attribute_value_xhtml_1FormCallback *Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback) {
	renamed_attribute_value_xhtml_1FormCallback = new(Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback)
	renamed_attribute_value_xhtml_1FormCallback.probe = probe
	renamed_attribute_value_xhtml_1FormCallback.renamed_attribute_value_xhtml_1 = renamed_attribute_value_xhtml_1
	renamed_attribute_value_xhtml_1FormCallback.formGroup = formGroup

	renamed_attribute_value_xhtml_1FormCallback.CreationMode = (renamed_attribute_value_xhtml_1 == nil)

	return
}

type Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback struct {
	renamed_attribute_value_xhtml_1 *models.Renamed_ATTRIBUTE_VALUE_XHTML_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_attribute_value_xhtml_1FormCallback *Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback) OnSave() {

	log.Println("Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_attribute_value_xhtml_1FormCallback.probe.formStage.Checkout()

	if renamed_attribute_value_xhtml_1FormCallback.renamed_attribute_value_xhtml_1 == nil {
		renamed_attribute_value_xhtml_1FormCallback.renamed_attribute_value_xhtml_1 = new(models.Renamed_ATTRIBUTE_VALUE_XHTML_1).Stage(renamed_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
	}
	renamed_attribute_value_xhtml_1_ := renamed_attribute_value_xhtml_1FormCallback.renamed_attribute_value_xhtml_1
	_ = renamed_attribute_value_xhtml_1_

	for _, formDiv := range renamed_attribute_value_xhtml_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_attribute_value_xhtml_1_.Name), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML:DEFAULT_VALUE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_XHTMLOwner *models.ATTRIBUTE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "DEFAULT_VALUE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_attribute_value_xhtml_1FormCallback.probe.stageOfInterest,
				renamed_attribute_value_xhtml_1FormCallback.probe.backRepoOfInterest,
				renamed_attribute_value_xhtml_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, renamed_attribute_value_xhtml_1_)
					pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](renamed_attribute_value_xhtml_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_XHTMLOwner := _attribute_definition_xhtml // we have a match
						if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
							if newATTRIBUTE_DEFINITION_XHTMLOwner != pastATTRIBUTE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, renamed_attribute_value_xhtml_1_)
								pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, idx, idx+1)
								newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, renamed_attribute_value_xhtml_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE = append(newATTRIBUTE_DEFINITION_XHTMLOwner.DEFAULT_VALUE, renamed_attribute_value_xhtml_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_attribute_value_xhtml_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_xhtml_1_.Unstage(renamed_attribute_value_xhtml_1FormCallback.probe.stageOfInterest)
	}

	renamed_attribute_value_xhtml_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_ATTRIBUTE_VALUE_XHTML_1](
		renamed_attribute_value_xhtml_1FormCallback.probe,
	)
	renamed_attribute_value_xhtml_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_attribute_value_xhtml_1FormCallback.CreationMode || renamed_attribute_value_xhtml_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_attribute_value_xhtml_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_attribute_value_xhtml_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_ATTRIBUTE_VALUE_XHTML_1FormCallback(
			nil,
			renamed_attribute_value_xhtml_1FormCallback.probe,
			newFormGroup,
		)
		renamed_attribute_value_xhtml_1 := new(models.Renamed_ATTRIBUTE_VALUE_XHTML_1)
		FillUpForm(renamed_attribute_value_xhtml_1, newFormGroup, renamed_attribute_value_xhtml_1FormCallback.probe)
		renamed_attribute_value_xhtml_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_attribute_value_xhtml_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback(
	renamed_datatype_definition_boolean_ref_1 *models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_boolean_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback) {
	renamed_datatype_definition_boolean_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback)
	renamed_datatype_definition_boolean_ref_1FormCallback.probe = probe
	renamed_datatype_definition_boolean_ref_1FormCallback.renamed_datatype_definition_boolean_ref_1 = renamed_datatype_definition_boolean_ref_1
	renamed_datatype_definition_boolean_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_boolean_ref_1FormCallback.CreationMode = (renamed_datatype_definition_boolean_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback struct {
	renamed_datatype_definition_boolean_ref_1 *models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_boolean_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_boolean_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_boolean_ref_1FormCallback.renamed_datatype_definition_boolean_ref_1 == nil {
		renamed_datatype_definition_boolean_ref_1FormCallback.renamed_datatype_definition_boolean_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1).Stage(renamed_datatype_definition_boolean_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_boolean_ref_1_ := renamed_datatype_definition_boolean_ref_1FormCallback.renamed_datatype_definition_boolean_ref_1
	_ = renamed_datatype_definition_boolean_ref_1_

	for _, formDiv := range renamed_datatype_definition_boolean_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_boolean_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_boolean_ref_1_.DATATYPE_DEFINITION_BOOLEAN_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_BOOLEANOwner *models.ATTRIBUTE_DEFINITION_BOOLEAN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_BOOLEAN"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_boolean_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_boolean_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_boolean_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_BOOLEANOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, renamed_datatype_definition_boolean_ref_1_)
					pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_boolean := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](renamed_datatype_definition_boolean_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_boolean.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_BOOLEANOwner := _attribute_definition_boolean // we have a match
						if pastATTRIBUTE_DEFINITION_BOOLEANOwner != nil {
							if newATTRIBUTE_DEFINITION_BOOLEANOwner != pastATTRIBUTE_DEFINITION_BOOLEANOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, renamed_datatype_definition_boolean_ref_1_)
								pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, renamed_datatype_definition_boolean_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE = append(newATTRIBUTE_DEFINITION_BOOLEANOwner.TYPE, renamed_datatype_definition_boolean_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_boolean_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_boolean_ref_1_.Unstage(renamed_datatype_definition_boolean_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_boolean_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1](
		renamed_datatype_definition_boolean_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_boolean_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_boolean_ref_1FormCallback.CreationMode || renamed_datatype_definition_boolean_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_boolean_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_boolean_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1FormCallback(
			nil,
			renamed_datatype_definition_boolean_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_boolean_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_BOOLEAN_REF_1)
		FillUpForm(renamed_datatype_definition_boolean_ref_1, newFormGroup, renamed_datatype_definition_boolean_ref_1FormCallback.probe)
		renamed_datatype_definition_boolean_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_boolean_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback(
	renamed_datatype_definition_date_ref_1 *models.Renamed_DATATYPE_DEFINITION_DATE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_date_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback) {
	renamed_datatype_definition_date_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback)
	renamed_datatype_definition_date_ref_1FormCallback.probe = probe
	renamed_datatype_definition_date_ref_1FormCallback.renamed_datatype_definition_date_ref_1 = renamed_datatype_definition_date_ref_1
	renamed_datatype_definition_date_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_date_ref_1FormCallback.CreationMode = (renamed_datatype_definition_date_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback struct {
	renamed_datatype_definition_date_ref_1 *models.Renamed_DATATYPE_DEFINITION_DATE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_date_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_date_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_date_ref_1FormCallback.renamed_datatype_definition_date_ref_1 == nil {
		renamed_datatype_definition_date_ref_1FormCallback.renamed_datatype_definition_date_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_DATE_REF_1).Stage(renamed_datatype_definition_date_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_date_ref_1_ := renamed_datatype_definition_date_ref_1FormCallback.renamed_datatype_definition_date_ref_1
	_ = renamed_datatype_definition_date_ref_1_

	for _, formDiv := range renamed_datatype_definition_date_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_date_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_date_ref_1_.DATATYPE_DEFINITION_DATE_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_DATE:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_DATEOwner *models.ATTRIBUTE_DEFINITION_DATE
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_DATE"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_date_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_date_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_date_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_DATEOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_DATE)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.TYPE, renamed_datatype_definition_date_ref_1_)
					pastATTRIBUTE_DEFINITION_DATEOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_date := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](renamed_datatype_definition_date_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_date.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_DATEOwner := _attribute_definition_date // we have a match
						if pastATTRIBUTE_DEFINITION_DATEOwner != nil {
							if newATTRIBUTE_DEFINITION_DATEOwner != pastATTRIBUTE_DEFINITION_DATEOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_DATEOwner.TYPE, renamed_datatype_definition_date_ref_1_)
								pastATTRIBUTE_DEFINITION_DATEOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_DATEOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_DATEOwner.TYPE = append(newATTRIBUTE_DEFINITION_DATEOwner.TYPE, renamed_datatype_definition_date_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_DATEOwner.TYPE = append(newATTRIBUTE_DEFINITION_DATEOwner.TYPE, renamed_datatype_definition_date_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_date_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_date_ref_1_.Unstage(renamed_datatype_definition_date_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_date_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_DATE_REF_1](
		renamed_datatype_definition_date_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_date_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_date_ref_1FormCallback.CreationMode || renamed_datatype_definition_date_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_date_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_date_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_DATE_REF_1FormCallback(
			nil,
			renamed_datatype_definition_date_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_date_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_DATE_REF_1)
		FillUpForm(renamed_datatype_definition_date_ref_1, newFormGroup, renamed_datatype_definition_date_ref_1FormCallback.probe)
		renamed_datatype_definition_date_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_date_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback(
	renamed_datatype_definition_enumeration_ref_1 *models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_enumeration_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback) {
	renamed_datatype_definition_enumeration_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback)
	renamed_datatype_definition_enumeration_ref_1FormCallback.probe = probe
	renamed_datatype_definition_enumeration_ref_1FormCallback.renamed_datatype_definition_enumeration_ref_1 = renamed_datatype_definition_enumeration_ref_1
	renamed_datatype_definition_enumeration_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_enumeration_ref_1FormCallback.CreationMode = (renamed_datatype_definition_enumeration_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback struct {
	renamed_datatype_definition_enumeration_ref_1 *models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_enumeration_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_enumeration_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_enumeration_ref_1FormCallback.renamed_datatype_definition_enumeration_ref_1 == nil {
		renamed_datatype_definition_enumeration_ref_1FormCallback.renamed_datatype_definition_enumeration_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1).Stage(renamed_datatype_definition_enumeration_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_enumeration_ref_1_ := renamed_datatype_definition_enumeration_ref_1FormCallback.renamed_datatype_definition_enumeration_ref_1
	_ = renamed_datatype_definition_enumeration_ref_1_

	for _, formDiv := range renamed_datatype_definition_enumeration_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_enumeration_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_enumeration_ref_1_.DATATYPE_DEFINITION_ENUMERATION_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_ENUMERATIONOwner *models.ATTRIBUTE_DEFINITION_ENUMERATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_ENUMERATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_enumeration_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_enumeration_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_enumeration_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_ENUMERATIONOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, renamed_datatype_definition_enumeration_ref_1_)
					pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_enumeration := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](renamed_datatype_definition_enumeration_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_enumeration.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_ENUMERATIONOwner := _attribute_definition_enumeration // we have a match
						if pastATTRIBUTE_DEFINITION_ENUMERATIONOwner != nil {
							if newATTRIBUTE_DEFINITION_ENUMERATIONOwner != pastATTRIBUTE_DEFINITION_ENUMERATIONOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, renamed_datatype_definition_enumeration_ref_1_)
								pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, renamed_datatype_definition_enumeration_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE = append(newATTRIBUTE_DEFINITION_ENUMERATIONOwner.TYPE, renamed_datatype_definition_enumeration_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_enumeration_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_enumeration_ref_1_.Unstage(renamed_datatype_definition_enumeration_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_enumeration_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1](
		renamed_datatype_definition_enumeration_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_enumeration_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_enumeration_ref_1FormCallback.CreationMode || renamed_datatype_definition_enumeration_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_enumeration_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_enumeration_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1FormCallback(
			nil,
			renamed_datatype_definition_enumeration_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_enumeration_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_ENUMERATION_REF_1)
		FillUpForm(renamed_datatype_definition_enumeration_ref_1, newFormGroup, renamed_datatype_definition_enumeration_ref_1FormCallback.probe)
		renamed_datatype_definition_enumeration_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_enumeration_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback(
	renamed_datatype_definition_integer_ref_1 *models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_integer_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback) {
	renamed_datatype_definition_integer_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback)
	renamed_datatype_definition_integer_ref_1FormCallback.probe = probe
	renamed_datatype_definition_integer_ref_1FormCallback.renamed_datatype_definition_integer_ref_1 = renamed_datatype_definition_integer_ref_1
	renamed_datatype_definition_integer_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_integer_ref_1FormCallback.CreationMode = (renamed_datatype_definition_integer_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback struct {
	renamed_datatype_definition_integer_ref_1 *models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_integer_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_integer_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_integer_ref_1FormCallback.renamed_datatype_definition_integer_ref_1 == nil {
		renamed_datatype_definition_integer_ref_1FormCallback.renamed_datatype_definition_integer_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1).Stage(renamed_datatype_definition_integer_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_integer_ref_1_ := renamed_datatype_definition_integer_ref_1FormCallback.renamed_datatype_definition_integer_ref_1
	_ = renamed_datatype_definition_integer_ref_1_

	for _, formDiv := range renamed_datatype_definition_integer_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_integer_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_integer_ref_1_.DATATYPE_DEFINITION_INTEGER_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_INTEGEROwner *models.ATTRIBUTE_DEFINITION_INTEGER
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_INTEGER"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_integer_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_integer_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_integer_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_INTEGEROwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_INTEGER)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, renamed_datatype_definition_integer_ref_1_)
					pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_integer := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](renamed_datatype_definition_integer_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_integer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_INTEGEROwner := _attribute_definition_integer // we have a match
						if pastATTRIBUTE_DEFINITION_INTEGEROwner != nil {
							if newATTRIBUTE_DEFINITION_INTEGEROwner != pastATTRIBUTE_DEFINITION_INTEGEROwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, renamed_datatype_definition_integer_ref_1_)
								pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_INTEGEROwner.TYPE = append(newATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, renamed_datatype_definition_integer_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_INTEGEROwner.TYPE = append(newATTRIBUTE_DEFINITION_INTEGEROwner.TYPE, renamed_datatype_definition_integer_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_integer_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_integer_ref_1_.Unstage(renamed_datatype_definition_integer_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_integer_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1](
		renamed_datatype_definition_integer_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_integer_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_integer_ref_1FormCallback.CreationMode || renamed_datatype_definition_integer_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_integer_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_integer_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_INTEGER_REF_1FormCallback(
			nil,
			renamed_datatype_definition_integer_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_integer_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_INTEGER_REF_1)
		FillUpForm(renamed_datatype_definition_integer_ref_1, newFormGroup, renamed_datatype_definition_integer_ref_1FormCallback.probe)
		renamed_datatype_definition_integer_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_integer_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback(
	renamed_datatype_definition_real_ref_1 *models.Renamed_DATATYPE_DEFINITION_REAL_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_real_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback) {
	renamed_datatype_definition_real_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback)
	renamed_datatype_definition_real_ref_1FormCallback.probe = probe
	renamed_datatype_definition_real_ref_1FormCallback.renamed_datatype_definition_real_ref_1 = renamed_datatype_definition_real_ref_1
	renamed_datatype_definition_real_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_real_ref_1FormCallback.CreationMode = (renamed_datatype_definition_real_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback struct {
	renamed_datatype_definition_real_ref_1 *models.Renamed_DATATYPE_DEFINITION_REAL_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_real_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_real_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_real_ref_1FormCallback.renamed_datatype_definition_real_ref_1 == nil {
		renamed_datatype_definition_real_ref_1FormCallback.renamed_datatype_definition_real_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_REAL_REF_1).Stage(renamed_datatype_definition_real_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_real_ref_1_ := renamed_datatype_definition_real_ref_1FormCallback.renamed_datatype_definition_real_ref_1
	_ = renamed_datatype_definition_real_ref_1_

	for _, formDiv := range renamed_datatype_definition_real_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_real_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_real_ref_1_.DATATYPE_DEFINITION_REAL_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_REAL:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_REALOwner *models.ATTRIBUTE_DEFINITION_REAL
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_REAL"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_real_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_real_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_real_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_REALOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_REAL)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_REALOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.TYPE, renamed_datatype_definition_real_ref_1_)
					pastATTRIBUTE_DEFINITION_REALOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_real := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](renamed_datatype_definition_real_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_real.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_REALOwner := _attribute_definition_real // we have a match
						if pastATTRIBUTE_DEFINITION_REALOwner != nil {
							if newATTRIBUTE_DEFINITION_REALOwner != pastATTRIBUTE_DEFINITION_REALOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_REALOwner.TYPE, renamed_datatype_definition_real_ref_1_)
								pastATTRIBUTE_DEFINITION_REALOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_REALOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_REALOwner.TYPE = append(newATTRIBUTE_DEFINITION_REALOwner.TYPE, renamed_datatype_definition_real_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_REALOwner.TYPE = append(newATTRIBUTE_DEFINITION_REALOwner.TYPE, renamed_datatype_definition_real_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_real_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_real_ref_1_.Unstage(renamed_datatype_definition_real_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_real_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_REAL_REF_1](
		renamed_datatype_definition_real_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_real_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_real_ref_1FormCallback.CreationMode || renamed_datatype_definition_real_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_real_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_real_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_REAL_REF_1FormCallback(
			nil,
			renamed_datatype_definition_real_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_real_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_REAL_REF_1)
		FillUpForm(renamed_datatype_definition_real_ref_1, newFormGroup, renamed_datatype_definition_real_ref_1FormCallback.probe)
		renamed_datatype_definition_real_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_real_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback(
	renamed_datatype_definition_string_ref_1 *models.Renamed_DATATYPE_DEFINITION_STRING_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_string_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback) {
	renamed_datatype_definition_string_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback)
	renamed_datatype_definition_string_ref_1FormCallback.probe = probe
	renamed_datatype_definition_string_ref_1FormCallback.renamed_datatype_definition_string_ref_1 = renamed_datatype_definition_string_ref_1
	renamed_datatype_definition_string_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_string_ref_1FormCallback.CreationMode = (renamed_datatype_definition_string_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback struct {
	renamed_datatype_definition_string_ref_1 *models.Renamed_DATATYPE_DEFINITION_STRING_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_string_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_string_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_string_ref_1FormCallback.renamed_datatype_definition_string_ref_1 == nil {
		renamed_datatype_definition_string_ref_1FormCallback.renamed_datatype_definition_string_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_STRING_REF_1).Stage(renamed_datatype_definition_string_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_string_ref_1_ := renamed_datatype_definition_string_ref_1FormCallback.renamed_datatype_definition_string_ref_1
	_ = renamed_datatype_definition_string_ref_1_

	for _, formDiv := range renamed_datatype_definition_string_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_string_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_string_ref_1_.DATATYPE_DEFINITION_STRING_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_STRING:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_STRINGOwner *models.ATTRIBUTE_DEFINITION_STRING
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_STRING"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_string_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_string_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_string_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_STRINGOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_STRING)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE, renamed_datatype_definition_string_ref_1_)
					pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_string := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](renamed_datatype_definition_string_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_string.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_STRINGOwner := _attribute_definition_string // we have a match
						if pastATTRIBUTE_DEFINITION_STRINGOwner != nil {
							if newATTRIBUTE_DEFINITION_STRINGOwner != pastATTRIBUTE_DEFINITION_STRINGOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE, renamed_datatype_definition_string_ref_1_)
								pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_STRINGOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_STRINGOwner.TYPE = append(newATTRIBUTE_DEFINITION_STRINGOwner.TYPE, renamed_datatype_definition_string_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_STRINGOwner.TYPE = append(newATTRIBUTE_DEFINITION_STRINGOwner.TYPE, renamed_datatype_definition_string_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_string_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_string_ref_1_.Unstage(renamed_datatype_definition_string_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_string_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_STRING_REF_1](
		renamed_datatype_definition_string_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_string_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_string_ref_1FormCallback.CreationMode || renamed_datatype_definition_string_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_string_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_string_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_STRING_REF_1FormCallback(
			nil,
			renamed_datatype_definition_string_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_string_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_STRING_REF_1)
		FillUpForm(renamed_datatype_definition_string_ref_1, newFormGroup, renamed_datatype_definition_string_ref_1FormCallback.probe)
		renamed_datatype_definition_string_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_string_ref_1FormCallback.probe)
}
func __gong__New__Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback(
	renamed_datatype_definition_xhtml_ref_1 *models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_datatype_definition_xhtml_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback) {
	renamed_datatype_definition_xhtml_ref_1FormCallback = new(Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback)
	renamed_datatype_definition_xhtml_ref_1FormCallback.probe = probe
	renamed_datatype_definition_xhtml_ref_1FormCallback.renamed_datatype_definition_xhtml_ref_1 = renamed_datatype_definition_xhtml_ref_1
	renamed_datatype_definition_xhtml_ref_1FormCallback.formGroup = formGroup

	renamed_datatype_definition_xhtml_ref_1FormCallback.CreationMode = (renamed_datatype_definition_xhtml_ref_1 == nil)

	return
}

type Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback struct {
	renamed_datatype_definition_xhtml_ref_1 *models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_datatype_definition_xhtml_ref_1FormCallback *Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback) OnSave() {

	log.Println("Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_datatype_definition_xhtml_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_datatype_definition_xhtml_ref_1FormCallback.renamed_datatype_definition_xhtml_ref_1 == nil {
		renamed_datatype_definition_xhtml_ref_1FormCallback.renamed_datatype_definition_xhtml_ref_1 = new(models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1).Stage(renamed_datatype_definition_xhtml_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_datatype_definition_xhtml_ref_1_ := renamed_datatype_definition_xhtml_ref_1FormCallback.renamed_datatype_definition_xhtml_ref_1
	_ = renamed_datatype_definition_xhtml_ref_1_

	for _, formDiv := range renamed_datatype_definition_xhtml_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_datatype_definition_xhtml_ref_1_.Name), formDiv)
		case "DATATYPE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&(renamed_datatype_definition_xhtml_ref_1_.DATATYPE_DEFINITION_XHTML_REF), formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML:TYPE":
			// we need to retrieve the field owner before the change
			var pastATTRIBUTE_DEFINITION_XHTMLOwner *models.ATTRIBUTE_DEFINITION_XHTML
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ATTRIBUTE_DEFINITION_XHTML"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_datatype_definition_xhtml_ref_1FormCallback.probe.stageOfInterest,
				renamed_datatype_definition_xhtml_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_datatype_definition_xhtml_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastATTRIBUTE_DEFINITION_XHTMLOwner = reverseFieldOwner.(*models.ATTRIBUTE_DEFINITION_XHTML)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
					idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, renamed_datatype_definition_xhtml_ref_1_)
					pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attribute_definition_xhtml := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](renamed_datatype_definition_xhtml_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attribute_definition_xhtml.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newATTRIBUTE_DEFINITION_XHTMLOwner := _attribute_definition_xhtml // we have a match
						if pastATTRIBUTE_DEFINITION_XHTMLOwner != nil {
							if newATTRIBUTE_DEFINITION_XHTMLOwner != pastATTRIBUTE_DEFINITION_XHTMLOwner {
								idx := slices.Index(pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, renamed_datatype_definition_xhtml_ref_1_)
								pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE = slices.Delete(pastATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, idx, idx+1)
								newATTRIBUTE_DEFINITION_XHTMLOwner.TYPE = append(newATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, renamed_datatype_definition_xhtml_ref_1_)
							}
						} else {
							newATTRIBUTE_DEFINITION_XHTMLOwner.TYPE = append(newATTRIBUTE_DEFINITION_XHTMLOwner.TYPE, renamed_datatype_definition_xhtml_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_datatype_definition_xhtml_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_xhtml_ref_1_.Unstage(renamed_datatype_definition_xhtml_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_datatype_definition_xhtml_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1](
		renamed_datatype_definition_xhtml_ref_1FormCallback.probe,
	)
	renamed_datatype_definition_xhtml_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_datatype_definition_xhtml_ref_1FormCallback.CreationMode || renamed_datatype_definition_xhtml_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_datatype_definition_xhtml_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_datatype_definition_xhtml_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_DATATYPE_DEFINITION_XHTML_REF_1FormCallback(
			nil,
			renamed_datatype_definition_xhtml_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_datatype_definition_xhtml_ref_1 := new(models.Renamed_DATATYPE_DEFINITION_XHTML_REF_1)
		FillUpForm(renamed_datatype_definition_xhtml_ref_1, newFormGroup, renamed_datatype_definition_xhtml_ref_1FormCallback.probe)
		renamed_datatype_definition_xhtml_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_datatype_definition_xhtml_ref_1FormCallback.probe)
}
func __gong__New__Renamed_RELATION_GROUP_TYPE_REF_1FormCallback(
	renamed_relation_group_type_ref_1 *models.Renamed_RELATION_GROUP_TYPE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_relation_group_type_ref_1FormCallback *Renamed_RELATION_GROUP_TYPE_REF_1FormCallback) {
	renamed_relation_group_type_ref_1FormCallback = new(Renamed_RELATION_GROUP_TYPE_REF_1FormCallback)
	renamed_relation_group_type_ref_1FormCallback.probe = probe
	renamed_relation_group_type_ref_1FormCallback.renamed_relation_group_type_ref_1 = renamed_relation_group_type_ref_1
	renamed_relation_group_type_ref_1FormCallback.formGroup = formGroup

	renamed_relation_group_type_ref_1FormCallback.CreationMode = (renamed_relation_group_type_ref_1 == nil)

	return
}

type Renamed_RELATION_GROUP_TYPE_REF_1FormCallback struct {
	renamed_relation_group_type_ref_1 *models.Renamed_RELATION_GROUP_TYPE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_relation_group_type_ref_1FormCallback *Renamed_RELATION_GROUP_TYPE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_RELATION_GROUP_TYPE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_relation_group_type_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_relation_group_type_ref_1FormCallback.renamed_relation_group_type_ref_1 == nil {
		renamed_relation_group_type_ref_1FormCallback.renamed_relation_group_type_ref_1 = new(models.Renamed_RELATION_GROUP_TYPE_REF_1).Stage(renamed_relation_group_type_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_relation_group_type_ref_1_ := renamed_relation_group_type_ref_1FormCallback.renamed_relation_group_type_ref_1
	_ = renamed_relation_group_type_ref_1_

	for _, formDiv := range renamed_relation_group_type_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_relation_group_type_ref_1_.Name), formDiv)
		case "RELATION_GROUP_TYPE_REF":
			FormDivBasicFieldToField(&(renamed_relation_group_type_ref_1_.RELATION_GROUP_TYPE_REF), formDiv)
		case "RELATION_GROUP:TYPE":
			// we need to retrieve the field owner before the change
			var pastRELATION_GROUPOwner *models.RELATION_GROUP
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RELATION_GROUP"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_relation_group_type_ref_1FormCallback.probe.stageOfInterest,
				renamed_relation_group_type_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_relation_group_type_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastRELATION_GROUPOwner = reverseFieldOwner.(*models.RELATION_GROUP)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRELATION_GROUPOwner != nil {
					idx := slices.Index(pastRELATION_GROUPOwner.TYPE, renamed_relation_group_type_ref_1_)
					pastRELATION_GROUPOwner.TYPE = slices.Delete(pastRELATION_GROUPOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _relation_group := range *models.GetGongstructInstancesSet[models.RELATION_GROUP](renamed_relation_group_type_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _relation_group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRELATION_GROUPOwner := _relation_group // we have a match
						if pastRELATION_GROUPOwner != nil {
							if newRELATION_GROUPOwner != pastRELATION_GROUPOwner {
								idx := slices.Index(pastRELATION_GROUPOwner.TYPE, renamed_relation_group_type_ref_1_)
								pastRELATION_GROUPOwner.TYPE = slices.Delete(pastRELATION_GROUPOwner.TYPE, idx, idx+1)
								newRELATION_GROUPOwner.TYPE = append(newRELATION_GROUPOwner.TYPE, renamed_relation_group_type_ref_1_)
							}
						} else {
							newRELATION_GROUPOwner.TYPE = append(newRELATION_GROUPOwner.TYPE, renamed_relation_group_type_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_relation_group_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_relation_group_type_ref_1_.Unstage(renamed_relation_group_type_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_relation_group_type_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_RELATION_GROUP_TYPE_REF_1](
		renamed_relation_group_type_ref_1FormCallback.probe,
	)
	renamed_relation_group_type_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_relation_group_type_ref_1FormCallback.CreationMode || renamed_relation_group_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_relation_group_type_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_relation_group_type_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_RELATION_GROUP_TYPE_REF_1FormCallback(
			nil,
			renamed_relation_group_type_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_relation_group_type_ref_1 := new(models.Renamed_RELATION_GROUP_TYPE_REF_1)
		FillUpForm(renamed_relation_group_type_ref_1, newFormGroup, renamed_relation_group_type_ref_1FormCallback.probe)
		renamed_relation_group_type_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_relation_group_type_ref_1FormCallback.probe)
}
func __gong__New__Renamed_SPECIFICATION_TYPE_REF_1FormCallback(
	renamed_specification_type_ref_1 *models.Renamed_SPECIFICATION_TYPE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_specification_type_ref_1FormCallback *Renamed_SPECIFICATION_TYPE_REF_1FormCallback) {
	renamed_specification_type_ref_1FormCallback = new(Renamed_SPECIFICATION_TYPE_REF_1FormCallback)
	renamed_specification_type_ref_1FormCallback.probe = probe
	renamed_specification_type_ref_1FormCallback.renamed_specification_type_ref_1 = renamed_specification_type_ref_1
	renamed_specification_type_ref_1FormCallback.formGroup = formGroup

	renamed_specification_type_ref_1FormCallback.CreationMode = (renamed_specification_type_ref_1 == nil)

	return
}

type Renamed_SPECIFICATION_TYPE_REF_1FormCallback struct {
	renamed_specification_type_ref_1 *models.Renamed_SPECIFICATION_TYPE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_specification_type_ref_1FormCallback *Renamed_SPECIFICATION_TYPE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_SPECIFICATION_TYPE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_specification_type_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_specification_type_ref_1FormCallback.renamed_specification_type_ref_1 == nil {
		renamed_specification_type_ref_1FormCallback.renamed_specification_type_ref_1 = new(models.Renamed_SPECIFICATION_TYPE_REF_1).Stage(renamed_specification_type_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_specification_type_ref_1_ := renamed_specification_type_ref_1FormCallback.renamed_specification_type_ref_1
	_ = renamed_specification_type_ref_1_

	for _, formDiv := range renamed_specification_type_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_specification_type_ref_1_.Name), formDiv)
		case "SPECIFICATION_TYPE_REF":
			FormDivBasicFieldToField(&(renamed_specification_type_ref_1_.SPECIFICATION_TYPE_REF), formDiv)
		case "SPECIFICATION:TYPE":
			// we need to retrieve the field owner before the change
			var pastSPECIFICATIONOwner *models.SPECIFICATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPECIFICATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_specification_type_ref_1FormCallback.probe.stageOfInterest,
				renamed_specification_type_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_specification_type_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPECIFICATIONOwner = reverseFieldOwner.(*models.SPECIFICATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPECIFICATIONOwner != nil {
					idx := slices.Index(pastSPECIFICATIONOwner.TYPE, renamed_specification_type_ref_1_)
					pastSPECIFICATIONOwner.TYPE = slices.Delete(pastSPECIFICATIONOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _specification := range *models.GetGongstructInstancesSet[models.SPECIFICATION](renamed_specification_type_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _specification.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPECIFICATIONOwner := _specification // we have a match
						if pastSPECIFICATIONOwner != nil {
							if newSPECIFICATIONOwner != pastSPECIFICATIONOwner {
								idx := slices.Index(pastSPECIFICATIONOwner.TYPE, renamed_specification_type_ref_1_)
								pastSPECIFICATIONOwner.TYPE = slices.Delete(pastSPECIFICATIONOwner.TYPE, idx, idx+1)
								newSPECIFICATIONOwner.TYPE = append(newSPECIFICATIONOwner.TYPE, renamed_specification_type_ref_1_)
							}
						} else {
							newSPECIFICATIONOwner.TYPE = append(newSPECIFICATIONOwner.TYPE, renamed_specification_type_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_specification_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_specification_type_ref_1_.Unstage(renamed_specification_type_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_specification_type_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_SPECIFICATION_TYPE_REF_1](
		renamed_specification_type_ref_1FormCallback.probe,
	)
	renamed_specification_type_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_specification_type_ref_1FormCallback.CreationMode || renamed_specification_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_specification_type_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_specification_type_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_SPECIFICATION_TYPE_REF_1FormCallback(
			nil,
			renamed_specification_type_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_specification_type_ref_1 := new(models.Renamed_SPECIFICATION_TYPE_REF_1)
		FillUpForm(renamed_specification_type_ref_1, newFormGroup, renamed_specification_type_ref_1FormCallback.probe)
		renamed_specification_type_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_specification_type_ref_1FormCallback.probe)
}
func __gong__New__Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback(
	renamed_spec_object_type_ref_1 *models.Renamed_SPEC_OBJECT_TYPE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_spec_object_type_ref_1FormCallback *Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback) {
	renamed_spec_object_type_ref_1FormCallback = new(Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback)
	renamed_spec_object_type_ref_1FormCallback.probe = probe
	renamed_spec_object_type_ref_1FormCallback.renamed_spec_object_type_ref_1 = renamed_spec_object_type_ref_1
	renamed_spec_object_type_ref_1FormCallback.formGroup = formGroup

	renamed_spec_object_type_ref_1FormCallback.CreationMode = (renamed_spec_object_type_ref_1 == nil)

	return
}

type Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback struct {
	renamed_spec_object_type_ref_1 *models.Renamed_SPEC_OBJECT_TYPE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_spec_object_type_ref_1FormCallback *Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_spec_object_type_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_spec_object_type_ref_1FormCallback.renamed_spec_object_type_ref_1 == nil {
		renamed_spec_object_type_ref_1FormCallback.renamed_spec_object_type_ref_1 = new(models.Renamed_SPEC_OBJECT_TYPE_REF_1).Stage(renamed_spec_object_type_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_spec_object_type_ref_1_ := renamed_spec_object_type_ref_1FormCallback.renamed_spec_object_type_ref_1
	_ = renamed_spec_object_type_ref_1_

	for _, formDiv := range renamed_spec_object_type_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_spec_object_type_ref_1_.Name), formDiv)
		case "SPEC_OBJECT_TYPE_REF":
			FormDivBasicFieldToField(&(renamed_spec_object_type_ref_1_.SPEC_OBJECT_TYPE_REF), formDiv)
		case "SPEC_OBJECT:TYPE":
			// we need to retrieve the field owner before the change
			var pastSPEC_OBJECTOwner *models.SPEC_OBJECT
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_OBJECT"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_spec_object_type_ref_1FormCallback.probe.stageOfInterest,
				renamed_spec_object_type_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_spec_object_type_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_OBJECTOwner = reverseFieldOwner.(*models.SPEC_OBJECT)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_OBJECTOwner != nil {
					idx := slices.Index(pastSPEC_OBJECTOwner.TYPE, renamed_spec_object_type_ref_1_)
					pastSPEC_OBJECTOwner.TYPE = slices.Delete(pastSPEC_OBJECTOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_object := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](renamed_spec_object_type_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_object.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_OBJECTOwner := _spec_object // we have a match
						if pastSPEC_OBJECTOwner != nil {
							if newSPEC_OBJECTOwner != pastSPEC_OBJECTOwner {
								idx := slices.Index(pastSPEC_OBJECTOwner.TYPE, renamed_spec_object_type_ref_1_)
								pastSPEC_OBJECTOwner.TYPE = slices.Delete(pastSPEC_OBJECTOwner.TYPE, idx, idx+1)
								newSPEC_OBJECTOwner.TYPE = append(newSPEC_OBJECTOwner.TYPE, renamed_spec_object_type_ref_1_)
							}
						} else {
							newSPEC_OBJECTOwner.TYPE = append(newSPEC_OBJECTOwner.TYPE, renamed_spec_object_type_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_spec_object_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_spec_object_type_ref_1_.Unstage(renamed_spec_object_type_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_spec_object_type_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_SPEC_OBJECT_TYPE_REF_1](
		renamed_spec_object_type_ref_1FormCallback.probe,
	)
	renamed_spec_object_type_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_spec_object_type_ref_1FormCallback.CreationMode || renamed_spec_object_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_spec_object_type_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_spec_object_type_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_SPEC_OBJECT_TYPE_REF_1FormCallback(
			nil,
			renamed_spec_object_type_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_spec_object_type_ref_1 := new(models.Renamed_SPEC_OBJECT_TYPE_REF_1)
		FillUpForm(renamed_spec_object_type_ref_1, newFormGroup, renamed_spec_object_type_ref_1FormCallback.probe)
		renamed_spec_object_type_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_spec_object_type_ref_1FormCallback.probe)
}
func __gong__New__Renamed_SPEC_RELATION_TYPE_REF_1FormCallback(
	renamed_spec_relation_type_ref_1 *models.Renamed_SPEC_RELATION_TYPE_REF_1,
	probe *Probe,
	formGroup *table.FormGroup,
) (renamed_spec_relation_type_ref_1FormCallback *Renamed_SPEC_RELATION_TYPE_REF_1FormCallback) {
	renamed_spec_relation_type_ref_1FormCallback = new(Renamed_SPEC_RELATION_TYPE_REF_1FormCallback)
	renamed_spec_relation_type_ref_1FormCallback.probe = probe
	renamed_spec_relation_type_ref_1FormCallback.renamed_spec_relation_type_ref_1 = renamed_spec_relation_type_ref_1
	renamed_spec_relation_type_ref_1FormCallback.formGroup = formGroup

	renamed_spec_relation_type_ref_1FormCallback.CreationMode = (renamed_spec_relation_type_ref_1 == nil)

	return
}

type Renamed_SPEC_RELATION_TYPE_REF_1FormCallback struct {
	renamed_spec_relation_type_ref_1 *models.Renamed_SPEC_RELATION_TYPE_REF_1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (renamed_spec_relation_type_ref_1FormCallback *Renamed_SPEC_RELATION_TYPE_REF_1FormCallback) OnSave() {

	log.Println("Renamed_SPEC_RELATION_TYPE_REF_1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	renamed_spec_relation_type_ref_1FormCallback.probe.formStage.Checkout()

	if renamed_spec_relation_type_ref_1FormCallback.renamed_spec_relation_type_ref_1 == nil {
		renamed_spec_relation_type_ref_1FormCallback.renamed_spec_relation_type_ref_1 = new(models.Renamed_SPEC_RELATION_TYPE_REF_1).Stage(renamed_spec_relation_type_ref_1FormCallback.probe.stageOfInterest)
	}
	renamed_spec_relation_type_ref_1_ := renamed_spec_relation_type_ref_1FormCallback.renamed_spec_relation_type_ref_1
	_ = renamed_spec_relation_type_ref_1_

	for _, formDiv := range renamed_spec_relation_type_ref_1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(renamed_spec_relation_type_ref_1_.Name), formDiv)
		case "SPEC_RELATION_TYPE_REF":
			FormDivBasicFieldToField(&(renamed_spec_relation_type_ref_1_.SPEC_RELATION_TYPE_REF), formDiv)
		case "SPEC_RELATION:TYPE":
			// we need to retrieve the field owner before the change
			var pastSPEC_RELATIONOwner *models.SPEC_RELATION
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SPEC_RELATION"
			rf.Fieldname = "TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				renamed_spec_relation_type_ref_1FormCallback.probe.stageOfInterest,
				renamed_spec_relation_type_ref_1FormCallback.probe.backRepoOfInterest,
				renamed_spec_relation_type_ref_1_,
				&rf)

			if reverseFieldOwner != nil {
				pastSPEC_RELATIONOwner = reverseFieldOwner.(*models.SPEC_RELATION)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSPEC_RELATIONOwner != nil {
					idx := slices.Index(pastSPEC_RELATIONOwner.TYPE, renamed_spec_relation_type_ref_1_)
					pastSPEC_RELATIONOwner.TYPE = slices.Delete(pastSPEC_RELATIONOwner.TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _spec_relation := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](renamed_spec_relation_type_ref_1FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _spec_relation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSPEC_RELATIONOwner := _spec_relation // we have a match
						if pastSPEC_RELATIONOwner != nil {
							if newSPEC_RELATIONOwner != pastSPEC_RELATIONOwner {
								idx := slices.Index(pastSPEC_RELATIONOwner.TYPE, renamed_spec_relation_type_ref_1_)
								pastSPEC_RELATIONOwner.TYPE = slices.Delete(pastSPEC_RELATIONOwner.TYPE, idx, idx+1)
								newSPEC_RELATIONOwner.TYPE = append(newSPEC_RELATIONOwner.TYPE, renamed_spec_relation_type_ref_1_)
							}
						} else {
							newSPEC_RELATIONOwner.TYPE = append(newSPEC_RELATIONOwner.TYPE, renamed_spec_relation_type_ref_1_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if renamed_spec_relation_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_spec_relation_type_ref_1_.Unstage(renamed_spec_relation_type_ref_1FormCallback.probe.stageOfInterest)
	}

	renamed_spec_relation_type_ref_1FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Renamed_SPEC_RELATION_TYPE_REF_1](
		renamed_spec_relation_type_ref_1FormCallback.probe,
	)
	renamed_spec_relation_type_ref_1FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if renamed_spec_relation_type_ref_1FormCallback.CreationMode || renamed_spec_relation_type_ref_1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		renamed_spec_relation_type_ref_1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(renamed_spec_relation_type_ref_1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Renamed_SPEC_RELATION_TYPE_REF_1FormCallback(
			nil,
			renamed_spec_relation_type_ref_1FormCallback.probe,
			newFormGroup,
		)
		renamed_spec_relation_type_ref_1 := new(models.Renamed_SPEC_RELATION_TYPE_REF_1)
		FillUpForm(renamed_spec_relation_type_ref_1, newFormGroup, renamed_spec_relation_type_ref_1FormCallback.probe)
		renamed_spec_relation_type_ref_1FormCallback.probe.formStage.Commit()
	}

	fillUpTree(renamed_spec_relation_type_ref_1FormCallback.probe)
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
		case "A_SPECIFICATIONS:SPECIFICATION":
			// we need to retrieve the field owner before the change
			var pastA_SPECIFICATIONSOwner *models.A_SPECIFICATIONS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPECIFICATIONS"
			rf.Fieldname = "SPECIFICATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specificationFormCallback.probe.stageOfInterest,
				specificationFormCallback.probe.backRepoOfInterest,
				specification_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPECIFICATIONSOwner = reverseFieldOwner.(*models.A_SPECIFICATIONS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPECIFICATIONSOwner != nil {
					idx := slices.Index(pastA_SPECIFICATIONSOwner.SPECIFICATION, specification_)
					pastA_SPECIFICATIONSOwner.SPECIFICATION = slices.Delete(pastA_SPECIFICATIONSOwner.SPECIFICATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_specifications := range *models.GetGongstructInstancesSet[models.A_SPECIFICATIONS](specificationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_specifications.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPECIFICATIONSOwner := _a_specifications // we have a match
						if pastA_SPECIFICATIONSOwner != nil {
							if newA_SPECIFICATIONSOwner != pastA_SPECIFICATIONSOwner {
								idx := slices.Index(pastA_SPECIFICATIONSOwner.SPECIFICATION, specification_)
								pastA_SPECIFICATIONSOwner.SPECIFICATION = slices.Delete(pastA_SPECIFICATIONSOwner.SPECIFICATION, idx, idx+1)
								newA_SPECIFICATIONSOwner.SPECIFICATION = append(newA_SPECIFICATIONSOwner.SPECIFICATION, specification_)
							}
						} else {
							newA_SPECIFICATIONSOwner.SPECIFICATION = append(newA_SPECIFICATIONSOwner.SPECIFICATION, specification_)
						}
					}
				}
			}
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
		case "A_SPEC_TYPES:SPECIFICATION_TYPE":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_TYPESOwner *models.A_SPEC_TYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPECIFICATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				specification_typeFormCallback.probe.stageOfInterest,
				specification_typeFormCallback.probe.backRepoOfInterest,
				specification_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_TYPESOwner = reverseFieldOwner.(*models.A_SPEC_TYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_TYPESOwner != nil {
					idx := slices.Index(pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE, specification_type_)
					pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](specification_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_types.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_TYPESOwner := _a_spec_types // we have a match
						if pastA_SPEC_TYPESOwner != nil {
							if newA_SPEC_TYPESOwner != pastA_SPEC_TYPESOwner {
								idx := slices.Index(pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE, specification_type_)
								pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPECIFICATION_TYPE, idx, idx+1)
								newA_SPEC_TYPESOwner.SPECIFICATION_TYPE = append(newA_SPEC_TYPESOwner.SPECIFICATION_TYPE, specification_type_)
							}
						} else {
							newA_SPEC_TYPESOwner.SPECIFICATION_TYPE = append(newA_SPEC_TYPESOwner.SPECIFICATION_TYPE, specification_type_)
						}
					}
				}
			}
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
		case "A_CHILDREN:SPEC_HIERARCHY":
			// we need to retrieve the field owner before the change
			var pastA_CHILDRENOwner *models.A_CHILDREN
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_CHILDREN"
			rf.Fieldname = "SPEC_HIERARCHY"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_hierarchyFormCallback.probe.stageOfInterest,
				spec_hierarchyFormCallback.probe.backRepoOfInterest,
				spec_hierarchy_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_CHILDRENOwner = reverseFieldOwner.(*models.A_CHILDREN)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_CHILDRENOwner != nil {
					idx := slices.Index(pastA_CHILDRENOwner.SPEC_HIERARCHY, spec_hierarchy_)
					pastA_CHILDRENOwner.SPEC_HIERARCHY = slices.Delete(pastA_CHILDRENOwner.SPEC_HIERARCHY, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_children := range *models.GetGongstructInstancesSet[models.A_CHILDREN](spec_hierarchyFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_children.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_CHILDRENOwner := _a_children // we have a match
						if pastA_CHILDRENOwner != nil {
							if newA_CHILDRENOwner != pastA_CHILDRENOwner {
								idx := slices.Index(pastA_CHILDRENOwner.SPEC_HIERARCHY, spec_hierarchy_)
								pastA_CHILDRENOwner.SPEC_HIERARCHY = slices.Delete(pastA_CHILDRENOwner.SPEC_HIERARCHY, idx, idx+1)
								newA_CHILDRENOwner.SPEC_HIERARCHY = append(newA_CHILDRENOwner.SPEC_HIERARCHY, spec_hierarchy_)
							}
						} else {
							newA_CHILDRENOwner.SPEC_HIERARCHY = append(newA_CHILDRENOwner.SPEC_HIERARCHY, spec_hierarchy_)
						}
					}
				}
			}
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
		case "A_SPEC_OBJECTS:SPEC_OBJECT":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_OBJECTSOwner *models.A_SPEC_OBJECTS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_OBJECTS"
			rf.Fieldname = "SPEC_OBJECT"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_objectFormCallback.probe.stageOfInterest,
				spec_objectFormCallback.probe.backRepoOfInterest,
				spec_object_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_OBJECTSOwner = reverseFieldOwner.(*models.A_SPEC_OBJECTS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_OBJECTSOwner != nil {
					idx := slices.Index(pastA_SPEC_OBJECTSOwner.SPEC_OBJECT, spec_object_)
					pastA_SPEC_OBJECTSOwner.SPEC_OBJECT = slices.Delete(pastA_SPEC_OBJECTSOwner.SPEC_OBJECT, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_objects := range *models.GetGongstructInstancesSet[models.A_SPEC_OBJECTS](spec_objectFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_objects.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_OBJECTSOwner := _a_spec_objects // we have a match
						if pastA_SPEC_OBJECTSOwner != nil {
							if newA_SPEC_OBJECTSOwner != pastA_SPEC_OBJECTSOwner {
								idx := slices.Index(pastA_SPEC_OBJECTSOwner.SPEC_OBJECT, spec_object_)
								pastA_SPEC_OBJECTSOwner.SPEC_OBJECT = slices.Delete(pastA_SPEC_OBJECTSOwner.SPEC_OBJECT, idx, idx+1)
								newA_SPEC_OBJECTSOwner.SPEC_OBJECT = append(newA_SPEC_OBJECTSOwner.SPEC_OBJECT, spec_object_)
							}
						} else {
							newA_SPEC_OBJECTSOwner.SPEC_OBJECT = append(newA_SPEC_OBJECTSOwner.SPEC_OBJECT, spec_object_)
						}
					}
				}
			}
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
		case "A_SPEC_TYPES:SPEC_OBJECT_TYPE":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_TYPESOwner *models.A_SPEC_TYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPEC_OBJECT_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_object_typeFormCallback.probe.stageOfInterest,
				spec_object_typeFormCallback.probe.backRepoOfInterest,
				spec_object_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_TYPESOwner = reverseFieldOwner.(*models.A_SPEC_TYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_TYPESOwner != nil {
					idx := slices.Index(pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, spec_object_type_)
					pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](spec_object_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_types.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_TYPESOwner := _a_spec_types // we have a match
						if pastA_SPEC_TYPESOwner != nil {
							if newA_SPEC_TYPESOwner != pastA_SPEC_TYPESOwner {
								idx := slices.Index(pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, spec_object_type_)
								pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, idx, idx+1)
								newA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE = append(newA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, spec_object_type_)
							}
						} else {
							newA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE = append(newA_SPEC_TYPESOwner.SPEC_OBJECT_TYPE, spec_object_type_)
						}
					}
				}
			}
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
		case "A_SPEC_RELATIONS:SPEC_RELATION":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_RELATIONSOwner *models.A_SPEC_RELATIONS
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_RELATIONS"
			rf.Fieldname = "SPEC_RELATION"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_relationFormCallback.probe.stageOfInterest,
				spec_relationFormCallback.probe.backRepoOfInterest,
				spec_relation_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_RELATIONSOwner = reverseFieldOwner.(*models.A_SPEC_RELATIONS)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_RELATIONSOwner != nil {
					idx := slices.Index(pastA_SPEC_RELATIONSOwner.SPEC_RELATION, spec_relation_)
					pastA_SPEC_RELATIONSOwner.SPEC_RELATION = slices.Delete(pastA_SPEC_RELATIONSOwner.SPEC_RELATION, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_relations := range *models.GetGongstructInstancesSet[models.A_SPEC_RELATIONS](spec_relationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_relations.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_RELATIONSOwner := _a_spec_relations // we have a match
						if pastA_SPEC_RELATIONSOwner != nil {
							if newA_SPEC_RELATIONSOwner != pastA_SPEC_RELATIONSOwner {
								idx := slices.Index(pastA_SPEC_RELATIONSOwner.SPEC_RELATION, spec_relation_)
								pastA_SPEC_RELATIONSOwner.SPEC_RELATION = slices.Delete(pastA_SPEC_RELATIONSOwner.SPEC_RELATION, idx, idx+1)
								newA_SPEC_RELATIONSOwner.SPEC_RELATION = append(newA_SPEC_RELATIONSOwner.SPEC_RELATION, spec_relation_)
							}
						} else {
							newA_SPEC_RELATIONSOwner.SPEC_RELATION = append(newA_SPEC_RELATIONSOwner.SPEC_RELATION, spec_relation_)
						}
					}
				}
			}
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
		case "A_SPEC_TYPES:SPEC_RELATION_TYPE":
			// we need to retrieve the field owner before the change
			var pastA_SPEC_TYPESOwner *models.A_SPEC_TYPES
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "A_SPEC_TYPES"
			rf.Fieldname = "SPEC_RELATION_TYPE"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				spec_relation_typeFormCallback.probe.stageOfInterest,
				spec_relation_typeFormCallback.probe.backRepoOfInterest,
				spec_relation_type_,
				&rf)

			if reverseFieldOwner != nil {
				pastA_SPEC_TYPESOwner = reverseFieldOwner.(*models.A_SPEC_TYPES)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastA_SPEC_TYPESOwner != nil {
					idx := slices.Index(pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, spec_relation_type_)
					pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _a_spec_types := range *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](spec_relation_typeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _a_spec_types.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newA_SPEC_TYPESOwner := _a_spec_types // we have a match
						if pastA_SPEC_TYPESOwner != nil {
							if newA_SPEC_TYPESOwner != pastA_SPEC_TYPESOwner {
								idx := slices.Index(pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, spec_relation_type_)
								pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE = slices.Delete(pastA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, idx, idx+1)
								newA_SPEC_TYPESOwner.SPEC_RELATION_TYPE = append(newA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, spec_relation_type_)
							}
						} else {
							newA_SPEC_TYPESOwner.SPEC_RELATION_TYPE = append(newA_SPEC_TYPESOwner.SPEC_RELATION_TYPE, spec_relation_type_)
						}
					}
				}
			}
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
