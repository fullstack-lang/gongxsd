// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxsd/test/books/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.BookDetailsGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.BookType:
		switch reverseField.GongstructName {
		// insertion point
		case "Books":
			switch reverseField.Fieldname {
			case "Book":
				if _books, ok := stage.Books_Book_reverseMap[inst]; ok {
					res = _books.Name
				}
			}
		}

	case *models.Books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Credit:
		switch reverseField.GongstructName {
		// insertion point
		case "BookType":
			switch reverseField.Fieldname {
			case "Credit":
				if _booktype, ok := stage.BookType_Credit_reverseMap[inst]; ok {
					res = _booktype.Name
				}
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Link":
				if _credit, ok := stage.Credit_Link_reverseMap[inst]; ok {
					res = _credit.Name
				}
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.BookDetailsGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.BookType:
		switch reverseField.GongstructName {
		// insertion point
		case "Books":
			switch reverseField.Fieldname {
			case "Book":
				res = stage.Books_Book_reverseMap[inst]
			}
		}

	case *models.Books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Credit:
		switch reverseField.GongstructName {
		// insertion point
		case "BookType":
			switch reverseField.Fieldname {
			case "Credit":
				res = stage.BookType_Credit_reverseMap[inst]
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Link":
				res = stage.Credit_Link_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
