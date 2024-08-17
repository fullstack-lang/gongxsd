// insertion point for imports
import { DocumentationAPI } from './documentation-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class AnnotationAPI {

	static GONGSTRUCT_NAME = "Annotation"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	AnnotationPointersEncoding: AnnotationPointersEncoding = new AnnotationPointersEncoding
}

export class AnnotationPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Documentations: number[] = []
}
