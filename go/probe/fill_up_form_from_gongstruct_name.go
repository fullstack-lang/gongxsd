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
	case "All":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "All Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllFormCallback(
			nil,
			probe,
			formGroup,
		)
		all := new(models.All)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(all, formGroup, probe)
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
	case "Attribute":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Attribute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute := new(models.Attribute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute, formGroup, probe)
	case "AttributeGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AttributeGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributegroup := new(models.AttributeGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributegroup, formGroup, probe)
	case "Choice":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Choice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChoiceFormCallback(
			nil,
			probe,
			formGroup,
		)
		choice := new(models.Choice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(choice, formGroup, probe)
	case "ComplexContent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ComplexContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		complexcontent := new(models.ComplexContent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(complexcontent, formGroup, probe)
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
	case "Extension":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Extension Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtensionFormCallback(
			nil,
			probe,
			formGroup,
		)
		extension := new(models.Extension)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extension, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
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
	case "SimpleContent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SimpleContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		simplecontent := new(models.SimpleContent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(simplecontent, formGroup, probe)
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
	case "Union":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Union Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnionFormCallback(
			nil,
			probe,
			formGroup,
		)
		union := new(models.Union)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(union, formGroup, probe)
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
