// insertion point for imports
import { BookTypeAPI } from './booktype-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BooksAPI {

	static GONGSTRUCT_NAME = "Books"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	BooksPointersEncoding: BooksPointersEncoding = new BooksPointersEncoding
}

export class BooksPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Book: number[] = []
}
