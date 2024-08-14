// generated code - do not edit

import { BookDetailsGroupAPI } from './bookdetailsgroup-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BookDetailsGroup {

	static GONGSTRUCT_NAME = "BookDetailsGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Title: string = ""
	Author: string = ""
	Year: number = 0
	Format: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyBookDetailsGroupToBookDetailsGroupAPI(bookdetailsgroup: BookDetailsGroup, bookdetailsgroupAPI: BookDetailsGroupAPI) {

	bookdetailsgroupAPI.CreatedAt = bookdetailsgroup.CreatedAt
	bookdetailsgroupAPI.DeletedAt = bookdetailsgroup.DeletedAt
	bookdetailsgroupAPI.ID = bookdetailsgroup.ID

	// insertion point for basic fields copy operations
	bookdetailsgroupAPI.Name = bookdetailsgroup.Name
	bookdetailsgroupAPI.Title = bookdetailsgroup.Title
	bookdetailsgroupAPI.Author = bookdetailsgroup.Author
	bookdetailsgroupAPI.Year = bookdetailsgroup.Year
	bookdetailsgroupAPI.Format = bookdetailsgroup.Format

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBookDetailsGroupAPIToBookDetailsGroup update basic, pointers and slice of pointers fields of bookdetailsgroup
// from respectively the basic fields and encoded fields of pointers and slices of pointers of bookdetailsgroupAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBookDetailsGroupAPIToBookDetailsGroup(bookdetailsgroupAPI: BookDetailsGroupAPI, bookdetailsgroup: BookDetailsGroup, frontRepo: FrontRepo) {

	bookdetailsgroup.CreatedAt = bookdetailsgroupAPI.CreatedAt
	bookdetailsgroup.DeletedAt = bookdetailsgroupAPI.DeletedAt
	bookdetailsgroup.ID = bookdetailsgroupAPI.ID

	// insertion point for basic fields copy operations
	bookdetailsgroup.Name = bookdetailsgroupAPI.Name
	bookdetailsgroup.Title = bookdetailsgroupAPI.Title
	bookdetailsgroup.Author = bookdetailsgroupAPI.Author
	bookdetailsgroup.Year = bookdetailsgroupAPI.Year
	bookdetailsgroup.Format = bookdetailsgroupAPI.Format

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
