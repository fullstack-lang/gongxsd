// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/books/go/models"
	"github.com/fullstack-lang/gongxsd/test/books/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__BookDetailsGroupFormCallback(
	bookdetailsgroup *models.BookDetailsGroup,
	probe *Probe,
	formGroup *table.FormGroup,
) (bookdetailsgroupFormCallback *BookDetailsGroupFormCallback) {
	bookdetailsgroupFormCallback = new(BookDetailsGroupFormCallback)
	bookdetailsgroupFormCallback.probe = probe
	bookdetailsgroupFormCallback.bookdetailsgroup = bookdetailsgroup
	bookdetailsgroupFormCallback.formGroup = formGroup

	bookdetailsgroupFormCallback.CreationMode = (bookdetailsgroup == nil)

	return
}

type BookDetailsGroupFormCallback struct {
	bookdetailsgroup *models.BookDetailsGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bookdetailsgroupFormCallback *BookDetailsGroupFormCallback) OnSave() {

	log.Println("BookDetailsGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bookdetailsgroupFormCallback.probe.formStage.Checkout()

	if bookdetailsgroupFormCallback.bookdetailsgroup == nil {
		bookdetailsgroupFormCallback.bookdetailsgroup = new(models.BookDetailsGroup).Stage(bookdetailsgroupFormCallback.probe.stageOfInterest)
	}
	bookdetailsgroup_ := bookdetailsgroupFormCallback.bookdetailsgroup
	_ = bookdetailsgroup_

	for _, formDiv := range bookdetailsgroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bookdetailsgroup_.Name), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(bookdetailsgroup_.Title), formDiv)
		case "Author":
			FormDivBasicFieldToField(&(bookdetailsgroup_.Author), formDiv)
		case "Year":
			FormDivBasicFieldToField(&(bookdetailsgroup_.Year), formDiv)
		case "Format":
			FormDivBasicFieldToField(&(bookdetailsgroup_.Format), formDiv)
		}
	}

	// manage the suppress operation
	if bookdetailsgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bookdetailsgroup_.Unstage(bookdetailsgroupFormCallback.probe.stageOfInterest)
	}

	bookdetailsgroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BookDetailsGroup](
		bookdetailsgroupFormCallback.probe,
	)
	bookdetailsgroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bookdetailsgroupFormCallback.CreationMode || bookdetailsgroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bookdetailsgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(bookdetailsgroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BookDetailsGroupFormCallback(
			nil,
			bookdetailsgroupFormCallback.probe,
			newFormGroup,
		)
		bookdetailsgroup := new(models.BookDetailsGroup)
		FillUpForm(bookdetailsgroup, newFormGroup, bookdetailsgroupFormCallback.probe)
		bookdetailsgroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bookdetailsgroupFormCallback.probe)
}
func __gong__New__BookTypeFormCallback(
	booktype *models.BookType,
	probe *Probe,
	formGroup *table.FormGroup,
) (booktypeFormCallback *BookTypeFormCallback) {
	booktypeFormCallback = new(BookTypeFormCallback)
	booktypeFormCallback.probe = probe
	booktypeFormCallback.booktype = booktype
	booktypeFormCallback.formGroup = formGroup

	booktypeFormCallback.CreationMode = (booktype == nil)

	return
}

