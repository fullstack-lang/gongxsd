// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Annotation_WOP struct {
	// insertion point
	Name string
}

func (from *Annotation) CopyBasicFields(to *Annotation) {
	// insertion point
	to.Name = from.Name
}

type Documentation_WOP struct {
	// insertion point
	Name string
	Text string
	Source string
	Lang string
}

func (from *Documentation) CopyBasicFields(to *Documentation) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
	to.Source = from.Source
	to.Lang = from.Lang
}

type Schema_WOP struct {
	// insertion point
	Name string
	Xs string
}

func (from *Schema) CopyBasicFields(to *Schema) {
	// insertion point
	to.Name = from.Name
	to.Xs = from.Xs
}

