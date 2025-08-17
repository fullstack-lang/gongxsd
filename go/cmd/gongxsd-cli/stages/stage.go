package main

import (
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
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

	const __write__local_time = "2025-08-17 17:44:28.708521 CEST"
	const __write__utc_time__ = "2025-08-17 15:44:28.708521 UTC"

	const __commitId__ = "0000000001"

	// Declaration of instances to stage

	__Annotation__000000_Schema_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000001_book_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000002_yearType_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000003_bookType_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000004_credit_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000005_link_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000006_isbn_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000007_bestseller_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000008_edition_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000009_title_Inlined := (&models.Annotation{}).Stage(stage)
	__Annotation__000010_format_Inlined := (&models.Annotation{}).Stage(stage)

	__Attribute__000000_page := (&models.Attribute{}).Stage(stage)
	__Attribute__000001_name := (&models.Attribute{}).Stage(stage)
	__Attribute__000002_isbn := (&models.Attribute{}).Stage(stage)
	__Attribute__000003_bestseller := (&models.Attribute{}).Stage(stage)
	__Attribute__000004_edition := (&models.Attribute{}).Stage(stage)

	__AttributeGroup__000000_ := (&models.AttributeGroup{}).Stage(stage)
	__AttributeGroup__000001_commonAttributes := (&models.AttributeGroup{}).Stage(stage)
	__AttributeGroup__000002_extendedAttributes := (&models.AttributeGroup{}).Stage(stage)
	__AttributeGroup__000003_ := (&models.AttributeGroup{}).Stage(stage)

	__Choice__000000_credit_S_C_ := (&models.Choice{}).Stage(stage)
	__Choice__000001_credit_S_C_S_S_C_ := (&models.Choice{}).Stage(stage)
	__Choice__000002_credit_S_C_S_C_ := (&models.Choice{}).Stage(stage)

	__ComplexType__000000_A_books := (&models.ComplexType{}).Stage(stage)
	__ComplexType__000001_bookType := (&models.ComplexType{}).Stage(stage)
	__ComplexType__000002_credit := (&models.ComplexType{}).Stage(stage)
	__ComplexType__000003_link := (&models.ComplexType{}).Stage(stage)

	__Documentation__000000_Schema_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000001_Schema_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000002__Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000003_yearType_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000004_bookType_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000005_credit_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000006_link_Inlined_Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000007__Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000008__Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000009__Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000010__Inlined := (&models.Documentation{}).Stage(stage)
	__Documentation__000011__Inlined := (&models.Documentation{}).Stage(stage)

	__Element__000000_books := (&models.Element{}).Stage(stage)
	__Element__000001_book := (&models.Element{}).Stage(stage)
	__Element__000002_credit := (&models.Element{}).Stage(stage)
	__Element__000003_credit_words := (&models.Element{}).Stage(stage)
	__Element__000004_credit_symbol := (&models.Element{}).Stage(stage)
	__Element__000005_link := (&models.Element{}).Stage(stage)
	__Element__000006_credit_words := (&models.Element{}).Stage(stage)
	__Element__000007_credit_symbol := (&models.Element{}).Stage(stage)
	__Element__000008_credit_type := (&models.Element{}).Stage(stage)
	__Element__000009_link := (&models.Element{}).Stage(stage)
	__Element__000010_title := (&models.Element{}).Stage(stage)
	__Element__000011_author := (&models.Element{}).Stage(stage)
	__Element__000012_year := (&models.Element{}).Stage(stage)
	__Element__000013_format := (&models.Element{}).Stage(stage)

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{}).Stage(stage)
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{}).Stage(stage)

	__Extension__000000_link_Inlined_Inlined := (&models.Extension{}).Stage(stage)

	__Group__000000_bookType_S_G_ := (&models.Group{}).Stage(stage)
	__Group__000001_bookDetailsGroup := (&models.Group{}).Stage(stage)

	__MaxInclusive__000000_yearType_Inlined_Inlined := (&models.MaxInclusive{}).Stage(stage)

	__MinInclusive__000000_yearType_Inlined_Inlined := (&models.MinInclusive{}).Stage(stage)

	__Pattern__000000_titleType_Inlined_Inlined := (&models.Pattern{}).Stage(stage)

	__Restriction__000000_yearType_Inlined := (&models.Restriction{}).Stage(stage)
	__Restriction__000001_bookFormatEnum_Inlined := (&models.Restriction{}).Stage(stage)
	__Restriction__000002_titleType_Inlined := (&models.Restriction{}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{}).Stage(stage)

	__Sequence__000000_ := (&models.Sequence{}).Stage(stage)
	__Sequence__000001_bookType_S_ := (&models.Sequence{}).Stage(stage)
	__Sequence__000002_credit_S_ := (&models.Sequence{}).Stage(stage)
	__Sequence__000003_credit_S_C_S_ := (&models.Sequence{}).Stage(stage)
	__Sequence__000004_credit_S_C_S_S_ := (&models.Sequence{}).Stage(stage)
	__Sequence__000005_bookDetailsGroup_S_ := (&models.Sequence{}).Stage(stage)

	__SimpleContent__000000_link_Inlined := (&models.SimpleContent{}).Stage(stage)

	__SimpleType__000000_yearType := (&models.SimpleType{}).Stage(stage)
	__SimpleType__000001_bookFormatEnum := (&models.SimpleType{}).Stage(stage)
	__SimpleType__000002_titleType := (&models.SimpleType{}).Stage(stage)

	__WhiteSpace__000000_titleType_Inlined_Inlined := (&models.WhiteSpace{}).Stage(stage)

	// Setup of values

	__Annotation__000000_Schema_Inlined.Name = `Schema_Inlined`

	__Annotation__000001_book_Inlined.Name = `book_Inlined`

	__Annotation__000002_yearType_Inlined.Name = `yearType_Inlined`

	__Annotation__000003_bookType_Inlined.Name = `bookType_Inlined`

	__Annotation__000004_credit_Inlined.Name = `credit_Inlined`

	__Annotation__000005_link_Inlined.Name = `link_Inlined`

	__Annotation__000006_isbn_Inlined.Name = `isbn_Inlined`

	__Annotation__000007_bestseller_Inlined.Name = `bestseller_Inlined`

	__Annotation__000008_edition_Inlined.Name = `edition_Inlined`

	__Annotation__000009_title_Inlined.Name = `title_Inlined`

	__Annotation__000010_format_Inlined.Name = `format_Inlined`

	__Attribute__000000_page.Name = `page`
	__Attribute__000000_page.NameXSD = `page`
	__Attribute__000000_page.Type = `positiveInteger`
	__Attribute__000000_page.HasNameConflict = false
	__Attribute__000000_page.GoIdentifier = `Page`
	__Attribute__000000_page.Default = ``
	__Attribute__000000_page.Use = ``
	__Attribute__000000_page.Form = ``
	__Attribute__000000_page.Fixed = ``
	__Attribute__000000_page.Ref = ``
	__Attribute__000000_page.TargetNamespace = ``
	__Attribute__000000_page.SimpleType = ``
	__Attribute__000000_page.IDXSD = ``

	__Attribute__000001_name.Name = `name`
	__Attribute__000001_name.NameXSD = `name`
	__Attribute__000001_name.Type = `token`
	__Attribute__000001_name.HasNameConflict = false
	__Attribute__000001_name.GoIdentifier = `NameXSD`
	__Attribute__000001_name.Default = ``
	__Attribute__000001_name.Use = ``
	__Attribute__000001_name.Form = ``
	__Attribute__000001_name.Fixed = ``
	__Attribute__000001_name.Ref = ``
	__Attribute__000001_name.TargetNamespace = ``
	__Attribute__000001_name.SimpleType = ``
	__Attribute__000001_name.IDXSD = ``

	__Attribute__000002_isbn.Name = `isbn`
	__Attribute__000002_isbn.NameXSD = `isbn`
	__Attribute__000002_isbn.Type = `string`
	__Attribute__000002_isbn.HasNameConflict = false
	__Attribute__000002_isbn.GoIdentifier = `Isbn`
	__Attribute__000002_isbn.Default = ``
	__Attribute__000002_isbn.Use = `required`
	__Attribute__000002_isbn.Form = ``
	__Attribute__000002_isbn.Fixed = ``
	__Attribute__000002_isbn.Ref = ``
	__Attribute__000002_isbn.TargetNamespace = ``
	__Attribute__000002_isbn.SimpleType = ``
	__Attribute__000002_isbn.IDXSD = ``

	__Attribute__000003_bestseller.Name = `bestseller`
	__Attribute__000003_bestseller.NameXSD = `bestseller`
	__Attribute__000003_bestseller.Type = `boolean`
	__Attribute__000003_bestseller.HasNameConflict = false
	__Attribute__000003_bestseller.GoIdentifier = `Bestseller`
	__Attribute__000003_bestseller.Default = ``
	__Attribute__000003_bestseller.Use = `optional`
	__Attribute__000003_bestseller.Form = ``
	__Attribute__000003_bestseller.Fixed = ``
	__Attribute__000003_bestseller.Ref = ``
	__Attribute__000003_bestseller.TargetNamespace = ``
	__Attribute__000003_bestseller.SimpleType = ``
	__Attribute__000003_bestseller.IDXSD = ``

	__Attribute__000004_edition.Name = `edition`
	__Attribute__000004_edition.NameXSD = `edition`
	__Attribute__000004_edition.Type = `string`
	__Attribute__000004_edition.HasNameConflict = false
	__Attribute__000004_edition.GoIdentifier = `Edition`
	__Attribute__000004_edition.Default = ``
	__Attribute__000004_edition.Use = `optional`
	__Attribute__000004_edition.Form = ``
	__Attribute__000004_edition.Fixed = ``
	__Attribute__000004_edition.Ref = ``
	__Attribute__000004_edition.TargetNamespace = ``
	__Attribute__000004_edition.SimpleType = ``
	__Attribute__000004_edition.IDXSD = ``

	__AttributeGroup__000000_.Name = ``
	__AttributeGroup__000000_.NameXSD = ``
	__AttributeGroup__000000_.HasNameConflict = true
	__AttributeGroup__000000_.GoIdentifier = `AttributeGroup__1`
	__AttributeGroup__000000_.Ref = `extendedAttributes`
	__AttributeGroup__000000_.Order = 5
	__AttributeGroup__000000_.Depth = 1

	__AttributeGroup__000001_commonAttributes.Name = `commonAttributes`
	__AttributeGroup__000001_commonAttributes.NameXSD = `commonAttributes`
	__AttributeGroup__000001_commonAttributes.HasNameConflict = false
	__AttributeGroup__000001_commonAttributes.GoIdentifier = `AttributeGroup_commonAttributes`
	__AttributeGroup__000001_commonAttributes.Ref = ``
	__AttributeGroup__000001_commonAttributes.Order = 16
	__AttributeGroup__000001_commonAttributes.Depth = 0

	__AttributeGroup__000002_extendedAttributes.Name = `extendedAttributes`
	__AttributeGroup__000002_extendedAttributes.NameXSD = `extendedAttributes`
	__AttributeGroup__000002_extendedAttributes.HasNameConflict = false
	__AttributeGroup__000002_extendedAttributes.GoIdentifier = `AttributeGroup_extendedAttributes`
	__AttributeGroup__000002_extendedAttributes.Ref = ``
	__AttributeGroup__000002_extendedAttributes.Order = 17
	__AttributeGroup__000002_extendedAttributes.Depth = 0

	__AttributeGroup__000003_.Name = ``
	__AttributeGroup__000003_.NameXSD = ``
	__AttributeGroup__000003_.HasNameConflict = false
	__AttributeGroup__000003_.GoIdentifier = `AttributeGroup_`
	__AttributeGroup__000003_.Ref = `commonAttributes`
	__AttributeGroup__000003_.Order = 18
	__AttributeGroup__000003_.Depth = 1

	__Choice__000000_credit_S_C_.Name = `credit_S__C_`
	__Choice__000000_credit_S_C_.OuterElementName = `credit_S__C_`
	__Choice__000000_credit_S_C_.Order = 0
	__Choice__000000_credit_S_C_.Depth = 0
	__Choice__000000_credit_S_C_.MinOccurs = ``
	__Choice__000000_credit_S_C_.MaxOccurs = ``
	__Choice__000000_credit_S_C_.IsDuplicatedInXSD = false

	__Choice__000001_credit_S_C_S_S_C_.Name = `credit_S__C__S__S__C_`
	__Choice__000001_credit_S_C_S_S_C_.OuterElementName = `credit_S__C__S__S__C_`
	__Choice__000001_credit_S_C_S_S_C_.Order = 0
	__Choice__000001_credit_S_C_S_S_C_.Depth = 0
	__Choice__000001_credit_S_C_S_S_C_.MinOccurs = ``
	__Choice__000001_credit_S_C_S_S_C_.MaxOccurs = ``
	__Choice__000001_credit_S_C_S_S_C_.IsDuplicatedInXSD = false

	__Choice__000002_credit_S_C_S_C_.Name = `credit_S__C__S__C_`
	__Choice__000002_credit_S_C_S_C_.OuterElementName = `credit_S__C__S__C_`
	__Choice__000002_credit_S_C_S_C_.Order = 0
	__Choice__000002_credit_S_C_S_C_.Depth = 0
	__Choice__000002_credit_S_C_S_C_.MinOccurs = ``
	__Choice__000002_credit_S_C_S_C_.MaxOccurs = ``
	__Choice__000002_credit_S_C_S_C_.IsDuplicatedInXSD = false

	__ComplexType__000000_A_books.Name = `A_books`
	__ComplexType__000000_A_books.HasNameConflict = false
	__ComplexType__000000_A_books.GoIdentifier = `A_books`
	__ComplexType__000000_A_books.IsAnonymous = true
	__ComplexType__000000_A_books.NameXSD = ``
	__ComplexType__000000_A_books.OuterElementName = `A_books`
	__ComplexType__000000_A_books.Order = 14
	__ComplexType__000000_A_books.Depth = 1
	__ComplexType__000000_A_books.MinOccurs = ``
	__ComplexType__000000_A_books.MaxOccurs = ``
	__ComplexType__000000_A_books.IsDuplicatedInXSD = false

	__ComplexType__000001_bookType.Name = `bookType`
	__ComplexType__000001_bookType.HasNameConflict = false
	__ComplexType__000001_bookType.GoIdentifier = `BookType`
	__ComplexType__000001_bookType.IsAnonymous = false
	__ComplexType__000001_bookType.NameXSD = `bookType`
	__ComplexType__000001_bookType.OuterElementName = `bookType`
	__ComplexType__000001_bookType.Order = 2
	__ComplexType__000001_bookType.Depth = 0
	__ComplexType__000001_bookType.MinOccurs = ``
	__ComplexType__000001_bookType.MaxOccurs = ``
	__ComplexType__000001_bookType.IsDuplicatedInXSD = false

	__ComplexType__000002_credit.Name = `credit`
	__ComplexType__000002_credit.HasNameConflict = false
	__ComplexType__000002_credit.GoIdentifier = `Credit`
	__ComplexType__000002_credit.IsAnonymous = false
	__ComplexType__000002_credit.NameXSD = `credit`
	__ComplexType__000002_credit.OuterElementName = `credit`
	__ComplexType__000002_credit.Order = 19
	__ComplexType__000002_credit.Depth = 0
	__ComplexType__000002_credit.MinOccurs = ``
	__ComplexType__000002_credit.MaxOccurs = ``
	__ComplexType__000002_credit.IsDuplicatedInXSD = false

	__ComplexType__000003_link.Name = `link`
	__ComplexType__000003_link.HasNameConflict = false
	__ComplexType__000003_link.GoIdentifier = `Link`
	__ComplexType__000003_link.IsAnonymous = false
	__ComplexType__000003_link.NameXSD = `link`
	__ComplexType__000003_link.OuterElementName = `link`
	__ComplexType__000003_link.Order = 27
	__ComplexType__000003_link.Depth = 0
	__ComplexType__000003_link.MinOccurs = ``
	__ComplexType__000003_link.MaxOccurs = ``
	__ComplexType__000003_link.IsDuplicatedInXSD = false

	__Documentation__000000_Schema_Inlined_Inlined.Name = `Schema_Inlined_Inlined`
	__Documentation__000000_Schema_Inlined_Inlined.Text = ` This schema defines
            the structure of a simple book collection. It includes types for book details, such as
            title, author, year, and format. `
	__Documentation__000000_Schema_Inlined_Inlined.Source = `http://example.com/schema-docs`
	__Documentation__000000_Schema_Inlined_Inlined.Lang = `en`

	__Documentation__000001_Schema_Inlined_Inlined.Name = `Schema_Inlined_Inlined`
	__Documentation__000001_Schema_Inlined_Inlined.Text = ` Ce schéma définit
            la structure d'une collection de livres simple. Il inclut des types pour les détails du
            livre, tels que le titre, l'auteur, l'année et le format. `
	__Documentation__000001_Schema_Inlined_Inlined.Source = `http://example.com/schema-docs`
	__Documentation__000001_Schema_Inlined_Inlined.Lang = `fr`

	__Documentation__000002__Inlined.Name = `_Inlined`
	__Documentation__000002__Inlined.Text = ` A book element representing a single book in the
                            collection. `
	__Documentation__000002__Inlined.Source = ``
	__Documentation__000002__Inlined.Lang = ``

	__Documentation__000003_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__Documentation__000003_yearType_Inlined_Inlined.Text = ` This type represents a year. It restricts the value to an integer
                between 1000 and 2100 inclusive. `
	__Documentation__000003_yearType_Inlined_Inlined.Source = ``
	__Documentation__000003_yearType_Inlined_Inlined.Lang = ``

	__Documentation__000004_bookType_Inlined_Inlined.Name = `bookType_Inlined_Inlined`
	__Documentation__000004_bookType_Inlined_Inlined.Text = ` This complex type defines the structure of a book, including title,
                author, year, and format. `
	__Documentation__000004_bookType_Inlined_Inlined.Source = ``
	__Documentation__000004_bookType_Inlined_Inlined.Lang = ``

	__Documentation__000005_credit_Inlined_Inlined.Name = `credit_Inlined_Inlined`
	__Documentation__000005_credit_Inlined_Inlined.Text = `The credit type .. `
	__Documentation__000005_credit_Inlined_Inlined.Source = ``
	__Documentation__000005_credit_Inlined_Inlined.Lang = ``

	__Documentation__000006_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Documentation__000006_link_Inlined_Inlined.Text = `The link type serves as an outgoing simple XLink. If a relative link
                is used within a document that is part of a compressed MusicXML file, the link is
                relative to the root folder of the zip file.`
	__Documentation__000006_link_Inlined_Inlined.Source = ``
	__Documentation__000006_link_Inlined_Inlined.Lang = ``

	__Documentation__000007__Inlined.Name = `_Inlined`
	__Documentation__000007__Inlined.Text = `The ISBN number of the book.`
	__Documentation__000007__Inlined.Source = ``
	__Documentation__000007__Inlined.Lang = ``

	__Documentation__000008__Inlined.Name = `_Inlined`
	__Documentation__000008__Inlined.Text = `Indicates if the book is a bestseller.`
	__Documentation__000008__Inlined.Source = ``
	__Documentation__000008__Inlined.Lang = ``

	__Documentation__000009__Inlined.Name = `_Inlined`
	__Documentation__000009__Inlined.Text = `The edition of the book.`
	__Documentation__000009__Inlined.Source = ``
	__Documentation__000009__Inlined.Lang = ``

	__Documentation__000010__Inlined.Name = `_Inlined`
	__Documentation__000010__Inlined.Text = `The title of the book, consisting of alphabetic characters and
                        spaces only.`
	__Documentation__000010__Inlined.Source = ``
	__Documentation__000010__Inlined.Lang = ``

	__Documentation__000011__Inlined.Name = `_Inlined`
	__Documentation__000011__Inlined.Text = `The format of the book, either Paperback or Hardcover.`
	__Documentation__000011__Inlined.Source = ``
	__Documentation__000011__Inlined.Lang = ``

	__Element__000000_books.Name = `books`
	__Element__000000_books.Order = 13
	__Element__000000_books.Depth = 0
	__Element__000000_books.HasNameConflict = false
	__Element__000000_books.GoIdentifier = `Books`
	__Element__000000_books.NameXSD = `books`
	__Element__000000_books.Type = ``
	__Element__000000_books.MinOccurs = ``
	__Element__000000_books.MaxOccurs = ``
	__Element__000000_books.Default = ``
	__Element__000000_books.Fixed = ``
	__Element__000000_books.Nillable = ``
	__Element__000000_books.Ref = ``
	__Element__000000_books.Abstract = ``
	__Element__000000_books.Form = ``
	__Element__000000_books.Block = ``
	__Element__000000_books.Final = ``
	__Element__000000_books.IsDuplicatedInXSD = false

	__Element__000001_book.Name = `book`
	__Element__000001_book.Order = 15
	__Element__000001_book.Depth = 2
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
	__Element__000001_book.IsDuplicatedInXSD = false

	__Element__000002_credit.Name = `credit`
	__Element__000002_credit.Order = 4
	__Element__000002_credit.Depth = 1
	__Element__000002_credit.HasNameConflict = false
	__Element__000002_credit.GoIdentifier = `Credit`
	__Element__000002_credit.NameXSD = `credit`
	__Element__000002_credit.Type = `credit`
	__Element__000002_credit.MinOccurs = `0`
	__Element__000002_credit.MaxOccurs = `unbounded`
	__Element__000002_credit.Default = ``
	__Element__000002_credit.Fixed = ``
	__Element__000002_credit.Nillable = ``
	__Element__000002_credit.Ref = ``
	__Element__000002_credit.Abstract = ``
	__Element__000002_credit.Form = ``
	__Element__000002_credit.Block = ``
	__Element__000002_credit.Final = ``
	__Element__000002_credit.IsDuplicatedInXSD = false

	__Element__000003_credit_words.Name = `credit-words`
	__Element__000003_credit_words.Order = 25
	__Element__000003_credit_words.Depth = 1
	__Element__000003_credit_words.HasNameConflict = false
	__Element__000003_credit_words.GoIdentifier = `Credit_words`
	__Element__000003_credit_words.NameXSD = `credit-words`
	__Element__000003_credit_words.Type = `string`
	__Element__000003_credit_words.MinOccurs = ``
	__Element__000003_credit_words.MaxOccurs = ``
	__Element__000003_credit_words.Default = ``
	__Element__000003_credit_words.Fixed = ``
	__Element__000003_credit_words.Nillable = ``
	__Element__000003_credit_words.Ref = ``
	__Element__000003_credit_words.Abstract = ``
	__Element__000003_credit_words.Form = ``
	__Element__000003_credit_words.Block = ``
	__Element__000003_credit_words.Final = ``
	__Element__000003_credit_words.IsDuplicatedInXSD = false

	__Element__000004_credit_symbol.Name = `credit-symbol`
	__Element__000004_credit_symbol.Order = 26
	__Element__000004_credit_symbol.Depth = 1
	__Element__000004_credit_symbol.HasNameConflict = false
	__Element__000004_credit_symbol.GoIdentifier = `Credit_symbol`
	__Element__000004_credit_symbol.NameXSD = `credit-symbol`
	__Element__000004_credit_symbol.Type = `string`
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
	__Element__000004_credit_symbol.IsDuplicatedInXSD = false

	__Element__000005_link.Name = `link`
	__Element__000005_link.Order = 24
	__Element__000005_link.Depth = 1
	__Element__000005_link.HasNameConflict = false
	__Element__000005_link.GoIdentifier = `Link`
	__Element__000005_link.NameXSD = `link`
	__Element__000005_link.Type = `link`
	__Element__000005_link.MinOccurs = `0`
	__Element__000005_link.MaxOccurs = `unbounded`
	__Element__000005_link.Default = ``
	__Element__000005_link.Fixed = ``
	__Element__000005_link.Nillable = ``
	__Element__000005_link.Ref = ``
	__Element__000005_link.Abstract = ``
	__Element__000005_link.Form = ``
	__Element__000005_link.Block = ``
	__Element__000005_link.Final = ``
	__Element__000005_link.IsDuplicatedInXSD = false

	__Element__000006_credit_words.Name = `credit-words`
	__Element__000006_credit_words.Order = 22
	__Element__000006_credit_words.Depth = 1
	__Element__000006_credit_words.HasNameConflict = false
	__Element__000006_credit_words.GoIdentifier = `Credit_words`
	__Element__000006_credit_words.NameXSD = `credit-words`
	__Element__000006_credit_words.Type = `xs:string`
	__Element__000006_credit_words.MinOccurs = ``
	__Element__000006_credit_words.MaxOccurs = ``
	__Element__000006_credit_words.Default = ``
	__Element__000006_credit_words.Fixed = ``
	__Element__000006_credit_words.Nillable = ``
	__Element__000006_credit_words.Ref = ``
	__Element__000006_credit_words.Abstract = ``
	__Element__000006_credit_words.Form = ``
	__Element__000006_credit_words.Block = ``
	__Element__000006_credit_words.Final = ``
	__Element__000006_credit_words.IsDuplicatedInXSD = false

	__Element__000007_credit_symbol.Name = `credit-symbol`
	__Element__000007_credit_symbol.Order = 23
	__Element__000007_credit_symbol.Depth = 1
	__Element__000007_credit_symbol.HasNameConflict = false
	__Element__000007_credit_symbol.GoIdentifier = `Credit_symbol`
	__Element__000007_credit_symbol.NameXSD = `credit-symbol`
	__Element__000007_credit_symbol.Type = `xs:string`
	__Element__000007_credit_symbol.MinOccurs = ``
	__Element__000007_credit_symbol.MaxOccurs = ``
	__Element__000007_credit_symbol.Default = ``
	__Element__000007_credit_symbol.Fixed = ``
	__Element__000007_credit_symbol.Nillable = ``
	__Element__000007_credit_symbol.Ref = ``
	__Element__000007_credit_symbol.Abstract = ``
	__Element__000007_credit_symbol.Form = ``
	__Element__000007_credit_symbol.Block = ``
	__Element__000007_credit_symbol.Final = ``
	__Element__000007_credit_symbol.IsDuplicatedInXSD = false

	__Element__000008_credit_type.Name = `credit-type`
	__Element__000008_credit_type.Order = 20
	__Element__000008_credit_type.Depth = 1
	__Element__000008_credit_type.HasNameConflict = false
	__Element__000008_credit_type.GoIdentifier = `Credit_type`
	__Element__000008_credit_type.NameXSD = `credit-type`
	__Element__000008_credit_type.Type = `string`
	__Element__000008_credit_type.MinOccurs = `0`
	__Element__000008_credit_type.MaxOccurs = `unbounded`
	__Element__000008_credit_type.Default = ``
	__Element__000008_credit_type.Fixed = ``
	__Element__000008_credit_type.Nillable = ``
	__Element__000008_credit_type.Ref = ``
	__Element__000008_credit_type.Abstract = ``
	__Element__000008_credit_type.Form = ``
	__Element__000008_credit_type.Block = ``
	__Element__000008_credit_type.Final = ``
	__Element__000008_credit_type.IsDuplicatedInXSD = false

	__Element__000009_link.Name = `link`
	__Element__000009_link.Order = 21
	__Element__000009_link.Depth = 1
	__Element__000009_link.HasNameConflict = false
	__Element__000009_link.GoIdentifier = `Link`
	__Element__000009_link.NameXSD = `link`
	__Element__000009_link.Type = `link`
	__Element__000009_link.MinOccurs = `0`
	__Element__000009_link.MaxOccurs = `unbounded`
	__Element__000009_link.Default = ``
	__Element__000009_link.Fixed = ``
	__Element__000009_link.Nillable = ``
	__Element__000009_link.Ref = ``
	__Element__000009_link.Abstract = ``
	__Element__000009_link.Form = ``
	__Element__000009_link.Block = ``
	__Element__000009_link.Final = ``
	__Element__000009_link.IsDuplicatedInXSD = false

	__Element__000010_title.Name = `title`
	__Element__000010_title.Order = 7
	__Element__000010_title.Depth = 1
	__Element__000010_title.HasNameConflict = false
	__Element__000010_title.GoIdentifier = `Title`
	__Element__000010_title.NameXSD = `title`
	__Element__000010_title.Type = `titleType`
	__Element__000010_title.MinOccurs = ``
	__Element__000010_title.MaxOccurs = ``
	__Element__000010_title.Default = ``
	__Element__000010_title.Fixed = ``
	__Element__000010_title.Nillable = ``
	__Element__000010_title.Ref = ``
	__Element__000010_title.Abstract = ``
	__Element__000010_title.Form = ``
	__Element__000010_title.Block = ``
	__Element__000010_title.Final = ``
	__Element__000010_title.IsDuplicatedInXSD = false

	__Element__000011_author.Name = `author`
	__Element__000011_author.Order = 8
	__Element__000011_author.Depth = 1
	__Element__000011_author.HasNameConflict = false
	__Element__000011_author.GoIdentifier = `Author`
	__Element__000011_author.NameXSD = `author`
	__Element__000011_author.Type = `string`
	__Element__000011_author.MinOccurs = ``
	__Element__000011_author.MaxOccurs = ``
	__Element__000011_author.Default = ``
	__Element__000011_author.Fixed = ``
	__Element__000011_author.Nillable = ``
	__Element__000011_author.Ref = ``
	__Element__000011_author.Abstract = ``
	__Element__000011_author.Form = ``
	__Element__000011_author.Block = ``
	__Element__000011_author.Final = ``
	__Element__000011_author.IsDuplicatedInXSD = false

	__Element__000012_year.Name = `year`
	__Element__000012_year.Order = 9
	__Element__000012_year.Depth = 1
	__Element__000012_year.HasNameConflict = false
	__Element__000012_year.GoIdentifier = `Year`
	__Element__000012_year.NameXSD = `year`
	__Element__000012_year.Type = `yearType`
	__Element__000012_year.MinOccurs = ``
	__Element__000012_year.MaxOccurs = ``
	__Element__000012_year.Default = ``
	__Element__000012_year.Fixed = ``
	__Element__000012_year.Nillable = ``
	__Element__000012_year.Ref = ``
	__Element__000012_year.Abstract = ``
	__Element__000012_year.Form = ``
	__Element__000012_year.Block = ``
	__Element__000012_year.Final = ``
	__Element__000012_year.IsDuplicatedInXSD = false

	__Element__000013_format.Name = `format`
	__Element__000013_format.Order = 10
	__Element__000013_format.Depth = 1
	__Element__000013_format.HasNameConflict = false
	__Element__000013_format.GoIdentifier = `Format`
	__Element__000013_format.NameXSD = `format`
	__Element__000013_format.Type = `bookFormatEnum`
	__Element__000013_format.MinOccurs = ``
	__Element__000013_format.MaxOccurs = ``
	__Element__000013_format.Default = ``
	__Element__000013_format.Fixed = ``
	__Element__000013_format.Nillable = ``
	__Element__000013_format.Ref = ``
	__Element__000013_format.Abstract = ``
	__Element__000013_format.Form = ``
	__Element__000013_format.Block = ``
	__Element__000013_format.Final = ``
	__Element__000013_format.IsDuplicatedInXSD = false

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Value = `Paperback`

	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Value = `Hardcover`

	__Extension__000000_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Extension__000000_link_Inlined_Inlined.OuterElementName = ``
	__Extension__000000_link_Inlined_Inlined.Order = 0
	__Extension__000000_link_Inlined_Inlined.Depth = 0
	__Extension__000000_link_Inlined_Inlined.MinOccurs = ``
	__Extension__000000_link_Inlined_Inlined.MaxOccurs = ``
	__Extension__000000_link_Inlined_Inlined.Base = `string`
	__Extension__000000_link_Inlined_Inlined.Ref = ``

	__Group__000000_bookType_S_G_.Name = `bookType_S__G_`
	__Group__000000_bookType_S_G_.NameXSD = ``
	__Group__000000_bookType_S_G_.Ref = `bookDetailsGroup`
	__Group__000000_bookType_S_G_.IsAnonymous = false
	__Group__000000_bookType_S_G_.HasNameConflict = false
	__Group__000000_bookType_S_G_.GoIdentifier = `BookDetailsGroup`
	__Group__000000_bookType_S_G_.OuterElementName = `bookType_S__G_`
	__Group__000000_bookType_S_G_.Order = 3
	__Group__000000_bookType_S_G_.Depth = 1
	__Group__000000_bookType_S_G_.MinOccurs = ``
	__Group__000000_bookType_S_G_.MaxOccurs = ``

	__Group__000001_bookDetailsGroup.Name = `bookDetailsGroup`
	__Group__000001_bookDetailsGroup.NameXSD = `bookDetailsGroup`
	__Group__000001_bookDetailsGroup.Ref = ``
	__Group__000001_bookDetailsGroup.IsAnonymous = false
	__Group__000001_bookDetailsGroup.HasNameConflict = false
	__Group__000001_bookDetailsGroup.GoIdentifier = `Group_bookDetailsGroup`
	__Group__000001_bookDetailsGroup.OuterElementName = `bookDetailsGroup`
	__Group__000001_bookDetailsGroup.Order = 6
	__Group__000001_bookDetailsGroup.Depth = 0
	__Group__000001_bookDetailsGroup.MinOccurs = ``
	__Group__000001_bookDetailsGroup.MaxOccurs = ``

	__MaxInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MaxInclusive__000000_yearType_Inlined_Inlined.Value = `2100`

	__MinInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MinInclusive__000000_yearType_Inlined_Inlined.Value = `1000`

	__Pattern__000000_titleType_Inlined_Inlined.Name = `titleType_Inlined_Inlined`
	__Pattern__000000_titleType_Inlined_Inlined.Value = `[0-9A-Za-z ]+`

	__Restriction__000000_yearType_Inlined.Name = `yearType_Inlined`
	__Restriction__000000_yearType_Inlined.Base = `integer`

	__Restriction__000001_bookFormatEnum_Inlined.Name = `bookFormatEnum_Inlined`
	__Restriction__000001_bookFormatEnum_Inlined.Base = `string`

	__Restriction__000002_titleType_Inlined.Name = `titleType_Inlined`
	__Restriction__000002_titleType_Inlined.Base = `string`

	__Schema__000000_Schema.Name = `Schema`
	__Schema__000000_Schema.Xs = `http://www.w3.org/2001/XMLSchema`
	__Schema__000000_Schema.Order = 0
	__Schema__000000_Schema.Depth = 0

	__Sequence__000000_.Name = ``
	__Sequence__000000_.OuterElementName = ``
	__Sequence__000000_.Order = 0
	__Sequence__000000_.Depth = 0
	__Sequence__000000_.MinOccurs = ``
	__Sequence__000000_.MaxOccurs = ``

	__Sequence__000001_bookType_S_.Name = `bookType_S_`
	__Sequence__000001_bookType_S_.OuterElementName = `bookType_S_`
	__Sequence__000001_bookType_S_.Order = 0
	__Sequence__000001_bookType_S_.Depth = 0
	__Sequence__000001_bookType_S_.MinOccurs = ``
	__Sequence__000001_bookType_S_.MaxOccurs = ``

	__Sequence__000002_credit_S_.Name = `credit_S_`
	__Sequence__000002_credit_S_.OuterElementName = `credit_S_`
	__Sequence__000002_credit_S_.Order = 0
	__Sequence__000002_credit_S_.Depth = 0
	__Sequence__000002_credit_S_.MinOccurs = ``
	__Sequence__000002_credit_S_.MaxOccurs = ``

	__Sequence__000003_credit_S_C_S_.Name = `credit_S__C__S_`
	__Sequence__000003_credit_S_C_S_.OuterElementName = `credit_S__C__S_`
	__Sequence__000003_credit_S_C_S_.Order = 0
	__Sequence__000003_credit_S_C_S_.Depth = 0
	__Sequence__000003_credit_S_C_S_.MinOccurs = ``
	__Sequence__000003_credit_S_C_S_.MaxOccurs = ``

	__Sequence__000004_credit_S_C_S_S_.Name = `credit_S__C__S__S_`
	__Sequence__000004_credit_S_C_S_S_.OuterElementName = `credit_S__C__S__S_`
	__Sequence__000004_credit_S_C_S_S_.Order = 0
	__Sequence__000004_credit_S_C_S_S_.Depth = 0
	__Sequence__000004_credit_S_C_S_S_.MinOccurs = `0`
	__Sequence__000004_credit_S_C_S_S_.MaxOccurs = `unbounded`

	__Sequence__000005_bookDetailsGroup_S_.Name = `bookDetailsGroup_S_`
	__Sequence__000005_bookDetailsGroup_S_.OuterElementName = `bookDetailsGroup_S_`
	__Sequence__000005_bookDetailsGroup_S_.Order = 0
	__Sequence__000005_bookDetailsGroup_S_.Depth = 0
	__Sequence__000005_bookDetailsGroup_S_.MinOccurs = ``
	__Sequence__000005_bookDetailsGroup_S_.MaxOccurs = ``

	__SimpleContent__000000_link_Inlined.Name = `link_Inlined`

	__SimpleType__000000_yearType.Name = `yearType`
	__SimpleType__000000_yearType.NameXSD = `yearType`
	__SimpleType__000000_yearType.Order = 1
	__SimpleType__000000_yearType.Depth = 0

	__SimpleType__000001_bookFormatEnum.Name = `bookFormatEnum`
	__SimpleType__000001_bookFormatEnum.NameXSD = `bookFormatEnum`
	__SimpleType__000001_bookFormatEnum.Order = 11
	__SimpleType__000001_bookFormatEnum.Depth = 0

	__SimpleType__000002_titleType.Name = `titleType`
	__SimpleType__000002_titleType.NameXSD = `titleType`
	__SimpleType__000002_titleType.Order = 12
	__SimpleType__000002_titleType.Depth = 0

	__WhiteSpace__000000_titleType_Inlined_Inlined.Name = `titleType_Inlined_Inlined`
	__WhiteSpace__000000_titleType_Inlined_Inlined.Value = `collapse`

	// Setup of pointers
	// setup of Annotation instances pointers
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000000_Schema_Inlined_Inlined)
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000001_Schema_Inlined_Inlined)
	__Annotation__000001_book_Inlined.Documentations = append(__Annotation__000001_book_Inlined.Documentations, __Documentation__000002__Inlined)
	__Annotation__000002_yearType_Inlined.Documentations = append(__Annotation__000002_yearType_Inlined.Documentations, __Documentation__000003_yearType_Inlined_Inlined)
	__Annotation__000003_bookType_Inlined.Documentations = append(__Annotation__000003_bookType_Inlined.Documentations, __Documentation__000004_bookType_Inlined_Inlined)
	__Annotation__000004_credit_Inlined.Documentations = append(__Annotation__000004_credit_Inlined.Documentations, __Documentation__000005_credit_Inlined_Inlined)
	__Annotation__000005_link_Inlined.Documentations = append(__Annotation__000005_link_Inlined.Documentations, __Documentation__000006_link_Inlined_Inlined)
	__Annotation__000006_isbn_Inlined.Documentations = append(__Annotation__000006_isbn_Inlined.Documentations, __Documentation__000007__Inlined)
	__Annotation__000007_bestseller_Inlined.Documentations = append(__Annotation__000007_bestseller_Inlined.Documentations, __Documentation__000008__Inlined)
	__Annotation__000008_edition_Inlined.Documentations = append(__Annotation__000008_edition_Inlined.Documentations, __Documentation__000009__Inlined)
	__Annotation__000009_title_Inlined.Documentations = append(__Annotation__000009_title_Inlined.Documentations, __Documentation__000010__Inlined)
	__Annotation__000010_format_Inlined.Documentations = append(__Annotation__000010_format_Inlined.Documentations, __Documentation__000011__Inlined)
	// setup of Attribute instances pointers
	__Attribute__000002_isbn.Annotation = __Annotation__000006_isbn_Inlined
	__Attribute__000003_bestseller.Annotation = __Annotation__000007_bestseller_Inlined
	__Attribute__000004_edition.Annotation = __Annotation__000008_edition_Inlined
	// setup of AttributeGroup instances pointers
	__AttributeGroup__000001_commonAttributes.Attributes = append(__AttributeGroup__000001_commonAttributes.Attributes, __Attribute__000002_isbn)
	__AttributeGroup__000001_commonAttributes.Attributes = append(__AttributeGroup__000001_commonAttributes.Attributes, __Attribute__000003_bestseller)
	__AttributeGroup__000002_extendedAttributes.AttributeGroups = append(__AttributeGroup__000002_extendedAttributes.AttributeGroups, __AttributeGroup__000003_)
	__AttributeGroup__000002_extendedAttributes.Attributes = append(__AttributeGroup__000002_extendedAttributes.Attributes, __Attribute__000004_edition)
	// setup of Choice instances pointers
	__Choice__000000_credit_S_C_.Sequences = append(__Choice__000000_credit_S_C_.Sequences, __Sequence__000003_credit_S_C_S_)
	__Choice__000001_credit_S_C_S_S_C_.Elements = append(__Choice__000001_credit_S_C_S_S_C_.Elements, __Element__000003_credit_words)
	__Choice__000001_credit_S_C_S_S_C_.Elements = append(__Choice__000001_credit_S_C_S_S_C_.Elements, __Element__000004_credit_symbol)
	__Choice__000002_credit_S_C_S_C_.Elements = append(__Choice__000002_credit_S_C_S_C_.Elements, __Element__000006_credit_words)
	__Choice__000002_credit_S_C_S_C_.Elements = append(__Choice__000002_credit_S_C_S_C_.Elements, __Element__000007_credit_symbol)
	// setup of ComplexType instances pointers
	__ComplexType__000000_A_books.OuterElement = __Element__000000_books
	__ComplexType__000000_A_books.Sequences = append(__ComplexType__000000_A_books.Sequences, __Sequence__000000_)
	__ComplexType__000001_bookType.Annotation = __Annotation__000003_bookType_Inlined
	__ComplexType__000001_bookType.Sequences = append(__ComplexType__000001_bookType.Sequences, __Sequence__000001_bookType_S_)
	__ComplexType__000001_bookType.AttributeGroups = append(__ComplexType__000001_bookType.AttributeGroups, __AttributeGroup__000000_)
	__ComplexType__000002_credit.Annotation = __Annotation__000004_credit_Inlined
	__ComplexType__000002_credit.Sequences = append(__ComplexType__000002_credit.Sequences, __Sequence__000002_credit_S_)
	__ComplexType__000002_credit.Attributes = append(__ComplexType__000002_credit.Attributes, __Attribute__000000_page)
	__ComplexType__000003_link.Annotation = __Annotation__000005_link_Inlined
	__ComplexType__000003_link.SimpleContent = __SimpleContent__000000_link_Inlined
	// setup of Documentation instances pointers
	// setup of Element instances pointers
	__Element__000000_books.ComplexType = __ComplexType__000000_A_books
	__Element__000001_book.Annotation = __Annotation__000001_book_Inlined
	__Element__000010_title.Annotation = __Annotation__000009_title_Inlined
	__Element__000013_format.Annotation = __Annotation__000010_format_Inlined
	// setup of Enumeration instances pointers
	// setup of Extension instances pointers
	__Extension__000000_link_Inlined_Inlined.Attributes = append(__Extension__000000_link_Inlined_Inlined.Attributes, __Attribute__000001_name)
	// setup of Group instances pointers
	__Group__000001_bookDetailsGroup.Sequences = append(__Group__000001_bookDetailsGroup.Sequences, __Sequence__000005_bookDetailsGroup_S_)
	// setup of MaxInclusive instances pointers
	// setup of MinInclusive instances pointers
	// setup of Pattern instances pointers
	// setup of Restriction instances pointers
	__Restriction__000000_yearType_Inlined.MinInclusive = __MinInclusive__000000_yearType_Inlined_Inlined
	__Restriction__000000_yearType_Inlined.MaxInclusive = __MaxInclusive__000000_yearType_Inlined_Inlined
	__Restriction__000001_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000001_bookFormatEnum_Inlined.Enumerations, __Enumeration__000000_bookFormatEnum_Inlined_Inlined)
	__Restriction__000001_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000001_bookFormatEnum_Inlined.Enumerations, __Enumeration__000001_bookFormatEnum_Inlined_Inlined)
	__Restriction__000002_titleType_Inlined.Pattern = __Pattern__000000_titleType_Inlined_Inlined
	__Restriction__000002_titleType_Inlined.WhiteSpace = __WhiteSpace__000000_titleType_Inlined_Inlined
	// setup of Schema instances pointers
	__Schema__000000_Schema.Annotation = __Annotation__000000_Schema_Inlined
	__Schema__000000_Schema.Elements = append(__Schema__000000_Schema.Elements, __Element__000000_books)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000000_yearType)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000001_bookFormatEnum)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000002_titleType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000001_bookType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000002_credit)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000003_link)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000001_commonAttributes)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000002_extendedAttributes)
	__Schema__000000_Schema.Groups = append(__Schema__000000_Schema.Groups, __Group__000001_bookDetailsGroup)
	// setup of Sequence instances pointers
	__Sequence__000000_.Elements = append(__Sequence__000000_.Elements, __Element__000001_book)
	__Sequence__000001_bookType_S_.Groups = append(__Sequence__000001_bookType_S_.Groups, __Group__000000_bookType_S_G_)
	__Sequence__000001_bookType_S_.Elements = append(__Sequence__000001_bookType_S_.Elements, __Element__000002_credit)
	__Sequence__000002_credit_S_.Choices = append(__Sequence__000002_credit_S_.Choices, __Choice__000000_credit_S_C_)
	__Sequence__000002_credit_S_.Elements = append(__Sequence__000002_credit_S_.Elements, __Element__000008_credit_type)
	__Sequence__000002_credit_S_.Elements = append(__Sequence__000002_credit_S_.Elements, __Element__000009_link)
	__Sequence__000003_credit_S_C_S_.Sequences = append(__Sequence__000003_credit_S_C_S_.Sequences, __Sequence__000004_credit_S_C_S_S_)
	__Sequence__000003_credit_S_C_S_.Choices = append(__Sequence__000003_credit_S_C_S_.Choices, __Choice__000002_credit_S_C_S_C_)
	__Sequence__000004_credit_S_C_S_S_.Choices = append(__Sequence__000004_credit_S_C_S_S_.Choices, __Choice__000001_credit_S_C_S_S_C_)
	__Sequence__000004_credit_S_C_S_S_.Elements = append(__Sequence__000004_credit_S_C_S_S_.Elements, __Element__000005_link)
	__Sequence__000005_bookDetailsGroup_S_.Elements = append(__Sequence__000005_bookDetailsGroup_S_.Elements, __Element__000010_title)
	__Sequence__000005_bookDetailsGroup_S_.Elements = append(__Sequence__000005_bookDetailsGroup_S_.Elements, __Element__000011_author)
	__Sequence__000005_bookDetailsGroup_S_.Elements = append(__Sequence__000005_bookDetailsGroup_S_.Elements, __Element__000012_year)
	__Sequence__000005_bookDetailsGroup_S_.Elements = append(__Sequence__000005_bookDetailsGroup_S_.Elements, __Element__000013_format)
	// setup of SimpleContent instances pointers
	__SimpleContent__000000_link_Inlined.Extension = __Extension__000000_link_Inlined_Inlined
	// setup of SimpleType instances pointers
	__SimpleType__000000_yearType.Annotation = __Annotation__000002_yearType_Inlined
	__SimpleType__000000_yearType.Restriction = __Restriction__000000_yearType_Inlined
	__SimpleType__000001_bookFormatEnum.Restriction = __Restriction__000001_bookFormatEnum_Inlined
	__SimpleType__000002_titleType.Restriction = __Restriction__000002_titleType_Inlined
	// setup of WhiteSpace instances pointers
}

