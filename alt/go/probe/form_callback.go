// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/alt/go/models"
	"github.com/fullstack-lang/gongxsd/alt/go/orm"
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
func __gong__New__ComplexContentFormCallback(
	complexcontent *models.ComplexContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (complexcontentFormCallback *ComplexContentFormCallback) {
	complexcontentFormCallback = new(ComplexContentFormCallback)
	complexcontentFormCallback.probe = probe
	complexcontentFormCallback.complexcontent = complexcontent
	complexcontentFormCallback.formGroup = formGroup

	complexcontentFormCallback.CreationMode = (complexcontent == nil)

	return
}

type ComplexContentFormCallback struct {
	complexcontent *models.ComplexContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complexcontentFormCallback *ComplexContentFormCallback) OnSave() {

	log.Println("ComplexContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complexcontentFormCallback.probe.formStage.Checkout()

	if complexcontentFormCallback.complexcontent == nil {
		complexcontentFormCallback.complexcontent = new(models.ComplexContent).Stage(complexcontentFormCallback.probe.stageOfInterest)
	}
	complexcontent_ := complexcontentFormCallback.complexcontent
	_ = complexcontent_

	for _, formDiv := range complexcontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complexcontent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontent_.Unstage(complexcontentFormCallback.probe.stageOfInterest)
	}

	complexcontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ComplexContent](
		complexcontentFormCallback.probe,
	)
	complexcontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if complexcontentFormCallback.CreationMode || complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(complexcontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexContentFormCallback(
			nil,
			complexcontentFormCallback.probe,
			newFormGroup,
		)
		complexcontent := new(models.ComplexContent)
		FillUpForm(complexcontent, newFormGroup, complexcontentFormCallback.probe)
		complexcontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(complexcontentFormCallback.probe)
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
		case "Text":
			FormDivBasicFieldToField(&(documentation_.Text), formDiv)
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
		case "Xs":
			FormDivBasicFieldToField(&(schema_.Xs), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(schema_.Annotation), schemaFormCallback.probe.stageOfInterest, formDiv)
		case "Schema_A_ComplexType_A_ComplexContentDummy":
			FormDivBasicFieldToField(&(schema_.Schema_A_ComplexType_A_ComplexContentDummy), formDiv)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy":
			FormDivBasicFieldToField(&(schema_.Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy), formDiv)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy":
			FormDivBasicFieldToField(&(schema_.Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy), formDiv)
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
