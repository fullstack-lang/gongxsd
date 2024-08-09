// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// BookType is generated from named complex type "bookType"
type BookType struct {
	Name string
	// insertion point for fields

	// generated from element title with type simple type titleType
	Title string `xml:"title"`

	// generated from element author with type xs:string
	Author string `xml:"author"`

	// generated from element year with type simple type yearType
	Year int `xml:"year"`

	// generated from element format with type simple type bookFormatEnum
	Format string `xml:"format"`
}

// Books is generated from inlined complex type within element "books"
type Books struct {
	Name string
	// insertion point for fields
}


