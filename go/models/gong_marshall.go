// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Annotation_Identifiers := make(map[*Annotation]string)
	_ = map_Annotation_Identifiers

	annotationOrdered := []*Annotation{}
	for annotation := range stage.Annotations {
		annotationOrdered = append(annotationOrdered, annotation)
	}
	sort.Slice(annotationOrdered[:], func(i, j int) bool {
		return annotationOrdered[i].Name < annotationOrdered[j].Name
	})
	if len(annotationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, annotation := range annotationOrdered {

		id = generatesIdentifier("Annotation", idx, annotation.Name)
		map_Annotation_Identifiers[annotation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Annotation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", annotation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(annotation.Name))
		initializerStatements += setValueField

	}

	map_ComplexType_Identifiers := make(map[*ComplexType]string)
	_ = map_ComplexType_Identifiers

	complextypeOrdered := []*ComplexType{}
	for complextype := range stage.ComplexTypes {
		complextypeOrdered = append(complextypeOrdered, complextype)
	}
	sort.Slice(complextypeOrdered[:], func(i, j int) bool {
		return complextypeOrdered[i].Name < complextypeOrdered[j].Name
	})
	if len(complextypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, complextype := range complextypeOrdered {

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complextype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.NameXSD))
		initializerStatements += setValueField

	}

	map_Documentation_Identifiers := make(map[*Documentation]string)
	_ = map_Documentation_Identifiers

	documentationOrdered := []*Documentation{}
	for documentation := range stage.Documentations {
		documentationOrdered = append(documentationOrdered, documentation)
	}
	sort.Slice(documentationOrdered[:], func(i, j int) bool {
		return documentationOrdered[i].Name < documentationOrdered[j].Name
	})
	if len(documentationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, documentation := range documentationOrdered {

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Documentation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", documentation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Source")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Source))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lang")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Lang))
		initializerStatements += setValueField

	}

	map_Element_Identifiers := make(map[*Element]string)
	_ = map_Element_Identifiers

	elementOrdered := []*Element{}
	for element := range stage.Elements {
		elementOrdered = append(elementOrdered, element)
	}
	sort.Slice(elementOrdered[:], func(i, j int) bool {
		return elementOrdered[i].Name < elementOrdered[j].Name
	})
	if len(elementOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, element := range elementOrdered {

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", element.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Type))
		initializerStatements += setValueField

	}

	map_Enumeration_Identifiers := make(map[*Enumeration]string)
	_ = map_Enumeration_Identifiers

	enumerationOrdered := []*Enumeration{}
	for enumeration := range stage.Enumerations {
		enumerationOrdered = append(enumerationOrdered, enumeration)
	}
	sort.Slice(enumerationOrdered[:], func(i, j int) bool {
		return enumerationOrdered[i].Name < enumerationOrdered[j].Name
	})
	if len(enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, enumeration := range enumerationOrdered {

		id = generatesIdentifier("Enumeration", idx, enumeration.Name)
		map_Enumeration_Identifiers[enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Enumeration")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumeration.Value))
		initializerStatements += setValueField

	}

	map_MaxInclusive_Identifiers := make(map[*MaxInclusive]string)
	_ = map_MaxInclusive_Identifiers

	maxinclusiveOrdered := []*MaxInclusive{}
	for maxinclusive := range stage.MaxInclusives {
		maxinclusiveOrdered = append(maxinclusiveOrdered, maxinclusive)
	}
	sort.Slice(maxinclusiveOrdered[:], func(i, j int) bool {
		return maxinclusiveOrdered[i].Name < maxinclusiveOrdered[j].Name
	})
	if len(maxinclusiveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, maxinclusive := range maxinclusiveOrdered {

		id = generatesIdentifier("MaxInclusive", idx, maxinclusive.Name)
		map_MaxInclusive_Identifiers[maxinclusive] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxInclusive")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", maxinclusive.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxinclusive.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxinclusive.Value))
		initializerStatements += setValueField

	}

	map_MinInclusive_Identifiers := make(map[*MinInclusive]string)
	_ = map_MinInclusive_Identifiers

	mininclusiveOrdered := []*MinInclusive{}
	for mininclusive := range stage.MinInclusives {
		mininclusiveOrdered = append(mininclusiveOrdered, mininclusive)
	}
	sort.Slice(mininclusiveOrdered[:], func(i, j int) bool {
		return mininclusiveOrdered[i].Name < mininclusiveOrdered[j].Name
	})
	if len(mininclusiveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, mininclusive := range mininclusiveOrdered {

		id = generatesIdentifier("MinInclusive", idx, mininclusive.Name)
		map_MinInclusive_Identifiers[mininclusive] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinInclusive")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mininclusive.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mininclusive.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mininclusive.Value))
		initializerStatements += setValueField

	}

	map_Pattern_Identifiers := make(map[*Pattern]string)
	_ = map_Pattern_Identifiers

	patternOrdered := []*Pattern{}
	for pattern := range stage.Patterns {
		patternOrdered = append(patternOrdered, pattern)
	}
	sort.Slice(patternOrdered[:], func(i, j int) bool {
		return patternOrdered[i].Name < patternOrdered[j].Name
	})
	if len(patternOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pattern := range patternOrdered {

		id = generatesIdentifier("Pattern", idx, pattern.Name)
		map_Pattern_Identifiers[pattern] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pattern")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pattern.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pattern.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pattern.Value))
		initializerStatements += setValueField

	}

	map_Restriction_Identifiers := make(map[*Restriction]string)
	_ = map_Restriction_Identifiers

	restrictionOrdered := []*Restriction{}
	for restriction := range stage.Restrictions {
		restrictionOrdered = append(restrictionOrdered, restriction)
	}
	sort.Slice(restrictionOrdered[:], func(i, j int) bool {
		return restrictionOrdered[i].Name < restrictionOrdered[j].Name
	})
	if len(restrictionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, restriction := range restrictionOrdered {

		id = generatesIdentifier("Restriction", idx, restriction.Name)
		map_Restriction_Identifiers[restriction] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Restriction")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", restriction.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(restriction.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(restriction.Base))
		initializerStatements += setValueField

	}

	map_Schema_Identifiers := make(map[*Schema]string)
	_ = map_Schema_Identifiers

	schemaOrdered := []*Schema{}
	for schema := range stage.Schemas {
		schemaOrdered = append(schemaOrdered, schema)
	}
	sort.Slice(schemaOrdered[:], func(i, j int) bool {
		return schemaOrdered[i].Name < schemaOrdered[j].Name
	})
	if len(schemaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, schema := range schemaOrdered {

		id = generatesIdentifier("Schema", idx, schema.Name)
		map_Schema_Identifiers[schema] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Schema")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", schema.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(schema.Name))
		initializerStatements += setValueField

	}

	map_Sequence_Identifiers := make(map[*Sequence]string)
	_ = map_Sequence_Identifiers

	sequenceOrdered := []*Sequence{}
	for sequence := range stage.Sequences {
		sequenceOrdered = append(sequenceOrdered, sequence)
	}
	sort.Slice(sequenceOrdered[:], func(i, j int) bool {
		return sequenceOrdered[i].Name < sequenceOrdered[j].Name
	})
	if len(sequenceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, sequence := range sequenceOrdered {

		id = generatesIdentifier("Sequence", idx, sequence.Name)
		map_Sequence_Identifiers[sequence] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sequence")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sequence.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sequence.Name))
		initializerStatements += setValueField

	}

	map_SimpleType_Identifiers := make(map[*SimpleType]string)
	_ = map_SimpleType_Identifiers

	simpletypeOrdered := []*SimpleType{}
	for simpletype := range stage.SimpleTypes {
		simpletypeOrdered = append(simpletypeOrdered, simpletype)
	}
	sort.Slice(simpletypeOrdered[:], func(i, j int) bool {
		return simpletypeOrdered[i].Name < simpletypeOrdered[j].Name
	})
	if len(simpletypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, simpletype := range simpletypeOrdered {

		id = generatesIdentifier("SimpleType", idx, simpletype.Name)
		map_SimpleType_Identifiers[simpletype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", simpletype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(simpletype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(simpletype.NameXSD))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, annotation := range annotationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Annotation", idx, annotation.Name)
		map_Annotation_Identifiers[annotation] = id

		// Initialisation of values
		for _, _documentation := range annotation.Documentations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Documentations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Documentation_Identifiers[_documentation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, complextype := range complextypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		// Initialisation of values
		if complextype.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[complextype.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Sequence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[complextype.Sequence])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, documentation := range documentationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		// Initialisation of values
	}

	for idx, element := range elementOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		// Initialisation of values
		if element.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[element.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if element.SimpleType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SimpleType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SimpleType_Identifiers[element.SimpleType])
			pointersInitializesStatements += setPointerField
		}

		if element.ComplexType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ComplexType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexType_Identifiers[element.ComplexType])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, enumeration := range enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Enumeration", idx, enumeration.Name)
		map_Enumeration_Identifiers[enumeration] = id

		// Initialisation of values
		if enumeration.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[enumeration.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, maxinclusive := range maxinclusiveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MaxInclusive", idx, maxinclusive.Name)
		map_MaxInclusive_Identifiers[maxinclusive] = id

		// Initialisation of values
		if maxinclusive.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[maxinclusive.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, mininclusive := range mininclusiveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MinInclusive", idx, mininclusive.Name)
		map_MinInclusive_Identifiers[mininclusive] = id

		// Initialisation of values
		if mininclusive.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[mininclusive.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, pattern := range patternOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pattern", idx, pattern.Name)
		map_Pattern_Identifiers[pattern] = id

		// Initialisation of values
		if pattern.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[pattern.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, restriction := range restrictionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Restriction", idx, restriction.Name)
		map_Restriction_Identifiers[restriction] = id

		// Initialisation of values
		if restriction.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[restriction.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _enumeration := range restriction.Enumerations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Enumerations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Enumeration_Identifiers[_enumeration])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MinInclusive != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MinInclusive")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MinInclusive_Identifiers[restriction.MinInclusive])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MaxInclusive != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MaxInclusive")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MaxInclusive_Identifiers[restriction.MaxInclusive])
			pointersInitializesStatements += setPointerField
		}

		if restriction.Pattern != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pattern")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Pattern_Identifiers[restriction.Pattern])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, schema := range schemaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Schema", idx, schema.Name)
		map_Schema_Identifiers[schema] = id

		// Initialisation of values
		if schema.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[schema.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range schema.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _simpletype := range schema.SimpleTypes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SimpleTypes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SimpleType_Identifiers[_simpletype])
			pointersInitializesStatements += setPointerField
		}

		for _, _complextype := range schema.ComplexTypes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ComplexTypes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexType_Identifiers[_complextype])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, sequence := range sequenceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Sequence", idx, sequence.Name)
		map_Sequence_Identifiers[sequence] = id

		// Initialisation of values
		if sequence.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[sequence.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range sequence.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, simpletype := range simpletypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SimpleType", idx, simpletype.Name)
		map_SimpleType_Identifiers[simpletype] = id

		// Initialisation of values
		if simpletype.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[simpletype.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if simpletype.Restriction != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Restriction")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Restriction_Identifiers[simpletype.Restriction])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
