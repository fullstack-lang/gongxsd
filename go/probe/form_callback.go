// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AnnotationFormCallback(
	annotation *models.Annotation,
	probe *Probe,
	formGroup *table.FormGroup,
) (annotationFormCallback *AnnotationFormCallback) {
	annotationFormCallback = new(AnnotationFormCallback)
	annotationFormCallback.probe = probe
	annotationFormCallback.annotation = annotation
	annotationFormCallback.formGroup = formGroup

	annotationFormCallback.CreationMode = (annotation == nil)

	return
}

type AnnotationFormCallback struct {
	annotation *models.Annotation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (annotationFormCallback *AnnotationFormCallback) OnSave() {

	log.Println("AnnotationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	annotationFormCallback.probe.formStage.Checkout()

	if annotationFormCallback.annotation == nil {
		annotationFormCallback.annotation = new(models.Annotation).Stage(annotationFormCallback.probe.stageOfInterest)
	}
	annotation_ := annotationFormCallback.annotation
	_ = annotation_

	for _, formDiv := range annotationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(annotation_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotation_.Unstage(annotationFormCallback.probe.stageOfInterest)
	}

	annotationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Annotation](
		annotationFormCallback.probe,
	)
	annotationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if annotationFormCallback.CreationMode || annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(annotationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnnotationFormCallback(
			nil,
			annotationFormCallback.probe,
			newFormGroup,
		)
		annotation := new(models.Annotation)
		FillUpForm(annotation, newFormGroup, annotationFormCallback.probe)
		annotationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(annotationFormCallback.probe)
}
func __gong__New__ComplexTypeFormCallback(
	complextype *models.ComplexType,
	probe *Probe,
	formGroup *table.FormGroup,
) (complextypeFormCallback *ComplexTypeFormCallback) {
	complextypeFormCallback = new(ComplexTypeFormCallback)
	complextypeFormCallback.probe = probe
	complextypeFormCallback.complextype = complextype
	complextypeFormCallback.formGroup = formGroup

	complextypeFormCallback.CreationMode = (complextype == nil)

	return
}

type ComplexTypeFormCallback struct {
	complextype *models.ComplexType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complextypeFormCallback *ComplexTypeFormCallback) OnSave() {

	log.Println("ComplexTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complextypeFormCallback.probe.formStage.Checkout()

	if complextypeFormCallback.complextype == nil {
		complextypeFormCallback.complextype = new(models.ComplexType).Stage(complextypeFormCallback.probe.stageOfInterest)
	}
	complextype_ := complextypeFormCallback.complextype
	_ = complextype_

	for _, formDiv := range complextypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complextype_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(complextype_.Annotation), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(complextype_.NameXSD), formDiv)
		case "Sequence":
			FormDivSelectFieldToField(&(complextype_.Sequence), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "Schema:ComplexTypes":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "ComplexTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				complextypeFormCallback.probe.stageOfInterest,
				complextypeFormCallback.probe.backRepoOfInterest,
				complextype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.ComplexTypes, complextype_)
					pastSchemaOwner.ComplexTypes = slices.Delete(pastSchemaOwner.ComplexTypes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](complextypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.ComplexTypes, complextype_)
								pastSchemaOwner.ComplexTypes = slices.Delete(pastSchemaOwner.ComplexTypes, idx, idx+1)
								newSchemaOwner.ComplexTypes = append(newSchemaOwner.ComplexTypes, complextype_)
							}
						} else {
							newSchemaOwner.ComplexTypes = append(newSchemaOwner.ComplexTypes, complextype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextype_.Unstage(complextypeFormCallback.probe.stageOfInterest)
	}

	complextypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ComplexType](
		complextypeFormCallback.probe,
	)
	complextypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if complextypeFormCallback.CreationMode || complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(complextypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			nil,
			complextypeFormCallback.probe,
			newFormGroup,
		)
		complextype := new(models.ComplexType)
		FillUpForm(complextype, newFormGroup, complextypeFormCallback.probe)
		complextypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(complextypeFormCallback.probe)
}
func __gong__New__DocumentationFormCallback(
	documentation *models.Documentation,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentationFormCallback *DocumentationFormCallback) {
	documentationFormCallback = new(DocumentationFormCallback)
	documentationFormCallback.probe = probe
	documentationFormCallback.documentation = documentation
	documentationFormCallback.formGroup = formGroup

	documentationFormCallback.CreationMode = (documentation == nil)

	return
}

