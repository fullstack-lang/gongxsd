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

	map_ComplexContent_Identifiers := make(map[*ComplexContent]string)
	_ = map_ComplexContent_Identifiers

	complexcontentOrdered := []*ComplexContent{}
	for complexcontent := range stage.ComplexContents {
		complexcontentOrdered = append(complexcontentOrdered, complexcontent)
	}
	sort.Slice(complexcontentOrdered[:], func(i, j int) bool {
		return complexcontentOrdered[i].Name < complexcontentOrdered[j].Name
	})
	if len(complexcontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, complexcontent := range complexcontentOrdered {

		id = generatesIdentifier("ComplexContent", idx, complexcontent.Name)
		map_ComplexContent_Identifiers[complexcontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexContent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complexcontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complexcontent.Name))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Text))
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

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Xs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(schema.Xs))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Schema_A_ComplexType_A_ComplexContentDummy")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", schema.Schema_A_ComplexType_A_ComplexContentDummy))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", schema.Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", schema.Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy))
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

	for idx, complexcontent := range complexcontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexContent", idx, complexcontent.Name)
		map_ComplexContent_Identifiers[complexcontent] = id

		// Initialisation of values
	}

	for idx, complextype := range complextypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		// Initialisation of values
	}

	for idx, documentation := range documentationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		// Initialisation of values
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

		if schema.Sequence2.ComplexType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence2.ComplexType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexType_Identifiers[schema.Sequence2.ComplexType])
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
