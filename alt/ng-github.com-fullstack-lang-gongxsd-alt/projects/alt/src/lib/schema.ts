// generated code - do not edit

import { SchemaAPI } from './schema-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Annotation } from './annotation'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Schema {

	static GONGSTRUCT_NAME = "Schema"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Xs: string = ""

	// insertion point for pointers and slices of pointers declarations
	Annotation?: Annotation

}

export function CopySchemaToSchemaAPI(schema: Schema, schemaAPI: SchemaAPI) {

	schemaAPI.CreatedAt = schema.CreatedAt
	schemaAPI.DeletedAt = schema.DeletedAt
	schemaAPI.ID = schema.ID

	// insertion point for basic fields copy operations
	schemaAPI.Name = schema.Name
	schemaAPI.Xs = schema.Xs

	// insertion point for pointer fields encoding
	schemaAPI.SchemaPointersEncoding.AnnotationID.Valid = true
	if (schema.Annotation != undefined) {
		schemaAPI.SchemaPointersEncoding.AnnotationID.Int64 = schema.Annotation.ID  
	} else {
		schemaAPI.SchemaPointersEncoding.AnnotationID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopySchemaAPIToSchema update basic, pointers and slice of pointers fields of schema
// from respectively the basic fields and encoded fields of pointers and slices of pointers of schemaAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySchemaAPIToSchema(schemaAPI: SchemaAPI, schema: Schema, frontRepo: FrontRepo) {

	schema.CreatedAt = schemaAPI.CreatedAt
	schema.DeletedAt = schemaAPI.DeletedAt
	schema.ID = schemaAPI.ID

	// insertion point for basic fields copy operations
	schema.Name = schemaAPI.Name
	schema.Xs = schemaAPI.Xs

	// insertion point for pointer fields encoding
	schema.Annotation = frontRepo.map_ID_Annotation.get(schemaAPI.SchemaPointersEncoding.AnnotationID.Int64)

	// insertion point for slice of pointers fields encoding
}
