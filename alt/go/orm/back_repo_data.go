// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AnnotationAPIs []*AnnotationAPI

	ComplexTypeAPIs []*ComplexTypeAPI

	DocumentationAPIs []*DocumentationAPI

	SchemaAPIs []*SchemaAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, annotationDB := range backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationDB {

		var annotationAPI AnnotationAPI
		annotationAPI.ID = annotationDB.ID
		annotationAPI.AnnotationPointersEncoding = annotationDB.AnnotationPointersEncoding
		annotationDB.CopyBasicFieldsToAnnotation_WOP(&annotationAPI.Annotation_WOP)

		backRepoData.AnnotationAPIs = append(backRepoData.AnnotationAPIs, &annotationAPI)
	}

	for _, complextypeDB := range backRepo.BackRepoComplexType.Map_ComplexTypeDBID_ComplexTypeDB {

		var complextypeAPI ComplexTypeAPI
		complextypeAPI.ID = complextypeDB.ID
		complextypeAPI.ComplexTypePointersEncoding = complextypeDB.ComplexTypePointersEncoding
		complextypeDB.CopyBasicFieldsToComplexType_WOP(&complextypeAPI.ComplexType_WOP)

		backRepoData.ComplexTypeAPIs = append(backRepoData.ComplexTypeAPIs, &complextypeAPI)
	}

	for _, documentationDB := range backRepo.BackRepoDocumentation.Map_DocumentationDBID_DocumentationDB {

		var documentationAPI DocumentationAPI
		documentationAPI.ID = documentationDB.ID
		documentationAPI.DocumentationPointersEncoding = documentationDB.DocumentationPointersEncoding
		documentationDB.CopyBasicFieldsToDocumentation_WOP(&documentationAPI.Documentation_WOP)

		backRepoData.DocumentationAPIs = append(backRepoData.DocumentationAPIs, &documentationAPI)
	}

	for _, schemaDB := range backRepo.BackRepoSchema.Map_SchemaDBID_SchemaDB {

		var schemaAPI SchemaAPI
		schemaAPI.ID = schemaDB.ID
		schemaAPI.SchemaPointersEncoding = schemaDB.SchemaPointersEncoding
		schemaDB.CopyBasicFieldsToSchema_WOP(&schemaAPI.Schema_WOP)

		backRepoData.SchemaAPIs = append(backRepoData.SchemaAPIs, &schemaAPI)
	}

}
