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
