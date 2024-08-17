// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/alt/go/models"
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
	case "Annotation":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Annotation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnnotationFormCallback(
			nil,
			probe,
			formGroup,
		)
		annotation := new(models.Annotation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(annotation, formGroup, probe)
	case "Documentation":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Documentation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentationFormCallback(
			nil,
			probe,
			formGroup,
		)
		documentation := new(models.Documentation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(documentation, formGroup, probe)
	case "Schema":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Schema Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SchemaFormCallback(
			nil,
			probe,
			formGroup,
		)
		schema := new(models.Schema)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(schema, formGroup, probe)
	}
	formStage.Commit()
}
