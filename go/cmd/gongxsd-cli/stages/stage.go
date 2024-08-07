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

	__ComplexType__000000_bookType := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__000001_books_E := (&models.ComplexType{Name: `books_E`}).Stage(stage)

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

	__Pattern__000000_ := (&models.Pattern{Name: ``}).Stage(stage)

	__Restriction__000000_bookFormatEnum_E := (&models.Restriction{Name: `bookFormatEnum_E`}).Stage(stage)
	__Restriction__000001_titleType_E := (&models.Restriction{Name: `titleType_E`}).Stage(stage)
	__Restriction__000002_yearType_E := (&models.Restriction{Name: `yearType_E`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__000000_bookType_E := (&models.Sequence{Name: `bookType_E`}).Stage(stage)
	__Sequence__000001_books_E_E := (&models.Sequence{Name: `books_E_E`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_titleType := (&models.SimpleType{Name: `titleType`}).Stage(stage)
	__SimpleType__000002_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	// Setup of values

	__ComplexType__000000_bookType.Name = `bookType`
	__ComplexType__000000_bookType.NameXSD = `bookType`

	__ComplexType__000001_books_E.Name = `books_E`
	__ComplexType__000001_books_E.NameXSD = ``

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

	__Pattern__000000_.Name = ``
	__Pattern__000000_.Value = `[A-Za-z ]+`

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

	// Setup of pointers
	__ComplexType__000000_bookType.Sequence = __Sequence__000000_bookType_E
	__ComplexType__000001_books_E.Sequence = __Sequence__000001_books_E_E
	__Element__000002_books.ComplexType = __ComplexType__000001_books_E
	__Restriction__000000_bookFormatEnum_E.Enumerations = append(__Restriction__000000_bookFormatEnum_E.Enumerations, __Enumeration__000000_bookFormatEnum_E_E)
	__Restriction__000000_bookFormatEnum_E.Enumerations = append(__Restriction__000000_bookFormatEnum_E.Enumerations, __Enumeration__000001_bookFormatEnum_E_E)
	__Restriction__000001_titleType_E.Pattern = __Pattern__000000_
	__Restriction__000002_yearType_E.MinInclusive = __MinInclusive__000000_yearType_E_E
	__Restriction__000002_yearType_E.MaxInclusive = __MaxInclusive__000000_yearType_E_E
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
	__SimpleType__000002_yearType.Restriction = __Restriction__000002_yearType_E
}
