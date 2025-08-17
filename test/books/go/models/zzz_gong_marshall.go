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

	"github.com/sergi/go-diff/diffmatchpatch"
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

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

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
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)
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
	map_BookType_Identifiers := make(map[*BookType]string)
	_ = map_BookType_Identifiers

	booktypeOrdered := []*BookType{}
	for booktype := range stage.BookTypes {
		booktypeOrdered = append(booktypeOrdered, booktype)
	}
	sort.Slice(booktypeOrdered[:], func(i, j int) bool {
		booktypei := booktypeOrdered[i]
		booktypej := booktypeOrdered[j]
		booktypei_order, oki := stage.BookTypeMap_Staged_Order[booktypei]
		booktypej_order, okj := stage.BookTypeMap_Staged_Order[booktypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return booktypei_order < booktypej_order
	})
	if len(booktypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, booktype := range booktypeOrdered {

		id = generatesIdentifier("BookType", idx, booktype.Name)
		map_BookType_Identifiers[booktype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BookType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", booktype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Edition")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Edition))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Isbn")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Isbn))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Bestseller")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", booktype.Bestseller))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Title")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Title))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Author")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Author))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Year")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", booktype.Year))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Format")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(booktype.Format))
		initializerStatements += setValueField

	}

	map_Books_Identifiers := make(map[*Books]string)
	_ = map_Books_Identifiers

	booksOrdered := []*Books{}
	for books := range stage.Bookss {
		booksOrdered = append(booksOrdered, books)
	}
	sort.Slice(booksOrdered[:], func(i, j int) bool {
		booksi := booksOrdered[i]
		booksj := booksOrdered[j]
		booksi_order, oki := stage.BooksMap_Staged_Order[booksi]
		booksj_order, okj := stage.BooksMap_Staged_Order[booksj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return booksi_order < booksj_order
	})
	if len(booksOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, books := range booksOrdered {

		id = generatesIdentifier("Books", idx, books.Name)
		map_Books_Identifiers[books] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Books")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", books.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(books.Name))
		initializerStatements += setValueField

	}

	map_Credit_Identifiers := make(map[*Credit]string)
	_ = map_Credit_Identifiers

	creditOrdered := []*Credit{}
	for credit := range stage.Credits {
		creditOrdered = append(creditOrdered, credit)
	}
	sort.Slice(creditOrdered[:], func(i, j int) bool {
		crediti := creditOrdered[i]
		creditj := creditOrdered[j]
		crediti_order, oki := stage.CreditMap_Staged_Order[crediti]
		creditj_order, okj := stage.CreditMap_Staged_Order[creditj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return crediti_order < creditj_order
	})
	if len(creditOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, credit := range creditOrdered {

		id = generatesIdentifier("Credit", idx, credit.Name)
		map_Credit_Identifiers[credit] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Credit")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", credit.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(credit.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Page")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", credit.Page))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Credit_type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(credit.Credit_type))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Credit_words")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(credit.Credit_words))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Credit_symbol")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(credit.Credit_symbol))
		initializerStatements += setValueField

	}

	map_Link_Identifiers := make(map[*Link]string)
	_ = map_Link_Identifiers

	linkOrdered := []*Link{}
	for link := range stage.Links {
		linkOrdered = append(linkOrdered, link)
	}
	sort.Slice(linkOrdered[:], func(i, j int) bool {
		linki := linkOrdered[i]
		linkj := linkOrdered[j]
		linki_order, oki := stage.LinkMap_Staged_Order[linki]
		linkj_order, okj := stage.LinkMap_Staged_Order[linkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linki_order < linkj_order
	})
	if len(linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, link := range linkOrdered {

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", link.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnclosedText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.EnclosedText))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(booktypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of BookType instances pointers"
	}
	for idx, booktype := range booktypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BookType", idx, booktype.Name)
		map_BookType_Identifiers[booktype] = id

		// Initialisation of values
		for _, _credit := range booktype.Credit {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Credit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Credit_Identifiers[_credit])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(booksOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Books instances pointers"
	}
	for idx, books := range booksOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Books", idx, books.Name)
		map_Books_Identifiers[books] = id

		// Initialisation of values
		for _, _booktype := range books.Book {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Book")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_BookType_Identifiers[_booktype])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(creditOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Credit instances pointers"
	}
	for idx, credit := range creditOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Credit", idx, credit.Name)
		map_Credit_Identifiers[credit] = id

		// Initialisation of values
		for _, _link := range credit.Link {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Link")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Link_Identifiers[_link])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(linkOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Link instances pointers"
	}
	for idx, link := range linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
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

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

	fmt.Fprintln(file, res)
}

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
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
