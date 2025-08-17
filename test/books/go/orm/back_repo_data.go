// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	A_booksAPIs []*A_booksAPI

	BookTypeAPIs []*BookTypeAPI

	BooksAPIs []*BooksAPI

	CreditAPIs []*CreditAPI

	LinkAPIs []*LinkAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, a_booksDB := range backRepo.BackRepoA_books.Map_A_booksDBID_A_booksDB {

		var a_booksAPI A_booksAPI
		a_booksAPI.ID = a_booksDB.ID
		a_booksAPI.A_booksPointersEncoding = a_booksDB.A_booksPointersEncoding
		a_booksDB.CopyBasicFieldsToA_books_WOP(&a_booksAPI.A_books_WOP)

		backRepoData.A_booksAPIs = append(backRepoData.A_booksAPIs, &a_booksAPI)
	}

	for _, booktypeDB := range backRepo.BackRepoBookType.Map_BookTypeDBID_BookTypeDB {

		var booktypeAPI BookTypeAPI
		booktypeAPI.ID = booktypeDB.ID
		booktypeAPI.BookTypePointersEncoding = booktypeDB.BookTypePointersEncoding
		booktypeDB.CopyBasicFieldsToBookType_WOP(&booktypeAPI.BookType_WOP)

		backRepoData.BookTypeAPIs = append(backRepoData.BookTypeAPIs, &booktypeAPI)
	}

	for _, booksDB := range backRepo.BackRepoBooks.Map_BooksDBID_BooksDB {

		var booksAPI BooksAPI
		booksAPI.ID = booksDB.ID
		booksAPI.BooksPointersEncoding = booksDB.BooksPointersEncoding
		booksDB.CopyBasicFieldsToBooks_WOP(&booksAPI.Books_WOP)

		backRepoData.BooksAPIs = append(backRepoData.BooksAPIs, &booksAPI)
	}

	for _, creditDB := range backRepo.BackRepoCredit.Map_CreditDBID_CreditDB {

		var creditAPI CreditAPI
		creditAPI.ID = creditDB.ID
		creditAPI.CreditPointersEncoding = creditDB.CreditPointersEncoding
		creditDB.CopyBasicFieldsToCredit_WOP(&creditAPI.Credit_WOP)

		backRepoData.CreditAPIs = append(backRepoData.CreditAPIs, &creditAPI)
	}

	for _, linkDB := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {

		var linkAPI LinkAPI
		linkAPI.ID = linkDB.ID
		linkAPI.LinkPointersEncoding = linkDB.LinkPointersEncoding
		linkDB.CopyBasicFieldsToLink_WOP(&linkAPI.Link_WOP)

		backRepoData.LinkAPIs = append(backRepoData.LinkAPIs, &linkAPI)
	}

}
