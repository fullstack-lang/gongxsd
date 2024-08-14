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
	map_BookDetailsGroup_Identifiers := make(map[*BookDetailsGroup]string)
	_ = map_BookDetailsGroup_Identifiers

	bookdetailsgroupOrdered := []*BookDetailsGroup{}
	for bookdetailsgroup := range stage.BookDetailsGroups {
		bookdetailsgroupOrdered = append(bookdetailsgroupOrdered, bookdetailsgroup)
	}
	sort.Slice(bookdetailsgroupOrdered[:], func(i, j int) bool {
		return bookdetailsgroupOrdered[i].Name < bookdetailsgroupOrdered[j].Name
	})
	if len(bookdetailsgroupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bookdetailsgroup := range bookdetailsgroupOrdered {

		id = generatesIdentifier("BookDetailsGroup", idx, bookdetailsgroup.Name)
		map_BookDetailsGroup_Identifiers[bookdetailsgroup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BookDetailsGroup")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bookdetailsgroup.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bookdetailsgroup.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Title")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bookdetailsgroup.Title))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Author")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bookdetailsgroup.Author))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Year")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bookdetailsgroup.Year))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Format")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bookdetailsgroup.Format))
		initializerStatements += setValueField

	}

	map_BookType_Identifiers := make(map[*BookType]string)
	_ = map_BookType_Identifiers

	booktypeOrdered := []*BookType{}
	for booktype := range stage.BookTypes {
		booktypeOrdered = append(booktypeOrdered, booktype)
	}
	sort.Slice(booktypeOrdered[:], func(i, j int) bool {
		return booktypeOrdered[i].Name < booktypeOrdered[j].Name
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
		return booksOrdered[i].Name < booksOrdered[j].Name
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
		return creditOrdered[i].Name < creditOrdered[j].Name
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
		return linkOrdered[i].Name < linkOrdered[j].Name
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnclosedText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.EnclosedText))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.NameXSD))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, bookdetailsgroup := range bookdetailsgroupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("BookDetailsGroup", idx, bookdetailsgroup.Name)
		map_BookDetailsGroup_Identifiers[bookdetailsgroup] = id

		// Initialisation of values
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
