// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type BookDetailsGroup_WOP struct {
	// insertion point
	Name string
	Title string
	Author string
	Year int
	Format string
}

func (from *BookDetailsGroup) CopyBasicFields(to *BookDetailsGroup) {
	// insertion point
	to.Name = from.Name
	to.Title = from.Title
	to.Author = from.Author
	to.Year = from.Year
	to.Format = from.Format
}

type BookType_WOP struct {
	// insertion point
	Name string
	Title string
	Author string
	Year int
	Format string
}

func (from *BookType) CopyBasicFields(to *BookType) {
	// insertion point
	to.Name = from.Name
	to.Title = from.Title
	to.Author = from.Author
	to.Year = from.Year
	to.Format = from.Format
}

type Books_WOP struct {
	// insertion point
	Name string
}

func (from *Books) CopyBasicFields(to *Books) {
	// insertion point
	to.Name = from.Name
}

type CommonAttributes_WOP struct {
	// insertion point
	Name string
	Isbn string
	Bestseller bool
}

func (from *CommonAttributes) CopyBasicFields(to *CommonAttributes) {
	// insertion point
	to.Name = from.Name
	to.Isbn = from.Isbn
	to.Bestseller = from.Bestseller
}

type Credit_WOP struct {
	// insertion point
	Name string
	Page int
	Credit_type string
	Credit_words string
	Credit_symbol string
}

func (from *Credit) CopyBasicFields(to *Credit) {
	// insertion point
	to.Name = from.Name
	to.Page = from.Page
	to.Credit_type = from.Credit_type
	to.Credit_words = from.Credit_words
	to.Credit_symbol = from.Credit_symbol
}

type ExtendedAttributes_WOP struct {
	// insertion point
	Name string
	Edition string
}

func (from *ExtendedAttributes) CopyBasicFields(to *ExtendedAttributes) {
	// insertion point
	to.Name = from.Name
	to.Edition = from.Edition
}

type Link_WOP struct {
	// insertion point
	Name string
	NameXSD string
	EnclosedText string
}

func (from *Link) CopyBasicFields(to *Link) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.EnclosedText = from.EnclosedText
}

