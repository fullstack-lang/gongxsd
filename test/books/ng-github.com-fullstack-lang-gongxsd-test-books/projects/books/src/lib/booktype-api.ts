// insertion point for imports
import { CreditAPI } from './credit-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BookTypeAPI {

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

	// insertion point for other decls

	BookTypePointersEncoding: BookTypePointersEncoding = new BookTypePointersEncoding
}

export class BookTypePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Credit: number[] = []
}
