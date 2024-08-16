// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BookDetailsGroupAPIs []*BookDetailsGroupAPI

	BookTypeAPIs []*BookTypeAPI

	BooksAPIs []*BooksAPI

	CommonAttributesAPIs []*CommonAttributesAPI

	CreditAPIs []*CreditAPI

	ExtendedAttributesAPIs []*ExtendedAttributesAPI

	LinkAPIs []*LinkAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, bookdetailsgroupDB := range backRepo.BackRepoBookDetailsGroup.Map_BookDetailsGroupDBID_BookDetailsGroupDB {

		var bookdetailsgroupAPI BookDetailsGroupAPI
		bookdetailsgroupAPI.ID = bookdetailsgroupDB.ID
		bookdetailsgroupAPI.BookDetailsGroupPointersEncoding = bookdetailsgroupDB.BookDetailsGroupPointersEncoding
		bookdetailsgroupDB.CopyBasicFieldsToBookDetailsGroup_WOP(&bookdetailsgroupAPI.BookDetailsGroup_WOP)

		backRepoData.BookDetailsGroupAPIs = append(backRepoData.BookDetailsGroupAPIs, &bookdetailsgroupAPI)
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

	for _, commonattributesDB := range backRepo.BackRepoCommonAttributes.Map_CommonAttributesDBID_CommonAttributesDB {

		var commonattributesAPI CommonAttributesAPI
		commonattributesAPI.ID = commonattributesDB.ID
		commonattributesAPI.CommonAttributesPointersEncoding = commonattributesDB.CommonAttributesPointersEncoding
		commonattributesDB.CopyBasicFieldsToCommonAttributes_WOP(&commonattributesAPI.CommonAttributes_WOP)

		backRepoData.CommonAttributesAPIs = append(backRepoData.CommonAttributesAPIs, &commonattributesAPI)
	}

	for _, creditDB := range backRepo.BackRepoCredit.Map_CreditDBID_CreditDB {

		var creditAPI CreditAPI
		creditAPI.ID = creditDB.ID
		creditAPI.CreditPointersEncoding = creditDB.CreditPointersEncoding
		creditDB.CopyBasicFieldsToCredit_WOP(&creditAPI.Credit_WOP)

		backRepoData.CreditAPIs = append(backRepoData.CreditAPIs, &creditAPI)
	}

	for _, extendedattributesDB := range backRepo.BackRepoExtendedAttributes.Map_ExtendedAttributesDBID_ExtendedAttributesDB {

		var extendedattributesAPI ExtendedAttributesAPI
		extendedattributesAPI.ID = extendedattributesDB.ID
		extendedattributesAPI.ExtendedAttributesPointersEncoding = extendedattributesDB.ExtendedAttributesPointersEncoding
		extendedattributesDB.CopyBasicFieldsToExtendedAttributes_WOP(&extendedattributesAPI.ExtendedAttributes_WOP)

		backRepoData.ExtendedAttributesAPIs = append(backRepoData.ExtendedAttributesAPIs, &extendedattributesAPI)
	}

	for _, linkDB := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {

		var linkAPI LinkAPI
		linkAPI.ID = linkDB.ID
		linkAPI.LinkPointersEncoding = linkDB.LinkPointersEncoding
		linkDB.CopyBasicFieldsToLink_WOP(&linkAPI.Link_WOP)

		backRepoData.LinkAPIs = append(backRepoData.LinkAPIs, &linkAPI)
	}

}
