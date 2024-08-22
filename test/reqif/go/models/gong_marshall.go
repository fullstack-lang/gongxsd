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
	map_ALTERNATIVE_ID_Identifiers := make(map[*ALTERNATIVE_ID]string)
	_ = map_ALTERNATIVE_ID_Identifiers

	alternative_idOrdered := []*ALTERNATIVE_ID{}
	for alternative_id := range stage.ALTERNATIVE_IDs {
		alternative_idOrdered = append(alternative_idOrdered, alternative_id)
	}
	sort.Slice(alternative_idOrdered[:], func(i, j int) bool {
		return alternative_idOrdered[i].Name < alternative_idOrdered[j].Name
	})
	if len(alternative_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, alternative_id := range alternative_idOrdered {

		id = generatesIdentifier("ALTERNATIVE_ID", idx, alternative_id.Name)
		map_ALTERNATIVE_ID_Identifiers[alternative_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ALTERNATIVE_ID")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", alternative_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternative_id.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(alternative_id.IDENTIFIER))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]string)
	_ = map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers

	attribute_definition_booleanOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN{}
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_booleanOrdered = append(attribute_definition_booleanOrdered, attribute_definition_boolean)
	}
	sort.Slice(attribute_definition_booleanOrdered[:], func(i, j int) bool {
		return attribute_definition_booleanOrdered[i].Name < attribute_definition_booleanOrdered[j].Name
	})
	if len(attribute_definition_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_boolean := range attribute_definition_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", idx, attribute_definition_boolean.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[attribute_definition_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_DATE_Identifiers := make(map[*ATTRIBUTE_DEFINITION_DATE]string)
	_ = map_ATTRIBUTE_DEFINITION_DATE_Identifiers

	attribute_definition_dateOrdered := []*ATTRIBUTE_DEFINITION_DATE{}
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		attribute_definition_dateOrdered = append(attribute_definition_dateOrdered, attribute_definition_date)
	}
	sort.Slice(attribute_definition_dateOrdered[:], func(i, j int) bool {
		return attribute_definition_dateOrdered[i].Name < attribute_definition_dateOrdered[j].Name
	})
	if len(attribute_definition_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_date := range attribute_definition_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", idx, attribute_definition_date.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Identifiers[attribute_definition_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]string)
	_ = map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers

	attribute_definition_enumerationOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION{}
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		attribute_definition_enumerationOrdered = append(attribute_definition_enumerationOrdered, attribute_definition_enumeration)
	}
	sort.Slice(attribute_definition_enumerationOrdered[:], func(i, j int) bool {
		return attribute_definition_enumerationOrdered[i].Name < attribute_definition_enumerationOrdered[j].Name
	})
	if len(attribute_definition_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", idx, attribute_definition_enumeration.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[attribute_definition_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MULTI_VALUED")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.MULTI_VALUED))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers := make(map[*ATTRIBUTE_DEFINITION_INTEGER]string)
	_ = map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers

	attribute_definition_integerOrdered := []*ATTRIBUTE_DEFINITION_INTEGER{}
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integerOrdered = append(attribute_definition_integerOrdered, attribute_definition_integer)
	}
	sort.Slice(attribute_definition_integerOrdered[:], func(i, j int) bool {
		return attribute_definition_integerOrdered[i].Name < attribute_definition_integerOrdered[j].Name
	})
	if len(attribute_definition_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_integer := range attribute_definition_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", idx, attribute_definition_integer.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[attribute_definition_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_REAL_Identifiers := make(map[*ATTRIBUTE_DEFINITION_REAL]string)
	_ = map_ATTRIBUTE_DEFINITION_REAL_Identifiers

	attribute_definition_realOrdered := []*ATTRIBUTE_DEFINITION_REAL{}
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		attribute_definition_realOrdered = append(attribute_definition_realOrdered, attribute_definition_real)
	}
	sort.Slice(attribute_definition_realOrdered[:], func(i, j int) bool {
		return attribute_definition_realOrdered[i].Name < attribute_definition_realOrdered[j].Name
	})
	if len(attribute_definition_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_real := range attribute_definition_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", idx, attribute_definition_real.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Identifiers[attribute_definition_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_STRING_Identifiers := make(map[*ATTRIBUTE_DEFINITION_STRING]string)
	_ = map_ATTRIBUTE_DEFINITION_STRING_Identifiers

	attribute_definition_stringOrdered := []*ATTRIBUTE_DEFINITION_STRING{}
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_stringOrdered = append(attribute_definition_stringOrdered, attribute_definition_string)
	}
	sort.Slice(attribute_definition_stringOrdered[:], func(i, j int) bool {
		return attribute_definition_stringOrdered[i].Name < attribute_definition_stringOrdered[j].Name
	})
	if len(attribute_definition_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_string := range attribute_definition_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", idx, attribute_definition_string.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Identifiers[attribute_definition_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_DEFINITION_XHTML_Identifiers := make(map[*ATTRIBUTE_DEFINITION_XHTML]string)
	_ = map_ATTRIBUTE_DEFINITION_XHTML_Identifiers

	attribute_definition_xhtmlOrdered := []*ATTRIBUTE_DEFINITION_XHTML{}
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		attribute_definition_xhtmlOrdered = append(attribute_definition_xhtmlOrdered, attribute_definition_xhtml)
	}
	sort.Slice(attribute_definition_xhtmlOrdered[:], func(i, j int) bool {
		return attribute_definition_xhtmlOrdered[i].Name < attribute_definition_xhtmlOrdered[j].Name
	})
	if len(attribute_definition_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", idx, attribute_definition_xhtml.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[attribute_definition_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LONG_NAME))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers := make(map[*ATTRIBUTE_VALUE_BOOLEAN]string)
	_ = map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers

	attribute_value_booleanOrdered := []*ATTRIBUTE_VALUE_BOOLEAN{}
	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		attribute_value_booleanOrdered = append(attribute_value_booleanOrdered, attribute_value_boolean)
	}
	sort.Slice(attribute_value_booleanOrdered[:], func(i, j int) bool {
		return attribute_value_booleanOrdered[i].Name < attribute_value_booleanOrdered[j].Name
	})
	if len(attribute_value_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_boolean := range attribute_value_booleanOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", idx, attribute_value_boolean.Name)
		map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_value_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_boolean.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_boolean.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_DATE_Identifiers := make(map[*ATTRIBUTE_VALUE_DATE]string)
	_ = map_ATTRIBUTE_VALUE_DATE_Identifiers

	attribute_value_dateOrdered := []*ATTRIBUTE_VALUE_DATE{}
	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		attribute_value_dateOrdered = append(attribute_value_dateOrdered, attribute_value_date)
	}
	sort.Slice(attribute_value_dateOrdered[:], func(i, j int) bool {
		return attribute_value_dateOrdered[i].Name < attribute_value_dateOrdered[j].Name
	})
	if len(attribute_value_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_date := range attribute_value_dateOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", idx, attribute_value_date.Name)
		map_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_value_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_date.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers := make(map[*ATTRIBUTE_VALUE_ENUMERATION]string)
	_ = map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers

	attribute_value_enumerationOrdered := []*ATTRIBUTE_VALUE_ENUMERATION{}
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		attribute_value_enumerationOrdered = append(attribute_value_enumerationOrdered, attribute_value_enumeration)
	}
	sort.Slice(attribute_value_enumerationOrdered[:], func(i, j int) bool {
		return attribute_value_enumerationOrdered[i].Name < attribute_value_enumerationOrdered[j].Name
	})
	if len(attribute_value_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_enumeration := range attribute_value_enumerationOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", idx, attribute_value_enumeration.Name)
		map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_value_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_enumeration.Name))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_INTEGER_Identifiers := make(map[*ATTRIBUTE_VALUE_INTEGER]string)
	_ = map_ATTRIBUTE_VALUE_INTEGER_Identifiers

	attribute_value_integerOrdered := []*ATTRIBUTE_VALUE_INTEGER{}
	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		attribute_value_integerOrdered = append(attribute_value_integerOrdered, attribute_value_integer)
	}
	sort.Slice(attribute_value_integerOrdered[:], func(i, j int) bool {
		return attribute_value_integerOrdered[i].Name < attribute_value_integerOrdered[j].Name
	})
	if len(attribute_value_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_integer := range attribute_value_integerOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", idx, attribute_value_integer.Name)
		map_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_value_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_integer.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attribute_value_integer.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_REAL_Identifiers := make(map[*ATTRIBUTE_VALUE_REAL]string)
	_ = map_ATTRIBUTE_VALUE_REAL_Identifiers

	attribute_value_realOrdered := []*ATTRIBUTE_VALUE_REAL{}
	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		attribute_value_realOrdered = append(attribute_value_realOrdered, attribute_value_real)
	}
	sort.Slice(attribute_value_realOrdered[:], func(i, j int) bool {
		return attribute_value_realOrdered[i].Name < attribute_value_realOrdered[j].Name
	})
	if len(attribute_value_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_real := range attribute_value_realOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", idx, attribute_value_real.Name)
		map_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_value_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_real.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", attribute_value_real.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_STRING_Identifiers := make(map[*ATTRIBUTE_VALUE_STRING]string)
	_ = map_ATTRIBUTE_VALUE_STRING_Identifiers

	attribute_value_stringOrdered := []*ATTRIBUTE_VALUE_STRING{}
	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		attribute_value_stringOrdered = append(attribute_value_stringOrdered, attribute_value_string)
	}
	sort.Slice(attribute_value_stringOrdered[:], func(i, j int) bool {
		return attribute_value_stringOrdered[i].Name < attribute_value_stringOrdered[j].Name
	})
	if len(attribute_value_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_string := range attribute_value_stringOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", idx, attribute_value_string.Name)
		map_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_value_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "THE_VALUE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_string.THE_VALUE))
		initializerStatements += setValueField

	}

	map_ATTRIBUTE_VALUE_XHTML_Identifiers := make(map[*ATTRIBUTE_VALUE_XHTML]string)
	_ = map_ATTRIBUTE_VALUE_XHTML_Identifiers

	attribute_value_xhtmlOrdered := []*ATTRIBUTE_VALUE_XHTML{}
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtmlOrdered = append(attribute_value_xhtmlOrdered, attribute_value_xhtml)
	}
	sort.Slice(attribute_value_xhtmlOrdered[:], func(i, j int) bool {
		return attribute_value_xhtmlOrdered[i].Name < attribute_value_xhtmlOrdered[j].Name
	})
	if len(attribute_value_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute_value_xhtml := range attribute_value_xhtmlOrdered {

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", idx, attribute_value_xhtml.Name)
		map_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_value_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute_value_xhtml.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_SIMPLIFIED")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_xhtml.IS_SIMPLIFIED))
		initializerStatements += setValueField

	}

	map_A_ALTERNATIVE_ID_Identifiers := make(map[*A_ALTERNATIVE_ID]string)
	_ = map_A_ALTERNATIVE_ID_Identifiers

	a_alternative_idOrdered := []*A_ALTERNATIVE_ID{}
	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		a_alternative_idOrdered = append(a_alternative_idOrdered, a_alternative_id)
	}
	sort.Slice(a_alternative_idOrdered[:], func(i, j int) bool {
		return a_alternative_idOrdered[i].Name < a_alternative_idOrdered[j].Name
	})
	if len(a_alternative_idOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_alternative_id := range a_alternative_idOrdered {

		id = generatesIdentifier("A_ALTERNATIVE_ID", idx, a_alternative_id.Name)
		map_A_ALTERNATIVE_ID_Identifiers[a_alternative_id] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ALTERNATIVE_ID")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_alternative_id.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_alternative_id.Name))
		initializerStatements += setValueField

	}

	map_A_CHILDREN_Identifiers := make(map[*A_CHILDREN]string)
	_ = map_A_CHILDREN_Identifiers

	a_childrenOrdered := []*A_CHILDREN{}
	for a_children := range stage.A_CHILDRENs {
		a_childrenOrdered = append(a_childrenOrdered, a_children)
	}
	sort.Slice(a_childrenOrdered[:], func(i, j int) bool {
		return a_childrenOrdered[i].Name < a_childrenOrdered[j].Name
	})
	if len(a_childrenOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_children := range a_childrenOrdered {

		id = generatesIdentifier("A_CHILDREN", idx, a_children.Name)
		map_A_CHILDREN_Identifiers[a_children] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CHILDREN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_children.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_children.Name))
		initializerStatements += setValueField

	}

	map_A_CORE_CONTENT_Identifiers := make(map[*A_CORE_CONTENT]string)
	_ = map_A_CORE_CONTENT_Identifiers

	a_core_contentOrdered := []*A_CORE_CONTENT{}
	for a_core_content := range stage.A_CORE_CONTENTs {
		a_core_contentOrdered = append(a_core_contentOrdered, a_core_content)
	}
	sort.Slice(a_core_contentOrdered[:], func(i, j int) bool {
		return a_core_contentOrdered[i].Name < a_core_contentOrdered[j].Name
	})
	if len(a_core_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_core_content := range a_core_contentOrdered {

		id = generatesIdentifier("A_CORE_CONTENT", idx, a_core_content.Name)
		map_A_CORE_CONTENT_Identifiers[a_core_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CORE_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_core_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_core_content.Name))
		initializerStatements += setValueField

	}

	map_A_DATATYPES_Identifiers := make(map[*A_DATATYPES]string)
	_ = map_A_DATATYPES_Identifiers

	a_datatypesOrdered := []*A_DATATYPES{}
	for a_datatypes := range stage.A_DATATYPESs {
		a_datatypesOrdered = append(a_datatypesOrdered, a_datatypes)
	}
	sort.Slice(a_datatypesOrdered[:], func(i, j int) bool {
		return a_datatypesOrdered[i].Name < a_datatypesOrdered[j].Name
	})
	if len(a_datatypesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_datatypes := range a_datatypesOrdered {

		id = generatesIdentifier("A_DATATYPES", idx, a_datatypes.Name)
		map_A_DATATYPES_Identifiers[a_datatypes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatypes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_datatypes.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_Identifiers := make(map[*A_DEFAULT_VALUE]string)
	_ = map_A_DEFAULT_VALUE_Identifiers

	a_default_valueOrdered := []*A_DEFAULT_VALUE{}
	for a_default_value := range stage.A_DEFAULT_VALUEs {
		a_default_valueOrdered = append(a_default_valueOrdered, a_default_value)
	}
	sort.Slice(a_default_valueOrdered[:], func(i, j int) bool {
		return a_default_valueOrdered[i].Name < a_default_valueOrdered[j].Name
	})
	if len(a_default_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value := range a_default_valueOrdered {

		id = generatesIdentifier("A_DEFAULT_VALUE", idx, a_default_value.Name)
		map_A_DEFAULT_VALUE_Identifiers[a_default_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_1_Identifiers := make(map[*A_DEFAULT_VALUE_1]string)
	_ = map_A_DEFAULT_VALUE_1_Identifiers

	a_default_value_1Ordered := []*A_DEFAULT_VALUE_1{}
	for a_default_value_1 := range stage.A_DEFAULT_VALUE_1s {
		a_default_value_1Ordered = append(a_default_value_1Ordered, a_default_value_1)
	}
	sort.Slice(a_default_value_1Ordered[:], func(i, j int) bool {
		return a_default_value_1Ordered[i].Name < a_default_value_1Ordered[j].Name
	})
	if len(a_default_value_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_1 := range a_default_value_1Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_1", idx, a_default_value_1.Name)
		map_A_DEFAULT_VALUE_1_Identifiers[a_default_value_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_1.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_2_Identifiers := make(map[*A_DEFAULT_VALUE_2]string)
	_ = map_A_DEFAULT_VALUE_2_Identifiers

	a_default_value_2Ordered := []*A_DEFAULT_VALUE_2{}
	for a_default_value_2 := range stage.A_DEFAULT_VALUE_2s {
		a_default_value_2Ordered = append(a_default_value_2Ordered, a_default_value_2)
	}
	sort.Slice(a_default_value_2Ordered[:], func(i, j int) bool {
		return a_default_value_2Ordered[i].Name < a_default_value_2Ordered[j].Name
	})
	if len(a_default_value_2Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_2 := range a_default_value_2Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_2", idx, a_default_value_2.Name)
		map_A_DEFAULT_VALUE_2_Identifiers[a_default_value_2] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_2")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_2.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_2.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_3_Identifiers := make(map[*A_DEFAULT_VALUE_3]string)
	_ = map_A_DEFAULT_VALUE_3_Identifiers

	a_default_value_3Ordered := []*A_DEFAULT_VALUE_3{}
	for a_default_value_3 := range stage.A_DEFAULT_VALUE_3s {
		a_default_value_3Ordered = append(a_default_value_3Ordered, a_default_value_3)
	}
	sort.Slice(a_default_value_3Ordered[:], func(i, j int) bool {
		return a_default_value_3Ordered[i].Name < a_default_value_3Ordered[j].Name
	})
	if len(a_default_value_3Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_3 := range a_default_value_3Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_3", idx, a_default_value_3.Name)
		map_A_DEFAULT_VALUE_3_Identifiers[a_default_value_3] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_3")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_3.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_3.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_4_Identifiers := make(map[*A_DEFAULT_VALUE_4]string)
	_ = map_A_DEFAULT_VALUE_4_Identifiers

	a_default_value_4Ordered := []*A_DEFAULT_VALUE_4{}
	for a_default_value_4 := range stage.A_DEFAULT_VALUE_4s {
		a_default_value_4Ordered = append(a_default_value_4Ordered, a_default_value_4)
	}
	sort.Slice(a_default_value_4Ordered[:], func(i, j int) bool {
		return a_default_value_4Ordered[i].Name < a_default_value_4Ordered[j].Name
	})
	if len(a_default_value_4Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_4 := range a_default_value_4Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_4", idx, a_default_value_4.Name)
		map_A_DEFAULT_VALUE_4_Identifiers[a_default_value_4] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_4")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_4.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_4.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_5_Identifiers := make(map[*A_DEFAULT_VALUE_5]string)
	_ = map_A_DEFAULT_VALUE_5_Identifiers

	a_default_value_5Ordered := []*A_DEFAULT_VALUE_5{}
	for a_default_value_5 := range stage.A_DEFAULT_VALUE_5s {
		a_default_value_5Ordered = append(a_default_value_5Ordered, a_default_value_5)
	}
	sort.Slice(a_default_value_5Ordered[:], func(i, j int) bool {
		return a_default_value_5Ordered[i].Name < a_default_value_5Ordered[j].Name
	})
	if len(a_default_value_5Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_5 := range a_default_value_5Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_5", idx, a_default_value_5.Name)
		map_A_DEFAULT_VALUE_5_Identifiers[a_default_value_5] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_5")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_5.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_5.Name))
		initializerStatements += setValueField

	}

	map_A_DEFAULT_VALUE_6_Identifiers := make(map[*A_DEFAULT_VALUE_6]string)
	_ = map_A_DEFAULT_VALUE_6_Identifiers

	a_default_value_6Ordered := []*A_DEFAULT_VALUE_6{}
	for a_default_value_6 := range stage.A_DEFAULT_VALUE_6s {
		a_default_value_6Ordered = append(a_default_value_6Ordered, a_default_value_6)
	}
	sort.Slice(a_default_value_6Ordered[:], func(i, j int) bool {
		return a_default_value_6Ordered[i].Name < a_default_value_6Ordered[j].Name
	})
	if len(a_default_value_6Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_default_value_6 := range a_default_value_6Ordered {

		id = generatesIdentifier("A_DEFAULT_VALUE_6", idx, a_default_value_6.Name)
		map_A_DEFAULT_VALUE_6_Identifiers[a_default_value_6] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFAULT_VALUE_6")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_default_value_6.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_default_value_6.Name))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_Identifiers := make(map[*A_DEFINITION]string)
	_ = map_A_DEFINITION_Identifiers

	a_definitionOrdered := []*A_DEFINITION{}
	for a_definition := range stage.A_DEFINITIONs {
		a_definitionOrdered = append(a_definitionOrdered, a_definition)
	}
	sort.Slice(a_definitionOrdered[:], func(i, j int) bool {
		return a_definitionOrdered[i].Name < a_definitionOrdered[j].Name
	})
	if len(a_definitionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition := range a_definitionOrdered {

		id = generatesIdentifier("A_DEFINITION", idx, a_definition.Name)
		map_A_DEFINITION_Identifiers[a_definition] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition.ATTRIBUTE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_1_Identifiers := make(map[*A_DEFINITION_1]string)
	_ = map_A_DEFINITION_1_Identifiers

	a_definition_1Ordered := []*A_DEFINITION_1{}
	for a_definition_1 := range stage.A_DEFINITION_1s {
		a_definition_1Ordered = append(a_definition_1Ordered, a_definition_1)
	}
	sort.Slice(a_definition_1Ordered[:], func(i, j int) bool {
		return a_definition_1Ordered[i].Name < a_definition_1Ordered[j].Name
	})
	if len(a_definition_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_1 := range a_definition_1Ordered {

		id = generatesIdentifier("A_DEFINITION_1", idx, a_definition_1.Name)
		map_A_DEFINITION_1_Identifiers[a_definition_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_1.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_1.ATTRIBUTE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_2_Identifiers := make(map[*A_DEFINITION_2]string)
	_ = map_A_DEFINITION_2_Identifiers

	a_definition_2Ordered := []*A_DEFINITION_2{}
	for a_definition_2 := range stage.A_DEFINITION_2s {
		a_definition_2Ordered = append(a_definition_2Ordered, a_definition_2)
	}
	sort.Slice(a_definition_2Ordered[:], func(i, j int) bool {
		return a_definition_2Ordered[i].Name < a_definition_2Ordered[j].Name
	})
	if len(a_definition_2Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_2 := range a_definition_2Ordered {

		id = generatesIdentifier("A_DEFINITION_2", idx, a_definition_2.Name)
		map_A_DEFINITION_2_Identifiers[a_definition_2] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_2")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_2.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_2.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_2.ATTRIBUTE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_3_Identifiers := make(map[*A_DEFINITION_3]string)
	_ = map_A_DEFINITION_3_Identifiers

	a_definition_3Ordered := []*A_DEFINITION_3{}
	for a_definition_3 := range stage.A_DEFINITION_3s {
		a_definition_3Ordered = append(a_definition_3Ordered, a_definition_3)
	}
	sort.Slice(a_definition_3Ordered[:], func(i, j int) bool {
		return a_definition_3Ordered[i].Name < a_definition_3Ordered[j].Name
	})
	if len(a_definition_3Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_3 := range a_definition_3Ordered {

		id = generatesIdentifier("A_DEFINITION_3", idx, a_definition_3.Name)
		map_A_DEFINITION_3_Identifiers[a_definition_3] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_3")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_3.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_3.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_3.ATTRIBUTE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_4_Identifiers := make(map[*A_DEFINITION_4]string)
	_ = map_A_DEFINITION_4_Identifiers

	a_definition_4Ordered := []*A_DEFINITION_4{}
	for a_definition_4 := range stage.A_DEFINITION_4s {
		a_definition_4Ordered = append(a_definition_4Ordered, a_definition_4)
	}
	sort.Slice(a_definition_4Ordered[:], func(i, j int) bool {
		return a_definition_4Ordered[i].Name < a_definition_4Ordered[j].Name
	})
	if len(a_definition_4Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_4 := range a_definition_4Ordered {

		id = generatesIdentifier("A_DEFINITION_4", idx, a_definition_4.Name)
		map_A_DEFINITION_4_Identifiers[a_definition_4] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_4")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_4.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_4.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_4.ATTRIBUTE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_5_Identifiers := make(map[*A_DEFINITION_5]string)
	_ = map_A_DEFINITION_5_Identifiers

	a_definition_5Ordered := []*A_DEFINITION_5{}
	for a_definition_5 := range stage.A_DEFINITION_5s {
		a_definition_5Ordered = append(a_definition_5Ordered, a_definition_5)
	}
	sort.Slice(a_definition_5Ordered[:], func(i, j int) bool {
		return a_definition_5Ordered[i].Name < a_definition_5Ordered[j].Name
	})
	if len(a_definition_5Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_5 := range a_definition_5Ordered {

		id = generatesIdentifier("A_DEFINITION_5", idx, a_definition_5.Name)
		map_A_DEFINITION_5_Identifiers[a_definition_5] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_5")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_5.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_5.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_5.ATTRIBUTE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

	}

	map_A_DEFINITION_6_Identifiers := make(map[*A_DEFINITION_6]string)
	_ = map_A_DEFINITION_6_Identifiers

	a_definition_6Ordered := []*A_DEFINITION_6{}
	for a_definition_6 := range stage.A_DEFINITION_6s {
		a_definition_6Ordered = append(a_definition_6Ordered, a_definition_6)
	}
	sort.Slice(a_definition_6Ordered[:], func(i, j int) bool {
		return a_definition_6Ordered[i].Name < a_definition_6Ordered[j].Name
	})
	if len(a_definition_6Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_definition_6 := range a_definition_6Ordered {

		id = generatesIdentifier("A_DEFINITION_6", idx, a_definition_6.Name)
		map_A_DEFINITION_6_Identifiers[a_definition_6] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DEFINITION_6")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_definition_6.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_6.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_definition_6.ATTRIBUTE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

	}

	map_A_EDITABLE_ATTS_Identifiers := make(map[*A_EDITABLE_ATTS]string)
	_ = map_A_EDITABLE_ATTS_Identifiers

	a_editable_attsOrdered := []*A_EDITABLE_ATTS{}
	for a_editable_atts := range stage.A_EDITABLE_ATTSs {
		a_editable_attsOrdered = append(a_editable_attsOrdered, a_editable_atts)
	}
	sort.Slice(a_editable_attsOrdered[:], func(i, j int) bool {
		return a_editable_attsOrdered[i].Name < a_editable_attsOrdered[j].Name
	})
	if len(a_editable_attsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_editable_atts := range a_editable_attsOrdered {

		id = generatesIdentifier("A_EDITABLE_ATTS", idx, a_editable_atts.Name)
		map_A_EDITABLE_ATTS_Identifiers[a_editable_atts] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_EDITABLE_ATTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_editable_atts.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_OBJECT_Identifiers := make(map[*A_OBJECT]string)
	_ = map_A_OBJECT_Identifiers

	a_objectOrdered := []*A_OBJECT{}
	for a_object := range stage.A_OBJECTs {
		a_objectOrdered = append(a_objectOrdered, a_object)
	}
	sort.Slice(a_objectOrdered[:], func(i, j int) bool {
		return a_objectOrdered[i].Name < a_objectOrdered[j].Name
	})
	if len(a_objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_object := range a_objectOrdered {

		id = generatesIdentifier("A_OBJECT", idx, a_object.Name)
		map_A_OBJECT_Identifiers[a_object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_OBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_object.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_object.SPEC_OBJECT_REF))
		initializerStatements += setValueField

	}

	map_A_PROPERTIES_Identifiers := make(map[*A_PROPERTIES]string)
	_ = map_A_PROPERTIES_Identifiers

	a_propertiesOrdered := []*A_PROPERTIES{}
	for a_properties := range stage.A_PROPERTIESs {
		a_propertiesOrdered = append(a_propertiesOrdered, a_properties)
	}
	sort.Slice(a_propertiesOrdered[:], func(i, j int) bool {
		return a_propertiesOrdered[i].Name < a_propertiesOrdered[j].Name
	})
	if len(a_propertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_properties := range a_propertiesOrdered {

		id = generatesIdentifier("A_PROPERTIES", idx, a_properties.Name)
		map_A_PROPERTIES_Identifiers[a_properties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_PROPERTIES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_properties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_properties.Name))
		initializerStatements += setValueField

	}

	map_A_SOURCE_Identifiers := make(map[*A_SOURCE]string)
	_ = map_A_SOURCE_Identifiers

	a_sourceOrdered := []*A_SOURCE{}
	for a_source := range stage.A_SOURCEs {
		a_sourceOrdered = append(a_sourceOrdered, a_source)
	}
	sort.Slice(a_sourceOrdered[:], func(i, j int) bool {
		return a_sourceOrdered[i].Name < a_sourceOrdered[j].Name
	})
	if len(a_sourceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_source := range a_sourceOrdered {

		id = generatesIdentifier("A_SOURCE", idx, a_source.Name)
		map_A_SOURCE_Identifiers[a_source] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source.SPEC_OBJECT_REF))
		initializerStatements += setValueField

	}

	map_A_SOURCE_SPECIFICATION_Identifiers := make(map[*A_SOURCE_SPECIFICATION]string)
	_ = map_A_SOURCE_SPECIFICATION_Identifiers

	a_source_specificationOrdered := []*A_SOURCE_SPECIFICATION{}
	for a_source_specification := range stage.A_SOURCE_SPECIFICATIONs {
		a_source_specificationOrdered = append(a_source_specificationOrdered, a_source_specification)
	}
	sort.Slice(a_source_specificationOrdered[:], func(i, j int) bool {
		return a_source_specificationOrdered[i].Name < a_source_specificationOrdered[j].Name
	})
	if len(a_source_specificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_source_specification := range a_source_specificationOrdered {

		id = generatesIdentifier("A_SOURCE_SPECIFICATION", idx, a_source_specification.Name)
		map_A_SOURCE_SPECIFICATION_Identifiers[a_source_specification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_SPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source_specification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source_specification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_source_specification.SPECIFICATION_REF))
		initializerStatements += setValueField

	}

	map_A_SPECIFICATIONS_Identifiers := make(map[*A_SPECIFICATIONS]string)
	_ = map_A_SPECIFICATIONS_Identifiers

	a_specificationsOrdered := []*A_SPECIFICATIONS{}
	for a_specifications := range stage.A_SPECIFICATIONSs {
		a_specificationsOrdered = append(a_specificationsOrdered, a_specifications)
	}
	sort.Slice(a_specificationsOrdered[:], func(i, j int) bool {
		return a_specificationsOrdered[i].Name < a_specificationsOrdered[j].Name
	})
	if len(a_specificationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_specifications := range a_specificationsOrdered {

		id = generatesIdentifier("A_SPECIFICATIONS", idx, a_specifications.Name)
		map_A_SPECIFICATIONS_Identifiers[a_specifications] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specifications.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specifications.Name))
		initializerStatements += setValueField

	}

	map_A_SPECIFIED_VALUES_Identifiers := make(map[*A_SPECIFIED_VALUES]string)
	_ = map_A_SPECIFIED_VALUES_Identifiers

	a_specified_valuesOrdered := []*A_SPECIFIED_VALUES{}
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		a_specified_valuesOrdered = append(a_specified_valuesOrdered, a_specified_values)
	}
	sort.Slice(a_specified_valuesOrdered[:], func(i, j int) bool {
		return a_specified_valuesOrdered[i].Name < a_specified_valuesOrdered[j].Name
	})
	if len(a_specified_valuesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_specified_values := range a_specified_valuesOrdered {

		id = generatesIdentifier("A_SPECIFIED_VALUES", idx, a_specified_values.Name)
		map_A_SPECIFIED_VALUES_Identifiers[a_specified_values] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFIED_VALUES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specified_values.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_specified_values.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_ATTRIBUTES_Identifiers := make(map[*A_SPEC_ATTRIBUTES]string)
	_ = map_A_SPEC_ATTRIBUTES_Identifiers

	a_spec_attributesOrdered := []*A_SPEC_ATTRIBUTES{}
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		a_spec_attributesOrdered = append(a_spec_attributesOrdered, a_spec_attributes)
	}
	sort.Slice(a_spec_attributesOrdered[:], func(i, j int) bool {
		return a_spec_attributesOrdered[i].Name < a_spec_attributesOrdered[j].Name
	})
	if len(a_spec_attributesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_attributes := range a_spec_attributesOrdered {

		id = generatesIdentifier("A_SPEC_ATTRIBUTES", idx, a_spec_attributes.Name)
		map_A_SPEC_ATTRIBUTES_Identifiers[a_spec_attributes] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_ATTRIBUTES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_attributes.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_attributes.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_OBJECTS_Identifiers := make(map[*A_SPEC_OBJECTS]string)
	_ = map_A_SPEC_OBJECTS_Identifiers

	a_spec_objectsOrdered := []*A_SPEC_OBJECTS{}
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		a_spec_objectsOrdered = append(a_spec_objectsOrdered, a_spec_objects)
	}
	sort.Slice(a_spec_objectsOrdered[:], func(i, j int) bool {
		return a_spec_objectsOrdered[i].Name < a_spec_objectsOrdered[j].Name
	})
	if len(a_spec_objectsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_objects := range a_spec_objectsOrdered {

		id = generatesIdentifier("A_SPEC_OBJECTS", idx, a_spec_objects.Name)
		map_A_SPEC_OBJECTS_Identifiers[a_spec_objects] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECTS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_objects.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_objects.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATIONS_Identifiers := make(map[*A_SPEC_RELATIONS]string)
	_ = map_A_SPEC_RELATIONS_Identifiers

	a_spec_relationsOrdered := []*A_SPEC_RELATIONS{}
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		a_spec_relationsOrdered = append(a_spec_relationsOrdered, a_spec_relations)
	}
	sort.Slice(a_spec_relationsOrdered[:], func(i, j int) bool {
		return a_spec_relationsOrdered[i].Name < a_spec_relationsOrdered[j].Name
	})
	if len(a_spec_relationsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_relations := range a_spec_relationsOrdered {

		id = generatesIdentifier("A_SPEC_RELATIONS", idx, a_spec_relations.Name)
		map_A_SPEC_RELATIONS_Identifiers[a_spec_relations] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relations.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relations.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_RELATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relations.SPEC_RELATION_REF))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATIONS_1_Identifiers := make(map[*A_SPEC_RELATIONS_1]string)
	_ = map_A_SPEC_RELATIONS_1_Identifiers

	a_spec_relations_1Ordered := []*A_SPEC_RELATIONS_1{}
	for a_spec_relations_1 := range stage.A_SPEC_RELATIONS_1s {
		a_spec_relations_1Ordered = append(a_spec_relations_1Ordered, a_spec_relations_1)
	}
	sort.Slice(a_spec_relations_1Ordered[:], func(i, j int) bool {
		return a_spec_relations_1Ordered[i].Name < a_spec_relations_1Ordered[j].Name
	})
	if len(a_spec_relations_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_relations_1 := range a_spec_relations_1Ordered {

		id = generatesIdentifier("A_SPEC_RELATIONS_1", idx, a_spec_relations_1.Name)
		map_A_SPEC_RELATIONS_1_Identifiers[a_spec_relations_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATIONS_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relations_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relations_1.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_RELATION_GROUPS_Identifiers := make(map[*A_SPEC_RELATION_GROUPS]string)
	_ = map_A_SPEC_RELATION_GROUPS_Identifiers

	a_spec_relation_groupsOrdered := []*A_SPEC_RELATION_GROUPS{}
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groupsOrdered = append(a_spec_relation_groupsOrdered, a_spec_relation_groups)
	}
	sort.Slice(a_spec_relation_groupsOrdered[:], func(i, j int) bool {
		return a_spec_relation_groupsOrdered[i].Name < a_spec_relation_groupsOrdered[j].Name
	})
	if len(a_spec_relation_groupsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_relation_groups := range a_spec_relation_groupsOrdered {

		id = generatesIdentifier("A_SPEC_RELATION_GROUPS", idx, a_spec_relation_groups.Name)
		map_A_SPEC_RELATION_GROUPS_Identifiers[a_spec_relation_groups] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_GROUPS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_groups.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_relation_groups.Name))
		initializerStatements += setValueField

	}

	map_A_SPEC_TYPES_Identifiers := make(map[*A_SPEC_TYPES]string)
	_ = map_A_SPEC_TYPES_Identifiers

	a_spec_typesOrdered := []*A_SPEC_TYPES{}
	for a_spec_types := range stage.A_SPEC_TYPESs {
		a_spec_typesOrdered = append(a_spec_typesOrdered, a_spec_types)
	}
	sort.Slice(a_spec_typesOrdered[:], func(i, j int) bool {
		return a_spec_typesOrdered[i].Name < a_spec_typesOrdered[j].Name
	})
	if len(a_spec_typesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_spec_types := range a_spec_typesOrdered {

		id = generatesIdentifier("A_SPEC_TYPES", idx, a_spec_types.Name)
		map_A_SPEC_TYPES_Identifiers[a_spec_types] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_TYPES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_types.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_spec_types.Name))
		initializerStatements += setValueField

	}

	map_A_THE_HEADER_Identifiers := make(map[*A_THE_HEADER]string)
	_ = map_A_THE_HEADER_Identifiers

	a_the_headerOrdered := []*A_THE_HEADER{}
	for a_the_header := range stage.A_THE_HEADERs {
		a_the_headerOrdered = append(a_the_headerOrdered, a_the_header)
	}
	sort.Slice(a_the_headerOrdered[:], func(i, j int) bool {
		return a_the_headerOrdered[i].Name < a_the_headerOrdered[j].Name
	})
	if len(a_the_headerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_the_header := range a_the_headerOrdered {

		id = generatesIdentifier("A_THE_HEADER", idx, a_the_header.Name)
		map_A_THE_HEADER_Identifiers[a_the_header] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_THE_HEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_the_header.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_the_header.Name))
		initializerStatements += setValueField

	}

	map_A_TOOL_EXTENSIONS_Identifiers := make(map[*A_TOOL_EXTENSIONS]string)
	_ = map_A_TOOL_EXTENSIONS_Identifiers

	a_tool_extensionsOrdered := []*A_TOOL_EXTENSIONS{}
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		a_tool_extensionsOrdered = append(a_tool_extensionsOrdered, a_tool_extensions)
	}
	sort.Slice(a_tool_extensionsOrdered[:], func(i, j int) bool {
		return a_tool_extensionsOrdered[i].Name < a_tool_extensionsOrdered[j].Name
	})
	if len(a_tool_extensionsOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_tool_extensions := range a_tool_extensionsOrdered {

		id = generatesIdentifier("A_TOOL_EXTENSIONS", idx, a_tool_extensions.Name)
		map_A_TOOL_EXTENSIONS_Identifiers[a_tool_extensions] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TOOL_EXTENSIONS")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_tool_extensions.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_tool_extensions.Name))
		initializerStatements += setValueField

	}

	map_A_TYPE_Identifiers := make(map[*A_TYPE]string)
	_ = map_A_TYPE_Identifiers

	a_typeOrdered := []*A_TYPE{}
	for a_type := range stage.A_TYPEs {
		a_typeOrdered = append(a_typeOrdered, a_type)
	}
	sort.Slice(a_typeOrdered[:], func(i, j int) bool {
		return a_typeOrdered[i].Name < a_typeOrdered[j].Name
	})
	if len(a_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type := range a_typeOrdered {

		id = generatesIdentifier("A_TYPE", idx, a_type.Name)
		map_A_TYPE_Identifiers[a_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type.DATATYPE_DEFINITION_DATE_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_1_Identifiers := make(map[*A_TYPE_1]string)
	_ = map_A_TYPE_1_Identifiers

	a_type_1Ordered := []*A_TYPE_1{}
	for a_type_1 := range stage.A_TYPE_1s {
		a_type_1Ordered = append(a_type_1Ordered, a_type_1)
	}
	sort.Slice(a_type_1Ordered[:], func(i, j int) bool {
		return a_type_1Ordered[i].Name < a_type_1Ordered[j].Name
	})
	if len(a_type_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_1 := range a_type_1Ordered {

		id = generatesIdentifier("A_TYPE_1", idx, a_type_1.Name)
		map_A_TYPE_1_Identifiers[a_type_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_1.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_1.RELATION_GROUP_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_10_Identifiers := make(map[*A_TYPE_10]string)
	_ = map_A_TYPE_10_Identifiers

	a_type_10Ordered := []*A_TYPE_10{}
	for a_type_10 := range stage.A_TYPE_10s {
		a_type_10Ordered = append(a_type_10Ordered, a_type_10)
	}
	sort.Slice(a_type_10Ordered[:], func(i, j int) bool {
		return a_type_10Ordered[i].Name < a_type_10Ordered[j].Name
	})
	if len(a_type_10Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_10 := range a_type_10Ordered {

		id = generatesIdentifier("A_TYPE_10", idx, a_type_10.Name)
		map_A_TYPE_10_Identifiers[a_type_10] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_10")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_10.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_10.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_10.SPECIFICATION_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_2_Identifiers := make(map[*A_TYPE_2]string)
	_ = map_A_TYPE_2_Identifiers

	a_type_2Ordered := []*A_TYPE_2{}
	for a_type_2 := range stage.A_TYPE_2s {
		a_type_2Ordered = append(a_type_2Ordered, a_type_2)
	}
	sort.Slice(a_type_2Ordered[:], func(i, j int) bool {
		return a_type_2Ordered[i].Name < a_type_2Ordered[j].Name
	})
	if len(a_type_2Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_2 := range a_type_2Ordered {

		id = generatesIdentifier("A_TYPE_2", idx, a_type_2.Name)
		map_A_TYPE_2_Identifiers[a_type_2] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_2")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_2.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_2.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_2.SPEC_OBJECT_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_3_Identifiers := make(map[*A_TYPE_3]string)
	_ = map_A_TYPE_3_Identifiers

	a_type_3Ordered := []*A_TYPE_3{}
	for a_type_3 := range stage.A_TYPE_3s {
		a_type_3Ordered = append(a_type_3Ordered, a_type_3)
	}
	sort.Slice(a_type_3Ordered[:], func(i, j int) bool {
		return a_type_3Ordered[i].Name < a_type_3Ordered[j].Name
	})
	if len(a_type_3Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_3 := range a_type_3Ordered {

		id = generatesIdentifier("A_TYPE_3", idx, a_type_3.Name)
		map_A_TYPE_3_Identifiers[a_type_3] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_3")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_3.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_3.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_3.DATATYPE_DEFINITION_REAL_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_4_Identifiers := make(map[*A_TYPE_4]string)
	_ = map_A_TYPE_4_Identifiers

	a_type_4Ordered := []*A_TYPE_4{}
	for a_type_4 := range stage.A_TYPE_4s {
		a_type_4Ordered = append(a_type_4Ordered, a_type_4)
	}
	sort.Slice(a_type_4Ordered[:], func(i, j int) bool {
		return a_type_4Ordered[i].Name < a_type_4Ordered[j].Name
	})
	if len(a_type_4Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_4 := range a_type_4Ordered {

		id = generatesIdentifier("A_TYPE_4", idx, a_type_4.Name)
		map_A_TYPE_4_Identifiers[a_type_4] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_4")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_4.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_4.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_4.SPEC_RELATION_TYPE_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_5_Identifiers := make(map[*A_TYPE_5]string)
	_ = map_A_TYPE_5_Identifiers

	a_type_5Ordered := []*A_TYPE_5{}
	for a_type_5 := range stage.A_TYPE_5s {
		a_type_5Ordered = append(a_type_5Ordered, a_type_5)
	}
	sort.Slice(a_type_5Ordered[:], func(i, j int) bool {
		return a_type_5Ordered[i].Name < a_type_5Ordered[j].Name
	})
	if len(a_type_5Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_5 := range a_type_5Ordered {

		id = generatesIdentifier("A_TYPE_5", idx, a_type_5.Name)
		map_A_TYPE_5_Identifiers[a_type_5] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_5")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_5.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_5.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_5.DATATYPE_DEFINITION_INTEGER_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_6_Identifiers := make(map[*A_TYPE_6]string)
	_ = map_A_TYPE_6_Identifiers

	a_type_6Ordered := []*A_TYPE_6{}
	for a_type_6 := range stage.A_TYPE_6s {
		a_type_6Ordered = append(a_type_6Ordered, a_type_6)
	}
	sort.Slice(a_type_6Ordered[:], func(i, j int) bool {
		return a_type_6Ordered[i].Name < a_type_6Ordered[j].Name
	})
	if len(a_type_6Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_6 := range a_type_6Ordered {

		id = generatesIdentifier("A_TYPE_6", idx, a_type_6.Name)
		map_A_TYPE_6_Identifiers[a_type_6] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_6")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_6.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_6.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_6.DATATYPE_DEFINITION_XHTML_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_7_Identifiers := make(map[*A_TYPE_7]string)
	_ = map_A_TYPE_7_Identifiers

	a_type_7Ordered := []*A_TYPE_7{}
	for a_type_7 := range stage.A_TYPE_7s {
		a_type_7Ordered = append(a_type_7Ordered, a_type_7)
	}
	sort.Slice(a_type_7Ordered[:], func(i, j int) bool {
		return a_type_7Ordered[i].Name < a_type_7Ordered[j].Name
	})
	if len(a_type_7Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_7 := range a_type_7Ordered {

		id = generatesIdentifier("A_TYPE_7", idx, a_type_7.Name)
		map_A_TYPE_7_Identifiers[a_type_7] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_7")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_7.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_7.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_7.DATATYPE_DEFINITION_BOOLEAN_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_8_Identifiers := make(map[*A_TYPE_8]string)
	_ = map_A_TYPE_8_Identifiers

	a_type_8Ordered := []*A_TYPE_8{}
	for a_type_8 := range stage.A_TYPE_8s {
		a_type_8Ordered = append(a_type_8Ordered, a_type_8)
	}
	sort.Slice(a_type_8Ordered[:], func(i, j int) bool {
		return a_type_8Ordered[i].Name < a_type_8Ordered[j].Name
	})
	if len(a_type_8Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_8 := range a_type_8Ordered {

		id = generatesIdentifier("A_TYPE_8", idx, a_type_8.Name)
		map_A_TYPE_8_Identifiers[a_type_8] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_8")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_8.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_8.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_8.DATATYPE_DEFINITION_STRING_REF))
		initializerStatements += setValueField

	}

	map_A_TYPE_9_Identifiers := make(map[*A_TYPE_9]string)
	_ = map_A_TYPE_9_Identifiers

	a_type_9Ordered := []*A_TYPE_9{}
	for a_type_9 := range stage.A_TYPE_9s {
		a_type_9Ordered = append(a_type_9Ordered, a_type_9)
	}
	sort.Slice(a_type_9Ordered[:], func(i, j int) bool {
		return a_type_9Ordered[i].Name < a_type_9Ordered[j].Name
	})
	if len(a_type_9Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_type_9 := range a_type_9Ordered {

		id = generatesIdentifier("A_TYPE_9", idx, a_type_9.Name)
		map_A_TYPE_9_Identifiers[a_type_9] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TYPE_9")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_type_9.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_9.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_type_9.DATATYPE_DEFINITION_ENUMERATION_REF))
		initializerStatements += setValueField

	}

	map_A_VALUES_Identifiers := make(map[*A_VALUES]string)
	_ = map_A_VALUES_Identifiers

	a_valuesOrdered := []*A_VALUES{}
	for a_values := range stage.A_VALUESs {
		a_valuesOrdered = append(a_valuesOrdered, a_values)
	}
	sort.Slice(a_valuesOrdered[:], func(i, j int) bool {
		return a_valuesOrdered[i].Name < a_valuesOrdered[j].Name
	})
	if len(a_valuesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_values := range a_valuesOrdered {

		id = generatesIdentifier("A_VALUES", idx, a_values.Name)
		map_A_VALUES_Identifiers[a_values] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_VALUES")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_values.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_values.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ENUM_VALUE_REF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_values.ENUM_VALUE_REF))
		initializerStatements += setValueField

	}

	map_A_VALUES_1_Identifiers := make(map[*A_VALUES_1]string)
	_ = map_A_VALUES_1_Identifiers

	a_values_1Ordered := []*A_VALUES_1{}
	for a_values_1 := range stage.A_VALUES_1s {
		a_values_1Ordered = append(a_values_1Ordered, a_values_1)
	}
	sort.Slice(a_values_1Ordered[:], func(i, j int) bool {
		return a_values_1Ordered[i].Name < a_values_1Ordered[j].Name
	})
	if len(a_values_1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, a_values_1 := range a_values_1Ordered {

		id = generatesIdentifier("A_VALUES_1", idx, a_values_1.Name)
		map_A_VALUES_1_Identifiers[a_values_1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_VALUES_1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_values_1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(a_values_1.Name))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_BOOLEAN_Identifiers := make(map[*DATATYPE_DEFINITION_BOOLEAN]string)
	_ = map_DATATYPE_DEFINITION_BOOLEAN_Identifiers

	datatype_definition_booleanOrdered := []*DATATYPE_DEFINITION_BOOLEAN{}
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		datatype_definition_booleanOrdered = append(datatype_definition_booleanOrdered, datatype_definition_boolean)
	}
	sort.Slice(datatype_definition_booleanOrdered[:], func(i, j int) bool {
		return datatype_definition_booleanOrdered[i].Name < datatype_definition_booleanOrdered[j].Name
	})
	if len(datatype_definition_booleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_boolean := range datatype_definition_booleanOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", idx, datatype_definition_boolean.Name)
		map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[datatype_definition_boolean] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_BOOLEAN")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_boolean.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_DATE_Identifiers := make(map[*DATATYPE_DEFINITION_DATE]string)
	_ = map_DATATYPE_DEFINITION_DATE_Identifiers

	datatype_definition_dateOrdered := []*DATATYPE_DEFINITION_DATE{}
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		datatype_definition_dateOrdered = append(datatype_definition_dateOrdered, datatype_definition_date)
	}
	sort.Slice(datatype_definition_dateOrdered[:], func(i, j int) bool {
		return datatype_definition_dateOrdered[i].Name < datatype_definition_dateOrdered[j].Name
	})
	if len(datatype_definition_dateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_date := range datatype_definition_dateOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", idx, datatype_definition_date.Name)
		map_DATATYPE_DEFINITION_DATE_Identifiers[datatype_definition_date] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_DATE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_date.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_ENUMERATION_Identifiers := make(map[*DATATYPE_DEFINITION_ENUMERATION]string)
	_ = map_DATATYPE_DEFINITION_ENUMERATION_Identifiers

	datatype_definition_enumerationOrdered := []*DATATYPE_DEFINITION_ENUMERATION{}
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		datatype_definition_enumerationOrdered = append(datatype_definition_enumerationOrdered, datatype_definition_enumeration)
	}
	sort.Slice(datatype_definition_enumerationOrdered[:], func(i, j int) bool {
		return datatype_definition_enumerationOrdered[i].Name < datatype_definition_enumerationOrdered[j].Name
	})
	if len(datatype_definition_enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", idx, datatype_definition_enumeration.Name)
		map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[datatype_definition_enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_ENUMERATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LONG_NAME))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_INTEGER_Identifiers := make(map[*DATATYPE_DEFINITION_INTEGER]string)
	_ = map_DATATYPE_DEFINITION_INTEGER_Identifiers

	datatype_definition_integerOrdered := []*DATATYPE_DEFINITION_INTEGER{}
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integerOrdered = append(datatype_definition_integerOrdered, datatype_definition_integer)
	}
	sort.Slice(datatype_definition_integerOrdered[:], func(i, j int) bool {
		return datatype_definition_integerOrdered[i].Name < datatype_definition_integerOrdered[j].Name
	})
	if len(datatype_definition_integerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_integer := range datatype_definition_integerOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", idx, datatype_definition_integer.Name)
		map_DATATYPE_DEFINITION_INTEGER_Identifiers[datatype_definition_integer] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_INTEGER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_integer.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MAX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MIN")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MIN))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_REAL_Identifiers := make(map[*DATATYPE_DEFINITION_REAL]string)
	_ = map_DATATYPE_DEFINITION_REAL_Identifiers

	datatype_definition_realOrdered := []*DATATYPE_DEFINITION_REAL{}
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		datatype_definition_realOrdered = append(datatype_definition_realOrdered, datatype_definition_real)
	}
	sort.Slice(datatype_definition_realOrdered[:], func(i, j int) bool {
		return datatype_definition_realOrdered[i].Name < datatype_definition_realOrdered[j].Name
	})
	if len(datatype_definition_realOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_real := range datatype_definition_realOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", idx, datatype_definition_real.Name)
		map_DATATYPE_DEFINITION_REAL_Identifiers[datatype_definition_real] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_REAL")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_real.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ACCURACY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_real.ACCURACY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MAX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MIN")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MIN))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_STRING_Identifiers := make(map[*DATATYPE_DEFINITION_STRING]string)
	_ = map_DATATYPE_DEFINITION_STRING_Identifiers

	datatype_definition_stringOrdered := []*DATATYPE_DEFINITION_STRING{}
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_stringOrdered = append(datatype_definition_stringOrdered, datatype_definition_string)
	}
	sort.Slice(datatype_definition_stringOrdered[:], func(i, j int) bool {
		return datatype_definition_stringOrdered[i].Name < datatype_definition_stringOrdered[j].Name
	})
	if len(datatype_definition_stringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_string := range datatype_definition_stringOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", idx, datatype_definition_string.Name)
		map_DATATYPE_DEFINITION_STRING_Identifiers[datatype_definition_string] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_STRING")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_string.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LONG_NAME))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MAX_LENGTH")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_string.MAX_LENGTH))
		initializerStatements += setValueField

	}

	map_DATATYPE_DEFINITION_XHTML_Identifiers := make(map[*DATATYPE_DEFINITION_XHTML]string)
	_ = map_DATATYPE_DEFINITION_XHTML_Identifiers

	datatype_definition_xhtmlOrdered := []*DATATYPE_DEFINITION_XHTML{}
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		datatype_definition_xhtmlOrdered = append(datatype_definition_xhtmlOrdered, datatype_definition_xhtml)
	}
	sort.Slice(datatype_definition_xhtmlOrdered[:], func(i, j int) bool {
		return datatype_definition_xhtmlOrdered[i].Name < datatype_definition_xhtmlOrdered[j].Name
	})
	if len(datatype_definition_xhtmlOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", idx, datatype_definition_xhtml.Name)
		map_DATATYPE_DEFINITION_XHTML_Identifiers[datatype_definition_xhtml] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_XHTML")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_xhtml.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LONG_NAME))
		initializerStatements += setValueField

	}

	map_EMBEDDED_VALUE_Identifiers := make(map[*EMBEDDED_VALUE]string)
	_ = map_EMBEDDED_VALUE_Identifiers

	embedded_valueOrdered := []*EMBEDDED_VALUE{}
	for embedded_value := range stage.EMBEDDED_VALUEs {
		embedded_valueOrdered = append(embedded_valueOrdered, embedded_value)
	}
	sort.Slice(embedded_valueOrdered[:], func(i, j int) bool {
		return embedded_valueOrdered[i].Name < embedded_valueOrdered[j].Name
	})
	if len(embedded_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, embedded_value := range embedded_valueOrdered {

		id = generatesIdentifier("EMBEDDED_VALUE", idx, embedded_value.Name)
		map_EMBEDDED_VALUE_Identifiers[embedded_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EMBEDDED_VALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embedded_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embedded_value.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "KEY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", embedded_value.KEY))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OTHER_CONTENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(embedded_value.OTHER_CONTENT))
		initializerStatements += setValueField

	}

	map_ENUM_VALUE_Identifiers := make(map[*ENUM_VALUE]string)
	_ = map_ENUM_VALUE_Identifiers

	enum_valueOrdered := []*ENUM_VALUE{}
	for enum_value := range stage.ENUM_VALUEs {
		enum_valueOrdered = append(enum_valueOrdered, enum_value)
	}
	sort.Slice(enum_valueOrdered[:], func(i, j int) bool {
		return enum_valueOrdered[i].Name < enum_valueOrdered[j].Name
	})
	if len(enum_valueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, enum_value := range enum_valueOrdered {

		id = generatesIdentifier("ENUM_VALUE", idx, enum_value.Name)
		map_ENUM_VALUE_Identifiers[enum_value] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ENUM_VALUE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enum_value.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enum_value.LONG_NAME))
		initializerStatements += setValueField

	}

	map_RELATION_GROUP_Identifiers := make(map[*RELATION_GROUP]string)
	_ = map_RELATION_GROUP_Identifiers

	relation_groupOrdered := []*RELATION_GROUP{}
	for relation_group := range stage.RELATION_GROUPs {
		relation_groupOrdered = append(relation_groupOrdered, relation_group)
	}
	sort.Slice(relation_groupOrdered[:], func(i, j int) bool {
		return relation_groupOrdered[i].Name < relation_groupOrdered[j].Name
	})
	if len(relation_groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, relation_group := range relation_groupOrdered {

		id = generatesIdentifier("RELATION_GROUP", idx, relation_group.Name)
		map_RELATION_GROUP_Identifiers[relation_group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group.LONG_NAME))
		initializerStatements += setValueField

	}

	map_RELATION_GROUP_TYPE_Identifiers := make(map[*RELATION_GROUP_TYPE]string)
	_ = map_RELATION_GROUP_TYPE_Identifiers

	relation_group_typeOrdered := []*RELATION_GROUP_TYPE{}
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		relation_group_typeOrdered = append(relation_group_typeOrdered, relation_group_type)
	}
	sort.Slice(relation_group_typeOrdered[:], func(i, j int) bool {
		return relation_group_typeOrdered[i].Name < relation_group_typeOrdered[j].Name
	})
	if len(relation_group_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, relation_group_type := range relation_group_typeOrdered {

		id = generatesIdentifier("RELATION_GROUP_TYPE", idx, relation_group_type.Name)
		map_RELATION_GROUP_TYPE_Identifiers[relation_group_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(relation_group_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_REQ_IF_Identifiers := make(map[*REQ_IF]string)
	_ = map_REQ_IF_Identifiers

	req_ifOrdered := []*REQ_IF{}
	for req_if := range stage.REQ_IFs {
		req_ifOrdered = append(req_ifOrdered, req_if)
	}
	sort.Slice(req_ifOrdered[:], func(i, j int) bool {
		return req_ifOrdered[i].Name < req_ifOrdered[j].Name
	})
	if len(req_ifOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, req_if := range req_ifOrdered {

		id = generatesIdentifier("REQ_IF", idx, req_if.Name)
		map_REQ_IF_Identifiers[req_if] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lang")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if.Lang))
		initializerStatements += setValueField

	}

	map_REQ_IF_CONTENT_Identifiers := make(map[*REQ_IF_CONTENT]string)
	_ = map_REQ_IF_CONTENT_Identifiers

	req_if_contentOrdered := []*REQ_IF_CONTENT{}
	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_contentOrdered = append(req_if_contentOrdered, req_if_content)
	}
	sort.Slice(req_if_contentOrdered[:], func(i, j int) bool {
		return req_if_contentOrdered[i].Name < req_if_contentOrdered[j].Name
	})
	if len(req_if_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, req_if_content := range req_if_contentOrdered {

		id = generatesIdentifier("REQ_IF_CONTENT", idx, req_if_content.Name)
		map_REQ_IF_CONTENT_Identifiers[req_if_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_content.Name))
		initializerStatements += setValueField

	}

	map_REQ_IF_HEADER_Identifiers := make(map[*REQ_IF_HEADER]string)
	_ = map_REQ_IF_HEADER_Identifiers

	req_if_headerOrdered := []*REQ_IF_HEADER{}
	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_headerOrdered = append(req_if_headerOrdered, req_if_header)
	}
	sort.Slice(req_if_headerOrdered[:], func(i, j int) bool {
		return req_if_headerOrdered[i].Name < req_if_headerOrdered[j].Name
	})
	if len(req_if_headerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, req_if_header := range req_if_headerOrdered {

		id = generatesIdentifier("REQ_IF_HEADER", idx, req_if_header.Name)
		map_REQ_IF_HEADER_Identifiers[req_if_header] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_HEADER")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_header.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "COMMENT")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.COMMENT))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CREATION_TIME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.CREATION_TIME))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REPOSITORY_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REPOSITORY_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQ_IF_TOOL_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_TOOL_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "REQ_IF_VERSION")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_VERSION))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SOURCE_TOOL_ID")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.SOURCE_TOOL_ID))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TITLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_header.TITLE))
		initializerStatements += setValueField

	}

	map_REQ_IF_TOOL_EXTENSION_Identifiers := make(map[*REQ_IF_TOOL_EXTENSION]string)
	_ = map_REQ_IF_TOOL_EXTENSION_Identifiers

	req_if_tool_extensionOrdered := []*REQ_IF_TOOL_EXTENSION{}
	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extensionOrdered = append(req_if_tool_extensionOrdered, req_if_tool_extension)
	}
	sort.Slice(req_if_tool_extensionOrdered[:], func(i, j int) bool {
		return req_if_tool_extensionOrdered[i].Name < req_if_tool_extensionOrdered[j].Name
	})
	if len(req_if_tool_extensionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, req_if_tool_extension := range req_if_tool_extensionOrdered {

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", idx, req_if_tool_extension.Name)
		map_REQ_IF_TOOL_EXTENSION_Identifiers[req_if_tool_extension] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_TOOL_EXTENSION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_tool_extension.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(req_if_tool_extension.Name))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_Identifiers := make(map[*SPECIFICATION]string)
	_ = map_SPECIFICATION_Identifiers

	specificationOrdered := []*SPECIFICATION{}
	for specification := range stage.SPECIFICATIONs {
		specificationOrdered = append(specificationOrdered, specification)
	}
	sort.Slice(specificationOrdered[:], func(i, j int) bool {
		return specificationOrdered[i].Name < specificationOrdered[j].Name
	})
	if len(specificationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specification := range specificationOrdered {

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPECIFICATION_TYPE_Identifiers := make(map[*SPECIFICATION_TYPE]string)
	_ = map_SPECIFICATION_TYPE_Identifiers

	specification_typeOrdered := []*SPECIFICATION_TYPE{}
	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_typeOrdered = append(specification_typeOrdered, specification_type)
	}
	sort.Slice(specification_typeOrdered[:], func(i, j int) bool {
		return specification_typeOrdered[i].Name < specification_typeOrdered[j].Name
	})
	if len(specification_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, specification_type := range specification_typeOrdered {

		id = generatesIdentifier("SPECIFICATION_TYPE", idx, specification_type.Name)
		map_SPECIFICATION_TYPE_Identifiers[specification_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(specification_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_HIERARCHY_Identifiers := make(map[*SPEC_HIERARCHY]string)
	_ = map_SPEC_HIERARCHY_Identifiers

	spec_hierarchyOrdered := []*SPEC_HIERARCHY{}
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchyOrdered = append(spec_hierarchyOrdered, spec_hierarchy)
	}
	sort.Slice(spec_hierarchyOrdered[:], func(i, j int) bool {
		return spec_hierarchyOrdered[i].Name < spec_hierarchyOrdered[j].Name
	})
	if len(spec_hierarchyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spec_hierarchy := range spec_hierarchyOrdered {

		id = generatesIdentifier("SPEC_HIERARCHY", idx, spec_hierarchy.Name)
		map_SPEC_HIERARCHY_Identifiers[spec_hierarchy] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_HIERARCHY")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_hierarchy.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_EDITABLE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_EDITABLE))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IS_TABLE_INTERNAL")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_TABLE_INTERNAL))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_OBJECT_Identifiers := make(map[*SPEC_OBJECT]string)
	_ = map_SPEC_OBJECT_Identifiers

	spec_objectOrdered := []*SPEC_OBJECT{}
	for spec_object := range stage.SPEC_OBJECTs {
		spec_objectOrdered = append(spec_objectOrdered, spec_object)
	}
	sort.Slice(spec_objectOrdered[:], func(i, j int) bool {
		return spec_objectOrdered[i].Name < spec_objectOrdered[j].Name
	})
	if len(spec_objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spec_object := range spec_objectOrdered {

		id = generatesIdentifier("SPEC_OBJECT", idx, spec_object.Name)
		map_SPEC_OBJECT_Identifiers[spec_object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_OBJECT_TYPE_Identifiers := make(map[*SPEC_OBJECT_TYPE]string)
	_ = map_SPEC_OBJECT_TYPE_Identifiers

	spec_object_typeOrdered := []*SPEC_OBJECT_TYPE{}
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_typeOrdered = append(spec_object_typeOrdered, spec_object_type)
	}
	sort.Slice(spec_object_typeOrdered[:], func(i, j int) bool {
		return spec_object_typeOrdered[i].Name < spec_object_typeOrdered[j].Name
	})
	if len(spec_object_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spec_object_type := range spec_object_typeOrdered {

		id = generatesIdentifier("SPEC_OBJECT_TYPE", idx, spec_object_type.Name)
		map_SPEC_OBJECT_TYPE_Identifiers[spec_object_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_object_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_RELATION_Identifiers := make(map[*SPEC_RELATION]string)
	_ = map_SPEC_RELATION_Identifiers

	spec_relationOrdered := []*SPEC_RELATION{}
	for spec_relation := range stage.SPEC_RELATIONs {
		spec_relationOrdered = append(spec_relationOrdered, spec_relation)
	}
	sort.Slice(spec_relationOrdered[:], func(i, j int) bool {
		return spec_relationOrdered[i].Name < spec_relationOrdered[j].Name
	})
	if len(spec_relationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spec_relation := range spec_relationOrdered {

		id = generatesIdentifier("SPEC_RELATION", idx, spec_relation.Name)
		map_SPEC_RELATION_Identifiers[spec_relation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation.LONG_NAME))
		initializerStatements += setValueField

	}

	map_SPEC_RELATION_TYPE_Identifiers := make(map[*SPEC_RELATION_TYPE]string)
	_ = map_SPEC_RELATION_TYPE_Identifiers

	spec_relation_typeOrdered := []*SPEC_RELATION_TYPE{}
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		spec_relation_typeOrdered = append(spec_relation_typeOrdered, spec_relation_type)
	}
	sort.Slice(spec_relation_typeOrdered[:], func(i, j int) bool {
		return spec_relation_typeOrdered[i].Name < spec_relation_typeOrdered[j].Name
	})
	if len(spec_relation_typeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, spec_relation_type := range spec_relation_typeOrdered {

		id = generatesIdentifier("SPEC_RELATION_TYPE", idx, spec_relation_type.Name)
		map_SPEC_RELATION_TYPE_Identifiers[spec_relation_type] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION_TYPE")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation_type.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DESC")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.DESC))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDENTIFIER")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.IDENTIFIER))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LAST_CHANGE")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LAST_CHANGE))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LONG_NAME")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LONG_NAME))
		initializerStatements += setValueField

	}

	map_XHTML_CONTENT_Identifiers := make(map[*XHTML_CONTENT]string)
	_ = map_XHTML_CONTENT_Identifiers

	xhtml_contentOrdered := []*XHTML_CONTENT{}
	for xhtml_content := range stage.XHTML_CONTENTs {
		xhtml_contentOrdered = append(xhtml_contentOrdered, xhtml_content)
	}
	sort.Slice(xhtml_contentOrdered[:], func(i, j int) bool {
		return xhtml_contentOrdered[i].Name < xhtml_contentOrdered[j].Name
	})
	if len(xhtml_contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xhtml_content := range xhtml_contentOrdered {

		id = generatesIdentifier("XHTML_CONTENT", idx, xhtml_content.Name)
		map_XHTML_CONTENT_Identifiers[xhtml_content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XHTML_CONTENT")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xhtml_content.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, alternative_id := range alternative_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ALTERNATIVE_ID", idx, alternative_id.Name)
		map_ALTERNATIVE_ID_Identifiers[alternative_id] = id

		// Initialisation of values
	}

	for idx, attribute_definition_boolean := range attribute_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_BOOLEAN", idx, attribute_definition_boolean.Name)
		map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[attribute_definition_boolean] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_1 := range attribute_definition_boolean.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_1_Identifiers[_a_default_value_1])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_7 := range attribute_definition_boolean.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_7_Identifiers[_a_type_7])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_date := range attribute_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_DATE", idx, attribute_definition_date.Name)
		map_ATTRIBUTE_DEFINITION_DATE_Identifiers[attribute_definition_date] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_date.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_3 := range attribute_definition_date.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_3_Identifiers[_a_default_value_3])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type := range attribute_definition_date.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_Identifiers[_a_type])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_ENUMERATION", idx, attribute_definition_enumeration.Name)
		map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[attribute_definition_enumeration] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_5 := range attribute_definition_enumeration.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_5_Identifiers[_a_default_value_5])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_9 := range attribute_definition_enumeration.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_9_Identifiers[_a_type_9])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_integer := range attribute_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_INTEGER", idx, attribute_definition_integer.Name)
		map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[attribute_definition_integer] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_integer.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_6 := range attribute_definition_integer.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_6_Identifiers[_a_default_value_6])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_5 := range attribute_definition_integer.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_5_Identifiers[_a_type_5])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_real := range attribute_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_REAL", idx, attribute_definition_real.Name)
		map_ATTRIBUTE_DEFINITION_REAL_Identifiers[attribute_definition_real] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_real.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_4 := range attribute_definition_real.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_4_Identifiers[_a_default_value_4])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_3 := range attribute_definition_real.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_3_Identifiers[_a_type_3])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_string := range attribute_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_STRING", idx, attribute_definition_string.Name)
		map_ATTRIBUTE_DEFINITION_STRING_Identifiers[attribute_definition_string] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_string.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value := range attribute_definition_string.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_Identifiers[_a_default_value])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_8 := range attribute_definition_string.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_8_Identifiers[_a_type_8])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_DEFINITION_XHTML", idx, attribute_definition_xhtml.Name)
		map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[attribute_definition_xhtml] = id

		// Initialisation of values
		for _, _a_alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_default_value_2 := range attribute_definition_xhtml.DEFAULT_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFAULT_VALUE_2_Identifiers[_a_default_value_2])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_6 := range attribute_definition_xhtml.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_6_Identifiers[_a_type_6])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_boolean := range attribute_value_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_BOOLEAN", idx, attribute_value_boolean.Name)
		map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[attribute_value_boolean] = id

		// Initialisation of values
		for _, _a_definition_4 := range attribute_value_boolean.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_4_Identifiers[_a_definition_4])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_date := range attribute_value_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_DATE", idx, attribute_value_date.Name)
		map_ATTRIBUTE_VALUE_DATE_Identifiers[attribute_value_date] = id

		// Initialisation of values
		for _, _a_definition_2 := range attribute_value_date.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_2_Identifiers[_a_definition_2])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_enumeration := range attribute_value_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_ENUMERATION", idx, attribute_value_enumeration.Name)
		map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[attribute_value_enumeration] = id

		// Initialisation of values
		for _, _a_definition_5 := range attribute_value_enumeration.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_5_Identifiers[_a_definition_5])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_values := range attribute_value_enumeration.VALUES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_VALUES_Identifiers[_a_values])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_integer := range attribute_value_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_INTEGER", idx, attribute_value_integer.Name)
		map_ATTRIBUTE_VALUE_INTEGER_Identifiers[attribute_value_integer] = id

		// Initialisation of values
		for _, _a_definition := range attribute_value_integer.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_Identifiers[_a_definition])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_real := range attribute_value_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_REAL", idx, attribute_value_real.Name)
		map_ATTRIBUTE_VALUE_REAL_Identifiers[attribute_value_real] = id

		// Initialisation of values
		for _, _a_definition_6 := range attribute_value_real.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_6_Identifiers[_a_definition_6])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_string := range attribute_value_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_STRING", idx, attribute_value_string.Name)
		map_ATTRIBUTE_VALUE_STRING_Identifiers[attribute_value_string] = id

		// Initialisation of values
		for _, _a_definition_3 := range attribute_value_string.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_3_Identifiers[_a_definition_3])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute_value_xhtml := range attribute_value_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ATTRIBUTE_VALUE_XHTML", idx, attribute_value_xhtml.Name)
		map_ATTRIBUTE_VALUE_XHTML_Identifiers[attribute_value_xhtml] = id

		// Initialisation of values
		for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[_xhtml_content])
			pointersInitializesStatements += setPointerField
		}

		for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_ORIGINAL_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XHTML_CONTENT_Identifiers[_xhtml_content])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_definition_1 := range attribute_value_xhtml.DEFINITION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DEFINITION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DEFINITION_1_Identifiers[_a_definition_1])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_alternative_id := range a_alternative_idOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_ALTERNATIVE_ID", idx, a_alternative_id.Name)
		map_A_ALTERNATIVE_ID_Identifiers[a_alternative_id] = id

		// Initialisation of values
		for _, _alternative_id := range a_alternative_id.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ALTERNATIVE_ID_Identifiers[_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_children := range a_childrenOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_CHILDREN", idx, a_children.Name)
		map_A_CHILDREN_Identifiers[a_children] = id

		// Initialisation of values
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_HIERARCHY")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_HIERARCHY_Identifiers[_spec_hierarchy])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_core_content := range a_core_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_CORE_CONTENT", idx, a_core_content.Name)
		map_A_CORE_CONTENT_Identifiers[a_core_content] = id

		// Initialisation of values
		for _, _req_if_content := range a_core_content.REQ_IF_CONTENT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_CONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_CONTENT_Identifiers[_req_if_content])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_datatypes := range a_datatypesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DATATYPES", idx, a_datatypes.Name)
		map_A_DATATYPES_Identifiers[a_datatypes] = id

		// Initialisation of values
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[_datatype_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_DATE_Identifiers[_datatype_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[_datatype_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_INTEGER_Identifiers[_datatype_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_REAL_Identifiers[_datatype_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_STRING_Identifiers[_datatype_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DATATYPE_DEFINITION_XHTML_Identifiers[_datatype_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value := range a_default_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE", idx, a_default_value.Name)
		map_A_DEFAULT_VALUE_Identifiers[a_default_value] = id

		// Initialisation of values
		for _, _attribute_value_string := range a_default_value.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_1 := range a_default_value_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_1", idx, a_default_value_1.Name)
		map_A_DEFAULT_VALUE_1_Identifiers[a_default_value_1] = id

		// Initialisation of values
		for _, _attribute_value_boolean := range a_default_value_1.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_2 := range a_default_value_2Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_2", idx, a_default_value_2.Name)
		map_A_DEFAULT_VALUE_2_Identifiers[a_default_value_2] = id

		// Initialisation of values
		for _, _attribute_value_xhtml := range a_default_value_2.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_3 := range a_default_value_3Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_3", idx, a_default_value_3.Name)
		map_A_DEFAULT_VALUE_3_Identifiers[a_default_value_3] = id

		// Initialisation of values
		for _, _attribute_value_date := range a_default_value_3.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_4 := range a_default_value_4Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_4", idx, a_default_value_4.Name)
		map_A_DEFAULT_VALUE_4_Identifiers[a_default_value_4] = id

		// Initialisation of values
		for _, _attribute_value_real := range a_default_value_4.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_5 := range a_default_value_5Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_5", idx, a_default_value_5.Name)
		map_A_DEFAULT_VALUE_5_Identifiers[a_default_value_5] = id

		// Initialisation of values
		for _, _attribute_value_enumeration := range a_default_value_5.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_default_value_6 := range a_default_value_6Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFAULT_VALUE_6", idx, a_default_value_6.Name)
		map_A_DEFAULT_VALUE_6_Identifiers[a_default_value_6] = id

		// Initialisation of values
		for _, _attribute_value_integer := range a_default_value_6.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_definition := range a_definitionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION", idx, a_definition.Name)
		map_A_DEFINITION_Identifiers[a_definition] = id

		// Initialisation of values
	}

	for idx, a_definition_1 := range a_definition_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_1", idx, a_definition_1.Name)
		map_A_DEFINITION_1_Identifiers[a_definition_1] = id

		// Initialisation of values
	}

	for idx, a_definition_2 := range a_definition_2Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_2", idx, a_definition_2.Name)
		map_A_DEFINITION_2_Identifiers[a_definition_2] = id

		// Initialisation of values
	}

	for idx, a_definition_3 := range a_definition_3Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_3", idx, a_definition_3.Name)
		map_A_DEFINITION_3_Identifiers[a_definition_3] = id

		// Initialisation of values
	}

	for idx, a_definition_4 := range a_definition_4Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_4", idx, a_definition_4.Name)
		map_A_DEFINITION_4_Identifiers[a_definition_4] = id

		// Initialisation of values
	}

	for idx, a_definition_5 := range a_definition_5Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_5", idx, a_definition_5.Name)
		map_A_DEFINITION_5_Identifiers[a_definition_5] = id

		// Initialisation of values
	}

	for idx, a_definition_6 := range a_definition_6Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_DEFINITION_6", idx, a_definition_6.Name)
		map_A_DEFINITION_6_Identifiers[a_definition_6] = id

		// Initialisation of values
	}

	for idx, a_editable_atts := range a_editable_attsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_EDITABLE_ATTS", idx, a_editable_atts.Name)
		map_A_EDITABLE_ATTS_Identifiers[a_editable_atts] = id

		// Initialisation of values
	}

	for idx, a_object := range a_objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_OBJECT", idx, a_object.Name)
		map_A_OBJECT_Identifiers[a_object] = id

		// Initialisation of values
	}

	for idx, a_properties := range a_propertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_PROPERTIES", idx, a_properties.Name)
		map_A_PROPERTIES_Identifiers[a_properties] = id

		// Initialisation of values
		for _, _embedded_value := range a_properties.EMBEDDED_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EMBEDDED_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_EMBEDDED_VALUE_Identifiers[_embedded_value])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_source := range a_sourceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SOURCE", idx, a_source.Name)
		map_A_SOURCE_Identifiers[a_source] = id

		// Initialisation of values
	}

	for idx, a_source_specification := range a_source_specificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SOURCE_SPECIFICATION", idx, a_source_specification.Name)
		map_A_SOURCE_SPECIFICATION_Identifiers[a_source_specification] = id

		// Initialisation of values
	}

	for idx, a_specifications := range a_specificationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPECIFICATIONS", idx, a_specifications.Name)
		map_A_SPECIFICATIONS_Identifiers[a_specifications] = id

		// Initialisation of values
		for _, _specification := range a_specifications.SPECIFICATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_Identifiers[_specification])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_specified_values := range a_specified_valuesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPECIFIED_VALUES", idx, a_specified_values.Name)
		map_A_SPECIFIED_VALUES_Identifiers[a_specified_values] = id

		// Initialisation of values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ENUM_VALUE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ENUM_VALUE_Identifiers[_enum_value])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_spec_attributes := range a_spec_attributesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_ATTRIBUTES", idx, a_spec_attributes.Name)
		map_A_SPEC_ATTRIBUTES_Identifiers[a_spec_attributes] = id

		// Initialisation of values
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_BOOLEAN_Identifiers[_attribute_definition_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_DATE_Identifiers[_attribute_definition_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_ENUMERATION_Identifiers[_attribute_definition_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_INTEGER_Identifiers[_attribute_definition_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_REAL_Identifiers[_attribute_definition_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_STRING_Identifiers[_attribute_definition_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_DEFINITION_XHTML_Identifiers[_attribute_definition_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_spec_objects := range a_spec_objectsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_OBJECTS", idx, a_spec_objects.Name)
		map_A_SPEC_OBJECTS_Identifiers[a_spec_objects] = id

		// Initialisation of values
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_Identifiers[_spec_object])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_spec_relations := range a_spec_relationsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATIONS", idx, a_spec_relations.Name)
		map_A_SPEC_RELATIONS_Identifiers[a_spec_relations] = id

		// Initialisation of values
	}

	for idx, a_spec_relations_1 := range a_spec_relations_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATIONS_1", idx, a_spec_relations_1.Name)
		map_A_SPEC_RELATIONS_1_Identifiers[a_spec_relations_1] = id

		// Initialisation of values
		for _, _spec_relation := range a_spec_relations_1.SPEC_RELATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_Identifiers[_spec_relation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_spec_relation_groups := range a_spec_relation_groupsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_RELATION_GROUPS", idx, a_spec_relation_groups.Name)
		map_A_SPEC_RELATION_GROUPS_Identifiers[a_spec_relation_groups] = id

		// Initialisation of values
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATION_GROUP")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_Identifiers[_relation_group])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_spec_types := range a_spec_typesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_SPEC_TYPES", idx, a_spec_types.Name)
		map_A_SPEC_TYPES_Identifiers[a_spec_types] = id

		// Initialisation of values
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RELATION_GROUP_TYPE_Identifiers[_relation_group_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_OBJECT_TYPE_Identifiers[_spec_object_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPEC_RELATION_TYPE_Identifiers[_spec_relation_type])
			pointersInitializesStatements += setPointerField
		}

		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SPECIFICATION_TYPE_Identifiers[_specification_type])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_the_header := range a_the_headerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_THE_HEADER", idx, a_the_header.Name)
		map_A_THE_HEADER_Identifiers[a_the_header] = id

		// Initialisation of values
		for _, _req_if_header := range a_the_header.REQ_IF_HEADER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_HEADER_Identifiers[_req_if_header])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_tool_extensions := range a_tool_extensionsOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TOOL_EXTENSIONS", idx, a_tool_extensions.Name)
		map_A_TOOL_EXTENSIONS_Identifiers[a_tool_extensions] = id

		// Initialisation of values
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "REQ_IF_TOOL_EXTENSION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_REQ_IF_TOOL_EXTENSION_Identifiers[_req_if_tool_extension])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, a_type := range a_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE", idx, a_type.Name)
		map_A_TYPE_Identifiers[a_type] = id

		// Initialisation of values
	}

	for idx, a_type_1 := range a_type_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_1", idx, a_type_1.Name)
		map_A_TYPE_1_Identifiers[a_type_1] = id

		// Initialisation of values
	}

	for idx, a_type_10 := range a_type_10Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_10", idx, a_type_10.Name)
		map_A_TYPE_10_Identifiers[a_type_10] = id

		// Initialisation of values
	}

	for idx, a_type_2 := range a_type_2Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_2", idx, a_type_2.Name)
		map_A_TYPE_2_Identifiers[a_type_2] = id

		// Initialisation of values
	}

	for idx, a_type_3 := range a_type_3Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_3", idx, a_type_3.Name)
		map_A_TYPE_3_Identifiers[a_type_3] = id

		// Initialisation of values
	}

	for idx, a_type_4 := range a_type_4Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_4", idx, a_type_4.Name)
		map_A_TYPE_4_Identifiers[a_type_4] = id

		// Initialisation of values
	}

	for idx, a_type_5 := range a_type_5Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_5", idx, a_type_5.Name)
		map_A_TYPE_5_Identifiers[a_type_5] = id

		// Initialisation of values
	}

	for idx, a_type_6 := range a_type_6Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_6", idx, a_type_6.Name)
		map_A_TYPE_6_Identifiers[a_type_6] = id

		// Initialisation of values
	}

	for idx, a_type_7 := range a_type_7Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_7", idx, a_type_7.Name)
		map_A_TYPE_7_Identifiers[a_type_7] = id

		// Initialisation of values
	}

	for idx, a_type_8 := range a_type_8Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_8", idx, a_type_8.Name)
		map_A_TYPE_8_Identifiers[a_type_8] = id

		// Initialisation of values
	}

	for idx, a_type_9 := range a_type_9Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_TYPE_9", idx, a_type_9.Name)
		map_A_TYPE_9_Identifiers[a_type_9] = id

		// Initialisation of values
	}

	for idx, a_values := range a_valuesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_VALUES", idx, a_values.Name)
		map_A_VALUES_Identifiers[a_values] = id

		// Initialisation of values
	}

	for idx, a_values_1 := range a_values_1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("A_VALUES_1", idx, a_values_1.Name)
		map_A_VALUES_1_Identifiers[a_values_1] = id

		// Initialisation of values
		for _, _attribute_value_boolean := range a_values_1.ATTRIBUTE_VALUE_BOOLEAN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_BOOLEAN_Identifiers[_attribute_value_boolean])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_date := range a_values_1.ATTRIBUTE_VALUE_DATE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_DATE_Identifiers[_attribute_value_date])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_enumeration := range a_values_1.ATTRIBUTE_VALUE_ENUMERATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_ENUMERATION_Identifiers[_attribute_value_enumeration])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_integer := range a_values_1.ATTRIBUTE_VALUE_INTEGER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_INTEGER_Identifiers[_attribute_value_integer])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_real := range a_values_1.ATTRIBUTE_VALUE_REAL {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_REAL_Identifiers[_attribute_value_real])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_string := range a_values_1.ATTRIBUTE_VALUE_STRING {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_STRING_Identifiers[_attribute_value_string])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute_value_xhtml := range a_values_1.ATTRIBUTE_VALUE_XHTML {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ATTRIBUTE_VALUE_XHTML_Identifiers[_attribute_value_xhtml])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_boolean := range datatype_definition_booleanOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_BOOLEAN", idx, datatype_definition_boolean.Name)
		map_DATATYPE_DEFINITION_BOOLEAN_Identifiers[datatype_definition_boolean] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_date := range datatype_definition_dateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_DATE", idx, datatype_definition_date.Name)
		map_DATATYPE_DEFINITION_DATE_Identifiers[datatype_definition_date] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_date.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_ENUMERATION", idx, datatype_definition_enumeration.Name)
		map_DATATYPE_DEFINITION_ENUMERATION_Identifiers[datatype_definition_enumeration] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_specified_values := range datatype_definition_enumeration.SPECIFIED_VALUES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFIED_VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPECIFIED_VALUES_Identifiers[_a_specified_values])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_integer := range datatype_definition_integerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_INTEGER", idx, datatype_definition_integer.Name)
		map_DATATYPE_DEFINITION_INTEGER_Identifiers[datatype_definition_integer] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_integer.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_real := range datatype_definition_realOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_REAL", idx, datatype_definition_real.Name)
		map_DATATYPE_DEFINITION_REAL_Identifiers[datatype_definition_real] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_real.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_string := range datatype_definition_stringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_STRING", idx, datatype_definition_string.Name)
		map_DATATYPE_DEFINITION_STRING_Identifiers[datatype_definition_string] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_string.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DATATYPE_DEFINITION_XHTML", idx, datatype_definition_xhtml.Name)
		map_DATATYPE_DEFINITION_XHTML_Identifiers[datatype_definition_xhtml] = id

		// Initialisation of values
		for _, _a_alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, embedded_value := range embedded_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("EMBEDDED_VALUE", idx, embedded_value.Name)
		map_EMBEDDED_VALUE_Identifiers[embedded_value] = id

		// Initialisation of values
	}

	for idx, enum_value := range enum_valueOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ENUM_VALUE", idx, enum_value.Name)
		map_ENUM_VALUE_Identifiers[enum_value] = id

		// Initialisation of values
		for _, _a_alternative_id := range enum_value.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_properties := range enum_value.PROPERTIES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PROPERTIES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_PROPERTIES_Identifiers[_a_properties])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, relation_group := range relation_groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP", idx, relation_group.Name)
		map_RELATION_GROUP_Identifiers[relation_group] = id

		// Initialisation of values
		for _, _a_alternative_id := range relation_group.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_source_specification := range relation_group.SOURCE_SPECIFICATION {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCE_SPECIFICATION")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_SPECIFICATION_Identifiers[_a_source_specification])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_relations := range relation_group.SPEC_RELATIONS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATIONS_Identifiers[_a_spec_relations])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_1 := range relation_group.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_1_Identifiers[_a_type_1])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, relation_group_type := range relation_group_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RELATION_GROUP_TYPE", idx, relation_group_type.Name)
		map_RELATION_GROUP_TYPE_Identifiers[relation_group_type] = id

		// Initialisation of values
		for _, _a_alternative_id := range relation_group_type.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_attributes := range relation_group_type.SPEC_ATTRIBUTES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[_a_spec_attributes])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, req_if := range req_ifOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF", idx, req_if.Name)
		map_REQ_IF_Identifiers[req_if] = id

		// Initialisation of values
		for _, _a_the_header := range req_if.THE_HEADER {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "THE_HEADER")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_THE_HEADER_Identifiers[_a_the_header])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_core_content := range req_if.CORE_CONTENT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CORE_CONTENT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CORE_CONTENT_Identifiers[_a_core_content])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_tool_extensions := range req_if.TOOL_EXTENSIONS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TOOL_EXTENSIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TOOL_EXTENSIONS_Identifiers[_a_tool_extensions])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, req_if_content := range req_if_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_CONTENT", idx, req_if_content.Name)
		map_REQ_IF_CONTENT_Identifiers[req_if_content] = id

		// Initialisation of values
		for _, _a_datatypes := range req_if_content.DATATYPES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DATATYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_DATATYPES_Identifiers[_a_datatypes])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_types := range req_if_content.SPEC_TYPES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_TYPES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_TYPES_Identifiers[_a_spec_types])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_objects := range req_if_content.SPEC_OBJECTS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_OBJECTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_OBJECTS_Identifiers[_a_spec_objects])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_relations_1 := range req_if_content.SPEC_RELATIONS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATIONS_1_Identifiers[_a_spec_relations_1])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_specifications := range req_if_content.SPECIFICATIONS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPECIFICATIONS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPECIFICATIONS_Identifiers[_a_specifications])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_relation_groups := range req_if_content.SPEC_RELATION_GROUPS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_RELATION_GROUPS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_RELATION_GROUPS_Identifiers[_a_spec_relation_groups])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, req_if_header := range req_if_headerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_HEADER", idx, req_if_header.Name)
		map_REQ_IF_HEADER_Identifiers[req_if_header] = id

		// Initialisation of values
	}

	for idx, req_if_tool_extension := range req_if_tool_extensionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("REQ_IF_TOOL_EXTENSION", idx, req_if_tool_extension.Name)
		map_REQ_IF_TOOL_EXTENSION_Identifiers[req_if_tool_extension] = id

		// Initialisation of values
	}

	for idx, specification := range specificationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION", idx, specification.Name)
		map_SPECIFICATION_Identifiers[specification] = id

		// Initialisation of values
		for _, _a_alternative_id := range specification.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_children := range specification.CHILDREN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CHILDREN_Identifiers[_a_children])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_values_1 := range specification.VALUES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_VALUES_1_Identifiers[_a_values_1])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_10 := range specification.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_10_Identifiers[_a_type_10])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, specification_type := range specification_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPECIFICATION_TYPE", idx, specification_type.Name)
		map_SPECIFICATION_TYPE_Identifiers[specification_type] = id

		// Initialisation of values
		for _, _a_alternative_id := range specification_type.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_attributes := range specification_type.SPEC_ATTRIBUTES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[_a_spec_attributes])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spec_hierarchy := range spec_hierarchyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_HIERARCHY", idx, spec_hierarchy.Name)
		map_SPEC_HIERARCHY_Identifiers[spec_hierarchy] = id

		// Initialisation of values
		for _, _a_alternative_id := range spec_hierarchy.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_children := range spec_hierarchy.CHILDREN {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CHILDREN")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_CHILDREN_Identifiers[_a_children])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_editable_atts := range spec_hierarchy.EDITABLE_ATTS {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EDITABLE_ATTS")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_EDITABLE_ATTS_Identifiers[_a_editable_atts])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_object := range spec_hierarchy.OBJECT {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OBJECT")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_OBJECT_Identifiers[_a_object])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spec_object := range spec_objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT", idx, spec_object.Name)
		map_SPEC_OBJECT_Identifiers[spec_object] = id

		// Initialisation of values
		for _, _a_alternative_id := range spec_object.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_values_1 := range spec_object.VALUES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_VALUES_1_Identifiers[_a_values_1])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_2 := range spec_object.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_2_Identifiers[_a_type_2])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spec_object_type := range spec_object_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_OBJECT_TYPE", idx, spec_object_type.Name)
		map_SPEC_OBJECT_TYPE_Identifiers[spec_object_type] = id

		// Initialisation of values
		for _, _a_alternative_id := range spec_object_type.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_attributes := range spec_object_type.SPEC_ATTRIBUTES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[_a_spec_attributes])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spec_relation := range spec_relationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION", idx, spec_relation.Name)
		map_SPEC_RELATION_Identifiers[spec_relation] = id

		// Initialisation of values
		for _, _a_alternative_id := range spec_relation.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_values_1 := range spec_relation.VALUES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "VALUES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_VALUES_1_Identifiers[_a_values_1])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_source := range spec_relation.SOURCE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SOURCE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SOURCE_Identifiers[_a_source])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_type_4 := range spec_relation.TYPE {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TYPE")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_TYPE_4_Identifiers[_a_type_4])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, spec_relation_type := range spec_relation_typeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SPEC_RELATION_TYPE", idx, spec_relation_type.Name)
		map_SPEC_RELATION_TYPE_Identifiers[spec_relation_type] = id

		// Initialisation of values
		for _, _a_alternative_id := range spec_relation_type.ALTERNATIVE_ID {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_ALTERNATIVE_ID_Identifiers[_a_alternative_id])
			pointersInitializesStatements += setPointerField
		}

		for _, _a_spec_attributes := range spec_relation_type.SPEC_ATTRIBUTES {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_A_SPEC_ATTRIBUTES_Identifiers[_a_spec_attributes])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, xhtml_content := range xhtml_contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XHTML_CONTENT", idx, xhtml_content.Name)
		map_XHTML_CONTENT_Identifiers[xhtml_content] = id

		// Initialisation of values
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
