// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
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
	case "ComplexType":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ComplexType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		complextype := new(models.ComplexType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(complextype, formGroup, probe)
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
	case "Element":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Element Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			probe,
			formGroup,
		)
		element := new(models.Element)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(element, formGroup, probe)
	case "Enumeration":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Enumeration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EnumerationFormCallback(
			nil,
			probe,
			formGroup,
		)
		enumeration := new(models.Enumeration)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(enumeration, formGroup, probe)
	case "Length":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Length Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		length := new(models.Length)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(length, formGroup, probe)
	case "MaxInclusive":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MaxInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			nil,
			probe,
			formGroup,
		)
		maxinclusive := new(models.MaxInclusive)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(maxinclusive, formGroup, probe)
	case "MaxLength":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MaxLength Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxLengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		maxlength := new(models.MaxLength)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(maxlength, formGroup, probe)
	case "MinInclusive":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MinInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			nil,
			probe,
			formGroup,
		)
		mininclusive := new(models.MinInclusive)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mininclusive, formGroup, probe)
	case "MinLength":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "MinLength Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinLengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		minlength := new(models.MinLength)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(minlength, formGroup, probe)
	case "Pattern":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pattern Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PatternFormCallback(
			nil,
			probe,
			formGroup,
		)
		pattern := new(models.Pattern)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pattern, formGroup, probe)
	case "Restriction":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Restriction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestrictionFormCallback(
			nil,
			probe,
			formGroup,
		)
		restriction := new(models.Restriction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(restriction, formGroup, probe)
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
	case "Sequence":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Sequence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SequenceFormCallback(
			nil,
			probe,
			formGroup,
		)
		sequence := new(models.Sequence)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sequence, formGroup, probe)
	case "SimpleType":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SimpleType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		simpletype := new(models.SimpleType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(simpletype, formGroup, probe)
	case "TotalDigit":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "TotalDigit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TotalDigitFormCallback(
			nil,
			probe,
			formGroup,
		)
		totaldigit := new(models.TotalDigit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(totaldigit, formGroup, probe)
	case "WhiteSpace":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "WhiteSpace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WhiteSpaceFormCallback(
			nil,
			probe,
			formGroup,
		)
		whitespace := new(models.WhiteSpace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(whitespace, formGroup, probe)
	}
	formStage.Commit()
}
