// generated code - do not edit

import { DocumentationAPI } from './documentation-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Documentation {

	static GONGSTRUCT_NAME = "Documentation"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Text: string = ""
	Source: string = ""
	Lang: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDocumentationToDocumentationAPI(documentation: Documentation, documentationAPI: DocumentationAPI) {

	documentationAPI.CreatedAt = documentation.CreatedAt
	documentationAPI.DeletedAt = documentation.DeletedAt
	documentationAPI.ID = documentation.ID

	// insertion point for basic fields copy operations
	documentationAPI.Name = documentation.Name
	documentationAPI.Text = documentation.Text
	documentationAPI.Source = documentation.Source
	documentationAPI.Lang = documentation.Lang

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDocumentationAPIToDocumentation update basic, pointers and slice of pointers fields of documentation
// from respectively the basic fields and encoded fields of pointers and slices of pointers of documentationAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDocumentationAPIToDocumentation(documentationAPI: DocumentationAPI, documentation: Documentation, frontRepo: FrontRepo) {

	documentation.CreatedAt = documentationAPI.CreatedAt
	documentation.DeletedAt = documentationAPI.DeletedAt
	documentation.ID = documentationAPI.ID

	// insertion point for basic fields copy operations
	documentation.Name = documentationAPI.Name
	documentation.Text = documentationAPI.Text
	documentation.Source = documentationAPI.Source
	documentation.Lang = documentationAPI.Lang

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
