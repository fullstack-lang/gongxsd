// insertion point for imports
import { AnnotationAPI } from './annotation-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SchemaAPI {

	static GONGSTRUCT_NAME = "Schema"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Xs: string = ""

	// insertion point for other decls

	SchemaPointersEncoding: SchemaPointersEncoding = new SchemaPointersEncoding
}

export class SchemaPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	AnnotationID: NullInt64 = new NullInt64 // if pointer is null, Annotation.ID = 0

}
