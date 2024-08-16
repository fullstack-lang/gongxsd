// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// AnonymousComplexTypeInline_books is generated from outer element "books"
type AnonymousComplexTypeInline_books struct {

	// insertion point for fields

	// generated from element "book" of type bookType
	Book []*BookType `xml:"book"`
}

// BookType is generated from named complex type "bookType"
type BookType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "extendedAttributes
	AttributeGroup_extendedAttributes

	// generated from element "credit" of type credit
	Credit []*Credit `xml:"credit"`

	// generated from element "title" of type titleType
	Title string `xml:"title"`

	// generated from element "author" of type xs:string
	Author string `xml:"author"`

	// generated from element "year" of type yearType
	Year int `xml:"year"`

	// generated from element "format" of type bookFormatEnum
	Format string `xml:"format"`
}

// Credit is generated from named complex type "credit"
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page" of type xs:positiveInteger
	Page int `xml:"page,attr"`

	// generated from element "credit-type" of type xs:string
	Credit_type string `xml:"credit-type"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "credit-words" of type xs:string
	Credit_words string `xml:"credit-words"`

	// generated from element "credit-symbol" of type xs:string
	Credit_symbol string `xml:"credit-symbol"`
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_bookDetailsGroup is generated from named group "bookDetailsGroup"
type Group_bookDetailsGroup struct {

	// insertion point for fields

	// generated from element "title" of type titleType
	Title string `xml:"title"`

	// generated from element "author" of type xs:string
	Author string `xml:"author"`

	// generated from element "year" of type yearType
	Year int `xml:"year"`

	// generated from element "format" of type bookFormatEnum
	Format string `xml:"format"`
}

// AttributeGroup_commonAttributes is generated from named attribute group "commonAttributes"
type AttributeGroup_commonAttributes struct {

	// insertion point for fields

	// generated from attribute "isbn" of type xs:string
	Isbn string `xml:"isbn,attr"`

	// generated from attribute "bestseller" of type xs:boolean
	Bestseller bool `xml:"bestseller,attr"`
}

// AttributeGroup_extendedAttributes is generated from named attribute group "extendedAttributes"
type AttributeGroup_extendedAttributes struct {

	// insertion point for fields

	// generated from attribute "edition" of type xs:string
	Edition string `xml:"edition,attr"`

	// generated from attribute group "commonAttributes
	AttributeGroup_commonAttributes
}

// Books is generated from element books within root schema
type Books struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from inline complex type
	AnonymousComplexTypeInline_books
}
