// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type ComplexType_WOP struct {
	// insertion point
	Name string
	NameXSD string
}

func (from *ComplexType) CopyBasicFields(to *ComplexType) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
}

type Element_WOP struct {
	// insertion point
	Name string
	NameXSD string
	Type string
}

func (from *Element) CopyBasicFields(to *Element) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.Type = from.Type
}

type Enumeration_WOP struct {
	// insertion point
	Name string
	Value string
}

func (from *Enumeration) CopyBasicFields(to *Enumeration) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type Restriction_WOP struct {
	// insertion point
	Name string
	Base string
}

func (from *Restriction) CopyBasicFields(to *Restriction) {
	// insertion point
	to.Name = from.Name
	to.Base = from.Base
}

type Schema_WOP struct {
	// insertion point
	Name string
}

func (from *Schema) CopyBasicFields(to *Schema) {
	// insertion point
	to.Name = from.Name
}

type Sequence_WOP struct {
	// insertion point
	Name string
}

func (from *Sequence) CopyBasicFields(to *Sequence) {
	// insertion point
	to.Name = from.Name
}

type SimpleType_WOP struct {
	// insertion point
	Name string
	NameXSD string
}

func (from *SimpleType) CopyBasicFields(to *SimpleType) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
}

