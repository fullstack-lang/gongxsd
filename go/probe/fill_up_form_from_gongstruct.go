// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
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
	case *models.Annotation:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Annotation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnnotationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ComplexType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ComplexType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Element:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Element Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Enumeration:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Enumeration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EnumerationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MaxInclusive:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MaxInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MinInclusive:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MinInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pattern:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Pattern Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PatternFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Restriction:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Restriction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestrictionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Schema:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Schema Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SchemaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sequence:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Sequence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SequenceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SimpleType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SimpleType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleTypeFormCallback(
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
