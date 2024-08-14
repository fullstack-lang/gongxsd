// generated code - do not edit

import { BookTypeAPI } from './booktype-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Credit } from './credit'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BookType {

	static GONGSTRUCT_NAME = "BookType"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Edition: string = ""
	Isbn: string = ""
	Bestseller: boolean = false
	Title: string = ""
	Author: string = ""
	Year: number = 0
	Format: string = ""

	// insertion point for pointers and slices of pointers declarations
	Credit: Array<Credit> = []
}

export function CopyBookTypeToBookTypeAPI(booktype: BookType, booktypeAPI: BookTypeAPI) {

	booktypeAPI.CreatedAt = booktype.CreatedAt
	booktypeAPI.DeletedAt = booktype.DeletedAt
	booktypeAPI.ID = booktype.ID

	// insertion point for basic fields copy operations
	booktypeAPI.Name = booktype.Name
	booktypeAPI.Edition = booktype.Edition
	booktypeAPI.Isbn = booktype.Isbn
	booktypeAPI.Bestseller = booktype.Bestseller
	booktypeAPI.Title = booktype.Title
	booktypeAPI.Author = booktype.Author
	booktypeAPI.Year = booktype.Year
	booktypeAPI.Format = booktype.Format

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	booktypeAPI.BookTypePointersEncoding.Credit = []
	for (let _credit of booktype.Credit) {
		booktypeAPI.BookTypePointersEncoding.Credit.push(_credit.ID)
	}

}

// CopyBookTypeAPIToBookType update basic, pointers and slice of pointers fields of booktype
// from respectively the basic fields and encoded fields of pointers and slices of pointers of booktypeAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBookTypeAPIToBookType(booktypeAPI: BookTypeAPI, booktype: BookType, frontRepo: FrontRepo) {

	booktype.CreatedAt = booktypeAPI.CreatedAt
	booktype.DeletedAt = booktypeAPI.DeletedAt
	booktype.ID = booktypeAPI.ID

	// insertion point for basic fields copy operations
	booktype.Name = booktypeAPI.Name
	booktype.Edition = booktypeAPI.Edition
	booktype.Isbn = booktypeAPI.Isbn
	booktype.Bestseller = booktypeAPI.Bestseller
	booktype.Title = booktypeAPI.Title
	booktype.Author = booktypeAPI.Author
	booktype.Year = booktypeAPI.Year
	booktype.Format = booktypeAPI.Format

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	booktype.Credit = new Array<Credit>()
	for (let _id of booktypeAPI.BookTypePointersEncoding.Credit) {
		let _credit = frontRepo.map_ID_Credit.get(_id)
		if (_credit != undefined) {
			booktype.Credit.push(_credit!)
		}
	}
}
