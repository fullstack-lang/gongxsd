// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BookDetailsGroupAPI {

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

	// insertion point for other decls

	BookDetailsGroupPointersEncoding: BookDetailsGroupPointersEncoding = new BookDetailsGroupPointersEncoding
}

export class BookDetailsGroupPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
