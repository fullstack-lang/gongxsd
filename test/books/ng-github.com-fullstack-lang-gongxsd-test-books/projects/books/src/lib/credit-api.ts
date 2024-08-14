// insertion point for imports
import { LinkAPI } from './link-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CreditAPI {

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

	// insertion point for other decls

	CreditPointersEncoding: CreditPointersEncoding = new CreditPointersEncoding
}

export class CreditPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Link: number[] = []
}