type DocumentationFormCallback struct {
	documentation *models.Documentation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentationFormCallback *DocumentationFormCallback) OnSave() {

	log.Println("DocumentationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentationFormCallback.probe.formStage.Checkout()

	if documentationFormCallback.documentation == nil {
		documentationFormCallback.documentation = new(models.Documentation).Stage(documentationFormCallback.probe.stageOfInterest)
	}
	documentation_ := documentationFormCallback.documentation
	_ = documentation_

	for _, formDiv := range documentationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(documentation_.Name), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(documentation_.Source), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(documentation_.Lang), formDiv)
		case "Annotation:Documentations":
			// we need to retrieve the field owner before the change
			var pastAnnotationOwner *models.Annotation
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Annotation"
			rf.Fieldname = "Documentations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				documentationFormCallback.probe.stageOfInterest,
				documentationFormCallback.probe.backRepoOfInterest,
				documentation_,
				&rf)

			if reverseFieldOwner != nil {
				pastAnnotationOwner = reverseFieldOwner.(*models.Annotation)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAnnotationOwner != nil {
					idx := slices.Index(pastAnnotationOwner.Documentations, documentation_)
					pastAnnotationOwner.Documentations = slices.Delete(pastAnnotationOwner.Documentations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _annotation := range *models.GetGongstructInstancesSet[models.Annotation](documentationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _annotation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAnnotationOwner := _annotation // we have a match
						if pastAnnotationOwner != nil {
							if newAnnotationOwner != pastAnnotationOwner {
								idx := slices.Index(pastAnnotationOwner.Documentations, documentation_)
								pastAnnotationOwner.Documentations = slices.Delete(pastAnnotationOwner.Documentations, idx, idx+1)
								newAnnotationOwner.Documentations = append(newAnnotationOwner.Documentations, documentation_)
							}
						} else {
							newAnnotationOwner.Documentations = append(newAnnotationOwner.Documentations, documentation_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentation_.Unstage(documentationFormCallback.probe.stageOfInterest)
	}

	documentationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Documentation](
		documentationFormCallback.probe,
	)
	documentationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if documentationFormCallback.CreationMode || documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(documentationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentationFormCallback(
			nil,
			documentationFormCallback.probe,
			newFormGroup,
		)
		documentation := new(models.Documentation)
		FillUpForm(documentation, newFormGroup, documentationFormCallback.probe)
		documentationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(documentationFormCallback.probe)
}
func __gong__New__ElementFormCallback(
	element *models.Element,
	probe *Probe,
	formGroup *table.FormGroup,
) (elementFormCallback *ElementFormCallback) {
	elementFormCallback = new(ElementFormCallback)
	elementFormCallback.probe = probe
	elementFormCallback.element = element
	elementFormCallback.formGroup = formGroup

	elementFormCallback.CreationMode = (element == nil)

	return
}

type ElementFormCallback struct {
	element *models.Element

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (elementFormCallback *ElementFormCallback) OnSave() {

	log.Println("ElementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	elementFormCallback.probe.formStage.Checkout()

	if elementFormCallback.element == nil {
		elementFormCallback.element = new(models.Element).Stage(elementFormCallback.probe.stageOfInterest)
	}
	element_ := elementFormCallback.element
	_ = element_

	for _, formDiv := range elementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(element_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(element_.Annotation), elementFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(element_.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(element_.Type), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(element_.SimpleType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "ComplexType":
			FormDivSelectFieldToField(&(element_.ComplexType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "Schema:Elements":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.Elements, element_)
					pastSchemaOwner.Elements = slices.Delete(pastSchemaOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.Elements, element_)
								pastSchemaOwner.Elements = slices.Delete(pastSchemaOwner.Elements, idx, idx+1)
								newSchemaOwner.Elements = append(newSchemaOwner.Elements, element_)
							}
						} else {
							newSchemaOwner.Elements = append(newSchemaOwner.Elements, element_)
						}
					}
				}
			}
		case "Sequence:Elements":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Elements, element_)
					pastSequenceOwner.Elements = slices.Delete(pastSequenceOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Elements, element_)
								pastSequenceOwner.Elements = slices.Delete(pastSequenceOwner.Elements, idx, idx+1)
								newSequenceOwner.Elements = append(newSequenceOwner.Elements, element_)
							}
						} else {
							newSequenceOwner.Elements = append(newSequenceOwner.Elements, element_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		element_.Unstage(elementFormCallback.probe.stageOfInterest)
	}

	elementFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Element](
		elementFormCallback.probe,
	)
	elementFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if elementFormCallback.CreationMode || elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(elementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			elementFormCallback.probe,
			newFormGroup,
		)
		element := new(models.Element)
		FillUpForm(element, newFormGroup, elementFormCallback.probe)
		elementFormCallback.probe.formStage.Commit()
	}

	fillUpTree(elementFormCallback.probe)
}
func __gong__New__EnumerationFormCallback(
	enumeration *models.Enumeration,
	probe *Probe,
	formGroup *table.FormGroup,
) (enumerationFormCallback *EnumerationFormCallback) {
	enumerationFormCallback = new(EnumerationFormCallback)
	enumerationFormCallback.probe = probe
	enumerationFormCallback.enumeration = enumeration
	enumerationFormCallback.formGroup = formGroup

	enumerationFormCallback.CreationMode = (enumeration == nil)

	return
}

type EnumerationFormCallback struct {
	enumeration *models.Enumeration

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enumerationFormCallback *EnumerationFormCallback) OnSave() {

	log.Println("EnumerationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enumerationFormCallback.probe.formStage.Checkout()

	if enumerationFormCallback.enumeration == nil {
		enumerationFormCallback.enumeration = new(models.Enumeration).Stage(enumerationFormCallback.probe.stageOfInterest)
	}
	enumeration_ := enumerationFormCallback.enumeration
	_ = enumeration_

	for _, formDiv := range enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enumeration_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(enumeration_.Annotation), enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(enumeration_.Value), formDiv)
		case "Restriction:Enumerations":
			// we need to retrieve the field owner before the change
			var pastRestrictionOwner *models.Restriction
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Restriction"
			rf.Fieldname = "Enumerations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				enumerationFormCallback.probe.stageOfInterest,
				enumerationFormCallback.probe.backRepoOfInterest,
				enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastRestrictionOwner = reverseFieldOwner.(*models.Restriction)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRestrictionOwner != nil {
					idx := slices.Index(pastRestrictionOwner.Enumerations, enumeration_)
					pastRestrictionOwner.Enumerations = slices.Delete(pastRestrictionOwner.Enumerations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _restriction := range *models.GetGongstructInstancesSet[models.Restriction](enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _restriction.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRestrictionOwner := _restriction // we have a match
						if pastRestrictionOwner != nil {
							if newRestrictionOwner != pastRestrictionOwner {
								idx := slices.Index(pastRestrictionOwner.Enumerations, enumeration_)
								pastRestrictionOwner.Enumerations = slices.Delete(pastRestrictionOwner.Enumerations, idx, idx+1)
								newRestrictionOwner.Enumerations = append(newRestrictionOwner.Enumerations, enumeration_)
							}
						} else {
							newRestrictionOwner.Enumerations = append(newRestrictionOwner.Enumerations, enumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumeration_.Unstage(enumerationFormCallback.probe.stageOfInterest)
	}

	enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Enumeration](
		enumerationFormCallback.probe,
	)
	enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if enumerationFormCallback.CreationMode || enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EnumerationFormCallback(
			nil,
			enumerationFormCallback.probe,
			newFormGroup,
		)
		enumeration := new(models.Enumeration)
		FillUpForm(enumeration, newFormGroup, enumerationFormCallback.probe)
		enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(enumerationFormCallback.probe)
}
func __gong__New__MaxInclusiveFormCallback(
	maxinclusive *models.MaxInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (maxinclusiveFormCallback *MaxInclusiveFormCallback) {
	maxinclusiveFormCallback = new(MaxInclusiveFormCallback)
	maxinclusiveFormCallback.probe = probe
	maxinclusiveFormCallback.maxinclusive = maxinclusive
	maxinclusiveFormCallback.formGroup = formGroup

	maxinclusiveFormCallback.CreationMode = (maxinclusive == nil)

	return
}

type MaxInclusiveFormCallback struct {
	maxinclusive *models.MaxInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (maxinclusiveFormCallback *MaxInclusiveFormCallback) OnSave() {

	log.Println("MaxInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	maxinclusiveFormCallback.probe.formStage.Checkout()

	if maxinclusiveFormCallback.maxinclusive == nil {
		maxinclusiveFormCallback.maxinclusive = new(models.MaxInclusive).Stage(maxinclusiveFormCallback.probe.stageOfInterest)
	}
	maxinclusive_ := maxinclusiveFormCallback.maxinclusive
	_ = maxinclusive_

	for _, formDiv := range maxinclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(maxinclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(maxinclusive_.Annotation), maxinclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(maxinclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusive_.Unstage(maxinclusiveFormCallback.probe.stageOfInterest)
	}

	maxinclusiveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MaxInclusive](
		maxinclusiveFormCallback.probe,
	)
	maxinclusiveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if maxinclusiveFormCallback.CreationMode || maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(maxinclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			nil,
			maxinclusiveFormCallback.probe,
			newFormGroup,
		)
		maxinclusive := new(models.MaxInclusive)
		FillUpForm(maxinclusive, newFormGroup, maxinclusiveFormCallback.probe)
		maxinclusiveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(maxinclusiveFormCallback.probe)
}
func __gong__New__MinInclusiveFormCallback(
	mininclusive *models.MinInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (mininclusiveFormCallback *MinInclusiveFormCallback) {
	mininclusiveFormCallback = new(MinInclusiveFormCallback)
	mininclusiveFormCallback.probe = probe
	mininclusiveFormCallback.mininclusive = mininclusive
	mininclusiveFormCallback.formGroup = formGroup

	mininclusiveFormCallback.CreationMode = (mininclusive == nil)

	return
}

type MinInclusiveFormCallback struct {
	mininclusive *models.MinInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mininclusiveFormCallback *MinInclusiveFormCallback) OnSave() {

	log.Println("MinInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mininclusiveFormCallback.probe.formStage.Checkout()

	if mininclusiveFormCallback.mininclusive == nil {
		mininclusiveFormCallback.mininclusive = new(models.MinInclusive).Stage(mininclusiveFormCallback.probe.stageOfInterest)
	}
	mininclusive_ := mininclusiveFormCallback.mininclusive
	_ = mininclusive_

	for _, formDiv := range mininclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mininclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(mininclusive_.Annotation), mininclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(mininclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusive_.Unstage(mininclusiveFormCallback.probe.stageOfInterest)
	}

	mininclusiveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MinInclusive](
		mininclusiveFormCallback.probe,
	)
	mininclusiveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mininclusiveFormCallback.CreationMode || mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mininclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			nil,
			mininclusiveFormCallback.probe,
			newFormGroup,
		)
		mininclusive := new(models.MinInclusive)
		FillUpForm(mininclusive, newFormGroup, mininclusiveFormCallback.probe)
		mininclusiveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mininclusiveFormCallback.probe)
}
func __gong__New__PatternFormCallback(
	pattern *models.Pattern,
	probe *Probe,
	formGroup *table.FormGroup,
) (patternFormCallback *PatternFormCallback) {
	patternFormCallback = new(PatternFormCallback)
	patternFormCallback.probe = probe
	patternFormCallback.pattern = pattern
	patternFormCallback.formGroup = formGroup

	patternFormCallback.CreationMode = (pattern == nil)

	return
}

type PatternFormCallback struct {
	pattern *models.Pattern

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (patternFormCallback *PatternFormCallback) OnSave() {

	log.Println("PatternFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	patternFormCallback.probe.formStage.Checkout()

	if patternFormCallback.pattern == nil {
		patternFormCallback.pattern = new(models.Pattern).Stage(patternFormCallback.probe.stageOfInterest)
	}
	pattern_ := patternFormCallback.pattern
	_ = pattern_

	for _, formDiv := range patternFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pattern_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(pattern_.Annotation), patternFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(pattern_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pattern_.Unstage(patternFormCallback.probe.stageOfInterest)
	}

	patternFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pattern](
		patternFormCallback.probe,
	)
	patternFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if patternFormCallback.CreationMode || patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		patternFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(patternFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PatternFormCallback(
			nil,
			patternFormCallback.probe,
			newFormGroup,
		)
		pattern := new(models.Pattern)
		FillUpForm(pattern, newFormGroup, patternFormCallback.probe)
		patternFormCallback.probe.formStage.Commit()
	}

	fillUpTree(patternFormCallback.probe)
}
func __gong__New__RestrictionFormCallback(
	restriction *models.Restriction,
	probe *Probe,
	formGroup *table.FormGroup,
) (restrictionFormCallback *RestrictionFormCallback) {
	restrictionFormCallback = new(RestrictionFormCallback)
	restrictionFormCallback.probe = probe
	restrictionFormCallback.restriction = restriction
	restrictionFormCallback.formGroup = formGroup

	restrictionFormCallback.CreationMode = (restriction == nil)

	return
}

type RestrictionFormCallback struct {
	restriction *models.Restriction

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (restrictionFormCallback *RestrictionFormCallback) OnSave() {

	log.Println("RestrictionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	restrictionFormCallback.probe.formStage.Checkout()

	if restrictionFormCallback.restriction == nil {
		restrictionFormCallback.restriction = new(models.Restriction).Stage(restrictionFormCallback.probe.stageOfInterest)
	}
	restriction_ := restrictionFormCallback.restriction
	_ = restriction_

	for _, formDiv := range restrictionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(restriction_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(restriction_.Annotation), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Base":
			FormDivBasicFieldToField(&(restriction_.Base), formDiv)
		case "MinInclusive":
			FormDivSelectFieldToField(&(restriction_.MinInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MaxInclusive":
			FormDivSelectFieldToField(&(restriction_.MaxInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Pattern":
			FormDivSelectFieldToField(&(restriction_.Pattern), restrictionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restriction_.Unstage(restrictionFormCallback.probe.stageOfInterest)
	}

	restrictionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Restriction](
		restrictionFormCallback.probe,
	)
	restrictionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if restrictionFormCallback.CreationMode || restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restrictionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(restrictionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RestrictionFormCallback(
			nil,
			restrictionFormCallback.probe,
			newFormGroup,
		)
		restriction := new(models.Restriction)
		FillUpForm(restriction, newFormGroup, restrictionFormCallback.probe)
		restrictionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(restrictionFormCallback.probe)
}
func __gong__New__SchemaFormCallback(
	schema *models.Schema,
	probe *Probe,
	formGroup *table.FormGroup,
) (schemaFormCallback *SchemaFormCallback) {
	schemaFormCallback = new(SchemaFormCallback)
	schemaFormCallback.probe = probe
	schemaFormCallback.schema = schema
	schemaFormCallback.formGroup = formGroup

	schemaFormCallback.CreationMode = (schema == nil)

	return
}

type SchemaFormCallback struct {
	schema *models.Schema

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (schemaFormCallback *SchemaFormCallback) OnSave() {

	log.Println("SchemaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	schemaFormCallback.probe.formStage.Checkout()

	if schemaFormCallback.schema == nil {
		schemaFormCallback.schema = new(models.Schema).Stage(schemaFormCallback.probe.stageOfInterest)
	}
	schema_ := schemaFormCallback.schema
	_ = schema_

	for _, formDiv := range schemaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(schema_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(schema_.Annotation), schemaFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schema_.Unstage(schemaFormCallback.probe.stageOfInterest)
	}

	schemaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Schema](
		schemaFormCallback.probe,
	)
	schemaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if schemaFormCallback.CreationMode || schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schemaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(schemaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SchemaFormCallback(
			nil,
			schemaFormCallback.probe,
			newFormGroup,
		)
		schema := new(models.Schema)
		FillUpForm(schema, newFormGroup, schemaFormCallback.probe)
		schemaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(schemaFormCallback.probe)
}
func __gong__New__SequenceFormCallback(
	sequence *models.Sequence,
	probe *Probe,
	formGroup *table.FormGroup,
) (sequenceFormCallback *SequenceFormCallback) {
	sequenceFormCallback = new(SequenceFormCallback)
	sequenceFormCallback.probe = probe
	sequenceFormCallback.sequence = sequence
	sequenceFormCallback.formGroup = formGroup

	sequenceFormCallback.CreationMode = (sequence == nil)

	return
}

type SequenceFormCallback struct {
	sequence *models.Sequence

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sequenceFormCallback *SequenceFormCallback) OnSave() {

	log.Println("SequenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sequenceFormCallback.probe.formStage.Checkout()

	if sequenceFormCallback.sequence == nil {
		sequenceFormCallback.sequence = new(models.Sequence).Stage(sequenceFormCallback.probe.stageOfInterest)
	}
	sequence_ := sequenceFormCallback.sequence
	_ = sequence_

	for _, formDiv := range sequenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sequence_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(sequence_.Annotation), sequenceFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequence_.Unstage(sequenceFormCallback.probe.stageOfInterest)
	}

	sequenceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Sequence](
		sequenceFormCallback.probe,
	)
	sequenceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sequenceFormCallback.CreationMode || sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sequenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SequenceFormCallback(
			nil,
			sequenceFormCallback.probe,
			newFormGroup,
		)
		sequence := new(models.Sequence)
		FillUpForm(sequence, newFormGroup, sequenceFormCallback.probe)
		sequenceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sequenceFormCallback.probe)
}
func __gong__New__SimpleTypeFormCallback(
	simpletype *models.SimpleType,
	probe *Probe,
	formGroup *table.FormGroup,
) (simpletypeFormCallback *SimpleTypeFormCallback) {
	simpletypeFormCallback = new(SimpleTypeFormCallback)
	simpletypeFormCallback.probe = probe
	simpletypeFormCallback.simpletype = simpletype
	simpletypeFormCallback.formGroup = formGroup

	simpletypeFormCallback.CreationMode = (simpletype == nil)

	return
}

type SimpleTypeFormCallback struct {
	simpletype *models.SimpleType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (simpletypeFormCallback *SimpleTypeFormCallback) OnSave() {

	log.Println("SimpleTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	simpletypeFormCallback.probe.formStage.Checkout()

	if simpletypeFormCallback.simpletype == nil {
		simpletypeFormCallback.simpletype = new(models.SimpleType).Stage(simpletypeFormCallback.probe.stageOfInterest)
	}
	simpletype_ := simpletypeFormCallback.simpletype
	_ = simpletype_

	for _, formDiv := range simpletypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(simpletype_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(simpletype_.Annotation), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(simpletype_.NameXSD), formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(simpletype_.Restriction), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "Schema:SimpleTypes":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "SimpleTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				simpletypeFormCallback.probe.stageOfInterest,
				simpletypeFormCallback.probe.backRepoOfInterest,
				simpletype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.SimpleTypes, simpletype_)
					pastSchemaOwner.SimpleTypes = slices.Delete(pastSchemaOwner.SimpleTypes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](simpletypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.SimpleTypes, simpletype_)
								pastSchemaOwner.SimpleTypes = slices.Delete(pastSchemaOwner.SimpleTypes, idx, idx+1)
								newSchemaOwner.SimpleTypes = append(newSchemaOwner.SimpleTypes, simpletype_)
							}
						} else {
							newSchemaOwner.SimpleTypes = append(newSchemaOwner.SimpleTypes, simpletype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletype_.Unstage(simpletypeFormCallback.probe.stageOfInterest)
	}

	simpletypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SimpleType](
		simpletypeFormCallback.probe,
	)
	simpletypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if simpletypeFormCallback.CreationMode || simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(simpletypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			nil,
			simpletypeFormCallback.probe,
			newFormGroup,
		)
		simpletype := new(models.SimpleType)
		FillUpForm(simpletype, newFormGroup, simpletypeFormCallback.probe)
		simpletypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(simpletypeFormCallback.probe)
}
