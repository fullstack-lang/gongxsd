// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ComplexTypeAPIs []*ComplexTypeAPI

	ElementAPIs []*ElementAPI

	EnumerationAPIs []*EnumerationAPI

	MaxInclusiveAPIs []*MaxInclusiveAPI

	MinInclusiveAPIs []*MinInclusiveAPI

	PatternAPIs []*PatternAPI

	RestrictionAPIs []*RestrictionAPI

	SchemaAPIs []*SchemaAPI

	SequenceAPIs []*SequenceAPI

	SimpleTypeAPIs []*SimpleTypeAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, complextypeDB := range backRepo.BackRepoComplexType.Map_ComplexTypeDBID_ComplexTypeDB {

		var complextypeAPI ComplexTypeAPI
		complextypeAPI.ID = complextypeDB.ID
		complextypeAPI.ComplexTypePointersEncoding = complextypeDB.ComplexTypePointersEncoding
		complextypeDB.CopyBasicFieldsToComplexType_WOP(&complextypeAPI.ComplexType_WOP)

		backRepoData.ComplexTypeAPIs = append(backRepoData.ComplexTypeAPIs, &complextypeAPI)
	}

	for _, elementDB := range backRepo.BackRepoElement.Map_ElementDBID_ElementDB {

		var elementAPI ElementAPI
		elementAPI.ID = elementDB.ID
		elementAPI.ElementPointersEncoding = elementDB.ElementPointersEncoding
		elementDB.CopyBasicFieldsToElement_WOP(&elementAPI.Element_WOP)

		backRepoData.ElementAPIs = append(backRepoData.ElementAPIs, &elementAPI)
	}

	for _, enumerationDB := range backRepo.BackRepoEnumeration.Map_EnumerationDBID_EnumerationDB {

		var enumerationAPI EnumerationAPI
		enumerationAPI.ID = enumerationDB.ID
		enumerationAPI.EnumerationPointersEncoding = enumerationDB.EnumerationPointersEncoding
		enumerationDB.CopyBasicFieldsToEnumeration_WOP(&enumerationAPI.Enumeration_WOP)

		backRepoData.EnumerationAPIs = append(backRepoData.EnumerationAPIs, &enumerationAPI)
	}

	for _, maxinclusiveDB := range backRepo.BackRepoMaxInclusive.Map_MaxInclusiveDBID_MaxInclusiveDB {

		var maxinclusiveAPI MaxInclusiveAPI
		maxinclusiveAPI.ID = maxinclusiveDB.ID
		maxinclusiveAPI.MaxInclusivePointersEncoding = maxinclusiveDB.MaxInclusivePointersEncoding
		maxinclusiveDB.CopyBasicFieldsToMaxInclusive_WOP(&maxinclusiveAPI.MaxInclusive_WOP)

		backRepoData.MaxInclusiveAPIs = append(backRepoData.MaxInclusiveAPIs, &maxinclusiveAPI)
	}

	for _, mininclusiveDB := range backRepo.BackRepoMinInclusive.Map_MinInclusiveDBID_MinInclusiveDB {

		var mininclusiveAPI MinInclusiveAPI
		mininclusiveAPI.ID = mininclusiveDB.ID
		mininclusiveAPI.MinInclusivePointersEncoding = mininclusiveDB.MinInclusivePointersEncoding
		mininclusiveDB.CopyBasicFieldsToMinInclusive_WOP(&mininclusiveAPI.MinInclusive_WOP)

		backRepoData.MinInclusiveAPIs = append(backRepoData.MinInclusiveAPIs, &mininclusiveAPI)
	}

	for _, patternDB := range backRepo.BackRepoPattern.Map_PatternDBID_PatternDB {

		var patternAPI PatternAPI
		patternAPI.ID = patternDB.ID
		patternAPI.PatternPointersEncoding = patternDB.PatternPointersEncoding
		patternDB.CopyBasicFieldsToPattern_WOP(&patternAPI.Pattern_WOP)

		backRepoData.PatternAPIs = append(backRepoData.PatternAPIs, &patternAPI)
	}

	for _, restrictionDB := range backRepo.BackRepoRestriction.Map_RestrictionDBID_RestrictionDB {

		var restrictionAPI RestrictionAPI
		restrictionAPI.ID = restrictionDB.ID
		restrictionAPI.RestrictionPointersEncoding = restrictionDB.RestrictionPointersEncoding
		restrictionDB.CopyBasicFieldsToRestriction_WOP(&restrictionAPI.Restriction_WOP)

		backRepoData.RestrictionAPIs = append(backRepoData.RestrictionAPIs, &restrictionAPI)
	}

	for _, schemaDB := range backRepo.BackRepoSchema.Map_SchemaDBID_SchemaDB {

		var schemaAPI SchemaAPI
		schemaAPI.ID = schemaDB.ID
		schemaAPI.SchemaPointersEncoding = schemaDB.SchemaPointersEncoding
		schemaDB.CopyBasicFieldsToSchema_WOP(&schemaAPI.Schema_WOP)

		backRepoData.SchemaAPIs = append(backRepoData.SchemaAPIs, &schemaAPI)
	}

	for _, sequenceDB := range backRepo.BackRepoSequence.Map_SequenceDBID_SequenceDB {

		var sequenceAPI SequenceAPI
		sequenceAPI.ID = sequenceDB.ID
		sequenceAPI.SequencePointersEncoding = sequenceDB.SequencePointersEncoding
		sequenceDB.CopyBasicFieldsToSequence_WOP(&sequenceAPI.Sequence_WOP)

		backRepoData.SequenceAPIs = append(backRepoData.SequenceAPIs, &sequenceAPI)
	}

	for _, simpletypeDB := range backRepo.BackRepoSimpleType.Map_SimpleTypeDBID_SimpleTypeDB {

		var simpletypeAPI SimpleTypeAPI
		simpletypeAPI.ID = simpletypeDB.ID
		simpletypeAPI.SimpleTypePointersEncoding = simpletypeDB.SimpleTypePointersEncoding
		simpletypeDB.CopyBasicFieldsToSimpleType_WOP(&simpletypeAPI.SimpleType_WOP)

		backRepoData.SimpleTypeAPIs = append(backRepoData.SimpleTypeAPIs, &simpletypeAPI)
	}

}
