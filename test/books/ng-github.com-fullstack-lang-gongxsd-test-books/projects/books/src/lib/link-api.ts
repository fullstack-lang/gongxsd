// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinkAPI {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EnclosedText: string = ""
	NameXSD: string = ""

	// insertion point for other decls

	LinkPointersEncoding: LinkPointersEncoding = new LinkPointersEncoding
}

export class LinkPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
