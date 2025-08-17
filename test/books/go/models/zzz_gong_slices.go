// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct BookType
	// insertion point per field
	clear(stage.BookType_Credit_reverseMap)
	stage.BookType_Credit_reverseMap = make(map[*Credit]*BookType)
	for booktype := range stage.BookTypes {
		_ = booktype
		for _, _credit := range booktype.Credit {
			stage.BookType_Credit_reverseMap[_credit] = booktype
		}
	}

	// Compute reverse map for named struct Books
	// insertion point per field
	clear(stage.Books_Book_reverseMap)
	stage.Books_Book_reverseMap = make(map[*BookType]*Books)
	for books := range stage.Bookss {
		_ = books
		for _, _booktype := range books.Book {
			stage.Books_Book_reverseMap[_booktype] = books
		}
	}

	// Compute reverse map for named struct Credit
	// insertion point per field
	clear(stage.Credit_Link_reverseMap)
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

}
