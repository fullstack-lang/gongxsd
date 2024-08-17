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
	__Annotation__000003_credit_Inlined := (&models.Annotation{Name: `credit_Inlined`}).Stage(stage)
	__Annotation__000004_edition_Inlined := (&models.Annotation{Name: `edition_Inlined`}).Stage(stage)
	__Annotation__000005_isbn_Inlined := (&models.Annotation{Name: `isbn_Inlined`}).Stage(stage)
	__Annotation__000006_link_Inlined := (&models.Annotation{Name: `link_Inlined`}).Stage(stage)
	__Annotation__000007_yearType_Inlined := (&models.Annotation{Name: `yearType_Inlined`}).Stage(stage)

	__Attribute__000000_bestseller := (&models.Attribute{Name: `bestseller`}).Stage(stage)
	__Attribute__000001_edition := (&models.Attribute{Name: `edition`}).Stage(stage)
	__Attribute__000002_isbn := (&models.Attribute{Name: `isbn`}).Stage(stage)
	__Attribute__000003_name := (&models.Attribute{Name: `name`}).Stage(stage)
	__Attribute__000004_page := (&models.Attribute{Name: `page`}).Stage(stage)

	__AttributeGroup__000000_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000001_ := (&models.AttributeGroup{Name: ``}).Stage(stage)
	__AttributeGroup__000002_commonAttributes := (&models.AttributeGroup{Name: `commonAttributes`}).Stage(stage)
	__AttributeGroup__000003_extendedAttributes := (&models.AttributeGroup{Name: `extendedAttributes`}).Stage(stage)

	__ComplexType__000000_AnonymousComplexTypeInline_books := (&models.ComplexType{Name: `AnonymousComplexTypeInline_books`}).Stage(stage)
	__ComplexType__000001_bookType := (&models.ComplexType{Name: `bookType`}).Stage(stage)
	__ComplexType__000002_credit := (&models.ComplexType{Name: `credit`}).Stage(stage)
	__ComplexType__000003_link := (&models.ComplexType{Name: `link`}).Stage(stage)

	__Documentation__000000_Schema_Inlined_Inlined := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__000001_Schema_Inlined_Inlined := (&models.Documentation{Name: `Schema_Inlined_Inlined`}).Stage(stage)
	__Documentation__000002__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000003__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000004__Inlined := (&models.Documentation{Name: `_Inlined`}).Stage(stage)
	__Documentation__000005_bookType_Inlined_Inlined := (&models.Documentation{Name: `bookType_Inlined_Inlined`}).Stage(stage)
	__Documentation__000006_credit_Inlined_Inlined := (&models.Documentation{Name: `credit_Inlined_Inlined`}).Stage(stage)
	__Documentation__000007_link_Inlined_Inlined := (&models.Documentation{Name: `link_Inlined_Inlined`}).Stage(stage)
	__Documentation__000008_yearType_Inlined_Inlined := (&models.Documentation{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__Element__000000_books := (&models.Element{Name: `books`}).Stage(stage)

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined := (&models.Enumeration{Name: `bookFormatEnum_Inlined_Inlined`}).Stage(stage)

	__Extension__000000_link_Inlined_Inlined := (&models.Extension{Name: `link_Inlined_Inlined`}).Stage(stage)

	__Group__000000_bookDetailsGroup := (&models.Group{Name: `bookDetailsGroup`}).Stage(stage)

	__MaxInclusive__000000_yearType_Inlined_Inlined := (&models.MaxInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__MinInclusive__000000_yearType_Inlined_Inlined := (&models.MinInclusive{Name: `yearType_Inlined_Inlined`}).Stage(stage)

	__Pattern__000000_titleType_Inlined_Inlined := (&models.Pattern{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	__Restriction__000000_bookFormatEnum_Inlined := (&models.Restriction{Name: `bookFormatEnum_Inlined`}).Stage(stage)
	__Restriction__000001_titleType_Inlined := (&models.Restriction{Name: `titleType_Inlined`}).Stage(stage)
	__Restriction__000002_yearType_Inlined := (&models.Restriction{Name: `yearType_Inlined`}).Stage(stage)

	__Schema__000000_Schema := (&models.Schema{Name: `Schema`}).Stage(stage)

	__SimpleContent__000000_link_Inlined := (&models.SimpleContent{Name: `link_Inlined`}).Stage(stage)

	__SimpleType__000000_bookFormatEnum := (&models.SimpleType{Name: `bookFormatEnum`}).Stage(stage)
	__SimpleType__000001_titleType := (&models.SimpleType{Name: `titleType`}).Stage(stage)
	__SimpleType__000002_yearType := (&models.SimpleType{Name: `yearType`}).Stage(stage)

	__WhiteSpace__000000_titleType_Inlined_Inlined := (&models.WhiteSpace{Name: `titleType_Inlined_Inlined`}).Stage(stage)

	// Setup of values

	__Annotation__000000_Schema_Inlined.Name = `Schema_Inlined`

	__Annotation__000001_bestseller_Inlined.Name = `bestseller_Inlined`

	__Annotation__000002_bookType_Inlined.Name = `bookType_Inlined`

	__Annotation__000003_credit_Inlined.Name = `credit_Inlined`

	__Annotation__000004_edition_Inlined.Name = `edition_Inlined`

	__Annotation__000005_isbn_Inlined.Name = `isbn_Inlined`

	__Annotation__000006_link_Inlined.Name = `link_Inlined`

	__Annotation__000007_yearType_Inlined.Name = `yearType_Inlined`

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
	__AttributeGroup__000000_.HasNameConflict = false
	__AttributeGroup__000000_.GoIdentifier = `AttributeGroup_`
	__AttributeGroup__000000_.Ref = `extendedAttributes`

	__AttributeGroup__000001_.Name = ``
	__AttributeGroup__000001_.NameXSD = ``
	__AttributeGroup__000001_.HasNameConflict = true
	__AttributeGroup__000001_.GoIdentifier = `AttributeGroup__1`
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

	__ComplexType__000000_AnonymousComplexTypeInline_books.Name = `AnonymousComplexTypeInline_books`
	__ComplexType__000000_AnonymousComplexTypeInline_books.HasNameConflict = false
	__ComplexType__000000_AnonymousComplexTypeInline_books.GoIdentifier = `AnonymousComplexTypeInline_books`
	__ComplexType__000000_AnonymousComplexTypeInline_books.IsAnonymous = true
	__ComplexType__000000_AnonymousComplexTypeInline_books.NameXSD = ``

	__ComplexType__000001_bookType.Name = `bookType`
	__ComplexType__000001_bookType.HasNameConflict = false
	__ComplexType__000001_bookType.GoIdentifier = `BookType`
	__ComplexType__000001_bookType.IsAnonymous = false
	__ComplexType__000001_bookType.NameXSD = `bookType`

	__ComplexType__000002_credit.Name = `credit`
	__ComplexType__000002_credit.HasNameConflict = false
	__ComplexType__000002_credit.GoIdentifier = `Credit`
	__ComplexType__000002_credit.IsAnonymous = false
	__ComplexType__000002_credit.NameXSD = `credit`

	__ComplexType__000003_link.Name = `link`
	__ComplexType__000003_link.HasNameConflict = false
	__ComplexType__000003_link.GoIdentifier = `Link`
	__ComplexType__000003_link.IsAnonymous = false
	__ComplexType__000003_link.NameXSD = `link`

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
	__Documentation__000002__Inlined.Text = `The ISBN number of the book.`
	__Documentation__000002__Inlined.Source = ``
	__Documentation__000002__Inlined.Lang = ``

	__Documentation__000003__Inlined.Name = `_Inlined`
	__Documentation__000003__Inlined.Text = `The edition of the book.`
	__Documentation__000003__Inlined.Source = ``
	__Documentation__000003__Inlined.Lang = ``

	__Documentation__000004__Inlined.Name = `_Inlined`
	__Documentation__000004__Inlined.Text = `Indicates if the book is a bestseller.`
	__Documentation__000004__Inlined.Source = ``
	__Documentation__000004__Inlined.Lang = ``

	__Documentation__000005_bookType_Inlined_Inlined.Name = `bookType_Inlined_Inlined`
	__Documentation__000005_bookType_Inlined_Inlined.Text = ` This complex type defines the structure of a book, including title,
                author, year, and format. `
	__Documentation__000005_bookType_Inlined_Inlined.Source = ``
	__Documentation__000005_bookType_Inlined_Inlined.Lang = ``

	__Documentation__000006_credit_Inlined_Inlined.Name = `credit_Inlined_Inlined`
	__Documentation__000006_credit_Inlined_Inlined.Text = `The credit type .. `
	__Documentation__000006_credit_Inlined_Inlined.Source = ``
	__Documentation__000006_credit_Inlined_Inlined.Lang = ``

	__Documentation__000007_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Documentation__000007_link_Inlined_Inlined.Text = `The link type serves as an outgoing simple XLink. If a relative link
                is used within a document that is part of a compressed MusicXML file, the link is
                relative to the root folder of the zip file.`
	__Documentation__000007_link_Inlined_Inlined.Source = ``
	__Documentation__000007_link_Inlined_Inlined.Lang = ``

	__Documentation__000008_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__Documentation__000008_yearType_Inlined_Inlined.Text = ` This type represents a year. It restricts the value to an integer
                between 1000 and 2100 inclusive. `
	__Documentation__000008_yearType_Inlined_Inlined.Source = ``
	__Documentation__000008_yearType_Inlined_Inlined.Lang = ``

	__Element__000000_books.Name = `books`
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

	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000000_bookFormatEnum_Inlined_Inlined.Value = `Paperback`

	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Name = `bookFormatEnum_Inlined_Inlined`
	__Enumeration__000001_bookFormatEnum_Inlined_Inlined.Value = `Hardcover`

	__Extension__000000_link_Inlined_Inlined.Name = `link_Inlined_Inlined`
	__Extension__000000_link_Inlined_Inlined.Base = `xs:string`

	__Group__000000_bookDetailsGroup.Name = `bookDetailsGroup`
	__Group__000000_bookDetailsGroup.NameXSD = `bookDetailsGroup`
	__Group__000000_bookDetailsGroup.Ref = ``
	__Group__000000_bookDetailsGroup.IsAnonymous = false
	__Group__000000_bookDetailsGroup.HasNameConflict = false
	__Group__000000_bookDetailsGroup.GoIdentifier = `Group_bookDetailsGroup`

	__MaxInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MaxInclusive__000000_yearType_Inlined_Inlined.Value = `2100`

	__MinInclusive__000000_yearType_Inlined_Inlined.Name = `yearType_Inlined_Inlined`
	__MinInclusive__000000_yearType_Inlined_Inlined.Value = `1000`

	__Pattern__000000_titleType_Inlined_Inlined.Name = `titleType_Inlined_Inlined`
	__Pattern__000000_titleType_Inlined_Inlined.Value = `[0-9A-Za-z ]+`

	__Restriction__000000_bookFormatEnum_Inlined.Name = `bookFormatEnum_Inlined`
	__Restriction__000000_bookFormatEnum_Inlined.Base = `xs:string`

	__Restriction__000001_titleType_Inlined.Name = `titleType_Inlined`
	__Restriction__000001_titleType_Inlined.Base = `xs:string`

	__Restriction__000002_yearType_Inlined.Name = `yearType_Inlined`
	__Restriction__000002_yearType_Inlined.Base = `xs:integer`

	__Schema__000000_Schema.Name = `Schema`
	__Schema__000000_Schema.Xs = `http://www.w3.org/2001/XMLSchema`

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
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000000_Schema_Inlined_Inlined)
	__Annotation__000000_Schema_Inlined.Documentations = append(__Annotation__000000_Schema_Inlined.Documentations, __Documentation__000001_Schema_Inlined_Inlined)
	__Annotation__000001_bestseller_Inlined.Documentations = append(__Annotation__000001_bestseller_Inlined.Documentations, __Documentation__000004__Inlined)
	__Annotation__000002_bookType_Inlined.Documentations = append(__Annotation__000002_bookType_Inlined.Documentations, __Documentation__000005_bookType_Inlined_Inlined)
	__Annotation__000003_credit_Inlined.Documentations = append(__Annotation__000003_credit_Inlined.Documentations, __Documentation__000006_credit_Inlined_Inlined)
	__Annotation__000004_edition_Inlined.Documentations = append(__Annotation__000004_edition_Inlined.Documentations, __Documentation__000003__Inlined)
	__Annotation__000005_isbn_Inlined.Documentations = append(__Annotation__000005_isbn_Inlined.Documentations, __Documentation__000002__Inlined)
	__Annotation__000006_link_Inlined.Documentations = append(__Annotation__000006_link_Inlined.Documentations, __Documentation__000007_link_Inlined_Inlined)
	__Annotation__000007_yearType_Inlined.Documentations = append(__Annotation__000007_yearType_Inlined.Documentations, __Documentation__000008_yearType_Inlined_Inlined)
	__Attribute__000000_bestseller.Annotation = __Annotation__000001_bestseller_Inlined
	__Attribute__000001_edition.Annotation = __Annotation__000004_edition_Inlined
	__Attribute__000002_isbn.Annotation = __Annotation__000005_isbn_Inlined
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000002_isbn)
	__AttributeGroup__000002_commonAttributes.Attributes = append(__AttributeGroup__000002_commonAttributes.Attributes, __Attribute__000000_bestseller)
	__AttributeGroup__000003_extendedAttributes.AttributeGroups = append(__AttributeGroup__000003_extendedAttributes.AttributeGroups, __AttributeGroup__000001_)
	__AttributeGroup__000003_extendedAttributes.Attributes = append(__AttributeGroup__000003_extendedAttributes.Attributes, __Attribute__000001_edition)
	__ComplexType__000000_AnonymousComplexTypeInline_books.OuterElement = __Element__000000_books
	__ComplexType__000001_bookType.Annotation = __Annotation__000002_bookType_Inlined
	__ComplexType__000001_bookType.AttributeGroups = append(__ComplexType__000001_bookType.AttributeGroups, __AttributeGroup__000000_)
	__ComplexType__000002_credit.Annotation = __Annotation__000003_credit_Inlined
	__ComplexType__000002_credit.Attributes = append(__ComplexType__000002_credit.Attributes, __Attribute__000004_page)
	__ComplexType__000003_link.Annotation = __Annotation__000006_link_Inlined
	__ComplexType__000003_link.SimpleContent = __SimpleContent__000000_link_Inlined
	__Element__000000_books.ComplexType = __ComplexType__000000_AnonymousComplexTypeInline_books
	__Extension__000000_link_Inlined_Inlined.Attributes = append(__Extension__000000_link_Inlined_Inlined.Attributes, __Attribute__000003_name)
	__Restriction__000000_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000000_bookFormatEnum_Inlined.Enumerations, __Enumeration__000000_bookFormatEnum_Inlined_Inlined)
	__Restriction__000000_bookFormatEnum_Inlined.Enumerations = append(__Restriction__000000_bookFormatEnum_Inlined.Enumerations, __Enumeration__000001_bookFormatEnum_Inlined_Inlined)
	__Restriction__000001_titleType_Inlined.Pattern = __Pattern__000000_titleType_Inlined_Inlined
	__Restriction__000001_titleType_Inlined.WhiteSpace = __WhiteSpace__000000_titleType_Inlined_Inlined
	__Restriction__000002_yearType_Inlined.MinInclusive = __MinInclusive__000000_yearType_Inlined_Inlined
	__Restriction__000002_yearType_Inlined.MaxInclusive = __MaxInclusive__000000_yearType_Inlined_Inlined
	__Schema__000000_Schema.Annotation = __Annotation__000000_Schema_Inlined
	__Schema__000000_Schema.Elements = append(__Schema__000000_Schema.Elements, __Element__000000_books)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000002_yearType)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000000_bookFormatEnum)
	__Schema__000000_Schema.SimpleTypes = append(__Schema__000000_Schema.SimpleTypes, __SimpleType__000001_titleType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000001_bookType)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000002_credit)
	__Schema__000000_Schema.ComplexTypes = append(__Schema__000000_Schema.ComplexTypes, __ComplexType__000003_link)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000002_commonAttributes)
	__Schema__000000_Schema.AttributeGroups = append(__Schema__000000_Schema.AttributeGroups, __AttributeGroup__000003_extendedAttributes)
	__Schema__000000_Schema.Groups = append(__Schema__000000_Schema.Groups, __Group__000000_bookDetailsGroup)
	__SimpleContent__000000_link_Inlined.Extension = __Extension__000000_link_Inlined_Inlined
	__SimpleType__000000_bookFormatEnum.Restriction = __Restriction__000000_bookFormatEnum_Inlined
	__SimpleType__000001_titleType.Restriction = __Restriction__000001_titleType_Inlined
	__SimpleType__000002_yearType.Annotation = __Annotation__000007_yearType_Inlined
	__SimpleType__000002_yearType.Restriction = __Restriction__000002_yearType_Inlined
}
