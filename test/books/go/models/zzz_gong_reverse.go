// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *A_books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *BookType:
		switch reverseField.GongstructName {
		// insertion point
		case "A_books":
			switch reverseField.Fieldname {
			case "Book":
				if _a_books, ok := stage.A_books_Book_reverseMap[inst]; ok {
					res = _a_books.Name
				}
			}
		case "Books":
			switch reverseField.Fieldname {
			case "Book":
				if _books, ok := stage.Books_Book_reverseMap[inst]; ok {
					res = _books.Name
				}
			}
		}

	case *Books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Credit:
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

	case *Link:
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

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *A_books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *BookType:
		switch reverseField.GongstructName {
		// insertion point
		case "A_books":
			switch reverseField.Fieldname {
			case "Book":
				res = stage.A_books_Book_reverseMap[inst]
			}
		case "Books":
			switch reverseField.Fieldname {
			case "Book":
				res = stage.Books_Book_reverseMap[inst]
			}
		}

	case *Books:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Credit:
		switch reverseField.GongstructName {
		// insertion point
		case "BookType":
			switch reverseField.Fieldname {
			case "Credit":
				res = stage.BookType_Credit_reverseMap[inst]
			}
		}

	case *Link:
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
