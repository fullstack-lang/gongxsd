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

	// insertion point for declaration of instances to stage

	__Annotation__00000000_ := (&models.Annotation{Name: `Schema_Inlined`}).Stage(stage)
	__Annotation__00000001_ := (&models.Annotation{Name: `book_Inlined`}).Stage(stage)
	__Annotation__00000002_ := (&models.Annotation{Name: `yearType_Inlined`}).Stage(stage)
	__Annotation__00000003_ := (&models.Annotation{Name: `bookType_Inlined`}).Stage(stage)
	__Annotation__00000004_ := (&models.Annotation{Name: `credit_Inlined`}).Stage(stage)
	__Annotation__00000005_ := (&models.Annotation{Name: `link_Inlined`}).Stage(stage)
	__Annotation__00000006_ := (&models.Annotation{Name: `isbn_Inlined`}).Stage(stage)
	__Annotation__00000007_ := (&models.Annotation{Name: `bestseller_Inlined`}).Stage(stage)
	__Annotation__00000008_ := (&models.Annotation{Name: `edition_Inlined`}).Stage(stage)
	__Annotation__00000009_ := (&models.Annotation{Name: `title_Inlined`}).Stage(stage)
	__Annotation__00000010_ := (&models.Annotation{Name: `format_Inlined`}).Stage(stage)

	__Attribute__00000000_ := (&models.Attribute{Name: `page`}).Stage(stage)
	__Attribute__00000001_ := (&models.Attribute{Name: `name`}).Stage(stage)
	__Attribute__00000002_ := (&models.Attribute{Name: `isbn`}).Stage(stage)
	__Attribute__00000003_ := (&models.Attribute{Name: `bestseller`}).Stage(stage)
	__Attribute__00000004_ := (&models.Attribute{Name: `edition`}).Stage(stage)

	__AttributeGroup__00000000_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__00000001_ := (&models.AttributeGroup{Name: `commonAttributes`}).Stage(stage)
	__AttributeGroup__00000002_ := (&models.AttributeGroup{Name: `extendedAttributes`}).Stage(stage)
	__AttributeGroup__00000003_ := (&models.AttributeGroup{Name: ``}).Stage(stage)

	__Choice__00000000_ := (&models.Choice{Name: `credit_S__C_`}).Stage(stage)
	__Choice__00000001_ := (&models.Choice{Name: `credit_S__C__S__S__C_`}).Stage(stage)
	__Choice__00000002_ := (&models.Choice{Name: `credit_S__C__S__C_`}).Stage(stage)

	__ComplexType__00000000_ := (&models.ComplexType{Name: `A_books`}).Stage(stage)
	__ComplexType__00000001_ := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__00000002_ := (&models.ComplexType{Name: `credit`}).Stage(stage)
	__ComplexType__00000003_ := (&models.ComplexType{Name: `link`}).Stage(stage)

	__Documentation__00000000_ := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000001_ := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000002_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__00000003_ := (&models.Documentation{Name: `yearType_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000004_ := (&models.Documentation{Name: `bookType_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000005_ := (&models.Documentation{Name: `credit_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000006_ := (&models.Documentation{Name: `link_Inlined_Inlined`}).Stage(stage)
	__Documentation__00000007_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__00000008_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__00000009_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__00000010_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__00000011_ := (&models.Documentation{Name: `_Inlined`}).Stage(stage)

	__Element__00000000_ := (&models.Element{Name: `books`}).Stage(stage)
	__Element__00000001_ := (&models.Element{Name: `book`}).Stage(stage)
	__Element__00000002_ := (&models.Element{Name: `credit`}).Stage(stage)
	__Element__00000003_ := (&models.Element{Name: `credit-words`}).Stage(stage)
	__Element__00000004_ := (&models.Element{Name: `credit-symbol`}).Stage(stage)
	__Element__00000005_ := (&models.Element{Name: `link`}).Stage(stage)
	__Element__00000006_ := (&models.Element{Name: `credit-words`}).Stage(stage)
	__Element__00000007_ := (&models.Element{Name: `credit-symbol`}).Stage(stage)
	__Element__00000008_ := (&models.Element{Name: `credit-type`}).Stage(stage)
	__Element__00000009_ := (&models.Element{Name: `link`}).Stage(stage)
	__Element__00000010_ := (&models.Element{Name: `title`}).Stage(stage)
	__Element__00000011_ := (&models.Element{Name: `author`}).Stage(stage)
	__Element__00000012_ := (&models.Element{Name: `year`}).Stage(stage)
	__Element__00000013_ := (&models.Element{Name: `format`}).Stage(stage)

	__Enumeration__00000000_ := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)
	__Enumeration__00000001_ := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)

	__Extension__00000000_ := (&models.Extension{Name: `link_Inlined_Inlined`}).Stage(stage)

	__Group__00000000_ := (&models.Group{Name: `bookType_S__G_`}).Stage(stage)
	__Group__00000001_ := (&models.Group{Name: `bookDetailsGroup`}).Stage(stage)

	__MaxInclusive__00000000_ := (&models.MaxInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__MinInclusive__00000000_ := (&models.MinInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__Pattern__00000000_ := (&models.Pattern{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	__Restriction__00000000_ := (&models.Restriction{Name: `yearType_Inlined`}).Stage(stage)
	__Restriction__00000001_ := (&models.Restriction{Name: `bookFormatEnum_Inlined`}).Stage(stage)
	__Restriction__00000002_ := (&models.Restriction{Name: `titleType_Inlined`}).Stage(stage)

	__Schema__00000000_ := (&models.Schema{Name: `Schema`}).Stage(stage)

	__Sequence__00000000_ := (&models.Sequence{Name: ``}).Stage(stage)
	__Sequence__00000001_ := (&models.Sequence{Name: `bookType_S_`}).Stage(stage)
	__Sequence__00000002_ := (&models.Sequence{Name: `credit_S_`}).Stage(stage)
	__Sequence__00000003_ := (&models.Sequence{Name: `credit_S__C__S_`}).Stage(stage)
	__Sequence__00000004_ := (&models.Sequence{Name: `credit_S__C__S__S_`}).Stage(stage)
	__Sequence__00000005_ := (&models.Sequence{Name: `bookDetailsGroup_S_`}).Stage(stage)

	__SimpleContent__00000000_ := (&models.SimpleContent{Name: `link_Inlined`}).Stage(stage)

	__SimpleType__00000000_ := (&models.SimpleType{Name: `yearType`}).Stage(stage)
	__SimpleType__00000001_ := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__00000002_ := (&models.SimpleType{Name: `titleType`}).Stage(stage)

	__WhiteSpace__00000000_ := (&models.WhiteSpace{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	// insertion point for initialization of values

	__Annotation__00000000_.Name = `Schema_Inlined`

	__Annotation__00000001_.Name = `book_Inlined`

	__Annotation__00000002_.Name = `yearType_Inlined`

	__Annotation__00000003_.Name = `bookType_Inlined`

	__Annotation__00000004_.Name = `credit_Inlined`

	__Annotation__00000005_.Name = `link_Inlined`

	__Annotation__00000006_.Name = `isbn_Inlined`

	__Annotation__00000007_.Name = `bestseller_Inlined`

	__Annotation__00000008_.Name = `edition_Inlined`

	__Annotation__00000009_.Name = `title_Inlined`

	__Annotation__00000010_.Name = `format_Inlined`

	__Attribute__00000000_.Name = `page`
	__Attribute__00000000_.NameXSD = `page`
	__Attribute__00000000_.Type = `positiveInteger`
	__Attribute__00000000_.HasNameConflict = false
	__Attribute__00000000_.GoIdentifier = `Page`
	__Attribute__00000000_.Default = ``
	__Attribute__00000000_.Use = ``
	__Attribute__00000000_.Form = ``
	__Attribute__00000000_.Fixed = ``
	__Attribute__00000000_.Ref = ``
	__Attribute__00000000_.TargetNamespace = ``
	__Attribute__00000000_.SimpleType = ``
	__Attribute__00000000_.IDXSD = ``

	__Attribute__00000001_.Name = `name`
	__Attribute__00000001_.NameXSD = `name`
	__Attribute__00000001_.Type = `token`
	__Attribute__00000001_.HasNameConflict = false
	__Attribute__00000001_.GoIdentifier = `NameXSD`
	__Attribute__00000001_.Default = ``
	__Attribute__00000001_.Use = ``
	__Attribute__00000001_.Form = ``
	__Attribute__00000001_.Fixed = ``
	__Attribute__00000001_.Ref = ``
	__Attribute__00000001_.TargetNamespace = ``
	__Attribute__00000001_.SimpleType = ``
	__Attribute__00000001_.IDXSD = ``

	__Attribute__00000002_.Name = `isbn`
	__Attribute__00000002_.NameXSD = `isbn`
	__Attribute__00000002_.Type = `string`
	__Attribute__00000002_.HasNameConflict = false
	__Attribute__00000002_.GoIdentifier = `Isbn`
	__Attribute__00000002_.Default = ``
	__Attribute__00000002_.Use = `required`
	__Attribute__00000002_.Form = ``
	__Attribute__00000002_.Fixed = ``
	__Attribute__00000002_.Ref = ``
	__Attribute__00000002_.TargetNamespace = ``
	__Attribute__00000002_.SimpleType = ``
	__Attribute__00000002_.IDXSD = ``

	__Attribute__00000003_.Name = `bestseller`
	__Attribute__00000003_.NameXSD = `bestseller`
	__Attribute__00000003_.Type = `boolean`
	__Attribute__00000003_.HasNameConflict = false
	__Attribute__00000003_.GoIdentifier = `Bestseller`
	__Attribute__00000003_.Default = ``
	__Attribute__00000003_.Use = `optional`
	__Attribute__00000003_.Form = ``
	__Attribute__00000003_.Fixed = ``
	__Attribute__00000003_.Ref = ``
	__Attribute__00000003_.TargetNamespace = ``
	__Attribute__00000003_.SimpleType = ``
	__Attribute__00000003_.IDXSD = ``

	__Attribute__00000004_.Name = `edition`
	__Attribute__00000004_.NameXSD = `edition`
	__Attribute__00000004_.Type = `string`
	__Attribute__00000004_.HasNameConflict = false
	__Attribute__00000004_.GoIdentifier = `Edition`
	__Attribute__00000004_.Default = ``
	__Attribute__00000004_.Use = `optional`
	__Attribute__00000004_.Form = ``
	__Attribute__00000004_.Fixed = ``
	__Attribute__00000004_.Ref = ``
	__Attribute__00000004_.TargetNamespace = ``
	__Attribute__00000004_.SimpleType = ``
	__Attribute__00000004_.IDXSD = ``

	__AttributeGroup__00000000_.Name = ``
	__AttributeGroup__00000000_.NameXSD = ``
	__AttributeGroup__00000000_.HasNameConflict = true
	__AttributeGroup__00000000_.GoIdentifier = `AttributeGroup__1`
	__AttributeGroup__00000000_.Ref = `extendedAttributes`
	__AttributeGroup__00000000_.Order = 5
	__AttributeGroup__00000000_.Depth = 1

	__AttributeGroup__00000001_.Name = `commonAttributes`
	__AttributeGroup__00000001_.NameXSD = `commonAttributes`
	__AttributeGroup__00000001_.HasNameConflict = false
	__AttributeGroup__00000001_.GoIdentifier = `AttributeGroup_commonAttributes`
	__AttributeGroup__00000001_.Ref = ``
	__AttributeGroup__00000001_.Order = 16
	__AttributeGroup__00000001_.Depth = 0

	__AttributeGroup__00000002_.Name = `extendedAttributes`
	__AttributeGroup__00000002_.NameXSD = `extendedAttributes`
	__AttributeGroup__00000002_.HasNameConflict = false
	__AttributeGroup__00000002_.GoIdentifier = `AttributeGroup_extendedAttributes`
	__AttributeGroup__00000002_.Ref = ``
	__AttributeGroup__00000002_.Order = 17
	__AttributeGroup__00000002_.Depth = 0

	__AttributeGroup__00000003_.Name = ``
	__AttributeGroup__00000003_.NameXSD = ``
	__AttributeGroup__00000003_.HasNameConflict = false
	__AttributeGroup__00000003_.GoIdentifier = `AttributeGroup_`
	__AttributeGroup__00000003_.Ref = `commonAttributes`
	__AttributeGroup__00000003_.Order = 18
	__AttributeGroup__00000003_.Depth = 1

	__Choice__00000000_.Name = `credit_S__C_`
	__Choice__00000000_.OuterElementName = `credit_S__C_`
	__Choice__00000000_.Order = 0
	__Choice__00000000_.Depth = 0
	__Choice__00000000_.MinOccurs = ``
	__Choice__00000000_.MaxOccurs = ``
	__Choice__00000000_.IsDuplicatedInXSD = false

	__Choice__00000001_.Name = `credit_S__C__S__S__C_`
	__Choice__00000001_.OuterElementName = `credit_S__C__S__S__C_`
	__Choice__00000001_.Order = 0
	__Choice__00000001_.Depth = 0
	__Choice__00000001_.MinOccurs = ``
	__Choice__00000001_.MaxOccurs = ``
	__Choice__00000001_.IsDuplicatedInXSD = false

	__Choice__00000002_.Name = `credit_S__C__S__C_`
	__Choice__00000002_.OuterElementName = `credit_S__C__S__C_`
	__Choice__00000002_.Order = 0
	__Choice__00000002_.Depth = 0
	__Choice__00000002_.MinOccurs = ``
	__Choice__00000002_.MaxOccurs = ``
	__Choice__00000002_.IsDuplicatedInXSD = false

	__ComplexType__00000000_.Name = `A_books`
	__ComplexType__00000000_.HasNameConflict = false
	__ComplexType__00000000_.GoIdentifier = `A_books`
	__ComplexType__00000000_.IsAnonymous = true
	__ComplexType__00000000_.NameXSD = ``
	__ComplexType__00000000_.OuterElementName = `A_books`
	__ComplexType__00000000_.Order = 14
	__ComplexType__00000000_.Depth = 1
	__ComplexType__00000000_.MinOccurs = ``
	__ComplexType__00000000_.MaxOccurs = ``
	__ComplexType__00000000_.IsDuplicatedInXSD = false

	__ComplexType__00000001_.Name = `bookType`
	__ComplexType__00000001_.HasNameConflict = false
	__ComplexType__00000001_.GoIdentifier = `BookType`
	__ComplexType__00000001_.IsAnonymous = false
	__ComplexType__00000001_.NameXSD = `bookType`
	__ComplexType__00000001_.OuterElementName = `bookType`
	__ComplexType__00000001_.Order = 2
	__ComplexType__00000001_.Depth = 0
	__ComplexType__00000001_.MinOccurs = ``
	__ComplexType__00000001_.MaxOccurs = ``
	__ComplexType__00000001_.IsDuplicatedInXSD = false

	__ComplexType__00000002_.Name = `credit`
	__ComplexType__00000002_.HasNameConflict = false
	__ComplexType__00000002_.GoIdentifier = `Credit`
	__ComplexType__00000002_.IsAnonymous = false
	__ComplexType__00000002_.NameXSD = `credit`
	__ComplexType__00000002_.OuterElementName = `credit`
	__ComplexType__00000002_.Order = 19
	__ComplexType__00000002_.Depth = 0
	__ComplexType__00000002_.MinOccurs = ``
	__ComplexType__00000002_.MaxOccurs = ``
	__ComplexType__00000002_.IsDuplicatedInXSD = false

	__ComplexType__00000003_.Name = `link`
	__ComplexType__00000003_.HasNameConflict = false
	__ComplexType__00000003_.GoIdentifier = `Link`
	__ComplexType__00000003_.IsAnonymous = false
	__ComplexType__00000003_.NameXSD = `link`
	__ComplexType__00000003_.OuterElementName = `link`
	__ComplexType__00000003_.Order = 27
	__ComplexType__00000003_.Depth = 0
	__ComplexType__00000003_.MinOccurs = ``
	__ComplexType__00000003_.MaxOccurs = ``
	__ComplexType__00000003_.IsDuplicatedInXSD = false

	__Documentation__00000000_.Name = `Schema_Inlined_Inlined`
	__Documentation__00000000_.Text = ` This schema defines
            the structure of a simple book collection. It includes types for book details, such as
            title, author, year, and format. `
	__Documentation__00000000_.Source = `http://example.com/schema-docs`
	__Documentation__00000000_.Lang = `en`

	__Documentation__00000001_.Name = `Schema_Inlined_Inlined`
	__Documentation__00000001_.Text = ` Ce schéma définit
            la structure d'une collection de livres simple. Il inclut des types pour les détails du
            livre, tels que le titre, l'auteur, l'année et le format. `
	__Documentation__00000001_.Source = `http://example.com/schema-docs`
	__Documentation__00000001_.Lang = `fr`

	__Documentation__00000002_.Name = `_Inlined`
	__Documentation__00000002_.Text = ` A book element representing a single book in the
                            collection. `
	__Documentation__00000002_.Source = ``
	__Documentation__00000002_.Lang = ``

	__Documentation__00000003_.Name = `yearType_Inlined_Inlined`
	__Documentation__00000003_.Text = ` This type represents a year. It restricts the value to an integer
                between 1000 and 2100 inclusive. `
	__Documentation__00000003_.Source = ``
	__Documentation__00000003_.Lang = ``

	__Documentation__00000004_.Name = `bookType_Inlined_Inlined`
	__Documentation__00000004_.Text = ` This complex type defines the structure of a book, including title,
                author, year, and format. `
	__Documentation__00000004_.Source = ``
	__Documentation__00000004_.Lang = ``

	__Documentation__00000005_.Name = `credit_Inlined_Inlined`
	__Documentation__00000005_.Text = `The credit type .. `
	__Documentation__00000005_.Source = ``
	__Documentation__00000005_.Lang = ``

	__Documentation__00000006_.Name = `link_Inlined_Inlined`
	__Documentation__00000006_.Text = `The link type serves as an outgoing simple XLink. If a relative link
                is used within a document that is part of a compressed MusicXML file, the link is
                relative to the root folder of the zip file.`
	__Documentation__00000006_.Source = ``
	__Documentation__00000006_.Lang = ``

	__Documentation__00000007_.Name = `_Inlined`
	__Documentation__00000007_.Text = `The ISBN number of the book.`
	__Documentation__00000007_.Source = ``
	__Documentation__00000007_.Lang = ``

	__Documentation__00000008_.Name = `_Inlined`
	__Documentation__00000008_.Text = `Indicates if the book is a bestseller.`
	__Documentation__00000008_.Source = ``
	__Documentation__00000008_.Lang = ``

	__Documentation__00000009_.Name = `_Inlined`
	__Documentation__00000009_.Text = `The edition of the book.`
	__Documentation__00000009_.Source = ``
	__Documentation__00000009_.Lang = ``

	__Documentation__00000010_.Name = `_Inlined`
	__Documentation__00000010_.Text = `The title of the book, consisting of alphabetic characters and
                        spaces only.`
	__Documentation__00000010_.Source = ``
	__Documentation__00000010_.Lang = ``

	__Documentation__00000011_.Name = `_Inlined`
	__Documentation__00000011_.Text = `The format of the book, either Paperback or Hardcover.`
	__Documentation__00000011_.Source = ``
	__Documentation__00000011_.Lang = ``

	__Element__00000000_.Name = `books`
	__Element__00000000_.Order = 13
	__Element__00000000_.Depth = 0
	__Element__00000000_.HasNameConflict = false
	__Element__00000000_.GoIdentifier = `Books`
	__Element__00000000_.NameXSD = `books`
	__Element__00000000_.Type = ``
	__Element__00000000_.MinOccurs = ``
	__Element__00000000_.MaxOccurs = ``
	__Element__00000000_.Default = ``
	__Element__00000000_.Fixed = ``
	__Element__00000000_.Nillable = ``
	__Element__00000000_.Ref = ``
	__Element__00000000_.Abstract = ``
	__Element__00000000_.Form = ``
	__Element__00000000_.Block = ``
	__Element__00000000_.Final = ``
	__Element__00000000_.IsDuplicatedInXSD = false

	__Element__00000001_.Name = `book`
	__Element__00000001_.Order = 15
	__Element__00000001_.Depth = 2
	__Element__00000001_.HasNameConflict = false
	__Element__00000001_.GoIdentifier = `Book`
	__Element__00000001_.NameXSD = `book`
	__Element__00000001_.Type = `bookType`
	__Element__00000001_.MinOccurs = ``
	__Element__00000001_.MaxOccurs = `unbounded`
	__Element__00000001_.Default = ``
	__Element__00000001_.Fixed = ``
	__Element__00000001_.Nillable = ``
	__Element__00000001_.Ref = ``
	__Element__00000001_.Abstract = ``
	__Element__00000001_.Form = ``
	__Element__00000001_.Block = ``
	__Element__00000001_.Final = ``
	__Element__00000001_.IsDuplicatedInXSD = false

	__Element__00000002_.Name = `credit`
	__Element__00000002_.Order = 4
	__Element__00000002_.Depth = 1
	__Element__00000002_.HasNameConflict = false
	__Element__00000002_.GoIdentifier = `Credit`
	__Element__00000002_.NameXSD = `credit`
	__Element__00000002_.Type = `credit`
	__Element__00000002_.MinOccurs = `0`
	__Element__00000002_.MaxOccurs = `unbounded`
	__Element__00000002_.Default = ``
	__Element__00000002_.Fixed = ``
	__Element__00000002_.Nillable = ``
	__Element__00000002_.Ref = ``
	__Element__00000002_.Abstract = ``
	__Element__00000002_.Form = ``
	__Element__00000002_.Block = ``
	__Element__00000002_.Final = ``
	__Element__00000002_.IsDuplicatedInXSD = false

	__Element__00000003_.Name = `credit-words`
	__Element__00000003_.Order = 25
	__Element__00000003_.Depth = 1
	__Element__00000003_.HasNameConflict = false
	__Element__00000003_.GoIdentifier = `Credit_words`
	__Element__00000003_.NameXSD = `credit-words`
	__Element__00000003_.Type = `string`
	__Element__00000003_.MinOccurs = ``
	__Element__00000003_.MaxOccurs = ``
	__Element__00000003_.Default = ``
	__Element__00000003_.Fixed = ``
	__Element__00000003_.Nillable = ``
	__Element__00000003_.Ref = ``
	__Element__00000003_.Abstract = ``
	__Element__00000003_.Form = ``
	__Element__00000003_.Block = ``
	__Element__00000003_.Final = ``
	__Element__00000003_.IsDuplicatedInXSD = false

	__Element__00000004_.Name = `credit-symbol`
	__Element__00000004_.Order = 26
	__Element__00000004_.Depth = 1
	__Element__00000004_.HasNameConflict = false
	__Element__00000004_.GoIdentifier = `Credit_symbol`
	__Element__00000004_.NameXSD = `credit-symbol`
	__Element__00000004_.Type = `string`
	__Element__00000004_.MinOccurs = ``
	__Element__00000004_.MaxOccurs = ``
	__Element__00000004_.Default = ``
	__Element__00000004_.Fixed = ``
	__Element__00000004_.Nillable = ``
	__Element__00000004_.Ref = ``
	__Element__00000004_.Abstract = ``
	__Element__00000004_.Form = ``
	__Element__00000004_.Block = ``
	__Element__00000004_.Final = ``
	__Element__00000004_.IsDuplicatedInXSD = false

	__Element__00000005_.Name = `link`
	__Element__00000005_.Order = 24
	__Element__00000005_.Depth = 1
	__Element__00000005_.HasNameConflict = false
	__Element__00000005_.GoIdentifier = `Link`
	__Element__00000005_.NameXSD = `link`
	__Element__00000005_.Type = `link`
	__Element__00000005_.MinOccurs = `0`
	__Element__00000005_.MaxOccurs = `unbounded`
	__Element__00000005_.Default = ``
	__Element__00000005_.Fixed = ``
	__Element__00000005_.Nillable = ``
	__Element__00000005_.Ref = ``
	__Element__00000005_.Abstract = ``
	__Element__00000005_.Form = ``
	__Element__00000005_.Block = ``
	__Element__00000005_.Final = ``
	__Element__00000005_.IsDuplicatedInXSD = false

	__Element__00000006_.Name = `credit-words`
	__Element__00000006_.Order = 22
	__Element__00000006_.Depth = 1
	__Element__00000006_.HasNameConflict = false
	__Element__00000006_.GoIdentifier = `Credit_words`
	__Element__00000006_.NameXSD = `credit-words`
	__Element__00000006_.Type = `xs:string`
	__Element__00000006_.MinOccurs = ``
	__Element__00000006_.MaxOccurs = ``
	__Element__00000006_.Default = ``
	__Element__00000006_.Fixed = ``
	__Element__00000006_.Nillable = ``
	__Element__00000006_.Ref = ``
	__Element__00000006_.Abstract = ``
	__Element__00000006_.Form = ``
	__Element__00000006_.Block = ``
	__Element__00000006_.Final = ``
	__Element__00000006_.IsDuplicatedInXSD = false

	__Element__00000007_.Name = `credit-symbol`
	__Element__00000007_.Order = 23
	__Element__00000007_.Depth = 1
	__Element__00000007_.HasNameConflict = false
	__Element__00000007_.GoIdentifier = `Credit_symbol`
	__Element__00000007_.NameXSD = `credit-symbol`
	__Element__00000007_.Type = `xs:string`
	__Element__00000007_.MinOccurs = ``
	__Element__00000007_.MaxOccurs = ``
	__Element__00000007_.Default = ``
	__Element__00000007_.Fixed = ``
	__Element__00000007_.Nillable = ``
	__Element__00000007_.Ref = ``
	__Element__00000007_.Abstract = ``
	__Element__00000007_.Form = ``
	__Element__00000007_.Block = ``
	__Element__00000007_.Final = ``
	__Element__00000007_.IsDuplicatedInXSD = false

	__Element__00000008_.Name = `credit-type`
	__Element__00000008_.Order = 20
	__Element__00000008_.Depth = 1
	__Element__00000008_.HasNameConflict = false
	__Element__00000008_.GoIdentifier = `Credit_type`
	__Element__00000008_.NameXSD = `credit-type`
	__Element__00000008_.Type = `string`
	__Element__00000008_.MinOccurs = `0`
	__Element__00000008_.MaxOccurs = `unbounded`
	__Element__00000008_.Default = ``
	__Element__00000008_.Fixed = ``
	__Element__00000008_.Nillable = ``
	__Element__00000008_.Ref = ``
	__Element__00000008_.Abstract = ``
	__Element__00000008_.Form = ``
	__Element__00000008_.Block = ``
	__Element__00000008_.Final = ``
	__Element__00000008_.IsDuplicatedInXSD = false

	__Element__00000009_.Name = `link`
	__Element__00000009_.Order = 21
	__Element__00000009_.Depth = 1
	__Element__00000009_.HasNameConflict = false
	__Element__00000009_.GoIdentifier = `Link`
	__Element__00000009_.NameXSD = `link`
	__Element__00000009_.Type = `link`
	__Element__00000009_.MinOccurs = `0`
	__Element__00000009_.MaxOccurs = `unbounded`
	__Element__00000009_.Default = ``
	__Element__00000009_.Fixed = ``
	__Element__00000009_.Nillable = ``
	__Element__00000009_.Ref = ``
	__Element__00000009_.Abstract = ``
	__Element__00000009_.Form = ``
	__Element__00000009_.Block = ``
	__Element__00000009_.Final = ``
	__Element__00000009_.IsDuplicatedInXSD = false

	__Element__00000010_.Name = `title`
	__Element__00000010_.Order = 7
	__Element__00000010_.Depth = 1
	__Element__00000010_.HasNameConflict = false
	__Element__00000010_.GoIdentifier = `Title`
	__Element__00000010_.NameXSD = `title`
	__Element__00000010_.Type = `titleType`
	__Element__00000010_.MinOccurs = ``
	__Element__00000010_.MaxOccurs = ``
	__Element__00000010_.Default = ``
	__Element__00000010_.Fixed = ``
	__Element__00000010_.Nillable = ``
	__Element__00000010_.Ref = ``
	__Element__00000010_.Abstract = ``
	__Element__00000010_.Form = ``
	__Element__00000010_.Block = ``
	__Element__00000010_.Final = ``
	__Element__00000010_.IsDuplicatedInXSD = false

	__Element__00000011_.Name = `author`
	__Element__00000011_.Order = 8
	__Element__00000011_.Depth = 1
	__Element__00000011_.HasNameConflict = false
	__Element__00000011_.GoIdentifier = `Author`
	__Element__00000011_.NameXSD = `author`
	__Element__00000011_.Type = `string`
	__Element__00000011_.MinOccurs = ``
	__Element__00000011_.MaxOccurs = ``
	__Element__00000011_.Default = ``
	__Element__00000011_.Fixed = ``
	__Element__00000011_.Nillable = ``
	__Element__00000011_.Ref = ``
	__Element__00000011_.Abstract = ``
	__Element__00000011_.Form = ``
	__Element__00000011_.Block = ``
	__Element__00000011_.Final = ``
	__Element__00000011_.IsDuplicatedInXSD = false

	__Element__00000012_.Name = `year`
	__Element__00000012_.Order = 9
	__Element__00000012_.Depth = 1
	__Element__00000012_.HasNameConflict = false
	__Element__00000012_.GoIdentifier = `Year`
	__Element__00000012_.NameXSD = `year`
	__Element__00000012_.Type = `yearType`
	__Element__00000012_.MinOccurs = ``
	__Element__00000012_.MaxOccurs = ``
	__Element__00000012_.Default = ``
	__Element__00000012_.Fixed = ``
	__Element__00000012_.Nillable = ``
	__Element__00000012_.Ref = ``
	__Element__00000012_.Abstract = ``
	__Element__00000012_.Form = ``
	__Element__00000012_.Block = ``
	__Element__00000012_.Final = ``
	__Element__00000012_.IsDuplicatedInXSD = false

	__Element__00000013_.Name = `format`
	__Element__00000013_.Order = 10
	__Element__00000013_.Depth = 1
	__Element__00000013_.HasNameConflict = false
	__Element__00000013_.GoIdentifier = `Format`
	__Element__00000013_.NameXSD = `format`
	__Element__00000013_.Type = `bookFormatEnum`
	__Element__00000013_.MinOccurs = ``
	__Element__00000013_.MaxOccurs = ``
	__Element__00000013_.Default = ``
	__Element__00000013_.Fixed = ``
	__Element__00000013_.Nillable = ``
	__Element__00000013_.Ref = ``
	__Element__00000013_.Abstract = ``
	__Element__00000013_.Form = ``
	__Element__00000013_.Block = ``
	__Element__00000013_.Final = ``
	__Element__00000013_.IsDuplicatedInXSD = false

	__Enumeration__00000000_.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__00000000_.Value = `Paperback`

	__Enumeration__00000001_.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__00000001_.Value = `Hardcover`

	__Extension__00000000_.Name = `link_Inlined_Inlined`
	__Extension__00000000_.OuterElementName = ``
	__Extension__00000000_.Order = 0
	__Extension__00000000_.Depth = 0
	__Extension__00000000_.MinOccurs = ``
	__Extension__00000000_.MaxOccurs = ``
	__Extension__00000000_.Base = `string`
	__Extension__00000000_.Ref = ``

	__Group__00000000_.Name = `bookType_S__G_`
	__Group__00000000_.NameXSD = ``
	__Group__00000000_.Ref = `bookDetailsGroup`
	__Group__00000000_.IsAnonymous = false
	__Group__00000000_.HasNameConflict = false
	__Group__00000000_.GoIdentifier = `BookDetailsGroup`
	__Group__00000000_.OuterElementName = `bookType_S__G_`
	__Group__00000000_.Order = 3
	__Group__00000000_.Depth = 1
	__Group__00000000_.MinOccurs = ``
	__Group__00000000_.MaxOccurs = ``

	__Group__00000001_.Name = `bookDetailsGroup`
	__Group__00000001_.NameXSD = `bookDetailsGroup`
	__Group__00000001_.Ref = ``
	__Group__00000001_.IsAnonymous = false
	__Group__00000001_.HasNameConflict = false
	__Group__00000001_.GoIdentifier = `Group_bookDetailsGroup`
	__Group__00000001_.OuterElementName = `bookDetailsGroup`
	__Group__00000001_.Order = 6
	__Group__00000001_.Depth = 0
	__Group__00000001_.MinOccurs = ``
	__Group__00000001_.MaxOccurs = ``

	__MaxInclusive__00000000_.Name = `yearType_Inlined_Inlined`
	__MaxInclusive__00000000_.Value = `2100`

	__MinInclusive__00000000_.Name = `yearType_Inlined_Inlined`
	__MinInclusive__00000000_.Value = `1000`

	__Pattern__00000000_.Name = `titleType_Inlined_Inlined`
	__Pattern__00000000_.Value = `[0-9A-Za-z ]+`

	__Restriction__00000000_.Name = `yearType_Inlined`
	__Restriction__00000000_.Base = `integer`

	__Restriction__00000001_.Name = `bookFormatEnum_Inlined`
	__Restriction__00000001_.Base = `string`

	__Restriction__00000002_.Name = `titleType_Inlined`
	__Restriction__00000002_.Base = `string`

	__Schema__00000000_.Name = `Schema`
	__Schema__00000000_.Xs = `http://www.w3.org/2001/XMLSchema`
	__Schema__00000000_.Order = 0
	__Schema__00000000_.Depth = 0

	__Sequence__00000000_.Name = ``
	__Sequence__00000000_.OuterElementName = ``
	__Sequence__00000000_.Order = 0
	__Sequence__00000000_.Depth = 0
	__Sequence__00000000_.MinOccurs = ``
	__Sequence__00000000_.MaxOccurs = ``

	__Sequence__00000001_.Name = `bookType_S_`
	__Sequence__00000001_.OuterElementName = `bookType_S_`
	__Sequence__00000001_.Order = 0
	__Sequence__00000001_.Depth = 0
	__Sequence__00000001_.MinOccurs = ``
	__Sequence__00000001_.MaxOccurs = ``

	__Sequence__00000002_.Name = `credit_S_`
	__Sequence__00000002_.OuterElementName = `credit_S_`
	__Sequence__00000002_.Order = 0
	__Sequence__00000002_.Depth = 0
	__Sequence__00000002_.MinOccurs = ``
	__Sequence__00000002_.MaxOccurs = ``

	__Sequence__00000003_.Name = `credit_S__C__S_`
	__Sequence__00000003_.OuterElementName = `credit_S__C__S_`
	__Sequence__00000003_.Order = 0
	__Sequence__00000003_.Depth = 0
	__Sequence__00000003_.MinOccurs = ``
	__Sequence__00000003_.MaxOccurs = ``

	__Sequence__00000004_.Name = `credit_S__C__S__S_`
	__Sequence__00000004_.OuterElementName = `credit_S__C__S__S_`
	__Sequence__00000004_.Order = 0
	__Sequence__00000004_.Depth = 0
	__Sequence__00000004_.MinOccurs = `0`
	__Sequence__00000004_.MaxOccurs = `unbounded`

	__Sequence__00000005_.Name = `bookDetailsGroup_S_`
	__Sequence__00000005_.OuterElementName = `bookDetailsGroup_S_`
	__Sequence__00000005_.Order = 0
	__Sequence__00000005_.Depth = 0
	__Sequence__00000005_.MinOccurs = ``
	__Sequence__00000005_.MaxOccurs = ``

	__SimpleContent__00000000_.Name = `link_Inlined`

	__SimpleType__00000000_.Name = `yearType`
	__SimpleType__00000000_.NameXSD = `yearType`
	__SimpleType__00000000_.Order = 1
	__SimpleType__00000000_.Depth = 0

	__SimpleType__00000001_.Name = `bookFormatEnum`
	__SimpleType__00000001_.NameXSD = `bookFormatEnum`
	__SimpleType__00000001_.Order = 11
	__SimpleType__00000001_.Depth = 0

	__SimpleType__00000002_.Name = `titleType`
	__SimpleType__00000002_.NameXSD = `titleType`
	__SimpleType__00000002_.Order = 12
	__SimpleType__00000002_.Depth = 0

	__WhiteSpace__00000000_.Name = `titleType_Inlined_Inlined`
	__WhiteSpace__00000000_.Value = `collapse`

	// insertion point for setup of pointers
	__Annotation__00000000_.Documentations = append(__Annotation__00000000_.Documentations, __Documentation__00000000_)
	__Annotation__00000000_.Documentations = append(__Annotation__00000000_.Documentations, __Documentation__00000001_)
	__Annotation__00000001_.Documentations = append(__Annotation__00000001_.Documentations, __Documentation__00000002_)
	__Annotation__00000002_.Documentations = append(__Annotation__00000002_.Documentations, __Documentation__00000003_)
	__Annotation__00000003_.Documentations = append(__Annotation__00000003_.Documentations, __Documentation__00000004_)
	__Annotation__00000004_.Documentations = append(__Annotation__00000004_.Documentations, __Documentation__00000005_)
	__Annotation__00000005_.Documentations = append(__Annotation__00000005_.Documentations, __Documentation__00000006_)
	__Annotation__00000006_.Documentations = append(__Annotation__00000006_.Documentations, __Documentation__00000007_)
	__Annotation__00000007_.Documentations = append(__Annotation__00000007_.Documentations, __Documentation__00000008_)
	__Annotation__00000008_.Documentations = append(__Annotation__00000008_.Documentations, __Documentation__00000009_)
	__Annotation__00000009_.Documentations = append(__Annotation__00000009_.Documentations, __Documentation__00000010_)
	__Annotation__00000010_.Documentations = append(__Annotation__00000010_.Documentations, __Documentation__00000011_)
	__Attribute__00000000_.Annotation = nil
	__Attribute__00000001_.Annotation = nil
	__Attribute__00000002_.Annotation = __Annotation__00000006_
	__Attribute__00000003_.Annotation = __Annotation__00000007_
	__Attribute__00000004_.Annotation = __Annotation__00000008_
	__AttributeGroup__00000000_.Annotation = nil
	__AttributeGroup__00000001_.Annotation = nil
	__AttributeGroup__00000001_.Attributes = append(__AttributeGroup__00000001_.Attributes, __Attribute__00000002_)
	__AttributeGroup__00000001_.Attributes = append(__AttributeGroup__00000001_.Attributes, __Attribute__00000003_)
	__AttributeGroup__00000002_.Annotation = nil
	__AttributeGroup__00000002_.AttributeGroups = append(__AttributeGroup__00000002_.AttributeGroups, __AttributeGroup__00000003_)
	__AttributeGroup__00000002_.Attributes = append(__AttributeGroup__00000002_.Attributes, __Attribute__00000004_)
	__AttributeGroup__00000003_.Annotation = nil
	__Choice__00000000_.Annotation = nil
	__Choice__00000000_.Sequences = append(__Choice__00000000_.Sequences, __Sequence__00000003_)
	__Choice__00000001_.Annotation = nil
	__Choice__00000001_.Elements = append(__Choice__00000001_.Elements, __Element__00000003_)
	__Choice__00000001_.Elements = append(__Choice__00000001_.Elements, __Element__00000004_)
	__Choice__00000002_.Annotation = nil
	__Choice__00000002_.Elements = append(__Choice__00000002_.Elements, __Element__00000006_)
	__Choice__00000002_.Elements = append(__Choice__00000002_.Elements, __Element__00000007_)
	__ComplexType__00000000_.OuterElement = __Element__00000000_
	__ComplexType__00000000_.Annotation = nil
	__ComplexType__00000000_.Sequences = append(__ComplexType__00000000_.Sequences, __Sequence__00000000_)
	__ComplexType__00000000_.Extension = nil
	__ComplexType__00000000_.SimpleContent = nil
	__ComplexType__00000000_.ComplexContent = nil
	__ComplexType__00000001_.OuterElement = nil
	__ComplexType__00000001_.Annotation = __Annotation__00000003_
	__ComplexType__00000001_.Sequences = append(__ComplexType__00000001_.Sequences, __Sequence__00000001_)
	__ComplexType__00000001_.Extension = nil
	__ComplexType__00000001_.SimpleContent = nil
	__ComplexType__00000001_.ComplexContent = nil
	__ComplexType__00000001_.AttributeGroups = append(__ComplexType__00000001_.AttributeGroups, __AttributeGroup__00000000_)
	__ComplexType__00000002_.OuterElement = nil
	__ComplexType__00000002_.Annotation = __Annotation__00000004_
	__ComplexType__00000002_.Sequences = append(__ComplexType__00000002_.Sequences, __Sequence__00000002_)
	__ComplexType__00000002_.Extension = nil
	__ComplexType__00000002_.SimpleContent = nil
	__ComplexType__00000002_.ComplexContent = nil
	__ComplexType__00000002_.Attributes = append(__ComplexType__00000002_.Attributes, __Attribute__00000000_)
	__ComplexType__00000003_.OuterElement = nil
	__ComplexType__00000003_.Annotation = __Annotation__00000005_
	__ComplexType__00000003_.Extension = nil
	__ComplexType__00000003_.SimpleContent = __SimpleContent__00000000_
	__ComplexType__00000003_.ComplexContent = nil
	__Element__00000000_.Annotation = nil
	__Element__00000000_.SimpleType = nil
	__Element__00000000_.ComplexType = __ComplexType__00000000_
	__Element__00000001_.Annotation = __Annotation__00000001_
	__Element__00000001_.SimpleType = nil
	__Element__00000001_.ComplexType = nil
	__Element__00000002_.Annotation = nil
	__Element__00000002_.SimpleType = nil
	__Element__00000002_.ComplexType = nil
	__Element__00000003_.Annotation = nil
	__Element__00000003_.SimpleType = nil
	__Element__00000003_.ComplexType = nil
	__Element__00000004_.Annotation = nil
	__Element__00000004_.SimpleType = nil
	__Element__00000004_.ComplexType = nil
	__Element__00000005_.Annotation = nil
	__Element__00000005_.SimpleType = nil
	__Element__00000005_.ComplexType = nil
	__Element__00000006_.Annotation = nil
	__Element__00000006_.SimpleType = nil
	__Element__00000006_.ComplexType = nil
	__Element__00000007_.Annotation = nil
	__Element__00000007_.SimpleType = nil
	__Element__00000007_.ComplexType = nil
	__Element__00000008_.Annotation = nil
	__Element__00000008_.SimpleType = nil
	__Element__00000008_.ComplexType = nil
	__Element__00000009_.Annotation = nil
	__Element__00000009_.SimpleType = nil
	__Element__00000009_.ComplexType = nil
	__Element__00000010_.Annotation = __Annotation__00000009_
	__Element__00000010_.SimpleType = nil
	__Element__00000010_.ComplexType = nil
	__Element__00000011_.Annotation = nil
	__Element__00000011_.SimpleType = nil
	__Element__00000011_.ComplexType = nil
	__Element__00000012_.Annotation = nil
	__Element__00000012_.SimpleType = nil
	__Element__00000012_.ComplexType = nil
	__Element__00000013_.Annotation = __Annotation__00000010_
	__Element__00000013_.SimpleType = nil
	__Element__00000013_.ComplexType = nil
	__Enumeration__00000000_.Annotation = nil
	__Enumeration__00000001_.Annotation = nil
	__Extension__00000000_.Attributes = append(__Extension__00000000_.Attributes, __Attribute__00000001_)
	__Group__00000000_.Annotation = nil
	__Group__00000000_.OuterElement = nil
	__Group__00000001_.Annotation = nil
	__Group__00000001_.OuterElement = nil
	__Group__00000001_.Sequences = append(__Group__00000001_.Sequences, __Sequence__00000005_)
	__MaxInclusive__00000000_.Annotation = nil
	__MinInclusive__00000000_.Annotation = nil
	__Pattern__00000000_.Annotation = nil
	__Restriction__00000000_.Annotation = nil
	__Restriction__00000000_.MinInclusive = __MinInclusive__00000000_
	__Restriction__00000000_.MaxInclusive = __MaxInclusive__00000000_
	__Restriction__00000000_.Pattern = nil
	__Restriction__00000000_.WhiteSpace = nil
	__Restriction__00000000_.MinLength = nil
	__Restriction__00000000_.MaxLength = nil
	__Restriction__00000000_.Length = nil
	__Restriction__00000000_.TotalDigit = nil
	__Restriction__00000001_.Annotation = nil
	__Restriction__00000001_.Enumerations = append(__Restriction__00000001_.Enumerations, __Enumeration__00000000_)
	__Restriction__00000001_.Enumerations = append(__Restriction__00000001_.Enumerations, __Enumeration__00000001_)
	__Restriction__00000001_.MinInclusive = nil
	__Restriction__00000001_.MaxInclusive = nil
	__Restriction__00000001_.Pattern = nil
	__Restriction__00000001_.WhiteSpace = nil
	__Restriction__00000001_.MinLength = nil
	__Restriction__00000001_.MaxLength = nil
	__Restriction__00000001_.Length = nil
	__Restriction__00000001_.TotalDigit = nil
	__Restriction__00000002_.Annotation = nil
	__Restriction__00000002_.MinInclusive = nil
	__Restriction__00000002_.MaxInclusive = nil
	__Restriction__00000002_.Pattern = __Pattern__00000000_
	__Restriction__00000002_.WhiteSpace = __WhiteSpace__00000000_
	__Restriction__00000002_.MinLength = nil
	__Restriction__00000002_.MaxLength = nil
	__Restriction__00000002_.Length = nil
	__Restriction__00000002_.TotalDigit = nil
	__Schema__00000000_.Annotation = __Annotation__00000000_
	__Schema__00000000_.Elements = append(__Schema__00000000_.Elements, __Element__00000000_)
	__Schema__00000000_.SimpleTypes = append(__Schema__00000000_.SimpleTypes, __SimpleType__00000000_)
	__Schema__00000000_.SimpleTypes = append(__Schema__00000000_.SimpleTypes, __SimpleType__00000001_)
	__Schema__00000000_.SimpleTypes = append(__Schema__00000000_.SimpleTypes, __SimpleType__00000002_)
	__Schema__00000000_.ComplexTypes = append(__Schema__00000000_.ComplexTypes, __ComplexType__00000001_)
	__Schema__00000000_.ComplexTypes = append(__Schema__00000000_.ComplexTypes, __ComplexType__00000002_)
	__Schema__00000000_.ComplexTypes = append(__Schema__00000000_.ComplexTypes, __ComplexType__00000003_)
	__Schema__00000000_.AttributeGroups = append(__Schema__00000000_.AttributeGroups, __AttributeGroup__00000001_)
	__Schema__00000000_.AttributeGroups = append(__Schema__00000000_.AttributeGroups, __AttributeGroup__00000002_)
	__Schema__00000000_.Groups = append(__Schema__00000000_.Groups, __Group__00000001_)
	__Sequence__00000000_.Annotation = nil
	__Sequence__00000000_.Elements = append(__Sequence__00000000_.Elements, __Element__00000001_)
	__Sequence__00000001_.Annotation = nil
	__Sequence__00000001_.Groups = append(__Sequence__00000001_.Groups, __Group__00000000_)
	__Sequence__00000001_.Elements = append(__Sequence__00000001_.Elements, __Element__00000002_)
	__Sequence__00000002_.Annotation = nil
	__Sequence__00000002_.Choices = append(__Sequence__00000002_.Choices, __Choice__00000000_)
	__Sequence__00000002_.Elements = append(__Sequence__00000002_.Elements, __Element__00000008_)
	__Sequence__00000002_.Elements = append(__Sequence__00000002_.Elements, __Element__00000009_)
	__Sequence__00000003_.Annotation = nil
	__Sequence__00000003_.Sequences = append(__Sequence__00000003_.Sequences, __Sequence__00000004_)
	__Sequence__00000003_.Choices = append(__Sequence__00000003_.Choices, __Choice__00000002_)
	__Sequence__00000004_.Annotation = nil
	__Sequence__00000004_.Choices = append(__Sequence__00000004_.Choices, __Choice__00000001_)
	__Sequence__00000004_.Elements = append(__Sequence__00000004_.Elements, __Element__00000005_)
	__Sequence__00000005_.Annotation = nil
	__Sequence__00000005_.Elements = append(__Sequence__00000005_.Elements, __Element__00000010_)
	__Sequence__00000005_.Elements = append(__Sequence__00000005_.Elements, __Element__00000011_)
	__Sequence__00000005_.Elements = append(__Sequence__00000005_.Elements, __Element__00000012_)
	__Sequence__00000005_.Elements = append(__Sequence__00000005_.Elements, __Element__00000013_)
	__SimpleContent__00000000_.Extension = __Extension__00000000_
	__SimpleContent__00000000_.Restriction = nil
	__SimpleType__00000000_.Annotation = __Annotation__00000002_
	__SimpleType__00000000_.Restriction = __Restriction__00000000_
	__SimpleType__00000000_.Union = nil
	__SimpleType__00000001_.Annotation = nil
	__SimpleType__00000001_.Restriction = __Restriction__00000001_
	__SimpleType__00000001_.Union = nil
	__SimpleType__00000002_.Annotation = nil
	__SimpleType__00000002_.Restriction = __Restriction__00000002_
	__SimpleType__00000002_.Union = nil
	__WhiteSpace__00000000_.Annotation = nil
}
