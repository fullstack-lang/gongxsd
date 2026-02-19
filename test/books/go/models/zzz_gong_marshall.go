// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
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

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

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
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
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
		identifiersDecl.WriteString("\n")
	}
	for _, booktype := range booktypeOrdered {

		identifiersDecl.WriteString(booktype.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Edition"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Isbn"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Bestseller"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Author"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Year"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Format"))
		pointersInitializesStatements.WriteString(booktype.GongMarshallField(stage, "Credit"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, books := range booksOrdered {

		identifiersDecl.WriteString(books.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(books.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(books.GongMarshallField(stage, "Book"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, credit := range creditOrdered {

		identifiersDecl.WriteString(credit.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Page"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_type"))
		pointersInitializesStatements.WriteString(credit.GongMarshallField(stage, "Link"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_words"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_symbol"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, link := range linkOrdered {

		identifiersDecl.WriteString(link.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(link.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EnclosedText"))
	}

	// insertion initialization of objects to stage
	for _, booktype := range booktypeOrdered {
		_ = booktype
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, books := range booksOrdered {
		_ = books
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, credit := range creditOrdered {
		_ = credit
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, link := range linkOrdered {
		_ = link
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

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
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (booktype *BookType) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Name))
	case "Edition":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Edition")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Edition))
	case "Isbn":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Isbn")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Isbn))
	case "Bestseller":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bestseller")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", booktype.Bestseller))
	case "Title":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Title")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Title))
	case "Author":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Author")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Author))
	case "Year":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Year")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", booktype.Year))
	case "Format":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", booktype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Format")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(booktype.Format))

	case "Credit":
		var sb strings.Builder
		for _, _credit := range booktype.Credit {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", booktype.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Credit")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _credit.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct BookType", fieldName)
	}
	return
}

func (books *Books) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", books.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(books.Name))

	case "Book":
		var sb strings.Builder
		for _, _booktype := range books.Book {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", books.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Book")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _booktype.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Books", fieldName)
	}
	return
}

func (credit *Credit) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", credit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(credit.Name))
	case "Page":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", credit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Page")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", credit.Page))
	case "Credit_type":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", credit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Credit_type")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(credit.Credit_type))
	case "Credit_words":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", credit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Credit_words")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(credit.Credit_words))
	case "Credit_symbol":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", credit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Credit_symbol")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(credit.Credit_symbol))

	case "Link":
		var sb strings.Builder
		for _, _link := range credit.Link {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", credit.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Link")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Credit", fieldName)
	}
	return
}

func (link *Link) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.Name))
	case "NameXSD":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NameXSD")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.NameXSD))
	case "EnclosedText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosedText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.EnclosedText))

	default:
		log.Panicf("Unknown field %s for Gongstruct Link", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (booktype *BookType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Edition"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Isbn"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Bestseller"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Author"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Year"))
		initializerStatements.WriteString(booktype.GongMarshallField(stage, "Format"))
		pointersInitializesStatements.WriteString(booktype.GongMarshallField(stage, "Credit"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (books *Books) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(books.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(books.GongMarshallField(stage, "Book"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (credit *Credit) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Page"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_type"))
		pointersInitializesStatements.WriteString(credit.GongMarshallField(stage, "Link"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_words"))
		initializerStatements.WriteString(credit.GongMarshallField(stage, "Credit_symbol"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (link *Link) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(link.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
