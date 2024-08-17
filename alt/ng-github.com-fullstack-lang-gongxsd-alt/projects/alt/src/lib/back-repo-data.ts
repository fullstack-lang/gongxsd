// generated code - do not edit

//insertion point for imports
import { AnnotationAPI } from './annotation-api'

import { DocumentationAPI } from './documentation-api'

import { SchemaAPI } from './schema-api'


export class BackRepoData {
	// insertion point for declarations
	AnnotationAPIs = new Array<AnnotationAPI>()

	DocumentationAPIs = new Array<DocumentationAPI>()

	SchemaAPIs = new Array<SchemaAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.AnnotationAPIs = data?.AnnotationAPIs || [];

		this.DocumentationAPIs = data?.DocumentationAPIs || [];

		this.SchemaAPIs = data?.SchemaAPIs || [];

	}

}