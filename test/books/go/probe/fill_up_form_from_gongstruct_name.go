// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gongxsd/test/books/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "BookType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
			Name:  FormName,
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
	case "Credit":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
