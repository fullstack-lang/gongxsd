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

	__Annotation__000000_Schema_Inlined := (&models.Annotation{Name: `Schema_Inlined`}).Stage(stage)
	__Annotation__000001_bestseller_Inlined := (&models.Annotation{Name: `bestseller_Inlined`}).Stage(stage)
	__Annotation__000002_bookType_Inlined := (&models.Annotation{Name: `bookType_Inlined`}).Stage(stage)
	__Annotation__000003_book_Inlined := (&models.Annotation{Name: `book_Inlined`}).Stage(stage)
	__Annotation__000004_credit_Inlined := (&models.Annotation{Name: `credit_Inlined`}).Stage(stage)
	__Annotation__000005_edition_Inlined := (&models.Annotation{Name: `edition_Inlined`}).Stage(stage)
	__Annotation__000006_format_Inlined := (&models.Annotation{Name: `format_Inlined`}).Stage(stage)
	__Annotation__000007_isbn_Inlined := (&models.Annotation{Name: `isbn_Inlined`}).Stage(stage)
	__Annotation__000008_link_Inlined := (&models.Annotation{Name: `link_Inlined`}).Stage(stage)
	__Annotation__000009_title_Inlined := (&models.Annotation{Name: `title_Inlined`}).Stage(stage)
	__Annotation__000010_yearType_Inlined := (&models.Annotation{Name: `yearType_Inlined`}).Stage(stage)

	__Attribute__000000_bestseller := (&models.Attribute{Name: `bestseller`}).Stage(stage)
	__Attribute__000001_edition := (&models.Attribute{Name: `edition`}).Stage(stage)
	__Attribute__000002_isbn := (&models.Attribute{Name: `isbn`}).Stage(stage)
	__Attribute__000003_name := (&models.Attribute{Name: `name`}).Stage(stage)
	__Attribute__000004_page := (&models.Attribute{Name: `page`}).Stage(stage)

	__AttributeGroup__000000_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000001_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000002_commonAttributes := (&models.AttributeGroup{Name: `commonAttributes`}).Stage(stage)
	__AttributeGroup__000003_extendedAttributes := (&models.AttributeGroup{Name: `extendedAttributes`}).Stage(stage)

	__Choice__000000_credit_S_C_ := (&models.Choice{Name: `credit_S__C_`}).Stage(stage)
	__Choice__000001_credit_S_C_S_C_ := (&models.Choice{Name: `credit_S__C__S__C_`}).Stage(stage)
	__Choice__000002_credit_S_C_S_S_C_ := (&models.Choice{Name: `credit_S__C__S__S__C_`}).Stage(stage)

	__ComplexType__000000_A_books := (&models.ComplexType{Name: `A_books`}).Stage(stage)
	__ComplexType__000001_bookType := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__000002_credit := (&models.ComplexType{Name: `credit`}).Stage(stage)
	__ComplexType__000003_link := (&models.ComplexType{Name: `link`}).Stage(stage)

	__Documentation__000000_Schema_Inlined_Inlined := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__000001_Schema_Inlined_Inlined := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__000002__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000003__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000004__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000005__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000006__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000007__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000008_bookType_Inlined_Inlined := (&models.Documentation{Name: `bookType_Inlined_Inlined`}).Stage(stage)
	__Documentation__000009_credit_Inlined_Inlined := (&models.Documentation{Name: `credit_Inlined_Inlined`}).Stage(stage)
	__Documentation__000010_link_Inlined_Inlined := (&models.Documentation{Name: `link_Inlined_Inlined`}).Stage(stage)
	__Documentation__000011_yearType_Inlined_Inlined := (&models.Documentation{Name: `yearType_Inlined_Inlined`}).Stage(stage)

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

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)

	__Extension__000000_link_Inlined_Inlined := (&models.Extension{Name: `link_Inlined_Inlined`}).Stage(stage)

	__Group__000000_bookDetailsGroup := (&models.Group{Name: `bookDetailsGroup`}).Stage(stage)
	__Group__000001_bookType_S_G_ := (&models.Group{Name: `bookType_S__G_`}).Stage(stage)

	__MaxInclusive__000000_yearType_Inlined_Inlined := (&models.MaxInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__MinInclusive__000000_yearType_Inlined_Inlined := (&models.MinInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__Pattern__000000_titleType_Inlined_Inlined := (&models.Pattern{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	__Restriction__000000_bookFormatEnum_Inlined := (&models.Restriction{Name: `bookFormatEnum_Inlined`}).Stage(stage)
	__Restriction__000001_titleType_Inlined := (&models.Restriction{Name: `titleType_Inlined`}).Stage(stage)
	__Restriction__000002_yearType_Inlined := (&models.Restriction{Name: `yearType_Inlined`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__000000_ := (&models.Sequence{Name: ``}).Stage(stage)
	__Sequence__000001_bookDetailsGroup_S_ := (&models.Sequence{Name: `bookDetailsGroup_S_`}).Stage(stage)
	__Sequence__000002_bookType_S_ := (&models.Sequence{Name: `bookType_S_`}).Stage(stage)
	__Sequence__000003_credit_S_ := (&models.Sequence{Name: `credit_S_`}).Stage(stage)
	__Sequence__000004_credit_S_C_S_ := (&models.Sequence{Name: `credit_S__C__S_`}).Stage(stage)
	__Sequence__000005_credit_S_C_S_S_ := (&models.Sequence{Name: `credit_S__C__S__S_`}).Stage(stage)

	__SimpleContent__000000_link_Inlined := (&models.SimpleContent{Name: `link_Inlined`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_titleType := (&models.SimpleType{Name: `titleType`}).Stage(stage)
	__SimpleType__000002_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	__WhiteSpace__000000_titleType_Inlined_Inlined := (&models.WhiteSpace{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	// Setup of values

	__Annotation__000000_Schema_Inlined.Name = `Schema_Inlined`

	__Annotation__000001_bestseller_Inlined.Name = `bestseller_Inlined`

	__Annotation__000002_bookType_Inlined.Name = `bookType_Inlined`

	__Annotation__000003_book_Inlined.Name = `book_Inlined`

	__Annotation__000004_credit_Inlined.Name = `credit_Inlined`

	__Annotation__000005_edition_Inlined.Name = `edition_Inlined`

	__Annotation__000006_format_Inlined.Name = `format_Inlined`

	__Annotation__000007_isbn_Inlined.Name = `isbn_Inlined`

	__Annotation__000008_link_Inlined.Name = `link_Inlined`

	__Annotation__000009_title_Inlined.Name = `title_Inlined`

	__Annotation__000010_yearType_Inlined.Name = `yearType_Inlined`

	__Attribute__000000_bestseller.Name = `bestseller`
	__Attribute__000000_bestseller.NameXSD = `bestseller`
	__Attribute__000000_bestseller.Type = `boolean`
	__Attribute__000000_bestseller.HasNameConflict = false
	__Attribute__000000_bestseller.GoIdentifier = `Bestseller`
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
	__Attribute__000001_edition.Type = `string`
	__Attribute__000001_edition.HasNameConflict = false
	__Attribute__000001_edition.GoIdentifier = `Edition`
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

	__Attribute__000003_name.Name = `name`
	__Attribute__000003_name.NameXSD = `name`
	__Attribute__000003_name.Type = `token`
	__Attribute__000003_name.HasNameConflict = false
	__Attribute__000003_name.GoIdentifier = `NameXSD`
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
	__Attribute__000004_page.Type = `positiveInteger`
	__Attribute__000004_page.HasNameConflict = false
	__Attribute__000004_page.GoIdentifier = `Page`
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
	__AttributeGroup__000000_.HasNameConflict = true
	__AttributeGroup__000000_.GoIdentifier = `AttributeGroup__1`
	__AttributeGroup__000000_.Ref = `extendedAttributes`

	__AttributeGroup__000001_.Name = ``
	__AttributeGroup__000001_.NameXSD = ``
	__AttributeGroup__000001_.HasNameConflict = false
	__AttributeGroup__000001_.GoIdentifier = `AttributeGroup_`
	__AttributeGroup__000001_.Ref = `commonAttributes`

	__AttributeGroup__000002_commonAttributes.Name = `commonAttributes`
	__AttributeGroup__000002_commonAttributes.NameXSD = `commonAttributes`
	__AttributeGroup__000002_commonAttributes.HasNameConflict = false
	__AttributeGroup__000002_commonAttributes.GoIdentifier = `AttributeGroup_commonAttributes`
	__AttributeGroup__000002_commonAttributes.Ref = ``

	__AttributeGroup__000003_extendedAttributes.Name = `extendedAttributes`
	__AttributeGroup__000003_extendedAttributes.NameXSD = `extendedAttributes`
	__AttributeGroup__000003_extendedAttributes.HasNameConflict = false
	__AttributeGroup__000003_extendedAttributes.GoIdentifier = `AttributeGroup_extendedAttributes`
	__AttributeGroup__000003_extendedAttributes.Ref = ``

	__Choice__000000_credit_S_C_.Name = `credit_S__C_`
	__Choice__000000_credit_S_C_.MinOccurs = ``
	__Choice__000000_credit_S_C_.MaxOccurs = ``
	__Choice__000000_credit_S_C_.OuterElementName = `credit_S__C_`
	__Choice__000000_credit_S_C_.Order = 0
	__Choice__000000_credit_S_C_.Depth = 0
	__Choice__000000_credit_S_C_.IsDuplicatedInXSD = false

	__Choice__000001_credit_S_C_S_C_.Name = `credit_S__C__S__C_`
	__Choice__000001_credit_S_C_S_C_.MinOccurs = ``
	__Choice__000001_credit_S_C_S_C_.MaxOccurs = ``
	__Choice__000001_credit_S_C_S_C_.OuterElementName = `credit_S__C__S__C_`
	__Choice__000001_credit_S_C_S_C_.Order = 0
	__Choice__000001_credit_S_C_S_C_.Depth = 0
	__Choice__000001_credit_S_C_S_C_.IsDuplicatedInXSD = false

	__Choice__000002_credit_S_C_S_S_C_.Name = `credit_S__C__S__S__C_`
	__Choice__000002_credit_S_C_S_S_C_.MinOccurs = ``
	__Choice__000002_credit_S_C_S_S_C_.MaxOccurs = ``
	__Choice__000002_credit_S_C_S_S_C_.OuterElementName = `credit_S__C__S__S__C_`
	__Choice__000002_credit_S_C_S_S_C_.Order = 0
	__Choice__000002_credit_S_C_S_S_C_.Depth = 0
	__Choice__000002_credit_S_C_S_S_C_.IsDuplicatedInXSD = false

	__ComplexType__000000_A_books.Name = `A_books`
	__ComplexType__000000_A_books.HasNameConflict = false
	__ComplexType__000000_A_books.GoIdentifier = `A_books`
	__ComplexType__000000_A_books.IsAnonymous = true
	__ComplexType__000000_A_books.NameXSD = ``
	__ComplexType__000000_A_books.OuterElementName = `A_books`
	__ComplexType__000000_A_books.Order = 0
	__ComplexType__000000_A_books.Depth = 0
	__ComplexType__000000_A_books.IsDuplicatedInXSD = false

	__ComplexType__000001_bookType.Name = `bookType`
	__ComplexType__000001_bookType.HasNameConflict = false
	__ComplexType__000001_bookType.GoIdentifier = `BookType`
	__ComplexType__000001_bookType.IsAnonymous = false
	__ComplexType__000001_bookType.NameXSD = `bookType`
	__ComplexType__000001_bookType.OuterElementName = `bookType`
	__ComplexType__000001_bookType.Order = 0
	__ComplexType__000001_bookType.Depth = 0
	__ComplexType__000001_bookType.IsDuplicatedInXSD = false

	__ComplexType__000002_credit.Name = `credit`
	__ComplexType__000002_credit.HasNameConflict = false
	__ComplexType__000002_credit.GoIdentifier = `Credit`
	__ComplexType__000002_credit.IsAnonymous = false
	__ComplexType__000002_credit.NameXSD = `credit`
	__ComplexType__000002_credit.OuterElementName = `credit`
	__ComplexType__000002_credit.Order = 0
	__ComplexType__000002_credit.Depth = 0
	__ComplexType__000002_credit.IsDuplicatedInXSD = false

	__ComplexType__000003_link.Name = `link`
	__ComplexType__000003_link.HasNameConflict = false
	__ComplexType__000003_link.GoIdentifier = `Link`
	__ComplexType__000003_link.IsAnonymous = false
	__ComplexType__000003_link.NameXSD = `link`
	__ComplexType__000003_link.OuterElementName = `link`
	__ComplexType__000003_link.Order = 0
	__ComplexType__000003_link.Depth = 0
	__ComplexType__000003_link.IsDuplicatedInXSD = false

	__Documentation__000000_Schema_Inlined_Inlined.Name = `Schema_Inlined_Inlined`
	__Documentation__000000_Schema_Inlined_Inlined.Text = ` Ce schéma définit
            la structure d'une collection de livres simple. Il inclut des types pour les détails du
            livre, tels que le titre, l'auteur, l'année et le format. `
	__Documentation__000000_Schema_Inlined_Inlined.Source = `http://example.com/schema-docs`
	__Documentation__000000_Schema_Inlined_Inlined.Lang = `fr`

	__Documentation__000001_Schema_Inlined_Inlined.Name = `Schema_Inlined_Inlined`
	__Documentation__000001_Schema_Inlined_Inlined.Text = ` This schema defines
            the structure of a simple book collection. It includes types for book details, such as
            title, author, year, and format. `
	__Documentation__000001_Schema_Inlined_Inlined.Source = `http://example.com/schema-docs`
	__Documentation__000001_Schema_Inlined_Inlined.Lang = `en`

	__Documentation__000002__Inlined.Name = `_Inlined`
	__Documentation__000002__Inlined.Text = `Indicates if the book is a bestseller.`
	__Documentation__000002__Inlined.Source = ``
	__Documentation__000002__Inlined.Lang = ``

	__Documentation__000003__Inlined.Name = `_Inlined`
	__Documentation__000003__Inlined.Text = `The ISBN number of the book.`
	__Documentation__000003__Inlined.Source = ``
	__Documentation__000003__Inlined.Lang = ``

	__Documentation__000004__Inlined.Name = `_Inlined`
	__Documentation__000004__Inlined.Text = `The edition of the book.`
	__Documentation__000004__Inlined.Source = ``
	__Documentation__000004__Inlined.Lang = ``

	__Documentation__000005__Inlined.Name = `_Inlined`
	__Documentation__000005__Inlined.Text = ` A book element representing a single book in the
                            collection. `
	__Documentation__000005__Inlined.Source = ``
	__Documentation__000005__Inlined.Lang = ``

	__Documentation__000006__Inlined.Name = `_Inlined`
	__Documentation__000006__Inlined.Text = `The title of the book, consisting of alphabetic characters and
                        spaces only.`
	__Documentation__000006__Inlined.Source = ``
	__Documentation__000006__Inlined.Lang = ``

	__Documentation__000007__Inlined.Name = `_Inlined`
	__Documentation__000007__Inlined.Text = `The format of the book, either Paperback or Hardcover.`
	__Documentation__000007__Inlined.Source = ``
	__Documentation__000007__Inlined.Lang = ``

	__Documentation__000008_bookType_Inlined_Inlined.Name = `bookType_Inlined_Inlined`
	__Documentation__000008_bookType_Inlined_Inlined.Text = ` This complex type defines the structure of a book, including title,
                author, year, and format. `
	__Documentation__000008_bookType_Inlined_Inlined.Source = ``
	__Documentation__000008_bookType_Inlined_Inlined.Lang = ``

	__Documentation__000009_credit_Inlined_Inlined.Name = `credit_Inlined_Inlined`
	__Documentation__000009_credit_Inlined_Inlined.Text = `The credit type .. `
	__Documentation__000009_credit_Inlined_Inlined.Source = ``
	__Documentation__000009_credit_Inlined_Inlined.Lang = ``

	__Documentation__000010_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Documentation__000010_link_Inlined_Inlined.Text = `The link type serves as an outgoing simple XLink. If a relative link
                is used within a document that is part of a compressed MusicXML file, the link is
                relative to the root folder of the zip file.`
	__Documentation__000010_link_Inlined_Inlined.Source = ``
	__Documentation__000010_link_Inlined_Inlined.Lang = ``

	__Documentation__000011_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__Documentation__000011_yearType_Inlined_Inlined.Text = ` This type represents a year. It restricts the value to an integer
                between 1000 and 2100 inclusive. `
	__Documentation__000011_yearType_Inlined_Inlined.Source = ``
	__Documentation__000011_yearType_Inlined_Inlined.Lang = ``

	__Element__000000_author.Name = `author`
	__Element__000000_author.Order = 4
	__Element__000000_author.Depth = 1
	__Element__000000_author.HasNameConflict = false
	__Element__000000_author.GoIdentifier = `Author`
	__Element__000000_author.NameXSD = `author`
	__Element__000000_author.Type = `string`
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
	__Element__000000_author.IsDuplicatedInXSD = false

	__Element__000001_book.Name = `book`
	__Element__000001_book.Order = 8
	__Element__000001_book.Depth = 1
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

	__Element__000002_books.Name = `books`
	__Element__000002_books.Order = 7
	__Element__000002_books.Depth = 0
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
	__Element__000002_books.IsDuplicatedInXSD = false

	__Element__000003_credit.Name = `credit`
	__Element__000003_credit.Order = 1
	__Element__000003_credit.Depth = 0
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
	__Element__000003_credit.IsDuplicatedInXSD = false

	__Element__000004_credit_symbol.Name = `credit-symbol`
	__Element__000004_credit_symbol.Order = 15
	__Element__000004_credit_symbol.Depth = 0
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

	__Element__000005_credit_symbol.Name = `credit-symbol`
	__Element__000005_credit_symbol.Order = 12
	__Element__000005_credit_symbol.Depth = 0
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
	__Element__000005_credit_symbol.IsDuplicatedInXSD = false

	__Element__000006_credit_type.Name = `credit-type`
	__Element__000006_credit_type.Order = 9
	__Element__000006_credit_type.Depth = 0
	__Element__000006_credit_type.HasNameConflict = false
	__Element__000006_credit_type.GoIdentifier = `Credit_type`
	__Element__000006_credit_type.NameXSD = `credit-type`
	__Element__000006_credit_type.Type = `string`
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
	__Element__000006_credit_type.IsDuplicatedInXSD = false

	__Element__000007_credit_words.Name = `credit-words`
	__Element__000007_credit_words.Order = 11
	__Element__000007_credit_words.Depth = 0
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
	__Element__000007_credit_words.IsDuplicatedInXSD = false

	__Element__000008_credit_words.Name = `credit-words`
	__Element__000008_credit_words.Order = 14
	__Element__000008_credit_words.Depth = 0
	__Element__000008_credit_words.HasNameConflict = false
	__Element__000008_credit_words.GoIdentifier = `Credit_words`
	__Element__000008_credit_words.NameXSD = `credit-words`
	__Element__000008_credit_words.Type = `string`
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
	__Element__000008_credit_words.IsDuplicatedInXSD = false

	__Element__000009_format.Name = `format`
	__Element__000009_format.Order = 6
	__Element__000009_format.Depth = 1
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
	__Element__000009_format.IsDuplicatedInXSD = false

	__Element__000010_link.Name = `link`
	__Element__000010_link.Order = 13
	__Element__000010_link.Depth = 0
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
	__Element__000010_link.IsDuplicatedInXSD = false

	__Element__000011_link.Name = `link`
	__Element__000011_link.Order = 10
	__Element__000011_link.Depth = 0
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
	__Element__000011_link.IsDuplicatedInXSD = false

	__Element__000012_title.Name = `title`
	__Element__000012_title.Order = 3
	__Element__000012_title.Depth = 1
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
	__Element__000012_title.IsDuplicatedInXSD = false

	__Element__000013_year.Name = `year`
	__Element__000013_year.Order = 5
	__Element__000013_year.Depth = 1
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
	__Element__000013_year.IsDuplicatedInXSD = false

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Value = `Hardcover`

	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Value = `Paperback`

	__Extension__000000_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Extension__000000_link_Inlined_Inlined.OuterElementName = ``
	__Extension__000000_link_Inlined_Inlined.Base = `string`
	__Extension__000000_link_Inlined_Inlined.Ref = ``

	__Group__000000_bookDetailsGroup.Name = `bookDetailsGroup`
	__Group__000000_bookDetailsGroup.NameXSD = `bookDetailsGroup`
	__Group__000000_bookDetailsGroup.Ref = ``
	__Group__000000_bookDetailsGroup.IsAnonymous = false
	__Group__000000_bookDetailsGroup.HasNameConflict = false
	__Group__000000_bookDetailsGroup.GoIdentifier = `Group_bookDetailsGroup`
	__Group__000000_bookDetailsGroup.OuterElementName = `bookDetailsGroup`
	__Group__000000_bookDetailsGroup.Order = 2
	__Group__000000_bookDetailsGroup.Depth = 0

	__Group__000001_bookType_S_G_.Name = `bookType_S__G_`
	__Group__000001_bookType_S_G_.NameXSD = ``
	__Group__000001_bookType_S_G_.Ref = `bookDetailsGroup`
	__Group__000001_bookType_S_G_.IsAnonymous = false
	__Group__000001_bookType_S_G_.HasNameConflict = false
	__Group__000001_bookType_S_G_.GoIdentifier = `BookDetailsGroup`
	__Group__000001_bookType_S_G_.OuterElementName = `bookType_S__G_`
	__Group__000001_bookType_S_G_.Order = 0
	__Group__000001_bookType_S_G_.Depth = 0

	__MaxInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MaxInclusive__000000_yearType_Inlined_Inlined.Value = `2100`

	__MinInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MinInclusive__000000_yearType_Inlined_Inlined.Value = `1000`

	__Pattern__000000_titleType_Inlined_Inlined.Name = `titleType_Inlined_Inlined`
	__Pattern__000000_titleType_Inlined_Inlined.Value = `[0-9A-Za-z ]+`

	__Restriction__000000_bookFormatEnum_Inlined.Name = `bookFormatEnum_Inlined`
	__Restriction__000000_bookFormatEnum_Inlined.Base = `string`

	__Restriction__000001_titleType_Inlined.Name = `titleType_Inlined`
	__Restriction__000001_titleType_Inlined.Base = `string`

	__Restriction__000002_yearType_Inlined.Name = `yearType_Inlined`
	__Restriction__000002_yearType_Inlined.Base = `integer`

	__Schema__000000_Schema.Name = `Schema`
	__Schema__000000_Schema.Xs = `http://www.w3.org/2001/XMLSchema`

	__Sequence__000000_.Name = ``
	__Sequence__000000_.MinOccurs = ``
	__Sequence__000000_.MaxOccurs = ``
	__Sequence__000000_.OuterElementName = ``
	__Sequence__000000_.Order = 0
	__Sequence__000000_.Depth = 0

	__Sequence__000001_bookDetailsGroup_S_.Name = `bookDetailsGroup_S_`
	__Sequence__000001_bookDetailsGroup_S_.MinOccurs = ``
	__Sequence__000001_bookDetailsGroup_S_.MaxOccurs = ``
	__Sequence__000001_bookDetailsGroup_S_.OuterElementName = `bookDetailsGroup_S_`
	__Sequence__000001_bookDetailsGroup_S_.Order = 0
	__Sequence__000001_bookDetailsGroup_S_.Depth = 0

	__Sequence__000002_bookType_S_.Name = `bookType_S_`
	__Sequence__000002_bookType_S_.MinOccurs = ``
	__Sequence__000002_bookType_S_.MaxOccurs = ``
	__Sequence__000002_bookType_S_.OuterElementName = `bookType_S_`
	__Sequence__000002_bookType_S_.Order = 0
	__Sequence__000002_bookType_S_.Depth = 0

	__Sequence__000003_credit_S_.Name = `credit_S_`
	__Sequence__000003_credit_S_.MinOccurs = ``
	__Sequence__000003_credit_S_.MaxOccurs = ``
	__Sequence__000003_credit_S_.OuterElementName = `credit_S_`
	__Sequence__000003_credit_S_.Order = 0
	__Sequence__000003_credit_S_.Depth = 0

	__Sequence__000004_credit_S_C_S_.Name = `credit_S__C__S_`
	__Sequence__000004_credit_S_C_S_.MinOccurs = ``
	__Sequence__000004_credit_S_C_S_.MaxOccurs = ``
	__Sequence__000004_credit_S_C_S_.OuterElementName = `credit_S__C__S_`
	__Sequence__000004_credit_S_C_S_.Order = 0
	__Sequence__000004_credit_S_C_S_.Depth = 0

	__Sequence__000005_credit_S_C_S_S_.Name = `credit_S__C__S__S_`
	__Sequence__000005_credit_S_C_S_S_.MinOccurs = `0`
	__Sequence__000005_credit_S_C_S_S_.MaxOccurs = `unbounded`
	__Sequence__000005_credit_S_C_S_S_.OuterElementName = `credit_S__C__S__S_`
	__Sequence__000005_credit_S_C_S_S_.Order = 0
	__Sequence__000005_credit_S_C_S_S_.Depth = 0

	__SimpleContent__000000_link_Inlined.Name = `link_Inlined`

	__SimpleType__000000_bookFormatEnum.Name = `bookFormatEnum`
	__SimpleType__000000_bookFormatEnum.NameXSD = `bookFormatEnum`

	__SimpleType__000001_titleType.Name = `titleType`
	__SimpleType__000001_titleType.NameXSD = `titleType`

	__SimpleType__000002_yearType.Name = `yearType`
	__SimpleType__000002_yearType.NameXSD = `yearType`

	__WhiteSpace__000000_titleType_Inlined_Inlined.Name = `titleType_Inlined_Inlined`
	__WhiteSpace__000000_titleType_Inlined_Inlined.Value = `collapse`

	// Setup of pointers
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000001_Schema_Inlined_Inlined)
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000000_Schema_Inlined_Inlined)
	__Annotation__000001_bestseller_Inlined.Documentations = append(__Annotation__000001_bestseller_Inlined.Documentations, __Documentation__000002__Inlined)
	__Annotation__000002_bookType_Inlined.Documentations = append(__Annotation__000002_bookType_Inlined.Documentations, __Documentation__000008_bookType_Inlined_Inlined)
	__Annotation__000003_book_Inlined.Documentations = append(__Annotation__000003_book_Inlined.Documentations, __Documentation__000005__Inlined)
	__Annotation__000004_credit_Inlined.Documentations = append(__Annotation__000004_credit_Inlined.Documentations, __Documentation__000009_credit_Inlined_Inlined)
	__Annotation__000005_edition_Inlined.Documentations = append(__Annotation__000005_edition_Inlined.Documentations, __Documentation__000004__Inlined)
	__Annotation__000006_format_Inlined.Documentations = append(__Annotation__000006_format_Inlined.Documentations, __Documentation__000007__Inlined)
	__Annotation__000007_isbn_Inlined.Documentations = append(__Annotation__000007_isbn_Inlined.Documentations, __Documentation__000003__Inlined)
	__Annotation__000008_link_Inlined.Documentations = append(__Annotation__000008_link_Inlined.Documentations, __Documentation__000010_link_Inlined_Inlined)
	__Annotation__000009_title_Inlined.Documentations = append(__Annotation__000009_title_Inlined.Documentations, __Documentation__000006__Inlined)
	__Annotation__000010_yearType_Inlined.Documentations = append(__Annotation__000010_yearType_Inlined.Documentations, __Documentation__000011_yearType_Inlined_Inlined)
	__Attribute__000000_bestseller.Annotation = __Annotation__000001_bestseller_Inlined
	__Attribute__000001_edition.Annotation = __Annotation__000005_edition_Inlined
	__Attribute__000002_isbn.Annotation = __Annotation__000007_isbn_Inlined
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000002_isbn)
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000000_bestseller)
	__AttributeGroup__000003_extendedAttributes.AttributeGroups = append(__AttributeGroup__000003_extendedAttributes.AttributeGroups, __AttributeGroup__000001_)
	__AttributeGroup__000003_extendedAttributes.Attributes = append(__AttributeGroup__000003_extendedAttributes.Attributes, __Attribute__000001_edition)
	__Choice__000000_credit_S_C_.Sequences = append(__Choice__000000_credit_S_C_.Sequences, __Sequence__000004_credit_S_C_S_)
	__Choice__000001_credit_S_C_S_C_.Elements = append(__Choice__000001_credit_S_C_S_C_.Elements, __Element__000007_credit_words)
	__Choice__000001_credit_S_C_S_C_.Elements = append(__Choice__000001_credit_S_C_S_C_.Elements, __Element__000005_credit_symbol)
	__Choice__000002_credit_S_C_S_S_C_.Elements = append(__Choice__000002_credit_S_C_S_S_C_.Elements, __Element__000008_credit_words)
	__Choice__000002_credit_S_C_S_S_C_.Elements = append(__Choice__000002_credit_S_C_S_S_C_.Elements, __Element__000004_credit_symbol)
	__ComplexType__000000_A_books.OuterElement = __Element__000002_books
	__ComplexType__000000_A_books.Sequences = append(__ComplexType__000000_A_books.Sequences, __Sequence__000000_)
	__ComplexType__000001_bookType.Annotation = __Annotation__000002_bookType_Inlined
	__ComplexType__000001_bookType.Sequences = append(__ComplexType__000001_bookType.Sequences, __Sequence__000002_bookType_S_)
	__ComplexType__000001_bookType.AttributeGroups = append(__ComplexType__000001_bookType.AttributeGroups, __AttributeGroup__000000_)
	__ComplexType__000002_credit.Annotation = __Annotation__000004_credit_Inlined
	__ComplexType__000002_credit.Sequences = append(__ComplexType__000002_credit.Sequences, __Sequence__000003_credit_S_)
	__ComplexType__000002_credit.Attributes = append(__ComplexType__000002_credit.Attributes, __Attribute__000004_page)
	__ComplexType__000003_link.Annotation = __Annotation__000008_link_Inlined
	__ComplexType__000003_link.SimpleContent = __SimpleContent__000000_link_Inlined
	__Element__000001_book.Annotation = __Annotation__000003_book_Inlined
	__Element__000002_books.ComplexType = __ComplexType__000000_A_books
	__Element__000009_format.Annotation = __Annotation__000006_format_Inlined
	__Element__000012_title.Annotation = __Annotation__000009_title_Inlined
	__Extension__000000_link_Inlined_Inlined.Attributes = append(__Extension__000000_link_Inlined_Inlined.Attributes, __Attribute__000003_name)
	__Group__000000_bookDetailsGroup.Sequences = append(__Group__000000_bookDetailsGroup.Sequences, __Sequence__000001_bookDetailsGroup_S_)
	__Restriction__000000_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000000_bookFormatEnum_Inlined.Enumerations, __Enumeration__000001_bookFormatEnum_Inlined_Inlined)
	__Restriction__000000_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000000_bookFormatEnum_Inlined.Enumerations, __Enumeration__000000_bookFormatEnum_Inlined_Inlined)
	__Restriction__000001_titleType_Inlined.Pattern = __Pattern__000000_titleType_Inlined_Inlined
	__Restriction__000001_titleType_Inlined.WhiteSpace = __WhiteSpace__000000_titleType_Inlined_Inlined
	__Restriction__000002_yearType_Inlined.MinInclusive = __MinInclusive__000000_yearType_Inlined_Inlined
	__Restriction__000002_yearType_Inlined.MaxInclusive = __MaxInclusive__000000_yearType_Inlined_Inlined
	__Schema__000000_Schema.Annotation = __Annotation__000000_Schema_Inlined
	__Schema__000000_Schema.Elements = append(__Schema__000000_Schema.Elements, __Element__000002_books)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000002_yearType)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000000_bookFormatEnum)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000001_titleType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000001_bookType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000002_credit)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000003_link)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000002_commonAttributes)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000003_extendedAttributes)
	__Schema__000000_Schema.Groups = append(__Schema__000000_Schema.Groups, __Group__000000_bookDetailsGroup)
	__Sequence__000000_.Elements = append(__Sequence__000000_.Elements, __Element__000001_book)
	__Sequence__000001_bookDetailsGroup_S_.Elements = append(__Sequence__000001_bookDetailsGroup_S_.Elements, __Element__000012_title)
	__Sequence__000001_bookDetailsGroup_S_.Elements = append(__Sequence__000001_bookDetailsGroup_S_.Elements, __Element__000000_author)
	__Sequence__000001_bookDetailsGroup_S_.Elements = append(__Sequence__000001_bookDetailsGroup_S_.Elements, __Element__000013_year)
	__Sequence__000001_bookDetailsGroup_S_.Elements = append(__Sequence__000001_bookDetailsGroup_S_.Elements, __Element__000009_format)
	__Sequence__000002_bookType_S_.Groups = append(__Sequence__000002_bookType_S_.Groups, __Group__000001_bookType_S_G_)
	__Sequence__000002_bookType_S_.Elements = append(__Sequence__000002_bookType_S_.Elements, __Element__000003_credit)
	__Sequence__000003_credit_S_.Choices = append(__Sequence__000003_credit_S_.Choices, __Choice__000000_credit_S_C_)
	__Sequence__000003_credit_S_.Elements = append(__Sequence__000003_credit_S_.Elements, __Element__000006_credit_type)
	__Sequence__000003_credit_S_.Elements = append(__Sequence__000003_credit_S_.Elements, __Element__000011_link)
	__Sequence__000004_credit_S_C_S_.Sequences = append(__Sequence__000004_credit_S_C_S_.Sequences, __Sequence__000005_credit_S_C_S_S_)
	__Sequence__000004_credit_S_C_S_.Choices = append(__Sequence__000004_credit_S_C_S_.Choices, __Choice__000001_credit_S_C_S_C_)
	__Sequence__000005_credit_S_C_S_S_.Choices = append(__Sequence__000005_credit_S_C_S_S_.Choices, __Choice__000002_credit_S_C_S_S_C_)
	__Sequence__000005_credit_S_C_S_S_.Elements = append(__Sequence__000005_credit_S_C_S_S_.Elements, __Element__000010_link)
	__SimpleContent__000000_link_Inlined.Extension = __Extension__000000_link_Inlined_Inlined
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_bookFormatEnum_Inlined
	__SimpleType__000001_titleType.Restriction = __Restriction__000001_titleType_Inlined
	__SimpleType__000002_yearType.Annotation = __Annotation__000010_yearType_Inlined
	__SimpleType__000002_yearType.Restriction = __Restriction__000002_yearType_Inlined
}
