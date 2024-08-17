// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AllAPIs []*AllAPI

	AnnotationAPIs []*AnnotationAPI

	AttributeAPIs []*AttributeAPI

	AttributeGroupAPIs []*AttributeGroupAPI

	ChoiceAPIs []*ChoiceAPI

	ComplexContentAPIs []*ComplexContentAPI

	ComplexTypeAPIs []*ComplexTypeAPI

	DocumentationAPIs []*DocumentationAPI

	ElementAPIs []*ElementAPI

	EnumerationAPIs []*EnumerationAPI

	ExtensionAPIs []*ExtensionAPI

	GroupAPIs []*GroupAPI

	LengthAPIs []*LengthAPI

	MaxInclusiveAPIs []*MaxInclusiveAPI

	MaxLengthAPIs []*MaxLengthAPI

	MinInclusiveAPIs []*MinInclusiveAPI

	MinLengthAPIs []*MinLengthAPI

	ModelGroupElementAPIs []*ModelGroupElementAPI

	PatternAPIs []*PatternAPI

	RestrictionAPIs []*RestrictionAPI

	SchemaAPIs []*SchemaAPI

	SequenceAPIs []*SequenceAPI

	SimpleContentAPIs []*SimpleContentAPI

	SimpleTypeAPIs []*SimpleTypeAPI

	TotalDigitAPIs []*TotalDigitAPI

	UnionAPIs []*UnionAPI

	WhiteSpaceAPIs []*WhiteSpaceAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, allDB := range backRepo.BackRepoAll.Map_AllDBID_AllDB {

		var allAPI AllAPI
		allAPI.ID = allDB.ID
		allAPI.AllPointersEncoding = allDB.AllPointersEncoding
		allDB.CopyBasicFieldsToAll_WOP(&allAPI.All_WOP)

		backRepoData.AllAPIs = append(backRepoData.AllAPIs, &allAPI)
	}

	for _, annotationDB := range backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationDB {

		var annotationAPI AnnotationAPI
		annotationAPI.ID = annotationDB.ID
		annotationAPI.AnnotationPointersEncoding = annotationDB.AnnotationPointersEncoding
		annotationDB.CopyBasicFieldsToAnnotation_WOP(&annotationAPI.Annotation_WOP)

		backRepoData.AnnotationAPIs = append(backRepoData.AnnotationAPIs, &annotationAPI)
	}

	for _, attributeDB := range backRepo.BackRepoAttribute.Map_AttributeDBID_AttributeDB {

		var attributeAPI AttributeAPI
		attributeAPI.ID = attributeDB.ID
		attributeAPI.AttributePointersEncoding = attributeDB.AttributePointersEncoding
		attributeDB.CopyBasicFieldsToAttribute_WOP(&attributeAPI.Attribute_WOP)

		backRepoData.AttributeAPIs = append(backRepoData.AttributeAPIs, &attributeAPI)
	}

	for _, attributegroupDB := range backRepo.BackRepoAttributeGroup.Map_AttributeGroupDBID_AttributeGroupDB {

		var attributegroupAPI AttributeGroupAPI
		attributegroupAPI.ID = attributegroupDB.ID
		attributegroupAPI.AttributeGroupPointersEncoding = attributegroupDB.AttributeGroupPointersEncoding
		attributegroupDB.CopyBasicFieldsToAttributeGroup_WOP(&attributegroupAPI.AttributeGroup_WOP)

		backRepoData.AttributeGroupAPIs = append(backRepoData.AttributeGroupAPIs, &attributegroupAPI)
	}

	for _, choiceDB := range backRepo.BackRepoChoice.Map_ChoiceDBID_ChoiceDB {

		var choiceAPI ChoiceAPI
		choiceAPI.ID = choiceDB.ID
		choiceAPI.ChoicePointersEncoding = choiceDB.ChoicePointersEncoding
		choiceDB.CopyBasicFieldsToChoice_WOP(&choiceAPI.Choice_WOP)

		backRepoData.ChoiceAPIs = append(backRepoData.ChoiceAPIs, &choiceAPI)
	}

	for _, complexcontentDB := range backRepo.BackRepoComplexContent.Map_ComplexContentDBID_ComplexContentDB {

		var complexcontentAPI ComplexContentAPI
		complexcontentAPI.ID = complexcontentDB.ID
		complexcontentAPI.ComplexContentPointersEncoding = complexcontentDB.ComplexContentPointersEncoding
		complexcontentDB.CopyBasicFieldsToComplexContent_WOP(&complexcontentAPI.ComplexContent_WOP)

		backRepoData.ComplexContentAPIs = append(backRepoData.ComplexContentAPIs, &complexcontentAPI)
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

	for _, extensionDB := range backRepo.BackRepoExtension.Map_ExtensionDBID_ExtensionDB {

		var extensionAPI ExtensionAPI
		extensionAPI.ID = extensionDB.ID
		extensionAPI.ExtensionPointersEncoding = extensionDB.ExtensionPointersEncoding
		extensionDB.CopyBasicFieldsToExtension_WOP(&extensionAPI.Extension_WOP)

		backRepoData.ExtensionAPIs = append(backRepoData.ExtensionAPIs, &extensionAPI)
	}

	for _, groupDB := range backRepo.BackRepoGroup.Map_GroupDBID_GroupDB {

		var groupAPI GroupAPI
		groupAPI.ID = groupDB.ID
		groupAPI.GroupPointersEncoding = groupDB.GroupPointersEncoding
		groupDB.CopyBasicFieldsToGroup_WOP(&groupAPI.Group_WOP)

		backRepoData.GroupAPIs = append(backRepoData.GroupAPIs, &groupAPI)
	}

	for _, lengthDB := range backRepo.BackRepoLength.Map_LengthDBID_LengthDB {

		var lengthAPI LengthAPI
		lengthAPI.ID = lengthDB.ID
		lengthAPI.LengthPointersEncoding = lengthDB.LengthPointersEncoding
		lengthDB.CopyBasicFieldsToLength_WOP(&lengthAPI.Length_WOP)

		backRepoData.LengthAPIs = append(backRepoData.LengthAPIs, &lengthAPI)
	}

	for _, maxinclusiveDB := range backRepo.BackRepoMaxInclusive.Map_MaxInclusiveDBID_MaxInclusiveDB {

		var maxinclusiveAPI MaxInclusiveAPI
		maxinclusiveAPI.ID = maxinclusiveDB.ID
		maxinclusiveAPI.MaxInclusivePointersEncoding = maxinclusiveDB.MaxInclusivePointersEncoding
		maxinclusiveDB.CopyBasicFieldsToMaxInclusive_WOP(&maxinclusiveAPI.MaxInclusive_WOP)

		backRepoData.MaxInclusiveAPIs = append(backRepoData.MaxInclusiveAPIs, &maxinclusiveAPI)
	}

	for _, maxlengthDB := range backRepo.BackRepoMaxLength.Map_MaxLengthDBID_MaxLengthDB {

		var maxlengthAPI MaxLengthAPI
		maxlengthAPI.ID = maxlengthDB.ID
		maxlengthAPI.MaxLengthPointersEncoding = maxlengthDB.MaxLengthPointersEncoding
		maxlengthDB.CopyBasicFieldsToMaxLength_WOP(&maxlengthAPI.MaxLength_WOP)

		backRepoData.MaxLengthAPIs = append(backRepoData.MaxLengthAPIs, &maxlengthAPI)
	}

	for _, mininclusiveDB := range backRepo.BackRepoMinInclusive.Map_MinInclusiveDBID_MinInclusiveDB {

		var mininclusiveAPI MinInclusiveAPI
		mininclusiveAPI.ID = mininclusiveDB.ID
		mininclusiveAPI.MinInclusivePointersEncoding = mininclusiveDB.MinInclusivePointersEncoding
		mininclusiveDB.CopyBasicFieldsToMinInclusive_WOP(&mininclusiveAPI.MinInclusive_WOP)

		backRepoData.MinInclusiveAPIs = append(backRepoData.MinInclusiveAPIs, &mininclusiveAPI)
	}

	for _, minlengthDB := range backRepo.BackRepoMinLength.Map_MinLengthDBID_MinLengthDB {

		var minlengthAPI MinLengthAPI
		minlengthAPI.ID = minlengthDB.ID
		minlengthAPI.MinLengthPointersEncoding = minlengthDB.MinLengthPointersEncoding
		minlengthDB.CopyBasicFieldsToMinLength_WOP(&minlengthAPI.MinLength_WOP)

		backRepoData.MinLengthAPIs = append(backRepoData.MinLengthAPIs, &minlengthAPI)
	}

	for _, modelgroupelementDB := range backRepo.BackRepoModelGroupElement.Map_ModelGroupElementDBID_ModelGroupElementDB {

		var modelgroupelementAPI ModelGroupElementAPI
		modelgroupelementAPI.ID = modelgroupelementDB.ID
		modelgroupelementAPI.ModelGroupElementPointersEncoding = modelgroupelementDB.ModelGroupElementPointersEncoding
		modelgroupelementDB.CopyBasicFieldsToModelGroupElement_WOP(&modelgroupelementAPI.ModelGroupElement_WOP)

		backRepoData.ModelGroupElementAPIs = append(backRepoData.ModelGroupElementAPIs, &modelgroupelementAPI)
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

	for _, simplecontentDB := range backRepo.BackRepoSimpleContent.Map_SimpleContentDBID_SimpleContentDB {

		var simplecontentAPI SimpleContentAPI
		simplecontentAPI.ID = simplecontentDB.ID
		simplecontentAPI.SimpleContentPointersEncoding = simplecontentDB.SimpleContentPointersEncoding
		simplecontentDB.CopyBasicFieldsToSimpleContent_WOP(&simplecontentAPI.SimpleContent_WOP)

		backRepoData.SimpleContentAPIs = append(backRepoData.SimpleContentAPIs, &simplecontentAPI)
	}

	for _, simpletypeDB := range backRepo.BackRepoSimpleType.Map_SimpleTypeDBID_SimpleTypeDB {

		var simpletypeAPI SimpleTypeAPI
		simpletypeAPI.ID = simpletypeDB.ID
		simpletypeAPI.SimpleTypePointersEncoding = simpletypeDB.SimpleTypePointersEncoding
		simpletypeDB.CopyBasicFieldsToSimpleType_WOP(&simpletypeAPI.SimpleType_WOP)

		backRepoData.SimpleTypeAPIs = append(backRepoData.SimpleTypeAPIs, &simpletypeAPI)
	}

	for _, totaldigitDB := range backRepo.BackRepoTotalDigit.Map_TotalDigitDBID_TotalDigitDB {

		var totaldigitAPI TotalDigitAPI
		totaldigitAPI.ID = totaldigitDB.ID
		totaldigitAPI.TotalDigitPointersEncoding = totaldigitDB.TotalDigitPointersEncoding
		totaldigitDB.CopyBasicFieldsToTotalDigit_WOP(&totaldigitAPI.TotalDigit_WOP)

		backRepoData.TotalDigitAPIs = append(backRepoData.TotalDigitAPIs, &totaldigitAPI)
	}

	for _, unionDB := range backRepo.BackRepoUnion.Map_UnionDBID_UnionDB {

		var unionAPI UnionAPI
		unionAPI.ID = unionDB.ID
		unionAPI.UnionPointersEncoding = unionDB.UnionPointersEncoding
		unionDB.CopyBasicFieldsToUnion_WOP(&unionAPI.Union_WOP)

		backRepoData.UnionAPIs = append(backRepoData.UnionAPIs, &unionAPI)
	}

	for _, whitespaceDB := range backRepo.BackRepoWhiteSpace.Map_WhiteSpaceDBID_WhiteSpaceDB {

		var whitespaceAPI WhiteSpaceAPI
		whitespaceAPI.ID = whitespaceDB.ID
		whitespaceAPI.WhiteSpacePointersEncoding = whitespaceDB.WhiteSpacePointersEncoding
		whitespaceDB.CopyBasicFieldsToWhiteSpace_WOP(&whitespaceAPI.WhiteSpace_WOP)

		backRepoData.WhiteSpaceAPIs = append(backRepoData.WhiteSpaceAPIs, &whitespaceAPI)
	}

}