type BookTypeFormCallback struct {
	booktype *models.BookType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (booktypeFormCallback *BookTypeFormCallback) OnSave() {

	log.Println("BookTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	booktypeFormCallback.probe.formStage.Checkout()

	if booktypeFormCallback.booktype == nil {
		booktypeFormCallback.booktype = new(models.BookType).Stage(booktypeFormCallback.probe.stageOfInterest)
	}
	booktype_ := booktypeFormCallback.booktype
	_ = booktype_

	for _, formDiv := range booktypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(booktype_.Name), formDiv)
		case "ExtendedAttributes":
			FormDivSelectFieldToField(&(booktype_.ExtendedAttributes), booktypeFormCallback.probe.stageOfInterest, formDiv)
		case "Title":
			FormDivBasicFieldToField(&(booktype_.Title), formDiv)
		case "Author":
			FormDivBasicFieldToField(&(booktype_.Author), formDiv)
		case "Year":
			FormDivBasicFieldToField(&(booktype_.Year), formDiv)
		case "Format":
			FormDivBasicFieldToField(&(booktype_.Format), formDiv)
		case "Books:Book":
			// we need to retrieve the field owner before the change
			var pastBooksOwner *models.Books
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Books"
			rf.Fieldname = "Book"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				booktypeFormCallback.probe.stageOfInterest,
				booktypeFormCallback.probe.backRepoOfInterest,
				booktype_,
				&rf)

			if reverseFieldOwner != nil {
				pastBooksOwner = reverseFieldOwner.(*models.Books)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBooksOwner != nil {
					idx := slices.Index(pastBooksOwner.Book, booktype_)
					pastBooksOwner.Book = slices.Delete(pastBooksOwner.Book, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _books := range *models.GetGongstructInstancesSet[models.Books](booktypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _books.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBooksOwner := _books // we have a match
						if pastBooksOwner != nil {
							if newBooksOwner != pastBooksOwner {
								idx := slices.Index(pastBooksOwner.Book, booktype_)
								pastBooksOwner.Book = slices.Delete(pastBooksOwner.Book, idx, idx+1)
								newBooksOwner.Book = append(newBooksOwner.Book, booktype_)
							}
						} else {
							newBooksOwner.Book = append(newBooksOwner.Book, booktype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if booktypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booktype_.Unstage(booktypeFormCallback.probe.stageOfInterest)
	}

	booktypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.BookType](
		booktypeFormCallback.probe,
	)
	booktypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if booktypeFormCallback.CreationMode || booktypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booktypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(booktypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BookTypeFormCallback(
			nil,
			booktypeFormCallback.probe,
			newFormGroup,
		)
		booktype := new(models.BookType)
		FillUpForm(booktype, newFormGroup, booktypeFormCallback.probe)
		booktypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(booktypeFormCallback.probe)
}
func __gong__New__BooksFormCallback(
	books *models.Books,
	probe *Probe,
	formGroup *table.FormGroup,
) (booksFormCallback *BooksFormCallback) {
	booksFormCallback = new(BooksFormCallback)
	booksFormCallback.probe = probe
	booksFormCallback.books = books
	booksFormCallback.formGroup = formGroup

	booksFormCallback.CreationMode = (books == nil)

	return
}

type BooksFormCallback struct {
	books *models.Books

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (booksFormCallback *BooksFormCallback) OnSave() {

	log.Println("BooksFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	booksFormCallback.probe.formStage.Checkout()

	if booksFormCallback.books == nil {
		booksFormCallback.books = new(models.Books).Stage(booksFormCallback.probe.stageOfInterest)
	}
	books_ := booksFormCallback.books
	_ = books_

	for _, formDiv := range booksFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(books_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		books_.Unstage(booksFormCallback.probe.stageOfInterest)
	}

	booksFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Books](
		booksFormCallback.probe,
	)
	booksFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if booksFormCallback.CreationMode || booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booksFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(booksFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BooksFormCallback(
			nil,
			booksFormCallback.probe,
			newFormGroup,
		)
		books := new(models.Books)
		FillUpForm(books, newFormGroup, booksFormCallback.probe)
		booksFormCallback.probe.formStage.Commit()
	}

	fillUpTree(booksFormCallback.probe)
}
func __gong__New__CommonAttributesFormCallback(
	commonattributes *models.CommonAttributes,
	probe *Probe,
	formGroup *table.FormGroup,
) (commonattributesFormCallback *CommonAttributesFormCallback) {
	commonattributesFormCallback = new(CommonAttributesFormCallback)
	commonattributesFormCallback.probe = probe
	commonattributesFormCallback.commonattributes = commonattributes
	commonattributesFormCallback.formGroup = formGroup

	commonattributesFormCallback.CreationMode = (commonattributes == nil)

	return
}

type CommonAttributesFormCallback struct {
	commonattributes *models.CommonAttributes

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (commonattributesFormCallback *CommonAttributesFormCallback) OnSave() {

	log.Println("CommonAttributesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	commonattributesFormCallback.probe.formStage.Checkout()

	if commonattributesFormCallback.commonattributes == nil {
		commonattributesFormCallback.commonattributes = new(models.CommonAttributes).Stage(commonattributesFormCallback.probe.stageOfInterest)
	}
	commonattributes_ := commonattributesFormCallback.commonattributes
	_ = commonattributes_

	for _, formDiv := range commonattributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(commonattributes_.Name), formDiv)
		case "Isbn":
			FormDivBasicFieldToField(&(commonattributes_.Isbn), formDiv)
		case "Bestseller":
			FormDivBasicFieldToField(&(commonattributes_.Bestseller), formDiv)
		}
	}

	// manage the suppress operation
	if commonattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		commonattributes_.Unstage(commonattributesFormCallback.probe.stageOfInterest)
	}

	commonattributesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CommonAttributes](
		commonattributesFormCallback.probe,
	)
	commonattributesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if commonattributesFormCallback.CreationMode || commonattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		commonattributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(commonattributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CommonAttributesFormCallback(
			nil,
			commonattributesFormCallback.probe,
			newFormGroup,
		)
		commonattributes := new(models.CommonAttributes)
		FillUpForm(commonattributes, newFormGroup, commonattributesFormCallback.probe)
		commonattributesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(commonattributesFormCallback.probe)
}
func __gong__New__CreditFormCallback(
	credit *models.Credit,
	probe *Probe,
	formGroup *table.FormGroup,
) (creditFormCallback *CreditFormCallback) {
	creditFormCallback = new(CreditFormCallback)
	creditFormCallback.probe = probe
	creditFormCallback.credit = credit
	creditFormCallback.formGroup = formGroup

	creditFormCallback.CreationMode = (credit == nil)

	return
}

type CreditFormCallback struct {
	credit *models.Credit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (creditFormCallback *CreditFormCallback) OnSave() {

	log.Println("CreditFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	creditFormCallback.probe.formStage.Checkout()

	if creditFormCallback.credit == nil {
		creditFormCallback.credit = new(models.Credit).Stage(creditFormCallback.probe.stageOfInterest)
	}
	credit_ := creditFormCallback.credit
	_ = credit_

	for _, formDiv := range creditFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(credit_.Name), formDiv)
		case "Page":
			FormDivBasicFieldToField(&(credit_.Page), formDiv)
		case "Credit_type":
			FormDivBasicFieldToField(&(credit_.Credit_type), formDiv)
		case "Credit_words":
			FormDivBasicFieldToField(&(credit_.Credit_words), formDiv)
		case "Credit_symbol":
			FormDivBasicFieldToField(&(credit_.Credit_symbol), formDiv)
		case "BookType:Credit":
			// we need to retrieve the field owner before the change
			var pastBookTypeOwner *models.BookType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "BookType"
			rf.Fieldname = "Credit"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				creditFormCallback.probe.stageOfInterest,
				creditFormCallback.probe.backRepoOfInterest,
				credit_,
				&rf)

			if reverseFieldOwner != nil {
				pastBookTypeOwner = reverseFieldOwner.(*models.BookType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBookTypeOwner != nil {
					idx := slices.Index(pastBookTypeOwner.Credit, credit_)
					pastBookTypeOwner.Credit = slices.Delete(pastBookTypeOwner.Credit, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _booktype := range *models.GetGongstructInstancesSet[models.BookType](creditFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _booktype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBookTypeOwner := _booktype // we have a match
						if pastBookTypeOwner != nil {
							if newBookTypeOwner != pastBookTypeOwner {
								idx := slices.Index(pastBookTypeOwner.Credit, credit_)
								pastBookTypeOwner.Credit = slices.Delete(pastBookTypeOwner.Credit, idx, idx+1)
								newBookTypeOwner.Credit = append(newBookTypeOwner.Credit, credit_)
							}
						} else {
							newBookTypeOwner.Credit = append(newBookTypeOwner.Credit, credit_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		credit_.Unstage(creditFormCallback.probe.stageOfInterest)
	}

	creditFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Credit](
		creditFormCallback.probe,
	)
	creditFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if creditFormCallback.CreationMode || creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		creditFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(creditFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			creditFormCallback.probe,
			newFormGroup,
		)
		credit := new(models.Credit)
		FillUpForm(credit, newFormGroup, creditFormCallback.probe)
		creditFormCallback.probe.formStage.Commit()
	}

	fillUpTree(creditFormCallback.probe)
}
func __gong__New__ExtendedAttributesFormCallback(
	extendedattributes *models.ExtendedAttributes,
	probe *Probe,
	formGroup *table.FormGroup,
) (extendedattributesFormCallback *ExtendedAttributesFormCallback) {
	extendedattributesFormCallback = new(ExtendedAttributesFormCallback)
	extendedattributesFormCallback.probe = probe
	extendedattributesFormCallback.extendedattributes = extendedattributes
	extendedattributesFormCallback.formGroup = formGroup

	extendedattributesFormCallback.CreationMode = (extendedattributes == nil)

	return
}

type ExtendedAttributesFormCallback struct {
	extendedattributes *models.ExtendedAttributes

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (extendedattributesFormCallback *ExtendedAttributesFormCallback) OnSave() {

	log.Println("ExtendedAttributesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	extendedattributesFormCallback.probe.formStage.Checkout()

	if extendedattributesFormCallback.extendedattributes == nil {
		extendedattributesFormCallback.extendedattributes = new(models.ExtendedAttributes).Stage(extendedattributesFormCallback.probe.stageOfInterest)
	}
	extendedattributes_ := extendedattributesFormCallback.extendedattributes
	_ = extendedattributes_

	for _, formDiv := range extendedattributesFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(extendedattributes_.Name), formDiv)
		case "Edition":
			FormDivBasicFieldToField(&(extendedattributes_.Edition), formDiv)
		case "CommonAttributes":
			FormDivSelectFieldToField(&(extendedattributes_.CommonAttributes), extendedattributesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if extendedattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extendedattributes_.Unstage(extendedattributesFormCallback.probe.stageOfInterest)
	}

	extendedattributesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ExtendedAttributes](
		extendedattributesFormCallback.probe,
	)
	extendedattributesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if extendedattributesFormCallback.CreationMode || extendedattributesFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extendedattributesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(extendedattributesFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExtendedAttributesFormCallback(
			nil,
			extendedattributesFormCallback.probe,
			newFormGroup,
		)
		extendedattributes := new(models.ExtendedAttributes)
		FillUpForm(extendedattributes, newFormGroup, extendedattributesFormCallback.probe)
		extendedattributesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(extendedattributesFormCallback.probe)
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(link_.NameXSD), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(link_.EnclosedText), formDiv)
		case "Credit:Link":
			// we need to retrieve the field owner before the change
			var pastCreditOwner *models.Credit
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Link"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkFormCallback.probe.stageOfInterest,
				linkFormCallback.probe.backRepoOfInterest,
				link_,
				&rf)

			if reverseFieldOwner != nil {
				pastCreditOwner = reverseFieldOwner.(*models.Credit)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCreditOwner != nil {
					idx := slices.Index(pastCreditOwner.Link, link_)
					pastCreditOwner.Link = slices.Delete(pastCreditOwner.Link, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _credit := range *models.GetGongstructInstancesSet[models.Credit](linkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _credit.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCreditOwner := _credit // we have a match
						if pastCreditOwner != nil {
							if newCreditOwner != pastCreditOwner {
								idx := slices.Index(pastCreditOwner.Link, link_)
								pastCreditOwner.Link = slices.Delete(pastCreditOwner.Link, idx, idx+1)
								newCreditOwner.Link = append(newCreditOwner.Link, link_)
							}
						} else {
							newCreditOwner.Link = append(newCreditOwner.Link, link_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.probe,
	)
	linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(linkFormCallback.probe)
}
