// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	SchemaAPIs []*SchemaAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, schemaDB := range backRepo.BackRepoSchema.Map_SchemaDBID_SchemaDB {

		var schemaAPI SchemaAPI
		schemaAPI.ID = schemaDB.ID
		schemaAPI.SchemaPointersEncoding = schemaDB.SchemaPointersEncoding
		schemaDB.CopyBasicFieldsToSchema_WOP(&schemaAPI.Schema_WOP)

		backRepoData.SchemaAPIs = append(backRepoData.SchemaAPIs, &schemaAPI)
	}

}
