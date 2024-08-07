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

	__Enumeration__000000_ := (&models.Enumeration{Name: ``}).Stage(stage)
	__Enumeration__000001_ := (&models.Enumeration{Name: ``}).Stage(stage)

	__Restriction__000000_ := (&models.Restriction{Name: ``}).Stage(stage)
	__Restriction__000001_ := (&models.Restriction{Name: ``}).Stage(stage)

	__Schema__000000_ := (&models.Schema{Name: ``}).Stage(stage)

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

	__Enumeration__000000_.Name = ``
	__Enumeration__000000_.Value = `Paperback`

	__Enumeration__000001_.Name = ``
	__Enumeration__000001_.Value = `Hardcover`

	__Restriction__000000_.Name = ``
	__Restriction__000000_.Base = `xs:string`

	__Restriction__000001_.Name = ``
	__Restriction__000001_.Base = `xs:integer`

	__Schema__000000_.Name = ``

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
	__Restriction__000000_.Enumerations = append(__Restriction__000000_.Enumerations, __Enumeration__000000_)
	__Restriction__000000_.Enumerations = append(__Restriction__000000_.Enumerations, __Enumeration__000001_)
	__Schema__000000_.Elements = append(__Schema__000000_.Elements, __Element__000002_books)
	__Schema__000000_.SimpleTypes = append(__Schema__000000_.SimpleTypes, __SimpleType__000001_yearType)
	__Schema__000000_.SimpleTypes = append(__Schema__000000_.SimpleTypes, __SimpleType__000000_bookFormatEnum)
	__Schema__000000_.ComplexTypes = append(__Schema__000000_.ComplexTypes, __ComplexType__000000_bookType)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000004_title)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000000_author)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000005_year)
	__Sequence__000000_within_bookType.Elements = append(__Sequence__000000_within_bookType.Elements, __Element__000003_format)
	__Sequence__000001_within_within_books.Elements = append(__Sequence__000001_within_within_books.Elements, __Element__000001_book)
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_
	__SimpleType__000001_yearType.Restriction = __Restriction__000001_
}
