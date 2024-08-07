package main

import (
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"

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

	// Declaration of instances to stage

	__Annotation__000000_Schema_E := (&models.Annotation{Name: `Schema_E`}).Stage(stage)
	__Annotation__000001_bookType_E := (&models.Annotation{Name: `bookType_E`}).Stage(stage)
	__Annotation__000002_book_E := (&models.Annotation{Name: `book_E`}).Stage(stage)
	__Annotation__000003_format_E := (&models.Annotation{Name: `format_E`}).Stage(stage)
	__Annotation__000004_title_E := (&models.Annotation{Name: `title_E`}).Stage(stage)
	__Annotation__000005_yearType_E := (&models.Annotation{Name: `yearType_E`}).Stage(stage)

	__ComplexType__000000_bookType := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__000001_books_E := (&models.ComplexType{Name: `books_E`}).Stage(stage)

	__Documentation__000000_Schema_E_E := (&models.Documentation{Name: `Schema_E_E`}).Stage(stage)
	__Documentation__000001_Schema_E_E := (&models.Documentation{Name: `Schema_E_E`}).Stage(stage)
	__Documentation__000002_bookType_E_E := (&models.Documentation{Name: `bookType_E_E`}).Stage(stage)
	__Documentation__000003_book_E_E := (&models.Documentation{Name: `book_E_E`}).Stage(stage)
	__Documentation__000004_format_E_E := (&models.Documentation{Name: `format_E_E`}).Stage(stage)
	__Documentation__000005_title_E_E := (&models.Documentation{Name: `title_E_E`}).Stage(stage)
	__Documentation__000006_yearType_E_E := (&models.Documentation{Name: `yearType_E_E`}).Stage(stage)

	__Element__000000_author := (&models.Element{Name: `author`}).Stage(stage)
	__Element__000001_book := (&models.Element{Name: `book`}).Stage(stage)
	__Element__000002_books := (&models.Element{Name: `books`}).Stage(stage)
	__Element__000003_format := (&models.Element{Name: `format`}).Stage(stage)
	__Element__000004_title := (&models.Element{Name: `title`}).Stage(stage)
	__Element__000005_year := (&models.Element{Name: `year`}).Stage(stage)

	__Enumeration__000000_bookFormatEnum_E_E := (&models.Enumeration{Name: `bookFormatEnum_E_E`}).Stage(stage)
	__Enumeration__000001_bookFormatEnum_E_E := (&models.Enumeration{Name: `bookFormatEnum_E_E`}).Stage(stage)

	__MaxInclusive__000000_yearType_E_E := (&models.MaxInclusive{Name: `yearType_E_E`}).Stage(stage)

	__MinInclusive__000000_yearType_E_E := (&models.MinInclusive{Name: `yearType_E_E`}).Stage(stage)

	__Pattern__000000_titleType_E_E := (&models.Pattern{Name: `titleType_E_E`}).Stage(stage)

	__Restriction__000000_bookFormatEnum_E := (&models.Restriction{Name: `bookFormatEnum_E`}).Stage(stage)
	__Restriction__000001_titleType_E := (&models.Restriction{Name: `titleType_E`}).Stage(stage)
	__Restriction__000002_yearType_E := (&models.Restriction{Name: `yearType_E`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__000000_bookType_E := (&models.Sequence{Name: `bookType_E`}).Stage(stage)
	__Sequence__000001_books_E_E := (&models.Sequence{Name: `books_E_E`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_titleType := (&models.SimpleType{Name: `titleType`}).Stage(stage)
	__SimpleType__000002_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	__WhiteSpace__000000_titleType_E_E := (&models.WhiteSpace{Name: `titleType_E_E`}).Stage(stage)

	// Setup of values

	__Annotation__000000_Schema_E.Name = `Schema_E`

	__Annotation__000001_bookType_E.Name = `bookType_E`

	__Annotation__000002_book_E.Name = `book_E`

	__Annotation__000003_format_E.Name = `format_E`

	__Annotation__000004_title_E.Name = `title_E`

	__Annotation__000005_yearType_E.Name = `yearType_E`

	__ComplexType__000000_bookType.Name = `bookType`
	__ComplexType__000000_bookType.NameXSD = `bookType`

	__ComplexType__000001_books_E.Name = `books_E`
	__ComplexType__000001_books_E.NameXSD = ``

	__Documentation__000000_Schema_E_E.Name = `Schema_E_E`
	__Documentation__000000_Schema_E_E.Source = `http://example.com/schema-docs`
	__Documentation__000000_Schema_E_E.Lang = `en`

	__Documentation__000001_Schema_E_E.Name = `Schema_E_E`
	__Documentation__000001_Schema_E_E.Source = `http://example.com/schema-docs`
	__Documentation__000001_Schema_E_E.Lang = `fr`

	__Documentation__000002_bookType_E_E.Name = `bookType_E_E`
	__Documentation__000002_bookType_E_E.Source = ``
	__Documentation__000002_bookType_E_E.Lang = ``

	__Documentation__000003_book_E_E.Name = `book_E_E`
	__Documentation__000003_book_E_E.Source = ``
	__Documentation__000003_book_E_E.Lang = ``

	__Documentation__000004_format_E_E.Name = `format_E_E`
	__Documentation__000004_format_E_E.Source = ``
	__Documentation__000004_format_E_E.Lang = ``

	__Documentation__000005_title_E_E.Name = `title_E_E`
	__Documentation__000005_title_E_E.Source = ``
	__Documentation__000005_title_E_E.Lang = ``

	__Documentation__000006_yearType_E_E.Name = `yearType_E_E`
	__Documentation__000006_yearType_E_E.Source = ``
	__Documentation__000006_yearType_E_E.Lang = ``

	__Element__000000_author.Name = `author`
	__Element__000000_author.NameXSD = `author`
	__Element__000000_author.Type = `xs:string`

	__Element__000001_book.Name = `book`
	__Element__000001_book.NameXSD = `book`
	__Element__000001_book.Type = `bookType`

	__Element__000002_books.Name = `books`
	__Element__000002_books.NameXSD = `books`
	__Element__000002_books.Type = ``

	__Element__000003_format.Name = `format`
	__Element__000003_format.NameXSD = `format`
	__Element__000003_format.Type = `bookFormatEnum`

	__Element__000004_title.Name = `title`
	__Element__000004_title.NameXSD = `title`
	__Element__000004_title.Type = `titleType`

	__Element__000005_year.Name = `year`
	__Element__000005_year.NameXSD = `year`
	__Element__000005_year.Type = `yearType`

	__Enumeration__000000_bookFormatEnum_E_E.Name = `bookFormatEnum_E_E`
	__Enumeration__000000_bookFormatEnum_E_E.Value = `Paperback`

	__Enumeration__000001_bookFormatEnum_E_E.Name = `bookFormatEnum_E_E`
	__Enumeration__000001_bookFormatEnum_E_E.Value = `Hardcover`

	__MaxInclusive__000000_yearType_E_E.Name = `yearType_E_E`
	__MaxInclusive__000000_yearType_E_E.Value = `2100`

	__MinInclusive__000000_yearType_E_E.Name = `yearType_E_E`
	__MinInclusive__000000_yearType_E_E.Value = `1900`

	__Pattern__000000_titleType_E_E.Name = `titleType_E_E`
	__Pattern__000000_titleType_E_E.Value = `[A-Za-z ]+`

	__Restriction__000000_bookFormatEnum_E.Name = `bookFormatEnum_E`
	__Restriction__000000_bookFormatEnum_E.Base = `xs:string`

	__Restriction__000001_titleType_E.Name = `titleType_E`
	__Restriction__000001_titleType_E.Base = `xs:string`

	__Restriction__000002_yearType_E.Name = `yearType_E`
	__Restriction__000002_yearType_E.Base = `xs:integer`

	__Schema__000000_Schema.Name = `Schema`

	__Sequence__000000_bookType_E.Name = `bookType_E`

	__Sequence__000001_books_E_E.Name = `books_E_E`

	__SimpleType__000000_bookFormatEnum.Name = `bookFormatEnum`
	__SimpleType__000000_bookFormatEnum.NameXSD = `bookFormatEnum`

	__SimpleType__000001_titleType.Name = `titleType`
	__SimpleType__000001_titleType.NameXSD = `titleType`

	__SimpleType__000002_yearType.Name = `yearType`
	__SimpleType__000002_yearType.NameXSD = `yearType`

	__WhiteSpace__000000_titleType_E_E.Name = `titleType_E_E`
	__WhiteSpace__000000_titleType_E_E.Value = `collapse`

	// Setup of pointers
	__Annotation__000000_Schema_E.Documentations = append(__Annotation__000000_Schema_E.Documentations, __Documentation__000000_Schema_E_E)
	__Annotation__000000_Schema_E.Documentations = append(__Annotation__000000_Schema_E.Documentations, __Documentation__000001_Schema_E_E)
	__Annotation__000001_bookType_E.Documentations = append(__Annotation__000001_bookType_E.Documentations, __Documentation__000002_bookType_E_E)
	__Annotation__000002_book_E.Documentations = append(__Annotation__000002_book_E.Documentations, __Documentation__000003_book_E_E)
	__Annotation__000003_format_E.Documentations = append(__Annotation__000003_format_E.Documentations, __Documentation__000004_format_E_E)
	__Annotation__000004_title_E.Documentations = append(__Annotation__000004_title_E.Documentations, __Documentation__000005_title_E_E)
	__Annotation__000005_yearType_E.Documentations = append(__Annotation__000005_yearType_E.Documentations, __Documentation__000006_yearType_E_E)
	__ComplexType__000000_bookType.Annotation = __Annotation__000001_bookType_E
	__ComplexType__000000_bookType.Sequence = __Sequence__000000_bookType_E
	__ComplexType__000001_books_E.Sequence = __Sequence__000001_books_E_E
	__Element__000001_book.Annotation = __Annotation__000002_book_E
	__Element__000002_books.ComplexType = __ComplexType__000001_books_E
	__Element__000003_format.Annotation = __Annotation__000003_format_E
	__Element__000004_title.Annotation = __Annotation__000004_title_E
	__Restriction__000000_bookFormatEnum_E.Enumerations = append(__Restriction__000000_bookFormatEnum_E.Enumerations, __Enumeration__000000_bookFormatEnum_E_E)
	__Restriction__000000_bookFormatEnum_E.Enumerations = append(__Restriction__000000_bookFormatEnum_E.Enumerations, __Enumeration__000001_bookFormatEnum_E_E)
	__Restriction__000001_titleType_E.Pattern = __Pattern__000000_titleType_E_E
	__Restriction__000001_titleType_E.WhiteSpace = __WhiteSpace__000000_titleType_E_E
	__Restriction__000002_yearType_E.MinInclusive = __MinInclusive__000000_yearType_E_E
	__Restriction__000002_yearType_E.MaxInclusive = __MaxInclusive__000000_yearType_E_E
	__Schema__000000_Schema.Annotation = __Annotation__000000_Schema_E
	__Schema__000000_Schema.Elements = append(__Schema__000000_Schema.Elements, __Element__000002_books)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000002_yearType)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000000_bookFormatEnum)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000001_titleType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000000_bookType)
	__Sequence__000000_bookType_E.Elements = append(__Sequence__000000_bookType_E.Elements, __Element__000004_title)
	__Sequence__000000_bookType_E.Elements = append(__Sequence__000000_bookType_E.Elements, __Element__000000_author)
	__Sequence__000000_bookType_E.Elements = append(__Sequence__000000_bookType_E.Elements, __Element__000005_year)
	__Sequence__000000_bookType_E.Elements = append(__Sequence__000000_bookType_E.Elements, __Element__000003_format)
	__Sequence__000001_books_E_E.Elements = append(__Sequence__000001_books_E_E.Elements, __Element__000001_book)
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_bookFormatEnum_E
	__SimpleType__000001_titleType.Restriction = __Restriction__000001_titleType_E
	__SimpleType__000002_yearType.Annotation = __Annotation__000005_yearType_E
	__SimpleType__000002_yearType.Restriction = __Restriction__000002_yearType_E
}
