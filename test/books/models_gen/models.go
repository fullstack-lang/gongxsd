// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// BookType is generated from named complex type "bookType"
type BookType struct {
	Name string `xml:"-"`
	
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

// Books is generated from inlined complex type within element "books"
type Books struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "book" of type bookType
	Book []*BookType `xml:"book"`
}


