// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DocumentationAPI {

	static GONGSTRUCT_NAME = "Documentation"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""
	Source: string = ""
	Lang: string = ""

	// insertion point for other decls

	DocumentationPointersEncoding: DocumentationPointersEncoding = new DocumentationPointersEncoding
}

export class DocumentationPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}
