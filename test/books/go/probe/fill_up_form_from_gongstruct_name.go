// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/books/go/models"
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
	case "BookDetailsGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BookDetailsGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookDetailsGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		bookdetailsgroup := new(models.BookDetailsGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bookdetailsgroup, formGroup, probe)
	case "BookType":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "BookType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		booktype := new(models.BookType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(booktype, formGroup, probe)
	case "Books":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Books Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BooksFormCallback(
			nil,
			probe,
			formGroup,
		)
		books := new(models.Books)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(books, formGroup, probe)
	case "CommonAttributes":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "CommonAttributes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CommonAttributesFormCallback(
			nil,
			probe,
			formGroup,
		)
		commonattributes := new(models.CommonAttributes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(commonattributes, formGroup, probe)
	case "Credit":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Credit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			probe,
			formGroup,
		)
		credit := new(models.Credit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(credit, formGroup, probe)
	case "ExtendedAttributes":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "ExtendedAttributes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtendedAttributesFormCallback(
			nil,
			probe,
			formGroup,
		)
		extendedattributes := new(models.ExtendedAttributes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extendedattributes, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	}
	formStage.Commit()
}
