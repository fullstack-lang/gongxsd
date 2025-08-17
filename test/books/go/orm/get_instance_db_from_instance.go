// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/test/books/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.A_books:
		a_booksInstance := any(concreteInstance).(*models.A_books)
		ret2 := backRepo.BackRepoA_books.GetA_booksDBFromA_booksPtr(a_booksInstance)
		ret = any(ret2).(*T2)
	case *models.BookType:
		booktypeInstance := any(concreteInstance).(*models.BookType)
		ret2 := backRepo.BackRepoBookType.GetBookTypeDBFromBookTypePtr(booktypeInstance)
		ret = any(ret2).(*T2)
	case *models.Books:
		booksInstance := any(concreteInstance).(*models.Books)
		ret2 := backRepo.BackRepoBooks.GetBooksDBFromBooksPtr(booksInstance)
		ret = any(ret2).(*T2)
	case *models.Credit:
		creditInstance := any(concreteInstance).(*models.Credit)
		ret2 := backRepo.BackRepoCredit.GetCreditDBFromCreditPtr(creditInstance)
		ret = any(ret2).(*T2)
	case *models.Link:
		linkInstance := any(concreteInstance).(*models.Link)
		ret2 := backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.A_books:
		tmp := GetInstanceDBFromInstance[models.A_books, A_booksDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BookType:
		tmp := GetInstanceDBFromInstance[models.BookType, BookTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Books:
		tmp := GetInstanceDBFromInstance[models.Books, BooksDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Credit:
		tmp := GetInstanceDBFromInstance[models.Credit, CreditDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.A_books:
		tmp := GetInstanceDBFromInstance[models.A_books, A_booksDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.BookType:
		tmp := GetInstanceDBFromInstance[models.BookType, BookTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Books:
		tmp := GetInstanceDBFromInstance[models.Books, BooksDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Credit:
		tmp := GetInstanceDBFromInstance[models.Credit, CreditDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
