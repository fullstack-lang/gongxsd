// generated code - do not edit

import { CreditAPI } from './credit-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Link } from './link'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Credit {

	static GONGSTRUCT_NAME = "Credit"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Page: number = 0
	Credit_type: string = ""
	Credit_words: string = ""
	Credit_symbol: string = ""

	// insertion point for pointers and slices of pointers declarations
	Link: Array<Link> = []
}

export function CopyCreditToCreditAPI(credit: Credit, creditAPI: CreditAPI) {

	creditAPI.CreatedAt = credit.CreatedAt
	creditAPI.DeletedAt = credit.DeletedAt
	creditAPI.ID = credit.ID

	// insertion point for basic fields copy operations
	creditAPI.Name = credit.Name
	creditAPI.Page = credit.Page
	creditAPI.Credit_type = credit.Credit_type
	creditAPI.Credit_words = credit.Credit_words
	creditAPI.Credit_symbol = credit.Credit_symbol

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	creditAPI.CreditPointersEncoding.Link = []
	for (let _link of credit.Link) {
		creditAPI.CreditPointersEncoding.Link.push(_link.ID)
	}

}

// CopyCreditAPIToCredit update basic, pointers and slice of pointers fields of credit
// from respectively the basic fields and encoded fields of pointers and slices of pointers of creditAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCreditAPIToCredit(creditAPI: CreditAPI, credit: Credit, frontRepo: FrontRepo) {

	credit.CreatedAt = creditAPI.CreatedAt
	credit.DeletedAt = creditAPI.DeletedAt
	credit.ID = creditAPI.ID

	// insertion point for basic fields copy operations
	credit.Name = creditAPI.Name
	credit.Page = creditAPI.Page
	credit.Credit_type = creditAPI.Credit_type
	credit.Credit_words = creditAPI.Credit_words
	credit.Credit_symbol = creditAPI.Credit_symbol

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	credit.Link = new Array<Link>()
	for (let _id of creditAPI.CreditPointersEncoding.Link) {
		let _link = frontRepo.map_ID_Link.get(_id)
		if (_link != undefined) {
			credit.Link.push(_link!)
		}
	}
}
