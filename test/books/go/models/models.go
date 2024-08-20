// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// A_books is generated from outer element "books"
type A_books struct {

	// insertion point for fields

	// generated from element "book" of type bookType order 8 depth 1
	Book []*BookType `xml:"book,omitempty"`
}

// BookType is generated from named complex type "bookType"
type BookType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "extendedAttributes
	AttributeGroup_extendedAttributes

	// generated from group with order 0 depth 0
	Group_bookDetailsGroup

	// generated from element "credit" of type credit order 1 depth 0
	Credit []*Credit `xml:"credit,omitempty"`
}

// Credit is generated from named complex type "credit"
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page
	Page int `xml:"page,attr,omitempty"`

	// generated from element "credit-type" of type string order 9 depth 0
	Credit_type string `xml:"credit-type,omitempty"`

	// generated from element "link" of type link order 13 depth 0
	Link []*Link `xml:"link,omitempty"`

	// generated from element "credit-words" of type string order 14 depth 0
	Credit_words string `xml:"credit-words,omitempty"`

	// generated from element "credit-symbol" of type string order 15 depth 0
	Credit_symbol string `xml:"credit-symbol,omitempty"`
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_bookDetailsGroup is generated from named group "bookDetailsGroup"
type Group_bookDetailsGroup struct {

	// insertion point for fields

	// generated from element "title" of type titleType order 3 depth 1
	Title string `xml:"title,omitempty"`

	// generated from element "author" of type string order 4 depth 1
	Author string `xml:"author,omitempty"`

	// generated from element "year" of type yearType order 5 depth 1
	Year int `xml:"year,omitempty"`

	// generated from element "format" of type bookFormatEnum order 6 depth 1
	Format string `xml:"format,omitempty"`
}

// AttributeGroup_commonAttributes is generated from named attribute group "commonAttributes"
type AttributeGroup_commonAttributes struct {

	// insertion point for fields

	// generated from attribute "isbn
	Isbn string `xml:"isbn,attr,omitempty"`

	// generated from attribute "bestseller
	Bestseller bool `xml:"bestseller,attr,omitempty"`
}

// AttributeGroup_extendedAttributes is generated from named attribute group "extendedAttributes"
type AttributeGroup_extendedAttributes struct {

	// insertion point for fields

	// generated from attribute "edition
	Edition string `xml:"edition,attr,omitempty"`

	// generated from attribute group "commonAttributes
	AttributeGroup_commonAttributes
}

// Books is generated from element books within root schema
type Books struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"books"`

	// generated from inline complex type
	A_books
}
