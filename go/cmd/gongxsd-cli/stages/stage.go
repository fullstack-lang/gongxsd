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
	__Annotation__000001_bestseller_E := (&models.Annotation{Name: `bestseller_E`}).Stage(stage)
	__Annotation__000002_bookType_E := (&models.Annotation{Name: `bookType_E`}).Stage(stage)
	__Annotation__000003_book_E := (&models.Annotation{Name: `book_E`}).Stage(stage)
	__Annotation__000004_credit_E := (&models.Annotation{Name: `credit_E`}).Stage(stage)
	__Annotation__000005_edition_E := (&models.Annotation{Name: `edition_E`}).Stage(stage)
	__Annotation__000006_format_E := (&models.Annotation{Name: `format_E`}).Stage(stage)
	__Annotation__000007_isbn_E := (&models.Annotation{Name: `isbn_E`}).Stage(stage)
	__Annotation__000008_link_E := (&models.Annotation{Name: `link_E`}).Stage(stage)
	__Annotation__000009_title_E := (&models.Annotation{Name: `title_E`}).Stage(stage)
	__Annotation__000010_yearType_E := (&models.Annotation{Name: `yearType_E`}).Stage(stage)

	__Attribute__000000_bestseller := (&models.Attribute{Name: `bestseller`}).Stage(stage)
	__Attribute__000001_edition := (&models.Attribute{Name: `edition`}).Stage(stage)
	__Attribute__000002_isbn := (&models.Attribute{Name: `isbn`}).Stage(stage)
	__Attribute__000003_name := (&models.Attribute{Name: `name`}).Stage(stage)
	__Attribute__000004_page := (&models.Attribute{Name: `page`}).Stage(stage)

	__AttributeGroup__000000_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000001_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000002_commonAttributes := (&models.AttributeGroup{Name: `commonAttributes`}).Stage(stage)
	__AttributeGroup__000003_extendedAttributes := (&models.AttributeGroup{Name: `extendedAttributes`}).Stage(stage)

	__Choice__000000_ := (&models.Choice{Name: ``}).Stage(stage)
	__Choice__000001_ := (&models.Choice{Name: ``}).Stage(stage)
	__Choice__000002_ := (&models.Choice{Name: ``}).Stage(stage)

	__ComplexType__000000_bookType := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__000001_books_E := (&models.ComplexType{Name: `books_E`}).Stage(stage)
	__ComplexType__000002_credit := (&models.ComplexType{Name: `credit`}).Stage(stage)
	__ComplexType__000003_link := (&models.ComplexType{Name: `link`}).Stage(stage)

	__Documentation__000000_Schema_E_E := (&models.Documentation{Name: `Schema_E_E`}).Stage(stage)
	__Documentation__000001_Schema_E_E := (&models.Documentation{Name: `Schema_E_E`}).Stage(stage)
	__Documentation__000002__E := (&models.Documentation{Name: `_E`}).Stage(stage)
	__Documentation__000003__E := (&models.Documentation{Name: `_E`}).Stage(stage)
	__Documentation__000004__E := (&models.Documentation{Name: `_E`}).Stage(stage)
	__Documentation__000005_bookType_E_E := (&models.Documentation{Name: `bookType_E_E`}).Stage(stage)
	__Documentation__000006_book_E_E := (&models.Documentation{Name: `book_E_E`}).Stage(stage)
	__Documentation__000007_credit_E_E := (&models.Documentation{Name: `credit_E_E`}).Stage(stage)
	__Documentation__000008_format_E_E := (&models.Documentation{Name: `format_E_E`}).Stage(stage)
	__Documentation__000009_link_E_E := (&models.Documentation{Name: `link_E_E`}).Stage(stage)
	__Documentation__000010_title_E_E := (&models.Documentation{Name: `title_E_E`}).Stage(stage)
	__Documentation__000011_yearType_E_E := (&models.Documentation{Name: `yearType_E_E`}).Stage(stage)

	__Element__000000_author := (&models.Element{Name: `author`}).Stage(stage)
	__Element__000001_book := (&models.Element{Name: `book`}).Stage(stage)
	__Element__000002_books := (&models.Element{Name: `books`}).Stage(stage)
	__Element__000003_credit := (&models.Element{Name: `credit`}).Stage(stage)
	__Element__000004_credit_symbol := (&models.Element{Name: `credit-symbol`}).Stage(stage)
	__Element__000005_credit_symbol := (&models.Element{Name: `credit-symbol`}).Stage(stage)
	__Element__000006_credit_type := (&models.Element{Name: `credit-type`}).Stage(stage)
	__Element__000007_credit_words := (&models.Element{Name: `credit-words`}).Stage(stage)
	__Element__000008_credit_words := (&models.Element{Name: `credit-words`}).Stage(stage)
	__Element__000009_format := (&models.Element{Name: `format`}).Stage(stage)
	__Element__000010_link := (&models.Element{Name: `link`}).Stage(stage)
	__Element__000011_link := (&models.Element{Name: `link`}).Stage(stage)
	__Element__000012_title := (&models.Element{Name: `title`}).Stage(stage)
	__Element__000013_year := (&models.Element{Name: `year`}).Stage(stage)

	__Enumeration__000000_bookFormatEnum_E_E := (&models.Enumeration{Name: `bookFormatEnum_E_E`}).Stage(stage)
	__Enumeration__000001_bookFormatEnum_E_E := (&models.Enumeration{Name: `bookFormatEnum_E_E`}).Stage(stage)

	__MaxInclusive__000000_yearType_E_E := (&models.MaxInclusive{Name: `yearType_E_E`}).Stage(stage)

	__MinInclusive__000000_yearType_E_E := (&models.MinInclusive{Name: `yearType_E_E`}).Stage(stage)

	__Pattern__000000_titleType_E_E := (&models.Pattern{Name: `titleType_E_E`}).Stage(stage)

	__Restriction__000000_bookFormatEnum_E := (&models.Restriction{Name: `bookFormatEnum_E`}).Stage(stage)
	__Restriction__000001_titleType_E := (&models.Restriction{Name: `titleType_E`}).Stage(stage)
	__Restriction__000002_yearType_E := (&models.Restriction{Name: `yearType_E`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__000000_ := (&models.Sequence{Name: ``}).Stage(stage)
	__Sequence__000001_ := (&models.Sequence{Name: ``}).Stage(stage)
	__Sequence__000002_bookType_E := (&models.Sequence{Name: `bookType_E`}).Stage(stage)
	__Sequence__000003_books_E_E := (&models.Sequence{Name: `books_E_E`}).Stage(stage)
	__Sequence__000004_credit_E := (&models.Sequence{Name: `credit_E`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_titleType := (&models.SimpleType{Name: `titleType`}).Stage(stage)
	__SimpleType__000002_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	__WhiteSpace__000000_titleType_E_E := (&models.WhiteSpace{Name: `titleType_E_E`}).Stage(stage)

	// Setup of values

	__Annotation__000000_Schema_E.Name = `Schema_E`

	__Annotation__000001_bestseller_E.Name = `bestseller_E`

	__Annotation__000002_bookType_E.Name = `bookType_E`

	__Annotation__000003_book_E.Name = `book_E`

	__Annotation__000004_credit_E.Name = `credit_E`

	__Annotation__000005_edition_E.Name = `edition_E`

	__Annotation__000006_format_E.Name = `format_E`

	__Annotation__000007_isbn_E.Name = `isbn_E`

	__Annotation__000008_link_E.Name = `link_E`

	__Annotation__000009_title_E.Name = `title_E`

	__Annotation__000010_yearType_E.Name = `yearType_E`

	__Attribute__000000_bestseller.Name = `bestseller`
	__Attribute__000000_bestseller.NameXSD = `bestseller`
	__Attribute__000000_bestseller.Type = `xs:boolean`
	__Attribute__000000_bestseller.HasNameConflict = false
	__Attribute__000000_bestseller.GoIdentifier = ``
	__Attribute__000000_bestseller.Default = ``
	__Attribute__000000_bestseller.Use = `optional`
	__Attribute__000000_bestseller.Form = ``
	__Attribute__000000_bestseller.Fixed = ``
	__Attribute__000000_bestseller.Ref = ``
	__Attribute__000000_bestseller.TargetNamespace = ``
	__Attribute__000000_bestseller.SimpleType = ``
	__Attribute__000000_bestseller.IDXSD = ``

	__Attribute__000001_edition.Name = `edition`
	__Attribute__000001_edition.NameXSD = `edition`
	__Attribute__000001_edition.Type = `xs:string`
	__Attribute__000001_edition.HasNameConflict = false
	__Attribute__000001_edition.GoIdentifier = ``
	__Attribute__000001_edition.Default = ``
	__Attribute__000001_edition.Use = `optional`
	__Attribute__000001_edition.Form = ``
	__Attribute__000001_edition.Fixed = ``
	__Attribute__000001_edition.Ref = ``
	__Attribute__000001_edition.TargetNamespace = ``
	__Attribute__000001_edition.SimpleType = ``
	__Attribute__000001_edition.IDXSD = ``

	__Attribute__000002_isbn.Name = `isbn`
	__Attribute__000002_isbn.NameXSD = `isbn`
	__Attribute__000002_isbn.Type = `xs:string`
	__Attribute__000002_isbn.HasNameConflict = false
	__Attribute__000002_isbn.GoIdentifier = ``
	__Attribute__000002_isbn.Default = ``
	__Attribute__000002_isbn.Use = `required`
	__Attribute__000002_isbn.Form = ``
	__Attribute__000002_isbn.Fixed = ``
	__Attribute__000002_isbn.Ref = ``
	__Attribute__000002_isbn.TargetNamespace = ``
	__Attribute__000002_isbn.SimpleType = ``
	__Attribute__000002_isbn.IDXSD = ``

	__Attribute__000003_name.Name = `name`
	__Attribute__000003_name.NameXSD = `name`
	__Attribute__000003_name.Type = `xs:token`
	__Attribute__000003_name.HasNameConflict = false
	__Attribute__000003_name.GoIdentifier = ``
	__Attribute__000003_name.Default = ``
	__Attribute__000003_name.Use = ``
	__Attribute__000003_name.Form = ``
	__Attribute__000003_name.Fixed = ``
	__Attribute__000003_name.Ref = ``
	__Attribute__000003_name.TargetNamespace = ``
	__Attribute__000003_name.SimpleType = ``
	__Attribute__000003_name.IDXSD = ``

	__Attribute__000004_page.Name = `page`
	__Attribute__000004_page.NameXSD = `page`
	__Attribute__000004_page.Type = `xs:positiveInteger`
	__Attribute__000004_page.HasNameConflict = false
	__Attribute__000004_page.GoIdentifier = ``
	__Attribute__000004_page.Default = ``
	__Attribute__000004_page.Use = ``
	__Attribute__000004_page.Form = ``
	__Attribute__000004_page.Fixed = ``
	__Attribute__000004_page.Ref = ``
	__Attribute__000004_page.TargetNamespace = ``
	__Attribute__000004_page.SimpleType = ``
	__Attribute__000004_page.IDXSD = ``

	__AttributeGroup__000000_.Name = ``
	__AttributeGroup__000000_.NameXSD = ``
	__AttributeGroup__000000_.Ref = `extendedAttributes`

	__AttributeGroup__000001_.Name = ``
	__AttributeGroup__000001_.NameXSD = ``
	__AttributeGroup__000001_.Ref = `commonAttributes`

	__AttributeGroup__000002_commonAttributes.Name = `commonAttributes`
	__AttributeGroup__000002_commonAttributes.NameXSD = `commonAttributes`
	__AttributeGroup__000002_commonAttributes.Ref = ``

	__AttributeGroup__000003_extendedAttributes.Name = `extendedAttributes`
	__AttributeGroup__000003_extendedAttributes.NameXSD = `extendedAttributes`
	__AttributeGroup__000003_extendedAttributes.Ref = ``

	__Choice__000000_.Name = ``
	__Choice__000000_.MinOccurs = ``
	__Choice__000000_.MaxOccurs = ``

	__Choice__000001_.Name = ``
	__Choice__000001_.MinOccurs = ``
	__Choice__000001_.MaxOccurs = ``

	__Choice__000002_.Name = ``
	__Choice__000002_.MinOccurs = ``
	__Choice__000002_.MaxOccurs = ``

	__ComplexType__000000_bookType.Name = `bookType`
	__ComplexType__000000_bookType.HasNameConflict = false
	__ComplexType__000000_bookType.GoIdentifier = `BookType`
	__ComplexType__000000_bookType.IsInlined = false
	__ComplexType__000000_bookType.NameXSD = `bookType`

	__ComplexType__000001_books_E.Name = `books_E`
	__ComplexType__000001_books_E.HasNameConflict = false
	__ComplexType__000001_books_E.GoIdentifier = ``
	__ComplexType__000001_books_E.IsInlined = true
	__ComplexType__000001_books_E.NameXSD = ``

	__ComplexType__000002_credit.Name = `credit`
	__ComplexType__000002_credit.HasNameConflict = false
	__ComplexType__000002_credit.GoIdentifier = `Credit`
	__ComplexType__000002_credit.IsInlined = false
	__ComplexType__000002_credit.NameXSD = `credit`

	__ComplexType__000003_link.Name = `link`
	__ComplexType__000003_link.HasNameConflict = false
	__ComplexType__000003_link.GoIdentifier = `Link`
	__ComplexType__000003_link.IsInlined = false
	__ComplexType__000003_link.NameXSD = `link`

	__Documentation__000000_Schema_E_E.Name = `Schema_E_E`
	__Documentation__000000_Schema_E_E.Text = ` This schema defines
            the structure of a simple book collection. It includes types for book details, such as
            title, author, year, and format. `
	__Documentation__000000_Schema_E_E.Source = `http://example.com/schema-docs`
	__Documentation__000000_Schema_E_E.Lang = `en`

	__Documentation__000001_Schema_E_E.Name = `Schema_E_E`
	__Documentation__000001_Schema_E_E.Text = ` Ce schéma définit
            la structure d'une collection de livres simple. Il inclut des types pour les détails du
            livre, tels que le titre, l'auteur, l'année et le format. `
	__Documentation__000001_Schema_E_E.Source = `http://example.com/schema-docs`
	__Documentation__000001_Schema_E_E.Lang = `fr`

	__Documentation__000002__E.Name = `_E`
	__Documentation__000002__E.Text = `The ISBN number of the book.`
	__Documentation__000002__E.Source = ``
	__Documentation__000002__E.Lang = ``

	__Documentation__000003__E.Name = `_E`
	__Documentation__000003__E.Text = `Indicates if the book is a bestseller.`
	__Documentation__000003__E.Source = ``
	__Documentation__000003__E.Lang = ``

	__Documentation__000004__E.Name = `_E`
	__Documentation__000004__E.Text = `The edition of the book.`
	__Documentation__000004__E.Source = ``
	__Documentation__000004__E.Lang = ``

	__Documentation__000005_bookType_E_E.Name = `bookType_E_E`
	__Documentation__000005_bookType_E_E.Text = ` This complex type defines the structure of a book, including title,
                author, year, and format. `
	__Documentation__000005_bookType_E_E.Source = ``
	__Documentation__000005_bookType_E_E.Lang = ``

	__Documentation__000006_book_E_E.Name = `book_E_E`
	__Documentation__000006_book_E_E.Text = ` A book element representing a single book in the
                            collection. `
	__Documentation__000006_book_E_E.Source = ``
	__Documentation__000006_book_E_E.Lang = ``

	__Documentation__000007_credit_E_E.Name = `credit_E_E`
	__Documentation__000007_credit_E_E.Text = `The credit type .. `
	__Documentation__000007_credit_E_E.Source = ``
	__Documentation__000007_credit_E_E.Lang = ``

	__Documentation__000008_format_E_E.Name = `format_E_E`
	__Documentation__000008_format_E_E.Text = `The format of the book, either Paperback or Hardcover.`
	__Documentation__000008_format_E_E.Source = ``
	__Documentation__000008_format_E_E.Lang = ``

	__Documentation__000009_link_E_E.Name = `link_E_E`
	__Documentation__000009_link_E_E.Text = `The link type serves as an outgoing simple XLink. If a relative link
                is used within a document that is part of a compressed MusicXML file, the link is
                relative to the root folder of the zip file.`
	__Documentation__000009_link_E_E.Source = ``
	__Documentation__000009_link_E_E.Lang = ``

	__Documentation__000010_title_E_E.Name = `title_E_E`
	__Documentation__000010_title_E_E.Text = `The title of the book, consisting of alphabetic characters and
                        spaces only.`
	__Documentation__000010_title_E_E.Source = ``
	__Documentation__000010_title_E_E.Lang = ``

	__Documentation__000011_yearType_E_E.Name = `yearType_E_E`
	__Documentation__000011_yearType_E_E.Text = ` This type represents a year. It restricts the value to an integer
                between 1900 and 2100 inclusive. `
	__Documentation__000011_yearType_E_E.Source = ``
	__Documentation__000011_yearType_E_E.Lang = ``

	__Element__000000_author.Name = `author`
	__Element__000000_author.HasNameConflict = false
	__Element__000000_author.GoIdentifier = `Author`
	__Element__000000_author.NameXSD = `author`
	__Element__000000_author.Type = `xs:string`
	__Element__000000_author.MinOccurs = ``
	__Element__000000_author.MaxOccurs = ``
	__Element__000000_author.Default = ``
	__Element__000000_author.Fixed = ``
	__Element__000000_author.Nillable = ``
	__Element__000000_author.Ref = ``
	__Element__000000_author.Abstract = ``
	__Element__000000_author.Form = ``
	__Element__000000_author.Block = ``
	__Element__000000_author.Final = ``

	__Element__000001_book.Name = `book`
	__Element__000001_book.HasNameConflict = false
	__Element__000001_book.GoIdentifier = `Book`
	__Element__000001_book.NameXSD = `book`
	__Element__000001_book.Type = `bookType`
	__Element__000001_book.MinOccurs = ``
	__Element__000001_book.MaxOccurs = `unbounded`
	__Element__000001_book.Default = ``
	__Element__000001_book.Fixed = ``
	__Element__000001_book.Nillable = ``
	__Element__000001_book.Ref = ``
	__Element__000001_book.Abstract = ``
	__Element__000001_book.Form = ``
	__Element__000001_book.Block = ``
	__Element__000001_book.Final = ``

	__Element__000002_books.Name = `books`
	__Element__000002_books.HasNameConflict = false
	__Element__000002_books.GoIdentifier = `Books`
	__Element__000002_books.NameXSD = `books`
	__Element__000002_books.Type = ``
	__Element__000002_books.MinOccurs = ``
	__Element__000002_books.MaxOccurs = ``
	__Element__000002_books.Default = ``
	__Element__000002_books.Fixed = ``
	__Element__000002_books.Nillable = ``
	__Element__000002_books.Ref = ``
	__Element__000002_books.Abstract = ``
	__Element__000002_books.Form = ``
	__Element__000002_books.Block = ``
	__Element__000002_books.Final = ``

	__Element__000003_credit.Name = `credit`
	__Element__000003_credit.HasNameConflict = false
	__Element__000003_credit.GoIdentifier = `Credit`
	__Element__000003_credit.NameXSD = `credit`
	__Element__000003_credit.Type = `credit`
	__Element__000003_credit.MinOccurs = `0`
	__Element__000003_credit.MaxOccurs = `unbounded`
	__Element__000003_credit.Default = ``
	__Element__000003_credit.Fixed = ``
	__Element__000003_credit.Nillable = ``
	__Element__000003_credit.Ref = ``
	__Element__000003_credit.Abstract = ``
	__Element__000003_credit.Form = ``
	__Element__000003_credit.Block = ``
	__Element__000003_credit.Final = ``

	__Element__000004_credit_symbol.Name = `credit-symbol`
	__Element__000004_credit_symbol.HasNameConflict = false
	__Element__000004_credit_symbol.GoIdentifier = `Credit_symbol`
	__Element__000004_credit_symbol.NameXSD = `credit-symbol`
	__Element__000004_credit_symbol.Type = `xs:string`
	__Element__000004_credit_symbol.MinOccurs = ``
	__Element__000004_credit_symbol.MaxOccurs = ``
	__Element__000004_credit_symbol.Default = ``
	__Element__000004_credit_symbol.Fixed = ``
	__Element__000004_credit_symbol.Nillable = ``
	__Element__000004_credit_symbol.Ref = ``
	__Element__000004_credit_symbol.Abstract = ``
	__Element__000004_credit_symbol.Form = ``
	__Element__000004_credit_symbol.Block = ``
	__Element__000004_credit_symbol.Final = ``

	__Element__000005_credit_symbol.Name = `credit-symbol`
	__Element__000005_credit_symbol.HasNameConflict = false
	__Element__000005_credit_symbol.GoIdentifier = `Credit_symbol`
	__Element__000005_credit_symbol.NameXSD = `credit-symbol`
	__Element__000005_credit_symbol.Type = `xs:string`
	__Element__000005_credit_symbol.MinOccurs = ``
	__Element__000005_credit_symbol.MaxOccurs = ``
	__Element__000005_credit_symbol.Default = ``
	__Element__000005_credit_symbol.Fixed = ``
	__Element__000005_credit_symbol.Nillable = ``
	__Element__000005_credit_symbol.Ref = ``
	__Element__000005_credit_symbol.Abstract = ``
	__Element__000005_credit_symbol.Form = ``
	__Element__000005_credit_symbol.Block = ``
	__Element__000005_credit_symbol.Final = ``

	__Element__000006_credit_type.Name = `credit-type`
	__Element__000006_credit_type.HasNameConflict = false
	__Element__000006_credit_type.GoIdentifier = `Credit_type`
	__Element__000006_credit_type.NameXSD = `credit-type`
	__Element__000006_credit_type.Type = `xs:string`
	__Element__000006_credit_type.MinOccurs = `0`
	__Element__000006_credit_type.MaxOccurs = `unbounded`
	__Element__000006_credit_type.Default = ``
	__Element__000006_credit_type.Fixed = ``
	__Element__000006_credit_type.Nillable = ``
	__Element__000006_credit_type.Ref = ``
	__Element__000006_credit_type.Abstract = ``
	__Element__000006_credit_type.Form = ``
	__Element__000006_credit_type.Block = ``
	__Element__000006_credit_type.Final = ``

	__Element__000007_credit_words.Name = `credit-words`
	__Element__000007_credit_words.HasNameConflict = false
	__Element__000007_credit_words.GoIdentifier = `Credit_words`
	__Element__000007_credit_words.NameXSD = `credit-words`
	__Element__000007_credit_words.Type = `xs:string`
	__Element__000007_credit_words.MinOccurs = ``
	__Element__000007_credit_words.MaxOccurs = ``
	__Element__000007_credit_words.Default = ``
	__Element__000007_credit_words.Fixed = ``
	__Element__000007_credit_words.Nillable = ``
	__Element__000007_credit_words.Ref = ``
	__Element__000007_credit_words.Abstract = ``
	__Element__000007_credit_words.Form = ``
	__Element__000007_credit_words.Block = ``
	__Element__000007_credit_words.Final = ``

	__Element__000008_credit_words.Name = `credit-words`
	__Element__000008_credit_words.HasNameConflict = false
	__Element__000008_credit_words.GoIdentifier = `Credit_words`
	__Element__000008_credit_words.NameXSD = `credit-words`
	__Element__000008_credit_words.Type = `xs:string`
	__Element__000008_credit_words.MinOccurs = ``
	__Element__000008_credit_words.MaxOccurs = ``
	__Element__000008_credit_words.Default = ``
	__Element__000008_credit_words.Fixed = ``
	__Element__000008_credit_words.Nillable = ``
	__Element__000008_credit_words.Ref = ``
	__Element__000008_credit_words.Abstract = ``
	__Element__000008_credit_words.Form = ``
	__Element__000008_credit_words.Block = ``
	__Element__000008_credit_words.Final = ``

	__Element__000009_format.Name = `format`
	__Element__000009_format.HasNameConflict = false
	__Element__000009_format.GoIdentifier = `Format`
	__Element__000009_format.NameXSD = `format`
	__Element__000009_format.Type = `bookFormatEnum`
	__Element__000009_format.MinOccurs = ``
	__Element__000009_format.MaxOccurs = ``
	__Element__000009_format.Default = ``
	__Element__000009_format.Fixed = ``
	__Element__000009_format.Nillable = ``
	__Element__000009_format.Ref = ``
	__Element__000009_format.Abstract = ``
	__Element__000009_format.Form = ``
	__Element__000009_format.Block = ``
	__Element__000009_format.Final = ``

	__Element__000010_link.Name = `link`
	__Element__000010_link.HasNameConflict = false
	__Element__000010_link.GoIdentifier = `Link`
	__Element__000010_link.NameXSD = `link`
	__Element__000010_link.Type = `link`
	__Element__000010_link.MinOccurs = `0`
	__Element__000010_link.MaxOccurs = `unbounded`
	__Element__000010_link.Default = ``
	__Element__000010_link.Fixed = ``
	__Element__000010_link.Nillable = ``
	__Element__000010_link.Ref = ``
	__Element__000010_link.Abstract = ``
	__Element__000010_link.Form = ``
	__Element__000010_link.Block = ``
	__Element__000010_link.Final = ``

	__Element__000011_link.Name = `link`
	__Element__000011_link.HasNameConflict = false
	__Element__000011_link.GoIdentifier = `Link`
	__Element__000011_link.NameXSD = `link`
	__Element__000011_link.Type = `link`
	__Element__000011_link.MinOccurs = `0`
	__Element__000011_link.MaxOccurs = `unbounded`
	__Element__000011_link.Default = ``
	__Element__000011_link.Fixed = ``
	__Element__000011_link.Nillable = ``
	__Element__000011_link.Ref = ``
	__Element__000011_link.Abstract = ``
	__Element__000011_link.Form = ``
	__Element__000011_link.Block = ``
	__Element__000011_link.Final = ``

	__Element__000012_title.Name = `title`
	__Element__000012_title.HasNameConflict = false
	__Element__000012_title.GoIdentifier = `Title`
	__Element__000012_title.NameXSD = `title`
	__Element__000012_title.Type = `titleType`
	__Element__000012_title.MinOccurs = ``
	__Element__000012_title.MaxOccurs = ``
	__Element__000012_title.Default = ``
	__Element__000012_title.Fixed = ``
	__Element__000012_title.Nillable = ``
	__Element__000012_title.Ref = ``
	__Element__000012_title.Abstract = ``
	__Element__000012_title.Form = ``
	__Element__000012_title.Block = ``
	__Element__000012_title.Final = ``

	__Element__000013_year.Name = `year`
	__Element__000013_year.HasNameConflict = false
	__Element__000013_year.GoIdentifier = `Year`
	__Element__000013_year.NameXSD = `year`
	__Element__000013_year.Type = `yearType`
	__Element__000013_year.MinOccurs = ``
	__Element__000013_year.MaxOccurs = ``
	__Element__000013_year.Default = ``
	__Element__000013_year.Fixed = ``
	__Element__000013_year.Nillable = ``
	__Element__000013_year.Ref = ``
	__Element__000013_year.Abstract = ``
	__Element__000013_year.Form = ``
	__Element__000013_year.Block = ``
	__Element__000013_year.Final = ``

	__Enumeration__000000_bookFormatEnum_E_E.Name = `bookFormatEnum_E_E`
	__Enumeration__000000_bookFormatEnum_E_E.Value = `Paperback`

	__Enumeration__000001_bookFormatEnum_E_E.Name = `bookFormatEnum_E_E`
	__Enumeration__000001_bookFormatEnum_E_E.Value = `Hardcover`

	__MaxInclusive__000000_yearType_E_E.Name = `yearType_E_E`
	__MaxInclusive__000000_yearType_E_E.Value = `2100`

	__MinInclusive__000000_yearType_E_E.Name = `yearType_E_E`
	__MinInclusive__000000_yearType_E_E.Value = `1000`

	__Pattern__000000_titleType_E_E.Name = `titleType_E_E`
	__Pattern__000000_titleType_E_E.Value = `[0-9A-Za-z ]+`

	__Restriction__000000_bookFormatEnum_E.Name = `bookFormatEnum_E`
	__Restriction__000000_bookFormatEnum_E.Base = `xs:string`

	__Restriction__000001_titleType_E.Name = `titleType_E`
	__Restriction__000001_titleType_E.Base = `xs:string`

	__Restriction__000002_yearType_E.Name = `yearType_E`
	__Restriction__000002_yearType_E.Base = `xs:integer`

	__Schema__000000_Schema.Name = `Schema`
	__Schema__000000_Schema.Xs = `http://www.w3.org/2001/XMLSchema`

	__Sequence__000000_.Name = ``
	__Sequence__000000_.MinOccurs = ``
	__Sequence__000000_.MaxOccurs = ``

	__Sequence__000001_.Name = ``
	__Sequence__000001_.MinOccurs = `0`
	__Sequence__000001_.MaxOccurs = `unbounded`

	__Sequence__000002_bookType_E.Name = `bookType_E`
	__Sequence__000002_bookType_E.MinOccurs = ``
	__Sequence__000002_bookType_E.MaxOccurs = ``

	__Sequence__000003_books_E_E.Name = `books_E_E`
	__Sequence__000003_books_E_E.MinOccurs = ``
	__Sequence__000003_books_E_E.MaxOccurs = ``

	__Sequence__000004_credit_E.Name = `credit_E`
	__Sequence__000004_credit_E.MinOccurs = ``
	__Sequence__000004_credit_E.MaxOccurs = ``

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
	__Annotation__000001_bestseller_E.Documentations = append(__Annotation__000001_bestseller_E.Documentations, __Documentation__000003__E)
	__Annotation__000002_bookType_E.Documentations = append(__Annotation__000002_bookType_E.Documentations, __Documentation__000005_bookType_E_E)
	__Annotation__000003_book_E.Documentations = append(__Annotation__000003_book_E.Documentations, __Documentation__000006_book_E_E)
	__Annotation__000004_credit_E.Documentations = append(__Annotation__000004_credit_E.Documentations, __Documentation__000007_credit_E_E)
	__Annotation__000005_edition_E.Documentations = append(__Annotation__000005_edition_E.Documentations, __Documentation__000004__E)
	__Annotation__000006_format_E.Documentations = append(__Annotation__000006_format_E.Documentations, __Documentation__000008_format_E_E)
	__Annotation__000007_isbn_E.Documentations = append(__Annotation__000007_isbn_E.Documentations, __Documentation__000002__E)
	__Annotation__000008_link_E.Documentations = append(__Annotation__000008_link_E.Documentations, __Documentation__000009_link_E_E)
	__Annotation__000009_title_E.Documentations = append(__Annotation__000009_title_E.Documentations, __Documentation__000010_title_E_E)
	__Annotation__000010_yearType_E.Documentations = append(__Annotation__000010_yearType_E.Documentations, __Documentation__000011_yearType_E_E)
	__Attribute__000000_bestseller.Annotation = __Annotation__000001_bestseller_E
	__Attribute__000001_edition.Annotation = __Annotation__000005_edition_E
	__Attribute__000002_isbn.Annotation = __Annotation__000007_isbn_E
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000002_isbn)
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000000_bestseller)
	__AttributeGroup__000003_extendedAttributes.AttributeGroups = append(__AttributeGroup__000003_extendedAttributes.AttributeGroups, __AttributeGroup__000001_)
	__AttributeGroup__000003_extendedAttributes.Attributes = append(__AttributeGroup__000003_extendedAttributes.Attributes, __Attribute__000001_edition)
	__Choice__000000_.Sequences = append(__Choice__000000_.Sequences, __Sequence__000000_)
	__Choice__000001_.Elements = append(__Choice__000001_.Elements, __Element__000008_credit_words)
	__Choice__000001_.Elements = append(__Choice__000001_.Elements, __Element__000005_credit_symbol)
	__Choice__000002_.Elements = append(__Choice__000002_.Elements, __Element__000007_credit_words)
	__Choice__000002_.Elements = append(__Choice__000002_.Elements, __Element__000004_credit_symbol)
	__ComplexType__000000_bookType.Annotation = __Annotation__000002_bookType_E
	__ComplexType__000000_bookType.Sequences = append(__ComplexType__000000_bookType.Sequences, __Sequence__000002_bookType_E)
	__ComplexType__000000_bookType.AttributeGroups = append(__ComplexType__000000_bookType.AttributeGroups, __AttributeGroup__000000_)
	__ComplexType__000001_books_E.EnclosingElement = __Element__000002_books
	__ComplexType__000001_books_E.Sequences = append(__ComplexType__000001_books_E.Sequences, __Sequence__000003_books_E_E)
	__ComplexType__000002_credit.Annotation = __Annotation__000004_credit_E
	__ComplexType__000002_credit.Sequences = append(__ComplexType__000002_credit.Sequences, __Sequence__000004_credit_E)
	__ComplexType__000002_credit.Attributes = append(__ComplexType__000002_credit.Attributes, __Attribute__000004_page)
	__ComplexType__000003_link.Annotation = __Annotation__000008_link_E
	__ComplexType__000003_link.Attributes = append(__ComplexType__000003_link.Attributes, __Attribute__000003_name)
	__Element__000001_book.Annotation = __Annotation__000003_book_E
	__Element__000002_books.ComplexType = __ComplexType__000001_books_E
	__Element__000009_format.Annotation = __Annotation__000006_format_E
	__Element__000012_title.Annotation = __Annotation__000009_title_E
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
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000002_credit)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000003_link)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000002_commonAttributes)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000003_extendedAttributes)
	__Sequence__000000_.Sequences = append(__Sequence__000000_.Sequences, __Sequence__000001_)
	__Sequence__000000_.Choices = append(__Sequence__000000_.Choices, __Choice__000002_)
	__Sequence__000001_.Elements = append(__Sequence__000001_.Elements, __Element__000010_link)
	__Sequence__000001_.Choices = append(__Sequence__000001_.Choices, __Choice__000001_)
	__Sequence__000002_bookType_E.Elements = append(__Sequence__000002_bookType_E.Elements, __Element__000012_title)
	__Sequence__000002_bookType_E.Elements = append(__Sequence__000002_bookType_E.Elements, __Element__000000_author)
	__Sequence__000002_bookType_E.Elements = append(__Sequence__000002_bookType_E.Elements, __Element__000013_year)
	__Sequence__000002_bookType_E.Elements = append(__Sequence__000002_bookType_E.Elements, __Element__000009_format)
	__Sequence__000002_bookType_E.Elements = append(__Sequence__000002_bookType_E.Elements, __Element__000003_credit)
	__Sequence__000003_books_E_E.Elements = append(__Sequence__000003_books_E_E.Elements, __Element__000001_book)
	__Sequence__000004_credit_E.Elements = append(__Sequence__000004_credit_E.Elements, __Element__000006_credit_type)
	__Sequence__000004_credit_E.Elements = append(__Sequence__000004_credit_E.Elements, __Element__000011_link)
	__Sequence__000004_credit_E.Choices = append(__Sequence__000004_credit_E.Choices, __Choice__000000_)
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_bookFormatEnum_E
	__SimpleType__000001_titleType.Restriction = __Restriction__000001_titleType_E
	__SimpleType__000002_yearType.Annotation = __Annotation__000010_yearType_E
	__SimpleType__000002_yearType.Restriction = __Restriction__000002_yearType_E
}
