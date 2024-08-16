// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BookTypeAPIs []*BookTypeAPI

	CreditAPIs []*CreditAPI

	LinkAPIs []*LinkAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, booktypeDB := range backRepo.BackRepoBookType.Map_BookTypeDBID_BookTypeDB {

		var booktypeAPI BookTypeAPI
		booktypeAPI.ID = booktypeDB.ID
		booktypeAPI.BookTypePointersEncoding = booktypeDB.BookTypePointersEncoding
		booktypeDB.CopyBasicFieldsToBookType_WOP(&booktypeAPI.BookType_WOP)

		backRepoData.BookTypeAPIs = append(backRepoData.BookTypeAPIs, &booktypeAPI)
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
