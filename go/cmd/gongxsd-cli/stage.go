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
	__ComplexType__000001_within_books := (&models.ComplexType{Name: `within books`}).Stage(stage)

	__Element__000000_author := (&models.Element{Name: `author`}).Stage(stage)
	__Element__000001_book := (&models.Element{Name: `book`}).Stage(stage)
	__Element__000002_books := (&models.Element{Name: `books`}).Stage(stage)
	__Element__000003_format := (&models.Element{Name: `format`}).Stage(stage)
	__Element__000004_title := (&models.Element{Name: `title`}).Stage(stage)
	__Element__000005_year := (&models.Element{Name: `year`}).Stage(stage)

	__Enumeration__000000_within_within_bookFormatEnum := (&models.Enumeration{Name: `within within bookFormatEnum`}).Stage(stage)
	__Enumeration__000001_within_within_bookFormatEnum := (&models.Enumeration{Name: `within within bookFormatEnum`}).Stage(stage)

	__Restriction__000000_within_bookFormatEnum := (&models.Restriction{Name: `within bookFormatEnum`}).Stage(stage)
	__Restriction__000001_within_yearType := (&models.Restriction{Name: `within yearType`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__000000_within_bookType := (&models.Sequence{Name: `within bookType`}).Stage(stage)
	__Sequence__000001_within_within_books := (&models.Sequence{Name: `within within books`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	// Setup of values

	__ComplexType__000000_bookType.Name = `bookType`
	__ComplexType__000000_bookType.NameXSD = `bookType`

	__ComplexType__000001_within_books.Name = `within books`
	__ComplexType__000001_within_books.NameXSD = ``

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
	__Element__000004_title.Type = `xs:string`

	__Element__000005_year.Name = `year`
	__Element__000005_year.NameXSD = `year`
	__Element__000005_year.Type = `yearType`

	__Enumeration__000000_within_within_bookFormatEnum.Name = `within within bookFormatEnum`
	__Enumeration__000000_within_within_bookFormatEnum.Value = `Paperback`

	__Enumeration__000001_within_within_bookFormatEnum.Name = `within within bookFormatEnum`
	__Enumeration__000001_within_within_bookFormatEnum.Value = `Hardcover`

	__Restriction__000000_within_bookFormatEnum.Name = `within bookFormatEnum`
	__Restriction__000000_within_bookFormatEnum.Base = `xs:string`

	__Restriction__000001_within_yearType.Name = `within yearType`
	__Restriction__000001_within_yearType.Base = `xs:integer`

	__Schema__000000_Schema.Name = `Schema`

	__Sequence__000000_within_bookType.Name = `within bookType`

	__Sequence__000001_within_within_books.Name = `within within books`

	__SimpleType__000000_bookFormatEnum.Name = `bookFormatEnum`
	__SimpleType__000000_bookFormatEnum.NameXSD = `bookFormatEnum`

	__SimpleType__000001_yearType.Name = `yearType`
	__SimpleType__000001_yearType.NameXSD = `yearType`

	// Setup of pointers
	__ComplexType__000000_bookType.Sequence = __Sequence__000000_within_bookType
	__ComplexType__000001_within_books.Sequence = __Sequence__000001_within_within_books
	__Element__000002_books.ComplexType = __ComplexType__000001_within_books
	__Restriction__000000_within_bookFormatEnum.Enumerations = append(__Restriction__000000_within_bookFormatEnum.Enumerations, __Enumeration__000000_within_within_bookFormatEnum)
	__Restriction__000000_within_bookFormatEnum.Enumerations = append(__Restriction__000000_within_bookFormatEnum.Enumerations, __Enumeration__000001_within_within_bookFormatEnum)
	__Schema__000000_Schema.Elements = append(__Schema__000000_Schema.Elements, __Element__000002_books)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000001_yearType)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000000_bookFormatEnum)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000000_bookType)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000004_title)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000000_author)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000005_year)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000003_format)
	__Sequence__000001_within_within_books.Elements = append(__Sequence__000001_within_within_books.Elements, __Element__000001_book)
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_within_bookFormatEnum
	__SimpleType__000001_yearType.Restriction = __Restriction__000001_within_yearType
}
